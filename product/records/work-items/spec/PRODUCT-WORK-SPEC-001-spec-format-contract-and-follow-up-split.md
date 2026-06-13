# PRODUCT-WORK-SPEC-001: Spec format contract and follow-up split

- **id**: PRODUCT-WORK-SPEC-001
- **status**: done
- **date**: 2026-06-10
- **requirement_refs**:
  - PRODUCT-REQ-SPEC-001
- **investigation_refs**:
  - PRODUCT-INV-SPEC-001
- **task_refs**:
  - PRODUCT-TASK-SPEC-001-01
  - PRODUCT-TASK-SPEC-001-02
  - PRODUCT-TASK-SPEC-001-03
  - PRODUCT-TASK-SPEC-001-04

## Summary

Define the accepted PRODUCT-level spec document format contract needed for MCP-readable topic navigation, and produce the exact follow-up split for spec ID-as-ref compatibility, temporary PRODUCT validation tooling, authoring guide updates, ownership-sensitive traceability boundary decisions, existing spec migration, and later DRMCP implementation-phase validation.

This work starts from PRODUCT-INV-SPEC-001. It must not jump directly into bulk migration or DRMCP implementation.

## Scope

This work owns the first executable step after PRODUCT-REQ-SPEC-001 acceptance.

| area | in scope |
|---|---|
| spec kind set | Decide accepted initial kinds and deferred kinds with revisit conditions. |
| H1 format | Define `# <SpecKind>: <Title>` and parser-aware one-H1 rule outside YAML front matter and fenced code blocks. |
| required sections | Define required / recommended / prohibited sections by spec kind. |
| `## What this is` | Define canonical requirement and migration severity for existing specs. |
| `Index` / `## Topics` | Define table columns, ownership of parent relation, duplicate-parent rule, and Index purity vs Overview+Topics. |
| parent grammar | Define interim parent grammar before derived topic refs are finalized. |
| spec ID-as-ref derivation | Define creation-time path-derived `spec:` IDs, including the `index.md` omission rule, underscore segment policy, and validation warnings for mismatches. |
| migration phase | Define staging and severity policy without bulk-migrating specs. |
| follow-up split | Define exact follow-up WORK / INV / implementation work and ordering. |

## Initial Design Position

| decision area | initial recommendation | rationale |
|---|---|---|
| accepted spec kinds | `Overview`, `Index`, `Concept`, `Reference`, `Contract` | PRODUCT-INV-SPEC-001 found these five cover current PRODUCT / DRMCP / BPDSL samples. |
| deferred spec kinds | `Guide`, `Process`, `Architecture`, `Glossary` | Existing specs do not force standalone kinds yet. Keep sections under existing kinds until a real need appears. |
| H1 format | `# <SpecKind>: <Title>` | Parser-friendly and mechanically validateable. |
| one-H1 rule | Exactly one ATX H1 outside YAML front matter and fenced code blocks | Existing parser primitives already distinguish real headings from fenced examples. |
| `## What this is` | Required for new / migrated specs; warning during migration for existing specs | 0 sampled existing specs already have it, so hard-failing the tree immediately is too disruptive. |
| Index purity | Allow `Overview` to contain `## Topics`; keep pure `Index` navigation-first | Avoids forced split of hybrid files such as traceability index while preserving navigation semantics. |
| stable `spec:` IDs | Remain canonical | Current stable `spec:` identity contract requires rename / move / split stability. |
| path-derived `spec:` IDs | Creation-time suggestion only; not live binding | Validator must warn when a spec file has a missing or suspicious path-derived ID mismatch unless alias / redirect compatibility is defined. |
| derived topic refs | Not introduced in this work as canonical IDs | Compatibility and resolver role are deferred to PRODUCT-WORK-SPEC-002. |
| ownership boundary | PRODUCT-owned semantics remain default; migration needs separate ownership review | Format contract can proceed, but migration must not blur PRODUCT semantics and DRMCP tool contracts. |

## Spec Kind Matrix

| spec kind | status | intent | revisit condition |
|---|---|---|---|
| `Overview` | accepted initial | Entry point for a spec area; may summarize current contract and contain `## Topics` when needed. | N/A |
| `Index` | accepted initial | Navigation-first topic table with minimal substantive specification body. | N/A |
| `Concept` | accepted initial | Concept model, vocabulary, semantic boundary, and rules. | N/A |
| `Reference` | accepted initial | Field, enum, grammar, fixed rule, or catalog reference. | N/A |
| `Contract` | accepted initial | API/tool/request/response/validation/error behavior contract. | N/A |
| `Guide` | deferred | Usage, authoring, or operation procedure. | Revisit when current guides become spec records. |
| `Process` | deferred | Lifecycle / workflow / transition model. | Revisit when workflow lifecycle needs standalone spec representation. |
| `Architecture` | deferred | Standalone component / runtime / storage / dependency architecture. | Revisit when a standalone architecture spec appears. |
| `Glossary` | deferred | Term-only file. | Revisit when term-only specs appear. |

