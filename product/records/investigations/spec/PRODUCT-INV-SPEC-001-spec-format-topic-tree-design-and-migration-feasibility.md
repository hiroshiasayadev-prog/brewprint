# PRODUCT-INV-SPEC-001: Spec format topic tree design and migration feasibility

- **status**: concluded
- **date**: 2026-06-10
- **trigger**: PRODUCT-REQ-SPEC-001
- **scope**: Evaluate the proposed MCP-readable spec H1 format, spec kind set, Index Topics model, stable `spec:` compatibility, validation feasibility, and sampled migration impact.
- **non_scope**: Existing spec migration, Design Records MCP implementation, derived topic reference contract finalization, authoring guide updates, and migration execution.
- **source_refs**:
  - PRODUCT-REQ-SPEC-001
- **follow_up_candidates**:
  - PRODUCT-WORK-SPEC-001
  - PRODUCT-WORK-SPEC-002
  - PRODUCT-WORK-SPEC-003
  - PRODUCT-WORK-SPEC-004
  - PRODUCT-WORK-SPEC-005
  - DRMCP-WORK-SPEC-001
  - DRMCP-WORK-SPEC-002

## Investigation Scope

This investigation evaluates whether PRODUCT-REQ-SPEC-001 can be turned into a practical spec format and migration plan without prematurely rewriting current PRODUCT / DRMCP / BPDSL specs.

The investigation treats the proposed format as a candidate, not as an accepted contract.

## Non-Scope

- No existing spec is migrated.
- No requirement is modified.
- No MCP behavior is implemented.
- No derived topic reference contract is finalized.
- This investigation does not perform migration execution.

## Evidence Reviewed

Primary source:

- `product/records/requirements/PRODUCT-REQ-SPEC-001-mcp-readable-spec-format-and-topic-tree.md`

Representative spec files:

- `drmcp/records/spec/design-records-mcp/schema.md`
- `drmcp/records/spec/design-records-mcp/tools.md`
- `drmcp/records/spec/design-records-mcp/overview.md`
- `product/records/spec/concepts/project-artifact-model/index.md`
- `product/records/spec/concepts/namespace-model/index.md`
- `product/records/spec/concepts/traceability/index.md`
- `product/records/spec/concepts/traceability/artifact-refs.md`
- `product/records/spec/concepts/traceability/coverage-mapping.md`
- `product/records/spec/concepts/traceability/metadata-schema.md`
- `product/records/spec/concepts/traceability/out-of-scope.md`
- `product/records/spec/concepts/traceability/resolve-and-validation.md`
- `product/records/spec/concepts/traceability/semantic-ref.md`
- `bpdsl/records/spec/overview.md`
- `bpdsl/records/spec/nodes.md`
- `bpdsl/records/spec/edges.md`
- `bpdsl/records/spec/mcp/schema.md`
- `bpdsl/records/spec/mcp/tools/inspect.md`
- `bpdsl/records/spec/views/dag.md`
- `bpdsl/records/spec/views/sequence-diagram.md`

Supporting implementation / guide evidence:

- `drmcp/src/internal/designrecords/parser.go`
- `drmcp/src/internal/designrecords/validation.go`
- `drmcp/src/internal/designrecords/resolver.go`
- `drmcp/src/internal/designrecords/types.go`
- `drmcp/src/internal/designrecords/authoring_guidance.go`
- `drmcp/src/internal/designrecords/config.go`
- `v01/records/guides/investigation-authoring.md`
- `v01/records/guides/spec-authoring.md`
- `v01/records/investigations/README.md`
- corresponding `v01/records/spec/**` snapshots for the sampled specs

Guide availability finding:

- Design Records MCP returned no authoring guides in the current PRODUCT-oriented context.
- `get_authoring_guidance("investigation-authoring")` returned `guide_not_found`.
- `v01/records/guides/investigation-authoring.md` was used only as legacy format evidence. PRODUCT-REQ-SPEC-001 and the explicit requested path / ID take priority.

MCP resolver finding:

- `resolve_reference("PRODUCT-REQ-SPEC-001")` returned `unsupported_reference`.
- This matches PRODUCT-REQ-SPEC-001's warning that current MCP may not resolve PRODUCT records yet, so this investigation uses explicit filesystem paths as evidence.

## Parser Scan Summary

Parser-aware scan means YAML front matter and fenced code blocks were ignored when counting real headings.

The `fenced headings` column counts ATX heading-looking lines inside fenced code blocks. These lines are not real Markdown headings for validation purposes.

