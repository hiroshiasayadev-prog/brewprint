---
scope: docs/spec/concepts/traceability/coverage-mapping.md
status: draft
last_updated: 2026-05-18
summary: >
  coverage mapping set / group / individual mapping、relation vocabulary、
  source / target constraints の MVP schema を定義する。
depends_on:
  - docs/adr/081-requirement-artifacts-and-spec-traceability.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
semantic_refs:
  - spec:trace.coverage-mapping
sections:
  spec:trace.coverage-mapping: Coverage mapping
  spec:trace.coverage-mapping-relations: Relation vocabulary
  spec:trace.coverage-mapping-endpoints: Source / target constraints
---

# Coverage mapping

## Coverage mapping

Coverage mapping は、project-level trace artifact として、design spec / internal design / coverage artifact 間の semantic ref 対応を表す。

Coverage mapping は source of truth ではない。
Coverage mapping は、spec / internal design / YAML などの artifact が所有する意味を上書きしない。

MVP では YAML を active trace 対象にしないため、coverage mapping が扱う主対象は `spec:` / `internal-design:` の対応である。

## Mapping set / group / individual mapping

MVP では以下の3層を想定する。

| layer | identifier | meaning |
|---|---|---|
| mapping set | `coverage:` semantic ref | coverage artifact または coverage topic の semantic key |
| mapping group | `coverage:` semantic ref または file-local group | mapping set 内の任意 grouping |
| individual mapping | `COV-<DOMAIN>-NNN` | 個別 relation entry |

`coverage:` は mapping set / group を指す semantic ref である。
`COV-*` は個別 mapping を参照するための ID-as-ref である。

Mapping set / group の厳密な階層、file layout、複数 file への分割規則は MVP では固定しない。
ただし、individual mapping ID は同一 project 内で一意でなければならない。

## Minimal coverage mapping schema

Individual mapping の最小 schema は以下とする。

```yaml
id: COV-TRACE-001
source: spec:trace.semantic-ref
relation: maps_to
target: internal-design:resolver.semantic-ref-index
note: optional
```

| field | required | type | meaning |
|---|---:|---|---|
| `id` | yes | string | individual coverage mapping ID |
| `source` | yes | semantic ref | relation source |
| `relation` | yes | string | relation vocabulary |
| `target` | yes | semantic ref | relation target |
| `note` | no | string | human-readable rationale / explanation |

`id` は `COV-<DOMAIN>-NNN` 形式を使う。
`source` / `target` は active prefix を持つ semantic ref でなければならない。

## Relation vocabulary

MVP の relation vocabulary は以下に限定する。

```yaml
relations:
  - maps_to
  - covers
```

`validates` relation は MVP には含めない。
検証機能そのものを不要とする意味ではなく、coverage relation vocabulary としての `validates` edge を採用しないという意味である。

## `maps_to`

`maps_to` は、異なる artifact layer 間で、同じ概念・責務・実装写像に対応することを示す。

Example:

```yaml
id: COV-TRACE-001
source: spec:trace.semantic-ref
relation: maps_to
target: internal-design:resolver.semantic-ref-index
```

MVP では `maps_to` の方向性を完全には固定しない。
ただし、同一 project 内では一貫した方向を使うべきである。
推奨方向は、より外部 contract に近い artifact から、より internal な artifact へ向けることである。

```text
spec -> internal-design
```

## `covers`

`covers` は、ある artifact が対象 semantic ref の意味範囲を表現・包含していることを示す。

Example:

```yaml
id: COV-TRACE-002
source: coverage:trace.semantic-ref
relation: covers
target: spec:trace.semantic-ref
```

`covers` は MVP では使用頻度が低い可能性がある。
YAML が active 化されるまでは、`maps_to` が中心になる見込みである。

`covers` と `coverage:` prefix は語感が近いため、将来の運用で混乱が大きい場合は relation vocabulary の rename を検討する。

## Source / target constraints

MVP coverage mapping の `source` / `target` に置けるのは、active prefix を持つ semantic ref のみである。

Allowed:

```text
spec:...
internal-design:...
coverage:...
```

Not allowed in MVP:

```text
yaml:...
REQ-...
WORK-...
COV-...
fixture:...
```

`COV-*` は mapping entry の `id` に使うものであり、`source` / `target` には使わない。

`REQ-*` / `WORK-*` は requirement / work item metadata 内の参照として使えるが、MVP coverage edge endpoint にはしない。

## Coverage file shape

MVP の coverage file は、front matter と mappings list を持つ YAML または Markdown + YAML block として実装できる。

MVP 推奨形は YAML file である。

```yaml
---
scope: docs/coverage/traceability/semantic-ref.yaml
status: draft
last_updated: 2026-05-18
summary: >
  trace semantic ref の coverage mapping set。
depends_on:
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
semantic_refs:
  - coverage:trace.semantic-ref
---

mappings:
  - id: COV-TRACE-001
    source: spec:trace.semantic-ref
    relation: maps_to
    target: internal-design:resolver.semantic-ref-index
```

ただし、coverage file の physical layout は外部 contract ではない。
MCP や resolver の外部 contract は semantic query interface とする。

## Validation rules

Coverage mapping validator は少なくとも以下を検査する。

- `id` が存在する
- `id` が `COV-<DOMAIN>-NNN` 形式に合う
- `id` が project 内で一意である
- `source` / `target` が semantic ref grammar に従う
- `source` / `target` が active prefix を持つ
- `relation` が allowed relation vocabulary に含まれる
- `source` / `target` が resolver で解決できる

Validation は `validates` relation ではない。
これは trace metadata YAML / coverage mapping schema の検査である。

## Out of scope

MVP では以下を扱わない。

- `yaml:` endpoint
- fixture / golden endpoint
- `validates` relation
- requirement / work item endpoint
- relation ontology の拡張
- generated human-readable coverage table の形式
