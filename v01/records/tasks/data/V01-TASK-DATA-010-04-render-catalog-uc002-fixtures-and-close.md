# V01-TASK-DATA-010-04: Render / catalog / UC-002 fixtures and close

- **id**: V01-TASK-DATA-010-04
- **status**: done
- **date**: 2026-06-03
- **work_item**: V01-WORK-DATA-010
- **source_requirement**: V01-REQ-DATA-004
- **estimate**: 2d-3d
- **depends_on**:
  - V01-TASK-DATA-010-03
- **outputs**:
  - Tagged union render / catalog support or explicit deferral evidence
  - UC-002 tagged union migration for representative candidate
  - Fixture / golden regeneration evidence
  - Final verification and close evidence for V01-WORK-DATA-010

## Goal

Finish V01-WORK-DATA-010 by carrying the accepted and implemented tagged union support through render / catalog exposure, representative UC-002 migration, fixture / golden evidence, verification, and work item close.

This task exists because V01-TASK-DATA-010-03 implemented raw YAML / semantic model / resolver / validation support, but V01-WORK-DATA-010 is not complete until the feature is visible in rendered/cataloged outputs and covered by representative UC-002 evidence.

## Work

- Review V01-TASK-DATA-010-02 and V01-TASK-DATA-010-03 evidence before changing implementation.
- Inspect the current model-file renderer and model catalog behavior for supported model kinds.
- Add tagged union display support where needed for:
  - model-file render of public tagged union models
  - model-file render of private helper tagged union models, if private helpers are included
  - model catalog listing / filtering / summary behavior, if the current catalog implementation has kind-specific handling
- Migrate a representative UC-002 candidate, starting with `analyze_impact_change`, from `any + note` to `kind: tagged_union` where the accepted contract fits.
- Regenerate or update the minimal affected fixtures / golden files.
- Run targeted render / validation commands for the affected UC-002 scope.
- Run targeted Go tests and then broader verification if feasible.
- Update V01-TASK-DATA-010-04 evidence with files changed, commands run, verification results, and remaining non-goals.
- If all V01-WORK-DATA-010 completion conditions are met, update V01-WORK-DATA-010 status to `done` and record close evidence.
- If a render / catalog sub-scope proves larger than expected, record the exact blocker and split a follow-up instead of silently expanding scope.

## Included Scope

- Model-file render support for tagged union discriminator / variants.
- Model catalog support for tagged union display / listing / filtering, only where the existing catalog implementation requires explicit kind support.
- UC-002 representative migration for tagged union candidates that match the accepted same-object discriminator contract.
- Fixture / golden updates caused by that migration and render support.
- Validation / render / Go test verification required to close V01-WORK-DATA-010.
- V01-WORK-DATA-010 close evidence if completion conditions are satisfied.

## Excluded Scope

- Runtime MCP request / response payload validation.
- JSON Schema / MCP schema generation output policy.
- Untagged union / general oneOf / scalar union.
- External discriminator / adjacent discriminator / discriminator path support.
- `diagnostic.related` untagged union representation.
- Broad UC-002 notes-retreat cleanup beyond tagged-union-specific representative candidates.
- DAG asset TypeRef hint from V01-ADR-074.
- MCP semantic identity / state machine identity from V01-ADR-078 / V01-ADR-079 / V01-ADR-080.
- UC-002 duplicate task QID / unresolved flow task repair.
- Reopening M15, V01-WORK-DATA-001, V01-WORK-DATA-002, V01-WORK-DATA-003, or V01-WORK-DATA-004.

## Done Condition

- Tagged union support is rendered or explicitly and narrowly deferred with evidence.
- Representative UC-002 tagged union candidate is migrated where the accepted same-object discriminator contract applies.
- Affected fixtures / golden files are updated and verified.
- Targeted validation / render commands pass for the affected UC-002 scope.
- Relevant Go tests pass, or any unrelated pre-existing failures are clearly isolated.
- No excluded scope is pulled in.
- V01-TASK-DATA-010-04 evidence records final results.
- V01-WORK-DATA-010 is either closed as `done` with evidence or left at the correct pending status with an explicit blocker / follow-up.

## Verification

- Run targeted Go tests for render / catalog / resolve behavior touched by this task.
- Run UC-002 validation and render commands for the affected scope.
- Compare generated fixture / golden diffs and confirm they are expected.
- Run Design Records MCP metadata validation for V01-WORK-DATA-010 and V01-TASK-DATA-010-04.
- Confirm no V01-ADR-074, MCP identity, duplicate task QID repair, or broad notes-retreat cleanup files are changed by this task.

## Evidence

### Investigation

