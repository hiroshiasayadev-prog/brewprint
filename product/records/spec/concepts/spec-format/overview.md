# Overview: Spec format overview

- **id**: `spec:product.concepts.spec_format.overview`
- **status**: accepted
- **date**: 2026-06-11
- **parent**: `spec:product.concepts.spec_format`

## What this is

This spec introduces the PRODUCT-level document format contract for spec records under app namespace records trees such as `product/records/spec/**`, `drmcp/records/spec/**`, and `bpdsl/records/spec/**`.

The goal is to make spec files parser-friendly enough for Design Records MCP to extract a topic tree, validate document shape, and narrow review context without forcing immediate bulk migration of existing specs.

This overview owns the entry-level framing for the spec-format area. The detailed rules are split into focused child specs under `spec:product.concepts.spec_format`.

## Current contract

The spec-format area uses visible markdown structure as the source of truth:

| field class | policy | reason |
|---|---|---|
| YAML front matter | Prohibited for new or migrated specs under this format. | Hidden metadata stales easily and differs from ADR / workflow record authoring style. |
| record identity | H1-adjacent `- **id**:` marker. For new or migrated specs, the value is the path-derived canonical spec ref, for example `spec:product.concepts.spec_format`. | Keeps identity visible and directly referenceable while preserving one-to-one path/ref mapping. |
| spec kind | H1 prefix, for example `# Concept: ...`. | Avoids duplicating kind in hidden metadata. |
| status/date | H1-adjacent markers. | Matches ADR / workflow record style and keeps lifecycle metadata visible. |
| semantic traceability | Specs do not carry reverse traceability tables. Semantic spec changes are governed by ADR decisions; ADRs may later carry `target_specs`. | Keeps spec content focused on current contract instead of duplicating provenance. |
| work evidence | WORK/TASK `## Evidence` may list edited files and review outputs. | Evidence is work history, not the spec traceability source of truth. |
| topic structure | H1 / H2 / `## Topics` tables. | Topic structure must be visible where the content is read. |
| spec ID refs | Do not use a separate `ref` marker. The H1-adjacent `id` is the canonical spec ID-as-ref for the spec. | ID-as-ref means the ID itself is the reference target. |
| section refs | Parser-aware headings or explicit visible tables. | Section identity must not be hidden in front matter mappings. |

Existing specs may keep front matter during migration, but migrated/new specs under this format must not use front matter as a metadata source of truth. If current DRMCP requires front matter for spec discovery, that is a DRMCP compatibility gap to resolve in follow-up indexing work.

PRODUCT-WORK-SPEC-008 owns investigating which ADRs should receive `target_specs` metadata and how stale ADR-to-spec targets should be handled. That follow-up must not reintroduce spec-local reverse traceability tables under another name.

## Non-goals

| not owned here | owner / reason |
|---|---|
| DRMCP implementation behavior | Later DRMCP implementation-phase work owns durable parser and graph validation. |
| Authoring guide text | PRODUCT-WORK-SPEC-003 owns guide updates and examples. |
| Exceptional spec ref compatibility beyond the local contract | PRODUCT-WORK-SPEC-002 owns alias, redirect, split, merge, move, and legacy compatibility behavior when the path-derived canonical model is insufficient. |
| Existing spec migration | PRODUCT-WORK-SPEC-005 owns bulk migration after prerequisites are ready. |
| Temporary validation tooling | PRODUCT-WORK-SPEC-006 owns bridge tooling before migration. |

## Topic map

| topic | ref |
|---|---|
| Document shape | `spec:product.concepts.spec_format.document_shape` |
| Topics table | `spec:product.concepts.spec_format.topics_table` |
| Spec ID-as-ref | `spec:product.concepts.spec_format.spec_id_as_ref` |
| Validation policy | `spec:product.concepts.spec_format.validation_policy` |
| Follow-up boundary | `spec:product.concepts.spec_format.follow_up_boundary` |

## Related specs

| ref | relation |
|---|---|
| `spec:product.concepts.spec_format` | Parent Index and authoritative topic declaration for this child spec. |
| `spec:product.concepts.namespace_model` | Defines app and domain namespace model. |
| `spec:product.concepts.repository_layout` | Defines namespace-first repository layout. |
