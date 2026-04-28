# Go M5 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: Go実装 Milestone 5 第1段の State Diagram renderer vertical slice 実装チェックリスト（完了済み）

---

## 1. 目的

M5第1段では、残りrenderer群のうち State Diagram renderer だけを小さく縦切りで通す。

### M5第1段完了メモ

M5第1段は実装完了済みとして扱う。

- state / event / transition decodeを `rawyaml` / `source` に追加済み。
- `semantic.State` / `semantic.Event` / `semantic.Transition` と `TransitionsByFile` を追加済み。
- `resolve` でstate / event登録、transition構築、guard分岐validationを追加済み。
- `internal/render/state` を作成し、`stateDiagram-v2` rendererを実装済み。
- guard分岐はchoice pseudostateで描画済み。
- UC-001のState Diagram golden test 3本を追加済み。
- `go fmt ./...` / `go test ./...` は 2026-04-28 に通過済み。

対象は UC-001 の以下3本。

```text
auth/state.yaml      -> renders/auth/state-auth.md
order/state.yaml     -> renders/commerce/state-order.md
inventory/state.yaml -> renders/catalog/state-inventory.md
```

M5第1段の経路:

```text
state.yaml
  ↓ source / rawyaml decode
rawyaml.State / rawyaml.Event / rawyaml.Transition
  ↓ resolve / semantic build
semantic.Project
  ↓ render/state
stateDiagram-v2 markdown
  ↓ golden test
```

---

## 2. 前提ドキュメント

- `docs/doc-policy.md`
- `docs/impl/go-m4-task.md`
- `docs/TASKS.md` の Milestone 5
- `docs/adr/018-event-node.md`
- `docs/adr/019-state-node.md`
- `docs/adr/034-internal-event-source.md`
- `docs/adr/035-fsm-guard-branch-and-transition-identification.md`
- `docs/adr/046-render-output-placement-for-state-sequence-wireframe-preview.md`
- `docs/spec/views/state-diagram.md`

---

## 3. 実装原則

- State rendererも `rawyaml` をimportしない。
- State / Event / Transition decodeは `rawyaml` / `source` に置く。
- 名前解決とFSM構築は `resolve` / `semantic` に置く。
- State rendererは `semantic.Project` のstate/event/transition情報だけを読む。
- Sequence / Wireframe / ER / API rendererはM5第1段では作らない。

境界:

```text
source       -> rawyaml
resolve      -> rawyaml, semantic
render/state -> semantic
render/dag   -> semantic
query        -> semantic
```

禁止:

```text
render/state -> rawyaml
render/state内でRaw YAML走査
Sequence scenario decodeを先に作る
Wireframe rendererを混ぜる
```

---

## 4. M5第1段の対象範囲

- state node decode
- event node decode
- transitions decode
- semantic.State / semantic.Event / semantic.Transition
- state/event symbol登録
- transitions by file保持
- guard分岐 choice pseudostate render
- State / Events detail table render
- UC-001 state golden test

M5第1段では以下を最小に留める。

- transition actionのtask解決は、state renderにはraw action label相当で十分
- event payload / watches のreference indexはM3第2段候補
- transition indexはM3第2段候補
- wireframe fieldはdecode対象外でもよい

---

## 5. 受け入れ条件

- [x] `auth/state.yaml` をdecode / resolve / renderできる
- [x] `order/state.yaml` のguard分岐をchoice pseudostateでrenderできる
- [x] `inventory/state.yaml` の複数choice pseudostateをrenderできる
- [x] `render/state` が `rawyaml` をimportしていない
- [x] golden testが3本通る
- [x] 既存DAG / query / placement testが壊れていない
- [x] `go fmt ./...` が通る
- [x] `go test ./...` が通る

---

## 6. M5第1段で実装しないもの

- Sequence Diagram renderer
- Wireframe renderer
- ER renderer
- API Table renderer
- scenario decode
- preview harness
- event/reference index拡張
- transition reverse lookup index
