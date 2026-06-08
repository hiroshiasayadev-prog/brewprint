# get_signature_request

MCP tool `get_signature` のrequest model。
docs/spec/mcp/tools/get-signature.md §2 のInputに対応する。
selectorの有効なobject / kind組み合わせは docs/spec/mcp/schema.md §1.2 のSelector support matrixに従う。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/get_signature_request.yaml |

### Fields

| field | type | note |
|---|---|---|
| selector | object_selector | 必須。対象objectを指定するObject selector。 |

