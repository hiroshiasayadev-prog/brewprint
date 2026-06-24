# Overview: Design Records

- **id**: `spec:product.design_records`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product`

## What this is

Owner of app-independent Design Records semantics.
The area remains valid when Brewprint applications and implementation tools change.

## Current contract

| semantic area | ownership |
|---|---|
| Identity | Public artifact identity, canonical references, and app-independent namespace concepts. |
| Authoring | Record selection, authoring rules, lifecycle guidance, and prose standards. |
| Format | Visible document shape, metadata, topic organization, and format validation rules. |
| Record placement | App-independent placement and discovery rules for design records. |
| Traceability | Reference meaning, relation integrity, resolution inputs, and validation semantics. |
| Artifact responsibility | Responsibilities and boundaries among design record kinds and workflow artifacts. |

## Non-goals

- DRMCP parser, storage, UI, tool, or authoring behavior.
- Canonical BPDSL language, schema, resolver, render, generation, runtime, or MCP contracts.
- Brewprint current app or domain registry facts.
- Brewprint compatibility history or migration state.

## Rules

- Design Records contracts have no normative dependency on DRMCP.
- Design Records contracts have no normative dependency on BPDSL.
- DRMCP may implement Design Records contracts through app-local specifications.
- Generic examples should be app-neutral and non-normative where practical.
- Current Brewprint registry tables must remain under `spec:product.brewprint`.
- Cross-owner contracts must be referenced by pointer instead of copied.

## Boundary

| content | owner |
|---|---|
| DRMCP request, response, diagnostic, parser, persistence, UI, or tool behavior | DRMCP app-local specifications. |
| BPDSL language, schema, resolver, render, generation, runtime, or MCP behavior | BPDSL app-local specifications. |
| Current Brewprint repository layout and namespace assignments | `spec:product.brewprint`. |
| Brewprint V01 compatibility, historical attribution, issued-ID retention, and migration state | `brewprint/compatibility/`. |
| Unadopted future integration mechanisms | A follow-up design record or deletion after evidence transfer. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Namespace model | Overview | `spec:product.design_records.namespace_model` | App-independent namespace semantics, artifact identity grammar, and ownership selection. |
| Authoring standards | Index | `spec:product.design_records.authoring_standards` | Design record selection, authoring, lifecycle, and writing standards. |
| Spec format | Index | `spec:product.design_records.spec_format` | Visible spec shape, identity, topic organization, and validation policy. |
| Repository layout | Overview | `spec:product.design_records.repository_layout` | App-independent Design Records placement and discovery paths. |
| Traceability | Overview | `spec:product.design_records.traceability` | Canonical references, declared relations, lookup inputs, and validation semantics. |
| Artifact model | Overview | `spec:product.design_records.artifact_model` | Design record responsibilities, workflow roles, source-of-truth boundaries, and tool boundaries. |

## Topic map

The child areas are `namespace-model/`, `authoring-standards/`, `spec-format/`,
`repository-layout/`, `traceability/`, and `artifact-model/`.
All six areas are declared in `## Topics`.

## Related specs

| ref | relation |
|---|---|
| `spec:product` | PRODUCT placement router and cross-area dependency direction. |
| `spec:product.brewprint` | Brewprint-specific instantiation and current-state owner. |
| `spec:product.bpdsl` | Temporary non-canonical BPDSL preservation boundary. |
