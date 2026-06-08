# V01-TASK-DATA-015-04: Clean up UC-002 recursive and untagged-like surfaces

- **id**: V01-TASK-DATA-015-04
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-DATA-015
- **source_requirement**: V01-REQ-DATA-008
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-DATA-015-03
  - V01-TASK-DATA-015-06
- **outputs**:
  - UC-002 cleanup decision for N-044 recursive `object_ref.parent`
  - UC-002 cleanup decision for N-009 untagged-like related list
  - Fixture / render update input if migration is selected

## Goal

Apply the accepted recursive / untagged-union boundary to UC-002 surfaces after runtime behavior is understood.

## Work

- Review N-044 and decide whether `object_ref.parent` should be represented as a recursive named model reference in UC-002 YAML.
- Review N-009 and decide whether the untagged-like related list remains `any + note` / prose or becomes a tagged union envelope model.
- Keep untagged union / general `oneOf` unsupported.
- Identify any fixture / render regeneration required if YAML migration is selected.

## Included Scope

- UC-002 cleanup decision and possible YAML migration input.
- Tagged union envelope decision for untagged-like surfaces.
- Traceability to `V01-TASK-DATA-015-01` and `V01-TASK-DATA-015-02`.

## Excluded Scope

- Introducing untagged union / general `oneOf`.
- Broadening V01-ADR-073 or V01-WORK-DATA-010.
- Runtime implementation work.
- Golden regeneration unless selected as a later task.

## Done condition

- N-044 is either migrated, explicitly deferred, or marked no-action with rationale.
- N-009 is either modeled via tagged union envelope, explicitly left opaque, deferred, or marked no-action with rationale.
- Follow-up fixture / render tasks are identified if needed.

## Verification

- Validate affected records after update.
- If YAML is changed, run focused validate / render checks in the later implementation or cleanup task.

## Evidence
- N-044 migrated `object_ref.parent` from `any + note` to recursive named model reference `object_ref` in `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/object_ref.yaml`.
- N-009 migrated `diagnostic.related` from opaque `any + note` to `list<diagnostic_related>` in `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/diagnostic.yaml`.
- Added `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/diagnostic_related.yaml` as a `kind: tagged_union` envelope with discriminator `kind` and `source_location` / `object_ref` variants.
- No untagged union / general `oneOf` / `anyOf` / scalar union support was introduced, and V01-ADR-073 was not broadened.
- Validation command: `go run ./cmd/brewprint validate --yaml-root docs/uc/002-brewprint-self-hosting/yaml --format json` -> passed with `error_count: 0`, `warning_count: 0`.
- Render command: `go run ./cmd/brewprint render --yaml-root docs/uc/002-brewprint-self-hosting/yaml --out docs/uc/002-brewprint-self-hosting/render --clean` -> passed, rendered 47 files.
- Focused tests: `go test ./internal/resolve ./internal/render/model ./cmd/brewprint` -> passed for all three packages.
- Render output changed under `docs/uc/002-brewprint-self-hosting/render`: current renderer output is grouped under `render/mcp/`; generated model pages include `model-object_ref.md`, `model-diagnostic.md`, and new `model-diagnostic_related.md`. The render run also deleted the stale tracked `render/yaml/` pages.
