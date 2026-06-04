# get_signature_response

MCP tool `get_signature` のresponse model。
docs/spec/mcp/tools/get-signature.md §3 のOutput envelope、および §4〜§10 のkind別signatureに対応する。
`signature` は kind別のunion相当であり、brewprint v1 modelではunion / discriminated objectを厳密表現できないため any + note で保持する。
task signature の endpoint はHTTP endpoint taskのみが持ち、MCP tool task自体に endpoint: true を付ける意味ではない。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/yaml/mcp/model/get_signature_response.yaml |

### Fields

| field | type | note |
|---|---|---|
| object | object_ref | 必須。対象object identity。 |
| signature | any | 必須。task / model / store / event / state / transition / field でshapeが異なるためanyで暫定表現する。 |
| doc | str | 任意。note由来の説明。null相当を取りうる。 |
| diagnostics | diagnostic_list | 必須。Diagnostic list。 |

