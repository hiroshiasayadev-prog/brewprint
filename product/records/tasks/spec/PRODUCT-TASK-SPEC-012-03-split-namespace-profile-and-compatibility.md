# PRODUCT-TASK-SPEC-012-03: Split namespace, profile, and compatibility semantics

- **id**: PRODUCT-TASK-SPEC-012-03
- **status**: done
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-SPEC-012
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 2d
- **depends_on**:
  - PRODUCT-TASK-SPEC-012-02
- **outputs**:
  - Cleaned Design Records namespace-model content
  - `product/records/spec/brewprint/namespaces/`
  - `product/records/spec/brewprint/compatibility/`

## Goal

Separate generic namespace semantics from Brewprint registry facts and V01 compatibility material.

## Work

Process this semantic batch:

- `namespace-model/index.md`
- `namespace-model/app-namespaces.md`
- `namespace-model/domain-catalog.md`
- `namespace-model/existing-artifacts.md`
- `namespace-model/legacy-id-compatibility.md`
- `namespace-model/subdomain-model.md`

Keep app-independent namespace and subdomain contracts with Design Records.
Move current app and domain assignments into Brewprint profile specifications.
Place V01 attribution, issued-ID retention, and migration state under Brewprint compatibility.
Replace DRMCP operational descriptions with pointers or handoff notes.
Do not relocate cleaned Design Records files in this task.

## Done condition

- Generic namespace semantics contain no current Brewprint registry table.
- Brewprint namespace profile contains current project assignments.
- Compatibility content resides under `brewprint/compatibility/`.
- DRMCP operational behavior is removed from PRODUCT normative text.
- The six-file semantic batch is independently reviewable.

## Verification

- Compare each edited section against `PRODUCT-INV-SPEC-005`.
- Confirm artifact ID grammar ownership remains app-independent.
- Confirm V01 compatibility does not become a top-level PRODUCT area.
- Confirm no broad ref synchronization occurs.

## Evidence

### Execution summary

T03 processed the exact six namespace-model source files.
Generic Design Records namespace semantics remained under `product/records/spec/concepts/namespace-model/`.
Brewprint registry facts moved to `product/records/spec/brewprint/namespaces/`.
Brewprint V01 compatibility moved to `product/records/spec/brewprint/compatibility/`.
No `drmcp/records/spec/**`, `bpdsl/records/spec/**`, or `v01/**` files were changed.
T04 and later task files were not modified.

### Source files processed

| source file | result |
|---|---|
| `product/records/spec/concepts/namespace-model/index.md` | Rewritten in place to keep generic namespace semantics and preserve T05 pending placement material. |
| `product/records/spec/concepts/namespace-model/app-namespaces.md` | Rewritten in place to keep generic app namespace semantics and remove app-local architecture content. |
| `product/records/spec/concepts/namespace-model/domain-catalog.md` | Removed from generic namespace-model after Brewprint profile content was recorded under `brewprint/namespaces/`. |
| `product/records/spec/concepts/namespace-model/existing-artifacts.md` | Rewritten in place to keep only generic new-artifact ownership selection. |
| `product/records/spec/concepts/namespace-model/legacy-id-compatibility.md` | Removed from generic namespace-model after compatibility content was recorded under `brewprint/compatibility/`. |
| `product/records/spec/concepts/namespace-model/subdomain-model.md` | Rewritten in place to keep generic subdomain semantics and replace the DRMCP example with a T08 handoff note. |

### Files created, rewritten, moved, and removed

| action | files |
|---|---|
| Created | `product/records/spec/brewprint/namespaces/index.md`; `product/records/spec/brewprint/namespaces/app-namespaces.md`; `product/records/spec/brewprint/namespaces/domain-catalog.md`; `product/records/spec/brewprint/compatibility/index.md`; `product/records/spec/brewprint/compatibility/existing-artifacts.md`; `product/records/spec/brewprint/compatibility/legacy-id-compatibility.md`. |
| Rewritten | `product/records/spec/concepts/namespace-model/index.md`; `product/records/spec/concepts/namespace-model/app-namespaces.md`; `product/records/spec/concepts/namespace-model/existing-artifacts.md`; `product/records/spec/concepts/namespace-model/subdomain-model.md`; `product/records/spec/brewprint/index.md`. |
| Removed | `product/records/spec/concepts/namespace-model/domain-catalog.md`; `product/records/spec/concepts/namespace-model/legacy-id-compatibility.md`. |
| Moved | None as a filesystem move. Content was split by rewrite, create, and remove so the diff remains reviewable. |

### Section disposition map

