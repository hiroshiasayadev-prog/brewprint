# string_list

MCP schema内の文字列配列を表す共通list model。
例: reference kind filter、truncated_reasons、coverage.analyzed / not_analyzed など。
配列要素のenum値集合は利用箇所ごとのnoteで補足する。

## Public model

| property | value |
|---|---|
| kind | list |
| visibility | public |
| source | yaml/mcp/model/string_list.yaml |

### Element

| property | value |
|---|---|
| element | str |

