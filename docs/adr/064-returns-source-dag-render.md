# 064: returns.source の DAG render ルール

- **status**: accepted
- **date**: 2026-05-05

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-062 で `task.returns.source` が導入され、ADR-063 で initialized source も指定可能になった。
さらに ADR-063 では initialized source を flow 内部 wiring（step.params など）の bare token source としても参照可能とした。

これにより、DAG 上で表現すべき新しい source / edge 種別は以下のとおり拡張された。

`returns.source` から参照される source 種別:

1. node id / QualifiedID
2. collected asset source（`foreach.returns` 由来）
3. initialized source（`initializes[].name` 由来）
4. `$params.<name>`

flow 内部 wiring の bare token source 種別:

1. node id / QualifiedID
2. collected asset source（`foreach.returns` 由来）
3. initialized source（`initializes[].name` 由来）

ADR-024 で DAG render の境界ノード（`subgraph params` / `subgraph returns`）が定義されている。
しかし、`returns.source` および flow 内部 wiring からの initialized source 参照を Mermaid DAG 上でどう可視化するかは未定義であった。

本ADR起票時の検討過程で、以下の認識転換が起きた。

### 認識1: `subgraph returns` は単一 return 前提では冗長

ADR-009 / spec/nodes.md §returns オブジェクト で「returns は単一のみ」と定められている。
複数返しが必要な場合は struct model で wrap して単一にする。

`subgraph returns` は ADR-024 起票時に「複数 return になった場合に boundary を視覚化する」想定で導入されたが、現行仕様では returns は構造的に単一であり、subgraph で囲う情報量上のメリットは「asset を返す場合に名前が boundary に表示される」程度に縮退している。

一方、`subgraph params` / `subgraph initializes`（後者は本ADR新設）で囲まれた境界 node を `returns.source` で参照する場合、`source --> boundary asset --> boundary returns` のような3段 edge になり、冗長性が顕在化する。

### 認識2: returns は実装上 function call の戻り値

task は実装上 `task()` のように function call として呼ばれ、戻り値は呼び出し側で受け取られる。
`returns.name` は docstring 的な命名であり、言語仕様としての意味は `returns.model`（型）が本質である。

一方、`returns.source` は「内部 source を外向き task signature の return として返す」wiring である。DAG 上で `returns.name` を boundary node として描く必要性は低いが、return data line の role label として表示することで、内部 source が外向きにどの return name として返されるかを自然に示せる。

### 認識3: `_end` を ObjectFlow 終点として再利用できる

ADR-024 で `_end` は ControlFlow の終点として確立されている。
Mermaid 制約上、ActivityFinalNode と ActivityParameterNode (output) を別シンボルで表現することが難しいため、`_end` に ObjectFlow（data line）も収束させることで、単一 return という構造的事実を視覚化できる。

### 認識4: initialized store は subgraph で囲うべき

ADR-065 §決定 §5 の役割対比で、initialized store は store node と mutability 観点で等価とされた。
一方、ADR-014 §3 で initialized store は file-private とされており、外部 store node とはスコープが異なる。

`initializes[]` は ADR-014 起票時に「変数宣言 task の省略形」として設計され、ADR-063 で flow wiring source および `returns.source` の参照対象として正規化された。task 内部の名前付き source として位置付けられた以上、DAG 上でも task の「3境界」（params 入力 / initializes 内部宣言 / 戻り値出力）を視覚的に揃えるのが整合的である。

## 決定

### 1. `subgraph returns` を廃止する

ADR-024 §1〜§3 の `subgraph returns` 描画ルールを廃止する。

`returns.name` を表す boundary asset node は描かず、`returns.source` で指定された値を表す node から `_end` への label 付き data line で「これが return される値」を表現する。edge label は固定 prefix `returns as ` と `task.returns.name` を連結した文字列とする。

```txt
<value_node> -- "returns as <returns.name>" --> _end
```

このとき、task / join / collected asset のように asset を生成する source では、task node から直接 `_end` へ data line を引かない。通常どおり asset node を生成し、その asset node から `_end` へ data line を引く。data line は「値」を運ぶ線であり、task node そのものではなく、その task が生成した asset を経由する方が意味的に正確である。

