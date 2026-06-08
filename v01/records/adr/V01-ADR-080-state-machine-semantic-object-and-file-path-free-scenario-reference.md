# V01-ADR-080: state machine semantic object と file-path-free scenario reference

- **status**: proposed
- **date**: 2026-05-13
- **depends_on**: V01-ADR-030, V01-ADR-032, V01-ADR-035, V01-ADR-078, V01-ADR-079

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

V01-ADR-078 では、MCP public contract に出る synthetic ID を file path based ではなく semantic anchor based に寄せる方針を accepted とした。

V01-ADR-079 では、transition ID policy のうち V01-ADR-080 に依存しない制約を proposed として整理した。
特に、old file path based transition ID を canonical / compatibility selector として採用しないこと、`state_file` は transition identity ではなく metadata であること、transition は public QualifiedID を持たない synthetic object であることを定めた。

しかし、transition ID の final anchor を決めるには、state machine 自体を semantic object として扱うかどうかを決める必要がある。

現行の sequence scenario YAML は、以下のように `state_file` で state 定義 YAML file を参照している。

```yaml
as: sequence_diagram
id: payment_webhook_success
state_file: order/state.yaml
steps:
  - from_state: processing
    via: payment_webhook_received
    guard: "payload.status == 'succeeded'"
```

これは YAML authoring layer に file path reference が残っていることを意味する。
MCP だけを考えれば、semantic build 後に canonical transition ID へ解決すればよい。
しかし、brewprint では MCP とほぼ機能互換の UI を将来的に提供する想定がある。
LLM に有用な semantic query surface は人間にも有用であり、UI でも file path ではなく semantic object を主要概念として扱いたい。

その観点では、`state_file: order/state.yaml` のような YAML file path reference は UI / MCP 共通設計面に漏れる design debt になる。
UI 上で `order/state.yaml` を主要選択肢として表示するより、`order.state_machine.order_flow` のような semantic object を表示・選択できる方が自然である。

本ADRは、state machine を semantic object として導入し、sequence scenario が file path ではなく state machine semantic reference を持つ方向を定める。

## 決定

### 1. state machine を semantic object として導入する

brewprint は state machine を first-class semantic object として扱う。

state machine は、state / event / transitions を束ねる Application layer の semantic root である。
State Diagram render、Sequence Diagram scenario、MCP inspect、将来 UI は state file path ではなく state machine semantic object を主要な対象として扱う。

state machine は public QualifiedID を持つ。

形式:

```text
<module-path>.state_machine.<id>
```

例:

```text
order.state_machine.order_flow
auth.state_machine.login_flow
```

### 2. state machine YAML に semantic ID を持たせる

state machine 定義 YAML は、file path ではなく YAML 内の semantic ID で識別する。

初期案として、state machine YAML は `as: state_machine` と `id:` を持つ。

```yaml
as: state_machine
id: order_flow

nodes:
  - id: processing
    type: state

  - id: payment_webhook_received
    type: event
    source: external
    actor: stripe

transitions:
  - from: processing
    on: payment_webhook_received
    to: confirmed
    guard: "payload.status == 'succeeded'"
```

この場合、module `order` 配下にある `id: order_flow` は以下の QualifiedID を持つ。

```text
order.state_machine.order_flow
```

`as: state_machine` の詳細な file classification、既存 `nodes:` file との互換性、`state.yaml` からの migration 形式は spec 反映時に確定する。
ただし、本ADRでは state machine が semantic ID を持つべきことを決定する。

### 3. sequence scenario は `state_machine` で semantic reference する

sequence scenario YAML は、file path の `state_file` ではなく semantic reference の `state_machine` を使う。

```yaml
as: sequence_diagram
id: payment_webhook_success
state_machine: order.state_machine.order_flow
steps:
  - from_state: processing
    via: payment_webhook_received
    guard: "payload.status == 'succeeded'"
```

`state_machine` は state machine QualifiedID である。
同一 module 内では短縮参照を許すか、常に QualifiedID を要求するかは spec 反映時に定義する。

`state_file` は新規 authoring では使用しない。
既存 scenario YAML の migration / compatibility policy は後続 task で扱う。

### 4. `state_file` は source metadata に降格する

state machine 定義は引き続き YAML file に保存される。
そのため、MCP response / diagnostics / get_source / debug では source file path を返してよい。

ただし、state file path は semantic identity ではない。

```json
{
  "object": "node",
  "kind": "state_machine",
  "id": "order.state_machine.order_flow",
  "qualified_id": "order.state_machine.order_flow",
  "visibility": "public",
  "source": {
    "file": "order/state.yaml"
  }
}
```

file path は source metadata であり、MCP / UI の主要 identity は state machine QualifiedID である。

### 5. State Diagram render は state machine semantic object を root とする

現行 spec では、State Diagram は `1 file = 1 FSM = 1枚の図` としている。
本ADR後は、State Diagram render の semantic root は state machine object とする。

