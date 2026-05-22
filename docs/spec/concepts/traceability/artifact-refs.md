---
scope: docs/spec/concepts/traceability/artifact-refs.md
status: draft
last_updated: 2026-05-23
summary: >
  traceability MVP の artifact ref 種別、active / reserved prefix、
  ID-as-ref、scope 外の prefix 方針を定義する。
depends_on:
  - docs/adr/081-requirement-artifacts-and-spec-traceability.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
semantic_refs:
  - spec:trace.artifact-refs
sections:
  spec:trace.artifact-refs.active-prefixes: Active prefixes
  spec:trace.artifact-refs.reserved-prefixes: Reserved prefixes
  spec:trace.artifact-refs.id-as-ref: ID-as-ref
---

# Artifact refs

## Purpose

この file は、traceability MVP で扱う artifact ref の種別を定義する。

MVP は、すべての docs artifact を semantic trace graph の一級対象にしない。
まず spec / internal-design / coverage を active prefix として扱い、brewprint DSL YAML は reserve only とする。
Requirement / work item / coverage mapping に加え、Design Records MCP が解決する decision / spec / investigation record は、prefix-ref ではなく artifact ID-as-ref を使う。

## Active prefixes

MVP の active prefix は以下である。

```yaml
active_prefixes:
  - spec
  - internal-design
  - coverage
```

Active prefix を持つ semantic ref は、resolver が実在 artifact / section / mapping へ解決できることを前提とする。

MVP coverage edge の `source` / `target` に置けるのは、active prefix を持つ semantic ref のみである。

### `spec:`

`spec:` は design spec の document-level または section-level semantic ref を表す。

例:

```text
spec:trace
spec:trace.semantic-ref
spec:trace.coverage-mapping
```

`spec:` ref の section-level 解決は、spec front matter の `sections` mapping を使う。
Markdown heading に `{#anchor}` を直接書く方式は使わない。

### `internal-design:`

`internal-design:` は internal design topic / wiring story の semantic ref を表す。

例:

```text
internal-design:resolver.semantic-ref-index
internal-design:mcp.trace-query
```

MVP では、`internal-design:` ref の解決単位を document / section / topic のどれにするかは完全には固定しない。
具体的な解決単位と metadata schema は後続 spec refinement で定義する。

### `coverage:`

`coverage:` は coverage mapping set / mapping group の semantic ref を表す。

例:

```text
coverage:trace.semantic-ref
coverage:trace.internal-design-mapping
```

`coverage:` は coverage artifact の semantic key であり、個別 mapping ID である `COV-*` とは別である。

## Reserved prefixes

MVP の reserved prefix は以下である。

```yaml
reserved_prefixes:
  - yaml
```

### `yaml:`

`yaml:` は、対象 system / design model の brewprint DSL YAML を指す semantic ref として予約する。

ここでいう brewprint DSL YAML は、ADR-083 が primary DSL source と呼ぶ YAML である。
例えば、DAG / model / API / state machine / sequence / future self-hosting model などを表す YAML を指す。

`yaml:` は trace metadata YAML を指さない。

以下は `yaml:` prefix の対象ではない。

- spec front matter
- internal-design front matter
- coverage mapping YAML
- requirement metadata
- work item metadata

MVP では、`yaml:` ref を active trace 対象にしない。
`yaml:` ref の記述を schema 上許容するか、許容する場合に unresolved を error / warning / ignored のどれにするかは、後続 spec refinement で定義する。

## ID-as-ref

MVP では、requirement / work item / individual coverage mapping、および Design Records MCP の record は prefix-ref ではなく ID-as-ref を使う。

```yaml
id_as_ref:
  decision: "ADR-NNN"
  spec_record: "SPEC-<slug>"
  investigation: "INV-<DOMAIN>-NNN"
  requirement: "REQ-<DOMAIN>-NNN"
  work_item: "WORK-<DOMAIN>-NNN"
  coverage_mapping: "COV-<DOMAIN>-NNN"
```

### `ADR-*` / `SPEC-*` / `INV-*`

`ADR-*` は decision record ID、`SPEC-*` は spec record ID、`INV-*` は investigation record ID である。

ADR-087 により、investigation の `source_refs` および記載済み `follow_up_results` は、これらの artifact ID または active semantic ref を canonical reference として使用できる。physical path は canonical reference ではない。

`INV-*` は ADR-086 に従い `INV-<DOMAIN>-NNN` 形式とする。

### `REQ-*`

`REQ-*` は requirement ID である。
Requirement ID は ADR-081 に従い、ADR 番号と結合しない。

例:

```text
REQ-TRACE-001
REQ-MCP-001
REQ-SELFHOST-001
```

MVP では `requirement:` prefix は採用しない。

### `WORK-*`

`WORK-*` は work item ID である。
Work item は source requirement を必ず持つ。

例:

```text
WORK-TRACE-001
WORK-MCP-001
```

MVP では `work-item:` prefix は採用しない。

### `COV-*`

`COV-*` は individual coverage mapping ID である。

例:

```text
COV-TRACE-001
COV-TRACE-002
```

`coverage:` は mapping set / group の semantic ref を指し、`COV-*` は個別 mapping を指す。
Mapping set / group / individual mapping の階層関係は [`coverage-mapping.md`](coverage-mapping.md) で扱う。

## Coverage edge source / target

MVP coverage edge の `source` / `target` に置けるのは active prefix を持つ semantic ref のみである。

Allowed in MVP:

```text
spec:...
internal-design:...
coverage:...
```

Not allowed in MVP coverage edge source / target:

```text
yaml:...
REQ-...
WORK-...
COV-...
fixture:...
```

`ADR-*` / `SPEC-*` / `INV-*` / `REQ-*` / `WORK-*` は metadata 内の参照解決に使えるが、MVP coverage edge の endpoint にはしない。

## Scope-out prefixes

### `fixture:`

`fixture:` は MVP では active prefix でも reserved prefix でもない。
Fixture / golden は brewprint processor / renderer / validator の debug / regression test asset であり、project-level semantic trace graph には含めない。

Fixture-level traceability が必要になった場合は、後続 requirement / work item で別 layer として導入する。

### `requirement:` / `work-item:`

`requirement:` / `work-item:` prefix は MVP では採用しない。
Requirement / work item は ID-as-ref として扱う。

将来、requirement / work item を coverage edge の endpoint として扱う必要が出た場合は、後続 ADR / spec で prefix reserve または active 化を判断する。

## 由来

- ADR-081 §5: requirement ID は ADR 番号と結合しない
- ADR-083 §8: trace layer common principle
- ADR-084 §1〜§6: MVP active / reserved / out-of-scope boundary
- ADR-087 §4〜§8: Design Records MCP resolve responsibility と investigation canonical reference rule
