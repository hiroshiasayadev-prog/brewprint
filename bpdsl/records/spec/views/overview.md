# Overview: View render contracts

- **id**: `spec:bpdsl.views.overview`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.overview`

## What this is

Entry point for BPDSL view-layer render contracts. Each topic defines how a class of brewprint YAML is rendered into a human-facing output (Markdown table, Mermaid diagram, or HTML fragment) — node/edge selection, layout, labeling, and the conditions under which sections are included or omitted.

## Current contract

| view | input | output format | render scope |
|---|---|---|---|
| Model file | `model/*.yaml` | Markdown tables | One model YAML file → one render. |
| API Table | `as: api_table` view YAML | Markdown tables | Cross-module endpoint listing. |
| State Diagram | `state` / `event` nodes + `transitions:` | Mermaid `stateDiagram-v2` | One YAML file → one FSM diagram. |
| ER Diagram | `store.kind: db` + referenced models | Mermaid `erDiagram` | Module-default, or cross-module via `as: er_diagram` view YAML. |
| DAG | Processing-layer nodes + `flow:` | Mermaid `flowchart TD` + Markdown detail sections | One main-node file → one DAG. |
| Sequence Diagram | `as: sequence_diagram` scenario YAML | Mermaid `sequenceDiagram` | One scenario file → one sequence diagram. |
| Wireframe | `state.wireframe` tree | HTML fragment + fixed CSS | One wireframe tree → one HTML fragment. |

All views render from the same brewprint YAML that DSL nodes/edges specs define — no separate "view source" YAML duplicates node/edge data, except for the cross-cutting view YAML kinds (`api_table`, `er_diagram` (cross-module form), `sequence_diagram`) which select and compose existing nodes rather than defining new ones.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Model file render rules | Contract | `spec:bpdsl.views.model_file` | Markdown render for public models and file-private helper models. |
| API Table render rules | Contract | `spec:bpdsl.views.api_table` | Cross-module endpoint table: collection, route composition, output format. |
| State Diagram render rules | Contract | `spec:bpdsl.views.state_diagram` | Mermaid `stateDiagram-v2` generation from `state` / `event` / `transitions:`. |
| ER Diagram render rules | Contract | `spec:bpdsl.views.er` | Mermaid `erDiagram` generation from `store.kind: db` and referenced models. |
| DAG render rules | Contract | `spec:bpdsl.views.dag` | Mermaid `flowchart TD` plus Markdown detail sections for the Processing layer. |
| Sequence Diagram render rules | Contract | `spec:bpdsl.views.sequence_diagram` | Mermaid `sequenceDiagram` generation from scenario YAML and resolved transitions. |
| Wireframe render rules | Contract | `spec:bpdsl.views.wireframe` | HTML fragment + fixed CSS generation from the `state.wireframe` DSL tree. |
