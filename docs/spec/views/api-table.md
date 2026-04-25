---
scope: docs/spec/api-table.md
status: draft
last_updated: 2026-04-25
summary: >
  API Table viewのYAML schemaとrenderルール定義。
  endpoint taskの収集範囲・route構成・Markdown table出力形式を定義する。
depends_on:
  - docs/adr/002-folder-as-namespace.md
  - docs/adr/005-class-diagram-as-endpoint-view.md
  - docs/adr/009-task-io-design.md
  - docs/adr/017-diagram-layers-and-scope.md
  - docs/adr/028-api-table-route-composition.md
  - docs/adr/030-yaml-file-type-declaration.md
---

# API Table仕様

## スコープ

API TableはApplicationレイヤーのviewであり、`endpoint: true` なtaskを人間とLLMが俯瞰するための一覧を提供する。
Mermaid描画は持たず、Markdownとしてrenderし、`list_endpoints` MCPツールで出力する。

API Tableは独立ノードではないが、暗黙集約でもない。API Table自身のYAMLを持ち、
どのmodule群をどう集計し、どのHTTP root配下として見せるかを明示する。

---

## YAML Schema

```yaml
as: api_table
id: auth_api
note: 認証API一覧
http_root_path: /api
modules:
  - module: app.auth
    include_submodules: true
  - module: app.admin.audit
    include_submodules: false
```

### フィールド定義

| フィールド | 必須 | 型 | 内容 |
|---|---|---|---|
| `as` | ✓ | string | ファイル種別宣言。view定義ファイルの識別に使う。`api_table` 固定（ADR-030） |
| `id` | ✓ | string | API Tableの識別子。MarkdownのH1に使う |
| `note` | 任意 | string | API Table全体の説明。H1直下に出力する |
| `http_root_path` | ✓ | string | このAPI Tableが担当するHTTP root path。`/api` のように先頭 `/` を含む |
| `modules` | ✓ | list\<module-entry\> | 集計対象moduleの一覧 |

### module-entry オブジェクト

| フィールド | 必須 | 型 | 内容 |
|---|---|---|---|
| `module` | ✓ | string | 絶対module path |
| `include_submodules` | 任意 | bool | `true` の場合、配下submoduleのendpoint taskも収集する。省略時は `false` |

---

## 収集ルール

### 対象

以下をすべて満たすtaskのみをAPI Tableに含める。

- `type: task`
- `endpoint: true`
- `modules[]` で指定されたmoduleに属する
- `include_submodules: true` の場合は、その配下submoduleも含む

`endpoint: true` でないtaskは完全に除外する。

### module単位のsection

API Tableの出力sectionは `modules[]` の各要素ごとに1つ持つ。

- `include_submodules: false` の場合、そのmodule直下のendpoint taskのみを対象にする
- `include_submodules: true` の場合、配下submoduleのendpoint taskも同一sectionに含める
- submoduleごとに別sectionは切らない
- 収集対象endpointが0件のmodule-entryはsectionを出力しない。parser error / warning にはしない

例: `app.auth` をsection起点moduleとし、`app.auth.oauth` を含む場合

- `login`
- `oauth/start`
- `oauth/callback`

を同じsectionに出力する。

---

## route構成ルール

endpoint taskの `path` はfull pathではなくleaf pathである。
`path` は省略可能で、省略時は `task.id` をleaf nameとして使う。
`path` は `/` を含まないsingle segmentのみ許容する（例: `stripe`、`login`）。`/` を含む複数セグメントは不正であり、URL階層はmoduleディレクトリ構造で表現する。

API Tableは各endpointの最終routeを以下の要素から構成する。

```text
{http_root_path}/{section起点moduleからの相対module path}/{task.path}
```

- `http_root_path` はAPI Table YAMLから取得する
- `section起点moduleからの相対module path` は、section対象moduleを起点に算出する
- `task.path` はendpoint taskが持つleaf pathを使う

### route構成例

前提:

```yaml
http_root_path: /api
modules:
  - module: app.auth
    include_submodules: true
```

収集対象task:

```yaml
# app.auth
- id: login
  type: task
  endpoint: true
  method: POST
  path: login

# app.auth.oauth
- id: start
  type: task
  endpoint: true
  method: GET
  path: start
```

renderされるroute:

- `login` → `/api/login`
- `oauth/start` → `/api/oauth/start`

