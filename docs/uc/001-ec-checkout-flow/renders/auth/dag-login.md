# login

**API**: [POST /api/login](../_cross/api.md)

認証情報を検証しアクセストークンを発行する。
実装メモ:
- user_db で credential を引き当て、form.password と照合
- 成功時 session_store に token を発行・保存
- 試行履歴は成否に関わらず login_log_db に記録
- request_context_store からリクエスト元情報（IP等）を取得しログに付与

```mermaid
flowchart TD
  subgraph params
    form([form: login_form])
  end
  subgraph initializes
    login_log_db[(login_log_db)]
  end

  _start([Start]) ==> login[login]
  form --> login

  user_db[(user_db)] -- "read" --> login
  request_context_store[(request_context_store)] -- "read" --> login
  login -- "write" --> session_store[(session_store)]
  login -- "write" --> login_log_db[(login_log_db)]

  login --> auth_token([auth_token: token])
  login ==> _end([End])

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff
  classDef initStoreNode fill:#F0C674,stroke:#B07820,color:#000
  classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class login taskNode
  class user_db,request_context_store,session_store,login_log_db storeNode
  class login_log_db initStoreNode
  class auth_token assetNode
  class _start,_end terminalNode
  class form boundaryNode
```

## Tasks

### login

#### Params

| name | model | note |
|---|---|---|
| form | login_form | UIフォーム入力。auth.model.login_form |

#### Returns

| name | model | source |
|---|---|---|
| auth_token | token | — |

#### Store access

| access | store |
|---|---|
| read | user_db |
| read | request_context_store |
| write | session_store |
| write | login_log_db |
