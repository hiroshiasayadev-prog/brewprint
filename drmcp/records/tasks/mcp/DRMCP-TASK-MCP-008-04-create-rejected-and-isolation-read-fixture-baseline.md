# DRMCP-TASK-MCP-008-04: Create rejected and isolation read fixture baseline

- **id**: DRMCP-TASK-MCP-008-04
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-008
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-008-01
  - DRMCP-TASK-MCP-008-02
  - DRMCP-TASK-MCP-008-03
- **outputs**:
  - DRMCP-WORK-MCP-008
  - DRMCP-TASK-MCP-001-08
  - drmcp/src/internal/designrecords/testdata/read-baseline/manifest.json
  - drmcp/src/internal/designrecords/testdata/read-baseline/config/legacy-roots-empty/
  - drmcp/src/internal/designrecords/testdata/read-baseline/arrangements/

## Goal

Materialize T01 cases C15, C16, L10 through L13, and R01 through R24.

Represent rejected inputs, intentional invalidity, root conflicts, isolation expectations, leakage prohibitions, and invalid Topics graphs without asserting production runtime results.

## Work

- Confirm the bounded Task and fixture pre-creation inventory.
- Preserve all accepted T02 and T03 case objects and normal fixture sources.
- Add one intentionally invalid current source with a portable source location.
- Reuse approved legacy sources for isolation cases without duplicating V01 records.
- Isolate rejected `V01-SPEC-*`, YAML-front-matter, duplicate identity, invalid root, duplicate root, overlapping root, and invalid Topics sources below `arrangements/`.
- Add explicit empty `legacy_roots` configuration below `config/legacy-roots-empty/`.
- Extend the authoritative manifest with C15, C16, L10 through L13, and R01 through R24 in deterministic order.
- Record every intentional absence, conflict source, overlap relation, isolation flag, forbidden leakage observation, runtime owner, and planned Task.
- Keep runtime assertions in W009, W010, W-SPEC-001, and W-SPEC-002.
- Synchronize W008 and hub T08 without closing either record.
- Keep production implementation and existing Go tests unchanged.

## Done condition

- C15, C16, L10 through L13, and R01 through R24 exist once in the authoritative manifest.
- Existing C01 through C14, C17, and L01 through L09 case object content remains unchanged.
- Every T04 case has a non-empty runtime-owner list and `planned_task: DRMCP-TASK-MCP-008-04`.
- Every rejected or invalid arrangement declares `intentional_invalidity`.
- Normal current and legacy roots remain unchanged and disjoint.
- Invalid, duplicate, overlap, rejection, and graph arrangements remain outside normal roots.
- Duplicate sources list every conflict and select no winner.
- Equal, ancestor, and descendant root overlap relations are explicit.
- Omitted and empty `legacy_roots` remain separate disabled cases.
- Unresolved current, unresolved accepted legacy, and unsupported legacy-family cases remain distinct.
- L10 through L13 contain the required isolation flags.
- R17 through R20 and R24 express prohibited leakage rather than observed runtime leakage.
- R21 through R23 contain the required Topics graph source arrangements.
- Production implementation and existing Go tests remain unchanged.
- Independent review reports no blocking or major finding before this Task changes to `done`.

## Verification

- Parse `manifest.json` as JSON.
- Confirm the complete required case set, unique case IDs, and deterministic ordering.
- Confirm accepted T02 and T03 case IDs retain their accepted classifications, null invalidity, owners, planned Tasks, and existing paths.
- Resolve every declared existing path from the fixture root.
- Confirm every intentional absent path does not exist.
- Confirm all fixture paths are relative and use `/` separators.
- Confirm normal current and legacy roots remain disjoint.
- Confirm arrangement roots are not registered as normal roots.
- Confirm duplicate-current and duplicate-parent source sets are complete and have no winner.
- Confirm equal, ancestor, and descendant overlap relations match their physical paths.
- Confirm omitted and empty legacy configurations remain separate.
- Confirm disabled, unresolved, unsupported, duplicate, and rejected classifications remain distinct.
- Confirm the positive normal legacy tree contains only ADR, INV, REQ, WORK, and TASK families.
- Confirm no positive `V01-SPEC-*` source exists below the normal legacy root.
- Confirm L10 through L13 isolation flags.
- Confirm R17 through R20 and R24 forbidden-leakage metadata.
- Confirm R21 through R23 Topics graph arrangements.
- Confirm T01 through T03 are `done`.
- Confirm this Task is `done` while W008 and hub T08 remain `in_progress`.
- Confirm W008 Task order is T01, T02, T03, T04.
- Run scoped Git status and whitespace inspection without asserting repository-wide cleanliness.

