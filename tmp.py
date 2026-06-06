"""REQ-MCP-027 / REQ-MCP-024 / REQ-MCP-028 runtime smoke.

Uses a temp root (copy of docs/) so the real repository is never modified.
Smoke covers:
  [1] Real metadata update  →  bounded git-style modify diff
  [2] No-op metadata update →  proposal_created:false, no_op_update diagnostic
  [3] Accept real update in temp root  →  written:true
  [4] REQ-MCP-028: propose_record_create with missing required fields
      → proposal_created:false, missing_required_metadata_batch diagnostic
  [5] REQ-MCP-024: propose_record_create with invalid status
      → proposal_created:false, invalid_metadata_value with allowed_values
"""

import json
import os
import shutil
import subprocess
import sys
import tempfile
from pathlib import Path

ROOT = Path.cwd()
# Drive the REPO-LOCAL server so the new binary is always used.
_SERVER_BASE = ["go", "run", "./cmd/design-records-mcp", "--root"]

# Target: TASK-MCP-023-04 (status:done, has Evidence section)
TARGET_ID = "TASK-MCP-023-04"
TARGET_KIND = "task"


# ── JSON-RPC helpers ─────────────────────────────────────────────────────────

def send(proc, obj):
    proc.stdin.write(json.dumps(obj, ensure_ascii=False) + "\n")
    proc.stdin.flush()


def read_response(proc):
    line = proc.stdout.readline()
    if not line:
        raise RuntimeError("server stdout closed. stderr:\n" + proc.stderr.read())
    return json.loads(line)


def call_tool(proc, req_id, name, arguments):
    send(proc, {
        "jsonrpc": "2.0",
        "id": req_id,
        "method": "tools/call",
        "params": {"name": name, "arguments": arguments},
    })
    res = read_response(proc)
    if "error" in res:
        raise AssertionError(
            f"{name} JSON-RPC error:\n{json.dumps(res, ensure_ascii=False, indent=2)}"
        )
    result = res.get("result", {})
    content = result.get("content", [])
    if not content:
        raise AssertionError(
            f"{name} returned no content:\n{json.dumps(res, ensure_ascii=False, indent=2)}"
        )
    text = content[0].get("text", "")
    try:
        payload = json.loads(text)
    except json.JSONDecodeError as e:
        raise AssertionError(f"{name} content is not JSON: {text}") from e
    if result.get("isError") is True:
        raise AssertionError(
            f"{name} returned isError=true:\n{json.dumps(payload, ensure_ascii=False, indent=2)}"
        )
    return payload


# ── Assertion helpers ────────────────────────────────────────────────────────

def assert_bounded_git_modify_diff(payload, expected_fragment, *, max_diff_lines=40):
    """Verify proposal_created:true and a bounded git-style unified modify diff."""
    if not payload.get("proposal_created"):
        raise AssertionError(
            f"proposal was not created:\n{json.dumps(payload, ensure_ascii=False, indent=2)}"
        )

    diff_text = (payload.get("diff") or {}).get("text") or ""

    for header in ("diff --git ", "--- a/", "+++ b/", "@@"):
        if header not in diff_text:
            raise AssertionError(
                f"missing git-style header '{header}' in diff:\n{diff_text}"
            )

    if expected_fragment not in diff_text:
        raise AssertionError(
            f"expected fragment '{expected_fragment}' not found in diff:\n{diff_text}"
        )

    lines = diff_text.splitlines()
    if len(lines) > max_diff_lines:
        raise AssertionError(
            f"diff is unbounded: {len(lines)} lines > max {max_diff_lines}. "
            f"Looks like whole-file addition:\n{diff_text}"
        )

    # A real single-field metadata update should have very few added content lines.
    plus_content = [l for l in lines if l.startswith("+") and not l.startswith("+++")]
    if len(plus_content) > 5:
        raise AssertionError(
            f"too many '+' content lines ({len(plus_content)}) — looks like whole-file addition:\n"
            + "\n".join(l for l in lines if l.startswith("+"))
        )

    return payload["proposal_id"]


def assert_no_op(payload):
    """Verify no retained proposal and a no_op_update info diagnostic."""
    if payload.get("proposal_created"):
        raise AssertionError(
            f"expected no-op but proposal was created:\n"
            f"{json.dumps(payload, ensure_ascii=False, indent=2)}"
        )
    if payload.get("proposal_id"):
        raise AssertionError(
            f"no-op response must not include proposal_id: {payload.get('proposal_id')}"
        )
    if payload.get("diff") is not None:
        raise AssertionError(
            f"no-op response must not include diff: {payload.get('diff')}"
        )
    diagnostics = payload.get("diagnostics") or []
    diag = next((d for d in diagnostics if d.get("category") == "no_op_update"), None)
    if diag is None:
        raise AssertionError(f"missing no_op_update diagnostic: {diagnostics}")
    if diag.get("severity") != "info":
        raise AssertionError(f"no_op_update severity must be 'info': {diag}")
    if not diag.get("record_id"):
        raise AssertionError(f"no_op_update diagnostic must include record_id: {diag}")