出力単位は当面 `1 state machine = 1枚の図` とする。

file path は render 入力の内部探索や source 表示には使ってよいが、render title / index / MCP / UI の主要識別子は state machine QualifiedID を使う。

### 6. Sequence Diagram scenario の transition 解決は state machine semantic object から行う

sequence scenario は `state_machine` で対象 state machine を指定し、steps の `(from_state, via, guard?)` で transition を指定する。

```text
(state_machine QID, from_state, via, guard?) -> transition
```

この解決は semantic build / QueryService 側の責務である。
MCP 利用者や LLM が file path から transition を再解決することは想定しない。

### 7. transition canonical ID は state machine anchor を基本候補とする

V01-ADR-079 では transition final anchor を未決として残した。

本ADRにより state machine semantic object を導入するため、transition canonical ID は state machine QID を anchor とする方針を基本候補とする。

形式候補:

```text
<state-machine-qualified-id>#<from-state>:<event>[<guard>]
```

例:

```text
order.state_machine.order_flow#processing:payment_webhook_received[payload.status == 'succeeded']
```

transition ID の最終形式、escaping、`local_id`、ObjectRef shape は V01-ADR-079 の後続修正または transition ID spec 反映で確定する。
本ADRは、その anchor として state machine semantic object を使えるようにする。

### 8. state / event は state machine 配下の local member として扱う

state machine object は、state / event / transition を member として持つ。

state / event の既存 QualifiedID policy を維持するか、state machine 配下の member ID へ寄せるかは、本ADRでは確定しない。

ただし、UI / MCP 上では state machine を root として states / events / transitions を inspect できるようにする方向を採る。

例:

```json
{
  "object": {
    "object": "node",
    "kind": "state_machine",
    "id": "order.state_machine.order_flow"
  },
  "members": {
    "states": [
      { "id": "processing", "label": "processing" }
    ],
    "events": [
      { "id": "payment_webhook_received", "label": "payment_webhook_received" }
    ],
    "transitions": [
      { "id": "order.state_machine.order_flow#processing:payment_webhook_received[payload.status == 'succeeded']" }
    ]
  }
}
```

state / event ID policy の完全な整理は、spec 反映時または別ADRで扱う。

### 9. existing `state_file` based scenario は migration 対象とする

既存の sequence scenario YAML にある `state_file:` は migration 対象とする。

旧:

```yaml
state_file: order/state.yaml
```

新:

```yaml
state_machine: order.state_machine.order_flow
```

`state_file` を compatibility field として残すか、validation warning / error にするか、migration tool を用意するかは task で扱う。

ただし、新規 authoring の正は `state_machine` とする。

## Non-goals

本ADRでは以下を決めない。

- transition ID の最終文字列表現
- transition ID の escaping policy
- state / event の QualifiedID policy を維持するか、state machine member ID へ移行するか
- existing `state_file` field の migration 手順
- `state_file` を warning にするか error にするか
- state machine YAML の最終 file classification 詳細
- state machine object の MCP response schema 全体
- UI の具体画面仕様

これらは後続ADRまたは spec / task で扱う。

## 理由

### なぜ state machine を semantic object にするか

MCP と UI を同じ semantic query surface に寄せるなら、file path ではなく semantic object を主要単位にする必要がある。

`order/state.yaml` は source location としては有用だが、人間や LLM が設計対象として扱う概念ではない。
設計対象は「注文フローの state machine」であり、それには stable semantic ID が必要である。

state machine object を導入することで、MCP / UI / render / scenario authoring が同じ semantic root を共有できる。

### なぜ `state_file` を scenario YAML から消すか

`state_file` は YAML file path reference であり、authoring YAML の時点で file layout を参照している。
これは V01-ADR-078 の MCP identity policy とは別レイヤーの問題だが、UI / MCP parity を考えると同じ方向で解決すべきである。

scenario YAML が `state_machine` semantic reference を持てば、UI でも MCP でも同じ対象を扱える。

### なぜ state machine anchor が transition ID に向くか

transition は state machine 内の edge である。
state machine object が存在するなら、transition を state machine 配下の generated object として表すのが自然である。

```text
order.state_machine.order_flow#processing:payment_webhook_received[payload.status == 'succeeded']
```

この形式では、どの state machine の transition かが明示される。
UI でも `order_flow` の中の transition として表示しやすい。

### なぜ from state anchor だけでは不足しうるか

from state anchor 案は、既存 state node QID を使えるため小さい。

```text
order.state.processing#payment_webhook_received[payload.status == 'succeeded']
```

しかし UI / MCP の semantic root としては、transition がどの state machine に属するかが見えにくい。
また、state machine object を導入するなら、transition は state machine 配下の member として扱う方が一貫する。

そのため、本ADRでは state machine object 導入を優先し、transition ID の anchor 候補を state machine QID に寄せる。

## 却下した代替案

### 代替案A: `state_file` を維持する