## Evidence

### Bounded pre-creation inventory

The exact directory `drmcp/records/tasks/mcp/` was listed once before T04 creation.
The only W008 child Tasks were:

- `DRMCP-TASK-MCP-008-01-build-current-and-legacy-fixture-coverage-matrix.md`;
- `DRMCP-TASK-MCP-008-02-create-current-root-read-fixture-baseline.md`;
- `DRMCP-TASK-MCP-008-03-create-configured-legacy-root-read-fixture-baseline.md`.

No file or Task ID for `DRMCP-TASK-MCP-008-04` existed.
The selected exact ID is `DRMCP-TASK-MCP-008-04` because T01 allocated rejected and isolation materialization to the fourth W008 child Task.
No existing Task was overwritten.

The bounded fixture-root inventory contained only:

- `manifest.json`;
- normal `current/` roots;
- normal `legacy/` root;
- `config/current-only/`;
- `config/current-plus-legacy/`;
- `config/legacy-roots-omitted/`.

Before T04 creation:

- `arrangements/` did not exist;
- `config/legacy-roots-empty/` did not exist;
- no rejected `V01-SPEC-*` fixture existed;
- no invalid current source, duplicate-current arrangement, root-error arrangement, or invalid Topics graph arrangement existed.

### Accepted authority and runtime ownership

| concern | accepted authority or runtime owner | T04 boundary |
|---|---|---|
| Current discovery, identity, duplicate isolation, and active-index scope | `DRMCP-WORK-MCP-003` | Supplies current source and root arrangement semantics. |
| Listing, exact retrieval, and normal path hiding | `DRMCP-WORK-MCP-004`; runtime `DRMCP-WORK-MCP-009` | T04 records fixture expectations only. |
| Exact grammar, no repair, configured legacy fallback, and fallback states | `DRMCP-WORK-MCP-005`; runtime `DRMCP-WORK-MCP-010` | T04 records input and source arrangements only. |
| Validation subject scope and diagnostic source locations | `DRMCP-WORK-MCP-006` | T04 separates exceptional source location from normal path leakage. |
| Accepted legacy families | `spec:product.brewprint.compatibility.legacy_id_compatibility` | The normal legacy tree remains ADR, INV, REQ, WORK, and TASK only. |
| Current spec document shape | `spec:product.design_records.spec_format.document_shape`; runtime `DRMCP-WORK-SPEC-001` | T04 supplies invalid source material without detector assertions. |
| Topics graph contract | `spec:product.design_records.spec_format.topics_table`; runtime `DRMCP-WORK-SPEC-002` | T04 supplies edge and source arrangements without graph execution. |
| Current runtime assertions | `DRMCP-WORK-MCP-009` | Owns current read, invalid source, unresolved current, and normal path-hiding assertions. |
| Legacy runtime assertions | `DRMCP-WORK-MCP-010` | Owns fallback, root errors, unsupported families, isolation, and leakage assertions. |

### Case materialization mapping

