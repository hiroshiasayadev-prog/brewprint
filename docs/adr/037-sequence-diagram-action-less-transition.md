# 037: sequence diagram における action なし transition の描画ルール

- **status**: accepted
- **date**: 2026-04-24

## 背景

ADR-019 で `transition.action` は任意フィールドと定義されている。
action なし transition は「UIの状態変化のみ」を表し、API 呼び出しを伴わない画面遷移がその典型例。

例: `order/state.yaml` の `cart → checkout_screen`（`view_checkout` event）

```yaml
- from: cart
  on: view_checkout
  to: checkout_screen
```

ADR-004 / ADR-032 / ADR-036 はいずれも「transition.action → task → endpoint」という経路で
API participant と矢印ラベルを解決することを前提としており、
action なし transition を sequence diagram の step に指定した場合の描画ルールが未定義だった。

## 決定

action なし transition を step に指定した場合、`UI->>UI: {event ID}` の self-message として描画する。

### ルール

| transition.action | 起点→終点 | ラベル | API応答矢印 |
|---|---|---|---|
| あり（既存ルール） | `UI→API`（source=ui の場合） | `METHOD path` | `API→UI` |
| **なし（本ADR）** | **`UI->>UI`** | **event ID** | **なし** |

- API participant は action なし step では生成しない
- DB 操作 table への出力もなし（reads / writes がないため）
- step index は全 step 通し番号で付与。action なし step は table に行が出ないだけ

### ラベルを event ID にする根拠

action なし transition は HTTP 呼び出しを伴わないため `METHOD path` のソース（task.method / task.path）が存在しない。
`internal` / `er(kind≠db)` の self-message も同じ理由で event ID をラベルとしており、
ラベル選択の原則（ADR-036）を踏襲して event ID を採用する。

## 理由

### UI self-message（案 B）を選択した理由

検討した代替案:

| 案 | 内容 |
|---|---|
| A | action なし transition の step 指定をパーサーエラーにする |
| **B** | **`UI->>UI: {event ID}` で描画（本決定）** |
| C | render 時に暗黙スキップ |
| D | `Note over UI: {event ID}` として描画 |

A（パーサーエラー）は「ADR-004 の粗い粒度＝レイヤー越えのみ」という読み過ぎに基づく案。
ADR-004 は粒度の方針を定めるものであり、UI 内遷移の描画を禁止していない。
UML sequence diagram の一般的な慣習として UI 自己矢印で画面遷移を表現するパターンは確立されており、
シナリオの連続した state chain を 1 つの図で表現したいケースに対応できる。

C（暗黙スキップ）は step index と DB 操作 table の対応関係を不透明にするため却下。
D（Note）は ADR-004 で「例外・エラーフローの表現」として位置づけられており、
happy path のフロー記述に note を使うのは誤用になる。

### `uiChange` のような新 source 種別を設けなかった理由

`source: ui` かつ action なし、という条件は既存フィールドの組み合わせで表現できる。
新フィールドを増やすと ADR-018 の source 定義・ADR-036 の矢印ルール表を拡張する必要が生じ、
既存概念との整合コストが高い。最小変更で解決できるため新 source は設けない。

## 影響

- `spec/views/sequence-diagram.md` の矢印ラベル表に action なしのケースを追記する
- ADR-036 のルール表は `event.source` ベースであり action の有無は扱っていないため、
  本 ADR は ADR-036 の **補完** として位置づける（supersede はしない）

## Evidence
- commit: dba8a2b
- impl commit: tbd
- 参考: UML 2.x Sequence Diagram self-message 記法
