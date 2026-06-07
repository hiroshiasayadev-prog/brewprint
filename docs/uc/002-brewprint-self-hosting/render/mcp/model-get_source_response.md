# get_source_response

MCP tool `get_source` のresponse model。
docs/spec/mcp/tools/get-source.md §3 のOutputに対応する。
`snippet` は `{ language: "yaml", text: "..." }` を持つ。language literal制約はnoteで保持する。
Raw YAML ASTではなく、ResolvedProject上のsemantic objectに紐づくsource補助情報として扱う。
`fallback` marker と fallback時diagnosticはMCP tool specを正本とし、このYAMLでは要約だけを保持する。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/get_source_response.yaml |

### Fields

| field | type | note |
|---|---|---|
| object | object_ref | 必須。source取得対象のObjectRef。 |
| source | source_location | 必須。対応するSourceLocation。line / columnが取得できない場合はfileのみでもよい。 |
| snippet | get_source_snippet | 必須。language: yaml と text を持つobject。language literal制約はnoteで保持する。 |
| fallback | str | 任意。fallbackした場合はfile。fallback response contractはMCP tool specを正本とする。 |
| diagnostics | diagnostic_list | 必須。fallback時のsource_range_unavailable等を含む。 |

## Private models

File-private helper schemas defined in this model YAML file.
Promote a helper model to a public model file when it needs to be reused from other YAML files.

| model | kind | shape | note |
|---|---|---|---|
| get_source_snippet | struct | language: str<br/>text: str | `get_source_response.snippet` のresponse-local helper model。 |

