# Milestone 14a: subnode scope + return primitive 実装修正

- **status**: open
- **scope**: internal/resolve / spec / tests
- **source**: M14 Phase A中に発覚した実装バグ。ADR-058 / ADR-059 で決定
- **last_updated**: 2026-05-02

---

## Context

ADR-057でv1.0.0-spec凍結後、M14 Phase A（MCP公開contract のblueprint化）に着手したところ、
既存実装と既存specの間に2系統の乖離が判明した。

**B1: サブノードのファイル内private性が実装で保証されていない**

- ADR-011 §3 / spec/nodes.md「ファイル構造」がサブノードはファイル内privateと規定
- しかし `internal/resolve/symbols.go` の `addNode` がサブノードまでproject全体の単一QIDマップで重複チェックする
- 結果、別ファイル同名サブノードで `duplicate_node` が発生
- 詳細は ADR-058 を参照

**B2: task / join の return が primitive を許容しない**

- ADR-021 §3 / spec/nodes.md primitive予約語節が `any` 等7語の primitive を定義
- `validateParams` は primitive を許容、しかし `validateReturn` は project内 model のみ許容
- 結果、`returns.model: any` 等で `unresolved_model` が発生
- 詳細は ADR-059 を参照

UC-001ではどちらのパターンも偶然存在しなかったため M0〜M13 で顕在化しなかった。
UC-002 Phase A YAMLで両方が同時露呈し、Phase A render が失敗した。

このmilestoneは ADR-058 / ADR-059 の方針に従い、実装をspecに追従させる純粋なbug fix milestoneとして扱う。

---

## Non-goals

M14aでは以下を行わない。

- ADR-011 / ADR-021 / ADR-027 / spec/nodes.md ファイル構造節の遡及修正
  - doc-policy.md §3 / ADR-058 §理由 / ADR-059 §理由
- v1.0.0-spec タグの再発行
  - 本milestoneはpatch release扱い (`v1.0.1-spec`)。v1.0.0-spec タグは保持
- UC-002 Phase A YAML の直接 render 検証
  - UC-002 Phase A YAML は M15（data layer expressiveness v1.1）完了後に enum / discriminated object / inline struct を使った形で再構築する方針
  - M14a 完了直後の Phase A YAML 状態は draft 保留扱い (TASKS-UC-002.md 参照)
  - M14a の検証は **UC-001 と新規回帰テストのみ** で行う
- v1 model の表現力拡張（enum / discriminated object / inline struct / union）
  - すべて M15 のスコープ。本milestoneは spec通りに実装が動くことを保証するのみ
- ADR-010の複数論点混在分割（v1後検討事項として継続。ADR-057 §6）

---

## Tasks

### 仕様反映

- [ ] **`docs/spec/nodes.md` ファイル構造節を更新する** (ADR-058 §4)
  - 「サブノードはファイル内private」を以下まで強化する
    - サブノードIDは同一ファイル内で一意
    - 別ファイルにある同名サブノードとは衝突しない
    - サブノードは QualifiedID で外部参照されない
  - 由来注記に ADR-058 を追加

- [ ] **`docs/spec/nodes.md` task節 / `returns` オブジェクト節を更新する** (ADR-059 §3)
  - `model` に primitive を指定可能であることを明記
  - 由来注記に ADR-059 を追加

- [ ] **`docs/spec/diagnostics.md` を更新する** (ADR-058 §4 / ADR-059 §3)
  - `duplicate_node`: メインノード間衝突に限定する旨を明記、由来に ADR-058 追加
  - `unresolved_model`: primitive は対象外である旨を明記、由来に ADR-059 追加

- [ ] **`docs/spec/naming.md` を更新する** (ADR-058 §4)
  - メインノード = 外部参照可能な QualifiedID 保持者
  - サブノード = ファイル内private、QualifiedIDによる外部参照対象外
  - §2 周辺に追記、由来に ADR-058 追加

### 実装修正 — B1 (subnode scope)

- [ ] **`internal/resolve/symbols.go` の `addNode` を修正する**
  - メインノード（`IsMain() == true`）のみ `NodesByQID` への重複チェック対象とする
  - サブノードは `NodesByFile[fileID]` への登録のみ行う
  - サブノード間の同一ファイル内ID重複は別 diagnostic（同ファイル内重複）として扱うか検討する

- [ ] **サブノードを内部識別する経路を整理する**
  - render / query layerでサブノードを引く経路を `(FileID, kind, id)` 等のファイルスコープキーに統一する
  - 既存のNodesByQID参照箇所のうちサブノードを混在させていた箇所を洗い出す
  - 既存の `isPrivateSubNode` / `PrivateNodeID` 補助関数を活用する

- [ ] **flow / reads / writes / param の解決ロジックを確認する**
  - `internal/resolve/flow.go` の `resolveNodeQID` は既にファイルスコープ優先実装になっているため、現行動作を維持
  - 解決順序がADR-058 §3 と整合していることを再確認する
  - 同一moduleの別ファイルに同名メインnodeがある場合の解決優先順を回帰テストで確認する

- [ ] **reference / impact ロジックを確認する**
  - `internal/resolve/references.go` 等がサブノードまで巻き込んで graph traversal していないか確認
  - サブノードを除外した上で参照関係が正しく構築されることを確認する

