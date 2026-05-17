---
scope: docs/spec/concepts/traceability/index.md
status: draft
last_updated: 2026-05-18
summary: >
  semantic traceability spec の入口。
  MVP scope、分割 spec 構成、用語、active / reserved / scope 外の概要を定義する。
depends_on:
  - docs/adr/081-requirement-artifacts-and-spec-traceability.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
---

# Traceability spec index

## 目的

この spec set は、brewprint docs における semantic ref / trace schema の MVP を定義する。

MVP の目的は、project-level docs artifact 間の trace を physical path ではなく semantic ref で扱えるようにすることである。

この spec set は、まず以下の対応関係を対象にする。

- design spec と internal design の対応
- coverage mapping set / group / individual mapping の最小構造
- semantic ref の解決と trace metadata validation の基本方針

## MVP scope

MVP で active semantic ref prefix として扱うのは以下に限定する。

```yaml
active_prefixes:
  - spec
  - internal-design
  - coverage
```

MVP では、ADR-083 が `docs/coverage/` に課した問いのうち、YAML を含まない問いだけを扱う。
具体的には、`spec ↔ internal-design` の対応を主対象とする。

`spec ↔ brewprint DSL YAML`、`internal-design ↔ brewprint DSL YAML`、および internal design 変更時にどの YAML を compile / validate / render すべきかという問いは、`yaml:` prefix の active 化まで扱わない。

## 用語

### semantic ref

physical path、Markdown heading、directory layout ではなく、artifact が表す概念を安定して参照するための identifier。

例:

```text
spec:trace.semantic-ref
internal-design:resolver.semantic-ref-index
coverage:trace.semantic-ref
```

### brewprint DSL YAML

対象 system / design model を brewprint DSL で表す YAML。
ADR-083 が primary DSL source と呼ぶ YAML はこれを指す。

例:

- DAG / asset / task などを表す YAML
- model / API / state machine / sequence などを表す YAML
- future self-hosting で brewprint 本体を表す YAML

### trace metadata YAML

semantic trace を運用するための metadata YAML。
これは `yaml:` prefix の対象ではない。

例:

- spec front matter
- internal-design front matter
- coverage mapping YAML
- requirement metadata
- work item metadata

trace metadata YAML の schema validation は、coverage relation vocabulary ではなく、traceability spec / MCP tool contract の責務として扱う。

## 分割 spec

この spec set は、人間がレビュー可能な粒度を保つため、以下に分割する。

| file | owns |
|---|---|
| `index.md` | scope、用語、分割 spec の入口 |
| `semantic-ref.md` | semantic ref grammar、安定性、document / section ref |
| `artifact-refs.md` | active / reserved prefix、ID-as-ref、artifact別 ref 方針 |
| `metadata-schema.md` | front matter / trace metadata YAML の最小 schema |
| `coverage-mapping.md` | coverage mapping set / group / mapping、relation vocabulary |
| `resolve-and-validation.md` | ref resolve、duplicate / orphan detection、schema validation 方針 |
| `out-of-scope.md` | MVP scope 外、future extension 候補 |

## Source of truth boundary

この spec set は Design Records MCP の tool contract ではない。

Traceability spec は docs artifact の semantic trace model を定義する上位仕様である。
Design Records MCP は、この仕様に基づく index / query / validation / writer tool を実装しうるが、traceability そのものは Design Records MCP の内部仕様ではない。

## MVP outside summary

MVP では以下を扱わない。

- `yaml:` active 化
- brewprint DSL YAML entity-level semantic ref
- `fixture:` prefix
- fixture / golden traceability
- `validates` relation
- render expected comparison semantics
- full MCP writer tool schema

詳細は [`out-of-scope.md`](out-of-scope.md) を参照する。

## 由来

- ADR-081: requirements layer と semantic traceability
- ADR-083: project artifact boundary と YAML as primary implementation source
- ADR-084: semantic trace MVP scope と artifact boundary