`returns.name` は boundary node としては DAG に出さない。ただし return data line の label として表示する。Tasks 詳細セクションの `#### Returns` table でも参照可能であり、DAG 上では node を増やさず edge role として補助表示する。

### 2. `_end` は ControlFlow 終点 + ObjectFlow 終点を兼ねる

ADR-024 §4 で確立された `_end([End])` は、本ADRから以下の2役を担う。

- **ControlFlow 終点**: 最後の task / floating ノードからの制御線（`==>`）の収束先（既存挙動）
- **ObjectFlow 終点**: `returns.source` で指定された値を表す node からの label 付き data line（`-- "returns as <returns.name>" -->`）の収束先（新規）

UML 2.x では ActivityFinalNode と ActivityParameterNode (output) は別シンボルだが、Mermaid flowchart にはこの区別を表現する慣習記法がない。単一 return という brewprint の構造的制約（ADR-009）を踏まえ、`_end` に両 flow を収束させることでこの制約を逆に視覚化する。

`_end` への return data line と control line は両立する。同一の最終 task が `returns.source` の source でもある場合でも、control line（`task ==> _end`）は task node から引き、return data line（`asset -- "returns as <returns.name>" --> _end`）は task が生成した asset node から引く。task / join が `returns` を持たず asset node を生成しない場合は、return source として不正であり、この render ルールの対象外である。

### 3. source 種別ごとの edge 起点

`returns.source` の source 種別ごとに、`_end` への return data line 起点を以下のとおり定める。edge label は source 種別を問わず `returns as <returns.name>` とする。

| source 種別 | `_end` への return data line 起点 | 例 |
|---|---|---|
| node id / QualifiedID | 当該 task / join が生成する asset node | `result -- "returns as report" --> _end` |
| collected asset source（`foreach.returns`） | `foreach.returns` 名で生成される collected asset node | `results -- "returns as report" --> _end` |
| initialized source（`initializes[].name`） | `subgraph initializes` 内の store node | `report -- "returns as report" --> _end` |
| `$params.<name>` | `subgraph params` 内の boundary asset | `config -- "returns as report" --> _end` |

node id / QualifiedID を `returns.source` に指定した場合、DAG 上では従来どおり当該 node から `returns.name` の asset node を生成し、その asset node から `_end` へ data line を引く。

```txt
transform --> result([result])
result -- "returns as report" --> _end
transform ==> _end
```

collected asset source は ADR-061 で独立した名前を持つ source であり、DAG 上でも `foreach.returns` 名の collected asset node を生成する。foreach apply 先 task から collected asset node への data line を引き、その collected asset node から `_end` へ data line を引く。

```txt
process_item --> results([results])
results -- "returns as report" --> _end
process_item ==> _end
```

initialized source と `$params.<name>` はすでに値を表す node（initialized store node / params boundary asset）を持つため、その node から `_end` へ label 付き return data line を引く。間に identity を示す表現は挟まない。

### 4. `subgraph initializes` を新設する

`initializes[]` で宣言された file-private store を、`subgraph initializes` で囲って描画する。

```txt
subgraph initializes
  report[(report)]
  cache[(cache)]
end
```

`subgraph params` / `subgraph initializes` / 本体 / `_end` のソース上の記述順を推奨する。Mermaid flowchart は subgraph 配置順序のヒントとしてソース順を使うが、最終配置はレンダリングエンジンに委ねられる。

`initializes` が空の task では `subgraph initializes` を出力しない。

### 5. initialized store の node 形状

`subgraph initializes` 内の initialized store は、シリンダー形状 `[(label)]` で描画する。
これは spec/views/dag.md §store の store node と同じ形状である。

initialized store と通常の store node は ADR-065 §決定 §5 で mutability 観点で等価とされており、形状を分ける積極的理由はない。両者の区別は subgraph の枠と classDef の色で表現する（§6 参照）。

### 6. `initStoreNode` classDef を新設する

initialized store 専用の classDef を追加する。

```
classDef initStoreNode fill:#F0C674,stroke:#B07820,color:#000
```

storeNode の `fill:#E8A838` を彩度落とした薄オレンジ。`stroke` は storeNode と同色を維持し、store 系統であることを示す。`color:#000` は ADR-066 の WCAG 2.1 AA 準拠方針に従う（コントラスト比 13.03、AA 基準 4.5 を満たす）。

