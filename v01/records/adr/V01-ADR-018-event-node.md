# V01-ADR-018: eventノード設計

- **status**: accepted
- **date**: 2026-04-19

## 背景

V01-ADR-017にてeventはApplicationレイヤーの図（Sequence Diagram / State Diagram）にのみ登場し、DAGには出ないことが確定した。本ADRではeventノードのYAML上のフィールド定義を確定する。

eventの発生源として `ui / time / external / er` の4種は `spec/overview.md` に概念として記載済みだったが、スキーマは未定だった。

## 決定

### eventはnode typeとして`nodes:`に書く

```yaml
- id: login_submitted
  type: event
  source: ui
  payload:
    model: login_form
  note: "ログインフォームのsubmit"
```

### フィールド定義

| フィールド | 必須 | 内容 |
|-----------|------|------|
| `source` | ✓ | `ui` / `external` / `er` |
| `actor` | `external`のみ必須 | 発火元のactor ID。`source=external` の場合、対応する `type: actor` ノードの宣言が必要 |
| `payload` | 任意 | イベントが運ぶデータのmodel参照 |
| `watches` | `er`のみ | 変化を監視するstore ID |
| `note` | 任意 | 発火条件・詳細。複雑な条件はここに書く |

### sourceごとの使い方

**`ui`**: ユーザー操作（ボタンクリック・フォーム送信等）。payloadにフォームデータ等のmodelを指定する。

```yaml
- id: login_submitted
  type: event
  source: ui
  payload:
    model: login_form
  note: "ログインフォームのsubmit"
```

**`external`**: webhook・message queue・WebSocket等の外部システム（スケジューラー含む）からの入力。`actor:` フィールドで発火元actorを必ず明示する。参照先の `type: actor` ノードはbrewprintのいずれかのファイルに宣言されていること。

```yaml
- id: payment_webhook_received
  type: event
  source: external
  actor: stripe
  payload:
    model: payment_event
  note: "Stripeからの決済完了通知"

- id: daily_batch_triggered
  type: event
  source: external
  actor: scheduler
  note: "毎日0時に発火"
```

**`er`**: storeの値が変化したことによって発火するイベント。`watches`に監視対象のstore IDを指定する。

```yaml
- id: connection_status_changed
  type: event
  source: er
  watches: db_connection_store
  payload:
    model: connection_status
  note: "db_connection_storeのstatusが変化した時"
```

### DAGには登場しない

V01-ADR-017のレイヤー原則に従い、eventノードはDAGのflow:内で参照されない。
Sequence DiagramおよびState DiagramのYAMLにのみ登場する。

## 理由

### node typeとして`nodes:`に書く

Sequence Diagram・State Diagramの両方から参照されうるため、独立したノードとして定義する必要がある。main nodeのフィールド（案B）にすると、State Diagramのtransition起点としての参照が難しくなる。flow:の構文（案C）にすると、複数の図から参照する際に定義が重複する。

### payloadをmodel参照で持つ

eventのペイロード構造をnoteのみで記述すると、LLMがイベントの「何が来るか」を把握するためにnoteを解析する必要が生じる。model IDで参照することで機械的な参照解決が可能になり、Claude CodeがSequence Diagramを読む際のコンテキストが明確になる。

### `watches`フィールド（`er`専用）

`er`イベントは「どのstoreの変化を監視するか」が本質的な情報であり、noteに埋めると機械的検証ができない。store IDへの明示的な参照として独立フィールドにすることで、store定義との整合チェックが可能になる。

## 影響

- `spec/overview.md` の「ノード種別」テーブルに `event` を追加する
- `spec/overview.md` の「triggerの発生源」セクションを本ADRへの参照に更新する（`time` 削除・3種に変更）
- `spec/overview.md` のsequence diagramのparticipant対応表の `event(source=ui)` 記述を本ADRに基づき整理する
- stateノード（FSM）の設計はV01-ADR-019へ

## Evidence
- commit: 0a2759a
- impl commit: tbd
- 参考: 特になし
