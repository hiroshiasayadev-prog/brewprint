# analyze_impact

MCP tool `analyze_impact` の公開contract。
MCP toolでありHTTP endpointではないため endpoint: true は付けない。
対象objectとchange kindから、意味づけ済みのimpact listを返す。
raw reference graphはget_reference_tree、source snippetはget_sourceに委ねる。

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

### analyze_impact

#### Params

| name | model | note |
|---|---|---|
| request | analyze_impact_request | — |

#### Returns

| name | model | source |
|---|---|---|
| response | analyze_impact_response | — |

### validate_request

tool input schemaとselector / change / scope_modules / max_impactsを検証する。
change はkindをdiscriminatorとするobject。
rename / remove / change_type / change_contract / change_transition_target / add のpayload制約はv1 modelでは表現できないためrequest model側のnoteで保持する。
kindとpayloadの不正な組み合わせは invalid_change_payload tool error。
unsupported selectorはtool errorではなく、空impacts + unsupported_selector diagnostic + coverageを含む通常responseとして返す。

#### Params

| name | model | note |
|---|---|---|
| request | analyze_impact_request | — |

#### Returns

| name | model | source |
|---|---|---|
| validated_request | analyze_impact_request | — |

### query_service

QueryService境界。
Raw YAML ASTではなくResolvedProject上のsemantic object / reference / render output mappingを読む。
change kindごとに探索深さとcollectorをtool側が決める。
severity / fixability / coverage / recommended_action / suggested_fixes を付与する。
flow wiring、sequence step、render output file粒度、型signature identityなどのcoverage差分はv1 modelでは厳密表現できないためresponse model側のnoteで保持する。
複雑なcollector分岐は現時点ではnoteで表現し、flowのbranch化はしない。

#### Params

| name | model | note |
|---|---|---|
| request | analyze_impact_request | — |

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
target / change / summary / impacts / coverage / assumptions / truncated / truncated_reasons / diagnostics を返す。
fixability=mechanical はsource locationと置換内容が一意に決まる場合のみ返す。

#### Params

| name | model | note |
|---|---|---|
| query_result | any | — |

#### Returns

| name | model | source |
|---|---|---|
| response | analyze_impact_response | — |

