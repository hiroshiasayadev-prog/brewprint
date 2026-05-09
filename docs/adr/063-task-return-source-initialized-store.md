# 063: task return source への initialized source 追加と initializes の wiring source 化

- **status**: accepted
- **date**: 2026-05-05
- **supersedes**: なし（ADR-062 / ADR-061 を補追する）

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-062 で `task.returns.source` を導入し、task が外部へ返す値の source を明示する task return wiring を定義した。
ADR-062 §3 の指定可能 source は以下の3種に限定された。

| source | 意味 |
|---|---|
| node id / QualifiedID | task / join など returns を持つ node の出力全体 |
| collected asset source | `foreach.returns` で明示された collected asset |
| `$params.<name>` | main task params の `<name>` をそのまま返す |

しかし、ADR-014 で導入された `initializes` 由来の file-private store を、main task の return として返す、または flow 内の後続 task の入力として渡すユースケースが ADR-062 / ADR-060 のスコープでは扱えない。

具体的には、UC-001 の `process_report` のような構造で以下のような flow を表現したい。

```yaml
nodes:
  - id: process_report
    type: task
    main: true
    initializes:
      - name: report
        model: report
        note: "report.itemsは空リスト、report.summaryは空文字で初期化"
    params:
      - name: items
        model: item_list
    returns:
      name: report
      model: report
      source: report   # ← initializes[].name を returns.source として参照したい

  - id: append_item
    type: task
    writes: [report]
    params:
      - name: item
        model: item

flow:
  - foreach: append_item
    over: $params.items
    params:
      item: $item
```

このパターンは、task 内部で空 store を初期化 → flow で writes 蓄積 → task 完了時点でその store を return する、という composite task の自然な構造である。
ADR-062 のままでは `returns.source: report` が `unresolved_return_source` になる。

### `initializes` の設計意図再確認

ADR-014 で `initializes` を導入した際、「main node のフィールドとして持つことで、関数内変数宣言のように task 内で使う store をまとめて宣言する」という設計意図があった。
すなわち `initializes` は **「変数宣言 task の省略形」** として位置付けられている。

この意図に立ち戻ると、initialized source は task 内部の名前付き source として扱うべきであり、`returns.source` だけでなく flow 内部 wiring の bare token からも参照可能であるのが自然である。
ADR-062 / ADR-061 では initialized source を bare wiring source として扱うルールが未定義であり、cross-edge `reads` を経由しないと flow 内の step.params に initialized store の値を渡せない。

本ADRは ADR-062 を補追しつつ、initialized source を flow wiring source の正規メンバーとして加える範囲まで踏み込む。

### 概念上の懸念

「store を task return source として読む」「store を flow wiring source として読む」のは、ADR-007 / ADR-014 / ADR-020 で確立した store の概念汚染にならないか、という論点がある。

- store は ADR-007 で「runtime data instance」として定義され、cross-edge `reads` / `writes` で接続される
- ADR-014 で `initializes` 由来の store は「ファイル内 private、外部参照不可」と定義された
- ADR-020 で cross-edge は flow ステップではなく task node のフィールドに属するとされた

これらに対し、initialized source を `returns.source` および flow 内部 wiring から参照可能にすることは概念的に正当化できる。
理由は §理由 §1〜§3 で述べる。

## 決定

### 1. `returns.source` の指定可能 source に initialized source を追加する

ADR-062 §3 の指定可能 source 表に、initialized source を追加する。

| source | 意味 |
|---|---|
| node id / QualifiedID | task / join など returns を持つ node の出力全体 |
| collected asset source | `foreach.returns` で明示された collected asset |
| **initialized source** | **同一 file の main task `initializes[].name` で宣言された file-private source** |
| `$params.<name>` | main task params の `<name>` をそのまま返す |

initialized source は、bare token として `returns.source` に指定する。

```yaml
nodes:
  - id: process_report
    type: task
    main: true
    initializes:
      - name: report
        model: report
    returns:
      name: report
      model: report
      source: report   # initialized source 参照
```

