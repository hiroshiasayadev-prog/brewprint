---
scope: docs/spec/concepts/traceability/metadata-schema.md
status: draft
last_updated: 2026-05-18
summary: >
  trace metadata YAML と front matter の最小 schema、semantic_refs、sections、
  artifact metadata の責務境界を定義する。
depends_on:
  - docs/adr/081-requirement-artifacts-and-spec-traceability.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
semantic_refs:
  - spec:trace.metadata-schema
sections:
  spec:trace.metadata-schema.trace-metadata-yaml: Trace metadata YAML
  spec:trace.metadata-schema.semantic-refs: semantic_refs
  spec:trace.metadata-schema.sections: sections
---

# Trace metadata schema

## Trace metadata YAML

Trace metadata YAML は、semantic trace を運用するために docs artifact の front matter または YAML file として書かれる metadata である。

Trace metadata YAML は brewprint DSL YAML ではない。
`yaml:` semantic ref prefix の対象でもない。

Examples:

- spec front matter
- internal-design front matter
- coverage mapping YAML
- requirement metadata
- work item metadata

Trace metadata YAML の schema validation は、coverage relation vocabulary ではなく、traceability spec / MCP tool contract の責務である。

## Common fields

MVP の trace metadata は、必要に応じて以下の field を持てる。

| field | required | type | meaning |
|---|---:|---|---|
| `semantic_refs` | no | list<string> | document-level semantic ref list |
| `sections` | no | map<string,string> | section-level semantic ref to Markdown heading text mapping |

この file では、doc-policy 既存の `scope` / `status` / `last_updated` / `summary` / `depends_on` は再定義しない。
既存 spec front matter の基本 field は doc-policy に従う。

## `semantic_refs`

`semantic_refs` は、その document / artifact 全体が所有する document-level semantic ref の list である。

Example:

```yaml
semantic_refs:
  - spec:trace.semantic-ref
```

Rules:

- values は semantic ref grammar に従う
- duplicate value は invalid とする
- document-level ref は同一 project 内で一意でなければならない
- physical path を value にしてはならない

`semantic_refs` は section-level ref を置く場所ではない。
Section-level ref は `sections` に置く。

## `sections`

`sections` は section-level semantic ref と Markdown heading text の対応を表す map である。

Example:

```yaml
sections:
  spec:trace.semantic-ref: Semantic ref
  spec:trace.coverage-mapping: Coverage mapping
```

Map key は section-level semantic ref である。
Map value は該当 section の Markdown heading text である。

Rules:

- key は semantic ref grammar に従う
- key は同一 project 内で一意でなければならない
- value は heading marker `#` を含めない
- value は heading text と一致しなければならない
- heading rename では key を維持し、value を更新する
- section move では key を維持する

Markdown heading に `{#anchor}` を直接書く方式は使わない。

## Spec front matter

Spec file は、doc-policy の front matter に加えて `semantic_refs` / `sections` を持てる。

Example:

```yaml
---
scope: docs/spec/concepts/traceability/semantic-ref.md
status: draft
last_updated: 2026-05-18
summary: >
  semantic ref の grammar と安定性を定義する。
depends_on:
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
semantic_refs:
  - spec:trace.semantic-ref
sections:
  spec:trace.semantic-ref: Semantic ref
  spec:trace.semantic-ref-grammar: Semantic ref grammar
---
```

`semantic_refs` / `sections` は Design Records MCP の `design_record` metadata とは別である。
Design Records MCP record index が読む metadata と、traceability resolver が読む semantic refs は別責務として扱う。

## Internal-design front matter

Internal design file は、document-level semantic ref を `semantic_refs` に置ける。

Example:

```yaml
---
scope: docs/internal-design/resolver/semantic-ref-index.md
status: draft
last_updated: 2026-05-18
summary: >
  semantic ref index resolver の内部設計を定義する。
depends_on:
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
semantic_refs:
  - internal-design:resolver.semantic-ref-index
---
```

MVP では `internal-design:` ref の section-level 解決単位を固定しない。
必要になった場合は `sections` mapping を使える。

## Coverage metadata

Coverage artifact は、mapping set / group の semantic ref を `semantic_refs` に置ける。

Example:

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
```

Coverage mapping の具体 schema は [`coverage-mapping.md`](coverage-mapping.md) が定義する。

## Requirement / work item metadata

Requirement / work item は MVP では prefix-ref ではなく ID-as-ref を持つ。

Requirement metadata example:

```yaml
id: REQ-TRACE-001
title: semantic trace MVP schema を定義する
status: accepted
source_refs:
  specs:
    - spec:trace.semantic-ref
work_items:
  - WORK-TRACE-001
```

Work item metadata example:

```yaml
id: WORK-TRACE-001
source_requirement: REQ-TRACE-001
impact_refs:
  specs:
    - spec:trace.semantic-ref
  internal_design:
    - internal-design:resolver.semantic-ref-index
```

MVP では requirement / work item の完全 schema は定義しない。
この file は、metadata 内で prefix-ref と ID-as-ref が併存しうることだけを示す。

## Validation responsibility

Trace metadata YAML schema validation は、relation vocabulary の `validates` ではない。

Validation examples:

- `semantic_refs` の値が grammar に従っているか
- `sections` の key が grammar に従っているか
- section mapping の value が実在 heading と一致するか
- duplicate semantic ref がないか
- coverage mapping の source / target が active prefix を持つか

これらは resolver / validator / MCP tool contract が扱う。

## Out of scope

この file では以下を定義しない。

- brewprint DSL YAML schema
- brewprint DSL YAML entity-level semantic ref
- requirement / work item の完全 lifecycle schema
- MCP writer tool request / response
- fixture-local coverage schema
