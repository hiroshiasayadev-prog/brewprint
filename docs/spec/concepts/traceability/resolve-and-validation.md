---
scope: docs/spec/concepts/traceability/resolve-and-validation.md
status: draft
last_updated: 2026-05-18
summary: >
  semantic ref resolver、duplicate / orphan detection、trace metadata YAML validation、
  validates relation との境界を定義する。
depends_on:
  - docs/adr/081-requirement-artifacts-and-spec-traceability.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
semantic_refs:
  - spec:trace.resolve-and-validation
sections:
  spec:trace.resolve: Resolve
  spec:trace.validation: Validation
  spec:trace.validation-not-validates: Validation is not validates
---

# Resolve and validation

## Resolve

Resolve は、semantic ref または ID-as-ref が実在 artifact / section / mapping / metadata entry に解決できるかを確認する処理である。

Resolve は coverage relation ではない。
Resolve は tool / resolver behavior である。

Examples:

```text
spec:trace.semantic-ref -> docs/spec/concepts/traceability/semantic-ref.md § Semantic ref
internal-design:resolver.semantic-ref-index -> docs/internal-design/... 等
coverage:trace.semantic-ref -> docs/coverage/... 等
REQ-TRACE-001 -> docs/requirements/... 等
WORK-TRACE-001 -> docs/work-items/... 等
COV-TRACE-001 -> coverage mapping entry
```

MVP では、active prefix を持つ semantic ref が resolver で解決できることを前提とする。

## Resolver input

Resolver input は、少なくとも以下を扱う。

```text
spec:...
internal-design:...
coverage:...
REQ-...
WORK-...
COV-...
```

ただし、MVP coverage edge の source / target として許可されるのは active prefix を持つ semantic ref のみである。
ID-as-ref は metadata 内の参照解決には使えるが、coverage edge endpoint としては使わない。

`yaml:` は reserve only である。
MVP では `yaml:` ref の resolve behavior を固定しない。

## Resolver output

Resolver output は、少なくとも以下の情報を返せるべきである。

| field | meaning |
|---|---|
| `ref` | 入力 ref |
| `kind` | ref kind。例: `semantic_ref` / `id_ref` |
| `target_type` | 解決先種別。例: `document` / `section` / `mapping` |
| `path` | 解決先 file path |
| `section` | section-level ref の場合の heading text |
| `status` | `resolved` / `unresolved` / `reserved` |

MVP では concrete MCP response schema は定義しない。
Design Records MCP などの tool contract が必要になった時点で別途定義する。

## Lookup sources

Resolver は以下の trace metadata を index source として扱う。

- spec front matter の `semantic_refs`
- spec front matter の `sections`
- internal-design front matter の `semantic_refs`
- coverage artifact front matter の `semantic_refs`
- coverage mapping entry の `id`
- requirement metadata の `id`
- work item metadata の `id`

MVP では自然言語本文から ref を推定しない。

## Section anchor lookup

Section-level ref は front matter の `sections` mapping によって heading text と対応付ける。

Resolver は以下を検査できる。

- `sections` key が semantic ref grammar に従っている
- `sections` value と一致する Markdown heading が存在する
- 同一 file 内で同一 heading text が複数存在する場合の扱い

同一 heading text が複数存在する場合の解決規則は MVP では固定しない。
後続 spec refinement で定義する。

## Duplicate detection

Duplicate detection は、同一 project 内で同じ semantic ref / ID-as-ref が複数の対象を指していないかを検査する。

MVP で duplicate とするもの:

- 同じ semantic ref が複数 document の `semantic_refs` に現れる
- 同じ section-level semantic ref が複数 `sections` key に現れる
- 同じ `COV-*` ID が複数 mapping entry に現れる
- 同じ `REQ-*` ID が複数 requirement metadata に現れる
- 同じ `WORK-*` ID が複数 work item metadata に現れる

Duplicate は error とすることを第一候補とする。
確定した severity は MCP / validator spec で定義する。

## Orphan detection

Orphan detection は、metadata や mapping が参照する ref が解決できない状態を検出する。

Examples:

- coverage mapping の `source` が存在しない
- coverage mapping の `target` が存在しない
- requirement metadata の `source_refs.specs` が存在しない `spec:` ref を指す
- work item metadata の `source_requirement` が存在しない `REQ-*` を指す

MVP では orphan の severity を固定しない。
ただし、active prefix を持つ coverage mapping endpoint の unresolved は error とすることを第一候補とする。

## Reserved ref handling

`yaml:` は reserve only である。

MVP では以下を後続 spec refinement に委ねる。

- `yaml:` ref を trace metadata に書くことを許可するか
- 許可する場合、unresolved `yaml:` を error / warning / ignored のどれにするか
- `yaml:` active 化時の migration rule

`fixture:` は reserved ではない。
`fixture:` ref が project-level trace metadata に現れた場合の扱いは、後続 validator spec で定義する。

## Validation

Validation は、trace metadata YAML が期待 schema に合っているかを検査する処理である。

Examples:

- `semantic_refs` が list<string> であるか
- `sections` が map<string,string> であるか
- coverage mapping が `id` / `source` / `relation` / `target` を持つか
- coverage mapping の `relation` が allowed vocabulary に含まれるか
- `source` / `target` が active prefix を持つか
- ID-as-ref が expected pattern に合うか

Validation は relation vocabulary の `validates` ではない。
Validation は resolver / validator / MCP tool contract の責務である。

## Validation is not validates

`validates` relation を MVP から外すことは、resolve や validation を行わないという意味ではない。

MVP で外すのは、coverage mapping の relation vocabulary としての `validates` edge である。

Scope outside relation vocabulary:

- semantic ref の resolve 可否
- duplicate detection
- orphan detection
- trace metadata YAML schema validation
- MCP writer による artifact generation contract

これらは coverage edge ではなく、tool behavior / schema validation / authoring contract として扱う。

## MCP writer contract placeholder

MCP writer は、trace metadata YAML や coverage mapping を生成・更新する将来 tool である。

MVP では full MCP writer tool schema を定義しない。
ただし、writer tool が導入される場合、tool request schema が artifact authoring contract の一部になる。

将来候補:

- create requirement
- create work item
- add coverage mapping
- register semantic ref
- update section mapping

Writer tool を導入する場合は、dry-run diff、ユーザー確認、conflict handling、format preservation を別 spec / ADR で定義する。

## Out of scope

この file では以下を定義しない。

- concrete MCP request / response schema
- write tool args
- fixture / golden validation
- render expected comparison
- brewprint DSL YAML schema validation
