# 079: MCP transition ID policy の非 file-path 制約

- **status**: proposed
- **date**: 2026-05-13
- **depends_on**: ADR-035, ADR-048, ADR-078

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-078 では、MCP public contract に出る synthetic ID を file path based ではなく semantic anchor based に寄せる方針を accepted とした。

一方、ADR-078 では transition ID の最終 policy を意図的に未決として分離した。
理由は、現行 MCP spec / implementation / tests では transition ID が以下のような file path based ID として深く使われているためである。

```text
order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']
```

現行 `docs/spec/mcp/schema.md` では、transition ID は ADR-035 / ADR-048 の `(stateFileID, fromStateID, eventID, guard?)` に対応する文字列として定義されている。
また、`inspect` / `get_signature` / `get_references` / `get_reference_tree` / `analyze_impact` でもこの ID が selector / ObjectRef / Reference target として現れる。

しかし、MCP は Raw YAML / file browser ではなく ResolvedProject 上の semantic object query layer である。
transition ID の主成分に `state.yaml` のような file path を含めると、利用者は semantic transition ではなく YAML file layout を意識することになる。
これは ADR-078 の semantic anchor policy と整合しない。

さらに、MCP とほぼ機能互換の UI を将来提供する場合、MCP response / query surface に出る概念は UI 上にも現れる可能性が高い。
LLM に有用な semantic query surface は人間にも有用であるため、UI でも file path ではなく semantic object を主要概念として扱いたい。
この観点では、MCP ID だけでなく sequence scenario YAML の `state_file:` のような YAML file reference も将来的な見直し対象になりうる。

そのため、transition ID の最終 anchor をこのADRで確定すると、state machine semantic object の導入有無や sequence scenario authoring schema の file-path-free 化まで巻き込む。
これは本ADRのスコープを超える。

本ADRは、ADR-080 相当の後続ADRで state machine semantic object / file-path-free scenario reference を検討する前に、transition ID policy のうち ADR-080 に依存しない制約だけを先に固定する proposed ADR である。

## 決定

### 1. old file path based transition ID は canonical ID にしない

以下の現行 transition ID 形式は、MCP public contract の canonical transition ID として採用しない。

```text
<state-file-id>#<from-state>:<event>
<state-file-id>#<from-state>:<event>[<guard>]
```

例:

```text
order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']
```

この形式は file path を identity の主成分にしており、ADR-078 の semantic anchor based synthetic ID policy と整合しない。

### 2. old file path based transition ID は compatibility selector としても維持しない

brewprint はまだ安定公開済みの MCP contract を持つ段階ではない。
そのため、本変更では後方互換よりも semantic identity policy の一貫性を優先する。

旧 file path based transition ID は compatibility selector として維持しない。
旧形式が `selector.id` に渡された場合、MCP implementation は `invalid_selector` または `unsupported_transition_id` diagnostic を返してよい。
可能であれば diagnostic message で新しい canonical ID 形式を案内する。

例:

```text
old transition ID `order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']` is no longer supported; use the canonical transition ID form defined by the accepted transition ID policy.
```

最終的な canonical form は、state machine semantic object / transition anchor policy を決める後続ADRで確定する。

### 3. `state_file` は transition identity ではなく metadata とする

transition は現行 YAML では state file 内の `transitions[]` entry に由来する。
そのため、MCP response は `state_file` / `source.file` を返してよい。

ただし、`state_file` は canonical transition identity の主成分ではない。

`state_file` は以下のための metadata として扱う。

- source snippet 取得
- diagnostics
- debug
- 旧 authoring schema との対応
- migration / review 時の説明

MCP public contract における canonical transition identity は `id` であり、`state_file` ではない。

### 4. sequence scenario の transition 解決は MCP 利用者の責務ではない

現行 sequence scenario YAML は、以下のように `state_file` + `(from_state, via, guard?)` で transition を指定する。

```yaml
state_file: order/state.yaml
steps:
  - from_state: processing
    via: payment_webhook_received
    guard: "payload.status == 'succeeded'"
