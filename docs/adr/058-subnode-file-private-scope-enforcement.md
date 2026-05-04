# 058: サブノードのファイル内private性を実装で正しく保証する

- **status**: accepted
- **date**: 2026-05-02

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

UC-002（brewprint self-hosting / M14）Phase Aで MCP tool task群をblueprint化したところ、
複数ファイルが同名サブtaskを定義した場合に `brewprint render` が `duplicate_node` エラーで失敗することが判明した。

具体例: `mcp/task/get_references.yaml` と `mcp/task/get_signature.yaml` がそれぞれ
`validate_request` / `query_service` / `build_response` という同名のサブtaskを持つと、
`duplicate_node` エラーと、それに連鎖した `unresolved_flow_task` エラーが大量に出る。

これは ADR-011 §3 / spec/nodes.md「ファイル構造」が定める**サブノードはファイル内private**という規定と矛盾する。
サブノードがファイル内privateであるならば、別ファイルにある同名サブノードは衝突しないはず。

### 仕様側の規定（再掲）

- ADR-011 §3 サブノードの可視性スコープ: 「サブノードはファイル内にprivate。外部モジュールからの参照不可」
- spec/nodes.md ファイル構造節: 「サブノードはファイル内private。外部モジュールから参照不可」

### 実装側の現状

`internal/resolve/symbols.go` の `addNode` がサブノードに対しても
`module.kind.id` 形式の QualifiedID を組み立て、project全体の単一マップ `NodesByQID` に登録する。
この時点で別ファイルの同名サブノードと衝突し `duplicate_node` を出す。

flow解決ロジック（`internal/resolve/flow.go` の `resolveNodeQID`）自体は
`NodesByFile[fileID]` を優先走査するファイルスコープ実装になっており、
reference / impact layer にも `isPrivateSubNode` / `PrivateNodeID` 補助関数が存在する。
つまり**ファイル内private性を意図した実装は部分的に存在するが、addNode 段階の duplicate判定がproject全体スコープで先に走るため貫徹していない**。

### 顕在化のタイミング

UC-001（EC Checkout Flow）では「複数ファイルが同名サブtaskを持つ」パターンが偶然存在しなかったため、
M0〜M13までこのバグは顕在化しなかった。
ADR-057でUC-001を v1.0.0-spec のcanonical fixtureと固定したが、UC-001だけでは仕様の一部側面が未カバーだったことがUC-002 Phase Aで露呈した形である。

## 決定

### 1. サブノードはQualifiedIDの一意性制約から除外する

サブノードはファイル内privateであり、外部から QualifiedID 参照されない。
よって project全体での QualifiedID 一意性制約（`duplicate_node`）の対象から外す。

メインノードは従来通り project全体で QualifiedID 一意とする。

### 2. サブノードの内部識別

サブノードもrender / queryからは識別する必要があるため、内部識別子は持つ。
具体的なデータ構造はimplにて確定する（後述「実装方針」参照）。

外部参照可能な QualifiedID は引き続きメインノードのみ。
サブノードを指す QualifiedID は MCP の公開contract / spec / 他ファイルから書ける YAML 上の参照には**登場しない**。

### 3. flow / reads / writes 等のファイル内参照はファイルスコープで解決する

ファイル内に書かれた `step: <id>` / `reads: [<store>]` / `writes: [<store>]` 等の bare ID は、
そのファイル内のサブノード（task / store / branch / fork / join 等）を最優先で解決する。
ファイル内に該当ノードが見つからない場合のみ、同一moduleのメインノードへフォールバックする。

この優先順は現行実装（`resolveNodeQID` の `NodesByFile` 優先走査）と整合する。
本ADRは仕様上もこの解決順を明示する。

### 4. 仕様側更新

本ADR受理後、以下のspecを更新する。

- **`docs/spec/nodes.md` ファイル構造節**: 「サブノードはファイル内private」を以下まで強化する
  - サブノードIDは同一ファイル内で一意
  - 別ファイルにある同名サブノードとは衝突しない
  - サブノードは QualifiedID で外部参照されない
- **`docs/spec/diagnostics.md` の `duplicate_node` 説明**: メインノード間衝突に限定する旨を明記
- **`docs/spec/naming.md`**: メインノード = 外部参照可能な QualifiedID 保持者 / サブノード = ファイル内private、という対応を §2 周辺に追記

### 5. 実装方針（参考）

ADR-050 のspec-first方針上、実装詳細はADRの拘束事項ではない。以下は方針メモ:

