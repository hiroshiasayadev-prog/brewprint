---
scope: docs/impl/review/findings-source.md
status: draft
last_updated: 2026-04-29
summary: >
  source層レビュー（internal/source/ 3ファイル）で出たfindingsの最終整理。
  各findingの最終対応（spec移管 / ADR起票 / コード修正 / 保留）を記録する。
target_layer: source
target_files:
  - internal/source/loader.go
  - internal/source/classify.go
  - internal/source/file.go
---

# source層レビュー findings 最終整理

## 概要

source層レビュー（Sonnet担当）と Claude による追加発見で出た findings 8件を、V01-ADR-050 の spec-first 方針に基づいて整理。
新設 spec 3本（file-types.md / project-layout.md / naming.md）と新規 ADR 2本（051 / 052）、既存 ADR 漸進移行 4本（002 / 030 / 043 / 045）で対応した。

## findings サマリ

| ID | ラベル | 概要 | 最終対応 |
|---|---|---|---|
| F-S-001 | spec-migrate | V01-ADR-030 の `as:` 値一覧に `er_diagram` 未記載 | spec/file-types.md §3 に集約。V01-ADR-030 漸進移行 |
| F-S-002 | spec-decision + ADR | Unsupported の扱い未文書化 | warning diagnostic 採用。V01-ADR-051 起票、spec/file-types.md §4 + spec/diagnostics.md 反映 |
| F-S-003 | spec-migrate | V01-ADR-002 の `master.yaml` が V01-ADR-043 で覆されているが V01-ADR-002 に注記なし | V01-ADR-002 の partial supersede 行を更新、漸進移行 |
| F-S-004 | spec-migrate | render_index.yaml の loader 探索ルールが未文書化 | spec/project-layout.md §1〜§2 に記載 |
| F-S-005 | doc-only | V01-ADR-002 / V01-ADR-030 の Evidence の `impl commit: tbd` | **未対応**（後述）|
| F-S-006 | spec-decision + ADR | `File.Path` が abs/rel どちらか未規定 | プロジェクトルート相対 + slash 正規化を採用。V01-ADR-052 起票、spec/file-types.md §5 反映 |
| F-S-007 | doc-only | wireframe は file-level view を持たないが V01-ADR-030 に明示なし | spec/file-types.md §3 注記で対応 |
| F-S-008 | doc-only | V01-ADR-030 は `as:` 有無で振り分けと書くが、実装は `nodes:` キーも判定 | spec/file-types.md §2 で分類アルゴリズムを正確に記述 |

## 各findingの詳細

### F-S-001 [spec-migrate]: `as:` 値一覧

**問題**: V01-ADR-030 の `as:` 値表に `er_diagram`（V01-ADR-039 で追加）が反映されていなかった。

**対応**:
- `docs/spec/file-types.md` §3 に `as:` 値一覧を集約。`api_table` / `sequence_diagram` / `er_diagram` を全て記載
- 各値の由来 ADR を表に明記
- V01-ADR-030 漸進移行: 起票時点表を残しつつ、現行仕様は spec 参照と注記

**ステータス**: 完了。

### F-S-002 [spec-decision + ADR]: Unsupported ファイルの扱い

**問題**: `as:` も `nodes:` も持たないファイルが silent skip されていた。typo / 書き忘れの検出ができない。

**対応**:
- V01-ADR-051 起票（accepted）: `unsupported_file` warning diagnostic を loader フェーズで発行
- `docs/spec/file-types.md` §4 に方針を記載
- `docs/spec/diagnostics.md` の diagnostic code 一覧に `unsupported_file` (warning) を追加

**実装側への影響**:
- `internal/source` 層で warning diagnostic 配信パスが必要。M1 時点では diagnostic 配信機構が未整備のため、impl 側で実装は別タスク扱い。

**ステータス**: 仕様確定。実装は impl 側で別タスク。

### F-S-003 [spec-migrate]: `master.yaml` 不在

**問題**: V01-ADR-002 が定めた `master.yaml` が V01-ADR-043 で実質廃止されたが、V01-ADR-002 に注記がなかった。

**対応**:
- V01-ADR-002 の `partial supersede` 行に「V01-ADR-043により `master.yaml` は廃止」を追加
- 「決定」セクション内の `master.yaml` 記述に「V01-ADR-043 で廃止」の注記
- `docs/spec/project-layout.md` §1 に「`master.yaml` のようなプロジェクトルート台帳ファイルは存在しない」と明記
- `docs/spec/naming.md` §1 にフォルダ階層 = モジュール階層を移管

