# analyze_impact_response

MCP tool `analyze_impact` のresponse model。
docs/spec/mcp/tools/analyze-impact.md §4 Output、§5 Impact entry、§6 severity / fixability、§7 coverage、§8 suggested_fixes、§9 SourceLocation、§10 coverage scope details、§12 assumptionsに対応する。
`impact.severity` は breaking / warning / info。
`impact.fixability` は mechanical / suggested / manual_review / unknown。
`coverage.analyzed` / `coverage.not_analyzed` は `analyze_impact_coverage_vocabulary` enum の typed list で表現する。
`suggested_fixes[]` はfix kind依存payloadを持つためany内に含める。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/yaml/mcp/model/analyze_impact_response.yaml |

### Fields

| field | type | note |
|---|---|---|
| target | object_ref | 必須。分析対象ObjectRef。 |
| change | analyze_impact_change | 必須。inputで指定されたchangeをそのまま返す。 |
| summary | analyze_impact_summary | 必須。by_severity / by_fixability / by_kind の件数集計object。 |
| impacts | list<analyze_impact_impact> | 必須。impact entry配列。各entryは id / kind / severity / fixability / object / reason / via / source / recommended_action / suggested_fixes を持つ。 |
| coverage | analyze_impact_coverage | 必須。analyzed / not_analyzed / note を持つcoverage object。 |
| assumptions | string_list | 必須。tool側の前提・限界の文字列配列。 |
| truncated | bool | 必須。max_impactsにより打ち切ったか。 |
| truncated_reasons | string_list | 必須。打ち切り理由の文字列配列。 |
| diagnostics | diagnostic_list | 必須。unsupported selector時のunsupported_selector等を含む。 |

## Private models

File-private helper schemas defined in this model YAML file.
Promote a helper model to a public model file when it needs to be reused from other YAML files.

| model | kind | shape | note |
|---|---|---|---|
| analyze_impact_coverage | struct | analyzed: list<analyze_impact_coverage_vocabulary><br/>not_analyzed: list<analyze_impact_coverage_vocabulary><br/>note: str | `analyze_impact_response.coverage` のresponse-local helper model。 |
| analyze_impact_coverage_vocabulary | enum | direct_references<br/>reference_tree<br/>model_field_resolution<br/>transition_action_resolution<br/>flow_step_task_resolution<br/>flow_param_field_resolution<br/>sequence_step_task_resolution<br/>type_signature_identity<br/>render_output_files<br/>name_collision<br/>type_structural_compatibility<br/>semantic_contract_compatibility<br/>render_presentation_details<br/>wireframe_element_binding | analyze_impact coverage vocabulary。v1 標準語彙。docs/spec/mcp/tools/analyze-impact.md §7 に対応。 |
| analyze_impact_fixability_counts | struct | mechanical: int<br/>suggested: int<br/>manual_review: int<br/>unknown: int | `analyze_impact_summary.by_fixability` のhelper model。 |
| analyze_impact_impact | struct | id: str<br/>kind: str<br/>severity: str<br/>fixability: str<br/>object: object_ref<br/>reason: str<br/>via: string_list<br/>source: source_location<br/>recommended_action: str<br/>suggested_fixes: any | `analyze_impact_response.impacts[]` のresponse-local helper model。<br/>severity / fixability の値集合と suggested_fixes のkind依存payloadは、このmigrationではnoteに保持する。 |
| analyze_impact_severity_counts | struct | breaking: int<br/>warning: int<br/>info: int | `analyze_impact_summary.by_severity` のhelper model。 |
| analyze_impact_summary | struct | by_severity: analyze_impact_severity_counts<br/>by_fixability: analyze_impact_fixability_counts<br/>by_kind: any | `analyze_impact_response.summary` のhelper model。<br/>`by_kind` は impact kind 語彙が実装裁量であり、v1 model では dict key semantics を型として表現できないため `any` を維持する。 |