- `internal/resolve/symbols.go` の `addNode` を分岐させ、`IsMain() == true` のみ `NodesByQID` の duplicate チェック対象とする
- サブノードは `NodesByFile[fileID]` への登録は維持する（flow / reads / writes 解決はここから引く）
- サブノードを QID 経由で引く必要がある場面（render / query）は、ファイルスコープの内部キー（例: `(FileID, kind, id)`）でアクセスする経路を新設する
- 既存の `NodesByQID` ベースのreference / impactロジックは、サブノードを混在させない方向で整理する
- 既に存在する `isPrivateSubNode` / `PrivateNodeID` 補助関数は、上記整理の延長として活用する

具体構造はM14a（後続milestone）のtaskとして決定する。

## 理由

### サブノードを QualifiedID 一意性から外す根拠

ADR-011がそもそも「外部参照不可」と定義したサブノードが、
QualifiedID という外部参照のための識別子を持って一意性を競うのは設計矛盾である。
ファイル内privateを言葉どおり扱うのが素直。

### ファイル内ID直書きの解決順を明示する根拠

現行実装はファイルスコープを優先する。
ただし spec / ADR にはこの優先順が明示されておらず、
今後別レイヤー（query layer等）でフラット解決と取り違える事故を防ぐため明文化しておく。

### specを更新する根拠

ADR-011 / spec/nodes.md の現状記述「サブノードはファイル内private」は意図としては正しいが、
「同名サブノードが別ファイルにあっても衝突しない」までは読み切れない曖昧さが残っていた。
spec-first運用（ADR-050）に従い、現行仕様の唯一の正であるspecを強化する。

ADR-011本体は遡及修正しない（doc-policy.md §3 「ADRの記述は遡及修正しない」）。
本ADRが ADR-011 の意図を引き継ぎつつ実装側の整合を取るための補強ADRとなる。

### 却下した代替案

#### 代替案A: サブノードに `module.file.kind.id` 形式の QualifiedID を与えてフラット一意性を維持する

- サブノードは外部参照不可なので、外部参照用の QualifiedID を持たせる必然性がない
- LLMエルゴ優先（doc-policy.md §1）に反する。サブノードを表すIDが実質4階層になりLLMが混乱する
- 既存の reference / impact ロジックがサブノードまで巻き込んでgraph traversalする結果になり、外部表面積が広がる

#### 代替案B: YAML側でサブノードIDをファイル名でprefixする運用回避

- 例: `gr_validate_request` / `gs_validate_request`
- 仕様上認められた書き方を回避する形になり、ADR-011のLLMエルゴ目的に反する
- self-hostingでこの回避運用が必要になるなら、UC-002自体が v1 の表現力不足を示すことになるが、実際にはspec通り書けるべきもの

## 影響

### 既存実装への影響

- `internal/resolve/symbols.go` の `addNode` 修正が必要
- 関連テスト（`internal/resolve/validation_test.go` 等）にサブノード重複ケースを追加
- query / mcp layerでサブノード走査経路がある場合、整合性を確認する

### 既存UCへの影響

- UC-001: 影響なし（同名サブノードが別ファイルに存在するパターンが現存しない）
- UC-002: 本修正は UC-002 Phase A renderの前提条件を満たすが、UC-002 Phase A YAMLは
  M15（data layer expressiveness v1.1）完了後に enum / discriminated object / inline struct
  を使った形で再構築する方針となったため、本ADR完了直後のUC-002 Phase A 直接render検証は
  M14aのスコープに含めない。詳細は M14a / 後続 ADR 群を参照

### v1.0.0-spec タグへの影響

`v1.0.0-spec` 凍結条件（ADR-057 §2）の判定基準のうち「`go test ./...` パス」は
M14a で本修正を入れる際に**バグ修正テストを追加した状態**で再度満たす必要がある。

ただしADR-057の凍結基準の意図は「凍結時点での品質ベースライン」であり、
本ADRはそのベースラインを覆すものではなく**spec通りに実装が動くよう補強する**ものである。
ADR-050 §7 / ADR-057 Non-goals が禁じる「v1範囲のspec / ADRの遡及修正」には該当しない。

`v1.0.0-spec` タグ自体はすでに発行済みであり、本ADRおよびM14aの修正は次の `v1.0.1-spec`
（patch release: bug fix のみ）として扱う。
ADR-057 §4 のRelease snapshots運用に従い、本ADR受理時点では新タグは切らない。

### M14への影響

- M14（self-hosting）は本ADR起票時点で paused
- M14a (B1: subnode scope + B2: return primitive) 完了で v1.0.1-spec タグ発行
- その後 M15（data layer expressiveness v1.1: enum / discriminated object / inline struct）完了で v1.1.0-spec タグ発行
- M14 Phase A YAML は M14a + M15 完了後に enum等を使った形で初構築する

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: GoおよびRustの「1ファイル=1public型+privateヘルパー」構造慣習