storeNode と分けることで、設計者・LLM が DAG 上で「これは module-level の store か、task ファイル内に閉じた initialized store か」を視覚的に判別できる。両者は flow wiring / cross-edge の対象としては等価だが、スコープと宣言経路が異なる（ADR-065 §決定 §5）。

### 7. edge 種別は node の種類ではなく edge の役割で決める

initialized store は store 形状で描かれるが、それに接続される edge の種別は **node の種類ではなく edge の役割** で決定する。

これは ADR-063 §7 / ADR-065 §決定 §1 で確立された「flow param wiring = 値の受け渡し contract」「`reads`/`writes` = 副作用 / store access contract」という役割分担を、DAG render 上で具現化する原則である。

### 7-1. flow wiring からの initialized source 参照は data line で描く

flow 内部 wiring（step.params / branch.params / fork.branches[].steps[].params / foreach.params / branch.cases[].params）から initialized source を bare token で参照する場合、それは値の受け渡し contract である。
DAG 上では通常の data line（`-->`）で描画する。

```yaml
flow:
  - foreach: append_item
    over: $params.items
    params:
      report: report   # initialized source を bare token で参照
      item: $item
```

```txt
report --> append_item
```

cross-edge `reads` 表現に統合しない。

### 7-2. cross-edge `reads`/`writes` 宣言は store access line で描く

cross-edge `reads:` / `writes:` の宣言は store access contract であり、ADR-044 で確立された store access line（ラベル付き edge）で描画する。

```yaml
- id: append_item
  reads: [report]
  writes: [report]
```

`reads` のみ:
```txt
report -- "read" --> append_item
```

`writes` のみ:
```txt
append_item -- "write" --> report
```

両方:
```txt
append_item <-- "read/write" --> report
```

これは spec/views/dag.md §エッジのrender §データ線 で確立されたルールを initialized store にも適用するものである。

### 7-3. 同一 initialized store node に data line と store access line が両立する

§7-1 / §7-2 の原則により、同一の initialized store node には複数の edge が接続されうる。

例: `report` を flow wiring で `append_item` に渡し、かつ `append_item` が `reads`/`writes` を宣言する場合：

```txt
report --> append_item                          %% flow wiring (data line)
append_item <-- "read/write" --> report         %% cross-edge (store access line)
```

両者は独立した宣言として描画する。

### 7-4. 推論・吸収・省略は行わない

以下の暗黙挙動は **行わない**。

- flow wiring で initialized source を渡している事実から `reads`/`writes` を推論しない
- `reads`/`writes` を宣言している事実から flow wiring の data line 描画を省略しない
- initialized source を渡しているからといって、data line を store access line に吸収・置換しない
- store access line が引かれているからといって、data line を省略しない

両 edge は独立した契約であり、DAG はそれぞれを忠実に描画する。これは ADR-063 §7 の「flow param wiring と cross-edge `reads`/`writes` は役割が独立した宣言として併存する」原則の DAG render 表現である。

将来 lint レベルで「`reads` 宣言なしに wiring から参照している」「flow に渡していない store の writes 宣言」等の整合性検査を導入する余地はあるが、本ADRでは規定しない（ADR-063 §7 末尾と同方針）。

### 8. 境界 boundary 体系の更新

ADR-024 §1〜§3 の boundary 体系は本ADRで以下のとおり再構成される。

| 境界 | 表現 | 状態 |
|---|---|---|
| 入力 (`$params`) | `subgraph params` + boundary asset | 維持（ADR-024） |
| 内部宣言 (`initializes[]`) | `subgraph initializes` + initialized store node | 新設（本ADR §4） |
| 出力 (`returns`) | `returns.source` の値 node から `_end` への label 付き return data line | `subgraph returns` は廃止 + return data line として再構成（本ADR §1, §2, §3） |

`boundaryNode` classDef は `subgraph params` 内の boundary asset にのみ適用する。`subgraph returns` は廃止のため `boundaryNode` の使用範囲が縮小する。

### 9. classDef 一覧（DAG 全体）

ADR-066 + 本ADRの帰結として、DAG render の classDef は以下のとおり。

