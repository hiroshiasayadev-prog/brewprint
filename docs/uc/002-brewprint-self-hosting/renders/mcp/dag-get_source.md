# get_source

MCP tool `get_source` の公開contract。
MCP toolでありHTTP endpointではないため endpoint: true は付けない。
semantic objectに対応するYAML source snippetを返す。
Raw YAML AST公開APIではなく、ResolvedProject上のsemantic objectに紐づくsource補助情報として扱う。

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

### get_source

#### Params

| name | model | note |
|---|---|---|
| request | get_source_request | — |

#### Returns

| name | model | source |
|---|---|---|
| response | get_source_response | — |

### validate_request

tool input schemaとselector / fallbackを検証する。
fallback は file / error のenumだが、v1 modelではenum制約を表現できないためrequest model側のnoteで保持する。

#### Params

| name | model | note |
|---|---|---|
| request | get_source_request | — |

#### Returns

| name | model | source |
|---|---|---|
| validated_request | get_source_request | — |

### query_service

QueryService境界。
Raw YAML ASTではなくResolvedProject上のsemantic objectからsource locationを引く。
object単位rangeが特定できない場合は fallback 設定に従い、同一FileID全体へのfallbackまたはtool errorへ分岐する。
この分岐は現時点ではnoteで表現し、flowのbranch化はしない。

#### Params

| name | model | note |
|---|---|---|
| request | get_source_request | — |

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
object / source / snippet / fallback / diagnostics を返す。

#### Params

| name | model | note |
|---|---|---|
| query_result | any | — |

#### Returns

| name | model | source |
|---|---|---|
| response | get_source_response | — |