# ── Main ─────────────────────────────────────────────────────────────────────

def main():
    # Build a temp root with a copy of docs/ so the real repo is never written.
    temp_root = tempfile.mkdtemp(prefix="brewprint-smoke-req027-")
    print(f"Temp root: {temp_root}")
    try:
        shutil.copytree(str(ROOT / "docs"), os.path.join(temp_root, "docs"))

        proc = subprocess.Popen(
            _SERVER_BASE + [temp_root],
            cwd=str(ROOT),
            stdin=subprocess.PIPE,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            text=True,
            encoding="utf-8",
            errors="replace",
            bufsize=1,
        )

        proposal_ids_to_discard = []

        try:
            # ── initialize ──────────────────────────────────────────────────
            print("\n== initialize ==")
            send(proc, {
                "jsonrpc": "2.0",
                "id": 1,
                "method": "initialize",
                "params": {
                    "protocolVersion": "2024-11-05",
                    "capabilities": {},
                    "clientInfo": {
                        "name": "req-mcp-027-runtime-smoke",
                        "version": "0.1.0",
                    },
                },
            })
            init_res = read_response(proc)
            print(json.dumps(init_res, ensure_ascii=False))

            send(proc, {
                "jsonrpc": "2.0",
                "method": "notifications/initialized",
                "params": {},
            })

            # ── [1] Real metadata update → bounded modify diff ───────────────
            print(f"\n== [1] real metadata update → bounded git-style modify diff ==")
            real_payload = call_tool(proc, 2, "propose_record_update", {
                "id": TARGET_ID,
                "kind": TARGET_KIND,
                "update": {
                    "type": "metadata_fields_replace",
                    "metadata": {"status": "in_progress"},
                },
            })

            diff_text = (real_payload.get("diff") or {}).get("text", "")
            print(json.dumps({
                "proposal_created": real_payload.get("proposal_created"),
                "proposal_id": real_payload.get("proposal_id"),
                "diff_line_count": len(diff_text.splitlines()),
                "diff_snippet": diff_text[:400] if diff_text else None,
            }, ensure_ascii=False, indent=2))

            # Check for "in_progress" in diff (the new status value)
            real_proposal_id = assert_bounded_git_modify_diff(
                real_payload, "in_progress", max_diff_lines=40
            )
            proposal_ids_to_discard.append(real_proposal_id)
            print("[1] PASS")

            # ── [2] No-op metadata update → no retained proposal ─────────────
            print(f"\n== [2] no-op metadata update (status:done → done) → no proposal ==")
            noop_payload = call_tool(proc, 3, "propose_record_update", {
                "id": TARGET_ID,
                "kind": TARGET_KIND,
                "update": {
                    "type": "metadata_fields_replace",
                    "metadata": {"status": "done"},
                },
            })

            print(json.dumps({
                "proposal_created": noop_payload.get("proposal_created"),
                "proposal_id": noop_payload.get("proposal_id"),
                "diagnostics": noop_payload.get("diagnostics"),
                "diff": noop_payload.get("diff"),
            }, ensure_ascii=False, indent=2))

            assert_no_op(noop_payload)
            print("[2] PASS")

            # ── [3] Accept real update in temp root → written:true ───────────
            print(f"\n== [3] accept real update proposal in temp root ==")
            accept_payload = call_tool(proc, 4, "accept_proposed_write", {
                "proposal_id": real_proposal_id,
            })

            print(json.dumps({
                "proposal_id": accept_payload.get("proposal_id"),
                "state": accept_payload.get("state"),
                "written": accept_payload.get("written"),
                "files_written": accept_payload.get("files_written"),
            }, ensure_ascii=False, indent=2))

            if not accept_payload.get("written"):
                raise AssertionError(
                    f"expected written:true:\n"
                    f"{json.dumps(accept_payload, ensure_ascii=False, indent=2)}"
                )
            if accept_payload.get("state") != "accepted":
                raise AssertionError(
                    f"expected state:accepted, got: {accept_payload.get('state')}"
                )
            # Accepted — remove from discard list
            proposal_ids_to_discard.remove(real_proposal_id)
            print("[3] PASS")

            # ── [4] REQ-MCP-028: batch missing-field diagnostic ─────────────
            print("\n== [4] REQ-MCP-028: missing required fields → missing_required_metadata_batch ==")
            batch_payload = call_tool(proc, 5, "propose_record_create", {
                "kind": "work_item",
                "id": "WORK-MCP-new",
                "domain": "MCP",
                "title": "Batch validation smoke",
                # Deliberately omit source_requirement, impact_refs, tasks, status — provide only date
                "fields": {"date": "2026-06-07"},
            })
            print(json.dumps({
                "proposal_created": batch_payload.get("proposal_created"),
                "diagnostics": batch_payload.get("diagnostics"),
            }, ensure_ascii=False, indent=2))

            if batch_payload.get("proposal_created"):
                raise AssertionError("[4] proposal must NOT be created when required fields are absent")
            batch_diags = batch_payload.get("diagnostics") or []
            batch_d = next((d for d in batch_diags if d.get("category") == "missing_required_metadata_batch"), None)
            if batch_d is None:
                raise AssertionError(f"[4] missing missing_required_metadata_batch diagnostic: {batch_diags}")
            if batch_d.get("severity") != "error":
                raise AssertionError(f"[4] expected error severity on batch diagnostic: {batch_d}")
            if not batch_d.get("required_fields"):
                raise AssertionError(f"[4] required_fields must be non-empty: {batch_d}")
            for expected_field in ("status", "source_requirement"):
                if expected_field not in batch_d["required_fields"]:
                    raise AssertionError(f"[4] expected {expected_field!r} in required_fields: {batch_d['required_fields']}")
            print("[4] PASS")

            # ── [5] REQ-MCP-024: invalid status → allowed_values diagnostic ──
            print("\n== [5] REQ-MCP-024: invalid status → invalid_metadata_value with allowed_values ==")
            status_payload = call_tool(proc, 6, "propose_record_create", {
                "kind": "work_item",
                "id": "WORK-MCP-new",
                "domain": "MCP",
                "title": "Status validation smoke",
                "fields": {
                    "status": "implementation_pending",
                    "date": "2026-06-07",
                    "source_requirement": "REQ-MCP-001",
                    "impact_refs": [],
                    "tasks": [],
                },
            })
            print(json.dumps({
                "proposal_created": status_payload.get("proposal_created"),
                "diagnostics": status_payload.get("diagnostics"),
            }, ensure_ascii=False, indent=2))

            if status_payload.get("proposal_created"):
                raise AssertionError("[5] proposal must NOT be created for invalid status")
            status_diags = status_payload.get("diagnostics") or []
            status_d = next((d for d in status_diags if d.get("category") == "invalid_metadata_value"), None)
            if status_d is None:
                raise AssertionError(f"[5] missing invalid_metadata_value diagnostic: {status_diags}")
            if status_d.get("severity") != "error":
                raise AssertionError(f"[5] expected error severity: {status_d}")
            if not status_d.get("allowed_values"):
                raise AssertionError(f"[5] expected non-empty allowed_values: {status_d}")
            if status_d.get("field") != "status":
                raise AssertionError(f"[5] expected field=status: {status_d}")
            if not status_d.get("repair_suggestion"):
                raise AssertionError(f"[5] expected repair_suggestion: {status_d}")
            suggested_status = status_d["repair_suggestion"].get("status")
            if suggested_status not in status_d["allowed_values"]:
                raise AssertionError(
                    f"[5] repair_suggestion.status={suggested_status!r} not in allowed_values {status_d['allowed_values']}"
                )
            print("[5] PASS")

            # ── discard remaining proposals ──────────────────────────────────
            if proposal_ids_to_discard:
                print(f"\n== discard remaining proposals ==")
                for i, pid in enumerate(proposal_ids_to_discard, start=7):
                    d = call_tool(proc, i, "discard_proposed_write", {"proposal_id": pid})
                    print(json.dumps({
                        "proposal_id": pid,
                        "discarded": d.get("discarded"),
                        "state": d.get("state"),
                    }, ensure_ascii=False))

            print("\n== runtime smoke PASS ==")

        finally:
            try:
                proc.stdin.close()
            except Exception:
                pass
            try:
                proc.terminate()
                proc.wait(timeout=5)
            except Exception:
                proc.kill()

            stderr = proc.stderr.read()
            if stderr.strip():
                print("\n== server stderr ==", file=sys.stderr)
                print(stderr, file=sys.stderr)

    finally:
        shutil.rmtree(temp_root, ignore_errors=True)
        print(f"Temp root cleaned up.")


if __name__ == "__main__":
    main()
