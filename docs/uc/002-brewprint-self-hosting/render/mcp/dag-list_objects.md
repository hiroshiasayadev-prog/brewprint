# list_objects

MCP tool `list_objects` の公開contract。
MCP toolでありHTTP endpointではないため endpoint: true は付けない。
project内のsemantic object一覧を返す探索用tool。
詳細なsignature / references / inspect情報は別toolへ委ねる。

```mermaid
flowchart TD
  subgraph params
    request([request: list_objects_request])
  end

  _start([Start]) ==> validate_request[validate_request]
  request --> validate_request
  validate_request --> validated_request([validated_request: list_objects_request])

  validate_request ==> query_service[query_service]
  validated_request --> query_service
  resolved_project_store[(resolved_project_store)] -- "read" --> query_service
  query_service --> query_result([query_result: any])

  query_service ==> build_response[build_response]
  query_result --> build_response
  build_response --> response([response: list_objects_response])
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

### list_objects

#### Params

| name | model | note |
|---|---|---|
| request | list_objects_request | — |

#### Returns

| name | model | source |
|---|---|---|
| response | list_objects_response | — |

### validate_request

tool input schemaを検証する。
object / kind / module / file はすべて任意filter。
object / kind のenum制約はv1 modelでは表現できないためrequest model側のnoteで保持する。

#### Params

| name | model | note |
|---|---|---|
| request | list_objects_request | — |

#### Returns

| name | model | source |
|---|---|---|
| validated_request | list_objects_request | — |

### query_service

QueryService境界。
Raw YAML ASTではなくResolvedProject上のsemantic object indexを読む。
node / view / transition / field をfilter条件に応じて列挙する。

#### Params

| name | model | note |
|---|---|---|
| request | list_objects_request | — |

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
objects[] と diagnostics[] を返す。

#### Params

| name | model | note |
|---|---|---|
| query_result | any | — |

#### Returns

| name | model | source |
|---|---|---|
| response | list_objects_response | — |

