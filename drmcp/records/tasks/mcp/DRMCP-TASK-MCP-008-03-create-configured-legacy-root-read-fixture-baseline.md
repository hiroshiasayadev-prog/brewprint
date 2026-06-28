# DRMCP-TASK-MCP-008-03: Create configured legacy-root read fixture baseline

- **id**: DRMCP-TASK-MCP-008-03
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-008
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-008-01
  - DRMCP-TASK-MCP-008-02
- **outputs**:
  - DRMCP-WORK-MCP-008
  - DRMCP-TASK-MCP-001-08
  - drmcp/src/internal/designrecords/testdata/read-baseline/manifest.json
  - drmcp/src/internal/designrecords/testdata/read-baseline/legacy/
  - drmcp/src/internal/designrecords/testdata/read-baseline/config/current-plus-legacy/
  - drmcp/src/internal/designrecords/testdata/read-baseline/current/drmcp/records/requirements/mcp/DRMCP-REQ-MCP-903-configured-legacy-relation-fixture.md

## Goal

Create the configured legacy-root fixture baseline for T01 cases L01 through L09.

Keep approved V01 archive sources separate from current roots and limit fixture-local claims to source arrangement and manifest structure.

## Work

- Confirm the bounded Task, fixture-root, current-root, configuration, and candidate legacy inventories.
- Confirm the accepted V01 family grammar and existing archive placement shape.
- Create one separate normal legacy archive root.
- Create one fixture-local current-plus-legacy configuration declaration.
- Create one readable source for each approved V01 ADR, INV, REQ, WORK, and TASK family.
- Create one T03-owned current requirement with an exact relation to an approved V01 target.
- Extend the authoritative manifest with L01 through L09 after the existing current cases.
- Record exact legacy retrieval and current-first-miss-to-legacy-success arrangement metadata.
- Synchronize W008 and hub T08 without copying fixture details into the hub.
- Keep production implementation and existing Go tests unchanged.

## Done condition

- L01 through L09 are materialized once and in case-ID order.
- The normal legacy root is separate from every normal current root.
- Exactly the approved V01 ADR, INV, REQ, WORK, and TASK families are positive legacy sources.
- Every issued ID is derived from its filename and matches its manifest exact input.
- One exact legacy retrieval input has one readable source.
- One current source contains an exact relation to one approved configured V01 target.
- One fallback arrangement records a current-stage miss and one unique legacy-stage source for the same exact input.
- Existing C01 through C14 and C17 entries remain unchanged.
- No T04 rejection, disabled, unresolved, duplicate, overlap, or leakage fixture is created.
- No production implementation or existing Go test changes.
- Independent review reports no blocking or major finding before this Task changes to `done`.

## Verification

- Parse the authoritative manifest as JSON.
- Confirm existing current case IDs and content remain unchanged.
- Confirm L01 through L09 exist once and in order.
- Resolve every declared existing path from the fixture root.
- Confirm every internal path is fixture-root-relative and uses `/`.
- Confirm normal current and legacy roots are disjoint.
- Confirm the positive legacy family set is exactly ADR, INV, REQ, WORK, and TASK.
- Confirm no positive `V01-SPEC-*` source exists.
- Confirm filename-derived issued IDs match manifest exact inputs.
- Confirm the current relation source and legacy target both exist.
- Confirm L09 records no usable current target and one unique readable legacy source.
- Confirm every L01 through L09 `runtime_owner` is a non-empty list containing `DRMCP-WORK-MCP-010`.
- Confirm every L01 through L09 `planned_task` is `DRMCP-TASK-MCP-008-03`.
- Confirm T01, T02, and this Task are `done` while W008 and hub T08 remain `in_progress`.
- Run scoped Git status and whitespace inspection without asserting repository-wide cleanliness.

## Evidence

### Bounded pre-creation inventory

The exact Task directory search found only:

- `DRMCP-TASK-MCP-008-01-build-current-and-legacy-fixture-coverage-matrix.md`;
- `DRMCP-TASK-MCP-008-02-create-current-root-read-fixture-baseline.md`.

No T03 Task existed before creation.
The selected ID is `DRMCP-TASK-MCP-008-03`.

The fixture root contained only:

- `manifest.json`;
- `current/`;
- `config/`.

The bounded current-root inventory contained `current/product/` and `current/drmcp/`.
The bounded configuration inventory contained `config/current-only/` and `config/legacy-roots-omitted/`.
Candidate `legacy/` and `config/current-plus-legacy/` did not exist.
No existing file was overwritten by T03 creation.

