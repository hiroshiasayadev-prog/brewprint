# object_ref

MCP response内でsemantic objectを指す共通形式。
parentは本来ObjectRefだが、v1 modelでは再帰structを直接表せないためanyで暫定表現する。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/yaml/mcp/model/object_ref.yaml |

### Fields

| field | type | note |
|---|---|---|
| object | mcp_object_type | node / view / transition / asset / field / file / primitive |
| kind | str | object種別。nodeならtask/model/store等。値集合はobjectごとに異なるためnoteで扱う |
| id | str | object ID。QualifiedIDまたはsynthetic ID |
| qualified_id | str | QualifiedIDを持つobjectのみ。任意 |
| file | str | file-local objectの所属FileID。任意 |
| local_id | str | file-local objectのlocal ID。任意 |
| label | str | 人間向け短縮表示名。任意 |
| source | source_location | 定義元source。任意 |
| parent | any | field等の親ObjectRef。再帰型をv1 modelで表しきれないためany |

