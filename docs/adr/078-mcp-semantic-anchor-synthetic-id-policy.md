# 078: MCP synthetic ID を semantic anchor based に見直す

- **status**: accepted
- **date**: 2026-05-13
- **depends_on**: ADR-047, ADR-048, ADR-049, ADR-054, ADR-058, ADR-070

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

brewprint MCP は、Raw YAML AST ではなく ResolvedProject 上の semantic object を LLM へ公開する query layer である。

しかし、現行 `docs/spec/mcp/schema.md` では、QualifiedID を持たない file-local object の synthetic ID として、以下のような file path anchor が使われている。

```text
order/task/checkout.yaml#build_order
order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']
```

これは、ADR-058 で private sub node を project-wide QualifiedID から外し、file-private object として扱う方針を補強した流れから生じた表現である。
ADR-058 の主眼は、private sub node を外部参照可能な QualifiedID 一意性制約から外し、同一 file 内だけで解決できるようにすることだった。

一方で、MCP public contract として file path を object identity の主成分に出すと、以下の問題がある。

- MCP 利用者が semantic object ではなく YAML file layout を意識する必要がある
- `module.kind.object` 形式の public QualifiedID と、file path based synthetic ID の見た目が大きく乖離する
- LLM が ID と source path を混同しやすい
- file layout が query contract に漏れ、将来の project layout 変更や file migration の影響を受けやすくなる
- file-private helper model の MCP exposure を決める際にも、`<file-id>#<local-id>` を流用すべきか、semantic anchor を使うべきかが未決になる

ADR-070 では file-private helper model を導入し、MCP response 内では private sub node と同じく synthetic ID を使う方針を記録した。
しかし、通常の listing に helper model を含めるか、`include_private` のような option を持つかは未決であり、さらに synthetic ID の anchor を file path にするか semantic object にするかも未整理である。

Codex review では、ADR-078 の大方針は妥当だが、transition ID は現行 spec / implementation / tests に深く file path based ID として組み込まれており、ADR-078 に含めて acceptance するには未確定要素が大きすぎると指摘された。
また、`selector.file + local_id`、ObjectRef metadata、`list_objects(include_private)`、file-private helper model の未実装状態も、総論と詳細仕様を分ける必要がある論点として整理された。

このため、本ADRは MCP synthetic ID policy の総論を定める proposed ADR とする。
transition ID の最終 policy や helper model の具体 exposure schema は、本ADRでは確定しない。

## 決定

### 1. MCP public contract では semantic object identity を優先する

MCP は Raw YAML AST や file layout を直接操作する API ではなく、ResolvedProject 上の semantic object query layer である。

したがって、MCP response / selector の object identity は、可能な限り semantic ID を主とする。
YAML path は identity の主成分ではなく、source metadata として扱う。

```json
{
  "id": "order.task.checkout#build_order",
  "source": {
    "file": "order/task/checkout.yaml"
  }
}
```

上記のように、query 用 ID は semantic anchor based にし、file path は source 情報へ閉じ込める方向を基本方針とする。

### 2. file-local / generated object の synthetic ID は semantic anchor based を原則とする

QualifiedID を持たない private / generated object には、MCP response 上の安定参照として synthetic ID を使う。

ただし、synthetic ID の anchor は file path ではなく、可能な限り semantic anchor object の ID とする。

基本形:

```text
<semantic-anchor-id>#<local-id>
```

例:

```text
order.task.checkout#build_order
mcp.model.get_reference_tree_response#reference_tree_node
auth.task.login#auth_token
```

`.` は public semantic namespace を表し、`#` は anchor 配下の private / generated namespace を表す。

この ID は public QualifiedID ではない。
外部 YAML から参照可能であることも意味しない。
MCP query layer 上で file-local / generated object を安定して識別するための synthetic ID である。

### 3. `source.file` と object identity を分離する

MCP response は source location を返してよい。
しかし、source file path は object identity とは別の情報として扱う。

ObjectRef では以下の分離を基本とする。

```json
{
  "object": "node",
  "kind": "task",
  "id": "order.task.checkout#build_order",
  "anchor": "order.task.checkout",
  "local_id": "build_order",
  "visibility": "file-private",
  "source": {
    "file": "order/task/checkout.yaml"
  }
}
```

