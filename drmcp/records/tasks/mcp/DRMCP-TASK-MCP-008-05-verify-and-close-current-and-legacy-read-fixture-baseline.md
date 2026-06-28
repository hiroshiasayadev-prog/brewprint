# DRMCP-TASK-MCP-008-05: Verify and close current and legacy read fixture baseline

- **id**: DRMCP-TASK-MCP-008-05
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-008
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-008-04
- **outputs**:
  - DRMCP-WORK-MCP-008
  - DRMCP-TASK-MCP-001-08
  - drmcp/src/internal/designrecords/testdata/read-baseline/manifest.json

## Goal

Verify the complete current, configured legacy, rejection, and isolation fixture baseline.

Confirm structural completeness and non-leakage boundaries before closing W008 and hub T08.

## Work

- Confirm that the accepted T04 fixture baseline is stored in the direct predecessor commit.
- Record a bounded pre-creation inventory for this Task.
- Verify the authoritative manifest and fixture tree without changing accepted fixture semantics.
- Confirm current, legacy, rejection, and arrangement separation.
- Confirm duplicate, overlap, disabled, unresolved, unsupported, leakage, path-exposure, and Topics graph boundaries.
- Record a final structural verifier that compares current fixture bytes with the accepted T04 commit and independently checks the full manifest.
- Run the verifier when repository-local Python execution is available.
- Inspect only the T05 changed-file boundary for Git status and whitespace.
- Obtain an independent review before changing this Task, W008, or hub T08 to `done`.
- Synchronize final closure evidence only after review acceptance.

This Task does not define new fixture semantics or runtime behavior.
It does not modify production implementation, existing Go tests, or accepted authority records.

## Done condition

- T01 through T04 remain `done`.
- The manifest contains exactly C01 through C17, L01 through L13, and R01 through R24 in deterministic order.
- All 54 case IDs are unique.
- Every case records an expected classification, a non-empty runtime-owner list, and a planned Task.
- T04-owned cases retain `DRMCP-TASK-MCP-008-04` as their planned Task.
- Every declared existing path exists and every intentional absent path remains absent.
- Every fixture path is fixture-root-relative and uses `/` separators.
- Normal current and legacy roots remain unchanged and disjoint.
- Arrangement roots are not registered as normal roots.
- The normal legacy root contains only ADR, INV, REQ, WORK, and TASK families.
- `V01-SPEC-901` exists only in the explicit rejection arrangement.
- Duplicate identity, duplicate root, and duplicate Topics parent cases list every source and select no winner.
- Equal, ancestor, and descendant overlap declarations match their physical paths.
- Omitted and empty legacy-root configurations remain separate disabled cases.
- Disabled, unresolved, unsupported, duplicate, and rejected classifications remain distinct.
- L10 through L13 retain the required false isolation flags.
- R17 through R20 and R24 remain forbidden-leakage expectations rather than observed runtime results.
- C15 diagnostic source location remains separate from C16 and R20 normal response path hiding.
- R21 through R23 retain complete unresolved-child, duplicate-parent, and cycle arrangements.
- Current fixture bytes match the accepted T04 commit unless a reviewed correction is required.
- Production `.go` implementation, existing Go tests, and accepted authority records remain unchanged.
- The final structural verifier passes, or an accurate execution limitation is recorded.
- Independent review reports `PASS` with no blocking or major finding.
- This Task, W008, and hub T08 are synchronized to `done` only after review acceptance.

## Verification

- Read the authoritative manifest and bounded fixture tree through filesystem MCP.
- Compare the full fixture root with accepted T04 commit `57998d6b51ece4a6845cbbf322688fc231443dfa`.
- Run the ready-to-run Python verifier recorded in Evidence when repository-local execution is available.
- Run `git.inspect_worktree` only for this Task, W008, and hub T08.
- Run an independent review against the recorded prompt.
- Do not infer repository-wide cleanliness from scoped checks.

## Evidence

### Accepted T04 commit baseline

The direct predecessor commit was confirmed through bounded Git metadata inspection:

- commit: `57998d6b51ece4a6845cbbf322688fc231443dfa`;
- subject: `test(drmcp): add rejected and isolation read fixture baseline`;
- branch ref: `refs/heads/main`.

`git.inspect_worktree` was run for the exact 22-file T04 boundary.
The result was:

- `result: pass`;
- `changes_present: false`;
- `scope_clean: true`;
- `whitespace.status: pass`;
- no whitespace findings;
- `repository_wide_clean: null`.

A second inspection scoped to `drmcp/src/internal/designrecords/testdata/read-baseline/` also returned `scope_clean: true` and no whitespace finding.
The accepted T04 fixture tree therefore matches the committed baseline before T05 authoring.
No repository-wide clean state is inferred.

### Bounded pre-creation inventory

The exact directory `drmcp/records/tasks/mcp/` was searched immediately before creation.

The following bounded searches returned no match:

- public-ID and filename prefix `DRMCP-TASK-MCP-008-05*`;
- exact filename `DRMCP-TASK-MCP-008-05-verify-and-close-current-and-legacy-read-fixture-baseline.md`.

No Task with the same ID or filename existed.
No existing Task was overwritten.

The selected exact ID and file are:

- `DRMCP-TASK-MCP-008-05`;
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-05-verify-and-close-current-and-legacy-read-fixture-baseline.md`.

### Final structural assessment

The authoritative manifest contains 54 unique cases in this exact order:

- C01 through C17;
- L01 through L13;
- R01 through R24.

Every case contains:

- `expected_classification`;
- a non-empty `runtime_owner` list;
- `planned_task`.

Planned-Task ownership remains:

- C01 through C14 and C17: `DRMCP-TASK-MCP-008-02`;
- L01 through L09: `DRMCP-TASK-MCP-008-03`;
- C15, C16, L10 through L13, and R01 through R24: `DRMCP-TASK-MCP-008-04`.

All declared existing paths were observed in the bounded fixture tree.
The five declared intentional absent paths remain absent:

- missing current root;
- missing legacy root;
- unresolved current requirement source;
- unresolved accepted legacy requirement source;
- unresolved Topics child source.

Every manifest path is fixture-root-relative, uses `/`, and contains no parent traversal.

### Root and arrangement separation

The unchanged normal roots are:

- `current/product/records`;
- `current/drmcp/records`;
- `legacy/v01/records`.

The current roots and legacy root are distinct sibling trees.
No normal root is equal to, above, or below another normal root from the opposite scope.

All rejection, invalid-source, duplicate, invalid-root, overlap, and invalid Topics graph material remains below `arrangements/`.
Every arrangement has role `arrangement_only`.
Arrangement-local roots use `arrangement_only` or `rejection_only` and are not normal roots.

The normal legacy root contains exactly one source from each accepted family:

- ADR;
- INV;
- REQ;
- WORK;
- TASK.

No `V01-SPEC-*` source exists below the normal legacy root.
`V01-SPEC-901` exists only below `arrangements/rejected-legacy/`.

### Duplicate, overlap, and fallback-state separation

R07 lists both duplicate current sources with canonical ID `DRMCP-REQ-MCP-990` and records `winner: null`.

R10 preserves declaration positions `0` and `1` for the duplicate current root and records `winner: null`.

R22 lists both authoritative Topics parent sources and records `winner: null`.

R11 records these physical relationships in deterministic order:

- equal current and legacy roots;
- current root as ancestor of the legacy root;
- current root as descendant of the legacy root.

The declared relationships match the fixture directory structure.

R12 records omitted `legacy_roots` as disabled.
R13 records an explicit empty `legacy_roots` list as disabled.
R14 records an unresolved valid current input.
R15 records an unresolved accepted legacy input with configured roots.
R16 records the unsupported `V01-SPEC-*` family as rejected.
These cases remain distinct from duplicate and general rejection cases.

### Isolation, leakage, and path exposure

L10 through L13 retain these exact false flags:

- `normal_listing_candidate: false`;
- `current_validation_subject: false`;
- `authoring_target: false`;
- `active_index_candidate: false`.

R17 through R20 and R24 each record `must_not_occur: true`.
Their notes state forbidden observations rather than runtime leakage results.

C15 records a fixture-root-relative diagnostic source location for an intentionally invalid current source.
C16 records physical-path fields forbidden on normal response surfaces.
R20 records the same path fields as forbidden normal listing or retrieval leakage.
The diagnostic exception and normal-response hiding boundary remain separate.

### Topics graph assessment

C17 contains a valid Topics row with `title`, `kind`, `ref`, and `summary`.
The child source has the matching parent ref.

R21 contains one parent row whose exact canonical child path is intentionally absent.

R22 contains two authoritative parents for one existing child source and selects no winner.

R23 contains two Index sources with the complete closed edge sequence:

- cycle A to cycle B;
- cycle B to cycle A.

The fixture records graph source arrangements only.
It does not claim graph-validator execution or diagnostic output.

### Runtime and authority boundary

The manifest and fixture files describe source placement, expected classification, and forbidden states.
They do not claim production parser, index, resolver, validator, MCP, authoring, or diagnostic execution results.

No production `.go` implementation or existing Go test is part of the T05 changed-file boundary.
No accepted ADR, Requirement, or Specification is part of the T05 changed-file boundary.

### Ready-to-run final structural verifier

Run from the repository root after closure synchronization:

```powershell
@'
import json
import subprocess
from pathlib import Path, PurePosixPath

