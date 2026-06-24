# Overview: Brewprint project specifications

- **id**: `spec:product.brewprint`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product`

## What this is

Owner of Brewprint-specific profile, current state, namespace assignments, and compatibility history.
This area instantiates generic Design Records contracts without redefining them.

## Current contract

| area | ownership |
|---|---|
| Brewprint profile | Project-specific applications, domains, and ownership assignments. |
| Current repository state | Observed namespace directories and repository support areas. |
| Current namespace assignments | Active Brewprint app and domain registry facts. |
| Brewprint compatibility | V01 compatibility, historical attribution, issued-ID retention, and migration state. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Current Brewprint repository layout | Reference | `spec:product.brewprint.layout` | Current app namespace directories and repository support areas. |
| Brewprint namespace profile | Overview | `spec:product.brewprint.namespaces` | Current app and domain namespace assignments for Brewprint. |
| Brewprint compatibility | Overview | `spec:product.brewprint.compatibility` | V01 compatibility, historical attribution, issued-ID retention, and migration state. |

## Non-goals

- Generic Design Records identity, authoring, format, placement, traceability, or artifact responsibility rules.
- DRMCP parser, storage, UI, tool, or authoring behavior.
- Canonical BPDSL language, schema, resolver, render, generation, runtime, or MCP contracts.

## Rules

- Brewprint profile may instantiate `spec:product.design_records` contracts.
- Current app and domain registry tables belong only in this area.
- Project facts must be labeled as profile, inventory, or compatibility content.
- Cross-owner contracts remain pointers instead of copied contract text.
- BPDSL migration judgments do not belong in this area.

## Boundary

| content | owner |
|---|---|
| App-independent Design Records semantics | `spec:product.design_records`. |
| DRMCP operational contracts | DRMCP app-local specifications. |
| Canonical BPDSL internals | BPDSL app-local specifications. |
| Temporary unresolved BPDSL preservation | `spec:product.bpdsl`. |

## Topic map

| child area | purpose | current state |
|---|---|---|
| `layout/` | Current Brewprint repository inventory. | Existing. |
| `namespaces/` | Current Brewprint app and domain assignments. | Existing. |
| `compatibility/` | Brewprint V01 compatibility and migration history. | Existing. |

## Related specs

| ref | relation |
|---|---|
| `spec:product` | PRODUCT placement router and cross-area dependency direction. |
| `spec:product.design_records` | Generic contracts instantiated by Brewprint profile. |
| `spec:product.bpdsl` | Temporary BPDSL preservation boundary. |
