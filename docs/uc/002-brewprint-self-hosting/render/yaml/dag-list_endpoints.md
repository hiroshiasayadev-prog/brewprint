# list_endpoints

MCP tool `list_endpoints` の公開contract。
MCP toolでありHTTP endpointではないため endpoint: true は付けない。
API Table view YAMLに基づいてendpoint一覧を返す。
list_endpoints自体はHTTP endpointではなく、HTTP endpoint一覧を返すquery toolである。

```mermaid
flowchart TD
  subgraph params
    request([request: list_endpoints_request])
  end

  _start([Start]) ==> validate_request[validate_request]
  request --> validate_request
  validate_request --> validated_request([validated_request: list_endpoints_request])

  validate_request ==> query_service[query_service]
  validated_request --> query_service
  resolved_project_store[(resolved_project_store)] -- "read" --> query_service
  query_service --> query_result([query_result: any])

  query_service ==> build_response[build_response]
  query_result --> build_response
  build_response --> response([response: list_endpoints_response])
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

### list_endpoints

#### Params

| name | model | note |
|---|---|---|
| request | list_endpoints_request | — |

#### Returns

| name | model | source |
|---|---|---|
| response | list_endpoints_response | — |

### validate_request

tool input schemaとapi_table_idを検証する。
api_table_id は任意。
API Tableが複数存在し、api_table_idが省略された場合は全API Tableをtables[]に分けて返す。

#### Params

| name | model | note |
|---|---|---|
| request | list_endpoints_request | — |

#### Returns

| name | model | source |
|---|---|---|
| validated_request | list_endpoints_request | — |

### query_service

QueryService境界。
Raw YAML ASTではなくResolvedProject上のAPI Table viewとendpoint task情報を読む。
task(endpoint=true)の単純列挙ではなく、API Table view定義とADR-028のroute合成規則に従いfull pathを計算する。
収集対象endpointが0件のmodule-entryはsectionsには出さない。

#### Params

| name | model | note |
|---|---|---|
| request | list_endpoints_request | — |

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
tables[] / sections[] / endpoints[] / diagnostics を返す。
endpoint objectにはmethod / path / leaf_path / task / params / returns / sourceを含めうる。

#### Params

| name | model | note |
|---|---|---|
| query_result | any | — |

#### Returns

| name | model | source |
|---|---|---|
| response | list_endpoints_response | — |

