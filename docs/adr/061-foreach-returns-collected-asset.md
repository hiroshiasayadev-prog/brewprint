# 061: foreach.returns collected asset 参�Eルール

- **status**: accepted
- **date**: 2026-05-04

> こ�EADRは起票時点での決定を記録したスナップショチE��である、E> 現在の仕様�E spec を参照すること、E
## 背景

ADR-016 で `foreach` は node type ではなぁE`flow:` セクションの制御構文として定義された、E`foreach` は apply 允Etask を繰り返し実行し、各 iteration の入力�E `$item` で渡す、E
ADR-060 では TypeRef と flow wiring type compatibility を導�Eし、`$item` の型を `foreach.over` の `list<T>` から導�Eするルールを定めた、Eただし、ADR-060 は `foreach.returns` によって collect されぁEasset を後綁Eflow から参�Eする場合�E source id 解決ルールめEADR-061 に委�EてぁE��、E
こ�E未決領域を残すと、以下�Eような自然な flow が仕様上あぁE��ぁE��なる、E
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

また、同ぁEapply 允Etask を褁E��の入力集合に対して使ぁE��合、apply 允Etask の `returns.name` だけでは collect 結果を区別できなぁE��E
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

したがって、`foreach.returns` めEforeach invocation 単位�E collected asset 名として定義し、後綁Eflow からの参�Eルールを�E確化する、E
## 決宁E
### 1. `foreach.returns` は collected asset source 名である

`foreach.returns` は、apply 允Etask の `returns` めEiteration ごとに雁E��ぁEcollected asset に付けめEfile-local source 名とする、E
```yaml
flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items
```

上記�E `validated_items` は、`validate_item` の吁Eiteration の返り値めEcollect した asset source である、Eこれは apply 允Etask の `returns.name` とは別の名前であり、foreach invocation 単位で命名される、E
### 2. `foreach.returns` は任意である

`foreach.returns` は optional とする、E
- collect 結果を後綁Eflow から安定参照する場合�E持E��すめE- collect 結果めEtask return source として使ぁE��合�E持E��すめE- side-effect 目皁E�� apply 允Etask を繰り返すだけ�E場合�E省略できる

```yaml
# side-effect only: collect result を使わなぁEflow:
  - foreach: sync_item
    over: $params.items
    params:
      item: $item
```

apply 允Etask に `returns` がなぁE��もかかわらず `foreach.returns` を指定した場合�E invalid とする、E
`foreach.returns` を省略した場合、collected asset は semantic model 上に生�EされなぁE��Eforeach は side-effect / 個別 iteration の効果�Eみを持ち、collect 結果を表現したぁE��合�E `foreach.returns` を�E示する忁E��がある、E
renderer / inspect / MCP は、`foreach.returns` を省略した foreach を「side-effect only な繰り返し実行」として扱ぁE��E省略時に collected asset の冁E��名や擬似 source を露出してはならなぁE��E
### 3. 後綁Eflow は `foreach.returns` 名を wiring source として参�Eできる

`foreach.returns` で宣言されぁEcollected asset は、同一 flow file 冁E�E後綁Estep / branch / fork / foreach から wiring source として参�Eできる、E
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

こ�E参�Eは file-local source 参�Eであり、`$params` / `$item` のようなシジルは使わなぁE��E
`foreach.returns` で宣言されぁEcollected asset は、その foreach entry が完亁E��た後に生�Eされる、Eしたがって、その foreach 自身の `params` 冁E��ら�E刁E�E身の `returns` 名を参�EしてはならなぁE��E参�Eできるのは、当該 foreach entry より後ろの flow entry からのみである、E
### 4. `foreach.returns` の型�E `list<T>` とする

apply 允Etask の `returns.model` ぁETypeRef `T` の場合、`foreach.returns` の collected asset 型�E `list<T>` とする、E
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