`file` / `source.file` は、定義元の確認、diagnostic、source snippet 取得、source-oriented query、debug / migration に使う。
一方、通常の MCP query / reference / traversal では `id` を semantic anchor based にする。

### 4. primary selector は `selector.id` のみを canonical path とする

MCP tool の primary selector は `selector.id` のみを canonical path とする。

MCP response の `ObjectRef.id` は、後続 tool の `selector.id` としてそのまま利用できる stable canonical ID でなければならない。
LLM / MCP client / UI は、ObjectRef から受け取った `id` をそのまま次の query に渡せることを期待できる。

```json
{
  "selector": {
    "id": "order.task.checkout#build_order"
  }
}
```

private / generated object を直接 query する場合も、主導線は semantic anchor based synthetic ID を `selector.id` に指定する形とする。

primary selector では `selector.object` を要求しない。
object kind は resolver が canonical ID から解決し、response の `ObjectRef.object` / `ObjectRef.kind` として返す。

`selector.object` は primary selector から除外する。
object kind assertion が必要な場合は、debug / validation 用 selector として別枠で扱うか、後続 spec で明示する。
ただし、それは primary query path ではない。

`selector.file + local_id` は source-oriented compatibility selector として残す。
これは、source-driven inspect、debug、migration、または `inspect(file)` からの follow-up query のために許容する補助経路である。

```json
{
  "selector": {
    "file": "order/task/checkout.yaml",
    "local_id": "build_order"
  }
}
```

`selector.file + local_id` は canonical form ではない。
MCP response 内で object を返す場合は、可能な限り canonical `id` を返す。

field も primary selector では canonical field ID で問い合わせる方向に寄せる。

```json
{
  "selector": {
    "id": "order.model.order.id"
  }
}
```

親 model ID + `local_id` による field selector は compatibility / source-oriented selector として扱い、primary selector にはしない。
private helper model field の canonical ID 形式は後続ADRまたは spec task で扱う。

### 5. ObjectRef に `anchor` と `visibility` を追加する方向で確定する

semantic anchor based synthetic ID を導入する場合、ID文字列だけでは、その object が public QualifiedID を持つ object なのか、anchor 配下の private / generated object なのかを機械的に区別しにくい。

そのため、MCP ObjectRef には以下の metadata を追加する方向で確定する。

| field | 内容 |
|---|---|
| `anchor` | private / generated object が属する semantic anchor object の ID |
| `visibility` | `public` / `file-private` / `generated` |

`visibility` の意味は以下とする。

| visibility | 意味 |
|---|---|
| `public` | public QualifiedID を持ち、外部 YAML から参照可能な semantic object |
| `file-private` | 同一 file 内だけで参照可能な private object。MCP query 可能だが public QualifiedID は持たない |
| `generated` | YAML上の独立 node ではなく、ResolvedProject 上で生成される queryable object。例: implicit asset |

`qualified_id` は public object のみに返す。
file-private / generated object では `qualified_id` を返さない。

ObjectRef metadata の requiredness は以下を基本方針とする。

- `visibility` は全 ObjectRef で必須とする
- `anchor` は `file-private` / `generated` object で必須とする
- `anchor` は `public` object では省略する

理由:

- `visibility` は object が public surface なのか、file-private なのか、generated なのかを tool / LLM が文字列 parse なしに判断するために必要である
- `visibility` を全 ObjectRef で必須にすると、`visibility` 欠落時は public とみなす、という暗黙ルールを避けられる
- `anchor` は `#` の左側を構造化した metadata であり、private / generated object の所属先を明示するために必要である
- public object はそれ自体が semantic anchor になれるため、`anchor` を重複して持つ必要はない

既存 `file` / `local_id` との共存、response schema の細部は MCP spec 反映時に確定する。

### 6. file-private helper model は本方針に従うが、具体 exposure schema は後続で扱う

ADR-070 の file-private helper model 方針は維持する。

ただし、現行 MCP spec / Go implementation には file-private helper model の exposure schema がまだ反映されていない。
そのため、本ADRでは helper model の具体的な listing / inspect / signature / references response schema は確定しない。

本ADRで決めるのは、file-private helper model が MCP に露出される場合、その ID は file path based ではなく semantic anchor based にする、という総論である。

