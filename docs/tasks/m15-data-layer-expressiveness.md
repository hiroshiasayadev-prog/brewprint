# Milestone 15: data layer expressiveness (v1.1)

- **status**: open
- **scope**: internal/semantic / internal/resolve / spec / docs/adr / tests
- **source**: ADR-060 (TypeRef + flow wiring type compatibility) / ADR-061 (foreach.returns collected asset) / ADR-062 (task return source) / ADR-063 (initialized source の wiring source 化) / ADR-064 (returns.source の DAG render ルール, proposed) を起点とする v1.1 系の表現力拡張
- **last_updated**: 2026-05-05

---

## Context

M14a (ADR-058 / ADR-059) で v1.0 系の実装と spec の乖離を解消した時点で、
brewprint の型表現力には以下の制約が残っていた。

- flow wiring における型互換性ルールが存在しない（ADR-060 §背景）
- `any` の代入互換性が未定義
- `list<T>` / `dict<T>` 相当の shallow container を named model なしに表現できない
- `foreach` の `$item` 型解決が spec 上明文化されていない（UC-001 `validate_cart.yaml` 内コメントに記録された spec gap）
- foreach.returns の collected asset を後続 flow から参照する場合の source id 解決ルールが未定義
- enum / discriminated object / inline struct など、UC-002 self-hosting で必要になる data layer の表現力が不足

ADR-060 はこのうち **TypeRef 導入と flow wiring type compatibility** を v1.1 の基礎として確定した。
ADR-061 は **foreach.returns collected asset source** を Phase B として確定した。
ADR-062 は **task return source の明示化** を Phase B2 として確定した。
ADR-063 は **initialized source を returns.source および flow 内部 wiring の bare token source として参照可能にする** ことを Phase B3 として確定した。
ADR-064 は **`returns.source` / initialized source 参照の DAG render ルール** を Phase B4 で扱う（起票時点 proposed）。
M15 は ADR-060 / ADR-061 / ADR-062 / ADR-063 / ADR-064 を実装に落とし込みつつ、関連する v1.1 表現力拡張をまとめて扱う milestone である。

当初 M14b として独立 milestone を切る案もあったが、ADR-060 が v1.1 TypeRef 前提に拡張された結果、
M14b 独立の意義が薄れ、本 milestone（M15）に吸収する判断をした（ADR-060 §M15 への影響）。

---

## Non-goals

M15 では以下を行わない。

- **subtyping / structural typing / user-defined generics の導入**
  - ADR-060 §8 で明示的に却下。M15 もこの方針を踏襲する
- **variance rules の導入**
  - 同上
- **v1.0.0-spec タグの再発行 / v1.0 系 ADR の遡及修正**
  - ADR-050 §7 / ADR-057 Non-goals 準拠
- **`list<T>` / `dict<T>` 構文を v1.0 系に backport すること**
  - 本 milestone は v1.1.0-spec への前進であり、v1.0 系には影響しない
- **UC-002 self-hosting の完了**
  - M15 完了後に M14（self-hosting）を再開する想定。M15 自体は self-hosting の完了を目指さない

---

## Tag

- M15 完了時点で **`v1.1.0-spec`** タグを発行する（ADR-060 §タグ発行方針）
- M15 単独で v1.0.x patch を発行することはしない

---

## Phase 構成

M15 は以下の Phase に分けて進める。

- Phase A: TypeRef 導入 (ADR-060)
- Phase B: foreach.returns collected asset (ADR-061)
- Phase B2: task return source 明示化 (ADR-062)
- Phase B3: initialized source の wiring source 化 (ADR-063)
- Phase B4: returns.source の DAG render ルール (ADR-064, proposed)
- Phase C: enum / discriminated object など追加表現力

### Phase A: TypeRef + flow wiring type validation

ADR-060 を実装と spec に落とし込む。

#### 仕様反映

- [x] **`docs/spec/nodes.md` または新規 `docs/spec/type-ref.md` に TypeRef 構文を追加**
  - 構文形式: primitive / named model / inline `list<T>` / inline `dict<T>`
  - 再帰的定義の許容範囲（`list<dict<user>>` 等の入れ子）
  - 深さ制限または lint 方針は Phase C で扱うため、本 phase では「制限は M15 Phase C で扱う」とのみ明記
  - TypeRef を受け取るフィールド: `param.model` / `returns.model` / `field.type` / `model.element` / `model.value`
  - 由来: ADR-060 §1, §9

