# V01-TASK-DATA-015-03: Investigate recursive named reference runtime behavior

- **id**: V01-TASK-DATA-015-03
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-DATA-015
- **source_requirement**: V01-REQ-DATA-008
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-DATA-015-02
- **outputs**:
  - Current resolver / validator / renderer behavior for recursive named model references
  - Implementation gap classification if current behavior rejects or expands recursion incorrectly
  - Follow-up implementation / fixture task input

## Goal

Investigate whether current implementation already handles the `V01-TASK-DATA-015-01` / `V01-TASK-DATA-015-02` recursive named model reference boundary.

## Work

- Create or use a minimal fixture with a struct model field referencing the same named model, for example `object_ref.parent: object_ref`.
- Check parser / resolver behavior for recursive named model TypeRef.
- Check validator behavior for recursive named model TypeRef.
- Check renderer / model-file behavior and confirm it does not infinitely expand recursive references.
- Classify the result as already-supported, implementation-gap, or spec-only without runtime impact.
- Identify follow-up implementation and fixture tasks if needed.

## Included Scope

- Runtime behavior investigation.
- Focused fixture or temporary test input if needed.
- Gap classification for resolver / validator / renderer behavior.

## Excluded Scope

- Permanent UC-002 YAML migration.
- Golden regeneration unless a later implementation task selects it.
- Untagged union / general oneOf support.
- Broadening V01-ADR-073 or V01-WORK-DATA-010.

## Done condition

- Current recursive named model reference behavior is classified.
- Required implementation / fixture follow-up is identified, or explicit no-change evidence is recorded.
- No untagged union support is introduced.

## Verification

- Run focused validation / render checks or equivalent repo-local tests.
- Record commands and results in Evidence.

## Evidence

Investigated on 2026-06-08.

### Classification

Verdict: already-supported.

Current runtime accepts recursive named model references for the tested case where a struct model field references the same named model.

### Temporary fixture

Codex created a temporary fixture under `.codex-tmp/task-data-015-03/` with a public struct model `object_ref` whose `parent` field used `type: object_ref`.

Fixture shape:

```yaml
nodes:
  - id: object_ref
    type: model
    main: true
    kind: struct
    fields:
      - name: id
        type: str
      - name: parent
        type: object_ref
        note: "recursive named model reference"
```

The temporary fixture was cleaned up after investigation.

### Commands and results

- `go test ./internal/resolve ./internal/render/model ./cmd/brewprint`
  - Result: OK.
- `go run ./cmd/brewprint validate --yaml-root .\.codex-tmp\task-data-015-03\yaml --format json`
  - Result: `{"diagnostics":null,"error_count":0,"warning_count":0}`.
- `go run ./cmd/brewprint render --yaml-root .\.codex-tmp\task-data-015-03\yaml --out .\.codex-tmp\task-data-015-03\render-out --clean`
  - Result: `rendered 3 file(s)`.
- Focused tests also passed:
  - `go test ./internal/resolve -run "Test(ParseResolveNamedModelTypeRef|BuildSetsSemanticTypeRefs|TaggedUnionUsedAsStructFieldTypeRef)$"`
  - `go test ./internal/render/model -run Test`

### Findings by layer

- Parser / resolver: `object_ref.parent: object_ref` is accepted and resolved as a named model TypeRef. No `unresolved_field_type` or `invalid_type_ref` diagnostic was produced.
- Validator: recursive named model reference is valid for the tested case. Validation returned `error_count: 0` and `warning_count: 0`.
- Renderer: no infinite expansion occurred. The generated model markdown rendered the recursive field as a named reference:
  - `| parent | object_ref | recursive named model reference |`

### Follow-up

No implementation task is required for this specific recursive named model reference behavior.

Optional follow-up: add a focused regression fixture or Go test for `object_ref.parent: object_ref`, because the behavior is supported but not explicitly pinned by a dedicated self-recursive named-model test.

### Workspace note

Final `git status --short` from Codex still showed only pre-existing dirty files outside this task scope plus the already expected V01-TASK-DATA-015-06 update:

- `docs/spec/design-records-mcp/schema.md`
- `docs/tasks/data/TASK-DATA-015-06-review-spec-and-task-split-for-recursive-reference-boundary.md`
- `tmp.py`
