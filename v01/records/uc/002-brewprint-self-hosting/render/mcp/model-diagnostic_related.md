# diagnostic_related

diagnostic.related の要素を表す tagged union envelope。
以前の untagged SourceLocation | ObjectRef list を machine-readable な envelope model に置き換える。

## Public model

| property | value |
|---|---|
| kind | tagged_union |
| visibility | public |
| source | yaml/mcp/model/diagnostic_related.yaml |

### Discriminator

| property | value |
|---|---|
| discriminator | kind |

### Variants

#### `source_location`

| field | type | note |
|---|---|---|
| location | source_location | 関連source位置。 |

#### `object_ref`

| field | type | note |
|---|---|---|
| object | object_ref | 関連semantic object。 |