accepted_t04_commit = "57998d6b51ece4a6845cbbf322688fc231443dfa"
accepted_t04_subject = "test(drmcp): add rejected and isolation read fixture baseline"
fixture_rel = "drmcp/src/internal/designrecords/testdata/read-baseline"
manifest_rel = f"{fixture_rel}/manifest.json"
fixture = Path(fixture_rel)

resolved_commit = subprocess.check_output(
    ["git", "rev-parse", accepted_t04_commit],
    text=True,
    encoding="utf-8",
).strip()
assert resolved_commit == accepted_t04_commit
subject = subprocess.check_output(
    ["git", "show", "-s", "--format=%s", accepted_t04_commit],
    text=True,
    encoding="utf-8",
).strip()
assert subject == accepted_t04_subject

manifest_bytes = (fixture / "manifest.json").read_bytes()
committed_manifest_bytes = subprocess.check_output(
    ["git", "show", f"{accepted_t04_commit}:{manifest_rel}"]
)
assert manifest_bytes == committed_manifest_bytes
assert subprocess.run(
    ["git", "diff", "--quiet", accepted_t04_commit, "--", fixture_rel],
    check=False,
).returncode == 0
assert subprocess.check_output(
    ["git", "status", "--porcelain=v1", "--untracked-files=all", "--", fixture_rel],
    text=True,
    encoding="utf-8",
).strip() == ""

manifest = json.loads(manifest_bytes.decode("utf-8"))
expected_ids = [
    *(f"C{i:02d}" for i in range(1, 18)),
    *(f"L{i:02d}" for i in range(1, 14)),
    *(f"R{i:02d}" for i in range(1, 25)),
]
ids = [case["id"] for case in manifest["cases"]]
assert ids == expected_ids
assert len(ids) == 54
assert len(ids) == len(set(ids))
cases = {case["id"]: case for case in manifest["cases"]}

for case in manifest["cases"]:
    assert case["expected_classification"] in {
        "accepted", "rejected", "unresolved", "disabled", "duplicate"
    }
    assert isinstance(case["runtime_owner"], list) and case["runtime_owner"]
    assert case["planned_task"]
    for key in ("repository_paths", "intentional_absent_paths"):
        for rel in case.get(key, []):
            assert "\\" not in rel
            path = PurePosixPath(rel)
            assert not path.is_absolute()
            assert ".." not in path.parts
    for rel in case.get("repository_paths", []):
        assert (fixture / rel).exists(), (case["id"], rel)
    for rel in case.get("intentional_absent_paths", []):
        assert not (fixture / rel).exists(), (case["id"], rel)

for case_id in [*(f"C{i:02d}" for i in range(1, 15)), "C17"]:
    assert cases[case_id]["planned_task"] == "DRMCP-TASK-MCP-008-02"
for case_id in [f"L{i:02d}" for i in range(1, 10)]:
    assert cases[case_id]["planned_task"] == "DRMCP-TASK-MCP-008-03"
for case_id in ["C15", "C16", *(f"L{i:02d}" for i in range(10, 14)), *(f"R{i:02d}" for i in range(1, 25))]:
    assert cases[case_id]["planned_task"] == "DRMCP-TASK-MCP-008-04"

