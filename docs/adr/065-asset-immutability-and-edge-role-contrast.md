# 065: asset の immutability と edge 体系における役割対比

- **status**: accepted
- **date**: 2026-05-05

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-010 で `model` / `asset` 分離が確立し、ADR-014 で `initializes` による file-private store が導入され、ADR-063 で initialized source が flow wiring source 名前空間および `returns.source` の指定可能 source に追加された。

この体系の中で、`asset` の mutation 可否（asset が一度生成された後に書き換えられるか否か）について、現行 docs に明示的な規定が存在しない。

- ADR-010 §決定 §4 では asset を「task の `returns` 宣言から暗黙的に生まれるフロー上の存在」と定義しているが、生成後の状態については規定していない
- ADR-063 §理由 では initialized source を「mutation を前提にできる」と書きながら、asset との対比は明文化されていない
- spec/nodes.md §asset も「フロー上の存在」「task の returns から暗黙的に生まれる」と記すのみ

その結果、設計者・LLM の双方で「asset は append できるのか」「task が後続で asset を更新する記述は valid か」という疑問が暗黙に発生しうる。本ADRはこの空白を埋める。

### 現行 edge 体系の構造的事実

brewprint には2種類の edge 概念が共存している。

| edge 種別 | 対象 | 役割 | 出典 |
|---|---|---|---|
| flow wiring（dataflow） | task の returns / asset / initialized source / collected asset / `$params.<name>` / `$item` | 値の受け渡し contract | ADR-009, ADR-040, ADR-060〜063 |
| cross-edge（reads / writes） | `store-id` のみ | store access contract（task が store と関わる事実の宣言） | ADR-020, ADR-044 |

spec/edges.md §3 に「`reads` / `writes` は `list<store-id>`」と明記されており、cross-edge の対象は store のみである。asset は cross-edge の対象に含まれない。

### 違和感の発端

`initializes` 由来の store は ADR-063 で「mutation を前提にできる runtime instance」と位置付けられた。これと対比したとき、asset が「同じく mutation を前提にできる」のか「output snapshot として扱われるのか」が docs から読み取れない。

ADR-014 / ADR-063 で initialized store の mutation セマンティクスを規定するのと同等の整理を、asset 側にも行う必要がある。

## 決定

### 1. asset は cross-edge の対象外であることを明文化する

cross-edge `reads:` / `writes:` は `store-id` のみを受ける（spec/edges.md §3, ADR-020）。asset はこの対象に含まれない。

すなわち、brewprint は「task が asset を後続で書き換える」ことを表現する文法を持たない。task が値を生むのは `returns` 宣言を通してのみであり、その後の書き換え経路は語彙上存在しない。

### 2. asset は構造的に immutable な output snapshot として扱われる

§1 の構造的事実から、asset は brewprint の語彙レベルで以下のように扱われる。

- asset は task の単一実行が生んだ output snapshot として位置付けられる
- asset の生成後の状態変更を表す YAML 記述は brewprint には存在しない
- asset を「累積する箱」「後から append される対象」として記述することはできない

この immutability は、ADR で新たに「決定する」ものではなく、現行 edge 体系（ADR-020 / ADR-063 §7）から構造的に導かれる帰結である。本ADRはこの帰結を明文化する。

### 3. mutable な runtime instance が必要な場合は store を使う

mutation を前提にする runtime instance が設計上必要な場合、以下のいずれかを使う。

- module-level に共有される runtime instance → `store node`（`store/*.yaml` で `type: store` として宣言）
- task ファイル内に閉じた runtime instance → `initialized store`（`initializes[]` で宣言）

両者とも `reads` / `writes` cross-edge の対象であり、mutation を表現する語彙を持つ。詳細は spec/nodes.md §store / §init オブジェクト を参照。

### 4. 実装言語上の挙動は brewprint のスコープ外

本ADRが規定するのは brewprint の語彙レベルの contract である。実装言語側で asset 値が参照渡しによって意図せず書き換わるか否か、Go / Python 等の言語仕様レベルの mutation 挙動については brewprint としては規定しない。

これは ADR-063 §3 の「具体的 mutation semantics（参照渡し / コピー / append の挙動）は task 実装に委ねられ、brewprint としては規定しない」と同じトーンである。

### 5. 役割対比の整理

asset と store の役割対比を以下のとおり整理する。

| 観点 | asset | store node | initialized store |
|---|---|---|---|
| 宣言経路 | `task.returns` から暗黙生成 | `store/*.yaml` で `type: store` | `initializes[]` で宣言 |
| ファイル | 持たない | 持つ | 持たない（task ファイル内） |
| スコープ | task 1実行の output | module-level（QualifiedID 参照可） | file-private |
| edge 体系 | flow wiring の対象 | flow wiring + cross-edge の対象 | flow wiring + cross-edge の対象 |
| 書き込み文法 | なし（cross-edge 対象外） | あり（`writes:` で宣言） | あり（`writes:` で宣言） |
| brewprint 上の mutability | immutable な output snapshot | mutable | mutable |

