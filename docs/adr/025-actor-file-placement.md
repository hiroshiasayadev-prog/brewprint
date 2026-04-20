# 025: actorのファイル配置

- **status**: accepted
- **date**: 2026-04-20

## 背景

`actor` ノードのファイル配置が `spec/nodes.md` のノード種別一覧で `（未定）` のままだった。
他ノード種別（`task/*.yaml` / `model/*.yaml` / `store/*.yaml` / `state.yaml`）はすべて配置が確定しており、actorのみ未定だった。

## 決定

`actor` は独立ファイルを持たない。`state.yaml` 等、参照元ファイルのサブノードとして同居する。

```yaml
# state.yaml（例）
nodes:
  - id: end_user
    type: actor
    note: "サービスを利用するエンドユーザー"

  - id: idle
    type: state
    initial: true
    ...
```

## 理由

actorは「外部の人・システム」の宣言であり、それ自体がDAGや状態遷移のロジックを持たない。
独立ファイルを持つ意味が薄く、参照元ファイル（state.yaml 等）のサブノードとして定義するのが自然。
ADR-011の「サブノードはファイル内private」の仕組みをそのまま活用できる。

却下した代替案：
- `actor/*.yaml` の専用ディレクトリ → actorはロジックを持たず独立ファイルの意義がない

## 影響

- `spec/nodes.md` のノード種別一覧の `actor` 行のファイル配置を更新する
- actorが必要なファイル（state.yaml / sequence diagram ファイル等）にサブノードとして定義する

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: ADR-011（サブノード設計）