### Accepted legacy root and source placement

The normal legacy root is:

`drmcp/src/internal/designrecords/testdata/read-baseline/legacy/v01/records/`

The accepted lexical mapping comes from `spec:drmcp.design_records_mcp.namespace_scanning`.
The accepted family set comes from `spec:product.brewprint.compatibility.legacy_id_compatibility`.

Existing `v01/records/` placement was inspected before creation.
T03 follows that archive shape:

- ADR below `records/adr/`;
- INV below `records/investigations/<domain>/`;
- REQ below `records/requirements/<domain>/`;
- WORK below `records/work-items/<domain>/`;
- TASK below `records/tasks/<domain>/`.

The created issued IDs are:

- `V01-ADR-901`;
- `V01-INV-MCP-901`;
- `V01-REQ-MCP-901`;
- `V01-WORK-MCP-901`;
- `V01-TASK-MCP-901-01`.

Each filename begins with its complete issued ID and uses a lowercase slug.
The archive files preserve their V01 IDs and are not rewritten into current app-aware IDs.

### Configuration declaration

`config/current-plus-legacy/config.json` is a fixture-local arrangement declaration.
It records two separate current roots, one separate legacy root, and `legacy_roots_state: configured`.
It does not establish production runtime serialization authority.

### L01 through L09 mapping

| case | materialization |
|---|---|
| L01 | One configured normal legacy root at `legacy/v01/records`, separate from both current roots. |
| L02 | Readable `V01-ADR-901` archive source. |
| L03 | Readable `V01-INV-MCP-901` archive source. |
| L04 | Readable `V01-REQ-MCP-901` archive source. |
| L05 | Readable `V01-WORK-MCP-901` archive source. |
| L06 | Readable `V01-TASK-MCP-901-01` archive source. |
| L07 | Exact retrieval input `V01-REQ-MCP-901` with one readable configured legacy source. |
| L08 | Current source `DRMCP-REQ-MCP-903` has exact `source_refs` target `V01-REQ-MCP-901`. |
| L09 | Exact input `V01-REQ-MCP-901` has no usable current target and one unique readable configured legacy source. |

### Current-to-legacy relation arrangement

T03 adds one current requirement source:

`current/drmcp/records/requirements/mcp/DRMCP-REQ-MCP-903-configured-legacy-relation-fixture.md`

Its `source_refs` contains exact target `V01-REQ-MCP-901`.
The legacy source remains a relation target only and is not declared as a current validation subject.
T02-owned current files are unchanged.

### Fallback arrangement boundary

The manifest records fixture-local source facts only:

- exact input `V01-REQ-MCP-901`;
- zero usable current-stage targets;
- one unique readable configured legacy source.

The manifest does not assert resolver execution order, runtime response, target projection, warning, or diagnostic behavior.
Those assertions remain owned by `DRMCP-WORK-MCP-010`.

### Manifest extension

The existing authoritative `manifest.json` is extended after C17 with L01 through L09.
No second manifest is created.

Every legacy case records:

- case ID and fixture class;
- `current+legacy` root arrangement;
- current-root identities and `v01-legacy` legacy-root identity;
- fixture-root-relative repository paths;
- exact inputs;
- configured legacy-root state;
- accepted classification;
- null intentional invalidity;
- `DRMCP-WORK-MCP-010` as primary runtime owner;
- `DRMCP-TASK-MCP-008-03` as planned Task;
- a bounded note.

### Explicit exclusions

T03 does not create:

- L10 through L13;
- R01 through R24;
- `V01-SPEC-*`;
- unsupported V01 families;
- disabled fallback;
- omitted-versus-empty rejection arrangements;
- unresolved legacy targets;
- invalid, duplicate, or overlapping roots;
- duplicate legacy identity;
- listing, validation, authoring, active-index, or path leakage fixtures;
- production implementation changes;
- existing Go test changes.

### Lifecycle synchronization

- Independent review returned `PASS` with no blocking, major, or minor finding.
- This Task is `done` after review acceptance.
- `DRMCP-WORK-MCP-008` remains `in_progress` and lists T01, T02, then T03.
- Hub `DRMCP-TASK-MCP-001-08` remains `in_progress` and records only the T03 lifecycle and evidence pointer.
- T01 and T02 remain `done`.

### Changed-file boundary