## 理由

### なぜ「決定する」ではなく「明文化する」のか

asset の immutability は ADR-010 / ADR-020 / ADR-063 §7 が確立した時点で構造的に決まっていた。cross-edge の対象が store-id のみという事実が、asset の書き換え文法の不在を既に強制している。

本ADRが新たに導入する制約は無く、既存 edge 体系の帰結を明文化するだけである。よって supersedes は持たない。

### なぜ明文化が必要か

構造的事実だけでは、設計者・LLM が asset の扱いについて誤解する余地が残る。具体的には：

- LLM が「asset を append する task」を生成する記述パターンを抑止できない
- 設計レビューで「asset を mutate していいのか」という議論が無駄に発生する
- ADR-063 が initialized store の mutation を前提化したことで、asset 側の非対称性が暗黙化した

明文化することで、asset と store の役割対比を edge 体系から説明する基準点が生まれる。

### なぜ asset 側に reads/writes を導入しないか

仮に asset を mutable な実体とし cross-edge の対象に含めると、ADR-010 の model/asset 分離が崩れる。ADR-010 §決定 §4 では asset を「task の returns 宣言から暗黙的に生まれるフロー上の存在」と定義しており、独立ファイルを持たない。これは「asset の本体は task が単一実行で生む値である」という前提に立っている。

ここに「複数 task が後から書き換える対象」というセマンティクスを追加すると、asset は「フロー上の存在」を超えて「持続する runtime instance」になる。それは store の責務である。

mutable な runtime instance が必要な場合に store（特に initialized store）を選ぶ判断軸を維持するためにも、asset 側には書き込み文法を導入しない。

### なぜ「実装言語上の挙動」をスコープ外とするか

ADR-063 §3 と同じ理由。brewprint は YAML 語彙レベルの設計言語であり、Go / Python の参照渡しや mutation の言語仕様まで規定すると spec が際限なく複雑化する。語彙レベルで「mutation を表現する文法を持つか持たないか」を規定すれば、設計表現としての contract は十分に成立する。

### 却下した代替案

#### 代替案A: asset を mutable と定義し cross-edge の対象に含める

ADR-010 の model/asset 分離（asset = フロー上の存在、独立ファイルなし）を崩す。asset と store の責務境界が曖昧になり、ADR-007 supersede 時に整理した「型と実体の混在を防ぐ」原則に逆行する。

→ 却下。

#### 代替案B: ADR を立てず spec/nodes.md に1行追記するだけにする

「asset は immutable」と spec に書くだけでは、なぜそうなのか（edge 体系からの帰結）が伝わらない。設計者・LLM が「規約として決まった」と理解してしまうと、将来の議論で「変えていい？」が再発する。

ADR として残すことで「現行 edge 体系から構造的に導かれている」という根拠を保存できる。

→ 却下。ADR + spec 両方に反映する。

#### 代替案C: ADR-063 に追記する形で対応する

ADR-050 §3 「ADR は遡及修正しない」原則に反する。ADR-063 は accepted の起票時点スナップショットとして残すべきであり、追記による変更は行わない。

→ 却下。新規 ADR-065 として起票する。

#### 代替案D: 「brewprint 上 immutable」と言い切らず「規定しない」と非規定化する

asset の mutability を「規定しない」と書くと、設計者が「規定がないから mutate していい」と誤読する余地が残る。edge 体系の構造的事実から積極的に「immutable な output snapshot として扱う」と書く方が、設計表現としての contract が明確になる。

→ 却下。構造的事実の明文化として「immutable」と明示する。

## 影響

### spec への影響

- `docs/spec/nodes.md` の §asset セクションに、構造的 immutability と store との役割対比を追記する
- `docs/spec/nodes.md` の §store セクション冒頭に、store の2形態（store node / initialized store）の宣言経路・スコープ表を追加する
- `docs/spec/overview.md` の §ノード種別 末尾に、runtime data instance 総称としての store と、asset との役割対比を1段落追加する

### 既存 ADR への影響

- ADR-010 / ADR-014 / ADR-020 / ADR-063 はいずれも遡及修正しない。本ADRは既存体系の帰結を明文化する補追である
- supersedes はなし

### 既存実装への影響

- 実装変更なし。asset を cross-edge の対象に含めない既存の resolver / validator 実装はそのまま正しい
- 将来「asset を writes に書く」記述が現れた場合は既存の `unknown_store_id` 系の diagnostic で error になる（asset id は store-id 名前空間に存在しないため）

### 既存 UC への影響

- 既存 UC の YAML 記述は変更不要

### v1.1 への影響

本ADRは v1.0.0-spec / v1.1.0-spec の構造的帰結の明文化であり、forward 拡張ではない。v1.1.0-spec の凍結対象に含める。

## Evidence

- commit: tbd
- impl commit: 該当なし（既存 edge 体系の明文化のため実装変更なし）
- 参考: ADR-010 model/asset 分離, ADR-020 cross-edge management, ADR-063 task return source / initialized store
