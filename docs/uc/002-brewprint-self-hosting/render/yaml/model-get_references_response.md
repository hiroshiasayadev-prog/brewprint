# get_references_response

MCP tool `get_references` のresponse model。
docs/spec/mcp/tools/get-references.md §3 のOutput、および §4 のdepth固定ルールに対応する。
`references` は共通model `reference_list` を参照する。
MCP v1ではdirect referencesのみを返し、transitive traversalは `get_reference_tree` が担当する。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/yaml/mcp/model/get_references_response.yaml |

### Fields

| field | type | note |
|---|---|---|
| object | object_ref | 必須。query対象object。 |
| direction | reference_query_direction | 必須。実際に使ったdirection。 |
| depth | int | 必須。get_referencesでは常に1を返す。 |
| references | reference_list | 必須。対象objectの直接reference一覧。 |
| diagnostics | diagnostic_list | 必須。Diagnostic list。 |

