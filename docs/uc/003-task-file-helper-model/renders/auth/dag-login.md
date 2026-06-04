# login

**API**: [POST /api/login](../_cross/api.md)

Authenticate a user and return a token.
This task file intentionally declares helper models below the main task
so renderer behavior for `## Private models` can be inspected.

```mermaid
flowchart TD
  subgraph params
    form([form])
    audit_context([audit_context])
  end

  _start([Start]) ==> login[login]
  form --> login
  audit_context --> login

  login --> auth_token([auth_token])
  login ==> _end([End])

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class login taskNode
  class auth_token assetNode
  class _start,_end terminalNode
  class form,audit_context boundaryNode
```

## Tasks

### login

#### Params

| name | model | note |
|---|---|---|
| form | login_form | Login form submitted by the client. |
| audit_context | login_audit_context | Task-local helper schema used only by this login task file. |

#### Returns

| name | model | source |
|---|---|---|
| auth_token | login_token | — |

