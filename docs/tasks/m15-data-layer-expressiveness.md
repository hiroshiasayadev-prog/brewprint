# Milestone 15: data layer expressiveness (v1.1)

- **status**: open
- **scope**: internal/semantic / internal/resolve / spec / docs/adr / tests
- **source**: ADR-060 (TypeRef + flow wiring type compatibility) を起点とする v1.1 系の表現力拡張
- **last_updated**: 2026-05-03

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
M15 はその ADR-060 を実装に落とし込みつつ、関連する v1.1 表現力拡張をまとめて扱う milestone である。

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

M15 は3つの Phase に分けて進める。
Phase A は ADR-060 で確定済みのため即着手可能。
Phase B / C は ADR-061（予定）および追加 ADR の起票を伴う。

### Phase A: TypeRef + flow wiring type validation

ADR-060 を実装と spec に落とし込む。

#### 仕様反映

- [ ] **`docs/spec/nodes.md` または新規 `docs/spec/type-ref.md` に TypeRef 構文を追加**
  - 構文形式: primitive / named model / inline `list<T>` / inline `dict<T>`
  - 再帰的定義の許容範囲（`list<dict<user>>` 等の入れ子）
  - 深さ制限または lint 方針は Phase C で扱うため、本 phase では「制限は M15 Phase C で扱う」とのみ明記
  - TypeRef を受け取るフィールド: `param.model` / `returns.model` / `field.type` / `model.element` / `model.value`
  - 由来: ADR-060 §1, §9

- [ ] **`docs/spec/edges.md` §1（flow:セクション）末尾に「§型互換性ルール」節を追加**
  - ADR-060 §3 のルールを spec として記述
  - ADR-060 §4 の検証対象 wiring 一覧
  - ADR-060 §5 の wiring source 型解決ルール（node ID / `$params.<name>` / `$item`）
  - ADR-060 §7 の型解決失敗時の抑制ルール
  - 由来: ADR-060 §3〜§7, §9

- [ ] **`docs/spec/edges.md` §1-5 foreach.over の記述を ADR-060 と整合させる**
  - 現状: 「`$params.field` を指定した場合、parser/validator は `main.params.<field>.model` が `kind: list` であることを検証する」
  - ADR-060 後: 「`foreach.over` の解決結果が list を表す TypeRef でない場合 `invalid_foreach_over_type` を出す」
  - UC-001 `validate_cart.yaml` コメントに記録された spec gap（`foreach.over` に `$params.field` を指定可能か明文化されていない件）を併せて解消
  - 由来: ADR-060 §5-3

- [ ] **`docs/spec/diagnostics.md` に新 diagnostic 3件を追加**
  - `incompatible_wiring_type` (severity: error)
  - `invalid_wiring_source` (severity: error)
  - `invalid_foreach_over_type` (severity: error)
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
  - パースエラー時の diagnostic（仮称 `invalid_type_ref` 等）を要否含めて検討
  - 由来: ADR-060 §1

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

- [ ] **diagnostic コード一覧に新 3 件を追加**
  - `internal/resolve/diagnostics.go` 等に const 追加
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

ADR-061 (予定) を起票し、foreach の collected asset を後続 flow から参照する場合のルールを定める。

#### 検討範囲

- foreach.returns の意味整理（source id か label か）
- collected asset を後続 flow から参照する場合の source id 自動生成ルール
- `foreach.id` フィールド導入の要否
- source id 重複ルール
- ADR-016 との整合（supersedes か追補か）
- main task returns との暗黙接続パターン（UC-001 `validate_cart.yaml` の例: foreach.returns と main returns の name 一致）

#### Tasks

- [ ] **UC-001 / UC-002 で foreach.returns を後続 flow から参照する具体例を観測する**
  - UC-001 `validate_cart.yaml` 段階では後続参照の実例なし（M15 Phase A 進行中に Phase B 起票判断材料として再評価）
  - UC-002 self-hosting 開始後に collected asset を後続参照するケースが出現すれば、その時点で ADR-061 を起票

- [ ] **ADR-061 起票（タイミングは具体例観測後）**
  - status: proposed → 議論 → accepted

- [ ] **ADR-061 acceptance 後、spec / 実装に反映**
  - 詳細タスクは ADR-061 起票時に追記

#### Phase B 起票判断ポリシー

- UC-002 self-hosting で collected asset の後続参照が必要になった時点で起票
- それまでは設計を凍結し、Phase A / Phase C を優先する

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
- ADR-061: foreach.returns collected asset 参照ルール (予定 / Phase B で起票)
- ADR-021: primitive 予約語と list/dict model（v1.1 でも維持）
- ADR-016: foreach 構文（変更なし。`$item` 型解決ルールは ADR-060 で明文化）
- ADR-058 / ADR-059: M14a で確定した v1.0 系の実装バグ修正
- ADR-050: spec-first documentation policy
- ADR-057: self-hosting 思想 / Non-goals

---

## Evidence

- commit: tbd
- impl commit: tbd