上記では `validated_items` の型�E `list<cart_item>` である、E
apply 允Etask の `returns.model` ぁE`any` の場合、`foreach.returns` の型�E `list<any>` とする、E
apply 允Etask の `returns.model` が解決不�Eな場合、`foreach.returns` の TypeRef も解決不�Eとして扱ぁE��Eこ�E場合、後綁Ewiring が当該 collected asset を参照してぁE��も、`incompatible_wiring_type` は発行しなぁE��E未解決 TypeRef に対する一次診断を優先し、二重診断を避ける、E
### 5. task return source との関俁E
本 ADR は、`foreach.returns` によって flow 冁E�� collected asset source を作るルールを定める、Emain task / composite task がどの flow source を返すか、つまめEtask return source の明示方法�E本 ADR では定義しなぁE��E
`foreach.returns` は、後綁EADR で定義する task return source から参�E可能な flow source になりうる、Eただし、`main task returns.name` と `foreach.returns` の名前一致による暗黙接続�E、本 ADR では採用しなぁE��E
```yaml
# task return source の明示方法�E ADR-062 で扱ぁE��宁Enodes:
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

上記�E `returns.source` は ADR-062 で扱ぁE��定�E例であり、本 ADR の決定篁E��には含めなぁE��E
### 6. `foreach.returns` の名前空間と重褁E��ール

`foreach.returns` は同一 flow file 冁E�E wiring source 名前空間に参加する、Ebare token として参�EされめEsource は単一名前空間で解決され、衝突�E invalid とする、E
同一 flow file 冁E��、`foreach.returns` は以下と重褁E��てはならなぁE��E
- 同一ファイル冁E�E node id
- 他�E `foreach.returns`

重褁E��た場合、bare token (`some_name`) を後綁Ewiring から書ぁE��ときに「node output」と「collected asset source」�Eどちらに解決すべきかが決められなぁE��めEinvalid とする、E
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
    returns: validated_items # invalid: 同一ファイル冁Enode id と重褁E```

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
    returns: validated_items # invalid: collected asset source 名が重褁E```

一方、task の `returns.name` は通常 flow の wiring source ではなぁE��E通常 task の出力を参�Eする場合�E task の node id を使ぁE��め、task の `returns.name` と `foreach.returns` が同名であることは衝突とは扱わなぁE��E
```yaml
nodes:
  - id: validate_item
    type: task
    returns:
      name: validated_items   # task signature 上�E returns.name
      model: cart_item

flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items   # 衝突しなぁE(task.returns.name は wiring source ではなぁE
```

なお、file / module / private scope を趁E��た名前空間�E離�E�たとえ�E sub-task file 墁E��での衝突許容�E��E本 ADR の対象外であり、ADR-002 / ADR-003 / ADR-058 の名前空間ルールに従う、E本 ADR の重褁E��ールは「同一 flow file 冁E�� bare token として参�EされめEwiring source の解決」に限定される、E
### 7. `foreach.id` は導�EしなぁE
本 ADR では `foreach.id` フィールドを導�EしなぁE��E
同じ apply 允Etask を褁E��囁Eforeach する場合�E、`foreach.returns` に異なめEcollected asset 名を与えることで区別する、E
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

`foreach.id` を導�Eすると構文が重くなり、現時点の要件に対して過剰である、E
### 8. 制御フロースコープとの関俁E
ADR-023 は、制御フロー構文の冁E��で生�EされぁEasset をスコープ外から直接参�E不可とした、E
`foreach.returns` はこ�Eルールに対する明示皁E�� escape hatch とする、Eapply 允Etask の iteration 冁E��で生�Eされる個別 asset は外部参�E不可のままだが、`foreach.returns` で明示皁E�� collect した結果だけを flow 外�Eの source として公開できる、E
### 9. diagnostics

本 ADR 受理後、少なくとも以下�E診断めEspec / 実裁E��追加する、E
| code | severity | 意味 |
|---|---|---|
| `duplicate_flow_source` | error | 同一 flow file 冁E�� `foreach.returns` ぁEnode id また�E他�E `foreach.returns` と重褁E��てぁE�� |
| `invalid_foreach_returns` | error | apply 允Etask に `returns` がなぁE�Eに `foreach.returns` が指定されてぁE��、また�E当該 foreach 自身の `params` 冁E��ら�E刁E�E `returns` 名を参�EしてぁE�� |
| `unresolved_wiring_source` | error | wiring source ぁEnode id / `$params` / `$item` / collected asset のぁE��れとしても解決できなぁE|
| `invalid_wiring_source` | error | source は解決できたが、その斁E��では wiring source として使えなぁE|