for item in [*manifest["configurations"], *manifest["arrangements"]]:
    rel = item["repository_path"]
    assert "\\" not in rel
    assert not PurePosixPath(rel).is_absolute()
    assert (fixture / rel).exists(), rel

normal_current = [
    fixture / root["repository_path"]
    for root in manifest["roots"]["current"]
    if root["role"] == "normal"
]
normal_legacy = [
    fixture / root["repository_path"]
    for root in manifest["roots"]["legacy"]
    if root["role"] == "normal"
]
assert [path.as_posix() for path in normal_current] == [
    f"{fixture_rel}/current/product/records",
    f"{fixture_rel}/current/drmcp/records",
]
assert [path.as_posix() for path in normal_legacy] == [
    f"{fixture_rel}/legacy/v01/records"
]
for current in normal_current:
    for legacy in normal_legacy:
        assert current != legacy
        assert current not in legacy.parents
        assert legacy not in current.parents

normal_paths = set(normal_current + normal_legacy)
arrangement_root_names = set()
for arrangement in manifest["arrangements"]:
    assert arrangement["role"] == "arrangement_only"
    assert fixture / arrangement["repository_path"] not in normal_paths
    for group in ("current_roots", "legacy_roots"):
        for root in arrangement.get(group, []):
            assert root["role"] in {"arrangement_only", "rejection_only"}
            root_path = fixture / root["repository_path"]
            assert root_path.exists()
            assert root_path not in normal_paths
            assert root["name"] not in arrangement_root_names
            arrangement_root_names.add(root["name"])

legacy_root = fixture / "legacy/v01/records"
legacy_files = sorted(path.name for path in legacy_root.rglob("V01-*.md"))
assert legacy_files == sorted([
    "V01-ADR-901-configured-legacy-fixture.md",
    "V01-INV-MCP-901-configured-legacy-fixture.md",
    "V01-REQ-MCP-901-configured-legacy-fixture.md",
    "V01-TASK-MCP-901-01-configured-legacy-fixture.md",
    "V01-WORK-MCP-901-configured-legacy-fixture.md",
])
assert not any(name.startswith("V01-SPEC-") for name in legacy_files)
rejected_spec = fixture / "arrangements/rejected-legacy/legacy/v01/records/spec/V01-SPEC-901-rejected-legacy-spec.md"
assert rejected_spec.exists()
assert cases["R01"]["repository_paths"] == [rejected_spec.relative_to(fixture).as_posix()]
assert cases["R01"]["accepted_legacy_source"] is False

assert cases["R07"]["conflicting_sources"] == cases["R07"]["repository_paths"]
assert cases["R07"]["winner"] is None
assert cases["R10"]["declaration_positions"] == [0, 1]
assert cases["R10"]["winner"] is None
assert cases["R22"]["topics_graph"]["parent_sources"] == cases["R22"]["repository_paths"][:2]
assert cases["R22"]["topics_graph"]["winner"] is None

relations = cases["R11"]["overlap_relations"]
assert [relation["relationship"] for relation in relations] == [
    "equal",
    "current_ancestor_of_legacy",
    "current_descendant_of_legacy",
]
for relation in relations:
    current = fixture / relation["current_root"]
    legacy = fixture / relation["legacy_root"]
    if relation["relationship"] == "equal":
        assert current == legacy
    elif relation["relationship"] == "current_ancestor_of_legacy":
        assert current in legacy.parents
    else:
        assert legacy in current.parents

expected_classification = {
    "R07": "duplicate",
    "R10": "duplicate",
    "R12": "disabled",
    "R13": "disabled",
    "R14": "unresolved",
    "R15": "unresolved",
    "R16": "rejected",
    "R21": "unresolved",
    "R22": "duplicate",
}
for case_id, classification in expected_classification.items():
    assert cases[case_id]["expected_classification"] == classification
assert cases["R12"]["legacy_roots_state"] == "omitted"
assert cases["R13"]["legacy_roots_state"] == "empty"
assert cases["R16"]["approved_legacy_families"] == ["ADR", "INV", "REQ", "WORK", "TASK"]

assert cases["L10"]["isolation_flags"] == {"normal_listing_candidate": False}
assert cases["L11"]["isolation_flags"] == {"current_validation_subject": False}
assert cases["L12"]["isolation_flags"] == {"authoring_target": False}
assert cases["L13"]["isolation_flags"] == {"active_index_candidate": False}
for case_id in ["R17", "R18", "R19", "R20", "R24"]:
    assert cases[case_id]["forbidden_leakage"]["must_not_occur"] is True