- [x] **`docs/spec/edges.md` §1（flow:セクション）末尾に「§型互換性ルール」節を追加**
  - ADR-060 §3 のルールを spec として記述
  - ADR-060 §4 の検証対象 wiring 一覧
  - ADR-060 §5 の wiring source 型解決ルール（node ID / `$params.<name>` / `$item`）
  - ADR-060 §7 の型解決失敗時の抑制ルール
  - 由来: ADR-060 §3〜§7, §9

- [x] **`docs/spec/edges.md` §1-5 foreach.over の記述を ADR-060 と整合させる**
  - 現状: 「`$params.field` を指定した場合、parser/validator は `main.params.<field>.model` が `kind: list` であることを検証する」
  - ADR-060 後: 「`foreach.over` の解決結果が list を表す TypeRef でない場合 `invalid_foreach_over_type` を出す」
  - UC-001 `validate_cart.yaml` コメントに記録された spec gap（`foreach.over` に `$params.field` を指定可能か明文化されていない件）を併せて解消
  - 由来: ADR-060 §5-3

- [x] **`docs/spec/diagnostics.md` に新 diagnostic 4件を追加**
  - `incompatible_wiring_type` (severity: error)
  - `invalid_wiring_source` (severity: error)
  - `invalid_foreach_over_type` (severity: error)
  - `invalid_type_ref` (severity: error)
  - 各 diagnostic の発生条件・メッセージフォーマットを記述
  - 由来: ADR-060 §6, §9

#### 実装

- [ ] **`internal/semantic` に TypeRef 表現を追加**
  - 仮称: `semantic.TypeRef` 型。primitive / named model / inline list / inline dict の 4 variant
  - `Param` / `Return` / `ModelField` / `Model.Element` / `Model.Value` から TypeRef を引けるようにする
  - 既存の `ModelName` / `Model` (QualifiedID) は migration 期間中は保持してよい
  - 由来: ADR-060 §影響「既存実装への影響」

- [ ] **rawyaml / semantic で `list<T>` / `dict<T>` 構文をパース**
  - 既存の `ModelName` (raw string) パース箇所を拡張
  - パースエラー時の diagnostic は `invalid_type_ref` とする
  - 由来: ADR-060 §1, §6

- [ ] **既存の named list/dict model 解決ロジックを TypeRef に正規化**
  - `kind: list, element: T` → 内部表現として inline `list<T>` と同じ container shape を引けるようにする
  - `kind: dict, value: T` → 同様
  - 名前 (id) と note は意味付けとして保持
  - 由来: ADR-060 §2

- [ ] **`internal/resolve/validation.go` に `validateFlowWiringTypes` を追加**
  - validate phase で実行
  - ADR-060 §3 の互換ルール、§5 の source 型解決、§7 の抑制ルールを実装
  - 由来: ADR-060 §影響「既存実装への影響」

- [ ] **TypeRef 互換判定ヘルパーを追加**
  - 仮称: `typeRefsCompatible(a, b semantic.TypeRef) bool`
  - 由来: ADR-060 §3

- [ ] **diagnostic コード一覧に新 4 件を追加**
  - `internal/resolve/diagnostics.go` 等に `incompatible_wiring_type` / `invalid_wiring_source` / `invalid_foreach_over_type` / `invalid_type_ref` を追加
  - 由来: ADR-060 §6

#### テスト

- [ ] **互換ケースの単体テスト**
  - 同一 primitive / 同一 named model / any wildcard 両方向 / list 同士の element 互換 / dict 同士の value 互換 / named list ⇔ inline list の正規化互換
  - 由来: ADR-060 §3

- [ ] **非互換ケースの単体テスト**
  - 異なる primitive / 異なる named model / primitive と named model の混在 / list 同士で element 不一致 / dict 同士で value 不一致
  - 由来: ADR-060 §3

- [ ] **`$item` 解決のテスト**
  - over が node ID で list を返す場合 / over が `$params.field` で list を指す場合 / over が any の場合 / over が list でない場合 (invalid_foreach_over_type) / over の解決自体が失敗する場合
  - 由来: ADR-060 §5-3

- [ ] **`$item` 有効範囲のテスト**
  - foreach.params 外で `$item` を指定した場合に `invalid_wiring_source` が出る
  - 由来: ADR-060 §5-3

