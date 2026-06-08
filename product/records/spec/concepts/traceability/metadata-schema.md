---
scope: docs/spec/concepts/traceability/metadata-schema.md
status: draft
last_updated: 2026-06-02
summary: >
  canonical reference resolution foundation のための spec front matter と
  investigation reference metadata boundary を定義する。
depends_on:
  - docs/adr/081-requirement-artifacts-and-spec-traceability.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/adr/084-semantic-trace-mvp-scope-and-artifact-boundary.md
  - docs/adr/086-investigation-artifact-format-and-lifecycle.md
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
  - docs/adr/091-workflow-artifact-work-item-task-milestone.md
  - docs/adr/092-design-records-mcp-workflow-artifact-record-and-relation-boundary.md
semantic_refs:
  - spec:trace.metadata-schema
sections:
  spec:trace.metadata-schema.trace-metadata: Trace metadata
  spec:trace.metadata-schema.semantic-refs: `semantic_refs`
  spec:trace.metadata-schema.sections: `sections`
  spec:trace.metadata-schema.investigation-reference: Investigation reference metadata
---

# Trace metadata schema

## Trace metadata

Trace metadata は、canonical reference を宣言・参照するための docs metadata である。Brewprint DSL YAML ではなく、`yaml:` semantic ref の対象でもない。

MVP でこの spec が定義する metadata は次に限定する。

- spec front matter の `semantic_refs` / `sections`
- investigation metadata に記載される canonical reference の許容形と validation boundary
- workflow artifact metadata に記載される `REQ-*` / `WORK-*` / `TASK-*` declared relation の validation boundary

Internal-design relation metadata、coverage mapping YAML、relation entry schema は MVP では定義しない。

## Common fields

MVP の spec trace metadata は、必要に応じて以下の field を持てる。

| field | required | type | meaning |
|---|---:|---|---|
| `semantic_refs` | no | list<string> | document-level `spec:` semantic ref list |
| `sections` | no | map<string,string> | section-level `spec:` semantic ref と Markdown heading text の対応 |

`scope` / `status` / `last_updated` / `summary` / `depends_on` は doc-policy が所有する。

## `semantic_refs`

`semantic_refs` は、その spec document 全体が所有する document-level `spec:` ref の list である。

```yaml
semantic_refs:
  - spec:trace
```

Rules:

- value は active `spec:` semantic ref grammar に従う
- duplicate value は invalid とする
- document-level ref は同一 project 内で一意でなければならない
- physical path を value にしてはならない

## `sections`

`sections` は section-level `spec:` semantic ref と Markdown heading text の対応を表す map である。

```yaml
sections:
  spec:trace.semantic-ref.definition: Semantic ref definition
  spec:trace.resolve: Resolve
```

Rules:

- key は active `spec:` semantic ref grammar に従う
- key は同一 project 内で一意でなければならない
- value は heading marker `#` を含めず、実在 heading text と一致しなければならない
- heading rename / section move では semantic ref key を維持し、必要に応じて value のみ更新する

Markdown heading に `{#anchor}` を書いて canonical identity とする方式は用いない。

## Spec front matter

Spec file は doc-policy の front matter に加えて `semantic_refs` / `sections` を持てる。

```yaml
---
scope: docs/spec/concepts/traceability/semantic-ref.md
status: draft
last_updated: 2026-05-24
summary: >
  semantic ref の grammar と安定性を定義する。
depends_on:
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
semantic_refs:
  - spec:trace.semantic-ref
sections:
  spec:trace.semantic-ref.definition: Semantic ref definition
---
```

`semantic_refs` / `sections` は Design Records MCP の record metadata とは別責務である。

## Internal design metadata boundary

`docs/internal-design/` は implementation-facing artifact layer として存続するが、semantic trace MVP は internal design front matter に `internal-design:` ref や source `spec:` relation declaration を要求しない。

Internal design を canonical reference として解決すべき concrete consumer が生じた場合に、prefix、metadata field、resolver rule、validation contract を後続判断する。

## Coverage metadata boundary

MVP は external coverage artifact、`coverage:` ref、`COV-*` ID、coverage mapping YAML schema を定義しない。