`$item` を `returns.source` で禁止するルールは ADR-062 から継承する。`$item` は foreach iteration 内部の source であるため、`returns.source` で指定された場合は `invalid_return_source` を出す。

### 2. flow 内部 wiring の bare token source 種別に initialized source を追加する

ADR-061 §6 で確立された bare wiring source 名前空間（node id / collected asset source）に、initialized source を加える。
これにより、flow 内部 wiring（step.params / branch.params / fork.branches[].steps[].params / foreach.params / branch.cases[].params）の bare token source 種別は以下の3種となる。

1. node id / QualifiedID
2. collected asset source（`foreach.returns` 由来）
3. **initialized source（`initializes[].name` 由来）**

シジル付き source（`$params.<name>` / `$item`）は従来通り別系統で解決する。

```yaml
nodes:
  - id: process_report
    type: task
    main: true
    initializes:
      - name: cache
        model: result_cache

  - id: lookup_item
    type: task
    params:
      - name: cache
        model: result_cache
      - name: key
        model: str
    returns:
      name: hit
      model: cache_entry

flow:
  - step: lookup_item
    params:
      cache: cache       # bare token で initialized source を参照
      key: $params.key
```

initialized source の TypeRef は `initializes[].model` を named model TypeRef として導出する（§4 参照）。
これは ADR-060 §3 の TypeRef compatibility ルールで通常の wiring source と同様に target param と互換性検証される。

### 3. 評価時点

`returns.source` および flow 内部 wiring から initialized source を参照する場合、参照時点の source 状態は以下のとおり。

- `returns.source` から参照する場合: task 実行完了時点（flow END 時点）の store の値
- flow 内部 wiring（step.params など）から参照する場合: 当該 wiring 評価時点（その step / branch / foreach iteration が実行される直前）の store の値

`returns.source` の評価時点ルールは指定 source 種別を問わず共通である。

- node output: 当該 node の実行完了時点の output
- collected asset source: 当該 foreach 全体が完了した時点の collect 結果
- initialized source: task 実行完了時点（flow END 時点）の store の値
- `$params.<name>`: task 起動時点の入力（不変）

この時間軸の整理は ADR-062 §3「flow entry 順における前方参照という概念は適用しない」と整合する。

flow 内部 wiring から initialized source を参照する場合、その source は flow の構造上、当該 wiring に到達する時点での store の状態を指す。
具体的な mutation semantics（参照渡し / コピー / append の挙動）は task 実装に委ねられ、brewprint としては規定しない。

### 4. 型解決

initialized source の TypeRef は、`initializes[].model` を named model TypeRef として導出する。

`initializes[].model` は v1.1 でも model-id 参照のままであり、TypeRef 受け取り対象には含まれない（spec/nodes.md `init オブジェクト` 既述）。
inline `list<T>` / `dict<T>` は `initializes[].model` では受け取らない。

`returns.source` または flow 内部 wiring から参照された場合に限り、その model-id を named model TypeRef として扱い、ADR-060 §3 の TypeRef compatibility ルールで target TypeRef と互換性検証する。

```yaml
- id: process_report
  initializes:
    - name: report
      model: report   # named model 参照のみ。inline list<T>/dict<T> 不可
  returns:
    name: report
    model: report
    source: report   # source TypeRef = named model `report`
```

`initializes[].model` が解決不能な場合、initialized source の TypeRef も解決不能として扱い、ADR-060 §7 / ADR-062 §5 の二重診断抑制ルールに従って `incompatible_wiring_type` / `incompatible_return_type` を発行しない。

### 5. flow 内 writes 有無は valid 判定に影響しない

`returns.source` または flow 内部 wiring に指定した initialized store が、flow 内で `writes` されているか否かは本ADRでは validation 対象としない。

初期値のまま参照するユースケース（identity initializer / default fallback / placeholder return）が存在しうるため、writes されていない initialized source の参照は valid とする。

将来 lint レベルの diagnostic を導入する余地は残すが、本ADRでは規定しない。