**ステータス**: 完了。

### F-S-004 [spec-migrate]: render_index.yaml 探索ルール

**問題**: loader が「`yaml/` の親ディレクトリ直下の `render_index.yaml` を探す」挙動が V01-ADR-043 から自然に読み取れない。

**対応**:
- `docs/spec/project-layout.md` §1（プロジェクトルート構造）と §2（yaml/ の構造、loader 探索ルール）に明記

**ステータス**: 完了。

### F-S-005 [doc-only]: impl commit が tbd

**問題**: V01-ADR-002 / V01-ADR-030 の Evidence の `impl commit: tbd` が放置されている。本来は実装commit hash で埋めるべき。

**対応**: **未対応**。ユーザーに git log での該当 commit hash 提示を依頼する必要がある。

**ステータス**: 保留。次会話以降、または本会話の commit 提案時にユーザー側で対応。

### F-S-006 [spec-decision + ADR]: `File.Path` 正規化

**問題**: `rawyaml.File.Path` が OS ネイティブセパレータを含む絶対パスのままで、golden test の決定性が崩れる懸念。

**対応**:
- V01-ADR-052 起票（accepted）: プロジェクトルート相対 + slash 正規化を採用
- `docs/spec/file-types.md` §5 に方針を記載

**実装側への影響**:
- `internal/source/loader.go` の `loadFile` 呼び出し前後で `Path` を正規化する経路を追加する必要がある
- `internal/source/file.go` に正規化ヘルパーが必要になる可能性がある

**ステータス**: 仕様確定。実装は impl 側で別タスク。

### F-S-007 [doc-only]: wireframe の扱い

**問題**: V01-ADR-030 の `as:` 値表に wireframe が含まれないが、その理由が明示されていなかった。

**対応**:
- `docs/spec/file-types.md` §3 に注記: 「wireframe はファイルレベルの view を持たない。state ノード内に DSL として埋め込まれる形でのみ存在する（V01-ADR-029, V01-ADR-042）」

**ステータス**: 完了。

### F-S-008 [doc-only]: 分類判定アルゴリズム

**問題**: V01-ADR-030 は「`as:` 有無で振り分け」と書くが、実装は `as:` 優先 → `nodes:` チェック → Unsupported の3段判定をしている。

**対応**:
- `docs/spec/file-types.md` §2 に分類アルゴリズムを4段階（ファイル名 → `as:` → `nodes:` → unsupported）で正確に記述
- V01-ADR-030 漸進移行と合わせ、現行仕様は spec 参照とした

**ステータス**: 完了。

## 残タスク

### 本会話で着手しなかったもの

- **F-S-005**: V01-ADR-002 / V01-ADR-030 の `impl commit: tbd` の更新（ユーザー側で git log 提示が必要）
- **F-S-002 / F-S-006 の実装**: `unsupported_file` warning 発行と `File.Path` 正規化の実装。M1 で diagnostic 配信機構を実装する際に組み込む方針

### resolve 層レビューに引き継ぐもの

- **spec/naming.md の肉付け**: 現状最小骨格のみ。actor の global 定義 / FK 解決 / cross-edge 等は resolve 層レビューで追記
- **ネスト module 内 task の render 出力ファイル名**: V01-ADR-045 §8 で「Go 実装 ADR で規定」と明記済み。実装 ADR 起票時に確定

## 成果物

### 新規作成

- `docs/spec/file-types.md`
- `docs/spec/project-layout.md`
- `docs/spec/naming.md`（最小骨格、status: wip）
- `docs/adr/051-unsupported-yaml-file-warning.md`
- `docs/adr/052-source-file-path-normalization.md`

### 更新

- `docs/spec/diagnostics.md`（`unsupported_file` warning code 追加）
- `docs/adr/002-folder-as-namespace.md`（漸進移行、`master.yaml` 廃止注記）
- `docs/adr/030-yaml-file-type-declaration.md`（漸進移行）
- `docs/adr/043-project-root-layout-and-render-output.md`（漸進移行）
- `docs/adr/045-render-index-schema.md`（漸進移行）

### 削除

- `docs/impl/review/findings-source-handoff.md`（本作業で完了したため）
