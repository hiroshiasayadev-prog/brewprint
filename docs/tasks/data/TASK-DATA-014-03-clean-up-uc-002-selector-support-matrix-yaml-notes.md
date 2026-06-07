# TASK-DATA-014-03: Clean up UC-002 selector support matrix YAML notes

- **id**: TASK-DATA-014-03
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-DATA-014
- **source_requirement**: REQ-DATA-007
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-DATA-014-02
- **outputs**:
  - UC-002 YAML notes aligned to MCP selector support matrix specs
  - Evidence for N-020 N-031 N-037 N-040 N-042 YAML cleanup

## Goal

Clean up UC-002 YAML notes for the selector support matrix and object-dependent vocabulary bucket after the MCP schema/tool specs have been made canonical.

## Work

- Update UC-002 YAML notes for selector-related request and response models so they point to the MCP schema/tool specs instead of carrying broad prose notes.
- Cover the `REQ-DATA-007` candidates N-020, N-031, N-037, N-040, and N-042.
- Keep this task to YAML note cleanup and workflow evidence only.
- Do not change parser, renderer, validator, MCP implementation, generated renders, fixtures, or golden outputs unless the YAML note cleanup requires it.

## Done condition

- Relevant UC-002 YAML notes refer to the canonical MCP selector support matrix / object-dependent vocabulary specs.
- The YAML notes do not imply a DATA DSL dependent enum or selector-matrix validation feature.
- Candidate coverage for N-020, N-031, N-037, N-040, and N-042 is recorded in Evidence.

## Verification

- Review changed YAML notes against `docs/spec/mcp/schema.md`, `docs/spec/mcp/tools/get-signature.md`, `docs/spec/mcp/tools/list-objects.md`, `docs/spec/mcp/tools/get-references.md`, `docs/spec/mcp/tools/get-reference-tree.md`, `docs/spec/mcp/tools/analyze-impact.md`, and `docs/spec/mcp/tools/inspect.md`.
- Run Design Records MCP validation for this task and `WORK-DATA-014`.
- No Go tests are required unless non-document implementation files are changed.

## Evidence

Verdict: PASS.

Files updated:

- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/common.yaml`
  - Updated `object_selector.kind` note to point to `docs/spec/mcp/schema.md` §1.1 object-dependent kind vocabulary.
  - Updated `object_selector` model note to point selector field combination / support matrix ownership to MCP schema and tool specs.
  - Clarified this YAML does not introduce DATA DSL dependent enum or selector matrix validation.
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/object_ref.yaml`
  - Updated `object_ref.kind` note to point to MCP object-dependent kind vocabulary.
  - Clarified ObjectRef object/kind combination ownership and recursive `parent` limitation.
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_signature_request.yaml`
  - Updated selector note to point to MCP schema §1.1–§1.2 and `get-signature.md` §2.
  - Added unsupported selector behavior summary.
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_objects_request.yaml`
  - Updated `kind` note to point to `list-objects.md` §2 and MCP schema §1.1.
  - Removed stale wording that held object-dependent vocabulary only as broad REQ/WORK note.
  - Clarified no DATA DSL dependent enum is introduced.
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_references_request.yaml`
  - Added selector support pointer to MCP schema §1.2 and `get-references.md` §5.
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_request.yaml`
  - Added selector support pointer to MCP schema §1.2 and `get-reference-tree.md` §7.
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_request.yaml`
  - Added selector support pointer to MCP schema §1.2 and `analyze-impact.md` §13.
  - Clarified analyze_impact unsupported selector behavior as normal response plus `unsupported_selector` diagnostic.
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/inspect_request.yaml`
  - Added selector support pointer to MCP schema §1.2 and `inspect.md` §2.
- `docs/work-items/data/WORK-DATA-014-selector-support-matrix-and-object-dependent-vocabulary.md`
  - Updated Task Flow to include `TASK-DATA-014-03` without duplicating task completion status from task artifacts.

Candidate coverage:

| ID | Coverage |
|---|---|
| N-020 | `get_signature_request.selector` now points to canonical MCP selector support specs and unsupported behavior. |
| N-031 | `list_objects_request.kind` now points to canonical object-dependent kind filter semantics. |
| N-037 | `object_selector.kind` now points to canonical object-dependent vocabulary and kind matching contract. |
| N-040 | `object_selector` now points to canonical selector field combination / tool support matrix ownership. |
| N-042 | `object_ref.kind` now points to canonical ObjectRef object-dependent vocabulary. |

Verification performed:

- Reviewed changed YAML notes against MCP schema/tool specs updated by `TASK-DATA-014-02`.
- Confirmed YAML notes point to canonical MCP schema/tool specs instead of owning broad support matrix prose.
- Confirmed no parser, renderer, validator, MCP implementation, generated render, fixture, or golden output file was intentionally changed.
- No Go tests were run because this task changed only YAML notes and workflow docs.
