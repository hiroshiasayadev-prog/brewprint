# Overview: Brewprint DSL

- **id**: `spec:bpdsl.overview`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `-`

## What this is

Entry point for the brewprint specification. Brewprint is a YAML-based design language and MCP context layer for human-AI collaboration. The spec area is organized into three peer domains.

## Current contract

| domain | description |
|---|---|
| DSL | YAML language contract: node kinds, edge syntax, file types, name resolution, TypeRef, project layout, and diagnostics. |
| MCP | Tool contract: MCP query tools that supply design context to Claude Code. |
| Views | Render contract: diagram and document view output formats. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Brewprint DSL language | Overview | `spec:bpdsl.dsl.overview` | YAML design language spec — node kinds, edge syntax, file types, name resolution, TypeRef, project layout, diagnostics, and design philosophy. |
| Brewprint MCP | Overview | `spec:bpdsl.mcp.overview` | MCP tool overview, design principles, and tool selection guidance. |
| View render contracts | Overview | `spec:bpdsl.views.overview` | Render contracts for all diagram and document view output formats. |
