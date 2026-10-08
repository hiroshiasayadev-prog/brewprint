# PRODUCT-INV-SPEC-011 semantic-analysis evidence index

This index records product-side status for semantic-analysis evidence derived from PRODUCT-INV-SPEC-011.

The ignored tools output remains retrospective source material.
Repository-persistent evidence is valid only after PRODUCT-WORK-SPEC-028 captures or synchronizes it under `product/records/`.

## Ownership

| item | value |
|---|---|
| Source Investigation | PRODUCT-INV-SPEC-011 |
| Analysis Requirement | PRODUCT-REQ-SPEC-014 |
| Owning Work Item | PRODUCT-WORK-SPEC-028 |
| Route decision Task | PRODUCT-TASK-SPEC-028-01 |
| Raw observation corpus | `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/results/*.observations.jsonl` |
| Working output root | `tools/term-inventory-analysis/output/` |

## Accepted retrospective source stages

PRODUCT-TASK-SPEC-028-01 accepted these completed stages as retrospective source material.
They still require product-side capture or synchronization before repository history can rely on them as durable evidence.

| stage | working output | accepted completion evidence | product-side owner |
|---|---|---|---|
| Leaf semantic analysis | `PRODUCT-INV-SPEC-011-leaf-analysis/` | 71 of 71 valid results. | PRODUCT-TASK-SPEC-028-02 |
| Trigger-level reduction | `PRODUCT-INV-SPEC-011-trigger-reduction/` | 57 of 57 valid results. | PRODUCT-TASK-SPEC-028-03 |
| Tier A cross-trigger identity review | `PRODUCT-INV-SPEC-011-cross-trigger-review/` | 54 of 54 valid jobs and spot audit PASS. | PRODUCT-TASK-SPEC-028-04 |

## Product-side evidence status

| product-side path | purpose | status | owner |
|---|---|---|---|
| `leaf-analysis/` | Product-side leaf-analysis summary and machine-readable evidence. | directory exists; evidence pending | PRODUCT-TASK-SPEC-028-02 |
| `trigger-reduction/` | Product-side trigger-reduction summary and machine-readable evidence. | pending creation | PRODUCT-TASK-SPEC-028-03 |
| `cross-trigger-review/README.md` | Commit-safe summary of Tier A review scope, routing completion, validation, and storage policy. | captured; pending W028 evidence-policy synchronization | PRODUCT-TASK-SPEC-028-04 |
| `cross-trigger-review/tier-a-review-summary.json` | Machine-readable Tier A review completion, spot-audit, and ignored-output policy summary. | captured; pending W028 evidence-policy synchronization | PRODUCT-TASK-SPEC-028-04 |
| `semantic-analysis/index.md` | This index. Maps accepted source stages to product-side capture owners. | captured | PRODUCT-WORK-SPEC-028 |

## Working output inventory

| working output | product-side evidence policy |
|---|---|
| `tools/term-inventory-analysis/output/PRODUCT-INV-SPEC-011-leaf-analysis/` | Retrospective source material for PRODUCT-TASK-SPEC-028-02. |
| `tools/term-inventory-analysis/output/PRODUCT-INV-SPEC-011-trigger-reduction/` | Retrospective source material for PRODUCT-TASK-SPEC-028-03. |
| `tools/term-inventory-analysis/output/PRODUCT-INV-SPEC-011-cross-trigger-review/` | Retrospective source material already summarized under `cross-trigger-review/`; T04 synchronizes the summary with W028. |
| `tools/term-inventory-analysis/output/PRODUCT-INV-SPEC-011-cross-trigger-analysis/` | Working output only until a later Task explicitly admits or captures it. |
| `tools/term-inventory-analysis/output/PRODUCT-INV-SPEC-011-semantic-packets/` | Working output only until a later Task explicitly admits or captures it. |
| `tools/term-inventory-analysis/output/PRODUCT-INV-SPEC-011-profile/` | Working output only until a later Task explicitly admits or captures it. |

## Current route

| task | role | dependency | output boundary |
|---|---|---|---|
| PRODUCT-TASK-SPEC-028-02 | Capture leaf-analysis evidence. | PRODUCT-TASK-SPEC-028-01 | `leaf-analysis/` |
| PRODUCT-TASK-SPEC-028-03 | Capture trigger-reduction evidence. | PRODUCT-TASK-SPEC-028-01 | `trigger-reduction/` |
| PRODUCT-TASK-SPEC-028-04 | Synchronize cross-trigger-review evidence with W028 policy. | PRODUCT-TASK-SPEC-028-01 | `cross-trigger-review/` |
| PRODUCT-TASK-SPEC-028-05 | Coordinate follow-up routes. | T02, T03, T04 | PRODUCT-WORK-SPEC-028 |
| PRODUCT-TASK-SPEC-028-06 | Review captured evidence and routes. | T02, T03, T04, T05 | PRODUCT-WORK-SPEC-028 |
| PRODUCT-TASK-SPEC-028-07 | Synchronize closure after review route. | T06 | PRODUCT-WORK-SPEC-028 |

## Boundaries

This index does not approve canonical vocabulary.
This index does not define, deprecate, retire, or replace any term.
This index does not rewrite any source artifact.
This index does not make ignored tools output durable product evidence by itself.