Workflow files:

- new: `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-03-create-configured-legacy-root-read-fixture-baseline.md`;
- modified: `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md`;
- modified: `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md`.

Fixture files:

- modified: `drmcp/src/internal/designrecords/testdata/read-baseline/manifest.json`;
- new: `drmcp/src/internal/designrecords/testdata/read-baseline/config/current-plus-legacy/config.json`;
- new: five files below `drmcp/src/internal/designrecords/testdata/read-baseline/legacy/v01/records/`;
- new: `drmcp/src/internal/designrecords/testdata/read-baseline/current/drmcp/records/requirements/mcp/DRMCP-REQ-MCP-903-configured-legacy-relation-fixture.md`.

Production implementation and existing Go test files are outside this boundary.

### Filesystem static inspection

The bounded fixture tree was reread after all fixture writes.
It contains the existing T02 current roots, the new `config/current-plus-legacy/config.json`, one new current relation source, and exactly five positive legacy Markdown sources below `legacy/v01/records/`.

The five legacy filenames match the accepted ADR, INV, REQ, WORK, and TASK lexical forms and their manifest exact inputs.
A bounded search below `legacy/` found no `V01-SPEC-*` file.
The current roots and `legacy/v01/records` are distinct sibling trees, so neither is equal to, above, or below the other.

The manifest was reread after extension.
The edit diff preserves every existing C case object and adds only the top-level legacy root, the configured arrangement pointer, and L01 through L09 after C17.
Every L01 through L09 declared existing path was observed in the bounded fixture tree.
The current relation source contains exact `source_refs` value `V01-REQ-MCP-901`, and the target exists only below the legacy root.
L09 records zero usable current targets and one unique readable legacy source without recording runtime execution output.

W008 was reread with Task order T01, T02, T03.
T01, T02, and this Task are `done`.
W008 and hub T08 remain `in_progress`.

Filesystem static inspection does not execute a JSON parser, production parser, resolver, validator, runtime operation, or existing Go test.
The ready-to-run external verifier below owns the executable JSON and lifecycle assertions.

### Command execution state

`git.inspect_worktree` is used for bounded status and whitespace inspection when available.
No production runtime command, Go test, formatter, generator, or repository-wide validation is executed in this Task session.
Repository-wide cleanliness is not asserted.

### Ready-to-run external verifier

Run from the repository root after the final T03 edit:

```powershell
@'
import json
import re
from pathlib import Path

fixture = Path("drmcp/src/internal/designrecords/testdata/read-baseline")
manifest = json.loads((fixture / "manifest.json").read_text(encoding="utf-8"))

expected_current = [*(f"C{i:02d}" for i in range(1, 15)), "C17"]
expected_legacy = [f"L{i:02d}" for i in range(1, 10)]
ids = [case["id"] for case in manifest["cases"]]
assert ids[:len(expected_current)] == expected_current
assert ids[len(expected_current):] == expected_legacy
assert len(ids) == len(set(ids))

legacy_root = fixture / "legacy/v01/records"
current_roots = [fixture / root["repository_path"] for root in manifest["roots"]["current"] if root["role"] == "normal"]
legacy_roots = [fixture / root["repository_path"] for root in manifest["roots"]["legacy"] if root["role"] == "normal"]
assert legacy_roots == [legacy_root]
for current in current_roots:
    assert current != legacy_root
    assert current not in legacy_root.parents
    assert legacy_root not in current.parents

legacy_cases = manifest["cases"][len(expected_current):]
for case in legacy_cases:
    assert case["expected_classification"] == "accepted"
    assert case["intentional_invalidity"] is None
    assert isinstance(case["runtime_owner"], list) and case["runtime_owner"]
    assert "DRMCP-WORK-MCP-010" in case["runtime_owner"]
    assert case["planned_task"] == "DRMCP-TASK-MCP-008-03"
    assert case["legacy_roots_state"] == "configured"
    for rel in case["repository_paths"]:
        assert "\\" not in rel and not Path(rel).is_absolute()
        assert (fixture / rel).exists(), (case["id"], rel)

patterns = {
    "decision": re.compile(r"^(V01-ADR-[0-9]{3})(?:-[a-z0-9][a-z0-9-]*)?\.md$"),
    "investigation": re.compile(r"^(V01-INV-[A-Z][A-Z0-9]*-[0-9]{3})(?:-[a-z0-9][a-z0-9-]*)?\.md$"),
    "requirement": re.compile(r"^(V01-REQ-[A-Z][A-Z0-9]*-[0-9]{3})(?:-[a-z0-9][a-z0-9-]*)?\.md$"),
    "work_item": re.compile(r"^(V01-WORK-[A-Z][A-Z0-9]*-[0-9]{3})(?:-[a-z0-9][a-z0-9-]*)?\.md$"),
    "task": re.compile(r"^(V01-TASK-[A-Z][A-Z0-9]*-[0-9]{3}-[0-9]{2})(?:-[a-z0-9][a-z0-9-]*)?\.md$"),
}
source_by_case = {case["id"]: case for case in legacy_cases}
for case_id, kind in [("L02", "decision"), ("L03", "investigation"), ("L04", "requirement"), ("L05", "work_item"), ("L06", "task")]:
    path = Path(source_by_case[case_id]["repository_paths"][0])
    match = patterns[kind].match(path.name)
    assert match
    assert source_by_case[case_id]["exact_inputs"] == [match.group(1)]

assert not any(path.name.startswith("V01-SPEC-") for path in legacy_root.rglob("*.md"))
assert {"ADR", "INV", "REQ", "WORK", "TASK"} == {
    "ADR" if path.name.startswith("V01-ADR-") else
    "INV" if path.name.startswith("V01-INV-") else
    "REQ" if path.name.startswith("V01-REQ-") else
    "WORK" if path.name.startswith("V01-WORK-") else
    "TASK"
    for path in legacy_root.rglob("V01-*.md")
}

l08 = source_by_case["L08"]["relation"]
assert (fixture / l08["source_path"]).exists()
assert (fixture / l08["target_path"]).exists()
assert l08["target_ref"] == "V01-REQ-MCP-901"

l09 = source_by_case["L09"]["fallback_arrangement"]
assert l09["exact_input"] == "V01-REQ-MCP-901"
assert l09["current_stage"]["usable_target_count"] == 0
assert l09["legacy_stage"]["usable_source_count"] == 1
assert (fixture / l09["legacy_stage"]["source_path"]).exists()

work = Path("drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md").read_text(encoding="utf-8")
hub = Path("drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md").read_text(encoding="utf-8")
task = Path("drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-03-create-configured-legacy-root-read-fixture-baseline.md").read_text(encoding="utf-8")
assert "- **status**: done" in task
assert "- **status**: in_progress" in work
assert "- **status**: in_progress" in hub
assert work.index("  - DRMCP-TASK-MCP-008-01") < work.index("  - DRMCP-TASK-MCP-008-02") < work.index("  - DRMCP-TASK-MCP-008-03")
print("fixture_shape=OK")
print("lifecycle_shape=OK")
'@ | python -
```

Expected result:

```text
fixture_shape=OK
lifecycle_shape=OK
```

### Independent review prompt

```text
C:\Users\imved\projects\brewprint

`DRMCP-TASK-MCP-008-03-create-configured-legacy-root-read-fixture-baseline`の独立reviewを行う。

ファイルは変更しないこと。
DRMCPは現在利用できない。
repositoryの読み込みにはfilesystem MCPを使用すること。
Gitの限定status・whitespace確認には、利用可能なら`git.inspect_worktree`を使用すること。
sandboxへrepositoryを複製しないこと。
repository-wide traversalを行わないこと。
repository-wide clean statusを推測しないこと。
repository-local commandを実行できない場合、実行したと捏造しないこと。
production implementationと既存Go testは変更しないこと。

## 最初に読む

1. `prompt_chappy.md`
2. `product/records/spec/design-records/authoring-standards/task-authoring.md`
3. `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
4. `product/records/spec/design-records/authoring-standards/writing-standard.md`
5. `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`

## Primary review records

- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-03-create-configured-legacy-root-read-fixture-baseline.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-01-build-current-and-legacy-fixture-coverage-matrix.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-02-create-current-root-read-fixture-baseline.md`
- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md`
- `drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-design-records-mcp-contract-baseline-and-realignment.md`

## Accepted authorities

- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-005-resolver-and-configured-legacy-fallback-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`
- `product/records/spec/brewprint/compatibility/legacy-id-compatibility.md`
- `drmcp/records/spec/design-records-mcp/namespace-scanning.md`
- `drmcp/records/spec/design-records-mcp/resolver.md`
- `drmcp/records/spec/design-records-mcp/tools/get-records.md`
- `drmcp/records/spec/design-records-mcp/tools/resolve-reference.md`

## Fixture scope

Review only:

- `drmcp/src/internal/designrecords/testdata/read-baseline/manifest.json`;
- `drmcp/src/internal/designrecords/testdata/read-baseline/config/current-plus-legacy/config.json`;
- five files below `drmcp/src/internal/designrecords/testdata/read-baseline/legacy/v01/records/`;
- `drmcp/src/internal/designrecords/testdata/read-baseline/current/drmcp/records/requirements/mcp/DRMCP-REQ-MCP-903-configured-legacy-relation-fixture.md`;
- T03 lifecycle synchronization in W008 and hub T08.

Confirm at minimum:

- C01 through C14 and C17 remain unchanged from accepted T02 state;
- L01 through L09 exist once and in order after C17;
- all manifest internal paths are fixture-root-relative and `/`-separated;
- current and legacy normal roots are physically and logically disjoint;
- the positive legacy family set is exactly ADR, INV, REQ, WORK, and TASK;
- no positive `V01-SPEC-*` source exists;
- filename-derived issued IDs and exact inputs match accepted lexical grammar;
- existing V01 archive placement shape is followed;
- configuration is explicitly fixture-local and does not claim production serialization authority;
- L07 proves only exact grammar, configured root, one readable source, and exact ID match;
- L08 has an exact current-to-legacy relation, separate roots, and W010 ownership;
- the legacy target is not treated as a current validation subject;
- L09 records one current miss and one unique readable legacy source without asserting runtime execution order or response;
- every L01 through L09 owner is a non-empty list containing `DRMCP-WORK-MCP-010`;
- every L01 through L09 planned Task is T03;
- no L10-L13, R01-R24, disabled, empty-list, unresolved, invalid-root, duplicate, overlap, leakage, production implementation, or existing Go test change exists;
- T03, W008, and hub T08 remain `in_progress` before review;
- T01 and T02 remain `done`;
- W008 task order is T01, T02, T03;
- changed-file boundary and verification evidence are accurate;
- repository-wide clean status is not inferred.

When repository-local execution is available, run the ready-to-run verifier recorded in T03 Evidence and scoped whitespace checks.
Do not write review results into the files.

出力形式:

1. Verdict: PASS / NEEDS REVISION
2. Previous-finding disposition
3. Blocking findings
4. Major findings
5. Minor findings
6. Advisories
7. Legacy family and lexical mapping assessment
8. Root and configuration separation assessment
9. L01-L09 coverage assessment
10. Current-to-legacy relation assessment
11. Fallback arrangement boundary assessment
12. Manifest non-regression assessment
13. Lifecycle and changed-file assessment
14. Verification-evidence assessment
15. T03 closure readiness
16. T04 start readiness
```

### Independent review and closure

Independent review verdict: `PASS`.

- Previous T03 findings: none.
- T02 findings F-MAJ-01, F-MAJ-02, and F-MIN-01 remain closed with no regression.
- Blocking findings: none.
- Major findings: none.
- Minor findings: none.
- Legacy family and lexical mapping assessment: `PASS`.
- Root and configuration separation assessment: `PASS`.
- L01 through L09 coverage assessment: `PASS`.
- Current-to-legacy relation assessment: `PASS`.
- Fallback arrangement boundary assessment: `PASS`.
- Manifest non-regression assessment: `PASS`.
- Lifecycle and changed-file assessment: `PASS`.
- Verification-evidence assessment: `PASS` with the recorded command-execution limitation.

The reviewer independently ran scoped `git.inspect_worktree` checks for the eleven T03 files.
The result was `pass` with no whitespace finding.
LF-to-CRLF conversion warnings were non-blocking.
`repository_wide_clean` remained `null`, and no repository-wide cleanliness claim was made.

The ready-to-run Python verifier was `NOT RUN` during review because the available review tools did not support arbitrary repository-local command execution.
No review-time Python execution result is inferred.

After closure synchronization, the external verifier was run from the repository root and returned:

```text
fixture_shape=OK
lifecycle_shape=OK
```

This confirms the recorded fixture shape and closure lifecycle assertions against the final T03 state.

T03 closure readiness: `READY`.
T04 start readiness: `READY AFTER T03 CLOSURE`.

This Task is `done`.
W008 and hub T08 remain `in_progress` for T04 and T05.
One final scoped Git and whitespace check follows the last closure edit and is reported outside these files.
