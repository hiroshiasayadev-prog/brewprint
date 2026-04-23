# 034: eventのsourceに `internal` を追加

- **status**: accepted
- **date**: 2026-04-24

## 背景

ADR-018でeventの `source` を `ui` / `external` / `er` の3種で定義した。しかしUC-001 EC Checkout Flow の `auth/state.yaml` 設計中に、以下のイベントがいずれのsourceにも綺麗に当てはまらないことが判明した。

- `login_succeeded` / `login_failed`（`loading` → `authenticated` / `error` を駆動）
- 一般に「taskの完了/成否」をトリガーとしてFSM遷移を起こすevent全般

これらのeventは実行時にアプリ内部（FSM runtime）で発火するもので、以下の理由から既存の3カテゴリに無理に押し込めるのは設計思想を歪める。

- **`ui` に入れる案**: ADR-018で「ユーザー操作（クリック・フォーム送信等）」と定義されている。task完了はユーザー操作ではない。`ui` に入れるとカテゴリが膨張し、LLMが `source: ui` を見て「ユーザーが何かした」と推論できなくなる。
- **`external` に入れる案**: ADR-018で「webhook・MQ・WebSocket等の外部システムからの入力」と定義されている。task完了通知はアプリ内部で起きる事象であり、物理的に外部システムから来るわけではない。`spec/views/state-diagram.md` のYAML入力例が `login_succeeded` を `source: external, actor: auth_server` と書いていたのは、ADR-018策定時にFSM駆動まで設計が届いていなかった仕様の穴と見るべき。
- **`er` に入れる案**: storeの値変化監視だが、`login_failed` のようにstoreに書かないパスが表現不能。FSMの主駆動軸として `er` を使うのは `er` の本来の性格（変化の副次的観測）を歪める。

## 決定

eventの `source` enum に **`internal`** を追加する。

```yaml
- id: login_succeeded
  type: event
  source: internal
  payload:
    model: token
  note: "auth.task.login 成功時にFSM runtimeが発火"

- id: login_failed
  type: event
  source: internal
  note: "auth.task.login 失敗時にFSM runtimeが発火"
```

### フィールド要件

| source | actor | watches | payload |
|--------|-------|---------|---------|
| `ui` | 不要 | 不要 | 任意 |
| `external` | 必須 | 不要 | 任意 |
| `er` | 不要 | 必須 | 任意 |
| `internal` | 不要 | 不要 | 任意 |

`internal` は追加の必須フィールドを持たない。何を監視するか（どのtaskの完了か、どの内部タイマーか等）は `note:` で人間向けに書く。

### sourceカテゴリの意味（ADR-018からの拡張）

| source | 発火元 | 典型的用途 |
|--------|-------|----------|
| `ui` | ユーザー操作 | クリック・フォーム送信 |
| `external` | 外部システム | webhook・MQ・WebSocket・スケジューラー |
| `er` | store値変化 | DB状態変化の観測 |
| `internal` | **アプリ内部** | **task完了・FSM runtime内部処理・内部タイマー** |

## 理由

**4カテゴリの直交性**: 発火元の「所属」を軸に直交する分類になる。

- `ui`: ユーザー起点
- `external`: 外部システム起点
- `er`: データ変化起点
- `internal`: アプリ内部起点

**LLMへの説明力**: sourceを見るだけで「このeventを誰が駆動するか」を判別できる。`source: internal` を見たLLMは「このevent発火はアプリ実装側の責務」と即座に理解できる。

**FSM runtime責務の明示**: `source: internal` の存在は、brewprintの実行モデルにおいて「taskの戻り値やアプリ内部状態を見てeventを発火する層」（= FSM runtime）があることを言語レベルで示す。

**最小限の拡張**: enum追加のみで、新フィールドを導入しない。`er` の `watches:` のように監視対象を明示するフィールドを追加する案も検討したが、task ID・タイマーID・その他多様な内部発火源を統一的に表現するフィールド設計が困難なため、当面は `note:` での人間向け記述にとどめる。

### 却下した代替案

- **ADR-018をsupersedeして書き直す案**: 既存3カテゴリの定義は変わらないため、supersedeは過剰。拡張ADRとして独立させるほうが履歴が辿りやすい。
- **`source: task` + `from: <task_id>` 形式**: task完了以外の内部発火源（タイマー等）を排除してしまう。`internal` のほうが抽象度が適切。

## 影響

- `spec/nodes.md` の event セクションのsource enum表・「sourceの意味」表を更新
- ADR-018は拡張対象。supersedeはしない（既存3カテゴリの定義は変更なし）
- `spec/views/state-diagram.md` のYAML入力例は `source: external, actor: auth_server` と書かれているが、これは本ADR起票前の記述。整合のためには `source: internal` への書き換えが望ましいが、本ADRのスコープ外（別途更新可）
- UC-001 の `auth/state.yaml` は本ADRに基づいて `login_succeeded` / `login_failed` を `source: internal` として記述

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: UML 2.x Activity Diagramにおける内部トリガーeventの概念、FSMランタイム実装慣習
