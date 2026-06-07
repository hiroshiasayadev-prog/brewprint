# TASK-DATA-015-02: Update TypeRef spec for recursive named references and unsupported untagged unions

- **id**: TASK-DATA-015-02
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-DATA-015
- **source_requirement**: REQ-DATA-008
- **estimate**: 0.5d
- **depends_on**:
  - TASK-DATA-015-01
- **outputs**:
  - TypeRef spec update for recursive named model references
  - Explicit spec statement that untagged union / general oneOf / anyOf / scalar union is unsupported
  - Follow-up implementation and UC-002 cleanup input

## Goal

Reflect the `TASK-DATA-015-01` contract boundary decision into the TypeRef specification.

## Work

- Update `docs/spec/type-ref.md` to document that recursive references are allowed only through named model TypeRef references.
- Document that inline recursive shapes are not introduced.
- Document that untagged union / general `oneOf` / `anyOf` / scalar union syntax is unsupported.
- Document that untagged-like surfaces requiring machine-readable schema should use tagged union envelope models.
- Identify follow-up implementation / YAML cleanup tasks without performing them in this task.

## Included Scope

- Documentation-only spec update.
- Traceability back to `REQ-DATA-008`, `WORK-DATA-015`, and `TASK-DATA-015-01`.

## Excluded Scope

- Parser, resolver, renderer, validator, MCP, or fixture implementation changes.
- UC-002 YAML migration.
- Golden regeneration.
- Broadening ADR-073 or WORK-DATA-010.

## Done condition

- `docs/spec/type-ref.md` reflects the recursive named model reference boundary.
- `docs/spec/type-ref.md` explicitly rejects untagged union / general `oneOf` support.
- Follow-up work input is recorded in task evidence.

## Verification

- Review the edited spec section for consistency with ADR-073 and `TASK-DATA-015-01`.
- Validate the task record after update.

## Evidence

Completed on 2026-06-07.

### Files changed

- `docs/spec/type-ref.md`

### Spec updates

- Updated front matter `last_updated` to 2026-06-07.
- Added an explicit TypeRef syntax boundary: inline `union<...>` / `oneOf<...>` / `anyOf<...>` / scalar union syntax is unsupported.
- Added recursive named model reference semantics under named model TypeRef.
- Clarified that recursive structures must use named model TypeRef references and must not introduce inline recursive shapes.
- Clarified that resolver / renderer follow-up must avoid unbounded recursive expansion.
- Extended tagged union compatibility text to state that untagged union / general oneOf is unsupported as both TypeRef syntax and model kind.
- Added a tagged union envelope replacement pattern for `SourceLocation | ObjectRef` style surfaces.

### Follow-up work input

- Implementation follow-up is needed only if current resolver / validator / renderer behavior rejects recursive named model references or expands recursive references incorrectly.
- UC-002 cleanup for N-044 should wait until implementation behavior is confirmed.
- UC-002 cleanup for N-009 should decide between remaining `any + note` / prose and introducing an explicit tagged union envelope model.

### Verification

- `validate_records` for `TASK-DATA-015-02`: PASS.
- `validate_records` for `TASK-DATA-015-01` through `TASK-DATA-015-02`: PASS.
- `validate_records` for `WORK-DATA-015`: PASS.
- Repository-wide spec validation was checked and still reports pre-existing unrelated `missing_section_target` diagnostics under `docs/spec/concepts/namespace-model/index.md`; no new TypeRef-specific diagnostic was reported by that check.
