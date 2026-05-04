# 061: foreach.returns collected asset 参照ルール

- **status**: accepted
- **date**: 2026-05-04

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-016 で `foreach` は node type ではなく `flow:` セクションの制御構文として定義された。
`foreach` は apply 先 task を繰り返し実行し、各 iteration の入力は `$item` で渡す。

ADR-060 では TypeRef と flow wiring type compatibility を導入し、`$item` の型を `foreach.over` の `list<T>` から導出するルールを定めた。
ただし、ADR-060 は `foreach.returns` によって collect された asset を後続 flow から参照する場合の source id 解決ルールを ADR-061 に委ねていた。

この未決領域を残すと、以下のような自然な flow が仕様上あいまいになる。

```yaml
flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items

  - step: summarize
    params:
      items: validated_items
```

また、同じ apply 先 task を複数の入力集合に対して使う場合、apply 先 task の `returns.name` だけでは collect 結果を区別できない。

```yaml
flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      item: $item
    returns: validated_cart_items

  - foreach: validate_item
    over: $params.wishlist_items
    params:
      item: $item
    returns: validated_wishlist_items
```

したがって、`foreach.returns` を foreach invocation 単位の collected asset 名として定義し、後続 flow からの参照ルールを明確化する。

## 決定

### 1. `foreach.returns` は collected asset source 名である

`foreach.returns` は、apply 先 task の `returns` を iteration ごとに集めた collected asset に付ける file-local source 名とする。

```yaml
flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items
```

上記の `validated_items` は、`validate_item` の各 iteration の返り値を collect した asset source である。
これは apply 先 task の `returns.name` とは別の名前であり、foreach invocation 単位で命名される。

### 2. `foreach.returns` は任意である

`foreach.returns` は optional とする。

- collect 結果を後続 flow から安定参照する場合は指定する
- collect 結果を task return source として使う場合は指定する
- side-effect 目的で apply 先 task を繰り返すだけの場合は省略できる

```yaml
# side-effect only: collect result を使わない
flow:
  - foreach: sync_item
    over: $params.items
    params:
      item: $item
```

apply 先 task に `returns` がないにもかかわらず `foreach.returns` を指定した場合は invalid とする。

`foreach.returns` を省略した場合、collected asset は semantic model 上に生成されない。
foreach は side-effect / 個別 iteration の効果のみを持ち、collect 結果を表現したい場合は `foreach.returns` を明示する必要がある。

renderer / inspect / MCP は、`foreach.returns` を省略した foreach を「side-effect only な繰り返し実行」として扱う。
省略時に collected asset の内部名や擬似 source を露出してはならない。

### 3. 後続 flow は `foreach.returns` 名を wiring source として参照できる

`foreach.returns` で宣言された collected asset は、同一 flow file 内の後続 step / branch / fork / foreach から wiring source として参照できる。

```yaml
flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items

  - step: summarize_cart
    params:
      items: validated_items
```

この参照は file-local source 参照であり、`$params` / `$item` のようなシジルは使わない。

`foreach.returns` で宣言された collected asset は、その foreach entry が完了した後に生成される。
したがって、その foreach 自身の `params` 内から自分自身の `returns` 名を参照してはならない。
参照できるのは、当該 foreach entry より後ろの flow entry からのみである。

### 4. `foreach.returns` の型は `list<T>` とする

apply 先 task の `returns.model` が TypeRef `T` の場合、`foreach.returns` の collected asset 型は `list<T>` とする。

```yaml
nodes:
  - id: validate_item
    type: task
    params:
      - name: cart_item
        model: cart_item
    returns:
      name: validated
      model: cart_item

flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items
```

上記では `validated_items` の型は `list<cart_item>` である。

apply 先 task の `returns.model` が `any` の場合、`foreach.returns` の型は `list<any>` とする。

apply 先 task の `returns.model` が解決不能な場合、`foreach.returns` の TypeRef も解決不能として扱う。
この場合、後続 wiring が当該 collected asset を参照していても、`incompatible_wiring_type` は発行しない。
未解決 TypeRef に対する一次診断を優先し、二重診断を避ける。

### 5. task return source との関係

本 ADR は、`foreach.returns` によって flow 内に collected asset source を作るルールを定める。
main task / composite task がどの flow source を返すか、つまり task return source の明示方法は本 ADR では定義しない。

`foreach.returns` は、後続 ADR で定義する task return source から参照可能な flow source になりうる。
ただし、`main task returns.name` と `foreach.returns` の名前一致による暗黙接続は、本 ADR では採用しない。

```yaml
# task return source の明示方法は ADR-062 で扱う想定
nodes:
  - id: validate_cart
    type: task
    main: true
    params:
      - name: cart_items
        model: cart_item_list
    returns:
      name: validated_items
      model: cart_item_list
      source: validated_items

flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items
```

上記の `returns.source` は ADR-062 で扱う予定の例であり、本 ADR の決定範囲には含めない。

### 6. `foreach.returns` の名前空間と重複ルール