- [ ] **重複診断抑制のテスト**
  - source が unresolved の場合 incompatible_wiring_type が出ない
  - target param model が unresolved の場合 incompatible_wiring_type が出ない
  - foreach.over が invalid_foreach_over_type の場合、内側の wiring に incompatible_wiring_type が出ない
  - 由来: ADR-060 §7

- [ ] **UC-001 回帰テスト**
  - 既存の wiring がすべて新ルールで OK と判定されることを確認
  - 由来: ADR-060 §影響「UC-001」

---

### Phase B: foreach.returns collected asset 参照ルール

ADR-061 は accepted 済み。
foreach の collected asset を後続 flow から参照する場合のルールを spec に反映し、実装へ落とし込む。

#### 確定した仕様範囲

- `foreach.returns` は apply 先 task の `returns` を iteration ごとに collect した collected asset source 名である
- `foreach.returns` は optional。side-effect only の foreach では省略できる
- 省略時は collected asset source を semantic model に生成しない。renderer / inspect / MCP も internal pseudo source を露出しない
- apply 先 task に `returns` がないにもかかわらず `foreach.returns` が指定された場合は `invalid_foreach_returns`
- `foreach.returns` で宣言された collected asset は、同一 flow file 内の後続 step / branch / fork / foreach から bare wiring source として参照できる
- 当該 foreach 自身の `params` 内から自分の `returns` 名を参照した場合は `invalid_foreach_returns`
- collected asset source の TypeRef は apply 先 task の `returns.model` `T` から `list<T>` として導出する
- apply 先 task の `returns.model` が `any` の場合は `list<any>`
- apply 先 task の `returns.model` が解決不能な場合、collected asset source の TypeRef も解決不能として扱い、後続 wiring の `incompatible_wiring_type` は抑制する
- `foreach.returns` は同一 flow file 内の bare wiring source 名前空間に参加し、node id / 他の `foreach.returns` と重複してはならない。重複時は `duplicate_flow_source`
- task の `returns.name` と `foreach.returns` が同名でも衝突扱いしない
- ADR-023 の制御フロースコープは維持し、`foreach.returns` は collect 結果だけを外部 source として公開する escape hatch とする
- `foreach.id` は導入しない
- `task.returns.source` および main task `returns.name` と `foreach.returns` の名前一致による暗黙接続は本 Phase では扱わない。ADR-062 領域として未実装のままにする

#### 仕様反映

- [x] **ADR-061 を accepted として確定**
- [x] **`docs/spec/edges.md` に `foreach.returns` collected asset source 仕様を反映**
  - `foreach.returns` の意味、optional、省略時の扱い
  - 後続 flow からの bare source 参照
  - TypeRef = `list<T>` 導出
  - source 名前空間・重複ルール
  - 制御フロースコープの escape hatch
  - `foreach.id` 非導入、task return source 非対象
- [x] **`docs/spec/diagnostics.md` に ADR-061 diagnostics を反映**
  - `duplicate_flow_source`
  - `invalid_foreach_returns`
  - `unresolved_wiring_source`
  - `invalid_wiring_source` との区別

#### 実装

- [ ] **collected asset source を semantic / flow resolver に登録**
- [ ] **`duplicate_flow_source` を検出**
  - `foreach.returns` と同一 flow file 内 node id の重複
  - `foreach.returns` 同士の重複
  - task `returns.name` との同名は衝突扱いしない
- [ ] **`invalid_foreach_returns` を検出**
  - apply 先 task に `returns` がないのに `foreach.returns` が指定されている
  - 当該 foreach 自身の `params` 内から自分の `returns` 名を参照している
- [ ] **`unresolved_wiring_source` を追加**
  - node id / `$params.<name>` / `$item` / collected asset source のいずれとしても解決できない wiring source を診断する
  - `invalid_wiring_source` は「参照先は存在するが、その文脈では source として使えない」場合に限定する
- [ ] **collected asset source TypeRef = `list<T>` を導出**
  - apply 先 task `returns.model` が `T` の場合は `list<T>`
  - `any` の場合は `list<any>`
  - `returns.model` が解決不能な場合は collected asset source TypeRef も解決不能として扱う
- [ ] **後続 wiring source として collected asset を解決**
  - 同一 flow file 内の後続 step / branch / fork / foreach から bare source として参照可能にする
  - 自分自身の `foreach.params` から自分の `returns` 名を参照することは許可しない