## Required Section Matrix

| section / marker | Overview | Index | Concept | Reference | Contract |
|---|---|---|---|---|---|
| H1-adjacent `- **id**:` | required | required | required | required | required |
| H1-adjacent `- **status**:` | required | required | required | required | required |
| H1-adjacent `- **date**:` | required | required | required | required | required |
| H1-adjacent `- **parent**:` | required | required | required | required | required |
| `## What this is` | required | required | required | required | required |
| `## Source records` | recommended | recommended | recommended | recommended | recommended |
| `## Current contract` | required | prohibited | optional | optional | optional |
| `## Non-goals` | recommended | prohibited | recommended | optional | recommended |
| `## Topic map` | recommended | prohibited | optional | optional | optional |
| `## Topics` table | optional | required | prohibited | prohibited | prohibited |
| `## Concept model` | optional | prohibited | required | prohibited | prohibited |
| `## Rules` | optional | prohibited | recommended | optional | optional |
| `## Boundary` | optional | prohibited | recommended | optional | recommended |
| reference table with kind-specific title | prohibited | prohibited | optional | required | optional |
| `## Request` | prohibited | prohibited | prohibited | prohibited | required |
| `## Response` | prohibited | prohibited | prohibited | prohibited | required |
| `## Errors` | optional | prohibited | prohibited | optional | required |
| `## Validation rules` | optional | prohibited | optional | optional | required |
| `## Related specs` | recommended | recommended | recommended | recommended | recommended |

## Index and Topics Contract Decisions

| issue | decision to make in this work | initial recommendation |
|---|---|---|
| Pure Index vs Overview+Topics | Whether `Overview` may contain authoritative `## Topics`. | Allow it. Use `Overview` when the file has substantive body; use `Index` when navigation-first. |
| Required columns | Minimum `## Topics` table columns. | `title`, `kind`, `parent`, `file`, `summary`. |
| Duplicate parent | Whether a child may appear under multiple parents. | Prohibit. A child file must have exactly one authoritative parent declaration. |
| Summary validation | Whether `summary` is semantic-validated. | Presence / non-empty only. Truthfulness is human-reviewed. |
| File path in table | Whether `file` is allowed. | Allowed as child target path, but not as canonical parent reference. |
| Filename implication | Whether `index.md` forces `Index` kind. | No. Kind is determined by H1 and content role. |

## Spec ID-as-ref Derivation

| rule | contract |
|---|---|
| derivation timing | Default `spec:` IDs are suggested at spec creation time. They are not live-bound to the file path after creation. |
| base prefix | `<app>/records/spec/` maps to `spec:<app>.` using lower-case app namespace. |
| path separator | Directory separators under `records/spec/` map to `.`. |
| extension | `.md` is removed. |
| `index.md` | `index` is omitted because `index.md` represents the entrypoint topic for its containing directory. |
| non-index file | File stem is kept as the final spec ID segment. |
| word separator | Hyphenated path/title segments are normalized to underscore in spec ID segments. |
| write location | The derived ID is written to H1-adjacent `- **id**:`; no separate `ref` marker is used. |
| rename / move | Existing stable `spec:` IDs are not auto-changed. Alias / redirect / compatibility behavior belongs to PRODUCT-WORK-SPEC-002. |
| validation | Validator must emit a warning when the visible `id` does not match the path-derived default ID or otherwise looks inconsistent with the current path. |

Examples:

| path | default spec ID |
|---|---|
| `product/records/spec/concepts/spec-format/index.md` | `spec:product.concepts.spec_format` |
| `product/records/spec/concepts/traceability/index.md` | `spec:product.concepts.traceability` |
| `product/records/spec/concepts/traceability/semantic-ref.md` | `spec:product.concepts.traceability.semantic_ref` |
| `drmcp/records/spec/design-records-mcp/tools.md` | `spec:drmcp.design_records_mcp.tools` |

The warning is required because silent drift between path-derived expectation and declared stable IDs makes review context and topic navigation unreliable. The warning must not rewrite IDs automatically. The underscore segment policy conflicts with older hyphen-only `spec:` segment text and must be handled by PRODUCT-WORK-SPEC-002 or an explicit traceability spec update before enforcement.