```
classDef taskNode      fill:#4A90D9,stroke:#2C5F8A,color:#000
classDef assetNode     fill:#5BA55B,stroke:#3A6B3A,color:#000
classDef storeNode     fill:#E8A838,stroke:#B07820,color:#000
classDef initStoreNode fill:#F0C674,stroke:#B07820,color:#000
classDef branchNode    fill:#9B6BBD,stroke:#6B3D8F,color:#000
classDef forkNode      fill:#8A8A8A,stroke:#5A5A5A,color:#000
classDef terminalNode  fill:#2C2C2C,stroke:#000,color:#fff
classDef boundaryNode  fill:#2D7D9A,stroke:#1A5068,color:#fff
classDef external      fill:#E0E0E0,stroke:#999,color:#555
```

### 10. ADR-024 / ADR-022 / ADR-066 との関係

- **ADR-024**: §1〜§3 の `subgraph returns` 描画ルールは本ADRで superseded（partial）。`subgraph params` / boundaryNode classDef / `_start`/`_end` の `_` prefix ルール（§4）は維持。ADR-024 自体は本文を遡及修正せず、冒頭に partial superseded annotation を追加する
- **ADR-022**: `initStoreNode` 新設はADR-022 の色テーブル拡張になる。ADR-022 自体は遡及修正せず、本ADRが補追する
- **ADR-066**: `initStoreNode` の `color:#000` は ADR-066 の WCAG 2.1 AA 準拠方針に従う

## 理由

### なぜ `subgraph returns` を廃止するか

ADR-024 §1 の「複数 return を視覚化する」目的は、ADR-009 の単一 return 制約により構造的に達成不要となっている。`subgraph returns` で囲うメリットは「return signature 名が boundary に表示される」程度に縮退しており、ADR-063 で initialized source / collected asset source / `$params.<name>` 等の多様な source が `returns.source` の参照対象になった現状では、`value node --> returns boundary` という専用 boundary 表現が冗長になる。

一方で、return data line が単なる `value --> _end` だけでは「終了した」ことは分かっても「外向き return signature として返した」ことが弱く見える。そこで returns boundary node は復活させず、edge label として `returns as <returns.name>` を付与する。これは boundary node の再導入ではなく、ObjectFlow の role を明示する軽量な表現である。

ただし、`subgraph returns` の廃止は asset node の廃止を意味しない。task / join / foreach collected asset が値を生成する場合、その値を表す asset node は従来どおり描画する。廃止するのは `returns.name` を境界表示するためだけの returns boundary node である。

`returns.name` は DAG node としては表示しないが、return data line の label と Tasks 詳細セクションの `#### Returns` table で参照可能にする。これにより、node 数を増やさずに「どの外向き return name として返されるか」を読み取れる。

### なぜ `_end` に ObjectFlow を収束させるか

UML 2.x では ActivityFinalNode と ActivityParameterNode (output) を別シンボルで表現するが、Mermaid flowchart にはこの区別を表現する慣習記法がない。

brewprint は単一 return 制約（ADR-009）を持つため、「`_end` に来る data line が return」という解釈は一意に定まる。Mermaid 制約と単一 return 構造を両立させる解として、`_end` に両 flowを収束させる方法を採用する。

data line は値の受け渡しを表すため、task / join node から直接 `_end` へ引くのではなく、task / join が生成した asset node から `_end` へ引く。これにより、制御主体としての task node と、返される値としての asset node を混同しない。

さらに label を `returns as <returns.name>` に固定することで、単なる dataflow ではなく task return wiring であることを視覚的に区別する。label は source 名ではなく edge role と外向き return name を表す。

### なぜ initialized store を subgraph で囲うか

`initializes[]` は ADR-014 起票時に「変数宣言 task の省略形」として設計された。ADR-063 で flow wiring source および `returns.source` の参照対象として正規化されたことで、task 内部の名前付き source として位置付けが確立した。

DAG 上で task の3境界（params 入力 / initializes 内部宣言 / 戻り値出力）を視覚的に揃えることで、ADR-014 / ADR-063 / ADR-065 で確立された initialized store の役割（file-private、mutable、task 内部の名前付き source）が一目で理解できる。

### なぜ `initStoreNode` classDef を新設するか

ADR-065 §決定 §5 で initialized store と store node は mutability 観点で等価とされたが、宣言経路（`initializes[]` vs `store/*.yaml`）とスコープ（file-private vs module-level）は明確に異なる。