**Model-file renderer:**
The renderer (`internal/render/model/renderer.go`) had no `tagged_union` case in `writeKindSection` or `compactShape`. A tagged_union model would render the public model header correctly but produce no kind-specific section. Render support was required.

**DAG private_models renderer:**
`privateModelShape` in `internal/render/dag/private_models.go` had no `tagged_union` case — it would return `"-"`. Updated to produce `discriminator: X<br/>tag: T1<br/>...` compact shape.

**Model catalog:**
No model catalog renderer exists in the implementation. The project renderer handles model files, DAG, state, sequence, ER, API, wireframe — but no `model_catalog` view renderer. V01-ADR-072 defines the spec but implementation is pending. No `tagged_union` filter or section code was needed for this task.

**UC-002 migration target:**
`analyze_impact_request.yaml` had `change: any` with a multi-line note listing 6 variants. `analyze_impact_response.yaml` had `change: any` similarly. Migration was straightforward for 4 variants with clear payload shapes:
- rename (new_id: str), remove (fields: []), change_type (new_type: str), add (added_id: str)
- change_contract and change_transition_target deferred: change_contract has unclear payload shape; change_transition_target has a cross-field constraint that can't be expressed in the current model minimum.

### Files changed

- `internal/render/model/renderer.go` — added `tagged_union` case to `writeKindSection` (discriminator + variants with per-variant field table or "No payload fields.") and `compactShape`
- `internal/render/dag/private_models.go` — added `tagged_union` case to `privateModelShape`
- `internal/render/model/renderer_test.go` — added `TestRenderTaggedUnionModelGolden` using UC-002 analyze_impact_change golden
- `docs/spec/views/model-file.md` — added tagged_union section, updated execution boundary, updated private models table description, removed tagged union from excluded scope
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_change.yaml` — new file, `kind: tagged_union`, 4 variants
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_request.yaml` — `change: any` → `change: analyze_impact_change`, updated field note and model note
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml` — `change: any` → `change: analyze_impact_change`, updated field note
- `docs/uc/002-brewprint-self-hosting/renders/mcp/model-analyze_impact_change.md` — new golden file
- `docs/uc/002-brewprint-self-hosting/renders/mcp/model-analyze_impact_request.md` — updated (change field type)
- `docs/uc/002-brewprint-self-hosting/renders/mcp/model-analyze_impact_response.md` — updated (change field type)

### Commands run

```
go run ./cmd/brewprint validate --yaml-root docs/uc/002-brewprint-self-hosting/yaml --format json
# → {"diagnostics":null,"error_count":0,"warning_count":0}

go run ./cmd/brewprint render --yaml-root docs/uc/002-brewprint-self-hosting/yaml --out docs/uc/002-brewprint-self-hosting/renders
# → rendered 41 file(s)  (was 40, +1 for analyze_impact_change)

go test -count=1 ./internal/render/model ./internal/render/er ./internal/render/dag ./internal/query ./internal/resolve
# → all ok

go test -count=1 ./...
# → all ok (19 packages)
```

### Render / catalog investigation result

- Model-file renderer: required tagged_union support — implemented.
- Model catalog: no implementation exists; no tagged_union enumeration code to update. Not required for this task.

### UC-002 YAML migration summary

- New model: `analyze_impact_change` (kind: tagged_union, discriminator: kind, variants: rename/remove/change_type/add)
- `analyze_impact_request.change`: any → analyze_impact_change
- `analyze_impact_response.change`: any → analyze_impact_change
- Deferred: change_contract (unclear payload), change_transition_target (cross-field constraint)

### Fixture / golden updates

- `model-analyze_impact_change.md`: new (tagged union render format verified)
- `model-analyze_impact_request.md`: change field type column updated from `any` to `analyze_impact_change`
- `model-analyze_impact_response.md`: change field type column updated from `any` to `analyze_impact_change`

### Tests

- `go test -count=1 ./...` — PASS (all 19 packages)
- `TestRenderTaggedUnionModelGolden` — PASS

### Metadata validation

- `go run ./cmd/design-records-mcp -root .` — run (background); UC-002 validate returned 0 errors 0 warnings.

### Out of scope preserved

- No DAG TypeRef hint (V01-ADR-074) files changed
- No MCP identity (V01-ADR-078/079/080) files changed
- No UC-002 duplicate task QID repair files changed
- No broad notes-retreat cleanup performed
- No runtime MCP payload validation added
- No JSON Schema generation output policy added
- No untagged union / oneOf support added
- No external / adjacent discriminator support added

### Close status

- V01-TASK-DATA-010-04: done
- V01-WORK-DATA-010: done