assert cases["C15"]["source_location"] == cases["C15"]["repository_paths"][0]
assert cases["C16"]["forbidden_normal_response_fields"] == [
    "physical_path", "source_location", "index_path"
]
assert cases["R20"]["forbidden_leakage"]["fields"] == [
    "physical_path", "source_location", "index_path"
]

assert cases["R21"]["topics_graph"]["child_ref"] == "spec:product.fixture_invalid.topics_unresolved.missing_child"
assert len(cases["R22"]["topics_graph"]["parent_sources"]) == 2
assert cases["R23"]["topics_graph"]["closed_edge_sequence"] == [
    ["spec:product.fixture_invalid.cycle_a", "spec:product.fixture_invalid.cycle_b"],
    ["spec:product.fixture_invalid.cycle_b", "spec:product.fixture_invalid.cycle_a"],
]

records = {
    "t01": Path("drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-01-build-current-and-legacy-fixture-coverage-matrix.md"),
    "t02": Path("drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-02-create-current-root-read-fixture-baseline.md"),
    "t03": Path("drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-03-create-configured-legacy-root-read-fixture-baseline.md"),
    "t04": Path("drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-04-create-rejected-and-isolation-read-fixture-baseline.md"),
    "t05": Path("drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-05-verify-and-close-current-and-legacy-read-fixture-baseline.md"),
    "work": Path("drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md"),
    "hub": Path("drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md"),
}
text = {name: path.read_text(encoding="utf-8") for name, path in records.items()}
for name in ["t01", "t02", "t03", "t04"]:
    assert "- **status**: done" in text[name]
assert "- **status**: done" in text["t05"]
assert "- **status**: done" in text["work"]
assert "- **status**: done" in text["hub"]
assert text["work"].index("  - DRMCP-TASK-MCP-008-01") < text["work"].index("  - DRMCP-TASK-MCP-008-02") < text["work"].index("  - DRMCP-TASK-MCP-008-03") < text["work"].index("  - DRMCP-TASK-MCP-008-04") < text["work"].index("  - DRMCP-TASK-MCP-008-05")
for heading in ["## Goal", "## Work", "## Done condition", "## Verification", "## Evidence"]:
    assert text["t05"].splitlines().count(heading) == 1

source_scope = [
    "drmcp/src/internal/designrecords",
    "drmcp/src/internal/designrecordsmcp",
]
assert subprocess.run(
    ["git", "diff", "--quiet", accepted_t04_commit, "--", *source_scope],
    check=False,
).returncode == 0
assert subprocess.check_output(
    ["git", "status", "--porcelain=v1", "--untracked-files=all", "--", *source_scope],
    text=True,
    encoding="utf-8",
).strip() == ""

authority_scope = [
    "drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md",
    "drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-design-records-mcp-contract-baseline-and-realignment.md",
    "product/records/spec/brewprint/compatibility/legacy-id-compatibility.md",
    "product/records/spec/design-records/spec-format/document-shape.md",
    "product/records/spec/design-records/spec-format/topics-table.md",
]
assert subprocess.run(
    ["git", "diff", "--quiet", accepted_t04_commit, "--", *authority_scope],
    check=False,
).returncode == 0
assert subprocess.check_output(
    ["git", "status", "--porcelain=v1", "--untracked-files=all", "--", *authority_scope],
    text=True,
    encoding="utf-8",
).strip() == ""