- [ ] **`foreach.returns` 省略時は source を生成しない**
  - semantic model に collected asset source を作らない
  - renderer / inspect / MCP に internal pseudo source を露出しない
- [ ] **ADR-062 領域の task return source は未実装のままにする**
  - `task.returns.source` は追加しない
  - main task `returns.name` と `foreach.returns` の名前一致による暗黙接続は実装しない

#### テスト

- [ ] **collected asset source 解決テスト**
  - 後続 step / branch / fork / foreach から `foreach.returns` を参照できる
  - `foreach.returns` の TypeRef が `list<T>` として target param と互換判定される
- [ ] **省略時のテスト**
  - `foreach.returns` 省略時に source が生成されない
  - 省略時の internal pseudo source が renderer / inspect / MCP に露出しない
- [ ] **diagnostic テスト**
  - `duplicate_flow_source`
  - `invalid_foreach_returns`
  - `unresolved_wiring_source`
  - `invalid_wiring_source` との区別
  - TypeRef 解決不能時に `incompatible_wiring_type` が抑制される

---

### Phase B2: task return source の明示化

ADR-062 は accepted 済み。
`task.returns.source` によって、task の外向き return signature と内部 flow / input の source を明示的に接続する。

#### 確定した仕様範囲

> ADR-063 により initialized source も追加される。詳細は Phase B3 を参照する。下記は ADR-063 反映後の現行仕様である。

- `task.returns.source` は optional field である
- `returns.name` / `returns.model` は task signature として維持する
- `returns.source` は、その signature を満たす値を内部 flow / input からどこで得るかを指定する task return wiring である
- leaf task / note-only task / external boundary task では `returns.source` を省略できる
- `flow:` を持たない task でも `returns.source: $params.<name>` による pass-through return は正当である
- `returns.source` に指定できる source は node id / QualifiedID、collected asset source、initialized source、`$params.<name>` である
- `$item` は `returns.source` では使えない。指定時は `invalid_return_source`
- `returns.source` と `returns.model` は ADR-060 の TypeRef compatibility で検証する
- 型が互換しない場合は `incompatible_return_type`
- source / target TypeRef が解決不能な場合は `incompatible_return_type` を抑制する
- `returns.name` と flow source 名が一致していても暗黙接続は行わない
- 既存の `join.params` の `returns.name` 一致による暗黙接続は維持する
- `returns.source` は単一 return にのみ対応する。複数 return / named tuple / multi output task は扱わない

#### 仕様反映

- [x] **ADR-062 を accepted として確認**
- [x] **`docs/spec/nodes.md` の `returns` オブジェクトに `source` を追加**
  - optional field として定義
  - node id / collected asset source / `$params.<name>` を指定可能
  - `$item` は指定不可
  - `returns.name` と flow source 名の暗黙接続は行わないことを明記
- [x] **`docs/spec/edges.md` に task return wiring 節を追加**
  - source 解決規則
  - TypeRef compatibility
  - `$item` 禁止
  - 暗黙接続なし
  - `join.params` 暗黙接続は既存仕様として維持
- [x] **`docs/spec/diagnostics.md` に ADR-062 diagnostics を反映**
  - `unresolved_return_source`
  - `invalid_return_source`
  - `incompatible_return_type`
  - `unresolved_return_source` と `invalid_return_source` の区別
  - TypeRef 解決不能時の `incompatible_return_type` 抑制

#### 実装

- [ ] **raw YAML / semantic model の `returns` に `source` field を追加**
  - `internal/rawyaml.Return.Source string`
  - `internal/semantic.Return` に raw source 文字列または resolved source を保持する field を追加
  - 既存の `returns.name` / `returns.model` / `TypeRef` / `Asset` は維持
- [ ] **builder / resolver で `returns.source` を semantic model に反映**
  - raw `returns.source` を読み取る
  - source 記法は flow wiring source と同じ構文を使う
  - `$item` は構文上は保持してよいが validation で `invalid_return_source` にする
- [ ] **task return source resolver を追加**
  - node id / QualifiedID を task / join returns に解決
  - `$params.<name>` を main task params に解決
  - collected asset source を同一 flow file 内の `foreach.returns` として解決
  - `returns.source` が未指定の場合は何もしない
- [ ] **task return source は flow 末尾時点の collected source 集合で解決**
  - `returns.source` は task 全体の return を表すため、flow entry 順における前方参照という概念は適用しない
  - 同一 flow file 内に出現するすべての `foreach.returns` 由来 collected asset source を参照候補にする
  - flow 内部 wiring の entry順 visibility とは扱いを分ける
