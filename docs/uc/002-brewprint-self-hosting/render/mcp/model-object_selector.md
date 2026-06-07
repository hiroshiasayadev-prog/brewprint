# object_selector

MCP tool共通のselector。
JSON schema上は全field任意だが、toolごとに有効な組み合わせが制約される。
selector field combination / object-dependent kind vocabulary / tool support matrix の正本は docs/spec/mcp/schema.md §1.1〜§1.2 と各 tool spec。
このYAMLはUC-002上のmodel表現であり、DATA DSLにdependent enumやselector matrix validationを導入するものではない。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/common.yaml |

### Fields

| field | type | note |
|---|---|---|
| id | str | 対象object ID。QualifiedID / actor global ID / synthetic IDを許容。任意 |
| object | mcp_object_type | 任意。node / view / transition / asset / field / file / primitive。 |
| kind | str | 任意。期待kind。値集合はobject-dependentであり、正本は docs/spec/mcp/schema.md §1.1 object-dependent kind vocabulary。指定時は解決結果との一致を検証する。 |
| file | str | FileID。file-local object指定時に使う。任意 |
| local_id | str | file内local object ID。sub task / field / asset等で使う。任意 |

