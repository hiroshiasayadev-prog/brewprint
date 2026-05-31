# analyze_impact_response

MCP tool `analyze_impact` のresponse model。
docs/spec/mcp/tools/analyze-impact.md §4 Output、§5 Impact entry、§6 severity / fixability、§7 coverage、§8 suggested_fixes、§9 SourceLocation、§10 coverage scope details、§12 assumptionsに対応する。
`impact.severity` は breaking / warning / info。
`impact.fixability` は mechanical / suggested / manual_review / unknown。
`coverage.analyzed` / `coverage.not_analyzed` はv1標準語彙を持つが、brewprint v1 modelではenum listを厳密表現できないためany + noteで保持する。
`suggested_fixes[]` はfix kind依存payloadを持つためany内に含める。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/analyze_impact_response.yaml |

### Fields

| field | type | note |
|---|---|---|
| target | object_ref | 必須。分析対象ObjectRef。 |
| change | any | 必須。inputで指定されたchangeをそのまま返す。discriminated objectをv1 modelで表せないためany。 |
| summary | any | 必須。by_severity / by_fixability / by_kind の件数集計object。dict shapeを厳密表現しないためany。 |
| impacts | list<analyze_impact_impact> | 必須。impact entry配列。各entryは id / kind / severity / fixability / object / reason / via / source / recommended_action / suggested_fixes を持つ。 |
| coverage | analyze_impact_coverage | 必須。analyzed / not_analyzed / note を持つcoverage object。語彙制約はnoteで保持する。 |
| assumptions | any | 必須。tool側の前提・限界の文字列配列。専用list modelを作らずanyで暫定表現する。 |
| truncated | bool | 必須。max_impactsにより打ち切ったか。 |
| truncated_reasons | any | 必須。打ち切り理由の文字列配列。専用list modelを作らずanyで暫定表現する。 |
| diagnostics | diagnostic_list | 必須。unsupported selector時のunsupported_selector等を含む。 |

## Private models

File-private helper schemas defined in this model YAML file.
Promote a helper model to a public model file when it needs to be reused from other YAML files.

| model | kind | shape | note |
|---|---|---|---|
| analyze_impact_coverage | struct | analyzed: string_list<br/>not_analyzed: string_list<br/>note: str | `analyze_impact_response.coverage` のresponse-local helper model。<br/>coverage vocabularyのenum化はこのmigrationでは行わない。 |
| analyze_impact_impact | struct | id: str<br/>kind: str<br/>severity: str<br/>fixability: str<br/>object: object_ref<br/>reason: str<br/>via: string_list<br/>source: source_location<br/>recommended_action: str<br/>suggested_fixes: any | `analyze_impact_response.impacts[]` のresponse-local helper model。<br/>severity / fixability の値集合と suggested_fixes のkind依存payloadは、このmigrationではnoteに保持する。 |

