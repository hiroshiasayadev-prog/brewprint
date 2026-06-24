# PRODUCT-TASK-NAMESPACE-001-08: Perform independent namespace and ID-as-ref boundary review

- **id**: PRODUCT-TASK-NAMESPACE-001-08
- **status**: completed
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-NAMESPACE-001
- **source_requirement**: V01-REQ-PRODUCT-001
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-NAMESPACE-001-06
  - PRODUCT-TASK-NAMESPACE-001-07
- **outputs**:

## Goal

Obtain a review of the restructured namespace-model spec set from an independent reviewer — someone who did not implement T02–T07 — and classify all findings before passing to T09.

## Work

| area | required work |
|---|---|
| reviewer role | The reviewer must not have implemented T02–T07. The reviewer reads the actual repository state after T06 and T07 complete. |
| files to review | Every file changed by T02–T07. T06 and T07 may modify files beyond the minimum list below; the reviewer must identify and include all T02–T07 changed files. Minimum required: all files in `product/records/spec/concepts/namespace-model/`; `drmcp/records/spec/design-records-mcp/namespace-scanning.md`; `drmcp/records/spec/design-records-mcp/schema/id-normalization.md`; `product/records/spec/concepts/traceability/artifact-refs.md`; `product/records/spec/concepts/authoring-standards/task-authoring.md`; `product/records/spec/concepts/authoring-standards/work-item-authoring.md`. |
| review criteria | The reviewer evaluates against all eight areas: (1) **Ownership leakage** — any DRMCP parser, normalization, or scanning behavior remaining in PRODUCT namespace-model files. (2) **Compatibility completeness** — all six V01-* legacy families enumerated; spec identity not redefined; legacy forms stated as non-canonical for new records. (3) **Stale refs** — any remaining machine-readable ref to temporary topic IDs in PRODUCT or DRMCP specs. (4) **Semantic correctness** — sequence allocation scopes present and correct; task work-sequence inheritance stated; subdomain exclusion from public IDs stated; complete public ID stated as canonical record ID-as-ref; bare forms stated as non-canonical external refs. (5) **Duplicate normative authority** — no PRODUCT spec cites DRMCP normalization as the authority for public identity semantics. (6) **Historical ownership versus attribution** — V01-ADR-096 ownership judgment not conflated with issued identity or effective attribution. (7) **Scope drift** — no content outside the Work Item boundary has been changed without justification recorded in Evidence. (8) **Work Item evidence and closure readiness** — `PRODUCT-WORK-NAMESPACE-001.impact_refs` uses new canonical refs; V01 reciprocal-link exception is documented in Evidence. |
| finding classification | Classify each finding as `must-fix` (spec is incorrect, ambiguous, or violates an established contract) or `advisory` (improvement worth considering but not required for closure). |
| report | The reviewer returns a written report listing each finding with its classification, the quoted text, and the file and section it appears in. |
| cross-check | Before recording a finding in Evidence, the implementer verifies the quoted text against the current file. Findings citing text not present in the current files are not accepted. |

## Done condition

| item | done when |
|---|---|
| review completed | Review has been run by an independent reviewer who did not implement T02–T07. |
| findings classified | All findings are classified as `must-fix` or `advisory`. |
| findings verified | Each must-fix finding has been verified against the current file before being recorded in Evidence. |

## Verification

- Confirm the reviewer did not implement T02–T07.
- Confirm no finding was accepted without cross-checking the quoted content against the actual file.

## Evidence

**Reviewer:** External independent reviewer (not the T02–T07 implementer). Review run post-compaction; reviewer read actual repository state after T07 completion.

**Quoted text cross-check:** All T08 finding quotes verified against current files before recording.

**Findings:**

| id | classification | file | section | quoted text |
|---|---|---|---|---|
| M1 | must-fix | `namespace-model/artifact-id-grammar.md` | What this is | "Defines the canonical artifact ID grammar for new records:" |
| M1 | must-fix | `namespace-model/legacy-id-compatibility.md` | Retention policy | "New records must use the grammar defined in `spec:product.concepts.namespace_model.artifact_id_grammar`." |
| M1 | must-fix | `namespace-model/existing-artifacts.md` | New-artifact ownership | "For new artifacts, use the `<APP_NAMESPACE>-...` form" |
| M2 | must-fix | `namespace-model/existing-artifacts.md` | Effective attribution | Attribution table omits `PRODUCT` domain group; bare forms used instead of `V01-*` prefix; column labeled "existing artifacts" implies exhaustive family list |
| A1 | advisory | `authoring-standards/*.md` (all 5 guides) | ID grammar section | Guides repeat canonical grammar as independent MUST rules alongside `artifact_id_grammar` citation, creating drift risk |
| A2 | advisory | `PRODUCT-WORK-NAMESPACE-001` | metadata | `status: not_started` no longer describes reality with T01–T08 completed |

**Checks passed (no findings):**
- No deleted temporary refs (`v1_id_grammar`, `v1_namespace_algorithm`, `v2_grammar`) remain in active PRODUCT or DRMCP specs.
- No machine-readable metadata refs to stale topics remain.
- No DRMCP normalization or scanning behavior in PRODUCT namespace-model files.
- All six V01-* compatibility families enumerated.
- Legacy spec identity remains compatibility-only; path-derived `spec:` identity preserved.
- No PRODUCT spec cites DRMCP ID normalization as authority for public identity semantics.
- Historical PRODUCT ownership and effective attribution separated without contradiction.
- `PRODUCT-WORK-NAMESPACE-001.impact_refs` uses new stable refs.
- V01-REQ-PRODUCT-001 reciprocal-link exception documented in Work Item Evidence.

**Scope drift note:** Reviewer checked all files enumerated in T02–T07 Evidence via filesystem rather than raw git diff. No out-of-scope changes found.
