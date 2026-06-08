# V01-ADR-062: task return source の明示化

- **status**: accepted
- **date**: 2026-05-04

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

brewprint の `task.returns` は、task が外部へ返す値の名前と型を表す signature として定義されてきた。

```yaml
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
```

この `returns.name` / `returns.model` は、task の外向き contract として必要である。
特に leaf task / note-only task / external boundary task では、内部 `flow:` を持たなくても「この task は何を返すか」を signature として表現できなければならない。

一方で、`flow:` を持つ main task / composite task では、内部 flow のどの source を task の return として返すのかが未定義だった。

たとえば V01-ADR-061 で `foreach.returns` が collected asset source として定義された結果、以下のような flow source を作れるようになった。

```yaml
flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items
```

しかし、main task `validate_cart` がこの `validated_items` を返すことをどう明示するかは V01-ADR-061 の責務ではない。
`main task returns.name` と `foreach.returns` の名前一致で暗黙接続する案も考えられるが、task return source は foreach に限らない横断的な仕様であり、暗黙一致は分かりにくい。

したがって、task の return source を明示するために `returns.source` を導入する。

## 決定

### 1. `task.returns.source` を導入する

`task.returns` に optional field として `source` を追加する。

```yaml
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
```

`returns.source` は、task が外部へ返す値の source を明示する。

`returns.name` / `returns.model` は task signature として維持する。
`returns.source` は、その signature を満たす値を内部 flow / input からどこで得るかを指定する wiring である。

### 2. `returns.source` は任意である

`returns.source` は optional とする。

- leaf task / note-only task / external boundary task では不要
- `flow:` を持つ main task / composite task が返す値を明示したい場合に指定する
- side-effect only task は `returns` 自体を省略できる

```yaml
# leaf / external boundary task: signature のみでよい
nodes:
  - id: query_service
    type: task
    params:
      - name: request
        model: inspect_request
    returns:
      name: result
      model: any
    note: "QueryService.Inspect を呼び出す。実装詳細は Go 側。"
```

この例では内部 flow source が存在しないため、`returns.source` は不要である。

`flow:` を持たない task でも、`returns.source: $params.<name>` を指定して入力をそのまま返す pass-through task として表現することは正当である。
この場合 task は内部 flow を持たないが、外部 contract として「params の値をそのまま返す」ことを明示できる。

### 3. `returns.source` に指定できる source

`returns.source` の source 解決規則は、V01-ADR-060 §5（wiring source 解決）および V01-ADR-061 §3（collected asset source 参照）を継承する。
すなわち、flow wiring source として解決可能な任意の source を `returns.source` に指定できる。

ただし、以下2点で flow wiring とは扱いを変える。

- 解決失敗 / 文脈不正に対する diagnostic は flow wiring 用ではなく return wiring 用（§8 の `unresolved_return_source` / `invalid_return_source` / `incompatible_return_type`）を使う
- `$item` は `returns.source` に指定できない（§3 後半および §理由を参照）

指定可能な source は以下。

| source | 意味 |
|---|---|
| node id / QualifiedID | task / join など returns を持つ node の出力全体 |
| collected asset source | `foreach.returns` で明示された collected asset |
| `$params.<name>` | main task params の `<name>` をそのまま返す |

`$item` は `returns.source` では使えない。
`$item` は foreach iteration 内部の source であり、task return は foreach の外側で決まるためである。

```yaml
# node output を返す
returns:
  name: report
  model: report
  source: build_report

# foreach collected asset を返す
returns:
  name: validated_items
  model: cart_item_list
  source: validated_items

# 入力をそのまま返す
returns:
  name: unchanged_items
  model: cart_item_list
  source: $params.cart_items
```

### 4. `returns.source` と `returns.model` は型互換性を検証する

`returns.source` を指定した場合、source TypeRef と `returns.model` の TypeRef は V01-ADR-060 の TypeRef compatibility ルールで検証する。

