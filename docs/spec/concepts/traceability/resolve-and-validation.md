---
scope: docs/spec/concepts/traceability/resolve-and-validation.md
status: draft
last_updated: 2026-05-27
summary: >
  canonical semantic/artifact ref resolution、investigation reference validation、
  noncanonical path diagnostic の MVP boundary を定義する。
depends_on:
  - docs/adr/081-requirement-artifacts-and-spec-traceability.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
  - docs/adr/091-workflow-artifact-work-item-task-milestone.md
  - docs/adr/092-design-records-mcp-workflow-artifact-record-and-relation-boundary.md
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
REQ-MCP-003 -> docs/requirements/mcp/REQ-MCP-003-workflow-artifact-mcp-support.md
WORK-MCP-003 -> docs/work-items/mcp/WORK-MCP-003-workflow-artifact-mcp-support.md
TASK-MCP-003-01 -> docs/tasks/mcp/TASK-MCP-003-01-workflow-artifact-mcp-evidence-and-scope.md
```

ADR-087 により、investigation の `source_refs` および記載済み `follow_up_results` に現れる canonical reference は resolve / validation 対象である。ADR-092 により、investigation metadata が追加で扱える workflow ID-as-ref は `REQ-*` / `WORK-*` に限定し、`TASK-*` は investigation metadata canonical reference に含めない。また、workflow artifact 間の declared relation integrity validation を MVP に含める。ADR-088 により、MVP は realization relation や external coverage mapping の解決を要求しない。

## Resolver input

MVP resolver は少なくとも以下を扱う。

```text
spec:...
ADR-...
SPEC-...
INV-...
REQ-...
WORK-...
TASK-...
```

`REQ-*` / `WORK-*` / `TASK-*` は workflow artifact record の direct resolver input として扱う。Workflow artifact 間 relation は metadata の ID-as-ref から検証し、ID 文字列または physical path から親 relation を推測しない。

以下は MVP resolver input として要求しない。

```text
internal-design:...
coverage:...
COV-...
yaml:...
```

`yaml:` は reserve only である。その他は ADR-088 により semantic trace endpoint / mapping mechanism の判断を defer した対象である。

## Resolver output

Resolver の concrete request / response field と status vocabulary は [`docs/spec/design-records-mcp/tools.md`](../../design-records-mcp/tools.md) が所有する。

本 spec は supported canonical reference boundary、lookup source、validation responsibility、および MVP 外とする relation / diagnostic scope のみを定義する。

## Lookup sources

MVP resolver の lookup source は以下である。

- spec front matter の `semantic_refs`
- spec front matter の `sections`
- Design Records MCP が index する record artifact ID (`ADR-*` / `SPEC-*` / `INV-*` / `REQ-*` / `WORK-*` / `TASK-*`)

Resolver は以下の investigation metadata field に記載された ref を validation input として扱う。

- `source_refs`
- `follow_up_results`
- `follow_up_candidates` に artifact reference が記載された場合の canonical form、および unresolved candidate の `info` diagnostic

Investigation metadata で workflow ID-as-ref を参照できる範囲は `REQ-*` / `WORK-*` に限る。`TASK-*` は同 field では unsupported とする。

Workflow relation validation は、requirement の `work_items`、work item の `source_requirement` / `tasks`、task の `work_item` / `source_requirement` / `depends_on` を validation input とする。

Investigation metadata と workflow relation field は参照元であり、参照先を登録する lookup source ではない。

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
- 同じ Design Records MCP record ID (`ADR-*` / `SPEC-*` / `INV-*` / `REQ-*` / `WORK-*` / `TASK-*`) が複数 record に現れる

`COV-*`、`internal-design:`、`coverage:` の duplicate detection は MVP 対象外である。

## Unresolved reference and declared relation integrity

MVP で unresolved error とするもの:

- investigation metadata の `source_refs` が解決不能な supported record ID-as-ref または active `spec:` ref を指す
- investigation metadata の記載済み `follow_up_results` が解決不能な supported record ID-as-ref または active `spec:` ref を指す
- workflow relation field が解決不能な `REQ-*` / `WORK-*` / `TASK-*` を指す

`follow_up_candidates` は未作成 artifact を指しうる。Artifact reference を記載する場合の canonical form は検査するが、unresolved であること自体は error としない。Canonical form の unresolved candidate は、予定された後続 artifact がまだ存在しないことを可視化する `info` diagnostic として返す。

MVP が integrity error とする declared workflow relation mismatch は以下である。

- `requirement.work_items` と `work_item.source_requirement` の不一致
- `work_item.tasks` と `task.work_item` の不一致
- `task.source_requirement` と parent work item の `source_requirement` の不一致

これは metadata に宣言された relation の整合性検査であり、参照されていない requirement / work item / task を探索する orphan diagnostics ではない。Workflow orphan diagnostics は MVP 外とする。

## Reserved / deferred ref handling

`yaml:` は reserve only であり、MVP は記述許可や unresolved severity を固定しない。

`internal-design:` / `coverage:` / `COV-*` は ADR-088 により MVP contract から defer された。既存 draft / example に存在する場合、それらを新規 MVP acceptance target として解決・検証してはならない。

## Validation

MVP validation は以下を扱う。

- `spec:` ref grammar / uniqueness / section lookup。Active `spec:` declaration grammar は root document ref (`spec:trace`) と dot 付き ref (`spec:trace.semantic-ref`) の双方を valid とする
- record ID-as-ref resolution (`ADR-*` / `SPEC-*` / `INV-*` / `REQ-*` / `WORK-*` / `TASK-*`)
- investigation `source_refs` / 記載済み `follow_up_results` の canonicality と resolve。追加 workflow ID-as-ref は `REQ-*` / `WORK-*` に限定する
- investigation `follow_up_candidates` の canonical form と、canonical だが unresolved な candidate の `info` diagnostic。追加 workflow ID-as-ref は `REQ-*` / `WORK-*` に限定する
- investigation metadata に現れる `TASK-*` の unsupported diagnostic
- workflow relation field の target kind / resolution / declared bidirectional consistency
- physical path reference の noncanonical diagnostic。`source_refs` / `follow_up_results` に置かれた場合は error、`follow_up_candidates` に置かれた場合は `info` とする

MVP validation は以下を扱わない。

- coverage mapping YAML schema
- `maps_to` / `covers` / `validates` relation
- `spec:` → `internal-design:` endpoint constraint
- internal-design relation declaration / reverse graph
- workflow orphan diagnostics / progress projection / workflow traversal query / task dependency cycle detection

## Validation boundary

`validates` relation を導入しないことは、resolve / schema validation を行わないことを意味しない。Validation は tool behavior / metadata contract として扱う。

## MCP writer contract placeholder

MVP は writer tools の request / response schema を定義しない。将来 semantic ref registration や investigation metadata update を tool で行う場合は、dry-run diff、confirmation、conflict handling、format preservation を別 spec / ADR で定義する。

## Out of scope

- concrete MCP request / response schema
- writer tool args
- internal-design / coverage / YAML endpoint resolution
- realization relation validation
- workflow orphan diagnostics / progress projection / traversal query / dependency graph projection
- fixture / golden validation
- brewprint DSL YAML schema validation
