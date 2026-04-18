# 005: class diagram = endpoint view

- **status**: accepted
- **date**: 2026-04-17
- **depends on**: ADR 001, ADR 004

## 背景

class diagramをどう表現するかが未定だった。
Handler / Service / Repository をそれぞれclassとして定義する方式を検討したが、
それらはすでにDAG（task）・ER（store）で表現されている。
独立したノード種別としてclassを追加するとDAG/ERとの重複が生まれる。

## 決定

class diagramは独立したノード種別を持たない。
`endpoint: true` なtaskを抽出し、エンドポイント単位でグルーピングしたviewとしてrenderする。

```yaml
- id: login
  type: task
  endpoint: true
  method: POST
  path: /auth/login
  params: login_request    # assetのID
  returns: auth_token      # assetのID
```

上記がclass diagram viewでは以下のようにrenderされる：

```
class AuthAPI {
  +POST /auth/login(LoginRequest) AuthToken
  +DELETE /auth/logout(SessionId) void
}
```

グルーピングはモジュール（フォルダ階層）単位で行う（ADR 002に準拠）。

### endpointフラグの定義

| フィールド | 必須 | 内容 |
|---|---|---|
| `endpoint` | ✓ | `true` のとき class diagram viewに出力される |
| `method` | ✓ | HTTP method（GET / POST / PUT / DELETE / PATCH） |
| `path` | ✓ | URLパス（例：`/auth/login`） |
| `params` | 任意 | リクエストbodyのasset ID |
| `returns` | 任意 | レスポンスbodyのasset ID |

## 理由

- Handler / Service はDAGのtaskとして表現できる
- Repository はERのstoreとして表現できる
- classとして独立定義すべき要素はエンドポイントのI/Oシグネチャのみ
- これはtaskに`endpoint`フラグを追加するだけで表現できる
- spec/overviewの「classは独立ノード種別を持たない」方針（structのmethodsはnoteで足りる、それを超えるものはtaskとしてDAGに出す）と一致する

却下した代替案：
- classを独立ノード種別として定義する → DAG/ERとの重複が生まれ、管理コストが上がる
- class diagramをスコープ外にする → APIの全体像を俯瞰する手段がなくなる

## 影響

- `task` のスキーマに `endpoint` / `method` / `path` フィールドが追加される
- class diagram viewのrenderロジックは `endpoint: true` なtaskをモジュール単位でグルーピングする
- sequence diagramの `API` participantは class diagram viewへリンクする（ADR 004）
- spec/overview.md の「書ける図の一覧」を更新する必要がある
