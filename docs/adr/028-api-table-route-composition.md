# 028: API Tableのroute構成とview定義

- **status**: accepted
- **date**: 2026-04-21
- **depends on**: ADR 002, ADR 005, ADR 009, ADR 017

## 背景

API Tableは `task` のうち `endpoint: true` なものを一覧するviewとして位置づけられている。
しかし既存定義だけでは、以下が未確定だった。

- API Tableをどのmoduleを起点に生成するか
- どのsubmoduleまで含めるか
- endpoint taskの `path` がfull pathなのかleaf pathなのか
- API Table上のrouteをどの責務で構成するか

この状態では、API Tableのrender実装時に実装者が暗黙補完するしかなく、
brewprintの「YAMLをsingle source of truthにする」方針と矛盾する。

## 決定

### 1. API Tableは独立view YAMLを持つ

API Tableはendpoint taskの単なる暗黙集約ではなく、view定義用のYAMLを持つ。
このYAMLは少なくとも以下の情報を持つ。

```yaml
id: auth_api
note: 認証API一覧
http_root_path: /api
modules:
  - module: app.auth
    include_submodules: true
  - module: app.admin.audit
    include_submodules: false
```

- `id`: API Tableの識別子。H1に使う
- `note`: API Tableの説明
- `http_root_path`: このAPI Tableが担当するHTTP root path
- `modules`: 集計対象moduleの一覧
- `modules[].module`: 絶対module path
- `modules[].include_submodules`: 配下moduleを収集対象に含めるか

### 2. endpoint taskの `path` はleaf pathとする

`task.endpoint: true` の場合、taskの `path` はfull pathではなくleaf pathを持つ。

- **optional**: `path` は省略可能。省略時は `task.id` をleaf nameとして使う
- **single segment**: `path` は `/` を含まないsingle segmentのみ許容する。`webhooks/stripe` のような複数セグメントは不正
- module階層によるURL構造の表現はmoduleディレクトリ構造に委ねる（ADR-002）

```yaml
- id: login
  type: task
  endpoint: true
  method: POST
  path: login       # id と異なる場合のみ明示的に上書き
```

```yaml
- id: start
  type: task
  endpoint: true
  method: GET
  # path省略 → task.id "start" をleaf nameとして使う
```

```yaml
# NG: /を含む複数セグメントは不正
- id: process_payment
  type: task
  endpoint: true
  method: POST
  path: webhooks/stripe   # ← 禁止。webhooksはmoduleディレクトリで表現すること
```

full pathはtask単体では決まらない。API Table viewの `http_root_path` とmodule階層情報を用いて構成する。

### 3. routeの構成責務はAPI Table viewが持つ

最終的なAPI routeは、API Table viewが以下の要素を連結して構成する。

```text
{http_root_path}/{section起点moduleからの相対module path}/{task.path}
```

- `http_root_path` はAPI Table YAMLから取得する
- `section起点moduleからの相対module path` は、`modules[]` の各要素を起点として算出する
- `task.path` はendpoint taskが持つleaf pathを使う

これにより、route全体の責務をview側に集約しつつ、task側はendpointのleaf名だけを持つ。

### 4. sectionは `modules[]` の各要素ごとに1つ持つ

API Tableの出力sectionは、`modules[]` の各要素ごとに1つ作る。

- `include_submodules: true` の場合、配下submoduleのendpointも同一sectionに含める
- 配下submoduleのendpointは、section内で相対module path付きで表示する
- `include_submodules: false` の場合、そのmodule直下のendpointのみを対象にする

例: `app.auth` をsection起点とし、`app.auth.oauth` を含む場合

- `login`
- `oauth/start`
- `oauth/callback`

のように表示する。

### 5. section起点から外れる場合は絶対path表示とする

API Table内の表示名は、原則としてsection起点moduleからの相対pathを使う。
ただし、section起点から相対化できないmoduleを同一sectionで扱う場合は、絶対module pathを使って表示する。

## 理由

- endpoint taskだけではroute全体の所属と集計範囲が決まらない
- module構造とHTTP root pathの責務を明示的に分離したい
- 実装者やLLMが暗黙補完せずにAPI Tableをrenderできるようにしたい
- `include_submodules` をview側に置くことで、module階層を活かした柔軟な集約ができる
- endpoint taskにfull pathを持たせると、module移動やview再利用時に責務が混ざる

却下した代替案：
- endpoint taskがfull pathを持つ → view定義なしでも表示できるが、集計単位と責務境界が不明瞭になる
- API Tableは暗黙にmodule全体を収集する → 起点module・対象範囲がYAMLに現れず、single source of truthに反する
- include_submodules時にsubmoduleごとに別sectionを切る → `modules[]` の宣言と出力構造が1対1に対応せず、view定義が読みにくい

## 影響

- API Table view用のspecを追加する必要がある
- `task.endpoint: true` の `path` はleaf pathとして解釈する
- API Tableのrender処理は `http_root_path` とmodule階層から最終routeを合成する
- API Tableのsection構造は `modules[]` の宣言に従う
- `spec/nodes.md` のendpoint taskの `path` 説明をleaf pathに更新する必要がある
- `spec/overview.md` のAPI Table説明にview YAML前提を反映する必要がある

## Evidence
- commit: 17fc138
- impl commit: tbd
- 参考: 特になし