| file | real H1 | proposed H1 format | `## What this is` | fenced headings | v01 snapshot |
|---|---:|---|---|---:|---|
| `drmcp/records/spec/design-records-mcp/schema.md` | 1 | no | no | 3 | different |
| `drmcp/records/spec/design-records-mcp/tools.md` | 1 | no | no | 0 | different |
| `drmcp/records/spec/design-records-mcp/overview.md` | 1 | no | no | 0 | different |
| `product/records/spec/concepts/project-artifact-model/index.md` | 1 | no | no | 0 | same |
| `product/records/spec/concepts/namespace-model/index.md` | 1 | no | no | 0 | same |
| `product/records/spec/concepts/traceability/index.md` | 1 | no | no | 0 | same |
| `product/records/spec/concepts/traceability/artifact-refs.md` | 1 | no | no | 0 | same |
| `product/records/spec/concepts/traceability/coverage-mapping.md` | 1 | no | no | 0 | same |
| `product/records/spec/concepts/traceability/metadata-schema.md` | 1 | no | no | 0 | same |
| `product/records/spec/concepts/traceability/out-of-scope.md` | 1 | no | no | 0 | same |
| `product/records/spec/concepts/traceability/resolve-and-validation.md` | 1 | no | no | 0 | same |
| `product/records/spec/concepts/traceability/semantic-ref.md` | 1 | no | no | 0 | same |
| `bpdsl/records/spec/overview.md` | 1 | no | no | 0 | same |
| `bpdsl/records/spec/nodes.md` | 1 | no | no | 7 | same |
| `bpdsl/records/spec/edges.md` | 1 | no | no | 6 | same |
| `bpdsl/records/spec/mcp/schema.md` | 1 | no | no | 0 | same |
| `bpdsl/records/spec/mcp/tools/inspect.md` | 1 | no | no | 0 | same |
| `bpdsl/records/spec/views/dag.md` | 1 | no | no | 13 | same |
| `bpdsl/records/spec/views/sequence-diagram.md` | 1 | no | no | 7 | same |

Findings:

- All sampled spec files already satisfy the proposed "exactly one real ATX H1 outside front matter and fenced code" rule.
- None of the sampled spec files already use `# <SpecKind>: <Title>`.
- None of the sampled spec files already has `## What this is`.
- BPDSL specs contain real-looking headings inside fenced examples. `nodes.md`, `edges.md`, `views/dag.md`, and `views/sequence-diagram.md` prove that validation must be parser-aware.
- PRODUCT and BPDSL sampled specs are byte-identical to their `v01/records/spec/**` snapshots. They carry historical `docs/...` paths in front matter and links.
- DRMCP sampled specs differ from v01 and partially reflect namespace-prefix / records-root migration, but still contain some old `docs/...` references.

## Current Implementation Feasibility

Existing DRMCP code already has useful parser primitives:

- `parseSpecH1` reads H1 outside YAML front matter and fenced code.
- `extractHeadings` returns parser-aware headings.
- `contentLinesOutsideFrontMatterAndFences` performs the front matter / fenced code exclusion.
- `semanticRefDiagnostics` already validates `semantic_refs` / `sections` grammar, uniqueness, and section heading targets.
- `requiredNarrativeSectionDiagnostics` already implements status-gated required section checks for workflow artifacts.

Reuse of `semanticRefDiagnostics` for topic-related validation depends on the later compatibility design: derived topic refs must either remain separate from stable `spec:` refs or receive explicit alias / redirect / collision rules before resolver validation can share the same path.

Gaps:

- `parseSpecH1` returns the first spec H1 and does not validate exact H1 count.
- Existing spec H1 validation accepts `# <title>` and does not validate `# <SpecKind>: <Title>`.
- There is no Topics table parser or Index child graph validator.
- Investigation reference validation currently does not cleanly support PRODUCT-prefixed IDs in metadata, even though resolver behavior has some namespace-prefix support.

## Kind-Fit Matrix

