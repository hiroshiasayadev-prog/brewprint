# object_ref

MCP response内でsemantic objectを指す共通形式。
`object` と `kind` の組み合わせは docs/spec/mcp/schema.md §1.1 の object-dependent kind vocabulary を正本とする。
parentはfield等の親ObjectRefをrecursive named model referenceとして表す。
このYAMLはUC-002上のmodel表現であり、DATA DSLにdependent enumやinline recursive shapeを導入するものではない。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/object_ref.yaml |

### Fields

| field | type | note |
|---|---|---|
| object | mcp_object_type | node / view / transition / asset / field / file / primitive |
| kind | str | object-dependent kind。値集合の正本は docs/spec/mcp/schema.md §1.1 object-dependent kind vocabulary。nodeならtask/model/store等。 |
| id | str | object ID。QualifiedIDまたはsynthetic ID |
| qualified_id | str | QualifiedIDを持つobjectのみ。任意 |
| file | str | file-local objectの所属FileID。任意 |
| local_id | str | file-local objectのlocal ID。任意 |
| label | str | 人間向け短縮表示名。任意 |
| source | source_location | 定義元source。任意 |
| parent | object_ref | field等の親ObjectRef。recursive named model reference。 |