## Parent Grammar Decision Matrix

| candidate parent form | recommended status | reason |
|---|---|---|
| `root` or `-` root marker | allow | Needed for top-level topic root without inventing a parent. |
| stable `spec:` ID-as-ref | allow | Stable across rename / move / split and already part of traceability model. |
| physical path | prohibit | Stale on move and conflicts with canonical-reference principles. |
| file name | prohibit | Ambiguous and stale on rename. |
| H1 title string | prohibit | Stale on title change and not stable identity. |
| Markdown heading anchor | prohibit | Parser/render dependent and stale on heading change. |
| title-derived topic ref | defer / prohibit in initial contract | Compatibility and collision behavior belong to PRODUCT-WORK-SPEC-002. |

Interim grammar to evaluate in this work:

```text
parent := "-" | "root" | spec_ref
spec_ref := "spec:" segment ("." segment)*
segment := [a-z0-9][a-z0-9_]*
```

## Stable Spec ID-as-ref and Derived Topic Ref Boundary

| item | this work | follow-up |
|---|---|---|
| stable `spec:` ID-as-ref | Keep canonical and unchanged. | PRODUCT-WORK-SPEC-002 may extend compatibility rules, but must not silently break stability. |
| path-derived default ID mismatch | Warn only; do not auto-rewrite. | PRODUCT-WORK-SPEC-006 should provide temporary validation before migration; DRMCP-WORK-SPEC-001 can implement durable diagnostics in the later DRMCP reimplementation phase. |
| derived topic refs | Do not introduce as canonical IDs in this work. | PRODUCT-WORK-SPEC-002 decides display-only vs resolver input vs deferred. |
| alias / redirect | Out of scope except to identify need. | PRODUCT-WORK-SPEC-002 owns alias, redirect, superseded mapping, collision, split, merge, move behavior. |
| legacy `semantic_refs` / `sections` | Preserve existing meaning. | PRODUCT-WORK-SPEC-002 defines coexistence and migration compatibility. |

## Ownership Boundary Matrix

| concept / contract area | initial owner | rationale | follow-up |
|---|---|---|---|
| project artifact model | PRODUCT | Cross-app semantics; BPDSL may reuse it. | Review during PRODUCT-INV-SPEC-002, but default stays PRODUCT. |
| namespace model / repository layout | PRODUCT | Cross-app governance. | No immediate relocation expected. |
| stable `spec:` grammar / stability rule | PRODUCT | Identity semantics must not be owned by one tool implementation. | PRODUCT-WORK-SPEC-002. |
| trace metadata schema | PRODUCT | Shared metadata contract. | PRODUCT-INV-SPEC-002 confirms section-level boundary. |
| canonical resolution semantics | PRODUCT | Semantic contract, not a tool-specific request/response shape. | PRODUCT-INV-SPEC-002 / PRODUCT-WORK-SPEC-004. |
| resolver tool request / response | DRMCP | Tool boundary. | DRMCP work items after PRODUCT contract. |
| temporary spec-format validator / resolver | PRODUCT | Migration needs lightweight validation before DRMCP is redesigned. | PRODUCT-WORK-SPEC-006. |
| validation diagnostics exposed by DRMCP | DRMCP | Tool implementation / API contract. | DRMCP-WORK-SPEC-001 / 002 as implementation-phase follow-up, not as PRODUCT stabilization prerequisite. |
| parser-aware H1 / Topics validation implementation in DRMCP | DRMCP | Durable implementation of PRODUCT-owned format contract after DRMCP redesign. | DRMCP-WORK-SPEC-001 / 002. |

## Prerequisite Investigation Decision

| question | decision |
|---|---|
| Is a separate ownership investigation needed? | Yes: create `PRODUCT-INV-SPEC-002`. |
| Is it a gate for PRODUCT-WORK-SPEC-001? | No. Format contract applies to PRODUCT / DRMCP / BPDSL regardless of final ownership. |
| Is it a gate for migration? | Yes. Migration and relocation must wait for ownership boundary classification. |
| Target scope | `product/records/spec/concepts/traceability/**` and `product/records/spec/concepts/project-artifact-model/index.md`, with section-level classification into PRODUCT-owned semantics, DRMCP-owned tool contract, or hybrid. |

## Required Follow-up Split Output

This work must produce an exact follow-up plan and create or explicitly recommend the next artifacts. The follow-up split must be represented as a table like this:

