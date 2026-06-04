# inspect_response

MCP tool `inspect` のresponse model。
docs/spec/mcp/tools/inspect.md §3 Common output shape、および §4〜§14 のkind別inspect shapeに対応する。
`signature` と `members` は task / store / model / state / event / scenario / transition / field / api_table / er_diagram / file ごとに異なるunion相当のpayloadである。
brewprint v1 modelではunion / discriminated object / nested arbitrary objectを厳密表現できないため any + note で保持する。
`references` は主要referenceのみを返す任意fieldで、詳細なdirect referencesは `get_references` が担当する。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/yaml/mcp/model/inspect_response.yaml |

### Fields

| field | type | note |
|---|---|---|
| object | object_ref | 必須。inspect対象ObjectRef。 |
| signature | any | 必須。get_signature相当の外形。task / store / model / state / event / view / file 等でshapeが異なるためany。 |
| doc | str | 任意。note由来の説明。 |
| source | source_location | 任意。定義元SourceLocation。 |
| members | any | 任意。objectが内包する要素。kind別shapeが大きく異なるためanyで暫定表現する。 |
| references | reference_list | 任意。主要reference一覧。 |
| diagnostics | diagnostic_list | 必須。Diagnostic list。 |

