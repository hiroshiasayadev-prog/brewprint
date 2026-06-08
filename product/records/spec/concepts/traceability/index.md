---
scope: docs/spec/concepts/traceability/index.md
status: draft
last_updated: 2026-05-27
summary: >
  semantic traceability spec の入口。
  canonical reference resolution foundation と future realization scope の境界を定義する。
depends_on:
  - docs/adr/081-requirement-artifacts-and-spec-traceability.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
  - docs/adr/091-workflow-artifact-work-item-task-milestone.md
  - docs/adr/092-design-records-mcp-workflow-artifact-record-and-relation-boundary.md
semantic_refs:
  - spec:trace
---

# Traceability spec index

## 目的

この spec set は、brewprint docs の semantic trace MVP を定義する。

V01-ADR-088 / V01-ADR-092 により、MVP の目的は artifact 間の realization graph を先取りすることではなく、設計 record、investigation、および workflow artifact が physical path に依存せず canonical reference と宣言済み relation を解決・検証できる **canonical reference resolution foundation** を提供することになった。

## MVP scope

MVP が扱うもの:

- `spec:` semantic ref の宣言、安定性、document / section 解決
- `ADR-*` / `SPEC-*` / `INV-*` / `REQ-*` / `WORK-*` / `TASK-*` record ID-as-ref の解決
- investigation の `source_refs` および記載済み `follow_up_results` の canonical reference 解決と unresolved error。追加 workflow ID-as-ref は `REQ-*` / `WORK-*` に限定する
- `follow_up_candidates` に artifact reference を記載する場合の canonical form 検査。追加 workflow ID-as-ref は `REQ-*` / `WORK-*` に限定し、未作成候補の存在は要求しない
- workflow artifact 間の宣言済み ID-as-ref relation の存在確認と双方向整合性確認
- physical path を canonical reference としない boundary

MVP の active semantic ref prefix は以下に限定する。

```yaml
active_prefixes:
  - spec
```

`yaml:` は brewprint DSL YAML 用に reserve するが active 化しない。

## MVP で扱わないもの

MVP は以下を operational mechanism として扱わない。

- `internal-design:` semantic endpoint
- `coverage:` semantic endpoint
- `COV-*` mapping identity
- external relation / assurance artifact とその配置
- `maps_to` / `covers` / `validates` relation
- `spec:` → `internal-design:` realization mapping
- YAML endpoint、fixture / golden traceability、coverage / evidence matrix
- workflow artifact の orphan diagnostics、task status 由来 progress projection、workflow 専用 traversal query、task dependency cycle / execution order projection
- investigation metadata における `TASK-*` canonical reference support

`docs/internal-design/` artifact layer 自体は存続する。MVP 外なのは、同 layer を semantic trace endpoint として resolve / validate する contract である。External relation / assurance artifact は必要性が成立した時点で配置と責務を含めて新設判断し、MVP layout には directory を予約しない。

## 用語

### semantic ref

Physical path、Markdown heading、directory layout ではなく、artifact が表す概念を安定して参照する identifier。

MVP example:

```text
spec:trace
spec:trace.semantic-ref
spec:trace.resolve-and-validation
```

### record ID-as-ref

Design Records MCP が扱う record artifact を指す stable ID。

```text
V01-ADR-088
SPEC-<slug>
V01-INV-DOCS-003
V01-REQ-MCP-003
V01-WORK-MCP-003
V01-TASK-MCP-003-01
```

### brewprint DSL YAML

対象 system / design model を brewprint DSL で表す primary implementation source。`yaml:` semantic endpoint の active 化は future decision とする。

### trace metadata

Canonical reference を宣言・参照するための metadata。brewprint DSL YAML とは別責務である。MVP では spec front matter および investigation metadata の canonical reference rule を中心に扱う。

## 分割 spec

| file | owns |
|---|---|
| `index.md` | scope と spec set の入口 |
| `semantic-ref.md` | `spec:` grammar、安定性、document / section ref |
| `artifact-refs.md` | active / reserved / deferred ref と ID-as-ref 方針 |
| `metadata-schema.md` | MVP trace metadata と investigation reference boundary |
| `coverage-mapping.md` | realization mapping を MVP 外へ送る境界と再導入 trigger |
| `resolve-and-validation.md` | canonical resolve / validation 方針 |
| `out-of-scope.md` | future extension 候補 |

## Source of truth boundary

Traceability spec は docs artifact の canonical reference model を所有する。Design Records MCP は V01-ADR-087 に従い、その resolve / validation を実装する tool boundary であるが、traceability の意味モデル自体の owner ではない。

Artifact system 全体の責務境界は [`../project-artifact-model/index.md`](../project-artifact-model/index.md) が所有する。

## 由来

- V01-ADR-081: requirements layer と semantic traceability
- V01-ADR-083: project artifact boundary と YAML as primary implementation source
- V01-ADR-084: semantic trace MVP scope と artifact boundary
- V01-ADR-087: Design Records MCP investigation support and semantic ref resolve
- V01-ADR-088: Reduce semantic trace MVP to a canonical reference resolution foundation
- V01-ADR-091: Workflow artifact の work item / task 責務分離と legacy milestone 移行
- V01-ADR-092: Design Records MCP workflow artifact record and relation boundary