### 表示名のルール

section内でのtask表示名は、section起点moduleからの相対pathを使う。

- 同一module直下: `login`
- submodule配下: `oauth/start`
- section起点moduleから相対化できないものを同一sectionで扱う場合: 絶対path表示

---

## 出力フォーマット

```markdown
# {id}

{note}

## Routes

- [{module-1}](#{module-1-anchor})
- [{module-2}](#{module-2-anchor})

## {module-1}

| task id | method | path | params | returns |
|---|---|---|---|---|
| login | POST | /api/login | credential | token |
| oauth/start | GET | /api/oauth/start | - | oauth_redirect |

## {module-2}

| task id | method | path | params | returns |
|---|---|---|---|---|
| list_logs | GET | /api/admin/audit/list_logs | - | audit_log_list |
```

- H1は `id`
- `note` がある場合はH1直下に出力。ない場合は省略
- `## Routes` を先頭に置き、各sectionへのリンク一覧を出す
- 本文は `modules[]` の順に `##` section を並べる
- 各sectionにはそのmodule-entryに対応するendpoint一覧tableを置く

### Routes一覧

`## Routes` には、各sectionの見出しリンクを列挙する。
リンクの表示名はsection見出しと同じ値を使う。
Routes一覧の順序は、出力されるsectionの順序と同じく `modules[]` の記述順に従う。収集対象endpointが0件で省略されたsectionはRoutes一覧にも出力しない。
リンク先anchorは、Markdown rendererの見出しID生成規則に従う。

### section見出し

section見出しは `modules[].module` の絶対module pathを使う。

例:

- `## app.auth`
- `## app.admin.audit`

---

## table列のrender

| 列 | 内容 |
|---|---|
| `task id` | section起点moduleからの相対path。必要に応じて絶対path |
| `method` | endpoint taskの `method` |
| `path` | API Tableが構成したfull route |
| `params` | `params[].model` を `<br/>` 区切りで列挙 |
| `returns` | `returns.model` を表示 |

### absent値

以下の欠損値はすべて `-` で表示する。

- `params` がない
- `returns` がない

### paramsのrender

`params` は `params[].model` を宣言順に `<br/>` 区切りで表示する。

例:

```yaml
params:
  - name: refresh_token
    model: refresh_token
  - name: client_info
    model: client_info
```

→ `refresh_token<br/>client_info`

`params` がない場合は `-` を表示する。

### returnsのrender

`returns` は単一のみなので、`returns.model` のみを表示する。

例:

```yaml
returns:
  name: auth_token
  model: token
```

→ `token`

`returns` がない場合は `-` を表示する。

---

## ソート順

各section内の行は `task id` のASCII昇順で並べる。

- `login`
- `oauth/callback`
- `oauth/start`

の順になる。

`modules[]` 自体のsection順はYAML記述順を維持する。

---

## 除外ルール

以下はAPI Tableに含めない。

- `endpoint: true` でないtask
- `model`
- `store`
- `actor`
- `event`
- `state`
- `branch`
- `fork`
- `join`
- `asset`（taskの `returns` から導出されるがtable対象ではない）

---

## render例

### YAML

```yaml
as: api_table
id: auth_api
note: 認証API一覧
http_root_path: /api
modules:
  - module: app.auth
    include_submodules: true
```

対象endpoint task:

```yaml
# app.auth
- id: login
  type: task
  endpoint: true
  method: POST
  path: login
  params:
    - name: credentials
      model: credential
  returns:
    name: auth_token
    model: token

# app.auth.oauth
- id: start
  type: task
  endpoint: true
  method: GET
  path: start
  returns:
    name: oauth_redirect
    model: oauth_redirect

# app.auth.oauth
- id: callback
  type: task
  endpoint: true
  method: GET
  path: callback
  params:
    - name: code
      model: oauth_code
    - name: state
      model: oauth_state
  returns:
    name: auth_token
    model: token
```

### Markdown出力

```markdown
# auth_api

認証API一覧

## Routes

- [app.auth](#appauth)

## app.auth

| task id | method | path | params | returns |
|---|---|---|---|---|
| login | POST | /api/login | credential | token |
| oauth/callback | GET | /api/oauth/callback | oauth_code<br/>oauth_state | token |
| oauth/start | GET | /api/oauth/start | - | oauth_redirect |
```