### 6. 名前空間と衝突ルール

initialized source は同一 flow file 内の bare wiring source 名前空間に参加する。

ADR-061 §6 で定義された重複ルールに、initialized source も加える。
同一 flow file 内で、以下は重複してはならない。

- 同一ファイル内の node id
- `foreach.returns` で宣言された collected asset source 名
- **initialized source 名（`initializes[].name`）**

重複した場合、bare token として参照したときに解決先が一意でなくなるため `duplicate_flow_source` を出す。

なお、task の `returns.name` は通常 flow の wiring source ではないため、task `returns.name` と initialized source 名 / `foreach.returns` 名が同名でも衝突扱いしない（ADR-061 §6 の方針を継承）。

### 7. cross-edge との関係

initialized source の bare wiring source 化は、cross-edge `reads` / `writes`（ADR-020）の体系を置き換えない。
両者は分担する。

- **flow param wiring**: 値の受け渡し contract。step が必要とする param に値を供給する経路
- **cross-edge `reads` / `writes`**: 副作用 / store access contract。task が store を参照・更新することの宣言

initialized source を flow wiring source として param に渡すことと、その task が `reads` / `writes` で同じ source を読む / 更新することは、概念的には独立した宣言として併存する。

```yaml
- id: append_item
  type: task
  reads: [report]       # task が report を読むことを宣言（store access contract）
  writes: [report]      # task が report を更新することを宣言
  params:
    - name: report
      model: report
    - name: item
      model: item

flow:
  - foreach: append_item
    over: $params.items
    params:
      report: report    # initialized source を param 値として渡す（dataflow）
      item: $item
```

cross-edge は task signature レベルで「store と関わる」事実を宣言する。
flow param wiring は dataflow レベルで「具体的にどの値を渡すか」を表す。
両者は同じ store を扱っていても役割が異なる。

なお、両者の整合性検査（reads 宣言なしに wiring から参照する場合に warn する等）は本ADRでは規定しない。
将来 lint レベルで扱う余地を残す。

### 8. ADR-014 の外部参照不可ルールとの関係

ADR-014 §3 は「`initializes` で宣言された store は file-private、外部参照不可」と規定した。
本ADRはこの規則を遡及修正しない。

initialized source を `returns.source` / flow 内部 wiring から参照することは、いずれも **同一 file 内** の参照である。
外部 file からの QualifiedID 参照や module 跨ぎ参照は引き続き不可である。

`returns` を経由して store の **値** が外部に伝播するが、これは store それ自体への参照ではなく、task の return signature を満たす値として取り出されたものである。

### 9. diagnostics

本ADRで以下の診断を追加・拡張する。新規 diagnostic コードは追加しない。

| code | severity | 意味 |
|---|---|---|
| `unresolved_return_source` | error | `returns.source` が node id / `$params.<name>` / collected asset source / **initialized source** のいずれとしても解決できない（拡張） |
| `unresolved_wiring_source` | error | wiring source が node id / `$params.<name>` / `$item` / collected asset source / **initialized source** のいずれとしても解決できない（拡張） |
| `invalid_return_source` | error | source は解決できたが task return source として使えない。例: returns を持たない node、`$item`（変更なし） |
| `incompatible_return_type` | error | `returns.source` の TypeRef と `returns.model` の TypeRef が互換しない（変更なし） |
| `incompatible_wiring_type` | error | wiring source の TypeRef と target param の TypeRef が互換しない（変更なし。initialized source も対象に含む） |
| `duplicate_flow_source` | error | 同一 flow file 内で node id / `foreach.returns` / **initialized source 名** のいずれかが重複している（拡張） |

`unresolved_return_source` の意味は ADR-062 §8 から拡張される。
`unresolved_wiring_source` の意味は ADR-060 §6 / ADR-061 §9 から拡張される。
`duplicate_flow_source` の対象は ADR-061 §9 から拡張される。

`invalid_return_source` / `invalid_wiring_source` は initialized source に関しては発火しない。initialized source は valid な source 種別になる。

