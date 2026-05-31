# inspect_request

MCP tool `inspect` のrequest model。
docs/spec/mcp/tools/inspect.md §2 のInputに対応する。
detailは brief / normal / full のenum相当。
MCP v1ではdetailによる厳密な返却差分は実装任意だが、未知の値はerrorとする。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/inspect_request.yaml |

### Fields

| field | type | note |
|---|---|---|
| selector | object_selector | 必須。inspect対象のObject selector。 |
| detail | str | 任意。brief / normal / full。省略時はnormal。enum制約はnoteで保持する。 |