storeNode を流用した場合、設計者・LLM が DAG 上で「これは DB / 永続化対象の module-level store か、task 内部に閉じた initialized store か」を判別できなくなる。色を分けることで、`reads`/`writes` 宣言の意味（外部副作用 vs file-private mutation）を視覚的に区別できる。

色決めは公知技術根拠を持たない brewprint render 実装上の選択であり、ADR-022 §決定 §「ノードの色付け」の方針（公知技術根拠は形状とエッジ種別、色は実装上の選択）と整合する。

### なぜ edge 種別を node の種類ではなく edge の役割で決めるか

ADR-063 §7 / ADR-065 §決定 §1 で「flow param wiring = 値の受け渡し contract」「`reads`/`writes` = 副作用 / store access contract」という役割分担が確立された。

DAG 上で initialized store が store 形状で描かれることを理由に、flow wiring 経由の参照を store access line に吸収すると、ADR-063 / ADR-065 の役割分担が DAG render で消失する。

役割分担は contract レベルで独立しており、DAG はそれぞれを独立した edge として描画することで両 contract の存在を保存する。

### なぜ推論・吸収・省略を行わないか

flow wiring と cross-edge `reads`/`writes` は contract が独立しているため、片方から他方を推論すると、設計者が明示的に書いた宣言と推論結果が乖離した場合に診断が困難になる。

DAG render は YAML に書かれた事実を忠実に描画するレイヤーであり、契約の推論・正規化は別レイヤー（lint / validation）の責務である。本ADRはこの責務分離を維持する。

### 却下した代替案

#### 代替案A: `subgraph returns` を維持し initialized source / collected asset source の場合は edge 表現を分岐させる

利点: ADR-024 の boundary 体系が温存される。
欠点: source 種別ごとの edge 分岐ルールが複雑化する。冗長な3段 edge が温存される。`returns.name` を boundary node に表示する情報量上のメリットが、Tasks 詳細セクションの `#### Returns` table と重複する。

→ 却下。

#### 代替案B: edge label なしで `_end` へ data line を引く

`value --> _end` のように、return data line を通常 data line と同じ無ラベル edge として描画する案。

利点: edge label が増えず、図が簡潔。
欠点: `value --> _end` だけでは通常の dataflow と return wiring の区別が弱い。`_end` が ControlFlow 終点も兼ねるため、何が外向き return として返されるのかが読み取りづらい。

→ 却下。`returns as <returns.name>` label を付け、return wiring の role を明示する。

#### 代替案C: initialized store を subgraph で囲わず通常の store node と同じ位置に描く

利点: render 実装が単純化する。subgraph 数が増えない。
欠点: ADR-065 §決定 §5 の宣言経路 / スコープの差異が DAG から見えない。task の3境界が視覚的に揃わない。

→ 却下。

#### 代替案D: initStoreNode を新設せず storeNode を流用する

利点: classDef 増えない。ADR-065 で mutability 観点で等価とされた事実と整合する。
欠点: DB / 永続化対象の module-level store と、task 内部に閉じた initialized store の区別が色から消える。`reads`/`writes` 宣言の意味（外部副作用 vs file-private mutation）が視覚的に判別できない。

→ 却下。`subgraph initializes` の枠だけでは色情報の追加価値が出ないため、色も分ける。

#### 代替案E: flow wiring から initialized source を参照する場合、cross-edge `reads` 表現に統合する

利点: 同じ store に対する edge が1本に集約される。DAG が密にならない。
欠点: ADR-063 §7 / ADR-065 §決定 §1 の役割分担が DAG render で消失する。flow wiring と `reads`/`writes` 宣言が同じ edge に潰れることで、設計者が明示した contract と DAG の表現が乖離する。

→ 却下。

#### 代替案F: flow wiring から initialized source を参照している事実から `reads` 宣言を推論し DAG に出力する

利点: 設計者が `reads` を書き忘れても DAG が補完する。
欠点: YAML に書かれていない宣言を DAG が描くことで、DAG render layer が validation layer の責務を侵食する。設計者が明示した宣言と DAG の乖離が診断困難になる。

→ 却下。本ADRは render layer と lint / validation layer の責務分離を維持する。

## 影響

### spec への影響