### 10. ADR-062 / ADR-061 / ADR-060 との関係

本ADRは ADR-062 / ADR-061 / ADR-060 を superseded しない。

- ADR-062 §3 の指定可能 source 表に initialized source を追加する補追
- ADR-061 §6 の bare wiring source 名前空間に initialized source を追加する補追
- ADR-060 §5 の wiring source 種別に initialized source を追加する補追

各 ADR 自体は accepted 状態の起票時点スナップショットとして残す。現行仕様としては spec を参照する。

`$item` を `returns.source` で禁止するルール、ADR-060 TypeRef compatibility 適用、二重診断抑制ルールはすべて元の ADR から継承する。

### 11. DAG render ルール

`returns.source` / flow 内部 wiring に initialized source を指定した場合の DAG render ルール（initializes subgraph の表示有無、return boundary node との接続表現、bare wiring からの edge 表現など）は、本ADRでは規定しない。

`returns.source` / initialized source の DAG render ルールは別 ADR（ADR-064: returns.source の DAG render ルール）で扱う。

## 理由

### なぜ initialized source を返り値・flow 内部 wiring 両方で参照可能にするか

ADR-014 起票時の設計意図は「main node のフィールドとして変数宣言タスクを省略する」である。
すなわち initialized source は task 内部の名前付き source であり、本来は通常の wiring source と同列に扱うのが意図に合う。

ADR-062 / ADR-061 / ADR-060 では initialized source を wiring source として位置付けるルールが空白だったため、`returns.source` / flow 内部 wiring から参照したくても標準経路がなく、cross-edge `reads` 経由で迂回するか dummy task を挟む必要があった。

initialized source を bare wiring source の正規メンバーに加えることで、ADR-014 の設計意図と整合し、composite task の自然な記述ができる。

### なぜ概念汚染にならないか（論理1: ライフサイクル視点）

`initializes` 由来の store は、ADR-014 で「file-private、外部参照不可」と定義された。
この store は task の中で初期化され、task の中で更新され、task の完了とともにスコープが閉じる。
ライフサイクルが task に閉じている file-private store は、task の view から見れば「task 内部の名前付き source」として扱える。

`returns.source` / flow 内部 wiring から initialized source を参照することは、task 内部での値の取得操作であり、ADR-014 の外部参照不可ルールと矛盾しない。
store それ自体は依然として file-private であり、本ADRが解禁するのは「task 内部の名前付き source として bare token で参照する」ことだけである。

### なぜ概念汚染にならないか（論理2: 粒度・役割視点）

cross-edge `reads` / `writes`（ADR-020）と initialized source の wiring 参照は、扱う粒度・役割が異なる。

- **cross-edge `reads` / `writes`**: task signature レベルで「store と関わる」事実を宣言する。task が外部副作用を持つことの contract
- **flow param wiring（initialized source 参照を含む）**: dataflow レベルで「具体的にどの値を step に渡すか」を表す
- **`returns.source` の initialized source 参照**: task 全体の return signature を満たす値として store の最終値を取り出すこと

これらは同じ store を扱っていても役割が独立している。
cross-edge は flow 内の各 task が store と関わる事実を宣言し、wiring は具体的な値の受け渡しを記述し、`returns.source` は task 完了後に外向きに何が出るかを記述する。

3者は補完関係にあり、片方が片方を置き換えるものではない。

### なぜ概念汚染にならないか（論理3: 設計意図への回帰）

`initializes` は ADR-014 起票時に「変数宣言 task の省略形」として設計された。
task 内部で名前付き source として宣言されたものを、task 内部の wiring から bare token で参照するのは、その設計意図に合致する。

「store だから cross-edge 経由でしか参照できない」という制約は、`initializes` を「store の宣言」と硬く解釈した場合に生じる派生規則であり、起票時の設計意図そのものではない。
本ADRは ADR-014 起票時の設計意図に立ち返り、initialized source を task 内部の名前付き source として扱うことを明確化する。

### なぜ writes 有無を validation 対象にしないか