| source file | section | disposition | target or note |
|---|---|---|---|
| `namespace-model/index.md` | `## What this is` | retain generic | Reframed as app-independent Design Records namespace model. |
| `namespace-model/index.md` | `## Current contract` | split | Generic concept table retained; current Brewprint registry moved to `spec:product.brewprint.namespaces`. |
| `namespace-model/index.md` | `## Boundary` | split | Generic ownership retained; Brewprint profile, compatibility, DRMCP, and BPDSL ownership routed by pointer. |
| `namespace-model/index.md` | `## Current placement and future layout` | T05 pending | Preserved with a pending-disposition note. |
| `namespace-model/index.md` | `## App namespace and domain namespace` | retain generic | Brewprint examples replaced with app-neutral examples. |
| `namespace-model/index.md` | `## Topics` | split | Generic children retained; profile and compatibility children removed from generic topics. |
| `namespace-model/index.md` | `## Sources` | retain evidence | Historical source evidence preserved without defining current placement. |
| `app-namespaces.md` | `## What this is` | retain generic | Rewritten as generic app namespace semantics. |
| `app-namespaces.md` | `## App namespace definitions` | move to Brewprint profile | Recorded in `spec:product.brewprint.namespaces.app_namespaces`. |
| `app-namespaces.md` | `## DRMCP` | T08 handoff plus profile summary | Current assignment recorded in Brewprint profile; architecture/tool behavior replaced by T08 handoff. |
| `app-namespaces.md` | `## BPDSL` | T09 handoff plus profile summary | Current assignment recorded in Brewprint profile; type/resolver/render/self-hosting content replaced by T09 handoff. |
| `app-namespaces.md` | `## PRODUCT` | move to Brewprint profile | Recorded in app namespace and domain profile tables. |
| `domain-catalog.md` | `## What this is` | move to Brewprint profile | Recorded in `spec:product.brewprint.namespaces.domain_catalog`. |
| `domain-catalog.md` | `## Canonical domain namespaces` | move to Brewprint profile | Corrected as current, legacy-effective, or future candidate profile rows. |
| `domain-catalog.md` | `## Existing prefixes outside the canonical catalog` | move to compatibility | Recorded as cross-app and legacy prefix compatibility material. |
| `existing-artifacts.md` | `## What this is` | split | Generic new-artifact rule retained; compatibility framing moved. |
| `existing-artifacts.md` | `## Historical ownership decision` | move to compatibility | Recorded in `spec:product.brewprint.compatibility.existing_artifacts`. |
| `existing-artifacts.md` | `## Effective attribution` | move to compatibility plus T08 handoff | Attribution table moved; UI/MCP projection behavior replaced by T08 handoff. |
| `existing-artifacts.md` | `## New-artifact ownership` | retain generic | Kept as app-independent namespace selection rule. |
| `legacy-id-compatibility.md` | all sections | move to compatibility | Recorded in `spec:product.brewprint.compatibility.legacy_id_compatibility`. |
| `subdomain-model.md` | `## What this is` | retain generic | Rewritten without DRMCP-specific framing. |
| `subdomain-model.md` | `## Subdomain model` | retain generic | Kept generic grouping rule. |
| `subdomain-model.md` | `## Definition and representation` | retain generic | Kept metadata representation. |
| `subdomain-model.md` | `## Write-time advisory` | retain generic | Kept app-independent authoring advisory behavior. |
| `subdomain-model.md` | `## DRMCP MCP domain example` | T08 handoff | Removed concrete example and recorded T08 handoff note. |

### Current namespace assignments recorded

| assignment | status | evidence |
|---|---|---|
| `DRMCP` app namespace | Active app namespace. | Existing app-local DRMCP records. |
| `BPDSL` app namespace | Active app namespace. | Brewprint profile assignment retained without BPDSL internals. |
| `PRODUCT` namespace | Active product namespace. | Existing PRODUCT record IDs. |
| `DRUI` app namespace | Future candidate. | Explicitly not active. |
| `DRMCP` / `MCP` | Active assignment. | `DRMCP-REQ-MCP-001`, `DRMCP-REQ-MCP-002`, `DRMCP-INV-MCP-001`. |
| `DRMCP` / `SPEC` | Active assignment. | `DRMCP-WORK-SPEC-001`, `DRMCP-WORK-SPEC-002`. |
| `PRODUCT` / `NAMESPACE` | Active assignment. | `PRODUCT-WORK-NAMESPACE-001`, `PRODUCT-TASK-NAMESPACE-001-01`. |
| `PRODUCT` / `SPEC` | Active assignment. | `PRODUCT-REQ-SPEC-001`, `PRODUCT-WORK-SPEC-012`, `PRODUCT-TASK-SPEC-012-03`. |
| `BPDSL` / `DATA` | Legacy-effective assignment. | V01 compatibility attribution only. |
| `BPDSL` / `RESOLVE` | Legacy-effective assignment. | V01 compatibility attribution only. |
| `PRODUCT` / `GOVERNANCE` | Future candidate. | No scoped current app-aware record ID found. |
| `PRODUCT` / `MIGRATION` | Future candidate. | No scoped current app-aware record ID found. |
| `SELFHOST` | Cross-app legacy activity. | Recorded under compatibility. |