```

この schema が将来的に `state_machine:` などの semantic reference へ置き換わるかは、後続ADRで扱う。

ただし、いずれの authoring schema であっても、scenario YAML から transition への解決は ResolvedProject build / QueryService 側の責務である。
MCP 利用者や LLM が `state_file` + `(from_state, via, guard?)` を使って transition を再解決することは想定しない。

MCP response では、解決済み transition を canonical transition ID で返す。

### 5. transition は public QualifiedID を持たない

transition は public node ではない。
外部 YAML から `module.kind.id` 形式で参照可能な public QualifiedID は持たない。

transition は MCP query layer 上の synthetic ID で識別する。
したがって、transition ObjectRef では `qualified_id` を返さない。

### 6. transition ObjectRef の `visibility` は `generated` とする方針を採る

transition は YAML 上の `nodes:` entry ではなく、`transitions[]` entry から ResolvedProject 上に構築される queryable semantic object である。

そのため、ADR-078 の分類に従い、transition ObjectRef の `visibility` は `generated` とする方針を採る。

これは「YAMLに由来しない」という意味ではない。
`generated` は、public QualifiedID を持つ独立 node ではなく、ResolvedProject 上で queryable object として構築される semantic object を表す。

### 7. guard は transition identity に含める

ADR-035 により、FSM transition の識別キーは `(from, on, guard)` である。
同一 `(from, on)` に複数 transition が存在する場合、guard exact match により transition を一意特定する。

したがって、guard がある transition の canonical MCP ID には guard 文字列を含める。

最終的な ID anchor が state machine anchor であっても from state anchor であっても、この制約は変わらない。

guard 文字列は現行方針どおり YAML decode 後の文字列をそのまま使う。
trim、空白正規化、Unicode正規化、式AST比較は行わない。

これは ADR-035 / ADR-048 の guard exact match 方針を維持するためである。

### 8. guard なし transition は guard suffix を持たない

guard がない transition では、canonical ID に `[]` suffix を付けない。

最終的な canonical ID 形式は後続ADRで確定するが、guard がない場合に空の guard 表現を付けない方針は維持する。

### 9. transition query では `selector.object: "transition"` の併送を推奨する

ADR-078 と同じく、transition query の canonical selector は `selector.id` とする。

ただし、transition ID は `#` を含む synthetic ID であり、node QualifiedID ではない。
そのため、transition を selector.id で query する場合は、resolver の曖昧性を下げるために `selector.object: "transition"` の併送を推奨する。

```json
{
  "selector": {
    "object": "transition",
    "id": "<canonical-transition-id>"
  }
}
```

### 10. transition ID の final anchor は後続ADRに従う

本ADRでは、canonical transition ID の最終 anchor を確定しない。

以下は後続ADRで扱う。

- state machine semantic object を導入するか
- sequence scenario YAML の `state_file:` を file-path-free な semantic reference に置き換えるか
- transition ID の anchor を state machine QID にするか、from state QID にするか
- anchor に応じた `local_id` 形式

候補例:

```text
# state machine anchor 案
<state-machine-qualified-id>#<from-state>:<event>[<guard>]

# from state anchor 案
<from-state-qualified-id>#<event>[<guard>]
```

本ADRは、どちらの案になっても変わらない非 file-path 制約を先に固定する。

## Non-goals

本ADRでは以下を決めない。

- canonical transition ID の最終形式
- transition ID の semantic anchor
- state machine semantic object を導入するか
- sequence scenario YAML の `state_file:` を廃止するか
- `state_machine:` などの file-path-free scenario reference を導入するか
- transition ObjectRef の `anchor` 値
- transition ObjectRef の `local_id` 形式
- state diagram render の root semantic object
- backward compatibility adapter の具体実装

これらは後続ADRまたは MCP spec 反映時に扱う。

## 理由

### なぜ file path based transition ID を採用しないか

MCP query layer の利用者が欲しいのは、file path ではなく semantic transition である。

file path は source location として重要だが、object identity として前面に出すと、MCP が YAML file browser のように見えてしまう。
これは、MCP が Raw YAML AST ではなく semantic object query layer であるという ADR-047 / ADR-078 の基本方針と相性が悪い。

