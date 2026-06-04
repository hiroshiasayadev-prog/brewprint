# list_endpoints_response

MCP tool `list_endpoints` のresponse model。
docs/spec/mcp/tools/list-endpoints.md §3 Output、および §4 endpoint objectに対応する。
`tables[]` は API Table view単位の集約結果で、sections[] / endpoints[] を含むnested list objectである。
endpoint objectは method / path / leaf_path / task を必須とし、params / returns / source を任意で持つ。
route合成はADR-028の規則に従い、task(endpoint=true)の単純列挙ではない。
endpoint objectのoptional fieldはnoteで保持する。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/yaml/mcp/model/list_endpoints_response.yaml |

### Fields

| field | type | note |
|---|---|---|
| tables | list<list_endpoints_table> | 必須。API Tableごとのendpoint一覧。各tableは id / http_root_path / sections を持つ。 |
| diagnostics | diagnostic_list | 必須。Diagnostic list。 |

## Private models

File-private helper schemas defined in this model YAML file.
Promote a helper model to a public model file when it needs to be reused from other YAML files.

| model | kind | shape | note |
|---|---|---|---|
| list_endpoints_endpoint | struct | method: str<br/>path: str<br/>leaf_path: str<br/>task: str<br/>params: str<br/>returns: str<br/>source: source_location | `list_endpoints_section.endpoints[]` のresponse-local helper model。 |
| list_endpoints_section | struct | module: str<br/>include_submodules: bool<br/>endpoints: list<list_endpoints_endpoint> | `list_endpoints_table.sections[]` のresponse-local helper model。 |
| list_endpoints_table | struct | id: str<br/>http_root_path: str<br/>sections: list<list_endpoints_section> | `list_endpoints_response.tables[]` のresponse-local helper model。 |