| case | materialization |
|---|---|
| C15 | Invalid current spec under `arrangements/invalid-current-source/` with missing H1-adjacent `parent` and a fixture-root-relative source location. |
| C16 | Existing valid current sources plus forbidden normal-response path fields in manifest metadata. |
| L10 | Existing five approved legacy sources with `normal_listing_candidate: false`. |
| L11 | Existing five approved legacy sources with `current_validation_subject: false`. |
| L12 | Existing five approved legacy sources with `authoring_target: false`. |
| L13 | Existing five approved legacy sources with `active_index_candidate: false`. |
| R01 | `V01-SPEC-901` exists only under `arrangements/rejected-legacy/` and is marked non-accepted. |
| R02 | App-prefixless `REQ-MCP-901` literal with no accepted target or inference hint. |
| R03 | Repository-relative physical path recorded only as a rejected canonical input. |
| R04 | Partial `DRMCP-REQ-MCP-90` literal with no repaired value. |
| R05 | Valid-looking `TASK-MCP-901-01` suffix without an app prefix. |
| R06 | YAML-front-matter current spec under `arrangements/invalid-spec-format/`. |
| R07 | Two current sources with canonical ID `DRMCP-REQ-MCP-990`; both are listed and no winner exists. |
| R08 | Missing current-root path declared by `arrangements/invalid-root/config.json`. |
| R09 | Missing legacy-root path declared by the same bounded invalid-root arrangement. |
| R10 | Duplicate `current/product/records` declarations at positions 0 and 1 with no winner. |
| R11 | Equal, current-ancestor-of-legacy, and current-descendant-of-legacy arrangements. |
| R12 | Existing omitted-`legacy_roots` configuration classified as disabled. |
| R13 | New explicit empty-`legacy_roots` configuration classified as disabled. |
| R14 | Valid current input `DRMCP-REQ-MCP-999` with one intentional absent target. |
| R15 | Approved legacy input `V01-REQ-MCP-999` with configured roots and one intentional absent source. |
| R16 | `V01-SPEC-901` classified outside the exact five-family allowlist. |
| R17 | Forbidden normal-listing leakage metadata for an existing legacy source. |
| R18 | Forbidden current-validation-subject leakage metadata for an existing legacy source. |
| R19 | Forbidden authoring-target leakage metadata for an existing legacy source. |
| R20 | Forbidden physical-path fields on normal listing or retrieval surfaces. |
| R21 | Topics parent source with one intentionally absent canonical child. |
| R22 | Two authoritative Topics parents declaring one existing child; no parent winner. |
| R23 | Two source files forming the closed edge sequence A to B to A. |
| R24 | Forbidden active-current-index leakage metadata for an existing legacy source. |

### Existing accepted fixture non-regression

The authoritative `manifest.json` was extended in three bounded locations:

- T04 top-level materialized scope, configuration, and arrangement declarations;
- C15 and C16 immediately before existing C17;
- L10 through L13 and R01 through R24 after existing L09.

The existing C01 through C14, C17, and L01 through L09 case object bodies were not edited.
C15 and C16 insertion changed ordering around C17 without changing C17 fields.
No existing normal current or legacy source file was changed.
No approved V01 source was duplicated.

### Intentional invalidity representation

Every T04 rejected, disabled, unresolved, duplicate, overlap, and leakage case records a structured `intentional_invalidity` object.
Intentional absent paths are listed under `intentional_absent_paths` and must remain absent.
Duplicate cases list every conflicting source and declare `winner: null`.
Overlap cases list exact root pairs and one of:

- `equal`;
- `current_ancestor_of_legacy`;
- `current_descendant_of_legacy`.

Forbidden leakage cases record `must_not_occur: true`.
They do not claim that a runtime operation produced the forbidden observation.

### Normal-root and arrangement separation

Normal roots remain:

- `current/product/records`;
- `current/drmcp/records`;
- `legacy/v01/records`.

All invalid source, rejection, duplicate, root-error, overlap, and invalid Topics graph material lives below `arrangements/`.
No `arrangements/` path is added to `roots.current` or `roots.legacy`.
Arrangement-local current and rejection roots are declared only inside their `arrangements` entries with non-normal roles.
The rejected `V01-SPEC-901` source is outside `legacy/v01/records`.

### Changed-file boundary

Workflow files:

- new: `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-04-create-rejected-and-isolation-read-fixture-baseline.md`;
- modified: `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md`;
- modified: `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md`.

Fixture files:

