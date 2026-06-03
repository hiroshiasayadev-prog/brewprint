# TASK-DATA-010-03: Tagged union raw model / resolver / validation implementation

- **id**: TASK-DATA-010-03
- **status**: done
- **date**: 2026-06-03
- **work_item**: WORK-DATA-010
- **source_requirement**: REQ-DATA-004
- **estimate**: 2d-3d
- **depends_on**:
  - TASK-DATA-010-02
- **outputs**:
  - Raw YAML / semantic model support for tagged union models
  - Resolver support for tagged union variant field TypeRefs
  - Validation implementation for tagged union model definitions
  - TypeRef compatibility confirmation for tagged union named models
  - Follow-up input for render / catalog / UC-002 fixture work

## Goal

Implement the accepted tagged union minimum in the raw YAML / semantic model / resolver / validation layers, without pulling in render, catalog, UC-002 YAML migration, fixture, or golden regeneration work.

This task follows TASK-DATA-010-02, which accepted and specified the tagged union / discriminator payload contract in ADR-073 and the active specs.

## Work

- Add raw YAML representation for `kind: tagged_union`, `discriminator`, and `variants`.
- Add semantic model representation for tagged union models.
- Preserve variant order and variant field order from YAML.
- Resolve `variants[].fields[].type` as normal TypeRefs.
- Ensure tagged union models are usable as named model TypeRefs.
- Confirm TypeRef compatibility remains nominal for tagged union models.
- Preserve existing `any` wildcard behavior from ADR-060.
- Implement validation for:
  - missing / empty / dot-path `discriminator`
  - missing / empty / non-list `variants`
  - missing / empty / non-string `variants[].tag`
  - duplicate valid `variants[].tag`
  - missing / invalid `variants[].fields`
  - discriminator field redefinition inside variant payload fields
  - disallowed variant field attributes such as `pk`, `fk`, `unique`
  - duplicate variant payload field names
  - invalid / unresolved variant field TypeRefs
  - disallowed model fields for `kind: tagged_union`, such as `fields`, `element`, `value`, `values`
- Add or update unit tests for the above behavior.
- Record follow-up input for render / catalog / UC-002 fixture work.

## Included Scope

- Parser / raw YAML structs.
- Semantic model structs.
- Resolver changes needed for variant field TypeRefs.
- Validation changes and tests for tagged union definitions.
- TypeRef compatibility tests for nominal tagged union model behavior.
- Diagnostic behavior for:
  - `invalid_tagged_union_model`
  - `duplicate_variant_tag`
  - `invalid_variant_field`
  - reuse of `duplicate_model_field`
  - reuse of `invalid_type_ref`
  - reuse of `unresolved_field_type`

## Excluded Scope

- Model-file renderer implementation for tagged union discriminator / variants.
- Model catalog implementation for tagged union display / filtering.
- UC-002 YAML migration.
- Fixture / golden regeneration.
- JSON Schema / MCP schema generation output policy.
- Runtime MCP request / response payload validation.
- Untagged union / general oneOf / scalar union.
- External discriminator / adjacent discriminator / discriminator path support.
- DAG asset TypeRef hint from ADR-074.
- MCP semantic identity / state machine identity from ADR-078 / ADR-079 / ADR-080.
- UC-002 duplicate task QID / unresolved flow task repair.
- Broad remaining UC-002 notes-retreat cleanup.
- Reopening M15, WORK-DATA-001, WORK-DATA-002, WORK-DATA-003, or WORK-DATA-004.

## Done Condition

- Tagged union models can be parsed into raw and semantic model representations.
- Tagged union variant field TypeRefs are resolved consistently with struct field TypeRefs.
- Tagged union models can be referenced as named model TypeRefs.
- TypeRef compatibility treats tagged union models nominally, with existing `any` wildcard behavior preserved.
- Tagged union model validation emits the accepted diagnostics without ambiguous ownership.
- Tests cover representative valid and invalid tagged union model definitions.
- No render / catalog / UC-002 YAML / fixture / golden changes are performed in this task.
- Follow-up scope for render / catalog / UC-002 fixtures is recorded.

