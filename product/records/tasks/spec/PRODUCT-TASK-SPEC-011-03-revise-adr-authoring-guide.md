# PRODUCT-TASK-SPEC-011-03: Revise ADR authoring guide

- **id**: PRODUCT-TASK-SPEC-011-03
- **status**: done
- **date**: 2026-06-23
- **work_item**: PRODUCT-WORK-SPEC-011
- **source_requirement**: PRODUCT-REQ-SPEC-002
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-011-01
- **outputs**:
  - `product/records/spec/concepts/authoring-standards/adr-authoring.md`

## Goal

Revise the ADR authoring guide as the canonical per-artifact example for PRODUCT authoring standards.

## Work

- Apply the common per-artifact guide structure.
- Replace Japanese ADR body headings with canonical English headings.
- Separate create input, partial update input, and persisted metadata requirements.
- Define `date` as the decision-validity date.
- Use abstract v2 ID forms and the `new` placeholder.
- Separate author-supplied values from generated H1 and path values.
- Replace concrete DRMCP operation details with author-facing interface requirements.
- Reference the shared authoring boundary instead of duplicating the complete artifact ownership matrix.
- Remove current DRMCP operating status and V01-specific primary examples.

## Done condition

| item | done when |
|---|---|
| Common shape applied | The guide uses the common section structure defined by PRODUCT-WORK-SPEC-011. |
| English headings canonical | ADR body headings are `Context`, `Decision`, `Rationale`, `Rejected alternatives`, `Consequences`, and `Evidence`. |
| Metadata states separated | The metadata table distinguishes create, partial update, and persisted requirements. |
| Date semantics defined | `date` records when the decision became valid and is not changed for editorial updates. |
| Authoring boundary respected | PRODUCT semantics remain canonical and concrete DRMCP contracts remain outside the guide. |
| Legacy operational text removed | The guide contains no current DRMCP operating-status statement or V01 primary authoring form. |

## Verification

- Confirm the guide headings match the common per-artifact structure.
- Confirm all prescribed ADR body headings use English.
- Confirm metadata requiredness is separated by state.
- Confirm the authoring interface identifies generated ID, H1, and path values.
- Confirm old Japanese headings, `V01-ADR`, `Current operating mode`, and `superseded_by` are absent.

## Evidence

- `product/records/spec/concepts/authoring-standards/adr-authoring.md` was rewritten on 2026-06-23.
- The guide now uses `Status lifecycle`, `Kind-specific authoring rules`, and `Authoring interface requirements`.
- The body schema defines `Context`, `Decision`, `Rationale`, `Rejected alternatives`, `Consequences`, and `Evidence`.
- The metadata schema distinguishes create input, partial update, and persisted ADR state.
- The guide defines `date` as the decision-validity date and excludes automatic editorial-date updates.
- The guide uses `<APP>-ADR-<DOMAIN>-<NNN>` and `<APP>-ADR-<DOMAIN>-new` as primary forms.
- The guide references `spec:product.concepts.authoring_standards.artifact_boundary`; creating that shared spec remains PRODUCT-TASK-SPEC-011-02 scope.
- Verification search returned zero matches for the former Japanese headings, `V01-ADR`, `Current operating mode`, `non-operational`, `Status enum`, `Responsibility boundary`, `DRMCP authoring transaction`, and `superseded_by`.