- [ ] **return source diagnostic を追加**
  - `unresolved_return_source`
  - `invalid_return_source`
  - `incompatible_return_type`
- [ ] **return source TypeRef compatibility を検証**
  - source TypeRef と `returns.model` TypeRef を ADR-060 ルールで比較
  - named list/dict model と inline `list<T>` / `dict<T>` の正規化互換を使う
  - `any` は両方向 wildcard
  - source / target TypeRef 解決不能時は `incompatible_return_type` を抑制
- [ ] **name一致による暗黙 return 接続を実装しないことを確認**
  - `returns.name` と flow source / `foreach.returns` が一致しても、`returns.source` 未指定なら return wiring として扱わない
- [ ] **MCP / renderer / resolved index への露出方針を確認**
  - `returns.source` を MCP `get_signature` に含めるか、`inspect` のみに出すかを判断する
  - DAG renderer で `returns.source` を可視化するか、ADR-062 ではスコープ外とするかを判断する
  - semantic / resolved project に raw source 文字列だけを持つか、resolved source も持つかを ADR-048 と整合確認する
- [ ] **ADR-062 実装後に Evidence を更新**
  - `docs/adr/062-task-return-source.md` の `impl commit: tbd` を実装 commit hash で更新
  - `docs/tasks/m15-data-layer-expressiveness.md` の Evidence は必要に応じて追補する

#### テスト

- [ ] **`returns.source` 正常系テスト**
  - node output を task return source として返せる
  - collected asset source を task return source として返せる
  - initialized source を task return source として返せる
  - `$params.<name>` を pass-through return として返せる
  - `returns.source` 未指定の leaf task / external boundary task は既存通り valid
- [ ] **return source diagnostic テスト**
  - 未解決 source で `unresolved_return_source`
  - returns を持たない node で `invalid_return_source`
  - `$item` 指定で `invalid_return_source`
  - source / target 型不一致で `incompatible_return_type`
  - source TypeRef 解決不能時は `incompatible_return_type` を抑制
  - target `returns.model` 解決不能時は `incompatible_return_type` を抑制
- [ ] **collected asset visibility / 暗黙接続のテスト**
  - flow 内の `foreach.returns` は、flow entry 上の出現位置にかかわらず `returns.source` から参照できる
  - flow 内部 wiring では既存どおり entry順 visibility を維持する
  - `returns.name` と flow source 名が一致していても、`returns.source` 未指定なら return wiring として扱わない
- [ ] **UC-001 回帰更新**
  - `cart/task/validate_cart.yaml` で `foreach.returns: validated_items` を main task return として返す場合、`returns.source: validated_items` を明示する
  - 既存 render / validate が ADR-062 後も通ることを確認

---

### Phase B3: initialized source の wiring source 化

ADR-063 は accepted 済み。
`initializes[].name` を `returns.source` および flow 内部 wiring の bare token source として参照可能にする。

#### 確定した仕様範囲

- `returns.source` の指定可能 source に initialized source を追加する
- flow 内部 wiring（step.params / branch.params / fork.branches[].steps[].params / foreach.params / branch.cases[].params）の bare token source 種別に initialized source を追加する
- initialized source の TypeRef は `initializes[].model` を named model TypeRef として導出する。inline `list<T>` / `dict<T>` は `initializes[].model` では受け取らない
- `initializes[].model` が解決不能な場合、initialized source の TypeRef も解決不能として扱い、`incompatible_wiring_type` / `incompatible_return_type` を抑制する
- `returns.source` から参照する場合の評価時点は task 実行完了時点（flow END 時点）
- flow 内部 wiring から参照する場合、その source は flow 構造上の当該 wiring 到達時点の store 状態を指す。mutation semantics は task 実装に委ねる
- `initializes[].name` は同一 flow file 内の bare wiring source 名前空間に参加する。node id / `foreach.returns` / `initializes[].name` は重複不可。重複時は `duplicate_flow_source`
- task `returns.name` と initialized source 名 / `foreach.returns` 名が同名でも衝突扱いしない
- writes されていない initialized source の参照は valid。validation 対象としない
- cross-edge `reads` / `writes` と flow wiring 参照は分担する（flow param wiring は値の受け渡し contract、cross-edge は副作用 / store access contract）。両者の整合性検査は本 phase では実装しない
- ADR-014 の外部参照不可ルールは維持する。外部 file / module 跨ぎ参照は引き続き不可
- DAG render ルールは ADR-064（proposed）で扱う。本 phase では暫定方針（ADR-064 §暫定方針）に従う