model file 内 helper model:

```text
mcp.model.get_reference_tree_response#reference_tree_node
```

ここで `mcp.model.get_reference_tree_response` は public main model の QualifiedID、`reference_tree_node` は file-private helper model の local ID である。

同様に task file 内 helper model は、main task を anchor とする。

```text
mcp.task.analyze_impact#impact_entry
```

file-private helper model は public QualifiedID を持たない。
`qualified_id` は返さず、ObjectRef には `anchor` と `visibility: "file-private"` を含める方向とする。

### 7. `list_objects` は default public-only、`include_private` は private node / helper model に限定する

`list_objects` は project 内 semantic object の探索起点である。

file-private helper model や private sub node を default listing に混ぜると、LLM がそれらを外部参照可能な public object と誤認しやすい。

したがって、`list_objects` の default は public surface のみとする。
private object を listing に含める場合は `include_private: true` を明示する。

`include_private: true` の初期対象は以下に限定する。

- private sub node
- file-private helper model

以下は `include_private: true` の初期対象に含めない。

- transition
- asset
- field
- scenario step

これらは `include_private: true` の対象外であるだけでなく、初期 `list_objects` の default listing 対象にも含めない。
独立した探索起点ではなく、`inspect` / `get_references` / `get_reference_tree` から到達する queryable object として扱う。
listing に含めるかどうかは、必要が出た時点で別途拡張する。

### 8. transition ID の最終 policy は本ADRでは決めない

transition は本ADRの semantic anchor based synthetic ID 原則に関係するが、最終 ID policy は本ADRでは決定しない。

現行 spec / implementation / tests では、transition ID は以下の file path based ID として深く使われている。

```text
order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']
```

これを semantic anchor based にするには、少なくとも以下を決める必要がある。

- state file 自体を semantic object として扱うか
- state machine object を新設するか
- module-level state namespace を anchor にするか
- module 内に複数 state file が存在しうる場合の一意性をどう扱うか
- guard 文字列を ID に含め続けるか
- `state_file` を identity として扱うのか、source metadata として扱うのか

これらは transition / state model の identity policy そのものであり、本ADRに含めるとスコープが大きくなりすぎる。

したがって、transition ID policy は後続ADRで扱う。
本ADRでは、transition ID を semantic-anchor synthetic ID policy の既知の例外 / unresolved scope として明示する。

### 9. 本ADRは proposed とし、spec 反映前に二段階レビューを必須とする

本ADRは影響範囲が広いため、accepted に進める前に以下のレビューを行う。

1. Codex review
   - 現行実装 / spec / tests / MCP examples への影響範囲を洗い出す
   - file path based ID が残っている箇所を列挙する
   - semantic anchor based ID へ移行した場合の破壊範囲を確認する
   - transition ID を本ADRから分割すべきか確認する

2. Opus review
   - Codex review の影響範囲が漏れていないか確認する
   - semantic anchor の定義が task / model / asset / field / view に対して一貫するか確認する
   - transition ID policy を別ADRに分割する前提で、ADR-078 が総論として十分か確認する
   - spec 更新案として十分か、さらにADRを分割すべきか判断する

レビュー完了後、MCP spec の具体更新内容を確定する。

## Non-goals

本ADRでは以下を決めない。

- transition ID の最終形式
- state file / state machine の semantic anchor policy
- file-private helper model の具体的な MCP response schema
- field of private helper model の ID 形式
- parent object ID + `local_id` 型 selector の compatibility scope
- private sub task が生成した asset の producer ID 形式
- scenario step の selector / ID policy
- transition / field / scenario step の `visibility` 値の最終分類
- 既存実装の migration 手順
- backward compatibility の具体的な adapter / alias 実装

これらは後続ADRまたは MCP spec 反映時に扱う。

## 理由

### なぜ file path を MCP ID の主成分にしないか

MCP query layer の利用者が欲しいのは、file path ではなく semantic object である。

file path は source location として重要だが、object identity として前面に出すと、MCP が YAML file browser のように見えてしまう。
これは、MCP が Raw YAML AST ではなく semantic object query layer であるという ADR-047 / MCP spec の基本方針と相性が悪い。

また、file path based ID は project layout に強く依存する。
将来 file placement や module grouping を変えると、semantic object が同じでも MCP ID が変わりやすい。