初期値のまま task return として返す / flow に渡すユースケースが存在する。

- 設計途中で空 default を返す stub
- 入力がない場合の identity initializer
- placeholder として default 構造体を渡す API

これらは「writes されていないが意図的に initialized 値を使う」ケースであり、validation で error にすると合理的な設計を阻害する。

将来 lint レベルで「writes されていない initialized source を参照している」を warning として出す余地はあるが、本ADRでは規定しない。

### なぜ TypeRef を named model 参照に限定するか

`initializes[].model` は ADR-014 起票時から model-id 参照として定義され、v1.1 TypeRef 拡張対象（ADR-060 §1, §9）には含まれていない。

initialized store の型表現を inline `list<T>` / `dict<T>` に拡張するか否かは独立した設計判断であり、ADR-060 / 本ADRのスコープ外である。
将来必要になった場合は別 ADR で扱う。

本ADRでは、wiring source として参照されたときに限り named model TypeRef として扱うことで、最小限の追補に留める。

### なぜ initialized source 名前空間を bare wiring source と統合するか

initialized source は `returns.source` および flow 内部 wiring で bare token として参照される。
ADR-061 §6 で確立された bare wiring source 名前空間（node id / collected asset source）に initialized source も加えないと、同名の node / collected asset と initialized source が同居したときに解決先が決まらない。

統合により `duplicate_flow_source` で一括検出できる。
別 namespace にすると参照側で「どちらの source か」を構文上区別する記法が必要になり、brewprint の bare token 一意性が崩れる。

### なぜ cross-edge と flow wiring の整合性検査を本ADRで規定しないか

`reads` 宣言なしに initialized source を wiring から参照したり、`writes` 宣言なしに store を更新する step を書いた場合の検査は、設計の質を高める lint 対象である。

しかしこれは ADR-014 / ADR-020 の体系を跨ぐ大きい論点であり、本ADRのスコープ（initialized source の wiring source 化）を越える。
将来別 ADR で lint 方針として扱う余地を残す。

### 却下した代替案

#### 代替案A: initialized source を `returns.source` だけで参照可能にする

`returns.source` だけ解禁し、flow 内部 wiring からは引き続き cross-edge `reads` 経由を強制する案。

利点: 変更範囲が小さい。
欠点: ADR-014 起票時の「変数宣言 task の省略形」という設計意図と整合しない。flow 内で initialized store の値を step.params に渡したいたびに dummy reader task を挟む必要があり、設計表現が冗長になる。

→ 却下。設計意図への回帰を優先し、flow 内部 wiring も解禁する。

#### 代替案B: initialized source を `returns.source` / flow 内部 wiring の両方で禁止のままにする

ADR-062 / ADR-061 / ADR-060 のままで何も追加しない案。

利点: ADR 体系の変更がない。
欠点: composite task で初期化 store を return / 受け渡しする自然な構造が表現できない。flow 内に dummy pass-through / reader task を挟む迂回策が必要になる。

→ 却下。

#### 代替案C: writes されていない initialized source の参照を error にする

利点: 「使われていない初期化」を機械的に検出できる。
欠点: 初期値のまま参照する合理的なユースケース（stub / identity / placeholder）を阻害する。

→ 却下。warning レベルの lint 提案は将来検討の余地を残すが、本ADRでは error 化しない。

#### 代替案D: ADR-062 / ADR-061 / ADR-060 を遡及修正して指定可能 source 表を更新する

ADR-050 §7 / ADR-057 Non-goals が禁じる「ADR の遡及修正」に該当する。
各 ADR は accepted 状態の起票時点スナップショットとして残し、本ADRが補追する形を取る。

→ 却下。

#### 代替案E: initializes[].model 自体を v1.1 TypeRef に拡張する

ADR-060 §1 / §9 で TypeRef 受け取り対象は明示的に列挙されており、initializes[].model は含まれていない。
これを変更すると ADR-060 の決定の遡及修正に近くなる。

initialized store の型表現を拡張するか否かは独立した設計判断であり、本ADRとは別の ADR で扱う。