| kind candidate | evidence from current specs | fit | recommendation |
|---|---|---|---|
| `Overview` | `drmcp/.../overview.md`, `bpdsl/records/spec/overview.md`, traceability index scope sections | strong | Keep in the initial set. |
| `Index` | `product/.../traceability/index.md` has a split-spec table, but also owns real scope and term content | medium | Keep, but require `Index` to be navigation-first. Existing hybrid index files need split or an allowed Overview+Topics pattern. |
| `Concept` | PRODUCT concept specs and trace semantic-ref / coverage boundary docs | strong | Keep in the initial set. |
| `Reference` | metadata schema, artifact refs, BPDSL node / edge / schema tables | strong | Keep in the initial set. |
| `Contract` | DRMCP tools, DRMCP schema behavior, BPDSL MCP tools, render rules | strong | Keep in the initial set. |
| `Guide` | Current authoring guidance lives under `v01/records/guides`, not current spec samples | weak | Defer as an initial spec kind. Revisit if guides become spec records. |
| `Process` | Project artifact flow and workflow lifecycle content exist, but not as standalone process specs | weak | Defer. Model process-like sections under Concept until a standalone process spec appears. |
| `Architecture` | Namespace / DRMCP overview include architecture-like boundaries, but no standalone architecture spec is required | weak | Defer. Do not force an Architecture kind for current migration. |
| `Glossary` | Traceability index has terms, but no standalone glossary file is sampled | weak | Defer or allow later. Term sections can remain in Overview / Concept initially. |

Recommended initial supported kind set:

```text
Overview
Index
Concept
Reference
Contract
```

Deferred kind set:

```text
Guide
Process
Architecture
Glossary
```

## H1 Format Evaluation

The candidate H1 format is implementable:

```markdown
# <SpecKind>: <Title>
```

Evidence:

- Current sampled specs already have exactly one real H1, so the parser-aware count rule is low-risk.
- The H1 rewrite itself is mechanically small for direct migrations.
- Current files can be reasonably classified by content into the five recommended initial kinds.

Risk:

- Some existing files are hybrid. In particular, `product/records/spec/concepts/traceability/index.md` is both an index and an overview/scope document. If `Index` is defined as "principally navigation and no substantive specification body", this file cannot be migrated by an H1 rewrite alone.
- `schema.md` files mix reference tables, behavior contracts, and diagnostic categories. A strict one-kind-per-file model may force premature splits.

Recommendation:

- Accept the H1 format as the target for new and migrated specs.
- Do not use inferred kind as authoritative metadata before a file's H1 is migrated.
- Add a transition period where current `# <title>` specs are still readable but produce migration diagnostics.

## Required `What this is` Evaluation

The `## What this is` section is a good universal target, but it is unsafe as immediate hard validation.

Evidence:

- 0 of 19 sampled spec files currently have `## What this is`.
- Many files have equivalent first sections such as `## Purpose`, `## 目的`, or `## スコープ`.
- Adding the section is usually low-risk because it can summarize existing scope without changing contract.

Recommendation:

- Require `## What this is` for newly authored specs and migrated specs.
- During migration, classify missing `## What this is` as a warning or migration diagnostic, not a hard error across the existing tree.
- Avoid accepting localized aliases as canonical. Aliases make validation and future agents less reliable. Existing localized sections can remain, but the canonical section should be exact English text.

## Index Topics Table Evaluation

Proposed table:

| title | kind | parent | file | summary |
|---|---|---|---|---|

Evaluation:

- `title`: Useful for readable navigation, but stale unless validated against child H1 title.
- `kind`: Useful for topic classification, but stale unless validated against child H1 kind.
- `parent`: Can be the authoritative parent relation, but needs a stable parent reference grammar before derived topic refs are finalized.
- `file`: Practical child target, but stale on moves unless validation checks existence and normalized path.
- `summary`: Valuable for MCP preview, but cannot be fully mechanically validated against content.

Recommendation:

- Use an `Index` spec with `## Topics` as the authoritative topic parent source for a bounded topic tree.
- Require `title`, `kind`, `parent`, `file`, and `summary` columns.
- Validate `file`, child H1 kind/title, duplicate parents, and parent metadata mismatch mechanically.
- Treat `summary` as human-maintained visible metadata with only presence / non-empty validation.
- Do not make front matter `sections` or derived topic refs the parent source.

Open risk:

- A pure Index rule conflicts with current `index.md` files that also own overview and scope content. The follow-up spec must either split those files or allow an `Overview` file to contain a `## Topics` table.

## Stable `spec:` Refs and Derived Topic Refs

Current traceability specs define stable `spec:` refs as concept identifiers:

- They are not physical paths.
- They are not Markdown headings.
- They survive heading rename, file rename, section move, document split, and document merge.
- Section-level refs are currently declared via front matter `sections`.

Compatibility matrix:

| operation | stable `spec:` behavior today | derived topic ref risk | recommendation |
|---|---|---|---|
| Rename title / H1 | Existing `spec:` ref should remain stable | Title-derived ref changes unless alias exists | Do not replace stable `spec:` refs with title-derived refs. |
| Move file | Existing `spec:` ref should remain stable | Path-derived or parent-derived ref may change | Keep stable refs path-independent. |
| Split document | Existing ref remains on the nearest successor document / section | Derived refs may fan out or collide | Require explicit compatibility design before using derived refs as canonical. |
| Merge document | Multiple stable refs may point to one section | Derived ref may lose one identity | Allow multiple stable refs to target one topic. |
| Collision | Existing validation catches duplicate stable semantic refs | Normalized title collisions are likely | Derived refs need collision diagnostics and disambiguation rules. |
| Alias / redirect | Reserved in trace semantic-ref spec, not fully defined | Required if derived refs become public links | Define before any replacement. |
| Legacy `semantic_refs` / `sections` | Current resolver depends on them | Replacement would break current references | Keep them until migration compatibility is specified. |

Recommendation:

- Keep existing stable `spec:` semantic refs as canonical references.
- Treat derived topic refs as navigation/display identifiers until a later compatibility design says otherwise.
- If derived refs become resolver inputs, they should resolve to topics and may link to stable `spec:` refs, but they should not silently replace stable `spec:` refs.

## Validation Rule Matrix

| rule | class | feasibility | notes |
|---|---|---|---|
| missing required section | mechanical | high | Existing required-section validation pattern can be reused. Start as migration warning for old specs. |
| invalid spec H1 format | mechanical | high | Add regex for `# <SpecKind>: <Title>`. |
| invalid H1 count outside front matter and fenced code | mechanical | high | Existing parser-aware line filtering can be reused, but exact count is not implemented yet. |
| invalid or missing Index Topics table column | mechanical | medium | Requires Markdown table parser. Avoid ad hoc string parsing in implementation. |
| unresolved child target | mechanical | high | Check normalized path from Index row to existing file. |
| H1 / Topics table kind mismatch | mechanical | high | Requires child H1 parser and Index table parser. |
| duplicate child parent | mechanical | medium | Build child-file to parent mapping across Index tables. |
| orphan child topic | semantic | medium | Requires policy for which files must appear in which Index. Do not infer globally from path alone. |
| parent metadata mismatch | mechanical | medium | Valid only after MCP-managed visible parent metadata format is specified. |
| topic cycle | deferred mechanical | medium | Graph cycle detection is simple after parent reference grammar is fixed; defer until Index parent refs are specified. |
| summary stale relative to body | semantic / deferred | low | Presence can be checked; truthfulness cannot be reliably mechanical. |
| derived topic ref collision | mechanical after design | medium | Requires normalization and disambiguation contract. |
| stable `spec:` / derived topic ref conflict | semantic | medium | Needs compatibility design before enforcement. |

## Migration Classification

| sampled file | candidate kind | classification | reason |
|---|---|---|---|
| `drmcp/records/spec/design-records-mcp/overview.md` | Overview | directly migratable | Clear overview role. Needs H1 rewrite and `What this is`; also cleanup of residual old path wording. |
| `drmcp/records/spec/design-records-mcp/schema.md` | Reference | split recommended | Contains metadata source, ID normalization, field definitions, record model, diagnostics, and bootstrap metadata. Reference kind fits, but topic-level reading would benefit from later split. |
| `drmcp/records/spec/design-records-mcp/tools.md` | Contract | split recommended | Tool contracts are clear, but the file is large and per-tool topics are natural child contracts. |
| `product/records/spec/concepts/project-artifact-model/index.md` | Concept | directly migratable with cleanup | Concept role is clear. It is byte-identical to v01 and still contains `docs/...` history / links. Filename `index.md` should not force `Index` kind. |
| `product/records/spec/concepts/namespace-model/index.md` | Concept | directly migratable with cleanup | Concept role is clear. It contains broad catalog content but still fits Concept. v01-identical snapshot status means path / ownership wording needs review. |
| `product/records/spec/concepts/traceability/index.md` | Index or Overview | forcing format candidate revision | Current file is both topic entrypoint and substantive scope / terms spec. A pure Index format would require split or an Overview+Topics allowance; final classification depends on resolving the Index purity question in the follow-up format spec. |
| `product/records/spec/concepts/traceability/artifact-refs.md` | Reference | directly migratable | Prefix and ID-as-ref catalog fits Reference. |
| `product/records/spec/concepts/traceability/coverage-mapping.md` | Concept | directly migratable | Owns a boundary concept and reintroduction triggers. |
| `product/records/spec/concepts/traceability/metadata-schema.md` | Reference | directly migratable | Field tables and metadata rules fit Reference. |
| `product/records/spec/concepts/traceability/out-of-scope.md` | Reference | directly migratable | Catalog of deferred scope fits Reference. Could also remain Concept if grouped by boundary. |
| `product/records/spec/concepts/traceability/resolve-and-validation.md` | Contract | directly migratable | Resolver and validation behavior fit Contract. |
| `product/records/spec/concepts/traceability/semantic-ref.md` | Concept | directly migratable | Stable identity model and rules fit Concept; grammar subsections can remain inside the Concept. |
| `bpdsl/records/spec/overview.md` | Overview | directly migratable | Clear overview role. |
| `bpdsl/records/spec/nodes.md` | Reference | directly migratable, split candidate | Node kind definitions fit Reference. Large per-node sections could later become child Reference topics. |
| `bpdsl/records/spec/edges.md` | Reference | directly migratable | Edge rule catalog fits Reference. Fenced H1 examples make parser-aware validation mandatory. |
| `bpdsl/records/spec/mcp/schema.md` | Reference | directly migratable | Common schema tables fit Reference. |
| `bpdsl/records/spec/mcp/tools/inspect.md` | Contract | directly migratable | Single tool behavior contract fits Contract. |
| `bpdsl/records/spec/views/dag.md` | Contract | directly migratable | Render rule contract fits Contract. Fenced output-format headings must be ignored. |
| `bpdsl/records/spec/views/sequence-diagram.md` | Contract | directly migratable | Render rule contract fits Contract. Fenced example headings must be ignored. |