#### 仕様反映

- [x] **ADR-063 を accepted として確認**
- [x] **`docs/spec/nodes.md` の `returns オブジェクト` の `source` 説明に initialized source を追加**
- [x] **`docs/spec/nodes.md` の `init オブジェクト` 末尾に bare wiring source 化を追記**
  - `initializes[].name` が bare wiring source として参照可能であること
  - `returns.source` および flow 内部 wiring から参照可能であること
  - TypeRef は `initializes[].model` を named model TypeRef として扱うこと
  - 外部 file / module 跨ぎ参照は ADR-014 通り引き続き不可
- [x] **`docs/spec/edges.md` §1-7 wiring source の型解決表に initialized source を5種目として追加**
- [x] **`docs/spec/edges.md` §1-7 名前空間・重複ルール説明に initialized source を加える**
- [x] **`docs/spec/edges.md` §1-8 task return wiring の指定可能 source 表に initialized source を追加**
- [x] **`docs/spec/edges.md` §1-8 に評価時点（task 実行完了時点 / flow END 時点）を明記**
- [x] **`docs/spec/diagnostics.md` の `unresolved_return_source` / `unresolved_wiring_source` / `duplicate_flow_source` の意味を拡張**

#### 実装

- [ ] **flow wiring source resolver に initialized source 解決を追加**
  - 同一 file の main task の `initializes[]` を引き、`name == <bare_token>` の entry の `model` を named model TypeRef として返す
  - bare token 解決順序は node id / collected asset source / initialized source の3者を統一名前空間として扱う
- [ ] **task return source resolver に initialized source 解決を追加**
  - flow wiring source resolver と同じ initialized source 解決ロジックを返す
- [ ] **`duplicate_flow_source` の検出対象に initialized source 名を加える**
  - 同一 flow file 内の node id / collected asset source / initialized source の重複を検出
- [ ] **`unresolved_wiring_source` / `unresolved_return_source` の解決失敗判定を更新**
  - bare token が node id / collected asset source / initialized source のいずれにも該当しない場合に出す
- [ ] **TypeRef compatibility を initialized source にも適用**
  - `initializes[].model` を named model TypeRef として扱い ADR-060 ルールで検証
  - `initializes[].model` 解決不能時は `incompatible_wiring_type` / `incompatible_return_type` を抑制
- [ ] **writes 有無の検査は実装しない**
- [ ] **cross-edge `reads` / `writes` 宣言と flow wiring 参照の整合性検査は実装しない**
- [ ] **DAG render は ADR-064 §暫定方針に従う**
  - `returns.source` の DAG 上明示 edge は引かない
  - initialized source は file-private store 表現を維持
  - flow 内部 wiring からの initialized source 参照 edge は通常の dataflow edge として描く
- [ ] **ADR-063 実装後に Evidence を更新**
  - `docs/adr/063-task-return-source-initialized-store.md` の `impl commit: tbd` を実装 commit hash で更新

#### テスト

- [ ] **flow 内部 wiring から initialized source を bare token で参照できる**
  - step.params / branch.params / fork.branches[].steps[].params / foreach.params / branch.cases[].params 各箇所
  - TypeRef compatibility が ADR-060 ルールで通る
- [ ] **`returns.source` から initialized source を参照できる**
  - main task が initialized store を return する pattern が valid
  - TypeRef が `returns.model` と互換しない場合は `incompatible_return_type`
- [ ] **重複検出テスト**
  - node id と initialized source 名の重複で `duplicate_flow_source`
  - collected asset source 名と initialized source 名の重複で `duplicate_flow_source`
  - task `returns.name` と initialized source 名が同名でも衝突扱いしない
- [ ] **未解決テスト**
  - typo した bare token で `unresolved_wiring_source` / `unresolved_return_source`
- [ ] **TypeRef 解決抑制テスト**
  - `initializes[].model` が解決不能なときの `incompatible_wiring_type` / `incompatible_return_type` 抑制
- [ ] **writes 有無無関係テスト**
  - flow 内で writes されていない initialized source を参照しても valid
- [ ] **UC-001 回帰**
  - 既存 wiring に initialized source を bare token で参照しているものはない見込みだが、回帰として `process_report.yaml` 等で composite task の構造を検証