また、将来 MCP とほぼ同等の semantic UI を作る場合、file path based ID は UI の主要概念として表示しにくい。
そのため、transition ID から file path anchor を外す方針は、MCP だけでなく UI を見据えた設計としても重要である。

### なぜ compatibility selector を維持しないか

旧形式を compatibility selector として維持すると、implementation / tests / docs に old ID と new ID の両方が残る。
これにより、semantic identity policy が混ざり、後続の spec / implementation が複雑化する。

brewprint はまだ安定公開済みの MCP contract を持つ段階ではない。
そのため、後方互換を守るよりも、早い段階で file path based transition ID を仕様から外す方が適切である。

### なぜ guard を identity に含めるか

ADR-035 により、同一 `(from, on)` に複数 transition が存在できる。
その場合、guard が transition を一意に区別する。

guard を ID から外すと、以下の2つが同一 ID になってしまう。

```yaml
- from: processing
  on: payment_webhook_received
  to: confirmed
  guard: "payload.status == 'succeeded'"

- from: processing
  on: payment_webhook_received
  to: failed
  guard: "payload.status == 'failed'"
```

したがって、guard は canonical transition ID に残す。

### なぜ transition は `generated` か

transition は public QualifiedID を持つ node ではない。
一方で、private sub node のような file-private node でもない。

transition は `transitions[]` から ResolvedProject 上に構築され、MCPから query 可能な semantic edge object である。
ADR-078 の `generated` は、このような queryable derived object を表すために使う。

### なぜ final anchor を本ADRで決めないか

transition ID の anchor を決めるには、state machine semantic object を導入するかどうかを決める必要がある。

state machine object を導入する場合、transition ID は以下のように state machine anchor に寄せるのが自然である。

```text
<state-machine-qualified-id>#<from-state>:<event>[<guard>]
```

一方、state machine object を導入しない場合、from state anchor が最小変更案になる。

```text
<from-state-qualified-id>#<event>[<guard>]
```

この判断は、transition ID だけでなく sequence scenario YAML の `state_file:` 廃止、state diagram render の root、UI の semantic surface にも関わる。
そのため、本ADRでは final anchor を決めず、後続ADRで扱う。

## 却下した代替案

### 代替案A: 現行 `<state-file-id>#<from>:<event>[guard]` を維持する

- 利点: 現行 spec / implementation / tests への影響が小さい
- 欠点: MCP public contract の object identity に file path が残る。ADR-078 の semantic anchor policy と不整合

→ 却下。canonical ID としては採用しない。

### 代替案B: old file path based transition ID を compatibility selector として維持する

- 利点: 移行時の一時的な互換性を確保できる
- 欠点: old/new ID が併存し、implementation と docs が複雑化する。まだ安定公開済み contract ではない段階で守る価値が小さい

→ 却下。旧形式は受け付けず、diagnostic で案内する方針とする。

### 代替案C: guard を ID から外す

- 利点: IDが短くなる
- 欠点: guard分岐 transition を一意に識別できない

→ 却下。guard は transition identity に必要である。

### 代替案D: state machine object 非導入を本ADRで決める

- 利点: transition ID の final form まで本ADRで確定できる
- 欠点: UI / MCP semantic surface と sequence scenario YAML の file-path-free 化に関わる大きな判断を、transition ID policy に押し込むことになる

→ 却下。state machine semantic object の有無は後続ADRで扱う。

## 影響

### MCP spec への影響

本ADR受理後、以下の spec 更新が必要になる。

- `docs/spec/mcp/schema.md`
  - 現行 `<state-file-id>#<from>:<event>[guard]` 形式を canonical transition ID から外す
  - old file path based transition ID は compatibility selector として維持しないことを定義する
  - old ID が selector として渡された場合の diagnostic 方針を定義する
  - transition は public QualifiedID を持たない synthetic object であることを明記する
  - transition ObjectRef の `visibility` は `generated` とする方針を反映する
  - `state_file` は identity ではなく source / debug / authoring schema 対応 metadata として扱うことを明記する
  - final anchor / final local_id は後続ADRに従う unresolved scope として扱う