`unresolved_wiring_source` と `invalid_wiring_source` は区別する、E前老E�E typo などにより参�E先が存在しなぁE��合、後老E�E参�E先�E存在するぁEreturns 相当�E出力を持たなぁEnode めE��効篁E��外�E `$item` など、source として不正な場合に使ぁE��E
既孁Ediagnostic code へ統合する場合でも、上記�E状態を機械皁E��区別できる形にする、E
## 琁E��

### なぁE`foreach.returns` が忁E��か

apply 允Etask の `returns.name` は task 単体�E出力名であり、foreach invocation ごとの collect 結果名ではなぁE��E
同じ task を褁E��の入力集合に適用する場合、apply 允Etask の `returns.name` だけでは collect 結果を区別できなぁE��E`foreach.returns` によって invocation 単位�E結果名を与えることで、後綁Eflow から明確に参�Eできる、E
### なぁE`foreach.returns` めEoptional にするぁE
foreach は side-effect 目皁E��使われることもある、Eたとえ�E吁Eitem を外部サービスへ同期する、store に書き込む、E��知を送る、とぁE��た場合、collect 結果を生成しなぁEflow が�E然である、E
そ�Eため `foreach.returns` は常に忁E��にはせず、collect 結果を利用する場合にだけ指定する、E
### なぜ型めE`list<T>` にするぁE
foreach は apply 允Etask を要素ごとに実行する構文である、Eapply 允Etask の単一実行が `T` を返すなら、foreach 全体�E collect 結果は `list<T>` と老E��る�Eが�E然である、E
こ�Eルールは ADR-060 の TypeRef と相性がよく、named list model と inline `list<T>` の互換性も既存ルールで扱える、E
### なぁE`foreach.id` を�EれなぁE��

`foreach.id` は foreach invocation の identity を�E示する手段になりうるが、現時点では `foreach.returns` ぁEinvocation 単位�E source 名として機�Eする、E
source として忁E��なのは collect 結果の名前であり、foreach制御構文自体�EIDではなぁE��E不要な識別子を増やすと YAML が重くなり、brewprint の読みめE��さを損なぁE��め導�EしなぁE��E
### なぜ制御フロースコープ�E例外にするぁE
foreach冁E��の吁Eiteration asset を外部から参�E可能にすると、ループ�E部構造と外部flowが寁E��合する、E一方で、�E示皁E�� collect された結果は foreach の正当な外部出力であり、後綁Etask が利用できなければ foreach の表現力が不足する、E
したがって、個別 iteration asset は隠蔽し、`foreach.returns` で宣言されぁEcollected asset だけを外部 source として公開する、E
### 却下した代替桁E
#### 代替桁E: apply 允Etask の `returns.name` をそのまま collect 結果名にする

同じ task を褁E��囁Eforeach した場合に collect 結果を区別できなぁE��Eまた、task 単体�E return 名と foreach invocation の結果名が混同される、E
ↁE却下、E
#### 代替桁E: `foreach.id` を忁E��にし、source id めE`<foreach_id>.<returns>` にする

source id の一意性は拁E��しめE��ぁE��、単純な foreach にも余�EなIDが忁E��になる、E現時点のユースケースでは `foreach.returns` 名だけで十�Eに区別できる、E
ↁE却下、E
#### 代替桁E: `foreach.returns` を常に忁E��にする

side-effect only の foreach でも未使用の collect 名を強制することになり、設計ノイズが増える、E
ↁE却下、E
#### 代替桁E: collected asset を後綁Eflow から参�E不可にする

