---
scope: docs/spec/concepts/traceability/semantic-ref.md
status: draft
last_updated: 2026-05-23
summary: >
  semantic ref の grammar、安定性、document-level / section-level ref、
  redirect / superseded の基本方針を定義する。
depends_on:
  - docs/adr/081-requirement-artifacts-and-spec-traceability.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
semantic_refs:
  - spec:trace.semantic-ref
sections:
  spec:trace.semantic-ref: Semantic ref
  spec:trace.semantic-ref-grammar: Semantic ref grammar
  spec:trace.semantic-ref-stability: Stability rules
---

# Semantic ref

## Semantic ref

Semantic ref は、brewprint docs artifact が表す概念を安定して参照するための identifier である。

Semantic ref は physical path ではない。
Markdown heading でもない。
Directory layout でもない。

File rename、document split、document merge、section move が発生しても、同一概念を指す限り semantic ref は維持されるべきである。

例:

```text
spec:trace.semantic-ref
internal-design:resolver.semantic-ref-index
coverage:trace.semantic-ref
```

Semantic ref と artifact ID-as-ref は区別する。`ADR-*` / `SPEC-*` / `INV-*` は design record artifact を指す ID-as-ref であり、`spec:` / `internal-design:` / `coverage:` のような semantic ref prefix ではない。

ADR-087 により、investigation の `source_refs` および記載済み `follow_up_results` は、対象に応じて artifact ID-as-ref または semantic ref を canonical reference として用いる。physical path は canonical reference として用いない。

## Semantic ref grammar

MVP の prefix-ref grammar は以下とする。

```text
<prefix>:<domain>.<concept>[.<subconcept>...]
```

例:

```text
spec:trace.semantic-ref
spec:trace.coverage-mapping
internal-design:resolver.semantic-ref-index
coverage:trace.semantic-ref
```

### Character rules

MVP の semantic ref は以下の文字だけを使う。

```text
a-z 0-9 - . :
```

制約:

- prefix は lowercase ASCII とする
- domain / concept / subconcept は lowercase ASCII とする
- word separator は hyphen `-` とする
- namespace separator は dot `.` とする
- prefix separator は colon `:` とする
- whitespace は使わない
- slash `/` は使わない
- file extension は含めない
- physical path は含めない

MVP では non-ASCII semantic ref を許可しない。
Human-readable title は Markdown heading や metadata field に置く。

## Document-level ref and section-level ref

Semantic ref は document-level と section-level の両方に使える。

Document-level ref は、artifact 全体が表す概念を指す。
Section-level ref は、artifact 内の特定 section が表す概念を指す。

例:

```yaml
semantic_refs:
  - spec:trace
sections:
  spec:trace.semantic-ref: Semantic ref
  spec:trace.coverage-mapping: Coverage mapping
```

`semantic_refs` は document-level ref を宣言する。
`sections` は section-level ref と Markdown heading text の対応を宣言する。

Section-level ref は Markdown heading に `{#anchor}` を直接書かない。
Section anchor は front matter の `sections` mapping で管理する。

## Stability rules

Semantic ref は append-only に扱う。

- 一度発行した semantic ref は別概念に再利用しない
- heading rename では semantic ref を維持する
- section move では semantic ref を維持する
- file rename では semantic ref を維持する
- document split では既存 semantic ref を最も近い後継 document / section に残す
- section split では既存 semantic ref を最も近い後継 section に残し、新しい概念には新しい semantic ref を発行する
- section merge では複数 semantic ref が同一 section を指してよい

この rule は、physical layout 変更と trace の安定性を分離するためのものである。

## Redirect / superseded mapping

Semantic ref の削除や意味変更が必要になった場合、既存 ref を別概念へ再利用してはならない。

MVP では redirect / superseded mapping の完全 schema は定義しない。
ただし、後続 spec で以下を扱う余地を予約する。

```yaml
redirects:
  spec:old.ref: spec:new.ref
superseded:
  spec:old.ref:
    by: spec:new.ref
    reason: optional
```

Redirect / superseded mapping が導入されるまでは、既存 semantic ref はできるだけ残す。

## Out of scope

この file では以下を定義しない。

- active / reserved prefix の完全一覧
- coverage mapping schema
- resolver request / response
- MCP tool contract
- brewprint DSL YAML entity-level ref

これらは sibling spec files で扱う。
