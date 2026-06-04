# reference

ResolvedProject上のsemantic object間の直接参照。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/yaml/mcp/model/reference.yaml |

### Fields

| field | type | note |
|---|---|---|
| kind | reference_kind | semantic reference kind。 |
| direction | reference_direction | query対象から見た方向。out / in。 |
| from | object_ref | 参照元object |
| to | object_ref | 参照先object |
| source | source_location | このreferenceの定義元。任意 |
| doc | str | reference補足。任意 |

