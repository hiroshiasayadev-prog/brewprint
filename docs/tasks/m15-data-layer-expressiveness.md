# Milestone 15: data layer expressiveness (v1.1)

- **status**: open
- **scope**: internal/semantic / internal/resolve / spec / docs/adr / tests
- **source**: ADR-060 (TypeRef + flow wiring type compatibility) / ADR-061 (foreach.returns collected asset) を起点とする v1.1 系の表現力拡張
- **last_updated**: 2026-05-04

---

## Context

M14a (ADR-058 / ADR-059) で v1.0 系の実裁E�� spec の乖離を解消した時点で、Ebrewprint の型表現力には以下�E制紁E��残ってぁE��、E
- flow wiring における型互換性ルールが存在しなぁE��EDR-060 §背景�E�E- `any` の代入互換性が未定義
- `list<T>` / `dict<T>` 相当�E shallow container めEnamed model なしに表現できなぁE- `foreach` の `$item` 型解決ぁEspec 上�E斁E��されてぁE��ぁE��EC-001 `validate_cart.yaml` 冁E��メントに記録されぁEspec gap�E�E- foreach.returns の collected asset を後綁Eflow から参�Eする場合�E source id 解決ルールが未定義
- enum / discriminated object / inline struct など、UC-002 self-hosting で忁E��になめEdata layer の表現力が不足

ADR-060 はこ�EぁE�� **TypeRef 導�Eと flow wiring type compatibility** めEv1.1 の基礎として確定した、EADR-061 は **foreach.returns collected asset source** めEPhase B として確定した、EM15 は ADR-060 / ADR-061 を実裁E��落とし込みつつ、E��連する v1.1 表現力拡張をまとめて扱ぁEmilestone である、E
当�E M14b として独竁Emilestone を�Eる案もあったが、ADR-060 ぁEv1.1 TypeRef 前提に拡張された結果、EM14b 独立�E意義が薄れ、本 milestone�E�E15�E�に吸収する判断をした！EDR-060 §M15 への影響�E�、E
---

## Non-goals

M15 では以下を行わなぁE��E
- **subtyping / structural typing / user-defined generics の導�E**
  - ADR-060 §8 で明示皁E��却下、E15 もこの方針を踏襲する
- **variance rules の導�E**
  - 同丁E- **v1.0.0-spec タグの再発衁E/ v1.0 系 ADR の遡及修正**
  - ADR-050 §7 / ADR-057 Non-goals 準拠
- **`list<T>` / `dict<T>` 構文めEv1.0 系に backport すること**
  - 本 milestone は v1.1.0-spec への前進であり、v1.0 系には影響しなぁE- **UC-002 self-hosting の完亁E*
  - M15 完亁E��に M14�E�Eelf-hosting�E�を再開する想定、E15 自体�E self-hosting の完亁E��目持E��なぁE
---

## Tag

- M15 完亁E��点で **`v1.1.0-spec`** タグを発行する！EDR-060 §タグ発行方針！E- M15 単独で v1.0.x patch を発行することはしなぁE
---

## Phase 構�E

M15 は3つの Phase に刁E��て進める、EPhase A は ADR-060 で確定済み、EPhase B は ADR-061 で accepted 済み、EPhase C は追加 ADR の起票を伴ぁE��E
### Phase A: TypeRef + flow wiring type validation

ADR-060 を実裁E�� spec に落とし込む、E
#### 仕様反映

- [x] **`docs/spec/nodes.md` また�E新要E`docs/spec/type-ref.md` に TypeRef 構文を追加**
  - 構文形弁E primitive / named model / inline `list<T>` / inline `dict<T>`
  - 再帰皁E��義の許容篁E���E�Elist<dict<user>>` 等�E入れ子！E  - 深さ制限また�E lint 方針�E Phase C で扱ぁE��め、本 phase では「制限�E M15 Phase C で扱ぁE��とのみ明訁E  - TypeRef を受け取るフィールチE `param.model` / `returns.model` / `field.type` / `model.element` / `model.value`
  - 由来: ADR-060 §1, §9

- [x] **`docs/spec/edges.md` §1�E�Elow:セクション�E�末尾に「§型互換性ルール」節を追加**
  - ADR-060 §3 のルールめEspec として記述
  - ADR-060 §4 の検証対象 wiring 一覧
  - ADR-060 §5 の wiring source 型解決ルール�E�Eode ID / `$params.<name>` / `$item`�E�E  - ADR-060 §7 の型解決失敗時の抑制ルール
  - 由来: ADR-060 §3〜§7, §9

- [x] **`docs/spec/edges.md` §1-5 foreach.over の記述めEADR-060 と整合させる**
  - 現状: 「`$params.field` を指定した場合、parser/validator は `main.params.<field>.model` ぁE`kind: list` であることを検証する、E  - ADR-060 征E 「`foreach.over` の解決結果ぁElist を表ぁETypeRef でなぁE��吁E`invalid_foreach_over_type` を�Eす、E  - UC-001 `validate_cart.yaml` コメントに記録されぁEspec gap�E�Eforeach.over` に `$params.field` を指定可能か�E斁E��されてぁE��ぁE���E�を併せて解涁E  - 由来: ADR-060 §5-3

