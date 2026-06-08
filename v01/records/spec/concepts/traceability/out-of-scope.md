---
scope: docs/spec/concepts/traceability/out-of-scope.md
status: draft
last_updated: 2026-05-27
summary: >
  canonical reference resolution foundation の scope 外と、
  future extension を再判断する trigger を定義する。
depends_on:
  - docs/adr/081-requirement-artifacts-and-spec-traceability.md
  - docs/adr/082-golden-fixture-and-self-hosting-requirement-boundary.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
  - docs/adr/091-workflow-artifact-work-item-task-milestone.md
  - docs/adr/092-design-records-mcp-workflow-artifact-record-and-relation-boundary.md
semantic_refs:
  - spec:trace.out-of-scope
sections:
  spec:trace.out-of-scope.boundary: MVP out-of-scope boundary
  spec:trace.future-extensions: Future extensions
---

# Traceability out of scope

## MVP out-of-scope boundary

Semantic trace MVP は canonical reference resolution foundation に限定する。Scope 外の項目は不要という意味ではなく、concrete requirement が確認された時点で後続 ADR / requirement / work item により判断する。

## Internal-design semantic endpoint and realization relation

`docs/internal-design/` artifact layer は存続するが、MVP は以下を扱わない。

- `internal-design:` semantic ref の active 化
- internal design document の canonical ref resolve / validation
- internal design metadata による source `spec:` relation declaration
- `spec:` → `internal-design:` realization mapping
- spec から internal design への reverse graph / impact query

再判断 trigger:

- 複数の internal design document について spec からの機械的 navigation / impact analysis が必要になった場合
- investigation / work item / MCP query が internal design artifact の canonical resolve を必要とする場合
- YAML trace とともに cross-layer realization chain を扱う必要が生じた場合

## External coverage artifact and relation vocabulary

MVP は external relation / assurance artifact、その配置、`coverage:` ref、`COV-*` mapping identity、coverage mapping YAML schema を扱わない。

`maps_to` / `covers` / `validates` を semantic realization relation vocabulary として operational に採用することも延期する。Validation 自体は relation ではなく、canonical reference の tool behavior / metadata contract として MVP に残る。

再判断 trigger:

- gap / completeness / approved relation set を中央管理する必要が生じた場合
- evidence / sign-off / audit snapshot / release baseline を relation と結び付ける必要が生じた場合
- relation entry 自体に stable identity / approval / lifecycle / history が必要になった場合

将来 external artifact を導入する場合、名称を `coverage` とするか、semantic mapping と assurance matrix を分けるか、どの directory に配置するかも判断対象とする。MVP は external artifact 用 directory を予約しない。

## Brewprint DSL YAML entity-level refs

MVP では `yaml:` prefix を active 化しない。

Scope 外:

- brewprint DSL YAML file-level / entity-level semantic ref
- node / edge / view / model / task / asset 単位の semantic ref
- YAML 内 anchor と logical unit resolver rule
- spec / internal design / YAML realization chain

`yaml:` active 化は、self-hosting / UC-002 再構築、または YAML entity ref / resolver rule の concrete requirement が成立した時点で判断する。

## Fixture / golden traceability

MVP では `fixture:` prefix を定義せず、fixture / golden を project-level canonical reference foundation に含めない。

Scope 外:

- fixture semantic ref
- fixture-local coverage の project-level 統合
- golden output と spec semantic ref の対応
- render expected comparison semantics
- test harness schema / golden update workflow

## Workflow semantic prefixes and derived operations

Requirement / work item / task は `REQ-*` / `WORK-*` / `TASK-*` ID-as-ref により Design Records MCP の public record / resolver / declared relation validation 対象となる。一方、MVP は `requirement:` / `work-item:` / `task:` semantic prefix を定義しない。

Workflow artifact support に関して、以下は MVP 外である。

- orphan requirement / orphan work item / orphan task diagnostics
- task status から work item progress を導出する projection
- workflow 専用 traversal / tree / graph query tool
- task dependency cycle detection / execution order projection
- investigation metadata による `TASK-*` canonical reference

これらは declared relation integrity validation または direct ID-as-ref resolution の成立条件ではなく、具体的な運用要件が確認された時点で判断する。

## Full MCP writer tools

MVP では MCP writer tool の request / response schema を定義しない。

Future candidates:

- create requirement / work item
- register `spec:` semantic ref
- update section mapping
- update investigation reference metadata

Writer tool を導入する場合は、dry-run diff、user confirmation、conflict handling、format preservation、write permission boundary を別 spec / ADR で定義する。

## Future extensions

| extension | trigger |
|---|---|
| `internal-design:` active 化 | internal design の canonical navigation / validation が concrete requirement になった場合 |
| realization relation | spec と implementation-facing artifact 間の machine-readable relation が必要になった場合 |
| external relation artifact | gap / evidence / sign-off / lifecycle の中央管理が必要になった場合 |
| `yaml:` active 化 | YAML entity ref / cross-layer trace が必要になった場合 |
| fixture-level traceability | golden fixture と docs/spec の対応を長期管理する必要が生じた場合 |
| workflow semantic prefix | `REQ-*` / `WORK-*` / `TASK-*` ID-as-ref では不足する section-level addressing requirement が生じた場合 |
| workflow orphan / progress / traversal capability | 未接続診断、status 集約 view、専用 graph traversal の concrete requirement が生じた場合 |
| MCP resolve contract refinement | 採用済み canonical reference / workflow relation validation contract の追加 refinement が必要になった場合 |
| MCP writer tools | canonical metadata を tool で生成・更新する必要が生じた場合 |

## Follow-up artifact placement

Future extension が必要になった場合、要求・判断・進捗・contract の owner は既存 artifact boundary に従う。

- 要求・不足・要望: `docs/requirements/`
- 判断: `docs/adr/`
- 横断進捗: `docs/work-items/`
- 具体作業: `docs/tasks/`
- internal wiring route: `docs/internal-design/`
- canonical reference / tool-independent trace contract: `docs/spec/concepts/traceability/`
- MCP tool contract: `docs/spec/design-records-mcp/`

> 由来: V01-ADR-082, V01-ADR-083, V01-ADR-084, V01-ADR-087, V01-ADR-088, V01-ADR-091, V01-ADR-092; V01-INV-DOCS-002; V01-INV-DOCS-003
