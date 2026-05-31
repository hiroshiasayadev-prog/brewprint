# diagnostic

tool実行は成立したが注意すべき情報。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/diagnostic.yaml |

### Fields

| field | type | note |
|---|---|---|
| severity | mcp_diagnostic_severity | error / warning / info / hint |
| code | str | machine-readable diagnostic code |
| message | str | human-readable message |
| source | source_location | 関連source。任意 |
| related | any | 関連SourceLocationまたはObjectRef配列。v1 modelではunion listを表せないためany |