### 実装修正 — B2 (return primitive)

- [ ] **`internal/resolve/validation.go` の `validateReturn` を修正する**
  - `modelExists` から `modelOrPrimitiveExists` ベースの判定に切り替える
  - `validateParams` と対称な解決ロジックにする

- [ ] **task return / join return から primitive へのreference扱いを確認する**
  - `internal/resolve/references.go` の `buildReferences` 内の return 関連処理が
    primitive return を `modelOrPrimitiveKey` / `modelOrPrimitiveEndpoint` で扱えるか確認
  - 必要であれば対称化する

### テスト — B1 関連

- [ ] **回帰テスト追加: 別ファイル同名サブノードが衝突しない**
  - `internal/resolve/validation_test.go` に追加
  - 同一moduleの2ファイルが同名サブtaskを持つfixtureで `duplicate_node` が出ないことを確認

- [ ] **回帰テスト追加: ファイル内サブノード重複は検出される**
  - 同一ファイル内で同名サブノードを定義した場合、適切な diagnostic が出ること

- [ ] **回帰テスト追加: メインノード衝突は引き続き検出される**
  - 同一moduleの2ファイルが同名メインノードを持つ場合、`duplicate_node` が出ること

- [ ] **回帰テスト追加: flow解決のファイルスコープ優先**
  - 別ファイルにある同名task（main）と、自ファイル内のサブtaskが両方存在する場合、
    flow `step: <id>` は自ファイルのサブtaskを優先解決すること

### テスト — B2 関連

- [ ] **回帰テスト追加: task return が primitive を許容する**
  - `returns.model: any` / `str` / `int` 等で `unresolved_model` が出ないこと

- [ ] **回帰テスト追加: join return が primitive を許容する**

- [ ] **回帰テスト追加: task return が存在しない model を指す場合は依然として `unresolved_model` を出す**
  - primitive 許容と未定義 model 検出の両立を確認

### テスト — 全体

- [ ] **`go test ./...` パスを確認する**
- [ ] **`gofmt` 整形を確認する**
- [ ] **UC-001 が引き続き通ることを確認する**
  - `brewprint render --yaml-root docs/uc/001-ec-checkout-flow/yaml --out docs/uc/001-ec-checkout-flow/renders` が成功し、既存 fixture と一致すること

### M14a クローズ作業

- [ ] **`docs/impl/go-m14a-summary.md` を起票する**
  - 修正内容、追加した回帰テスト、ADR-058 / ADR-059との対応関係を記録
  - UC-002 Phase A 直接render検証が M14a スコープ外であること、M15 後に再構築する方針を引き継ぎとして記録

- [ ] **`docs/TASKS.md` の M14a status を closed にする**

- [ ] **本ファイル `docs/tasks/m14a-subnode-scope-fix.md` を closed にする**

- [ ] **`v1.0.1-spec` タグを発行する** (ADR-058 §影響 / ADR-059 §影響)
  - patch release: bug fix のみ
  - UC-001 で動作確認済み、UC-002 Phase A は draft 保留扱い
  - タグ実行はユーザーが行う:
    ```bash
    git tag -a v1.0.1-spec -m "brewprint v1.0.1-spec snapshot (bug fix)"
    git push origin v1.0.1-spec
    ```

---

## ADR-058 / ADR-059 との対応

| ADR | 決定 | M14a task |
|---|---|---|
| ADR-058 §1 | サブノードはQID一意性制約から除外 | impl修正 B1 — `addNode` 修正 |
| ADR-058 §2 | サブノードの内部識別 | impl修正 B1 — 内部識別経路の整理 |
| ADR-058 §3 | ファイル内bare ID解決のスコープ明示 | impl修正 B1 — flow解決確認 / 仕様反映 — naming.md |
| ADR-058 §4 | spec更新 (nodes / diagnostics / naming) | 仕様反映セクション |
| ADR-058 §5 | 実装方針（参考） | impl修正 B1 全体の指針 |
| ADR-059 §1 | task / join return が primitive を許容 | impl修正 B2 — `validateReturn` 修正 |
| ADR-059 §2 | validation の対称化 | impl修正 B2 — 同上 |
| ADR-059 §3 | spec更新 (nodes / diagnostics) | 仕様反映セクション |

---

## 後続milestoneへの引き継ぎ

M14a 完了後の進行は以下の予定:

1. **M15: data layer expressiveness v1.1**
   - ADR-060: model field enum type
   - ADR-061: discriminated object model kind
   - ADR-062: inline struct field type
   - free union (untagged) は導入しない方針 (ADR-061で却下案として明記予定)
   - 完了で `v1.1.0-spec` タグ発行

2. **M14: brewprint self-hosting 再開**
   - M14a + M15 完了後に Phase A YAML を enum / discriminated object / inline struct を使った形で初構築
   - 既存の Phase A YAML（`note` + `any` 暫定版）は M15 完了時点で全面書き直しとなる

3. **UC-002 Phase A 検証**
   - M14 再開後、Phase A render を初回成功させる
   - spec gap #1 / #2 の現行記述（暫定対応 = `any`）は M15 完了時点で大半が解消されるため、
    TASKS-UC-002.md のspec gap整理を更新する
