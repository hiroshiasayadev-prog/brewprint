# get_signature

MCP tool `get_signature` の公開contract。
MCP toolでありHTTP endpointではないため endpoint: true は付けない。
対象object単体の外形、signature、doc、diagnosticsを返す。

```mermaid
flowchart TD
  subgraph params
    request([request])
  end

  _start([Start]) ==> validate_request[validate_request]
  request --> validate_request
  validate_request --> validated_request([validated_request])

  validate_request ==> query_service[query_service]
  validated_request --> query_service
  resolved_project_store[(resolved_project_store)] -- "read" --> query_service
  query_service --> query_result([query_result])

  query_service ==> build_response[build_response]
  query_result --> build_response
  build_response --> response
  build_response ==> _end([End])

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff
  classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class validate_request,query_service,build_response taskNode
  class validated_request,query_result,response assetNode
  class resolved_project_store storeNode
  class _start,_end terminalNode
  class request boundaryNode
```

## Tasks

### get_signature

#### Params

| name | model | note |
|---|---|---|
| request | get_signature_request | — |

#### Returns

| name | model | source |
|---|---|---|
| response | get_signature_response | — |

### validate_request

tool input schemaとselectorを検証する。
selectorの有効な組み合わせや対象kind別の制約はv1 modelでは表現できないためrequest model側のnoteで保持する。

#### Params

| name | model | note |
|---|---|---|
| request | get_signature_request | — |

#### Returns

| name | model | source |
|---|---|---|
| validated_request | get_signature_request | — |

### query_service

QueryService境界。
Raw YAML ASTではなくResolvedProject上のsemantic objectを読む。
node / transition / field などselector対象ごとにsignatureを取得する。
対象kind別signatureのshape差分はv1 modelでは厳密表現できないためresponse model側のnoteで保持する。

#### Params

| name | model | note |
|---|---|---|
| request | get_signature_request | — |

#### Returns

| name | model | source |
|---|---|---|
| query_result | any | — |

#### Store access

| access | store |
|---|---|
| read | resolved_project_store |

### build_response

QueryService結果をMCP response payloadへ整形する。
object / signature / doc / diagnostics を返す。

#### Params

| name | model | note |
|---|---|---|
| query_result | any | — |

#### Returns

| name | model | source |
|---|---|---|
| response | get_signature_response | — |