- `docs/spec/views/dag.md` §ノードのrender に `subgraph initializes` 節を新設し、`initStoreNode` classDef を追加する
- `docs/spec/views/dag.md` §ノードのrender §start / end の `_end` 説明に「ObjectFlow 終点を兼ねる」旨を追記する
- `docs/spec/views/dag.md` §ノードのrender §subgraph params / returns から `subgraph returns` 部分を削除する
- `docs/spec/views/dag.md` §ノードの色付け に `initStoreNode` classDef を追加する
- `docs/spec/views/dag.md` §エッジのrender に `returns.source` の edge ルール節を新設する
- `docs/spec/views/dag.md` の全 render 例（基本DAG / fork-join / store / foreach / branch）を新ルールに従い書き換える
  - `subgraph returns` の削除
  - `returns.source` の値 node から `_end` への label 付き return data line 追加
  - task / join / collected asset source では asset node を経由することの反映
  - return data line label として `returns as <returns.name>` を出力することの反映
  - 該当する場合 `subgraph initializes` の追加
- `docs/spec/views/dag.md` §エッジのrender §データ線 の store access line 説明を、initialized store にも適用される旨に拡張する
- ADR-066 の color 修正と本ADRの initStoreNode 追加を classDef 一覧に反映する

### 既存 ADR への影響

- **ADR-024**: §1〜§3 の `subgraph returns` 描画ルールが本ADRで superseded（partial）。冒頭に partial superseded annotation を追加。本文は遡及修正しない（doc-policy.md §3）
- **ADR-022**: `initStoreNode` classDef 追加は色テーブルの拡張。ADR-022 自体は遡及修正しない（既に ADR-066 で partial superseded annotation 済み）
- **ADR-062 / ADR-063**: §「DAG render は別 ADR で扱う」の保留事項を本ADRで確定する
- ADR-014 / ADR-020 / ADR-044 / ADR-061 / ADR-065: 本文への影響なし

### 既存実装への影響

- `internal/render/dag` から `subgraph returns` 出力ロジックを削除する
- `internal/render/dag` に `subgraph initializes` 出力ロジックを追加する
- `internal/render/dag` に `returns.source` 解決と `_end` への label 付き return data line 出力を追加する
  - 4 source 種別（node id / collected asset source / initialized source / `$params.<name>`）に対応
  - node id / QualifiedID / collected asset source では asset node を経由して `_end` へ接続する
  - edge label は `returns as <returns.name>` とする
- `internal/render/dag` の classDef 出力に `initStoreNode` を追加する
- flow wiring と cross-edge `reads`/`writes` の両方が同じ initialized store node を参照する場合、両 edge をそれぞれ独立して出力する
- 出力の Mermaid 文字列が変わるため、文字列比較ベースの test が変わる

### 既存 UC への影響

- UC-001 の renders/ 配下 Mermaid 出力が以下のとおり変わる
  - `subgraph returns` 消滅
  - `returns.source` の値 node から `_end` へ label 付き return data line 追加
  - task / join / collected asset source では asset node から `_end` へ label 付き return data line 追加
  - initialized store を持つ task は `subgraph initializes` 追加
  - classDef に `initStoreNode` 追加
- ADR-066 の文字色変更と合わせて golden test を一括更新する

### v1.1 への影響

本ADRは ADR-062 / ADR-063 / ADR-066 と合わせて v1.1.0-spec の凍結対象に含める。
v1.0.0-spec の遡及修正ではない forward 拡張である。

### 後続 ADR への影響

- 本ADRで `subgraph returns` が廃止されたことで、ADR-024 の boundary 体系が `subgraph params` + `subgraph initializes` の2系統になる
- 将来 `reads` 宣言なしに wiring から initialized source を参照している等の整合性検査を lint レベルで導入する場合、別 ADR で扱う

## Evidence

- commit: eb891f2
- impl commit: tbd
- 参考: ADR-009 task IO design (single return), ADR-014 initializes field, ADR-020 cross-edge management, ADR-024 DAG boundary nodes, ADR-044 store access edge labels, ADR-061 foreach collected asset, ADR-062 task return source, ADR-063 task return source / initialized store, ADR-065 asset immutability and edge role contrast, ADR-066 DAG classDef WCAG fix, OMG UML 2.x ActivityFinalNode / ActivityParameterNode