| artifact | kind | depends on | purpose | gate |
|---|---|---|---|---|
| PRODUCT-WORK-SPEC-002 | WORK | PRODUCT-WORK-SPEC-001 | Stable `spec:` ID-as-ref and derived topic ref compatibility design. | User decision on derived topic ref role. |
| PRODUCT-WORK-SPEC-003 | WORK | PRODUCT-WORK-SPEC-001 | Spec authoring guide format update. | Format contract accepted. |
| PRODUCT-INV-SPEC-002 | INV | PRODUCT-WORK-SPEC-001 draft or accepted format | Artifact / traceability ownership boundary investigation. | Needed before migration / relocation. |
| PRODUCT-WORK-SPEC-004 | WORK | PRODUCT-INV-SPEC-002, PRODUCT-WORK-SPEC-001 | Ownership boundary decision and relocation plan. | User / ADR-level decision. |
| PRODUCT-WORK-SPEC-006 | WORK | PRODUCT-WORK-SPEC-001, PRODUCT-WORK-SPEC-002 | Temporary PRODUCT-level spec format validator / resolver tooling. | Needed before PRODUCT-WORK-SPEC-005 migration validation; not a DRMCP reimplementation. |
| PRODUCT-WORK-SPEC-005 | WORK | PRODUCT-WORK-SPEC-001, 002, 003, 004, 006 | Existing spec format migration and restructuring using temporary validation tooling. | Spec update review gates per migration task. |
| DRMCP-REQ-MCP-001 / DRMCP-INV-MCP-001 | REQ / INV | PRODUCT-WORK-SPEC-005 | Existing app namespace / multi-root MCP contract redesign input. | Revisit after spec migration/restructuring; create concrete WORK only when entering app namespace redesign. |
| DRMCP-WORK-SPEC-001 | WORK | PRODUCT spec-format stabilization and DRMCP redesign/reimplementation phase | Parser-aware spec format validation in the future DRMCP implementation. | Implementation-phase follow-up; not a prerequisite for PRODUCT spec-format stabilization. |
| DRMCP-WORK-SPEC-002 | WORK | PRODUCT-WORK-SPEC-002 and DRMCP-WORK-SPEC-001, during DRMCP redesign/reimplementation phase | Index Topics graph validation in the future DRMCP implementation. | Implementation-phase follow-up; not a prerequisite for PRODUCT spec-format stabilization. |

## Deliverables

| deliverable | expected path / form | notes |
|---|---|---|
| PRODUCT spec format contract | `product/records/spec/concepts/spec-format/index.md` as initial candidate | This work must confirm or adjust the path. |
| accepted kind set | In the format contract | Initial five kinds plus deferred four kinds and revisit conditions. |
| required section matrix | In the format contract | Use matrix/table form; do not bury as prose. |
| Topics table contract | In the format contract | Include required columns and duplicate-parent prohibition. |
| spec ID-as-ref derivation | In the format contract | Include `index.md` omission rule, underscore segment policy, no separate `ref`, and mismatch warning requirement. |
| parent grammar | In the format contract | Interim grammar must prohibit path / filename / H1 title / derived topic ref unless explicitly justified. |
| follow-up split | In this work item evidence and/or a dedicated section | Must include ordering and dependency gates. |
| follow-up records | `*-new` placeholder IDs when using MCP authoring | Create if tooling supports it; otherwise document why not. |

## Explicitly Excluded Scope

| excluded area | reason |
|---|---|
| Existing spec bulk migration | Owned by PRODUCT-WORK-SPEC-005. |
| Temporary validator/tooling implementation | Owned by PRODUCT-WORK-SPEC-006; this work only defines the need and boundary. |
| Current DRMCP implementation changes | Excluded. DRMCP-WORK-SPEC-001 / 002 are later implementation-phase follow-ups, not immediate patch targets. |
| Authoring guide updates | Owned by PRODUCT-WORK-SPEC-003. |
| Final stable `spec:` ID-as-ref / derived topic ref compatibility contract | Owned by PRODUCT-WORK-SPEC-002. |
| App namespace MCP contract redesign | Existing DRMCP-REQ-MCP-001 / DRMCP-INV-MCP-001 capture the issue; concrete WORK should be created later, after PRODUCT spec restructuring. |
| UI topic tree display | Not needed for document format contract. |
| v01 snapshot modification | `v01/` is historical and must remain unchanged. |

## Evidence

- Close synchronization completed on 2026-06-10.
- Product spec format contract created at `product/records/spec/concepts/spec-format/index.md`.
- Review gates completed:
  - Codex review: OK with minor fixes; TASK-SPEC-001-02 marked done.
  - Opus design review: PASS / TASK-SPEC-001-03 READY.