`foreach.returns` は同一 flow file 内の wiring source 名前空間に参加する。
bare token として参照される source は単一名前空間で解決され、衝突は invalid とする。

同一 flow file 内で、`foreach.returns` は以下と重複してはならない。

- 同一ファイル内の node id
- 他の `foreach.returns`

重複した場合、bare token (`some_name`) を後続 wiring から書いたときに「node output」と「collected asset source」のどちらに解決すべきかが決められないため invalid とする。

```yaml
nodes:
  - id: validated_items
    type: task
    returns:
      name: result
      model: cart_item_list

flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      item: $item
    returns: validated_items # invalid: 同一ファイル内 node id と重複
```

```yaml
flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      item: $item
    returns: validated_items

  - foreach: validate_item
    over: $params.wishlist_items
    params:
      item: $item
    returns: validated_items # invalid: collected asset source 名が重複
```

一方、task の `returns.name` は通常 flow の wiring source ではない。
通常 task の出力を参照する場合は task の node id を使うため、task の `returns.name` と `foreach.returns` が同名であることは衝突とは扱わない。

```yaml
nodes:
  - id: validate_item
    type: task
    returns:
      name: validated_items   # task signature 上の returns.name
      model: cart_item

flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items   # 衝突しない (task.returns.name は wiring source ではない)
```

なお、file / module / private scope を超えた名前空間分離（たとえば sub-task file 境界での衝突許容）は本 ADR の対象外であり、ADR-002 / ADR-003 / ADR-058 の名前空間ルールに従う。
本 ADR の重複ルールは「同一 flow file 内で bare token として参照される wiring source の解決」に限定される。

### 7. `foreach.id` は導入しない

本 ADR では `foreach.id` フィールドを導入しない。

同じ apply 先 task を複数回 foreach する場合は、`foreach.returns` に異なる collected asset 名を与えることで区別する。

```yaml
flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      item: $item
    returns: validated_cart_items

  - foreach: validate_item
    over: $params.wishlist_items
    params:
      item: $item
    returns: validated_wishlist_items
```

`foreach.id` を導入すると構文が重くなり、現時点の要件に対して過剰である。

### 8. 制御フロースコープとの関係

ADR-023 は、制御フロー構文の内部で生成された asset をスコープ外から直接参照不可とした。

`foreach.returns` はこのルールに対する明示的な escape hatch とする。
apply 先 task の iteration 内部で生成される個別 asset は外部参照不可のままだが、`foreach.returns` で明示的に collect した結果だけを flow 外側の source として公開できる。

### 9. diagnostics

本 ADR 受理後、少なくとも以下の診断を spec / 実装に追加する。

| code | severity | 意味 |
|---|---|---|
| `duplicate_flow_source` | error | 同一 flow file 内で `foreach.returns` が node id または他の `foreach.returns` と重複している |
| `invalid_foreach_returns` | error | apply 先 task に `returns` がないのに `foreach.returns` が指定されている、または当該 foreach 自身の `params` 内から自分の `returns` 名を参照している |
| `unresolved_wiring_source` | error | wiring source が node id / `$params` / `$item` / collected asset のいずれとしても解決できない |
| `invalid_wiring_source` | error | source は解決できたが、その文脈では wiring source として使えない |

`unresolved_wiring_source` と `invalid_wiring_source` は区別する。
前者は typo などにより参照先が存在しない場合、後者は参照先は存在するが returns 相当の出力を持たない node や有効範囲外の `$item` など、source として不正な場合に使う。

既存 diagnostic code へ統合する場合でも、上記の状態を機械的に区別できる形にする。

## 理由

### なぜ `foreach.returns` が必要か

apply 先 task の `returns.name` は task 単体の出力名であり、foreach invocation ごとの collect 結果名ではない。

同じ task を複数の入力集合に適用する場合、apply 先 task の `returns.name` だけでは collect 結果を区別できない。
`foreach.returns` によって invocation 単位の結果名を与えることで、後続 flow から明確に参照できる。

### なぜ `foreach.returns` を optional にするか

foreach は side-effect 目的で使われることもある。
たとえば各 item を外部サービスへ同期する、store に書き込む、通知を送る、といった場合、collect 結果を生成しない flow が自然である。

そのため `foreach.returns` は常に必須にはせず、collect 結果を利用する場合にだけ指定する。

### なぜ型を `list<T>` にするか

foreach は apply 先 task を要素ごとに実行する構文である。
apply 先 task の単一実行が `T` を返すなら、foreach 全体の collect 結果は `list<T>` と考えるのが自然である。

このルールは ADR-060 の TypeRef と相性がよく、named list model と inline `list<T>` の互換性も既存ルールで扱える。

### なぜ `foreach.id` を入れないか

`foreach.id` は foreach invocation の identity を明示する手段になりうるが、現時点では `foreach.returns` が invocation 単位の source 名として機能する。

source として必要なのは collect 結果の名前であり、foreach制御構文自体のIDではない。
不要な識別子を増やすと YAML が重くなり、brewprint の読みやすさを損なうため導入しない。

### なぜ制御フロースコープの例外にするか

