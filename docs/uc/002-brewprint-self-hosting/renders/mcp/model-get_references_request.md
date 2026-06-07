# get_references_request

MCP tool `get_references` のrequest model。
docs/spec/mcp/tools/get-references.md §2 のInputに対応する。
MCP v1ではdirect referencesのみを返し、inputにdepthは持たない。
`kinds` は Reference.kind のfilterであり、許容値は docs/spec/mcp/schema.md §2.2 のReference kindに従う。
directionの省略時defaultと未知値errorはMCP tool specを正本とし、このYAMLでは要約だけを保持する。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/get_references_request.yaml |

### Fields

| field | type | note |
|---|---|---|
| selector | object_selector | 必須。直接reference取得対象のObject selector。 |
| direction | str | 任意。out / in / both。省略時は MCP tool contract により out。 |
| kinds | string_list | 任意。reference kind filterの文字列配列。 |

