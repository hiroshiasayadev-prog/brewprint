# source_location

semantic objectに対応するsource位置。line/column欠落を許容する。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/yaml/mcp/model/source_location.yaml |

### Fields

| field | type | note |
|---|---|---|
| file | str | FileID。必須 |
| line | int | 1-origin line number。取得できる場合のみ |
| column | int | 1-origin column number。取得できる場合のみ |
| end_line | int | 範囲終端line。任意 |
| end_column | int | 範囲終端column。任意 |

