# TASK-DATA-014-02: Update MCP selector matrix and object-dependent vocabulary specs

- **id**: TASK-DATA-014-02
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-DATA-014
- **source_requirement**: REQ-DATA-007
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-DATA-014-01
- **outputs**:
  - Updated MCP schema selector support matrix contract
  - Updated MCP tool specs for selector support and unsupported behavior
  - Updated MCP error documentation if needed

## Goal

Reflect the `TASK-DATA-014-01` boundary decision into the MCP schema and relevant MCP tool specs.

This task makes selector support matrices, object-dependent `kind` vocabulary, and unsupported / limited selector behavior explicit in spec text.

## Work

- Update the shared MCP schema so `Object selector`, selector field combinations, selector support matrix meaning, and ObjectRef object/kind vocabulary are explicit.
- Update relevant MCP tool specs so each tool points to the shared selector matrix and states its unsupported / limited behavior.
- Update MCP error documentation if unsupported selector behavior needs common wording.
- Keep this task to spec documentation only. Do not edit UC-002 YAML, fixtures, renderer output, parser, validator, MCP implementation, or tests.

## Done condition

- `docs/spec/mcp/schema.md` documents object-dependent `kind` vocabulary and selector support matrix semantics clearly enough to remove broad UC-002 prose notes later.
- Relevant MCP tool specs explicitly reference or align with the shared selector support matrix.
- Unsupported selector behavior is documented without contradicting existing `analyze_impact` normal-response behavior.
- No YAML, fixture, implementation, or test files are intentionally changed.

## Verification

- Review the changed spec text against `TASK-DATA-014-01` Evidence.
- Confirm candidate coverage for N-020, N-031, N-037, N-040, and N-042.
- Confirm no unrelated DATA DSL, tagged union, recursive/untagged union, numeric/default behavior, DAG TypeRef hint, or MCP identity scope is introduced.
- Run Design Records MCP validation for this task and `WORK-DATA-014`.

## Evidence

Verdict: PASS for MCP selector matrix / object-dependent vocabulary spec reflection.

Files updated:

- `docs/spec/mcp/schema.md`
  - Updated `last_updated` to `2026-06-07`.
  - Clarified `Object selector.kind` is object-dependent and `kind_mismatch` is returned when the resolved kind differs.
  - Added selector error behavior for `invalid_selector`, `not_found`, `ambiguous`, and `kind_mismatch`.
  - Added shared object-dependent kind vocabulary for `node`, `view`, `transition`, `asset`, `field`, `file`, and `primitive`.
  - Clarified the vocabulary is MCP runtime selector / response contract, not a DATA DSL dependent enum extension.
  - Clarified selector support matrix cell semantics for `yes` / `no` / `limited`.
  - Clarified default unsupported-selector behavior: most selector tools return `unsupported_object`; `analyze_impact` returns a normal response with `unsupported_selector` diagnostic.
- `docs/spec/mcp/tools/get-signature.md`
  - Updated `last_updated` to `2026-06-07`.
  - Added pointer to shared selector support matrix.
  - Clarified `no` matrix entries return `unsupported_object` for `get_signature`.
  - Clarified `view: api_table` route listing is `list_endpoints` responsibility.
- `docs/spec/mcp/tools/list-objects.md`
  - Updated `last_updated` to `2026-06-07`.
  - Clarified listable objects are `node`, `view`, `transition`, and `field`.
  - Clarified `asset`, `file`, and `primitive` are not list_objects enumeration targets.
  - Clarified `kind` is an object-dependent filter using the shared vocabulary.
  - Clarified `object` omitted behavior and invalid filter handling.
- `docs/spec/mcp/tools/get-references.md`
  - Added pointer to shared selector support matrix.
  - Clarified `no` matrix entries return `unsupported_object` for `get_references`.
  - Clarified `limited` matrix entries may limit returned references.
- `docs/spec/mcp/tools/get-reference-tree.md`
  - Added pointer to shared selector support matrix.
  - Clarified `no` matrix entries return `unsupported_object` for `get_reference_tree`.
  - Clarified `limited` selector behavior for traversal roots / edge expansion.
  - Aligned field wording with `field` / `model_field` vocabulary.
- `docs/spec/mcp/tools/analyze-impact.md`
  - Updated `last_updated` to `2026-06-07`.
  - Added pointer to shared selector support matrix.
  - Clarified unsupported selector remains normal response + `unsupported_selector` diagnostic, not tool error.
  - Aligned field wording with `field` / `model_field` vocabulary.
- `docs/spec/mcp/tools/inspect.md`
  - Added pointer to shared selector support matrix.
  - Clarified `no` matrix entries return `unsupported_object` for `inspect`.
  - Clarified `limited` matrix entries may limit returned context.
- `docs/spec/mcp/errors.md`
  - Added selector support behavior summary for malformed selector, not found, ambiguous resolution, kind mismatch, unsupported object, and the `analyze_impact` exception.

Candidate coverage:

| ID | Coverage |
|---|---|
| N-020 | `get_signature_request.selector` now points to shared selector support matrix and states unsupported behavior. |
| N-031 | `list_objects_request.kind` now has object-dependent filter semantics and shared vocabulary ownership. |
| N-037 | `object_selector.kind` now states object-dependent vocabulary and `kind_mismatch` behavior. |
| N-040 | `object_selector` now has selector resolution / unsupported matrix behavior and tool-specific handling. |
| N-042 | `ObjectRef.kind` now has shared object-dependent response vocabulary. |

Verification performed:

- Compared the changes against `TASK-DATA-014-01` Evidence.
- Confirmed the changes stay in MCP schema/tool/error specs and do not introduce DATA DSL dependent enum support.
- Confirmed no UC-002 YAML, fixture, renderer output, parser, validator, MCP implementation, or test files were intentionally changed.
- No Go tests were run because this task changed only Markdown specs.

Review repair:

- Addressed Codex review finding for `get_references` limited selector behavior by adding a concrete `Selector support` section.
- Addressed Codex review finding for `inspect` ObjectRef examples by replacing invalid `object: model_field` examples with `object: field`.
- Addressed Codex review finding for stale `WORK-DATA-014` task flow by listing current tasks and remaining split.
- Repository-scope note: `tmp.py` is outside this task's intended scope and was not modified by this repair. Untracked task files are expected workflow artifact outputs until staged/committed.
