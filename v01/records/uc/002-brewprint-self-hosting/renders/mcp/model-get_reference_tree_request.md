# get_reference_tree_request

MCP tool `get_reference_tree` のrequest model。
docs/spec/mcp/tools/get-reference-tree.md §2 のInputに対応する。
directionは探索範囲を暗黙化しないため必須。
`kinds` は docs/spec/mcp/schema.md §2.2 のReference.kind filter。
directionのclosed vocabularyはenum modelで表現する。
depthの範囲制約と範囲外errorはMCP tool specを正本とし、このYAMLでは要約だけを保持する。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/get_reference_tree_request.yaml |

### Fields

| field | type | note |
|---|---|---|
| selector | object_selector | 必須。traversal root object selector。 |
| direction | reference_tree_direction | 必須。out / in / both。 |
| depth | int | 必須。0..4。範囲外は MCP tool contract により invalid_depth error。 |
| kinds | string_list | 任意。traversal / return対象のreference kind filter文字列配列。 |
| max_nodes | int | 任意。node返却上限。省略時200。 |
| max_edges | int | 任意。edge返却上限。省略時500。 |