- modified: `drmcp/src/internal/designrecords/testdata/read-baseline/manifest.json`;
- new: `drmcp/src/internal/designrecords/testdata/read-baseline/config/legacy-roots-empty/config.json`;
- new: `drmcp/src/internal/designrecords/testdata/read-baseline/arrangements/invalid-current-source/current/product/records/spec/invalid-source/missing-parent.md`;
- new: `drmcp/src/internal/designrecords/testdata/read-baseline/arrangements/rejected-legacy/legacy/v01/records/spec/V01-SPEC-901-rejected-legacy-spec.md`;
- new: `drmcp/src/internal/designrecords/testdata/read-baseline/arrangements/invalid-spec-format/current/product/records/spec/invalid-format/yaml-current-spec.md`;
- new: two files below `arrangements/duplicate-current/`;
- new: `arrangements/invalid-root/config.json`;
- new: `arrangements/duplicate-root/config.json`;
- new: `arrangements/overlapping-root/config.json` and three fixture-local `root-marker.json` files below its equal and nested leaf roots;
- new: six files below `arrangements/topics-graph/`.

Production `.go` implementation and existing Go test files are outside this boundary.

### Explicit exclusions

T04 does not:

- reopen T01 through T03;
- change an accepted ADR, Requirement, or Specification;
- change production `.go` implementation;
- change existing Go tests;
- call a production parser, index, resolver, validator, or MCP operation;
- claim runtime listing, retrieval, resolution, validation, authoring, indexing, or diagnostic results;
- add `V01-SPEC-*` to the accepted legacy tree;
- register invalid or overlapping paths as normal roots;
- select a duplicate identity or duplicate Topics parent winner;
- repair malformed, partial, fuzzy, or app-prefixless inputs;
- infer repository-wide clean status.

### Command execution state

The available filesystem MCP does not execute arbitrary repository-local commands.
The Python verifier below is therefore `NOT RUN` in this Task session.
No production parser, resolver, validator, MCP server, existing Go test, formatter, generator, or repository-wide validation was executed.

`git.inspect_worktree` is used separately for bounded status and whitespace inspection only.
It is not treated as a test or verifier substitute.

### Ready-to-run fixture and lifecycle verifier

Run from the repository root after the final T04 edit:

```powershell
@'
import json
import subprocess
from pathlib import Path, PurePosixPath

manifest_rel = "drmcp/src/internal/designrecords/testdata/read-baseline/manifest.json"
fixture = Path("drmcp/src/internal/designrecords/testdata/read-baseline")
manifest = json.loads((fixture / "manifest.json").read_text(encoding="utf-8"))
baseline = json.loads(subprocess.check_output(
    ["git", "show", f"HEAD:{manifest_rel}"],
    text=True,
    encoding="utf-8",
))

expected_ids = [
    *(f"C{i:02d}" for i in range(1, 18)),
    *(f"L{i:02d}" for i in range(1, 14)),
    *(f"R{i:02d}" for i in range(1, 25)),
]
ids = [case["id"] for case in manifest["cases"]]
assert ids == expected_ids, ids
assert len(ids) == len(set(ids))
cases = {case["id"]: case for case in manifest["cases"]}

accepted_t02 = [*(f"C{i:02d}" for i in range(1, 15)), "C17"]
accepted_t03 = [f"L{i:02d}" for i in range(1, 10)]
baseline_cases = {case["id"]: case for case in baseline["cases"]}
assert list(baseline_cases) == [*accepted_t02[:-1], "C17", *accepted_t03]
for case_id in accepted_t02 + accepted_t03:
    assert cases[case_id] == baseline_cases[case_id], case_id
for case_id in accepted_t02:
    case = cases[case_id]
    assert case["expected_classification"] == "accepted"
    assert case["intentional_invalidity"] is None
    assert case["planned_task"] == "DRMCP-TASK-MCP-008-02"
    assert isinstance(case["runtime_owner"], list) and case["runtime_owner"]
for case_id in accepted_t03:
    case = cases[case_id]
    assert case["expected_classification"] == "accepted"
    assert case["intentional_invalidity"] is None
    assert case["planned_task"] == "DRMCP-TASK-MCP-008-03"
    assert isinstance(case["runtime_owner"], list) and case["runtime_owner"]

t04_ids = ["C15", "C16", *(f"L{i:02d}" for i in range(10, 14)), *(f"R{i:02d}" for i in range(1, 25))]
for case_id in t04_ids:
    case = cases[case_id]
    assert case["planned_task"] == "DRMCP-TASK-MCP-008-04"
    assert isinstance(case["runtime_owner"], list) and case["runtime_owner"]

expected_t04_classification = {
    "C15": "accepted", "C16": "accepted",
    "L10": "accepted", "L11": "accepted", "L12": "accepted", "L13": "accepted",
    "R01": "rejected", "R02": "rejected", "R03": "rejected", "R04": "rejected",
    "R05": "rejected", "R06": "rejected", "R07": "duplicate", "R08": "rejected",
    "R09": "rejected", "R10": "duplicate", "R11": "rejected", "R12": "disabled",
    "R13": "disabled", "R14": "unresolved", "R15": "unresolved", "R16": "rejected",
    "R17": "rejected", "R18": "rejected", "R19": "rejected", "R20": "rejected",
    "R21": "unresolved", "R22": "duplicate", "R23": "rejected", "R24": "rejected",
}
assert {case_id: cases[case_id]["expected_classification"] for case_id in t04_ids} == expected_t04_classification
for case_id in ["C15", *(f"R{i:02d}" for i in range(1, 25))]:
    assert isinstance(cases[case_id]["intentional_invalidity"], dict)

for case in manifest["cases"]:
    for key in ("repository_paths", "intentional_absent_paths"):
        for rel in case.get(key, []):
            assert "\\" not in rel
            assert not Path(rel).is_absolute()
            assert not PurePosixPath(rel).is_absolute()
            assert ".." not in PurePosixPath(rel).parts
    for rel in case.get("repository_paths", []):
        assert (fixture / rel).exists(), (case["id"], rel)
    for rel in case.get("intentional_absent_paths", []):
        assert not (fixture / rel).exists(), (case["id"], rel)

for item in manifest["configurations"] + manifest["arrangements"]:
    rel = item["repository_path"]
    assert "\\" not in rel and not Path(rel).is_absolute()
    assert (fixture / rel).exists(), rel

normal_current = [fixture / root["repository_path"] for root in manifest["roots"]["current"] if root["role"] == "normal"]
normal_legacy = [fixture / root["repository_path"] for root in manifest["roots"]["legacy"] if root["role"] == "normal"]
for current in normal_current:
    for legacy in normal_legacy:
        assert current != legacy
        assert current not in legacy.parents
        assert legacy not in current.parents
arrangement_root_names = {}
for arrangement in manifest["arrangements"]:
    assert arrangement["role"] == "arrangement_only"
    arrangement_path = fixture / arrangement["repository_path"]
    assert arrangement_path not in normal_current + normal_legacy
    for root_group in ("current_roots", "legacy_roots"):
        for root in arrangement.get(root_group, []):
            assert root["role"] in {"arrangement_only", "rejection_only"}
            rel = root["repository_path"]
            assert "\\" not in rel and not Path(rel).is_absolute()
            assert ".." not in PurePosixPath(rel).parts
            root_path = fixture / rel
            assert root_path.exists(), root["repository_path"]
            assert root_path not in normal_current + normal_legacy
            for normal in normal_current + normal_legacy:
                assert root_path != normal
                assert root_path not in normal.parents
                assert normal not in root_path.parents
            assert root["name"] not in arrangement_root_names
            arrangement_root_names[root["name"]] = arrangement["name"]

assert cases["C15"]["fixture_arrangement"] == "invalid-current-source"
assert cases["C15"]["current_roots"] == ["invalid-source-product-current"]
assert cases["C15"]["source_location"] == cases["C15"]["repository_paths"][0]
assert (fixture / cases["C15"]["source_location"]).exists()
assert cases["R01"]["legacy_roots"] == ["rejected-v01-legacy"]
assert cases["R06"]["current_roots"] == ["invalid-format-product-current"]
assert cases["R07"]["root_arrangement"] == "current-only"
assert cases["R07"]["current_roots"] == ["duplicate-current-a", "duplicate-current-b"]
for case_id in ["R21", "R22", "R23"]:
    assert cases[case_id]["fixture_arrangement"] == "topics-graph"
    assert cases[case_id]["current_roots"] == ["topics-graph-product-current"]

assert cases["R07"]["conflicting_sources"] == cases["R07"]["repository_paths"]
assert cases["R07"]["winner"] is None
assert cases["R10"]["declaration_positions"] == [0, 1]
assert cases["R10"]["winner"] is None
assert cases["R22"]["topics_graph"]["parent_sources"] == cases["R22"]["repository_paths"][:2]
assert cases["R22"]["topics_graph"]["winner"] is None

relations = cases["R11"]["overlap_relations"]
assert [item["relationship"] for item in relations] == [
    "equal",
    "current_ancestor_of_legacy",
    "current_descendant_of_legacy",
]
for item in relations:
    current = fixture / item["current_root"]
    legacy = fixture / item["legacy_root"]
    if item["relationship"] == "equal":
        assert current == legacy
    elif item["relationship"] == "current_ancestor_of_legacy":
        assert current in legacy.parents
    else:
        assert legacy in current.parents

assert cases["R12"]["legacy_roots_state"] == "omitted"
assert cases["R13"]["legacy_roots_state"] == "empty"
assert cases["R12"]["expected_classification"] == "disabled"
assert cases["R13"]["expected_classification"] == "disabled"
assert cases["R14"]["expected_classification"] == "unresolved"
assert cases["R15"]["expected_classification"] == "unresolved"
assert cases["R16"]["expected_classification"] == "rejected"
assert cases["R16"]["approved_legacy_families"] == ["ADR", "INV", "REQ", "WORK", "TASK"]

legacy_root = fixture / "legacy/v01/records"
legacy_files = sorted(path.name for path in legacy_root.rglob("V01-*.md"))
assert len(legacy_files) == 5
assert not any(name.startswith("V01-SPEC-") for name in legacy_files)
assert {
    "ADR" if name.startswith("V01-ADR-") else
    "INV" if name.startswith("V01-INV-") else
    "REQ" if name.startswith("V01-REQ-") else
    "WORK" if name.startswith("V01-WORK-") else
    "TASK"
    for name in legacy_files
} == {"ADR", "INV", "REQ", "WORK", "TASK"}

assert cases["L10"]["isolation_flags"] == {"normal_listing_candidate": False}
assert cases["L11"]["isolation_flags"] == {"current_validation_subject": False}
assert cases["L12"]["isolation_flags"] == {"authoring_target": False}
assert cases["L13"]["isolation_flags"] == {"active_index_candidate": False}
for case_id in ["R17", "R18", "R19", "R20", "R24"]:
    assert cases[case_id]["forbidden_leakage"]["must_not_occur"] is True

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
    "work": Path("drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md"),
    "hub": Path("drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md"),
}
text = {name: path.read_text(encoding="utf-8") for name, path in records.items()}
for name in ["t01", "t02", "t03"]:
    assert "- **status**: done" in text[name]
assert "- **status**: done" in text["t04"]
assert "- **status**: in_progress" in text["work"]
assert "- **status**: in_progress" in text["hub"]
assert text["work"].index("  - DRMCP-TASK-MCP-008-01") < text["work"].index("  - DRMCP-TASK-MCP-008-02") < text["work"].index("  - DRMCP-TASK-MCP-008-03") < text["work"].index("  - DRMCP-TASK-MCP-008-04")
print("fixture_shape=OK")
print("lifecycle_shape=OK")
'@ | python -
```

