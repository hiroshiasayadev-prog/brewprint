# Go M5 Sequence 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: Go実装 Milestone 5 第2段の Sequence Diagram renderer vertical slice 実装チェックリスト

---

## 1. 目的

M5第2段では、Sequence Diagram rendererをUC-001の2シナリオで縦切り実装する。

対象:

```text
views/scenarios/checkout_flow.yaml        -> renders/commerce/seq-checkout_flow.md
views/scenarios/payment_webhook_flow.yaml -> renders/commerce/seq-payment_webhook_flow.md
```

M5第2段では、scenario decode / semantic scenario build / transition解決 / sequence render / golden testまでを通す。

---

## 2. 実装範囲

- `as: sequence_diagram` view file decode
- semantic.SequenceScenario / SequenceStep
- stepのtransition解決 `(from_state, via, guard?)`
- step連続性validation
- source別participant生成
- actionなしUI self-message
- external actor message
- task同一ファイル内sub task reads/writes集約
- DB操作table
- UC-001 golden test 2本

---

## 3. 境界

```text
source          -> rawyaml
resolve         -> rawyaml, semantic
render/sequence -> semantic
```

禁止:

```text
render/sequence -> rawyaml
render/sequence内でscenario YAMLを直接読む
Wireframe / ER / API rendererを混ぜる
```

---

## 4. 受け入れ条件

- [x] checkout_flowをrenderできる
- [x] payment_webhook_flowをrenderできる
- [x] actionなしtransitionをUI self-messageで描画できる
- [x] external eventをActor->APIで描画できる
- [x] guardでtransitionを一意特定できる
- [x] sub task DB reads/writesを集約できる
- [x] `render/sequence` が `rawyaml` をimportしていない
- [x] golden test 2本が通る
- [x] `go fmt ./...` が通る
- [x] `go test ./...` が通る