### なぜ semantic anchor + `#local` か

public QualifiedID は `module.kind.object` 形式であり、人間にも LLM にも semantic object として読みやすい。

file-local object は public QualifiedID を持たないが、多くの場合は public main node や producer object に従属している。
そのため、以下のように public semantic anchor の下に置くと、public / private の境界が読みやすい。

```text
order.task.checkout#build_order
mcp.model.get_reference_tree_response#reference_tree_node
```

`.` は外部 semantic namespace、`#` は anchor 配下の local namespace という読み分けができる。

これは file path based ID よりも MCP query layer の語彙として自然である。

### なぜ primary selector を `selector.id` のみにするか

MCP response の `ObjectRef.id` をそのまま後続 tool の `selector.id` に渡せることは、LLM / MCP client / UI にとって最も単純で誤りにくい query path である。

primary selector に `selector.object` を含めると、利用者は各 query のたびに object kind を判断して埋める必要がある。
これは、実装側の解決都合を LLM / UI 側へ押し付けることになり、余計な推論負荷と誤指定の原因になる。

そのため、canonical object ID が `selector.id` に指定された場合、resolver は `selector.object` なしで対象 object を解決できなければならない。
object kind は resolver が canonical ID から解決し、response の `ObjectRef.object` / `ObjectRef.kind` として返す。

現行 implementation / tests では、private sub node の lookup に `selector.file + local_id` と file path based synthetic ID が使われている。
これを即座に削ると、既存 MCP clients / tests / debug workflow を壊す可能性がある。

一方で、MCP public contract の主導線を file path based selector や parent + local_id selector にすると、semantic query layer としての抽象化が弱くなる。

そのため、`selector.id` を唯一の primary selector とし、`selector.file + local_id` や parent + `local_id` 型 selector は source-oriented compatibility selector として残す。
これにより、semantic ID policy へ移行しつつ、source-driven query の実用性も維持できる。

### なぜ `anchor` / `visibility` が必要か

`order.task.checkout#build_order` のような ID は、人間には private object だと読める。
しかし、機械的には ID 文字列を parse しない限り、その object が public なのか private なのか、anchor が何なのかを判断しにくい。

MCP response は LLM / tool が読む public contract であるため、ID 文字列だけに意味を詰め込まない。
`anchor` / `visibility` を持たせることで、private / generated object の性質を構造的に表現できる。

### なぜ `list_objects(include_private)` の対象を絞るか

`list_objects` は探索用 tool であり、何でも列挙すると public surface と internal detail の境界が曖昧になる。

private sub node と file-private helper model は、設計対話や実装レビューで直接探索したい需要がある。
一方、transition / asset / field / scenario step は、通常は親 object の inspect や reference traversal から到達する方が自然である。

そのため、初期 `include_private` は private sub node / helper model に限定し、その他は必要が出た時点で拡張する。

### なぜ transition ID を分割するか

transition ID は、単なる private object ID ではなく state model の identity policy と密接に関わる。

現行 ID は state file path、from state、event、guard exact match に依存している。
これを semantic anchor based にするには、state file や state machine の semantic identity を先に決める必要がある。

したがって、transition ID を ADR-078 に含めて確定しようとすると、MCP synthetic ID の総論が transition / state 設計全体に引きずられる。
本ADRでは transition を unresolved scope として切り出し、後続ADRで扱う。

### なぜ ADR-058 を直接修正しないか

ADR-058 は起票時点で、private sub node を project-wide QualifiedID から外し、file-private scope を実装で保証するための判断を記録したものである。

ADR-058 の本文はスナップショットとして残す。
本ADRは、ADR-058 の private object 方針を否定するものではなく、MCP public contract に出す synthetic ID の anchor を file path から semantic object へ寄せる補強・再整理である。

## 却下した代替案

### 代替案A: `<file-id>#<local-id>` を MCP public contract の主IDとして維持する

- 利点: 現行 spec / 実装に近く、file-local uniqueness を説明しやすい
- 欠点: MCP 利用者に YAML file layout を意識させる。semantic query layer としての抽象化が弱くなる。file migration に弱い

→ 却下。少なくとも主IDとしては使わず、source / compatibility selector へ下げる。

### 代替案B: private object を MCP から直接 query できないようにする