Gap / completeness / evidence / sign-off / audit / approved relation set を外部 artifact で保持する必要が生じた場合、名称と schema を含めて後続判断する。

## Workflow artifact metadata boundary

Requirement / work item / task は、Design Records MCP が扱う workflow record artifact である。Workflow artifact の完全 parser / response schema と diagnostic category は Design Records MCP spec が所有し、本 spec は canonical relation boundary のみを定義する。

Workflow artifact 間 relation は次の ID-as-ref field により宣言する。

| source artifact | field | canonical target |
|---|---|---|
| requirement | `work_items` | `WORK-*` |
| work item | `source_requirement` | `REQ-*` |
| work item | `tasks` | `TASK-*` |
| task | `work_item` | `WORK-*` |
| task | `source_requirement` | `REQ-*` |
| task | `depends_on` | `TASK-*` |

Workflow relation は metadata の ID-as-ref のみから読み、physical path または ID 文字列構造から親 relation を導出しない。`req:` / `work:` / `task:` semantic prefix も導入しない。

MVP は上記 field の参照先存在確認と、declared relation の以下の整合性確認を扱う。

- requirement と work item の相互 relation: `requirement.work_items` と `work_item.source_requirement`
- work item と task の相互 relation: `work_item.tasks` と `task.work_item`
- task の source requirement と parent work item の source requirement の一致

未接続 artifact の orphan diagnostics、task status 由来 progress projection、workflow traversal query、task dependency cycle / execution order projection は扱わない。

## Investigation reference metadata

Investigation metadata の field 構成、required / optional 区分、status、lifecycle、authoring format は guide ID `investigation-authoring` が authoring guidance として扱う。

この spec が所有するのは、V01-ADR-087 / V01-ADR-088 に基づく canonical reference rule である。

- `source_refs` は record ID-as-ref または active `spec:` semantic ref を用い、記載値は resolve 可能でなければならない
- `follow_up_results` は記載する場合、record ID-as-ref または active `spec:` semantic ref を用い、記載値は resolve 可能でなければならない
- `follow_up_candidates` に artifact reference を記載する場合、canonical form を用いる。未作成 artifact 候補を指しうるため、unresolved は error にせず `info` diagnostic として可視化する
- workflow artifact ID-as-ref のうち、investigation metadata が追加で使用できるものは `REQ-*` / `WORK-*` に限定する
- `TASK-*` は direct resolver input と workflow artifact 間 relation では supported だが、investigation metadata canonical reference には含めない。`source_refs` / `follow_up_results` に現れた場合は unsupported error、`follow_up_candidates` に現れた場合は unsupported info として扱う
- physical path は canonical reference としない。`source_refs` / `follow_up_results` に現れた場合は error diagnostic、`follow_up_candidates` に現れた場合は noncanonical candidate を示す `info` diagnostic として扱う
- `trigger` / `related_*` の resolve / validation rule は後続 contract で定義する

## Validation responsibility

MVP validation は canonical reference validation と、workflow artifact が明示的に宣言した ID-as-ref relation の integrity validation である。

- `semantic_refs` / `sections` の grammar と uniqueness
- `sections` value と実在 heading の一致
- investigation `source_refs` / 記載済み `follow_up_results` の canonicality と resolve
- investigation `follow_up_candidates` に記載された artifact reference の canonical form と、unresolved candidate の `info` diagnostic
- investigation metadata での `REQ-*` / `WORK-*` 許容と `TASK-*` 非対応 boundary
- physical path の noncanonical diagnostic。`source_refs` / `follow_up_results` は error、`follow_up_candidates` は `info` とする
- workflow relation field の canonical target kind / resolution / declared bidirectional consistency

Coverage mapping endpoint、semantic realization relation、`COV-*`、`internal-design:` resolve、workflow orphan diagnostics、progress projection、traversal query は MVP validation 対象外である。

## Out of scope

- brewprint DSL YAML schema / entity-level semantic ref
- internal-design semantic ref / relation metadata schema
- coverage mapping schema
- workflow artifact の orphan diagnostics / progress projection / traversal query / dependency cycle detection
- MCP writer tool request / response
- fixture-level traceability