foreach内部の各 iteration asset を外部から参照可能にすると、ループ内部構造と外部flowが密結合する。
一方で、明示的に collect された結果は foreach の正当な外部出力であり、後続 task が利用できなければ foreach の表現力が不足する。

したがって、個別 iteration asset は隠蔽し、`foreach.returns` で宣言された collected asset だけを外部 source として公開する。

### 却下した代替案

#### 代替案A: apply 先 task の `returns.name` をそのまま collect 結果名にする

同じ task を複数回 foreach した場合に collect 結果を区別できない。
また、task 単体の return 名と foreach invocation の結果名が混同される。

→ 却下。

#### 代替案B: `foreach.id` を必須にし、source id を `<foreach_id>.<returns>` にする

source id の一意性は担保しやすいが、単純な foreach にも余分なIDが必要になる。
現時点のユースケースでは `foreach.returns` 名だけで十分に区別できる。

→ 却下。

#### 代替案C: `foreach.returns` を常に必須にする

side-effect only の foreach でも未使用の collect 名を強制することになり、設計ノイズが増える。

→ 却下。

#### 代替案D: collected asset を後続 flow から参照不可にする

foreach 結果を後続 step で集計・変換する flow を表現できない。
また、後続ADRで task return source を明示する場合にも、foreach の collect 結果を return source 候補として扱えなくなる。

→ 却下。

#### 代替案E: main task returns.name と foreach.returns の名前一致で暗黙接続する

UC-001 のような例を短く書けるが、task return source の一般則を foreach 固有ADRに混ぜることになる。
main task / composite task が何を返すかは foreach に限らない横断的な仕様であり、`foreach.returns` の source 化とは別に扱うべきである。

→ 却下。本 ADR では採用せず、task return source の明示化は ADR-062 で扱う。

## 影響

### 既存 spec への影響

- `docs/spec/edges.md` §1-5 foreachエントリに、`foreach.returns` が collected asset source 名であることを追記する
- `docs/spec/edges.md` §1-7 flow wiring 型互換性に、wiring source 種別として collected asset source を追加する
  - source 記法: `foreach.returns` で宣言された collected asset source 名
  - source TypeRef: `list<T>`（`T` は apply 先 task の `returns.model`）
  - apply 先 task の `returns.model` が解決不能な場合は source TypeRef も解決不能として扱い、`incompatible_wiring_type` を抑制する
- `docs/spec/diagnostics.md` に `duplicate_flow_source` / `invalid_foreach_returns` / `unresolved_wiring_source` または同等の診断を追加する
- task return source の明示化は ADR-062 で扱うため、本 ADR の spec 反映では `task.returns.source` を追加しない

### 既存 ADR への影響

- ADR-016 の foreach 構文判断は維持する。本 ADR は ADR-016 を supersede せず、`foreach.returns` の意味を追補する
- ADR-023 の制御フロースコープに対し、`foreach.returns` を明示的な escape hatch として追加する。ADR-023 自体は遡及修正しないが、spec/edges.md の制御フロースコープ節には `foreach.returns` が例外として外部公開される旨を追記する
- ADR-060 §5-1 の注で ADR-061 に委ねられていた collected asset source 解決を、本 ADR で確定する。ADR-060 §5 の wiring source 解決ルールに、collected asset source（`foreach.returns` 由来）を4つ目の source 種別として追加する

### 既存実装への影響

- flow source resolver に `foreach.returns` 由来の collected asset source を登録する
- 同一 flow file 内で `foreach.returns` が node id または他の `foreach.returns` と重複していないかを検出する
- `foreach.returns` ありの場合、apply 先 task の `returns.model` から `list<T>` TypeRef を生成する
- `foreach.returns` 省略時は collected asset を semantic model に生成しない。renderer / inspect / MCP もこれを露出しない
- 後続 wiring が collected asset source を参照した場合、生成された `list<T>` を source TypeRef として使う
- apply 先 task の `returns.model` が解決不能な場合、collected asset source の TypeRef も解決不能として扱い、後続 wiring の `incompatible_wiring_type` を抑制する
- main task / composite task の return source 接続は本 ADR では実装対象に含めない。`task.returns.source` は ADR-062 で扱う

### 既存 UC への影響

- UC-001 `cart/task/validate_cart.yaml` の `foreach.returns: validated_items` は、foreach の collected asset source 名として正当化される
- UC-001 で main task がその collected asset を返す接続方法は、本 ADR では未定義とし、ADR-062 の task return source 明示化で扱う
- UC-002 self-hosting で同じ tool / normalize task を複数集合に適用する場合、`foreach.returns` によって collected asset を区別できる

### v1.1 への影響

本 ADR は M15 / data layer expressiveness の Phase B として、v1.1.0-spec に含める。
ADR-060 の TypeRef と flow wiring type compatibility を前提にした forward 拡張であり、v1.0.0-spec の遡及修正ではない。

## Evidence

- commit: a5032d1
- impl commit: 01e7127
- 参考: ADR-016 foreach flow construct, ADR-023 control flow scope, ADR-060 TypeRef / flow wiring type compatibility