- [x] **`docs/spec/diagnostics.md` に新 diagnostic 4件を追加**
  - `incompatible_wiring_type` (severity: error)
  - `invalid_wiring_source` (severity: error)
  - `invalid_foreach_over_type` (severity: error)
  - `invalid_type_ref` (severity: error)
  - 吁Ediagnostic の発生条件・メチE��ージフォーマットを記述
  - 由来: ADR-060 §6, §9

#### 実裁E
- [ ] **`internal/semantic` に TypeRef 表現を追加**
  - 仮称: `semantic.TypeRef` 型。primitive / named model / inline list / inline dict の 4 variant
  - `Param` / `Return` / `ModelField` / `Model.Element` / `Model.Value` から TypeRef を引けるよぁE��する
  - 既存�E `ModelName` / `Model` (QualifiedID) は migration 期間中は保持してよい
  - 由来: ADR-060 §影響「既存実裁E��の影響、E
- [ ] **rawyaml / semantic で `list<T>` / `dict<T>` 構文をパース**
  - 既存�E `ModelName` (raw string) パ�Eス箁E��を拡張
  - パ�Eスエラー時�E diagnostic は `invalid_type_ref` とする
  - 由来: ADR-060 §1, §6

- [ ] **既存�E named list/dict model 解決ロジチE��めETypeRef に正規化**
  - `kind: list, element: T` ↁE冁E��表現として inline `list<T>` と同じ container shape を引けるよぁE��する
  - `kind: dict, value: T` ↁE同槁E  - 名前 (id) と note は意味付けとして保持
  - 由来: ADR-060 §2

- [ ] **`internal/resolve/validation.go` に `validateFlowWiringTypes` を追加**
  - validate phase で実衁E  - ADR-060 §3 の互換ルール、E�5 の source 型解決、E�7 の抑制ルールを実裁E  - 由来: ADR-060 §影響「既存実裁E��の影響、E
- [ ] **TypeRef 互換判定�Eルパ�Eを追加**
  - 仮称: `typeRefsCompatible(a, b semantic.TypeRef) bool`
  - 由来: ADR-060 §3

- [ ] **diagnostic コード一覧に新 4 件を追加**
  - `internal/resolve/diagnostics.go` 等に `incompatible_wiring_type` / `invalid_wiring_source` / `invalid_foreach_over_type` / `invalid_type_ref` を追加
  - 由来: ADR-060 §6

#### チE��チE
- [ ] **互換ケースの単体テスチE*
  - 同一 primitive / 同一 named model / any wildcard 両方吁E/ list 同士の element 互換 / dict 同士の value 互換 / named list ⇁Einline list の正規化互換
  - 由来: ADR-060 §3

- [ ] **非互換ケースの単体テスチE*
  - 異なめEprimitive / 異なめEnamed model / primitive と named model の混在 / list 同士で element 不一致 / dict 同士で value 不一致
  - 由来: ADR-060 §3

- [ ] **`$item` 解決のチE��チE*
  - over ぁEnode ID で list を返す場吁E/ over ぁE`$params.field` で list を指す場吁E/ over ぁEany の場吁E/ over ぁElist でなぁE��吁E(invalid_foreach_over_type) / over の解決自体が失敗する場吁E  - 由来: ADR-060 §5-3

- [ ] **`$item` 有効篁E��のチE��チE*
  - foreach.params 外で `$item` を指定した場合に `invalid_wiring_source` が�EめE  - 由来: ADR-060 §5-3

- [ ] **重褁E��断抑制のチE��チE*
  - source ぁEunresolved の場吁Eincompatible_wiring_type が�EなぁE  - target param model ぁEunresolved の場吁Eincompatible_wiring_type が�EなぁE  - foreach.over ぁEinvalid_foreach_over_type の場合、�E側の wiring に incompatible_wiring_type が�EなぁE  - 由来: ADR-060 §7