→ 却下。本ADRは wiring source として参照されたときの扱いだけを規定する。

#### 代替案F: initialized source を returns.source 専用 namespace に閉じる

bare wiring source 名前空間とは別に、returns.source 専用 namespace を設ける案。

利点: 「returns.source からだけ参照できる」という機能要件を最小限で実現する。
欠点: 同じ initialized source 名なのに参照できる文脈で解決ルールが異なるという、利用者にとって不自然な体系になる。ADR-014 の設計意図にも反する。

→ 却下。

## 影響

### 既存 spec への影響

- `docs/spec/nodes.md` の `returns オブジェクト` の `source` 説明に initialized source を追加する
- `docs/spec/nodes.md` の `init オブジェクト` 末尾に「`initializes[].name` は同一 file の bare wiring source として参照可能。`returns.source` および flow 内部 wiring の両方から参照可。TypeRef は `initializes[].model` を named model TypeRef として扱う」旨を追記する
- `docs/spec/edges.md` §1-7 の wiring source 解決表に initialized source を追加し、5種目の bare wiring source 種別として記述する
- `docs/spec/edges.md` §1-7 の名前空間・重複ルール説明に initialized source を加える
- `docs/spec/edges.md` §1-8 task return wiring の指定可能 source 表に initialized source を追加する
- `docs/spec/edges.md` §1-8 に「`returns.source` は task 実行完了時点（flow END 時点）の source を参照する」旨を追加する
- `docs/spec/diagnostics.md` の `unresolved_return_source` / `unresolved_wiring_source` / `duplicate_flow_source` の意味を拡張（initialized source を含める）

### 既存 ADR への影響

- ADR-062 §3 の指定可能 source 表は補追される。ADR-062 自体は遡及修正しない（status: accepted のまま）
- ADR-061 §6 の重複ルール、§3 の wiring source 種別は本ADR §6 / §2 で initialized source も含めるよう拡張される。ADR-061 自体は遡及修正しない
- ADR-060 §5 の wiring source 種別は本ADR §2 で initialized source を5種目として加えるよう拡張される。ADR-060 自体は遡及修正しない
- ADR-014 の外部参照不可ルールは維持される。本ADR §8 で関係を整理する
- ADR-020 の cross-edge ルールは維持される。本ADR §7 で役割分担を整理する

### 既存実装への影響

- `internal/rawyaml` / `internal/semantic` で `returns.source` および flow wiring source の解決対象に initialized source を追加する
- task return source resolver / flow wiring source resolver の両方に initialized source 解決を追加する
  - 同一 file の main task の `initializes[]` を引き、`name == <bare_token>` の entry の `model` を named model TypeRef として返す
- `duplicate_flow_source` の検出対象に initialized source 名を加える
- `unresolved_return_source` / `unresolved_wiring_source` の解決失敗判定で initialized source も探索する
- TypeRef compatibility は ADR-060 ルールを継承する（initialized source も対象）
- writes 有無の検査は実装しない
- cross-edge `reads` / `writes` 宣言と flow wiring 参照の整合性検査は実装しない

### 既存 UC への影響

- UC-001 で initialized store を `returns.source` / flow 内 step.params に渡す記述が valid になる
- 既存 wiring の回帰は発生しない見込み（initialized source を bare token で参照している UC は現状存在しない）

### v1.1 への影響

本ADRは ADR-062 / ADR-061 / ADR-060 / M15 Phase B2 の補追であり、v1.1.0-spec の凍結対象に含める。
v1.0.0-spec の遡及修正ではない forward 拡張である。

### 後続 ADR への影響

- DAG render ルール（initializes subgraph と returns boundary node の接続、initialized source 参照時の edge 表現）は別 ADR（ADR-064）で扱う

## Evidence

- commit: ee0a48c
- impl commit: e7b8292
- 参考: ADR-014 initializes field, ADR-020 cross-edge management, ADR-060 TypeRef / flow wiring type compatibility, ADR-061 foreach collected asset, ADR-062 task return source