Cross-cutting migration findings:

- Ownership relocation is not finalized by this migration table. The table classifies format-migration feasibility, not final PRODUCT / DRMCP ownership.
- `product/records/spec/concepts/traceability/**` remains ownership-sensitive: the follow-up ownership relocation decision must decide whether traceability concept specs stay PRODUCT-owned as cross-app semantics or move partly under DRMCP as tool-owned resolver semantics.
- Direct migration usually means H1 rewrite plus `## What this is`; it does not mean the content is clean.
- PRODUCT and BPDSL sampled files being v01-identical is a stale-doc signal, not a blocker for format migration.
- Files named `index.md` must not automatically imply `Index` kind.
- Some large Reference / Contract files can be migrated first and split later. Requiring split before format adoption would cause unnecessary scope creep.

## Recommended Later WORK Items

Recommended follow-up work item candidates:

| candidate | purpose |
|---|---|
| PRODUCT-WORK-SPEC-001 | Define accepted H1 format, supported kinds, required sections, Index Topics table, migration phase, validation severity, Index purity vs Overview+Topics, parent reference grammar, and exact follow-up split. |
| PRODUCT-WORK-SPEC-003 | Update spec authoring guidance after the format spec is accepted. |
| DRMCP-WORK-SPEC-001 / DRMCP-WORK-SPEC-002 | Add parser-aware H1 count, spec kind H1 validation, required section validation, Index Topics validation, and parent mismatch diagnostics. |
| PRODUCT-WORK-SPEC-002 | Decide how stable `spec:` refs, derived topic refs, aliases, redirects, split, merge, and collisions coexist. |
| PRODUCT-WORK-SPEC-004 / PRODUCT-WORK-SPEC-005 | Decide ownership-sensitive traceability boundaries, then apply H1 / `What this is` migration, split hybrid index files where necessary, and clean stale `docs/...` references. |

## Recommendation

The proposed format is implementable, but the safe path is staged:

1. Start with PRODUCT-WORK-SPEC-001 to accept the initial kind set `Overview`, `Index`, `Concept`, `Reference`, and `Contract`, and to produce the exact follow-up WORK / prerequisite INV split.
2. Defer `Guide`, `Process`, `Architecture`, and `Glossary` until current specs prove a need.
3. Keep stable `spec:` refs as canonical references.
4. Treat derived topic refs as non-canonical navigation identifiers until a later compatibility design exists.
5. Make `## What this is` required for new / migrated specs, but do not hard-fail the existing tree immediately.
6. Permit direct migration for coherent large specs before optional later splitting.
7. Resolve the `Index` purity question before enforcing `## Topics` globally.

## Unresolved Questions

- Should an `Overview` spec be allowed to contain `## Topics`, or must all authoritative topic tables live only in `Index` specs?
- What exact grammar should `parent` use in an Index Topics table before derived topic refs are specified? This is the first design dependency for safe Index graph validation.
- Should topic cycle validation ship in the first DRMCP validation implementation or be explicitly deferred?
- How should PRODUCT-prefixed requirement / investigation IDs be validated in investigation metadata while current MCP validation still has namespace-prefix gaps?