Expected result:

```text
fixture_shape=OK
lifecycle_shape=OK
```

The verifier reads fixture and workflow files plus the committed T03 manifest baseline through `git show HEAD:`.
It compares every pre-existing accepted case object as a whole.
It does not import or call production parser, index, resolver, validator, MCP server, or Go tests.

### Post-closure external verifier result

After independent review acceptance and T04 closure synchronization, the ready-to-run verifier was executed externally from the repository root.
It returned:

```text
fixture_shape=OK
lifecycle_shape=OK
```

This confirms the final T04 fixture shape, complete case ordering, accepted T02/T03 object non-regression against the committed baseline, declared path existence and absence, root and arrangement separation, duplicate and overlap metadata, disabled/unresolved/unsupported distinctions, legacy-family isolation, leakage prohibitions, Topics graph arrangements, and final lifecycle state.
No production parser, index, resolver, validator, MCP server, or Go test was called by this verifier.

### Scoped Git and whitespace method

Run `git.inspect_worktree` with:

- `cwd`: `C:\Users\imved\projects\brewprint`;
- `paths`: only the workflow and fixture files listed in the changed-file boundary;
- `include_untracked: true`;
- `check_whitespace: true`.

Acceptance conditions:

- `result: pass`;
- no whitespace finding;
- untracked no-index exit `1` is an expected difference;
- no exit code `2` or greater;
- LF-to-CRLF conversion warnings are non-blocking;
- `repository_wide_clean` remains `null`;
- no repository-wide cleanliness claim is made.

