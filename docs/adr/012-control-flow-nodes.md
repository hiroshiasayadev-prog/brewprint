# 012: 制御フローノード（branch / fork / join）

- **status**: accepted
- **date**: 2026-04-19

## 背景

ADR-011にて `foreach` / `cond` / `initializes` の設計をADR-012に委譲していた。
議論の過程でこれらは概念が異なることが明確になったため、本ADRでは制御フローノード（条件分岐・並列実行・合流）のみを扱う。`foreach` はADR-013、`initializes` はADR-014に委譲。

また、ADR-011では `cond` という仮称を使っていたが、本ADRで正式名称を確定する。

## 決定

制御フローノードとして以下の3種を導入する。

| ノード | 意味 | ペア |
|--------|------|------|
| `branch` | 条件に応じて後続パスを**1本だけ**選ぶ（排他分岐） | ペア不要 |
| `fork` | 後続パスを**すべて並列実行**する | `join` と必ずペアで使う |
| `join` | `fork` で分岐した全ブランチが揃うまで待ち、合流する | `fork` と必ずペアで使う |

### branch

```yaml
- id: route_by_role
  type: branch
  params:
    - name: user
      model: user
  note: "user.roleに応じてadmin_flow / user_flowのどちらかに進む"
```

- 後続パスは1本のみ実行される（排他）
- 合流点は明示しない。後続taskがDAGのedge構造上で複数パスから受け取る形になる場合、それを暗黙の合流点として読む
- 分岐後に合流しないケース（パスがそのまま終端する）も許容

### fork / join

```yaml
- id: fan_out
  type: fork
  params:
    - name: request
      model: analysis_request
  note: "静的解析・動的解析・依存チェックを並列実行する"

- id: aggregate
  type: join
  params:
    - name: static_result
      model: static_result
    - name: dynamic_result
      model: dynamic_result
    - name: dep_result
      model: dep_result
  returns:
    name: full_report
    model: full_report
  note: "3ブランチの結果を結合してfull_reportを生成する"
```

- `fork` → `join` は必ずペアで使う。`fork` のみ・`join` のみは不正
- `join` は対応する `fork` のID（または後続で設計するedge定義）で紐付ける

## 理由

### `branch`（旧 `cond`）の命名

`cond` はLisp/関数型言語の慣習。DAGツールの文脈では Airflow の `BranchPythonOperator` に由来する `branch` が定着しており、ADR-006の方針（業界標準の語彙から採る）と一致する。

`if` は二択限定のイメージが強く多分岐への拡張が直感的でない。`switch` は命令型言語の構文語という印象が強い。

### `fork` の命名

Unix/並列計算全般で「並列実行の起点」として定着している。`branch`（排他）との対比が語彙レベルで明快になる。`parallel` は意味は正確だが形容詞的でノード名として冗長。`fan_out` はアンダースコア込みで長い。

### `join` の明示必須

`fork` の合流点は暗黙に読めない（どのbranchが揃えば進むか、何を待つかが不明確）。`branch` との違いとして、並列実行の完了条件は明示的に記述する必要がある。

却下した代替案：
- `join` 不要（edge構造から暗黙的に読む）→ 「全ブランチ完了待ち」という意味論がDAGから読めないため却下

## 影響

- `spec/overview.md` のノード種別表に `branch` / `fork` / `join` を追記する
- ADR-011で使われていた仮称 `cond` は本ADRにより `branch` に確定
- `foreach` の設計は ADR-013 へ
- `initializes` の設計は ADR-014 へ
- ファイル内edgeの記述構造（`fork`-`join` の紐付け方法を含む）は ADR-015 へ委譲

## Evidence
- commit: e11491a
- impl commit: tbd
- 参考: Airflow BranchPythonOperator参考、BPMN Exclusive Gateway / Parallel Gateway参考
