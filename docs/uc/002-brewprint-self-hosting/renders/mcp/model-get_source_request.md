# get_source_request

MCP tool `get_source` のrequest model。
docs/spec/mcp/tools/get-source.md §2 のInputに対応する。
fallback=file または省略時はobject単位rangeが取れない場合に同一FileID全体を返す。
fallback=error の場合はobject単位rangeが特定できないとtool errorを返す。
fallbackのdefault、分岐、未知値errorはMCP tool specを正本とし、このYAMLでは要約だけを保持する。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/get_source_request.yaml |

### Fields

| field | type | note |
|---|---|---|
| selector | object_selector | 必須。source snippet取得対象のsemantic object selector。 |
| fallback | str | 任意。file / error。省略時は MCP tool contract により file相当。 |

