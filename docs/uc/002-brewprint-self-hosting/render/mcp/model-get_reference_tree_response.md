# get_reference_tree_response

MCP tool `get_reference_tree` のresponse model。
docs/spec/mcp/tools/get-reference-tree.md §3 のOutput、§4 Node entry、§5 Edge entry、§6 Traversal semanticsに対応する。
tool名はtreeだが返却形式は nodes[] / edges[] からなるbounded reference graph。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/get_reference_tree_response.yaml |

### Fields

| field | type | note |
|---|---|---|
| root | object_ref | 必須。traversal root ObjectRef。 |
| direction | reference_tree_direction | 必須。実際に使ったdirection。out / in / both。 |
| depth | int | 必須。実際に使ったmax traversal depth。 |
| nodes | list<get_reference_tree_node> | 必須。到達object entry配列。各entryは object:ObjectRef / depth:int / via:reference_kind[] を持つ。 |
| edges | list<get_reference_tree_edge> | 必須。traversalで辿ったReference entry配列。Referenceにdepth:intを加えたshape。継承 / extensionは使わずsame-file helperで表す。 |
| truncated | bool | 必須。max_nodes / max_edges により打ち切ったか。 |
| truncated_reasons | string_list | 必須。max_nodes / max_edges 等の打ち切り理由文字列配列。 |
| diagnostics | diagnostic_list | 必須。Diagnostic list。 |

## Private models

File-private helper schemas defined in this model YAML file.
Promote a helper model to a public model file when it needs to be reused from other YAML files.

| model | kind | shape | note |
|---|---|---|---|
| get_reference_tree_edge | struct | kind: reference_kind<br/>direction: reference_direction<br/>from: object_ref<br/>to: object_ref<br/>source: source_location<br/>doc: str<br/>depth: int | `get_reference_tree_response.edges[]` のresponse-local helper model。<br/>v1 modelには継承 / extensionがないため、`reference`にdepthを足したshapeを明示的に再掲する。 |
| get_reference_tree_node | struct | object: object_ref<br/>depth: int<br/>via: list<reference_kind> | `get_reference_tree_response.nodes[]` のresponse-local helper model。 |