- [ ] **UC-001 回帰チE��チE*
  - 既存�E wiring がすべて新ルールで OK と判定されることを確誁E  - 由来: ADR-060 §影響「UC-001、E
---

### Phase B: foreach.returns collected asset 参�Eルール

ADR-061 は accepted 済み、Eforeach の collected asset を後綁Eflow から参�Eする場合�EルールめEspec に反映し、実裁E��落とし込む、E
#### 確定した仕様篁E��

- `foreach.returns` は apply 允Etask の `returns` めEiteration ごとに collect した collected asset source 名である
- `foreach.returns` は optional。side-effect only の foreach では省略できる
- 省略時�E collected asset source めEsemantic model に生�EしなぁE��renderer / inspect / MCP めEinternal pseudo source を露出しなぁE- apply 允Etask に `returns` がなぁE��もかかわらず `foreach.returns` が指定された場合�E `invalid_foreach_returns`
- `foreach.returns` で宣言されぁEcollected asset は、同一 flow file 冁E�E後綁Estep / branch / fork / foreach から bare wiring source として参�Eできる
- 当該 foreach 自身の `params` 冁E��ら�E刁E�E `returns` 名を参�Eした場合�E `invalid_foreach_returns`
- collected asset source の TypeRef は apply 允Etask の `returns.model` `T` から `list<T>` として導�Eする
- apply 允Etask の `returns.model` ぁE`any` の場合�E `list<any>`
- apply 允Etask の `returns.model` が解決不�Eな場合、collected asset source の TypeRef も解決不�Eとして扱ぁE��後綁Ewiring の `incompatible_wiring_type` は抑制する
- `foreach.returns` は同一 flow file 冁E�E bare wiring source 名前空間に参加し、node id / 他�E `foreach.returns` と重褁E��てはならなぁE��重褁E��は `duplicate_flow_source`
- task の `returns.name` と `foreach.returns` が同名でも衝突扱ぁE��なぁE- ADR-023 の制御フロースコープ�E維持し、`foreach.returns` は collect 結果だけを外部 source として公開すめEescape hatch とする
- `foreach.id` は導�EしなぁE- `task.returns.source` および main task `returns.name` と `foreach.returns` の名前一致による暗黙接続�E本 Phase では扱わなぁE��EDR-062 領域として未実裁E�Eままにする

#### 仕様反映

- [x] **ADR-061 めEaccepted として確宁E*
- [x] **`docs/spec/edges.md` に `foreach.returns` collected asset source 仕様を反映**
  - `foreach.returns` の意味、optional、省略時�E扱ぁE  - 後綁Eflow からの bare source 参�E
  - TypeRef = `list<T>` 導�E
  - source 名前空間�E重褁E��ール
  - 制御フロースコープ�E escape hatch
  - `foreach.id` 非導�E、task return source 非対象
- [x] **`docs/spec/diagnostics.md` に ADR-061 diagnostics を反映**
  - `duplicate_flow_source`
  - `invalid_foreach_returns`
  - `unresolved_wiring_source`
  - `invalid_wiring_source` との区別

#### 実裁E
- [ ] **collected asset source めEsemantic / flow resolver に登録**
- [ ] **`duplicate_flow_source` を検�E**
  - `foreach.returns` と同一 flow file 冁Enode id の重褁E  - `foreach.returns` 同士の重褁E  - task `returns.name` との同名は衝突扱ぁE��なぁE- [ ] **`invalid_foreach_returns` を検�E**
  - apply 允Etask に `returns` がなぁE�Eに `foreach.returns` が指定されてぁE��
  - 当該 foreach 自身の `params` 冁E��ら�E刁E�E `returns` 名を参�EしてぁE��
- [ ] **`unresolved_wiring_source` を追加**
  - node id / `$params.<name>` / `$item` / collected asset source のぁE��れとしても解決できなぁEwiring source を診断する
  - `invalid_wiring_source` は「参照先�E存在するが、その斁E��では source として使えなぁE��場合に限定すめE- [ ] **collected asset source TypeRef = `list<T>` を導�E**
  - apply 允Etask `returns.model` ぁE`T` の場合�E `list<T>`
  - `any` の場合�E `list<any>`
  - `returns.model` が解決不�Eな場合�E collected asset source TypeRef も解決不�Eとして扱ぁE- [ ] **後綁Ewiring source として collected asset を解決**
  - 同一 flow file 冁E�E後綁Estep / branch / fork / foreach から bare source として参�E可能にする
  - 自刁E�E身の `foreach.params` から自刁E�E `returns` 名を参�Eすることは許可しなぁE- [ ] **`foreach.returns` 省略時�E source を生成しなぁE*
  - semantic model に collected asset source を作らなぁE  - renderer / inspect / MCP に internal pseudo source を露出しなぁE- [ ] **ADR-062 領域の task return source は未実裁E�Eままにする**
  - `task.returns.source` は追加しなぁE  - main task `returns.name` と `foreach.returns` の名前一致による暗黙接続�E実裁E��なぁE
#### チE��チE
- [ ] **collected asset source 解決チE��チE*
  - 後綁Estep / branch / fork / foreach から `foreach.returns` を参照できる
  - `foreach.returns` の TypeRef ぁE`list<T>` として target param と互換判定される
- [ ] **省略時�EチE��チE*
  - `foreach.returns` 省略時に source が生成されなぁE  - 省略時�E internal pseudo source ぁErenderer / inspect / MCP に露出しなぁE- [ ] **diagnostic チE��チE*
  - `duplicate_flow_source`
  - `invalid_foreach_returns`
  - `unresolved_wiring_source`
  - `invalid_wiring_source` との区別
  - TypeRef 解決不�E時に `incompatible_wiring_type` が抑制されめE
---

### Phase C: enum / discriminated object / inline struct / 深さ制陁E
UC-002 self-hosting で忁E��になめEdata layer 表現力を ADR ベ�Eスで導�Eする、E
#### 検討篁E��

- **enum**: 限定された値の雁E��を持つ primitive 型！Etatus, role 等！E- **discriminated object** (tagged union): 褁E��のmodel variant を識別フィールドで刁E��替える垁E- **inline struct**: model として外部に定義しなぁE��param/return冁E�Ead-hoc struct
- **`list<T>` / `dict<T>` の深さ制限また�E lint 方釁E* (ADR-060 §1 後段)

#### Tasks

- [ ] **enum 導�Eの要否判断 ADR を起票**
  - UC-002 で `status` / `role` 等�E enum 候補を観測
  - ADR としては proposed ↁE議諁EↁEaccepted の流れ

- [ ] **discriminated object 導�Eの要否判断 ADR を起票**
  - 同丁E
- [ ] **inline struct 導�Eの要否判断 ADR を起票**
  - 同丁E
- [ ] **TypeRef 深さ制陁E/ lint 方釁EADR を起票**
  - `list<dict<list<...>>>` のような深ぁE�Eれ子をどぁE��ぁE��
  - lint レベルの diagnostic か、構文エラーぁE
- [ ] **吁EADR acceptance 後、spec / 実裁E��反映**
  - Phase A の TypeRef 表現に追加 variant を足す形になる想宁E
---

## 実裁E��釁E
### Phase 進行頁E
- Phase A から着手！EDR-060 確定済み�E�E- Phase A 完亁E��、UC-002 self-hosting を�E開しながら Phase B / C の忁E��性を観測
- Phase B / C は ADR 起票 ↁE議諁EↁEaccepted を�E度行う

### commit 墁E��

- ADR-060 spec 反映 ↁE1 commit
- TypeRef semantic 表現追加 ↁE1 commit
- パ�Eサ拡張 ↁE1 commit
- validateFlowWiringTypes 実裁E+ チE��チEↁE1 commit
- 吁EADR-061 / Phase C 吁EADR は別 commit

### review 単佁E
- Phase A 完亁E��点でレビュー実施�E�EDR-060 と実裁E�E整合確認！E- Phase B / C は吁EADR ごとにレビュー

---

## 関連 ADR

- ADR-060: v1.1 TypeRef と flow wiring の型互換性 (accepted)
- ADR-061: foreach.returns collected asset 参�Eルール (予宁E/ Phase B で起票)
- ADR-021: primitive 予紁E��と list/dict model�E�E1.1 でも維持E��E- ADR-016: foreach 構文�E�変更なし。`$item` 型解決ルールは ADR-060 で明文化！E- ADR-058 / ADR-059: M14a で確定しぁEv1.0 系の実裁E��グ修正
- ADR-050: spec-first documentation policy
- ADR-057: self-hosting 思想 / Non-goals

---

## Evidence

- commit: 01e7127
- impl commit: 01e7127