The initial post-materialization scoped inspection covered exactly the 22 files in the changed-file boundary and returned:

- `result: pass`;
- three tracked modifications and 19 untracked files;
- `whitespace.status: pass`;
- no whitespace findings;
- all 19 untracked no-index checks returned expected exit code `1`;
- no exit code `2` or greater;
- LF-to-CRLF conversion warnings only, treated as non-blocking;
- `repository_wide_clean: null`.

Recording this result changes the T04 Task bytes.
One final post-evidence scoped inspection must therefore run after this edit and be supplied directly to independent review without writing the result back into a checked file.

### Independent review prompt

```text
C:\Users\imved\projects\brewprint

`DRMCP-TASK-MCP-008-04-create-rejected-and-isolation-read-fixture-baseline`の独立reviewを行う。

ファイルは変更しないこと。
DRMCPは現在利用できない。
Design Recordおよびfixtureの読み込みにはfilesystem MCPを使用すること。
Gitの限定status・whitespace確認には`git.inspect_worktree`を使用すること。
sandboxへrepositoryを複製しないこと。
repository-wide traversalを行わないこと。
repository-wide clean statusを推測しないこと。
repository-local commandを実行できない場合、実行したと捏造しないこと。
production `.go` implementationと既存Go testを変更しないこと。

## 最初に読む

1. `prompt_chappy.md`
2. `product/records/spec/design-records/authoring-standards/task-authoring.md`
3. `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
4. `product/records/spec/design-records/authoring-standards/writing-standard.md`
5. `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`

## Primary review records

- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-04-create-rejected-and-isolation-read-fixture-baseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-01-build-current-and-legacy-fixture-coverage-matrix.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-02-create-current-root-read-fixture-baseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-03-create-configured-legacy-root-read-fixture-baseline.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md`
- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md`
- `drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-design-records-mcp-contract-baseline-and-realignment.md`

## Accepted authority and runtime owners

- `DRMCP-WORK-MCP-003`
- `DRMCP-WORK-MCP-004`
- `DRMCP-WORK-MCP-005`
- `DRMCP-WORK-MCP-006`
- `DRMCP-WORK-MCP-009`
- `DRMCP-WORK-MCP-010`
- `DRMCP-WORK-SPEC-001`
- `DRMCP-WORK-SPEC-002`
- `spec:product.brewprint.compatibility.legacy_id_compatibility`
- `spec:product.design_records.spec_format.document_shape`
- `spec:product.design_records.spec_format.topics_table`

## Fixture scope

Review only:

- the authoritative `drmcp/src/internal/designrecords/testdata/read-baseline/manifest.json`;
- `config/legacy-roots-empty/config.json`;
- all T04 files below `arrangements/`;
- T04 lifecycle synchronization in W008 and hub T08.

Confirm at minimum:

- C15, C16, L10 through L13, and R01 through R24 exist exactly once and in deterministic order;
- C01 through C14, C17, and L01 through L09 case objects retain their accepted content;
- all paths are fixture-root-relative and `/`-separated;
- declared existing paths exist and intentional absent paths do not exist;
- normal current and legacy roots remain unchanged and disjoint;
- arrangement paths are not normal roots;
- `V01-SPEC-*` exists only in an explicit rejection arrangement and never as an accepted legacy source;
- the normal legacy family set is exactly ADR, INV, REQ, WORK, and TASK;
- app-prefixless, physical-path, fuzzy or partial repair, and YAML-front-matter inputs are rejected arrangements only;
- duplicate current identity lists every conflicting source and selects no winner;
- invalid current root, invalid legacy root, duplicate root, and equal, ancestor, and descendant overlap are explicit;
- omitted and empty `legacy_roots` are separate disabled cases;
- unresolved current, unresolved accepted legacy, and unsupported legacy family are separate classifications;
- L10 through L13 contain the exact required false isolation flags;
- R17 through R20 and R24 contain forbidden-leakage metadata and do not claim observed runtime leakage;
- C15 diagnostic source-location permission is distinct from C16 and R20 normal response path hiding;
- R21 through R23 contain unresolved-child, duplicate-parent, and cycle source arrangements without graph-runtime claims;
- every T04 case has a non-empty runtime-owner list and planned Task T04;
- T01 through T03 remain `done`;
- T04, W008, and hub T08 remain `in_progress` before review;
- W008 Task order is T01, T02, T03, T04;
- production `.go`, existing Go tests, accepted ADR, Requirement, and Specification files did not change;
- command and validation evidence is not fabricated;
- repository-wide cleanliness is not inferred.

When repository-local command execution is available, run the ready-to-run verifier recorded in T04 Evidence.
Run `git.inspect_worktree` only for the T04 changed-file manifest.
Do not write review results into the checked files.

出力形式:

1. Verdict: PASS / NEEDS REVISION
2. Previous-finding disposition
3. Blocking findings
4. Major findings
5. Minor findings
6. Advisories
7. Existing accepted fixture non-regression assessment
8. T04 case completeness and ordering assessment
9. Invalid, duplicate, and overlap arrangement assessment
10. Disabled, unresolved, and unsupported separation assessment
11. Legacy isolation and leakage assessment
12. Path-exposure boundary assessment
13. Topics graph arrangement assessment
14. Lifecycle and changed-file assessment
15. Verification-evidence assessment
16. T04 closure readiness
17. T05 start readiness
```

### Independent review acceptance

Independent T04 review returned `PASS` on 2026-06-28.

- blocking findings: none;
- major findings: none;
- minor findings: none;
- T02 F-MAJ-01, F-MAJ-02, and F-MIN-01 remain closed;
- no T03 finding existed and no regression was introduced;
- existing accepted fixture non-regression: PASS by static review;
- T04 case completeness and ordering: PASS;
- invalid, duplicate, and overlap arrangements: PASS;
- disabled, unresolved, and unsupported separation: PASS;
- legacy isolation and leakage boundaries: PASS;
- path-exposure boundary: PASS;
- Topics graph arrangements: PASS;
- lifecycle and changed-file boundary: PASS;
- verification evidence: PASS with the recorded command-execution limitation.

The ready-to-run Python verifier remained `NOT RUN` because no available tool could execute an arbitrary repository-local command.
The review did not infer its result.
LF-to-CRLF conversion warnings affected the scoped files but were non-blocking and produced no whitespace finding.
The reviewer observed unrelated boundary-external changes but did not attribute them to T04 or infer repository-wide cleanliness.

### Closure synchronization

This Task is `done` after independent review acceptance.
W008 and hub T08 remain `in_progress` because T05 remains pending.
T01 through T03 remain `done` and are not reopened.
The closure verifier expectation is now T04 `done`, W008 `in_progress`, and hub T08 `in_progress`.
The post-closure external verifier returned `fixture_shape=OK` and `lifecycle_shape=OK` against this final state.