- 利点: public surface は単純になる
- 欠点: render / inspect / implementation review で private sub node や helper model の詳細を取得できない。ADR-054 の設計対話 coverage と ADR-070 の helper model query 需要に合わない

→ 却下。private object は外部 YAML 参照不可でも、MCP query layer では inspect 可能にする必要がある。

### 代替案C: private helper model だけ semantic anchor based にし、private sub node は file path based のままにする

- 利点: ADR-070 の未決だけを小さく解ける
- 欠点: private object の ID policy が model と task で分裂する。LLM が ID 体系を覚えにくくなる。MCP schema の一貫性が壊れる

→ 却下。helper model の問題は MCP synthetic ID policy 全体として扱う。

### 代替案D: すべての private object に public QualifiedID を与える

- 利点: ID体系は単純になる
- 欠点: private object が外部参照可能に見える。ADR-011 / ADR-058 / ADR-070 の file-private 方針に反する。public surface が過剰に広がる

→ 却下。private object は public QualifiedID を持たない。

### 代替案E: transition ID も ADR-078 で確定する

- 利点: MCP synthetic ID policy を1本のADRで完結できる
- 欠点: state file / state machine / guard exact match / transition uniqueness まで同時に決める必要があり、ADR-078 のスコープを超える。現行実装・テストへの破壊範囲も大きい

→ 却下。transition ID は後続ADRで扱う。

## 影響

### MCP spec への影響

本ADRが accepted になった場合、少なくとも以下の spec 更新が必要になる。

- `docs/spec/mcp/schema.md`
  - Object selector の `id` / `file` / `local_id` の責務整理
  - primary selector は `selector.id` のみを canonical path とすることを定義
  - MCP response の `ObjectRef.id` は後続 tool の `selector.id` としてそのまま利用できる stable canonical ID であることを定義
  - primary selector では `selector.object` を要求しないことを定義
  - `selector.file + local_id` および parent object ID + `local_id` 型 selector を source-oriented compatibility selector として定義
  - field も primary selector では canonical field ID で問い合わせる方向に整理する
  - Synthetic ID を semantic anchor based に再定義
  - ObjectRef に `anchor` / `visibility` metadata を追加し、`visibility` を全 ObjectRef 必須、`anchor` を file-private / generated object で必須、public object では省略、`qualified_id` を public object のみに返すことを定義する
  - transition / field / scenario step の `visibility` 値は、本ADRでは確定せず後続ADRまたは spec task で扱うことを明記する
  - FileID を identity ではなく source / file object / compatibility selector として整理
  - transition ID は後続ADRまで unresolved / existing policy として扱う

- `docs/spec/mcp/tools/list-objects.md`
  - `include_private` option を追加
  - default は public-only と定義
  - `include_private: true` の初期対象を private sub node / file-private helper model に限定
  - transition / asset / field / scenario step は初期 listing 対象外と定義

- `docs/spec/mcp/tools/inspect.md`
  - `members.sub_tasks[]` の ID examples を semantic anchor based に更新
  - `inspect(file)` が返す file-local object と canonical ID の関係を定義
  - `inspect(model)` が file-private helper model を扱う場合の shape を後続仕様に接続
  - transition inspect examples は後続ADRまで現行 policy または unresolved として扱う

- `docs/spec/mcp/tools/get-signature.md`
  - private sub node / field / asset selector examples の確認と更新
  - transition selector examples は後続ADRの対象として分離

- `docs/spec/mcp/tools/get-references.md`
  - Reference.from / Reference.to に現れる private object ID の更新

- `docs/spec/mcp/tools/get-reference-tree.md`
  - traversal root / reached nodes の private object ID 方針を更新

- `docs/spec/mcp/tools/analyze-impact.md`
  - target selector / impacted object ID に private object が現れる場合の扱いを更新
  - transition impact は後続ADRに依存する箇所として整理

### naming spec への影響

`docs/spec/naming.md` は YAML authoring と semantic model の名前解決を主に扱う。

MCP synthetic ID は YAML から書ける QualifiedID ではないため、naming spec に直接混ぜすぎるべきではない。
ただし、public QualifiedID と MCP synthetic ID の境界を補足する必要がある可能性がある。

特に以下を整理する必要がある。