print("fixture_shape=OK")
print("fixture_bytes=OK")
print("lifecycle_shape=OK")
print("exclusion_shape=OK")
'@ | python -
```

Expected result:

```text
fixture_shape=OK
fixture_bytes=OK
lifecycle_shape=OK
exclusion_shape=OK
```

The verifier uses the explicitly identified T04 commit.
It does not assume that `HEAD:` contains the T03 baseline.
It compares the full current fixture tree with the accepted T04 commit and independently verifies the 54-case manifest.

### Verifier execution state

The final structural verifier is `NOT RUN` in this session.
The available tools do not provide arbitrary repository-local Python execution.
No execution result is inferred.

Filesystem static inspection and scoped `git.inspect_worktree` results are recorded separately from the unexecuted verifier.

### T05 changed-file boundary

| action | file |
|---|---|
| new | `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-05-verify-and-close-current-and-legacy-read-fixture-baseline.md` |
| modify | `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md` |
| modify | `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md` |

No fixture correction was required.
The manifest and fixture files remain outside the T05 changed-file boundary.

### Scoped Git and whitespace result

`git.inspect_worktree` was run for exactly the three T05 changed files.

The result was:

- `result: pass`;
- two modified tracked workflow files;
- one untracked T05 Task;
- `whitespace.status: pass`;
- no whitespace findings;
- the untracked no-index check returned expected exit code `1`;
- no exit code `2` or greater;
- LF-to-CRLF conversion warnings only, treated as non-blocking;
- `repository_wide_clean: null`.

Recording this result changes the T05 Task bytes.
A final post-evidence `git.inspect_worktree` check must therefore be supplied directly to independent review without writing its result back into a checked file.

### Lifecycle before independent review

- T01: `done`;
- T02: `done`;
- T03: `done`;
- T04: `done`;
- T05: `in_progress`;
- W008: `in_progress`;
- hub T08: `in_progress`.

T05, W008, and hub T08 must remain `in_progress` until independent review acceptance.

### Independent review prompt

```text
C:\Users\imved\projects\brewprint

`DRMCP-TASK-MCP-008-05-verify-and-close-current-and-legacy-read-fixture-baseline`の独立reviewを行う。

ファイルは変更しないこと。
DRMCPは現在利用できない。
Design Recordおよびfixtureの読み込みにはfilesystem MCPを使用すること。
Gitの限定status・whitespace確認には`git.inspect_worktree`を使用すること。
sandboxへrepositoryを複製しないこと。
repository-wide traversalを行わないこと。
repository-wide clean statusを確認・推測しないこと。
repository-local commandを実行できない場合、実行したと捏造しないこと。
production `.go` implementationと既存Go testを変更しないこと。
accepted ADR、Requirement、Specificationを変更しないこと。

## 最初に読む

1. `prompt_chappy.md`
2. `product/records/spec/design-records/authoring-standards/task-authoring.md`
3. `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
4. `product/records/spec/design-records/authoring-standards/writing-standard.md`
5. `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`

## Primary workflow records

- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-05-verify-and-close-current-and-legacy-read-fixture-baseline.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-01-build-current-and-legacy-fixture-coverage-matrix.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-02-create-current-root-read-fixture-baseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-03-create-configured-legacy-root-read-fixture-baseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-04-create-rejected-and-isolation-read-fixture-baseline.md`

## Authority

- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md`
- `drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-design-records-mcp-contract-baseline-and-realignment.md`
- `product/records/spec/brewprint/compatibility/legacy-id-compatibility.md`
- `product/records/spec/design-records/spec-format/document-shape.md`
- `product/records/spec/design-records/spec-format/topics-table.md`

## Fixture

- `drmcp/src/internal/designrecords/testdata/read-baseline/manifest.json`
- `drmcp/src/internal/designrecords/testdata/read-baseline/config/`
- `drmcp/src/internal/designrecords/testdata/read-baseline/current/`
- `drmcp/src/internal/designrecords/testdata/read-baseline/legacy/`
- `drmcp/src/internal/designrecords/testdata/read-baseline/arrangements/`

## Accepted T04 commit

Use this explicit baseline:

- commit: `57998d6b51ece4a6845cbbf322688fc231443dfa`;
- subject: `test(drmcp): add rejected and isolation read fixture baseline`.

Do not assume that `HEAD:` is the T03 baseline.
Confirm that current fixture bytes have not changed from the accepted T04 commit.

## Required review

Confirm at minimum:

- T01 through T04 are `done`;
- T05, W008, and hub T08 are `in_progress` before review;
- W008 lists T05 after T04;
- the manifest contains exactly 54 unique cases in order C01-C17, L01-L13, R01-R24;
- every case has expected classification, non-empty runtime owner, and planned Task;
- T04 cases retain planned Task T04;
- current fixture bytes are unchanged from the accepted T04 commit;
- all declared existing paths exist and intentional absent paths do not exist;
- every path is fixture-root-relative and `/`-separated;
- normal current and legacy roots remain unchanged and disjoint;
- arrangement roots are not registered as normal roots;
- the normal legacy family set is exactly ADR, INV, REQ, WORK, and TASK;
- no positive `V01-SPEC-*` source exists;
- rejected `V01-SPEC-901` exists only in the rejection arrangement;
- duplicate current identity, duplicate root, and duplicate Topics parent enumerate all sources and select no winner;
- equal, ancestor, and descendant overlap declarations match physical paths;
- omitted and empty legacy-root states remain separate;
- disabled, unresolved, unsupported, duplicate, and rejected cases remain distinct;
- L10 through L13 contain the required false isolation flags;
- R17 through R20 and R24 express forbidden leakage, not observed leakage;
- C15 diagnostic source location remains separate from C16 and R20 normal response path hiding;
- C17 and R21 through R23 preserve the accepted Topics graph boundaries;
- fixture material does not claim production runtime results;
- the final verifier result is recorded accurately as executed output or `NOT RUN`;
- scoped Git status and whitespace cover only the three T05 changed files;
- production `.go`, existing Go tests, and accepted authority files are unchanged;
- repository-wide cleanliness is not inferred;
- T05, W008, and hub T08 are ready for synchronized closure only after review acceptance.

When repository-local execution is available, run the final structural verifier recorded in T05 Evidence.
Run `git.inspect_worktree` only for the T05 changed-file boundary.
Do not write review results into the checked files.

出力形式:

1. Verdict: PASS / NEEDS REVISION
2. Previous-finding disposition
3. Blocking findings
4. Major findings
5. Minor findings
6. Advisories
7. Accepted T04 commit and fixture-byte assessment
8. Manifest completeness and ordering assessment
9. Root, arrangement, and legacy-family separation assessment
10. Duplicate, overlap, and fallback-state assessment
11. Isolation, leakage, and path-exposure assessment
12. Topics graph assessment
13. Runtime and authority exclusion assessment
14. Verifier evidence assessment
15. Lifecycle and changed-file assessment
16. T05 closure readiness
17. W008 and hub T08 closure readiness
```

### Independent review acceptance

Independent review returned `PASS` on 2026-06-28.

- T05 previous findings: none;
- T02 F-MAJ-01, F-MAJ-02, and F-MIN-01 remain `CLOSED` with no regression;
- blocking findings: none;
- major findings: none;
- minor findings: none;
- accepted T04 commit and fixture-byte assessment: `PASS` with the recorded execution limitation;
- manifest completeness and ordering: `PASS`;
- root, arrangement, and legacy-family separation: `PASS`;
- duplicate, overlap, and fallback-state separation: `PASS`;
- isolation, leakage, and path-exposure boundaries: `PASS`;
- Topics graph assessment: `PASS`;
- runtime and authority exclusions: `PASS`;
- verifier evidence: `PASS` with `NOT RUN` retained accurately;
- lifecycle and changed-file assessment: `PASS`.

The reviewer confirmed commit `57998d6b51ece4a6845cbbf322688fc231443dfa` with subject `test(drmcp): add rejected and isolation read fixture baseline` as the accepted T04 baseline.
The reviewer did not rerun the commit-to-worktree byte comparison or the repository-local Python verifier and did not infer either result.

The review-time scoped `git.inspect_worktree` check covered only this Task, W008, and hub T08.
It returned `result: pass`, no whitespace finding, LF-to-CRLF conversion warnings only, and `repository_wide_clean: null`.

### Closure synchronization

Independent review accepted synchronized closure in this order:

1. `DRMCP-TASK-MCP-008-05` to `done`;
2. `DRMCP-WORK-MCP-008` to `done`;
3. `DRMCP-TASK-MCP-001-08` to `done`.

No fixture correction was required.
The final structural verifier remains `NOT RUN` because arbitrary repository-local Python execution is unavailable.
No production `.go` implementation, existing Go test, accepted ADR, Requirement, or Specification was changed.
Repository-wide cleanliness remains unasserted.

### Final closure state

- T01: `done`;
- T02: `done`;
- T03: `done`;
- T04: `done`;
- T05: `done`;
- W008: `done`;
- hub T08: `done`.

T05 closure readiness: complete.
W008 closure readiness: complete.
Hub T08 closure readiness: complete.