## Verification

- Run targeted Go tests for raw YAML / resolver / validation / TypeRef compatibility.
- Run any relevant design-record metadata validation.
- Confirm no render output, fixture, golden, or UC-002 YAML file is changed.
- Confirm excluded scope remains excluded.
- Confirm diagnostic chaining does not duplicate TypeRef compatibility diagnostics when TypeRef resolution already failed.

## Evidence
### Files changed

- `internal/rawyaml/model.go` — added `Discriminator string`, `Variants []ModelVariant`, `ModelVariant struct`
- `internal/semantic/model.go` — added `Discriminator string`, `Variants []ModelVariant`, `ModelVariant struct`
- `internal/resolve/builder.go` — added variant building in `buildModel`; variant field TypeRefs built via `mustBuildTypeRef`
- `internal/resolve/type_ref.go` — added variant field TypeRef scope resolution in `resolveScopedTypeRefs`
- `internal/resolve/validation.go` — added `invalid_tagged_union_model`, `duplicate_variant_tag`, `invalid_variant_field` diagnostic constants; added `tagged_union` to `validModelKinds`; added `validateTaggedUnionModel`; called from `validateModelDefinitions`
- `internal/resolve/tagged_union_test.go` — new file with full valid/invalid test coverage

### Commands run

```
go test -count=1 ./internal/rawyaml ./internal/resolve      # PASS
go test -count=1 ./internal/query ./internal/render/model ./internal/render/er  # PASS
go test -count=1 ./...    # all targeted packages PASS
```

Pre-existing failure in `internal/designrecords` (`TestReplaceNamedSectionSpacingPreservation`) is from `authoring_test.go` already modified before this task; confirmed by stash test.

### Implemented behavior

- `kind: tagged_union` accepted by resolver and validator
- Raw YAML: `rawyaml.Model` accepts `discriminator` and `variants`; `rawyaml.ModelVariant` has `tag` and `fields` (nil = absent key, non-nil empty = payload-less)
- Semantic model: `semantic.Model` gains `Discriminator string` and `Variants []ModelVariant`; variant fields reuse `semantic.ModelField`
- Builder maps raw variants to semantic variants; variant field TypeRefs built via `mustBuildTypeRef`
- TypeRef resolver: variant field TypeRefs resolved via `resolveScopedTypeRef` same as struct fields
- TypeRef compatibility: nominal via existing `TypeRefNamedModel` path; no new code needed
- `any` wildcard: unchanged

### Diagnostic behavior

| diagnostic | trigger |
|---|---|
| `invalid_tagged_union_model` | missing/empty/dot-path discriminator; missing/empty variants; missing/empty variant tag; disallowed `fields`/`element`/`value`/`values` on model |
| `duplicate_variant_tag` | duplicate valid tag within same tagged union |
| `invalid_variant_field` | missing `fields` key; discriminator redefined in payload; `pk`/`fk`/`unique`; missing `name` or `type` |
| `duplicate_model_field` | reused for duplicate variant payload field names |
| `invalid_type_ref` | reused for malformed variant field TypeRef syntax |
| `unresolved_field_type` | reused for unresolved named model in variant field TypeRef |

### Out of scope preserved

- No render file changed
- No model catalog changed
- No UC-002 YAML changed
- No fixture or golden file changed

### Follow-up recommendation for TASK-DATA-010-04

- Model-file renderer: display tagged union discriminator and variants
- DAG private model rendering: show `tagged_union` kind
- Model catalog: list/filter tagged union models per ADR-072
- UC-002 YAML migration: migrate `analyze_impact_change` from `any + note` to `kind: tagged_union`
- Fixture and golden regeneration for affected UC-002 model files
- Verification pass: confirm `go test ./...` fully green after render/catalog/fixture changes