- public QualifiedID は外部参照可能 object の ID である
- MCP synthetic ID は query layer 上の識別子であり、YAML 上の外部参照可能性を意味しない
- `#` suffix は MCP query layer の private / generated object namespace であり、YAML QualifiedID grammar ではない

### ADR-058 への影響

ADR-058 の private sub node scope 方針は維持する。

ただし、ADR-058 由来で MCP spec に導入された `<file-id>#<local-id>` 形式は、本ADRにより見直し対象となる。
ADR-058 本文は遡及修正しない。
必要であれば、ADR-058 の後続補強ADRとして本ADRを参照し、spec 側で現行仕様を更新する。

### ADR-070 への影響

ADR-070 の file-private helper model 方針は維持する。

MCP listing 方針は本ADRで以下のように確定する。

- `list_objects` の default は public-only とする
- private object を listing に含める場合は `include_private: true` を使う
- `include_private: true` の初期対象は private sub node / file-private helper model に限定する

ただし、file-private helper model の具体的な response schema、`inspect(file)` / `inspect(model)` / `get_signature` / `get_references` での詳細な返却形は、helper model の spec 反映時に確定する。

### 後続ADRへの影響

本ADRから、少なくとも以下の後続ADRが必要になる。

- MCP transition ID policy
  - state file / state machine / module-level state namespace の semantic anchor を決める
  - guard exact match を ID に含めるか決める
  - transition ID と `state_file` source metadata の境界を決める

必要に応じて、以下も別ADRまたは spec task として扱う。

- field of private helper model の ID policy
- private sub task produced asset の producer ID policy
- scenario step selector / ID policy

### 実装への影響

実装では、file-local object の内部識別に `(FileID, localID)` を使い続けてよい。
本ADRが扱うのは MCP public contract 上の ID policy であり、内部 index key を必ず semantic anchor based にすることを要求しない。

ただし、QueryService / MCP adapter は、内部 key から MCP response ID を生成する際に semantic anchor based ID を返す必要がある。

影響が予想される実装領域:

- ObjectRef builder
- selector resolver
- private sub node lookup
- private sub node synthetic ID parser / formatter
- asset lookup
- field selector resolver
- reference index serialization
- MCP tool response examples / golden tests
- `list_objects` request / response schema

transition lookup / transition ID parser / transition reference comparison は後続ADRの影響範囲として分離する。

具体的な作業項目は task file で管理する。

### UC / fixture への影響

MCP response golden や UC-002 self-hosting の MCP contract YAML に、file path based synthetic ID が含まれている場合は更新が必要になる可能性がある。

ただし、fixture migration の範囲と順序は本ADRでは管理しない。
Codex review / Opus review 後、M15 または後続 milestone の task file で追跡する。

## Acceptance 前の確認事項

本ADRは proposed として起票する。
accepted に進める前に、以下を確認する。

1. 現行 docs/spec/mcp 以下で `<file-id>#<local-id>` または file path based ID が使われている箇所を列挙する
2. 現行実装で private sub node / asset / field / view object の ID がどこで生成・parseされているか確認する
3. private sub node / helper model / asset / field は semantic anchor based ID へ自然に移行できるか確認する
4. `selector.file + local_id` を source-oriented compatibility selector として残す方針で問題ないか確認する
5. `ObjectRef.anchor` / `ObjectRef.visibility` の追加で既存 response schema と矛盾しないか確認する
6. `list_objects(include_private)` の option 名・default・対象範囲が private sub node / helper model に限定されていて十分か確認する
7. backward compatibility を考慮する必要があるか確認する
8. transition ID policy を後続ADRへ分割する前提で、ADR-078 が総論ADRとして十分に閉じているか確認する

transition ID の最終形は、本ADRの acceptance 条件には含めない。
ただし、後続ADRが必要であることは acceptance 前に確認する。

## Evidence

- commit: 5ae7769
- impl commit: tbd
- close boundary: M15 / `v1.1.0-spec` では follow-up scope として deferred。実装は含めない。
- 参考: ADR-047 MCP query layer boundary、ADR-048 ResolvedProject index、ADR-049 MCP reference vocabulary、ADR-054 design conversation coverage、ADR-058 private sub node scope、ADR-070 file-private helper model、Codex review of ADR-078
