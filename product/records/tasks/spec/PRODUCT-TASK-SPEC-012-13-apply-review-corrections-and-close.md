# PRODUCT-TASK-SPEC-012-13: Apply review corrections and close

- **id**: PRODUCT-TASK-SPEC-012-13
- **status**: done
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-SPEC-012
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-012-12
- **outputs**:
  - Required review corrections
  - PRODUCT-WORK-SPEC-012 completion evidence and status

## Goal

Resolve required review findings and close the restructuring with complete evidence.

## Work

- Apply all blocker and must-fix findings from T12.
- Re-run affected validation and stale-ref checks.
- Record the disposition of optional findings.
- Confirm every Work Item completion condition.
- Update `PRODUCT-WORK-SPEC-012` Evidence with exact files, review gates, validation commands, and results.
- Set the Work Item to `done` only after all evidence is complete.
- Do not reopen deferred BPDSL migration or integration design.

## Done condition

- All required T12 findings are resolved.
- Affected validation passes after corrections.
- Optional findings have recorded dispositions.
- `PRODUCT-WORK-SPEC-012` contains complete evidence.
- `PRODUCT-WORK-SPEC-012` status is `done`.
- The temporary BPDSL exit obligation remains visible.

## Verification

- Compare final files with T12 findings.
- Confirm no required finding remains open.
- Confirm Work Item evidence covers every child Task.
- Confirm no active PRODUCT spec remains under `concepts/`.
- Confirm `v01/` remains unchanged.

## Evidence

### 1. Execution summary

| item | value |
|---|---|
| Final status | `done` |
| T12 verdict consumed | `NEEDS REVISION` — 0 BLOCKER, 5 MUST-FIX, 4 ADVISORY |
| Required findings resolved | M1, M2, M3, M4, M5 — all five |
| Optional findings applied | none |
| Files changed by T13 | 14 |

**T13-edited files (explicit list)**:

1. `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-13-apply-review-corrections-and-close.md` — lifecycle status + this Evidence
2. `product/records/work-items/spec/PRODUCT-WORK-SPEC-012-product-spec-semantic-layer-restructuring.md` — lifecycle status + Work Item Evidence
3. `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md` — M1
4. `product/records/spec/design-records/artifact-model/traceability-boundary.md` — M2A and M3D
5. `product/records/spec/design-records/repository-layout/record-discovery-paths.md` — M2B
6. `product/records/spec/design-records/namespace-model/subdomain-model.md` — M3A
7. `product/records/spec/design-records/artifact-model/index.md` — M3B
8. `product/records/spec/design-records/artifact-model/change-and-investigation-flow.md` — M3C
9. `product/records/spec/design-records/traceability/resolve-and-validation.md` — M3E
10. `product/records/spec/brewprint/namespaces/app-namespaces.md` — M4A
11. `product/records/spec/brewprint/compatibility/existing-artifacts.md` — M4B
12. `product/records/spec/design-records/traceability/index.md` — M4C
13. `product/records/spec/design-records/traceability/semantic-ref.md` — M4D
14. `product/records/spec/design-records/namespace-model/app-namespaces.md` — M5

### 2. Required finding dispositions

| finding | file(s) | old wording or defect | new wording or correction | verification |
|---|---|---|---|---|
| M1 — Unresolved active cross-owner ref | `agent-authoring-policy.md` Non-goals and Related specs | `spec:drmcp.design_records_mcp.tools` (two occurrences) | `spec:drmcp.design_records_mcp.tools.overview` | grep confirms zero occurrences of bare `.tools` authority refs in active spec prose |
| M2 — T08 handoff section headings and adjacent owner claims | `traceability-boundary.md`; `record-discovery-paths.md` | `## T08 handoff` heading in both files; "Concrete DRMCP behavior belongs to T08." in traceability-boundary; "handed off to T08" intro in record-discovery-paths; T08 table cells | `## DRMCP boundary`; "belongs to DRMCP app-local specifications."; evergreen intro; canonical-owner table per M2 spec | grep confirms no `## T08 handoff` heading remains in any spec |
| M3 — Task vehicles as current owners in Design Records specs | `subdomain-model.md`; `artifact-model/index.md`; `change-and-investigation-flow.md`; `traceability-boundary.md`; `resolve-and-validation.md` | `PRODUCT-TASK-SPEC-012-08 handoff to DRMCP app-local specifications.`; "T09 classification."; "belongs to T08." (multiple); "T08 handoff."; "DRMCP, T08 handoff."; "belong to T08." | `DRMCP app-local specifications.`; "non-canonical staging under `spec:product.bpdsl`"; "DRMCP app-local specifications." (all) | grep confirms no `belongs to T08`, `T08 handoff`, or `PRODUCT-TASK-SPEC-012-08 handoff` in active spec prose |
| M4 — Task vehicles in Brewprint and traceability current wording | `brewprint/namespaces/app-namespaces.md`; `brewprint/compatibility/existing-artifacts.md`; `traceability/index.md`; `traceability/semantic-ref.md` | "T08 handoff to DRMCP app-local specifications."; "T09 handoff to BPDSL migration review or app-local specifications."; "removed by T05" in Topics summary and What this is; `## Obsolete assumptions removed by T05` heading | `DRMCP app-local specifications.`; `BPDSL app-local specifications.`; "dispositions for obsolete semantic-ref assumptions"; `## Obsolete assumptions` | grep confirms no `T08 handoff`, `T09 handoff`, or `removed by T05` in active spec prose |
| M5 — Malformed Markdown table separators | `design-records/namespace-model/app-namespaces.md` | `\|---\|---\|---\|` under two-column headers in `## Current contract` and `## Related specs` | `\|---\|---\|` (two-column) | grep confirms no three-column separator in `app-namespaces.md`; full file read confirms both separators fixed |