foreach 結果を後綁Estep で雁E���E変換する flow を表現できなぁE��Eまた、後続ADRで task return source を�E示する場合にも、foreach の collect 結果めEreturn source 候補として扱えなくなる、E
ↁE却下、E
#### 代替桁E: main task returns.name と foreach.returns の名前一致で暗黙接続すめE
UC-001 のような例を短く書けるが、task return source の一般剁E�� foreach 固有ADRに混ぜることになる、Emain task / composite task が何を返すか�E foreach に限らなぁE��断皁E��仕様であり、`foreach.returns` の source 化とは別に扱ぁE��きである、E
ↁE却下。本 ADR では採用せず、task return source の明示化�E ADR-062 で扱ぁE��E
## 影響

### 既孁Espec への影響

- `docs/spec/edges.md` §1-5 foreachエントリに、`foreach.returns` ぁEcollected asset source 名であることを追記すめE- `docs/spec/edges.md` §1-7 flow wiring 型互換性に、wiring source 種別として collected asset source を追加する
  - source 記況E `foreach.returns` で宣言されぁEcollected asset source 吁E  - source TypeRef: `list<T>`�E�ET` は apply 允Etask の `returns.model`�E�E  - apply 允Etask の `returns.model` が解決不�Eな場合�E source TypeRef も解決不�Eとして扱ぁE��`incompatible_wiring_type` を抑制する
- `docs/spec/diagnostics.md` に `duplicate_flow_source` / `invalid_foreach_returns` / `unresolved_wiring_source` また�E同等�E診断を追加する
- task return source の明示化�E ADR-062 で扱ぁE��め、本 ADR の spec 反映では `task.returns.source` を追加しなぁE
### 既孁EADR への影響

- ADR-016 の foreach 構文判断は維持する。本 ADR は ADR-016 めEsupersede せず、`foreach.returns` の意味を追補すめE- ADR-023 の制御フロースコープに対し、`foreach.returns` を�E示皁E�� escape hatch として追加する、EDR-023 自体�E遡及修正しなぁE��、spec/edges.md の制御フロースコープ節には `foreach.returns` が例外として外部公開される旨を追記すめE- ADR-060 §5-1 の注で ADR-061 に委�EられてぁE�� collected asset source 解決を、本 ADR で確定する、EDR-060 §5 の wiring source 解決ルールに、collected asset source�E�Eforeach.returns` 由来�E�を4つ目の source 種別として追加する

### 既存実裁E��の影響

- flow source resolver に `foreach.returns` 由来の collected asset source を登録する
- 同一 flow file 冁E�� `foreach.returns` ぁEnode id また�E他�E `foreach.returns` と重褁E��てぁE��ぁE��を検�Eする
- `foreach.returns` ありの場合、apply 允Etask の `returns.model` から `list<T>` TypeRef を生成すめE- `foreach.returns` 省略時�E collected asset めEsemantic model に生�EしなぁE��renderer / inspect / MCP もこれを露出しなぁE- 後綁Ewiring ぁEcollected asset source を参照した場合、生成された `list<T>` めEsource TypeRef として使ぁE- apply 允Etask の `returns.model` が解決不�Eな場合、collected asset source の TypeRef も解決不�Eとして扱ぁE��後綁Ewiring の `incompatible_wiring_type` を抑制する
- main task / composite task の return source 接続�E本 ADR では実裁E��象に含めなぁE��`task.returns.source` は ADR-062 で扱ぁE
### 既孁EUC への影響

- UC-001 `cart/task/validate_cart.yaml` の `foreach.returns: validated_items` は、foreach の collected asset source 名として正当化されめE- UC-001 で main task がその collected asset を返す接続方法�E、本 ADR では未定義とし、ADR-062 の task return source 明示化で扱ぁE- UC-002 self-hosting で同じ tool / normalize task を褁E��雁E��に適用する場合、`foreach.returns` によって collected asset を区別できる

### v1.1 への影響

本 ADR は M15 / data layer expressiveness の Phase B として、v1.1.0-spec に含める、EADR-060 の TypeRef と flow wiring type compatibility を前提にした forward 拡張であり、v1.0.0-spec の遡及修正ではなぁE��E
## Evidence

- commit: a5032d1
- impl commit: 01e7127
- 参老E ADR-016 foreach flow construct, ADR-023 control flow scope, ADR-060 TypeRef / flow wiring type compatibility