### Compatibility material recorded

| material | target |
|---|---|
| V01 ownership decision under V01-ADR-096 | `spec:product.brewprint.compatibility.existing_artifacts`. |
| Effective historical attribution map | `spec:product.brewprint.compatibility.existing_artifacts`. |
| Issued-ID retention and no-rename policy | `spec:product.brewprint.compatibility.legacy_id_compatibility`. |
| V01 accepted legacy ID families | `spec:product.brewprint.compatibility.legacy_id_compatibility`. |
| Compatibility-only `V01-SPEC-*` identity note | `spec:product.brewprint.compatibility.legacy_id_compatibility`. |

### Handoff notes

| source file | source section | downstream task | replacement |
|---|---|---|---|
| `app-namespaces.md` | `## DRMCP` | PRODUCT-TASK-SPEC-012-08 | Pointer and handoff note for DRMCP architecture and tool behavior. |
| `existing-artifacts.md` | `## Effective attribution` | PRODUCT-TASK-SPEC-012-08 | Pointer and handoff note for UI/MCP projection behavior. |
| `subdomain-model.md` | `## DRMCP MCP domain example` | PRODUCT-TASK-SPEC-012-08 | Pointer and handoff note for concrete DRMCP MCP-domain example. |
| `app-namespaces.md` | `## BPDSL` | PRODUCT-TASK-SPEC-012-09 | Pointer and handoff note for BPDSL type, resolver, render, YAML parser, and self-hosting behavior. |

### T05 material intentionally preserved

`namespace-model/index.md` preserves `## Current placement and future layout` with a pending-disposition note.
No final placement decision was made for the future machine-readable namespace registry.

### Objective contradiction found

The old generic domain catalog listed `PRODUCT` / `GOVERNANCE` and `PRODUCT` / `MIGRATION` as canonical domains.
Scoped current app-aware IDs show active `PRODUCT` / `SPEC`, `PRODUCT` / `NAMESPACE`, and `DRMCP` / `SPEC` usage.
The Brewprint profile records the corrected current facts.
The generic namespace model does not absorb those registry facts.

### Validation

| command | exit code | output |
|---|---:|---|
| `python -X utf8 product/src/tools/validate_spec.py product/records/spec --strict --no-color` | 0 | `[strict]  All 47 file(s) OK.` |

### Scoped checks

| command | exit code | output |
|---|---:|---|
| `git status --short -- product/records/spec` | 0 | ` M product/records/spec/brewprint/index.md`<br>` M product/records/spec/concepts/namespace-model/app-namespaces.md`<br>` D product/records/spec/concepts/namespace-model/domain-catalog.md`<br>` M product/records/spec/concepts/namespace-model/existing-artifacts.md`<br>` M product/records/spec/concepts/namespace-model/index.md`<br>` D product/records/spec/concepts/namespace-model/legacy-id-compatibility.md`<br>` M product/records/spec/concepts/namespace-model/subdomain-model.md`<br>`?? product/records/spec/bpdsl/`<br>`?? product/records/spec/brewprint/compatibility/`<br>`?? product/records/spec/brewprint/namespaces/`<br>`?? product/records/spec/design-records/`<br>`?? product/records/spec/index.md` |
| `git status --short -- drmcp/records/spec bpdsl/records/spec v01` | 0 | No output. |
| `git diff --name-status -- product/records/spec` | 0 | `M	product/records/spec/brewprint/index.md`<br>`M	product/records/spec/concepts/namespace-model/app-namespaces.md`<br>`D	product/records/spec/concepts/namespace-model/domain-catalog.md`<br>`M	product/records/spec/concepts/namespace-model/existing-artifacts.md`<br>`M	product/records/spec/concepts/namespace-model/index.md`<br>`D	product/records/spec/concepts/namespace-model/legacy-id-compatibility.md`<br>`M	product/records/spec/concepts/namespace-model/subdomain-model.md` |

### Scope confirmations

| check | result |
|---|---|
| Generic namespace-model files relocated to `design-records/namespace-model/` | No. |
| App-local specs under `drmcp/records/spec/**` changed | No. |
| App-local specs under `bpdsl/records/spec/**` changed | No. |
| `v01/**` changed | No. |
| Broad `spec:product.concepts` synchronization | Not performed. The word-diff check showed only local deletions and local pointer edits in the T03 batch. |
| New files inspected despite untracked status | Yes. New Brewprint namespace and compatibility files were inspected explicitly. |