```yaml
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

`validated_items` の source TypeRef が `list<cart_item>` で、`cart_item_list` が `kind: list, element: cart_item` の named list model であれば互換とみなす。

型が互換しない場合は `incompatible_return_type` diagnostic（severity: error）を出す。

### 5. 型解決失敗時は二重診断を抑制する

`returns.source` の TypeRef または `returns.model` の TypeRef が解決不能な場合、`incompatible_return_type` は発行しない。

未解決 source / invalid source / unresolved model / invalid TypeRef など、一次診断を優先する。
この方針は V01-ADR-060 の flow wiring type compatibility における二重診断抑制と同じである。

### 6. name 一致による暗黙 return 接続は採用しない

`returns.name` と flow source 名が一致していても、それだけでは task return source として扱わない。

```yaml
nodes:
  - id: validate_cart
    type: task
    main: true
    returns:
      name: validated_items
      model: cart_item_list

flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items
```

上記では、`foreach.returns: validated_items` と `task.returns.name: validated_items` が一致しているが、`validate_cart` がそれを返すとはみなさない。
返す値を明示する場合は以下のように書く。

```yaml
returns:
  name: validated_items
  model: cart_item_list
  source: validated_items
```

なお、V01-ADR-060 §4 で定義された `join.params` の `returns.name` 一致による暗黙接続は v1.0 由来の既存仕様として維持する。
本 ADR は task return source の明示化に限定したスコープを持ち、join wiring の明示化は本 ADR の範囲外である。
将来 `join.source` のような明示化を検討する場合は別 ADR で扱う。

### 7. `returns.source` は単一 return にのみ対応する

既存仕様どおり、task の `returns` は単一のみである。
`returns.source` もその単一 return に対応する source を指定する。

複数 return / named tuple / multi output task は本 ADR では扱わない。
必要な場合は struct model で wrap して単一 return として表現する。

### 8. diagnostics

本 ADR 受理後、少なくとも以下の診断を spec / 実装に追加する。

| code | severity | 意味 |
|---|---|---|
| `unresolved_return_source` | error | `returns.source` が node id / `$params` / collected asset のいずれとしても解決できない |
| `invalid_return_source` | error | source は解決できたが task return source として使えない。例: returns を持たない node、`$item` |
| `incompatible_return_type` | error | `returns.source` の TypeRef と `returns.model` の TypeRef が互換しない |

`unresolved_return_source` と `invalid_return_source` は区別する。
前者は typo などにより参照先が存在しない場合、後者は参照先は存在するが return source として不正な場合に使う。

既存の wiring diagnostic に統合する場合でも、task return source 上の問題であることを機械的に区別できる形にする。

## 理由

### なぜ `returns.source` が必要か

`returns.name` / `returns.model` は task signature であり、「外から見た返り値の名前と型」を表す。
一方、`flow:` を持つ task では「内部 flow のどの source を返すか」という別の情報が必要である。

この2つを `returns.name` の名前一致で兼用すると、暗黙的で分かりにくく、task の返り値がどこから来るのかを機械的に説明しづらい。
`returns.source` を明示することで、signature と実際の return wiring を分離できる。

### なぜ `returns.model` を維持するか

`returns.model` は task signature の一部である。
これを source から推論するだけにすると、内部 flow を持たない leaf task / note-only task / external boundary task が返り値 contract を表現できなくなる。

brewprint は実装済み関数だけでなく、人間とLLMが合意する設計 contract を表すための言語である。
したがって、`returns.model` は明示的に残す。

### なぜ `source` を optional にするか

すべての task が内部 flow を持つわけではない。
leaf task や external boundary task に `source` を強制すると、存在しない内部値を指定することになり、設計ノイズが増える。

`source` は composite / main task が内部 flow の結果を返すときにだけ必要である。

### なぜ name 一致の暗黙接続を採用しないか

name 一致は短く書けるが、以下の問題がある。

- `returns.name` が signature 名なのか内部 source 名なのか曖昧になる
- task output / collected asset / main return の境界が見えにくくなる
- source が複数ある場合に解決規則が複雑化する
- LLM や実装者が暗黙接続を見落としやすい

brewprint は人間とLLMの共通設計言語であるため、重要な接続は明示した方がよい。

### なぜ `$params.<name>` を許可するか

入力をそのまま返す task は自然に存在する。
たとえば validation / pass-through / no-op fallback / identity transform では、入力を変えずに返すことがある。

`$params.<name>` を `returns.source` として許可することで、このような task を余計な pass-through step なしに表現できる。

### なぜ `$item` を禁止するか

`$item` は foreach iteration 内部の source であり、task 全体の return source ではない。
foreach 全体として返す場合は、V01-ADR-061 の `foreach.returns` で collect した source を `returns.source` に指定する。

### 却下した代替案

#### 代替案A: `returns.name` と flow source 名の一致で暗黙接続する

短く書けるが、signature 名と内部 source 名を兼用することになる。
暗黙接続は見落としやすく、複数 source がある task で意図を読み取りづらい。

→ 却下。

#### 代替案B: `returns.model` を source から推論し、明示不要にする

内部 flow を持つ task では便利だが、leaf task / note-only task / external boundary task が返り値 contract を表せなくなる。
また、task signature が source 解決に依存してしまい、設計 contract と実装構造の境界が曖昧になる。

→ 却下。

#### 代替案C: task return source は `flow:` の最後の entry とする

DAG では最後に見える step が必ず返り値とは限らない。
branch / fork / foreach / side-effect step が混ざると、最後の entry という構文上の位置だけでは意味が決まらない。

→ 却下。

#### 代替案D: task return source は `initializes` / store のみから選ぶ

store や initialized value を返す task は表現できるが、flow 内で生成した通常 task output や foreach collected asset を返せない。
DAG上で計算した値を返す自然な構造を表現できないため不足。

→ 却下。

## 影響

### 既存 spec への影響

- `docs/spec/nodes.md` の `returns` オブジェクトに optional field `source` を追加する
- `docs/spec/edges.md` に「task return wiring」節を新設し、`returns.source` の source 解決ルールおよび V01-ADR-060 §1-7 と同じ TypeRef compatibility ルールを記述する
  - source 解決規則は V01-ADR-060 §5 / V01-ADR-061 §3 を継承する
  - 解決失敗 / 文脈不正に対する diagnostic は return wiring 専用コードを使う
- `docs/spec/diagnostics.md` に `unresolved_return_source` / `invalid_return_source` / `incompatible_return_type` を追加する

### 既存 ADR への影響

- V01-ADR-009 の task I/O contract に、`returns.source` による return wiring を追補する
- V01-ADR-015 の flow wiring とは別に、task return wiring を定義する。ただし source 解決体系は flow wiring source と揃える
- V01-ADR-060 の TypeRef compatibility を task return source 検証にも適用する
- V01-ADR-061 の `foreach.returns` は、`returns.source` から参照可能な collected asset source になる

### 既存実装への影響

- raw YAML / semantic model の `returns` に `source` field を追加する
- `returns.source` の source resolver を追加する
- source TypeRef と `returns.model` の互換性を検証する
- `unresolved_return_source` / `invalid_return_source` / `incompatible_return_type` を diagnostic として出せるようにする
- `returns.source` 未指定の leaf task / external boundary task はこれまで通り有効とする

### 既存 UC への影響

- UC-001 `cart/task/validate_cart.yaml` は、`foreach.returns: validated_items` を main task return として返す場合、`returns.source: validated_items` を明示する形へ更新する
- UC-002 self-hosting で composite task が内部 query / normalize / build_response の結果を返す場合、`returns.source` で明示できる

### v1.1 への影響

本 ADR は M15 / data layer expressiveness の Phase B として、v1.1.0-spec に含める。
V01-ADR-060 / V01-ADR-061 を前提にした forward 拡張であり、v1.0.0-spec の遡及修正ではない。

## Evidence

- commit: a5032d1
- impl commit: e7b8292
- 参考: V01-ADR-009 task I/O design, V01-ADR-015 flow wiring, V01-ADR-060 TypeRef compatibility, V01-ADR-061 foreach collected asset source
