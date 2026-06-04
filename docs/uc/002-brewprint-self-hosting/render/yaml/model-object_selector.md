# object_selector

MCP tool共通のselector。
JSON schema上は全field任意だが、toolごとに有効な組み合わせが制約される。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/yaml/mcp/model/common.yaml |

### Fields

| field | type | note |
|---|---|---|
| id | str | 対象object ID。QualifiedID / actor global ID / synthetic IDを許容。任意 |
| object | mcp_object_type | 任意。node / view / transition / asset / field / file / primitive。 |
| kind | str | 期待kind。指定時は解決結果との一致を検証する。任意 |
| file | str | FileID。file-local object指定時に使う。任意 |
| local_id | str | file内local object ID。sub task / field / asset等で使う。任意 |

