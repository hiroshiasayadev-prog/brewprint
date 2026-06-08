# list_objects_response

MCP tool `list_objects` のresponse model。
docs/spec/mcp/tools/list-objects.md §3 のOutputに対応する。
`objects[]` は object / kind / id / qualified_id / label / module / file / source 等を持つ。
ObjectRef identity semanticsの詳細化はV01-ADR-078以降のscope。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/list_objects_response.yaml |

### Fields

| field | type | note |
|---|---|---|
| objects | list<list_objects_object> | ObjectRef相当の一覧。module等のlist_objects固有summary fieldを含みうる。 |
| diagnostics | diagnostic_list | Diagnostic list。tool実行は成立したが注意すべき情報。 |

## Private models

File-private helper schemas defined in this model YAML file.
Promote a helper model to a public model file when it needs to be reused from other YAML files.

| model | kind | shape | note |
|---|---|---|---|
| list_objects_object | struct | object: mcp_object_type<br/>kind: str<br/>id: str<br/>qualified_id: str<br/>label: str<br/>module: str<br/>file: str<br/>source: source_location | `list_objects_response.objects[]` のresponse-local helper model。<br/>ObjectRef相当のidentity semantics整理はV01-ADR-078以降のscopeであり、このmigrationでは既存fieldをsame-file helperへ移す。 |