---

### Phase B4: returns.source の DAG render ルール

ADR-064 は proposed。
`returns.source` および initialized source 参照を Mermaid DAG 上でどう可視化するかを確定させる。

#### Tasks

- [ ] **UC-001 で `returns.source` のサンプルを揃える**
  - initialized source を `returns.source` に指定するサンプル（process_report 等）
  - collected asset source を `returns.source` に指定するサンプル（validate_cart）
  - `$params.<name>` を `returns.source` に指定するサンプル（pass-through task）
  - flow 内部 wiring から initialized source を参照するサンプル
- [ ] **ADR-064 §論点 各案を実 render で比較**
  - 論点1（edge 表現）の各案
  - 論点2（ID 衝突回避）の各案
  - 論点3（initialized source の DAG 表現）の各案
  - 論点4（collected asset return の edge 起点）の各案
  - 論点5（pass-through return の edge）の各案
  - 論点6（flow 内部 wiring からの initialized source edge）の各案
- [ ] **ADR-064 を accepted に確定**
  - 比較結果から採用案を決め、ADR-064 を accepted へ
- [ ] **`docs/spec/views/dag.md` に `returns.source` render ルールを反映**
- [ ] **DAG renderer に採用案を実装**
- [ ] **golden test を更新**

---

### Phase C: enum / discriminated object / inline struct / 深さ制限

UC-002 self-hosting で必要になる data layer 表現力を ADR ベースで導入する。

#### 検討範囲

- **enum**: 限定された値の集合を持つ primitive 型（status, role 等）
- **discriminated object** (tagged union): 複数のmodel variant を識別フィールドで切り替える型
- **inline struct**: model として外部に定義しない、param/return内のad-hoc struct
- **`list<T>` / `dict<T>` の深さ制限または lint 方針** (ADR-060 §1 後段)

#### Tasks

- [ ] **enum 導入の要否判断 ADR を起票**
  - UC-002 で `status` / `role` 等の enum 候補を観測
  - ADR としては proposed → 議論 → accepted の流れ

- [ ] **discriminated object 導入の要否判断 ADR を起票**
  - 同上

- [ ] **inline struct 導入の要否判断 ADR を起票**
  - 同上

- [ ] **TypeRef 深さ制限 / lint 方針 ADR を起票**
  - `list<dict<list<...>>>` のような深い入れ子をどう扱うか
  - lint レベルの diagnostic か、構文エラーか

- [ ] **各 ADR acceptance 後、spec / 実装に反映**
  - Phase A の TypeRef 表現に追加 variant を足す形になる想定

---

## 実装方針

### Phase 進行順

- Phase A から着手（ADR-060 確定済み）
- Phase A 完了後、UC-002 self-hosting を再開しながら Phase B / C の必要性を観測
- Phase B / C は ADR 起票 → 議論 → accepted を都度行う

### commit 境界

- ADR-060 spec 反映 → 1 commit
- TypeRef semantic 表現追加 → 1 commit
- パーサ拡張 → 1 commit
- validateFlowWiringTypes 実装 + テスト → 1 commit
- 各 ADR-061 / Phase C 各 ADR は別 commit

### review 単位

- Phase A 完了時点でレビュー実施（ADR-060 と実装の整合確認）
- Phase B / C は各 ADR ごとにレビュー

---

## 関連 ADR

- ADR-060: v1.1 TypeRef と flow wiring の型互換性 (accepted)
- ADR-061: foreach.returns collected asset 参照ルール (accepted)
- ADR-062: task return source の明示化 (accepted)
- ADR-063: task return source への initialized source 追加と initializes の wiring source 化 (accepted)
- ADR-064: returns.source の DAG render ルール (proposed)
- ADR-014: initializes フィールド設計（v1.1 で wiring source 対応に拡張、ADR-063）
- ADR-020: cross-edge management（v1.1 で flow wiring との分担を ADR-063 §7 で整理）
- ADR-021: primitive 予約語と list/dict model（v1.1 でも維持）
- ADR-016: foreach 構文（変更なし。`$item` 型解決ルールは ADR-060 で明文化）
- ADR-058 / ADR-059: M14a で確定した v1.0 系の実装バグ修正
- ADR-050: spec-first documentation policy
- ADR-057: self-hosting 思想 / Non-goals

---

## Evidence

- commit: 01e7127
- impl commit: 01e7127
