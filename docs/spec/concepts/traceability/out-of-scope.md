---
scope: docs/spec/concepts/traceability/out-of-scope.md
status: draft
last_updated: 2026-05-18
summary: >
  traceability MVP の scope 外、future extension 候補、後続 requirement / work item 化の基準を定義する。
depends_on:
  - docs/adr/081-requirement-artifacts-and-spec-traceability.md
  - docs/adr/082-golden-fixture-and-self-hosting-requirement-boundary.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
semantic_refs:
  - spec:trace.out-of-scope
sections:
  spec:trace.out-of-scope: Out of scope
  spec:trace.future-extensions: Future extensions
---

# Traceability out of scope

## Out of scope

この file は、traceability MVP で扱わないものを明示する。

MVP scope 外の項目は、不要という意味ではない。
必要になった場合は、ADR-081 の requirements layer に捕捉し、work item によって spec / internal design / coverage / implementation へ展開する。

## Brewprint DSL YAML entity-level refs

MVP では `yaml:` prefix を active 化しない。

以下は scope 外である。

- brewprint DSL YAML file-level ref
- brewprint DSL YAML entity-level ref
- node / edge / view / model / task / asset 単位の semantic ref
- YAML 内 anchor
- YAML logical unit の resolver rule

`yaml:` active 化は、self-hosting / UC-002 再構築、または YAML entity ref / resolver rule の spec 化が進んだ時点で後続 ADR / requirement により判断する。

## Fixture / golden traceability

MVP では `fixture:` prefix を定義しない。
`fixture:` は reserved prefix でもない。

Fixture / golden は brewprint processor / renderer / validator の debug / regression test asset であり、project-level semantic trace graph には含めない。

Scope outside MVP:

- fixture semantic ref
- fixture-local coverage の project-level graph 統合
- golden output と spec semantic ref の対応
- render expected comparison semantics
- test harness schema
- golden update workflow

Fixture-level traceability が必要になった場合は、project-level trace graph とは別 layer として後続 requirement / work item で導入する。

## `validates` relation

MVP relation vocabulary に `validates` は含めない。

これは validation 機能を不要とする意味ではない。
MVP で外すのは、coverage edge としての `validates` relation である。

以下は後続 spec / tool contract の責務として扱う。

- semantic ref resolver
- duplicate detection
- orphan detection
- trace metadata YAML schema validation
- MCP writer artifact generation contract

## Requirement / work-item prefixes

MVP では `requirement:` / `work-item:` prefix を定義しない。
Requirement / work item は `REQ-*` / `WORK-*` の ID-as-ref として扱う。

Requirement / work item を coverage edge の endpoint として扱う必要が出た場合は、後続 ADR / requirement で prefix reserve または active 化を判断する。

## Full MCP writer tools

MVP では MCP writer tool の request / response schema を定義しない。

Future writer tool candidates:

- create requirement
- update requirement status
- create work item
- add coverage mapping
- register semantic ref
- update section mapping

Writer tool を導入する場合は、以下を別 spec / ADR で定義する。

- dry-run diff
- user confirmation
- conflict handling
- formatting preservation
- write permission boundary
- generated metadata ownership

## Relation ontology expansion

MVP relation vocabulary は `maps_to` / `covers` のみである。

以下の relation は MVP では定義しない。

- `validates`
- `depends_on`
- `implements`
- `derives_from`
- `supersedes`
- `redirects_to`
- `affects`

既存 relation が複数の意味で運用されていることが coverage review で確認された場合、relation vocabulary の追加または rename を検討する。

## Future extensions

後続 extension 候補は以下である。

| extension | trigger |
|---|---|
| `yaml:` active 化 | self-hosting / UC-002 再構築、または YAML entity ref / resolver rule の spec 化 |
| requirement / work-item prefix | requirement / work item を coverage edge endpoint にする必要が出た場合 |
| fixture-level traceability | golden fixture と docs/spec 対応を長期管理する必要が出た場合 |
| validation relation | test / validation artifact を trace graph の一級対象にする必要が出た場合 |
| MCP trace tools | semantic ref query / resolve / validate を MCP contract として公開する必要が出た場合 |
| MCP writer tools | trace metadata artifact を tool で生成・更新する必要が出た場合 |

## Follow-up artifact placement

Future extension が必要になった場合、置き場所は ADR-083 の artifact placement decision rule に従う。

- 要求・不足・要望: `docs/requirements/`
- 横断進捗・影響範囲: `docs/work-items/`
- 具体作業手順: `docs/tasks/`
- internal wiring route: `docs/internal-design/`
- project-level trace relation: `docs/coverage/`
- MCP tool contract: `docs/spec/design-records-mcp/` または将来の tool spec 配下

Traceability spec は、future extension の要求や作業状態を所有しない。
