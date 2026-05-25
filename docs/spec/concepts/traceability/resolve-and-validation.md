---
scope: docs/spec/concepts/traceability/resolve-and-validation.md
status: draft
last_updated: 2026-05-25
summary: >
  canonical semantic/artifact ref resolution、investigation reference validation、
  noncanonical path diagnostic の MVP boundary を定義する。
depends_on:
  - docs/adr/081-requirement-artifacts-and-spec-traceability.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
semantic_refs:
  - spec:trace.resolve-and-validation
sections:
  spec:trace.resolve: Resolve
  spec:trace.validation: Validation
  spec:trace.validation-boundary: Validation boundary
---

# Resolve and validation

## Resolve

Resolve は、canonical semantic ref または artifact ID-as-ref を実在 artifact / spec section へ解決する tool behavior である。Resolve は artifact 間 relation ではない。

MVP example:

```text
spec:trace -> docs/spec/concepts/traceability/index.md
spec:trace.semantic-ref -> docs/spec/concepts/traceability/semantic-ref.md
spec:trace.semantic-ref.definition -> docs/spec/concepts/traceability/semantic-ref.md § Semantic ref definition
spec:project-artifact-model -> docs/spec/concepts/project-artifact-model/index.md
ADR-088 -> docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
INV-DOCS-003 -> docs/investigations/docs/INV-DOCS-003-internal-design-semantic-trace-mvp-necessity.md
```

ADR-087 により、investigation の `source_refs` および記載済み `follow_up_results` に現れる canonical reference は resolve / validation 対象である。ADR-088 により、MVP は realization relation や external coverage mapping の解決を要求しない。

## Resolver input

MVP resolver は少なくとも以下を扱う。

```text
spec:...
ADR-...
SPEC-...
INV-...
```

`REQ-*` / `WORK-*` を resolver public contract に含めるかは、concrete consumer と tool contract が必要になった時点で定義する。

以下は MVP resolver input として要求しない。

```text
internal-design:...
coverage:...
COV-...
yaml:...
```

`yaml:` は reserve only である。その他は ADR-088 により semantic trace endpoint / mapping mechanism の判断を defer した対象である。

## Resolver output

Resolver output は、少なくとも以下の情報を返せるべきである。

| field | meaning |
|---|---|
| `ref` | 入力 canonical ref |
| `kind` | `semantic_ref` または `id_ref` |
| `target_type` | `document` / `section` / `record` 等 |
| `path` | 解決先 file path |
| `section` | section-level `spec:` ref の場合の heading text |
| `status` | `resolved` / `unresolved` / `reserved` |

Concrete MCP request / response schema は Design Records MCP spec と M19 が確定する。

## Lookup sources

MVP resolver の lookup source は以下である。

- spec front matter の `semantic_refs`
- spec front matter の `sections`
- Design Records MCP が index する record artifact ID (`ADR-*` / `SPEC-*` / `INV-*`)

Resolver は以下の investigation metadata field に記載された ref を validation input として扱う。

- `source_refs`
- `follow_up_results`
- `follow_up_candidates` に artifact reference が記載された場合の canonical form、および unresolved candidate の `info` diagnostic

Investigation metadata の参照 field は参照元であり、参照先を登録する lookup source ではない。

MVP では自然言語本文から ref を推定しない。Internal-design front matter や coverage mapping artifact から relation graph を構築しない。

## Section anchor lookup

Section-level semantic ref を active に resolve する対象は `spec:` のみであり、front matter の `sections` mapping により heading text と対応付ける。

Resolver は以下を検査する。

- `sections` key が `spec:` grammar に従っている
- `sections` value と一致する Markdown heading が存在する
- 同一 file 内で解決先が一意である

## Duplicate detection

MVP で duplicate error 候補とするもの:

- 同じ `spec:` document-level ref が複数 document の `semantic_refs` に現れる
- 同じ `spec:` section-level ref が複数 `sections` key に現れる
- 同じ Design Records MCP record ID が複数 record に現れる

`COV-*`、`internal-design:`、`coverage:` の duplicate detection は MVP 対象外である。

## Orphan detection

MVP で unresolved error とするもの:

- investigation metadata の `source_refs` が解決不能な record ID-as-ref または active `spec:` ref を指す
- investigation metadata の記載済み `follow_up_results` が解決不能な record ID-as-ref または active `spec:` ref を指す

`follow_up_candidates` は未作成 artifact を指しうる。Artifact reference を記載する場合の canonical form は検査するが、unresolved であること自体は orphan error としない。Canonical form の unresolved candidate は、予定された後続 artifact がまだ存在しないことを可視化する `info` diagnostic として返す。

## Reserved / deferred ref handling

`yaml:` は reserve only であり、MVP は記述許可や unresolved severity を固定しない。

`internal-design:` / `coverage:` / `COV-*` は ADR-088 により MVP contract から defer された。既存 draft / example に存在する場合、それらを新規 MVP acceptance target として解決・検証してはならない。

## Validation

MVP validation は以下を扱う。

- `spec:` ref grammar / uniqueness / section lookup。Active `spec:` declaration grammar は root document ref (`spec:trace`) と dot 付き ref (`spec:trace.semantic-ref`) の双方を valid とする
- record ID-as-ref resolution
- investigation `source_refs` / 記載済み `follow_up_results` の canonicality と resolve
- investigation `follow_up_candidates` の canonical form と、canonical だが unresolved な candidate の `info` diagnostic
- physical path reference の noncanonical diagnostic。`source_refs` / `follow_up_results` に置かれた場合は error、`follow_up_candidates` に置かれた場合は `info` とする

MVP validation は以下を扱わない。

- coverage mapping YAML schema
- `maps_to` / `covers` / `validates` relation
- `spec:` → `internal-design:` endpoint constraint
- internal-design relation declaration / reverse graph

## Validation boundary

`validates` relation を導入しないことは、resolve / schema validation を行わないことを意味しない。Validation は tool behavior / metadata contract として扱う。

## MCP writer contract placeholder

MVP は writer tools の request / response schema を定義しない。将来 semantic ref registration や investigation metadata update を tool で行う場合は、dry-run diff、confirmation、conflict handling、format preservation を別 spec / ADR で定義する。

## Out of scope

- concrete MCP request / response schema
- writer tool args
- internal-design / coverage / YAML endpoint resolution
- realization relation validation
- fixture / golden validation
- brewprint DSL YAML schema validation