### 3. Optional finding dispositions

| advisory | disposition |
|---|---|
| A1 | Not applied. Illustrative path/ref example in `spec-authoring.md` and `spec-id-as-ref.md`; non-blocking and outside required T13 scope. |
| A2 | No action. Intentional obsolete hyphen-form illustration in `spec-id-as-ref.md` Boundary table. |
| A3 | Not applied. Provenance-only migration wording is non-blocking and remains readable as evidence. |
| A4 | No action. DRMCP-side authoring consideration; outside `PRODUCT-WORK-SPEC-012`. |

No follow-up records created for A1–A4.

### 4. Validation results

**Strict validator commands** (run from repo root):

```powershell
python -X utf8 product/src/tools/validate_spec.py product/records/spec --strict --no-color
python -X utf8 product/src/tools/validate_spec.py drmcp/records/spec/design-records-mcp --strict --no-color
python -X utf8 product/src/tools/validate_spec.py bpdsl/records/spec --strict --no-color
```

| root | expected files | output | exit |
|---|---|---|---|
| `product/records/spec` | 47 | `[strict]  All 47 file(s) OK.` | 0 |
| `drmcp/records/spec/design-records-mcp` | 30 | `[strict]  All 30 file(s) OK.` | 0 |
| `bpdsl/records/spec` | 37 | `[strict]  All 37 file(s) OK.` | 0 |

All three commands exit 0. Total: 114 spec files, zero errors.

**Combined graph results** (inline temporary script, deleted after run):

| check | result |
|---|---|
| Indexed specs | 114 (47 + 30 + 37) — pass |
| Duplicate IDs | 0 — pass |
| Unresolved parents | 0 — pass |
| Parent/Topics mismatches | 0 — pass |
| Unresolved Topics refs | 0 — pass |
| Unresolved active Related specs refs | 0 active — pass |
| Unresolved active body refs (after classification) | 0 — pass |

**False-positive classifications (5 total)**:

| ref | file | classification |
|---|---|---|
| `spec:drmcp.design_records_mcp.tools` | `spec-authoring.md` line 43 | A1 illustrative path-derivation example |
| `spec:drmcp.design_records_mcp.tools` | `spec-id-as-ref.md` line 53 | A1 illustrative path-derivation example |
| `spec:trace.semantic` | `spec-id-as-ref.md` line 88 | A2 partial regex match inside hyphen-form backtick example |
| `spec:trace.semantic` | `drmcp/tools/resolve-reference.md` | A2 same partial regex match, illustrative example |
| `spec:product.concepts.repository_layout` | `bpdsl/design-flow.md` line 97 | T11-accepted historical source-map preservation row |

### 5. Final-tree and BPDSL checks

**Top-level product/records/spec/ tree**:

```text
product/records/spec/
  index.md              (spec:product)
  design-records/       (spec:product.design_records)
  brewprint/            (spec:product.brewprint)
  bpdsl/                (spec:product.bpdsl — non-canonical staging)
  concepts/             (empty shell — no .md files)
```

| check | result |
|---|---|
| No `.md` files under `product/records/spec/concepts/` | confirmed — empty shell |
| No top-level `compatibility/` | confirmed absent |
| No top-level `deferred-integration/` | confirmed absent |
| `brewprint/compatibility/` exists | confirmed present |
| BPDSL non-canonical status declared | `bpdsl/index.md` Purpose and Ownership status rows confirm non-canonical staging |
| Migration trigger visible | `bpdsl/index.md` migration trigger row: "Review every staged item during BPDSL migration or an explicit integration requirement." |
| Exit conditions visible | `bpdsl/index.md` `## Exit conditions` section lists three disposition paths |

### 6. Work Item closure check

All Work Item completion conditions were verified. Full matrix is recorded in the Work Item Evidence `## Completion-condition matrix`. All conditions satisfied.

### 7. Scope and Git evidence

**T13-edited files**: exactly 14 files listed in §1.

| check | result |
|---|---|
| Nothing staged | `git diff --cached --name-status` produced no output |
| `v01/` unchanged | `git diff --name-status -- v01` produced no output |
| T12 unchanged | no T13 edit touched `PRODUCT-TASK-SPEC-012-12-*.md` |
| No DRMCP or BPDSL app-local spec changed by T13 | confirmed — DRMCP and BPDSL files appearing in `git status` were already modified before T13 started |
| No file relocated | confirmed — all edits are in-place content edits |
| No optional advisory cleanup applied | confirmed — A1–A4 not applied |
| No commit created | confirmed