- Follow-up split created and corrected for the current app namespace / DRMCP rebuild roadmap:
  - PRODUCT-WORK-SPEC-002: stable `spec:` ID-as-ref and derived topic compatibility.
  - PRODUCT-INV-SPEC-002: artifact / traceability ownership boundary investigation.
  - PRODUCT-WORK-SPEC-004: ownership boundary decision and relocation plan.
  - PRODUCT-WORK-SPEC-003: spec authoring guide update.
  - PRODUCT-WORK-SPEC-006: temporary PRODUCT-level spec format validator / resolver tooling.
  - PRODUCT-WORK-SPEC-005: existing spec migration and restructuring using temporary tooling.
  - Existing `DRMCP-REQ-MCP-001` / `DRMCP-INV-MCP-001`: app namespace redesign input after PRODUCT spec restructuring.
  - DRMCP-WORK-SPEC-001 / DRMCP-WORK-SPEC-002: future DRMCP implementation-phase validation work.
- PRODUCT-WORK-SPEC-006 was added as the required temporary validator/tooling bridge before existing spec migration.
- DRMCP-WORK-SPEC-001 / DRMCP-WORK-SPEC-002 are explicitly isolated as implementation-phase follow-ups and are not prerequisites for PRODUCT spec-format stabilization.
- PRODUCT-WORK-SPEC-005 was updated to depend on temporary validation tooling before migration.
- Scope guard confirmed:
  - No existing spec bulk migration was performed.
  - No DRMCP implementation code was changed.
  - No app namespace redesign was performed in this work.
  - No `v01/records/**` changes were made by close synchronization.
- Current MCP validation failures are non-blocking because the app namespace transition leaves PRODUCT namespace / metadata scaffolds unsupported.

| evidence | implication |
|---|---|
| PRODUCT-REQ-SPEC-001 accepted the need for MCP-readable spec format and topic tree support. | This work is justified. |
| PRODUCT-INV-SPEC-001 concluded that the format is implementable with a staged approach. | Contract drafting can proceed. |
| PRODUCT-INV-SPEC-001 recommended starting with `Overview`, `Index`, `Concept`, `Reference`, and `Contract`. | Initial kind set should not include all candidates. |
| PRODUCT-INV-SPEC-001 identified Index purity, parent reference grammar, stable `spec:` ID-as-ref compatibility, and ownership-sensitive traceability concepts as design dependencies. | This work must split dependencies explicitly. |
| Opus design review recommended PRODUCT-INV-SPEC-002 for ownership boundary. | Ownership boundary should be separated from the format contract gate. |
| Current roadmap requires PRODUCT spec-format stabilization before app namespace redesign and DRMCP reimplementation. | DRMCP-WORK-SPEC-001 / 002 must be isolated as later implementation-phase follow-ups, while PRODUCT-WORK-SPEC-006 provides temporary validation for migration. |

## Done Condition

| done item | required result |
|---|---|
| Format contract drafted | PRODUCT spec format contract exists or equivalent spec update is drafted. |
| Deliverable path decided | `product/records/spec/concepts/spec-format/index.md` is accepted or replaced with a documented alternative. |
| Kind set decided | Accepted five kinds and deferred four kinds are explicit. |
| Required section matrix defined | Matrix covers `Overview`, `Index`, `Concept`, `Reference`, and `Contract`. |
| Index / Topics contract defined | Required columns, duplicate-parent rule, and Overview+Topics policy are decided. |
| Spec ID-as-ref derivation defined | Path-derived creation-time suggestion, `index.md` omission rule, underscore segment policy, no separate `ref`, and mismatch warning behavior are defined. |
| Parent grammar defined | Safe interim grammar is defined; forbidden parent forms are explicit. |
| Stable spec ID-as-ref boundary preserved | Stable `spec:` IDs remain canonical; derived topic refs are deferred to PRODUCT-WORK-SPEC-002. |
| Follow-up split produced | PRODUCT-WORK-SPEC-002/003/004/005/006, PRODUCT-INV-SPEC-002, existing app namespace redesign refs, and DRMCP-WORK-SPEC-001/002 implementation-phase follow-ups are created or explicitly queued with ordering. |
| No bulk migration | Existing specs are not bulk-migrated in this work. |
| No DRMCP implementation | DRMCP code is not changed in this work. |
| No v01 modification | `v01/records/spec/**` is not modified. |
| Review gate satisfied | Spec drafting task receives review before close. |
