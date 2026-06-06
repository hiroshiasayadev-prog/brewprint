"""REQ-MCP-027 runtime smoke — accurate update diffs and no-op detection.

Uses a temp root (copy of docs/) so the real repository is never modified.
Smoke covers:
  [1] Real metadata update  →  bounded git-style modify diff
  [2] No-op metadata update →  proposal_created:false, no_op_update diagnostic
  [3] Accept real update in temp root  →  written:true
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

            # ── discard remaining proposals ──────────────────────────────────
            if proposal_ids_to_discard:
                print(f"\n== discard remaining proposals ==")
                for i, pid in enumerate(proposal_ids_to_discard, start=5):
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
