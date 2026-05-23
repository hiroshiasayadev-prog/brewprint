---
scope: docs/spec/concepts/traceability/artifact-refs.md
status: draft
last_updated: 2026-05-24
summary: >
  traceability MVP の active / reserved / deferred semantic ref と、
  canonical ID-as-ref の境界を定義する。
depends_on:
  - docs/adr/081-requirement-artifacts-and-spec-traceability.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
semantic_refs:
  - spec:trace.artifact-refs
sections:
  spec:trace.artifact-refs.active-prefixes: Active prefixes
  spec:trace.artifact-refs.reserved-prefixes: Reserved prefixes
  spec:trace.artifact-refs.deferred-prefixes: Deferred prefixes
  spec:trace.artifact-refs.id-as-ref: ID-as-ref
---

# Artifact refs

## Purpose

この file は、canonical reference resolution foundation としての traceability MVP が解決する ref 種別を定義する。

MVP は、全 artifact layer や realization relation を semantic trace graph の一級対象にしない。Active semantic ref は `spec:` のみに限定し、record artifact は ID-as-ref で解決する。

## Active prefixes

MVP の active semantic ref prefix は以下のみである。

```yaml
active_prefixes:
  - spec
```

### `spec:`

`spec:` は design spec の document-level または section-level semantic ref を表す。

```text
spec:trace
spec:trace.semantic-ref
spec:trace.resolve-and-validation
```

`spec:` section ref の解決は spec front matter の `sections` mapping を使う。Physical heading anchor を canonical identity として扱わない。

## Reserved prefixes

```yaml
reserved_prefixes:
  - yaml
```

### `yaml:`

`yaml:` は、対象 system / design model を表す brewprint DSL YAML の future semantic endpoint 用に予約する。MVP では active にせず、resolve behavior も固定しない。

Spec front matter や investigation metadata のような trace metadata は `yaml:` の対象ではない。

## Deferred prefixes

以下は artifact layer の存在を否定せず、semantic trace endpoint としての operational contract を将来判断へ送った prefix 候補である。

```yaml
deferred_prefixes:
  - internal-design
  - coverage
```

### `internal-design:`

`docs/internal-design/` は implementation-facing documentation layer として存続する。ただし MVP は internal design document を semantic ref で index / resolve / validate せず、`spec:` との realization relation も扱わない。

### `coverage:`

MVP は external coverage artifact を採用しない。したがって `coverage:` mapping set identity も active / reserved contract としない。External artifact の再導入時に名称を含めて判断する。

## ID-as-ref

MVP が canonical reference として扱う ID-as-ref は、少なくとも Design Records MCP の record artifact ID である。

```yaml
id_as_ref:
  decision: "ADR-NNN"
  spec_record: "SPEC-<slug>"
  investigation: "INV-<DOMAIN>-NNN"
```

### `ADR-*` / `SPEC-*` / `INV-*`

`ADR-*` / `SPEC-*` / `INV-*` は Design Records MCP が index / query / validation の対象とする record artifact ID である。Investigation の `source_refs` および記載済み `follow_up_results` は、これらの ID-as-ref または active `spec:` ref を canonical reference として使用できる。

### `REQ-*` / `WORK-*`

Requirement / work item は stable ID を持つ artifact である。MVP でそれらを Design Records MCP の record kind または semantic relation endpoint に含めるかは確定しない。Metadata 上で参照する具体 rule は、consumer requirement と tool contract が必要になった時点で定義する。

### `COV-*`

`COV-*` は MVP の canonical reference form ではない。External coverage artifact と individual mapping を導入する必要が生じた場合に再判断する。

## Relation endpoint boundary

MVP は semantic realization relation の endpoint を定義しない。`spec:` → `internal-design:` mapping、`maps_to`、`covers`、coverage mapping endpoint constraint は future scope である。

## Scope-out prefixes

`fixture:` は MVP の active prefix でも reserved prefix でもない。Fixture / golden は processor / renderer / validator の検証資産であり、project-level canonical reference foundation には含めない。

`requirement:` / `work-item:` prefix も MVP では採用しない。

## 由来

- ADR-081 §5: requirement ID は ADR 番号と結合しない
- ADR-083 §8: physical path から trace identity を分離する原則
- ADR-087: Design Records MCP resolve responsibility と investigation canonical reference rule
- ADR-088: realization endpoint / external coverage artifact を MVP 外へ送り、canonical reference resolution foundation に縮小