- `docs/spec/mcp/tools/get-signature.md`
  - transition signature の selector / ObjectRef examples は、後続ADRで final canonical ID が決まるまで更新を保留する
  - old ID の例には deprecated / unsupported 注記を追加するか、後続ADRで一括更新する

- `docs/spec/mcp/tools/inspect.md`
  - task / state / event / scenario / transition inspect に出る TransitionRef examples は、後続ADRで final canonical ID が決まるまで更新を保留する
  - scenario 解決は MCP 利用者ではなく semantic build 側の責務であることを補足する

- `docs/spec/mcp/tools/get-references.md`
  - `transition_from` / `transition_event` / `transition_to` / `transition_action` / `scenario_step_transition` の TransitionRef examples は後続ADRで final canonical ID が決まってから更新する

- `docs/spec/mcp/tools/get-reference-tree.md`
  - traversal root / reached node として transition が出る場合、old file path based ID を使わない方針を反映する

- `docs/spec/mcp/tools/analyze-impact.md`
  - transition impact 対象の selector / impacted object ID で old file path based ID を使わない方針を反映する

### naming spec への影響

MCP transition ID は YAML QualifiedID grammar ではない。

`docs/spec/naming.md` には、ADR-078 の MCP synthetic ID 境界と同じく、以下を補足する必要がある可能性がある。

- transition ID は MCP synthetic ID であり、YAML から書ける QualifiedID ではない
- `#` suffix は MCP query layer の private / generated object namespace である
- transition は public QualifiedID を持つ node ではない
- final anchor は後続ADRに従う

### 実装への影響

実装では、内部 index key として `(stateFileID, fromStateID, eventID, guard?)` を維持してよい。
本ADRが定めるのは MCP public contract 上の ID policy である。

ただし、QueryService / MCP adapter は transition ObjectRef を返す際、old file path based ID を返してはならない。
final canonical ID の formatter は後続ADRで確定する。

影響が予想される実装領域:

- transition ID formatter / parser
- selector resolver
- TransitionRef / ObjectRef builder
- `get_signature(transition)`
- `inspect(state)` / `inspect(event)` / `inspect(task)` / `inspect(transition)`
- `get_references`
- `get_reference_tree`
- `analyze_impact`
- MCP server request / response tests

old file path based ID は compatibility selector として受け付けない。
旧形式が渡された場合、resolver は `invalid_selector` または `unsupported_transition_id` diagnostic を返し、可能であれば canonical ID 形式を案内する。

### UC / fixture への影響

MCP response golden / UC-002 self-hosting YAML / MCP contract examples に old transition ID が含まれている場合、更新が必要になる。

sequence scenario YAML の `state_file` / `steps[].from_state` / `steps[].via` / `steps[].guard` をどうするかは後続ADRで扱う。
本ADRでは、少なくとも MCP public contract の transition identity から file path anchor を外す方針を固定する。

### ADR-078 への影響

ADR-078 で未決として残した transition ID policy のうち、非 file-path 制約を本ADRが具体化する。

ADR-078 の semantic anchor based synthetic ID 原則に従い、transition ID から file path anchor を外す。
final anchor は後続ADRに委ねる。

## Acceptance 前の確認事項

本ADRは proposed として起票する。
accepted に進める前に、以下を確認する。

1. old file path based transition ID を canonical / compatibility selector として採用しない方針で問題ないか
2. old ID が渡された場合の diagnostic message で十分に移行先を案内できるか
3. guard 文字列を ID に含める場合の escaping / JSON transport / Mermaid との相互作用に問題がないか
4. `visibility: generated` が transition に対して妥当か
5. internal index key `(stateFileID, fromStateID, eventID, guard?)` と MCP public ID を分離して実装できるか
6. final anchor を後続ADRへ送った状態でも、本ADRが非 file-path 制約として十分に閉じているか

## Evidence

- commit: 5ae7769
- impl commit: tbd
- close boundary: M15 / `v1.1.0-spec` では follow-up scope として deferred。実装は含めない。
- 参考: ADR-035 FSM guard分岐とtransition識別、ADR-048 transition index、ADR-078 semantic anchor synthetic ID policy
