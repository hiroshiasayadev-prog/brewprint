# analyze_impact_change

MCP tool `analyze_impact` のchange discriminated object。
discriminator field `kind` の値によってpayload shapeが変わる。
rename / remove / change_type / add はpayload shapeが明確なため構造化済み。
change_contract / change_transition_target はpayload制約が複雑なため後続migration対象。

## Public model

| property | value |
|---|---|
| kind | tagged_union |
| visibility | public |
| source | yaml/mcp/model/analyze_impact_change.yaml |

### Discriminator

| property | value |
|---|---|
| discriminator | kind |

### Variants

#### `rename`

| field | type | note |
|---|---|---|
| new_id | str | 必須。rename後のfull qualified ID。 |

#### `remove`

No payload fields.

#### `change_type`

| field | type | note |
|---|---|---|
| new_type | str | 必須。変更後の型名。 |

#### `add`

| field | type | note |
|---|---|---|
| added_id | str | 必須。追加後のfull qualified ID。 |