- 利点: 現行 scenario YAML への影響が小さい
- 欠点: YAML file path reference が authoring layer に残る。MCP / UI の semantic surface と合わない

→ 却下。新規 authoring の正は `state_machine` とする。

### 代替案B: state machine object を導入せず、from state anchor だけで transition ID を解決する

- 利点: 変更が小さい
- 欠点: UI / MCP の root object として state machine を扱えない。scenario YAML の file path reference 問題も解けない

→ 却下。state machine semantic object を導入する。

### 代替案C: sequence scenario に canonical transition ID を直接書かせる

```yaml
steps:
  - transition: order.state_machine.order_flow#processing:payment_webhook_received[payload.status == 'succeeded']
```

- 利点: 解決済み ID を直接参照できる
- 欠点: authoring YAML が長くなり、人間が書きづらい。scenario は state machine と `(from_state, via, guard?)` で書く方が読みやすい

→ 却下。scenario YAML は state machine を指定し、step は従来どおり `(from_state, via, guard?)` を使う。

### 代替案D: UI 実装だけで file path を隠す

- 利点: spec 変更を避けられる
- 欠点: MCP / YAML / UI の概念が分裂する。UI が file path を semantic label に変換する独自 layer を持つことになり、設計言語としての一貫性が落ちる

→ 却下。spec で semantic object を定義する。

## 影響

### spec への影響

本ADR受理後、以下の spec 更新が必要になる。

- `docs/spec/file-types.md`
  - `as: state_machine` file type を追加するか、state machine 定義 file の分類を定義する

- `docs/spec/nodes.md`
  - `state_machine` を semantic object として扱うことを追加する
  - state / event / transition が state machine 配下の member であることを補足する

- `docs/spec/naming.md`
  - state machine QualifiedID 形式 `<module>.state_machine.<id>` を追加する
  - `state_machine` sentinel を予約語に追加するか検討する

- `docs/spec/views/state-diagram.md`
  - render root を file ではなく state machine semantic object に更新する
  - `1 state machine = 1 State Diagram` と定義する

- `docs/spec/views/sequence-diagram.md`
  - `state_file` を `state_machine` に置き換える
  - transition 解決を `(state_machine, from_state, via, guard?)` に更新する
  - existing `state_file` migration policy への導線を追加する

- `docs/spec/mcp/schema.md`
  - state machine ObjectRef / selector support を追加する
  - transition ID anchor として state machine QID を使う方向へ接続する

- `docs/spec/mcp/tools/inspect.md`
  - `inspect(state_machine)` を追加する
  - state / event / transition members を返す shape を定義する

### V01-ADR-079 への影響

V01-ADR-079 は、本ADRの結論を受けて transition ID final policy を更新する。

具体的には、transition canonical ID は以下の形式を基本候補とする。

```text
<state-machine-qualified-id>#<from-state>:<event>[<guard>]
```

V01-ADR-079 側では、guard inclusion、old file path based ID unsupported、visibility generated などの非 file-path 制約を維持しつつ、anchor / local_id を本ADRに合わせて確定する。

### implementation への影響

実装では、state machine semantic object を ResolvedProject に持たせる必要がある。

影響が予想される領域:

- file classification / decode
- semantic build
- state / event / transition indexing
- sequence scenario resolution
- state diagram renderer
- MCP ObjectRef / selector
- inspect(state_machine)
- transition ID formatter
- migration / fixture update

内部 index key として `(stateFileID, fromStateID, eventID, guard?)` を一時的に維持してよいが、MCP / UI public contract では state machine QID を semantic root とする。

### UC / fixture への影響

既存 sequence scenario YAML の `state_file:` は `state_machine:` へ migration する必要がある。

既存 state definition YAML には `as: state_machine` / `id:` の追加、または state machine semantic object を表す別の migration が必要になる。

UC-001 / UC-002 の State Diagram / Sequence Diagram fixture は更新対象になる。

## Acceptance 前の確認事項

本ADRは proposed として起票する。
accepted に進める前に、以下を確認する。

1. `as: state_machine` file type でよいか、または `nodes:` file 内の main node として state machine を表すべきか
2. state machine QID の sentinel を `state_machine` とするか、別名にするか
3. state / event の既存 QID policy を維持するか、state machine member ID に寄せるか
4. existing `state_file:` を warning / error / compatibility field のどれとして扱うか
5. transition ID を `<state-machine-qid>#<from>:<event>[guard]` とする方針で問題ないか
6. State Diagram render output path / title / index が state machine semantic object に自然に移行できるか
7. MCP/UI の object list / inspect / navigation で state machine を first-class object として扱えるか

## Evidence

- commit: 5ae7769
- impl commit: tbd
- close boundary: M15 / `v1.1.0-spec` では follow-up scope として deferred。実装は含めない。
- 参考: V01-ADR-078 semantic anchor synthetic ID policy、V01-ADR-079 MCP transition ID non-file-path constraints、UI/MCP semantic surface parity 方針
