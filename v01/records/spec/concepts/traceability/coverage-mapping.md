---
scope: docs/spec/concepts/traceability/coverage-mapping.md
status: draft
last_updated: 2026-05-24
summary: >
  semantic realization mapping と external coverage artifact を MVP 外へ送る境界、
  および将来再導入を判断する trigger を定義する。
depends_on:
  - docs/adr/081-requirement-artifacts-and-spec-traceability.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
semantic_refs:
  - spec:trace.coverage-mapping
sections:
  spec:trace.coverage-mapping.boundary: MVP realization mapping boundary
  spec:trace.coverage-mapping.deferred-mechanisms: Deferred mechanisms
  spec:trace.coverage-mapping.reintroduction-triggers: Reintroduction triggers
---

# Semantic realization mapping boundary

## MVP realization mapping boundary

Semantic trace MVP は、semantic realization mapping を operational contract として扱わない。

MVP は `spec:` semantic ref と record artifact ID-as-ref の canonical resolution、および investigation reference validation に限定する。Artifact 間の implementation realization graph、coverage matrix、evidence matrix は構築しない。

そのため、MVP では以下を定義・要求しない。

- semantic realization relation を保持する external mapping artifact とその配置
- `coverage:` semantic ref prefix
- `COV-*` individual mapping ID
- coverage mapping YAML schema
- `maps_to` / `covers` relation の authoring / resolve / validation
- `spec:` → `internal-design:` mapping direction や endpoint constraint

## Internal design boundary

`docs/internal-design/` は implementation-facing design artifact layer として存続するが、`internal-design:` は MVP active endpoint ではない。

Internal design document が存在することは、MVP に semantic realization relation を要求する理由にはならない。Internal design metadata に source `spec:` relation を宣言させる schema や、spec から internal design への逆引き graph も MVP では定義しない。

## Deferred mechanisms

以下は必要性を否定したものではなく、concrete requirement が確認されるまで判断を延期した mechanism である。

| mechanism | MVP treatment |
|---|---|
| `internal-design:` semantic endpoint | deferred |
| `coverage:` semantic endpoint | deferred |
| `COV-*` mapping identity | deferred |
| external relation / assurance artifact とその配置 | deferred |
| `maps_to` semantic realization relation | deferred |
| `covers` semantic realization relation | deferred |
| `yaml:` semantic endpoint | reserved / inactive |

`maps_to` を external artifact から endpoint metadata へ移設することも MVP では行わない。Relation 自体を operational に導入する時点で、identity、direction、owner、schema、validation を一緒に判断する。

## Reintroduction triggers

以下のいずれかが requirement として確認された場合、後続 ADR / requirement / work item により、endpoint・relation・external artifact の必要性を再判断する。

- 複数の internal design document について、spec からの機械的 navigation / impact analysis が必要になった場合
- investigation / work item / MCP query が internal design artifact を canonical reference として解決・検証する必要を持った場合
- `yaml:` active 化により、spec / internal design / YAML の realization chain または cross-layer validation が必要になった場合
- gap / completeness / approved relation set を中央管理する必要が生じた場合
- evidence / sign-off / audit snapshot / release baseline を relation と結び付ける必要が生じた場合
- relation entry 自体に stable identity / approval / lifecycle / history が必要になった場合

## Validation boundary

MVP で validation 対象となるのは、canonical reference foundation に関する rule である。

- `spec:` semantic ref の grammar / uniqueness / resolve
- record ID-as-ref の resolve
- investigation の `source_refs` / 記載済み `follow_up_results` の canonicality と resolve
- investigation の `follow_up_candidates` に artifact reference を記載する場合の canonical form
- physical path が canonical reference ではないこと

Coverage mapping schema や realization endpoint の validation は MVP 対象外である。

## Future artifact placement

External relation artifact を将来導入する場合、その名称を `coverage` とするか、semantic mapping と assurance matrix を分離するか、どの directory に配置するかを再判断する。MVP は external artifact 用 directory を予約しない。

> 由来: V01-ADR-083, V01-ADR-084, V01-ADR-088; V01-INV-DOCS-002; V01-INV-DOCS-003
