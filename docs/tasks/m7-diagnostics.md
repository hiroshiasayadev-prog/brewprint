# Milestone 7: validation / diagnostics を強化する

- **status**: closed
- **scope**: validation / diagnostics
- **source**: migrated from docs/TASKS.md
- **last_updated**: 2026-04-30

---

## Tasks

- [x] **semantic validation / diagnostics 第1段を実装する**
  - `semantic.Diagnostic` に `code` を追加
  - resolver diagnostic helperをcode対応
  - model / store / task / control / event の主要参照validationを追加
  - unresolved reads/writesはvalidation層のcode付きerrorに一本化
  - invalid rawyaml.Project fixtureのunit testを追加
  - UC-001 fixtureのクロスモジュール短縮参照をADR-027準拠に修正
  - API Table goldenを更新
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-1差分後）

- [x] **diagnostics CLI 第1段を実装する**
  - `brewprint validate --yaml-root <path>` を追加
  - diagnosticsなしなら `ok` をstdoutへ出力
  - diagnosticsを `severity code file: message` 形式で表示
  - error diagnosticsが1件以上あればnon-zero error
  - MCP起動時のsemantic diagnostic errorにもcodeを含める
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-2差分後）

- [x] **validate failure path testを追加する**
  - `t.TempDir()` でinvalid YAML rootを作成
  - 未解決 `params.model` を含む最小node YAMLを配置
  - stdoutのcode付きdiagnosticを固定
  - validation failed errorのerror/warning件数を固定
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-3差分後）

- [x] **validate JSON formatを追加する**
  - `brewprint validate --yaml-root <path> --format json` を追加
  - 既定のtext出力を維持
  - valid projectのJSON出力をCLI testで固定
  - invalid projectのJSON出力をCLI testで固定
  - unsupported formatをerrorにする
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-4差分後）

- [x] **diagnostics出力順を安定化する**
  - `resolve.Build` 返却前にdiagnosticsを安定ソート
  - severity / file id / code / message の順で固定
  - unit testで順序を固定
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-5差分後）

- [x] **validate warning-only path testを追加する**
  - warning-only YAML root helperを追加
  - text validateでwarning diagnosticを固定
  - json validateで `error_count=0` / `warning_count=1` を固定
  - warning-only projectではnon-zero errorにしないことを固定
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-6差分後）

- [x] **required fields validation 第1段を実装する**
  - `missing_required_field` diagnostic codeを追加
  - node id欠落を検出
  - model kind / model field name/type欠落を検出
  - store kind欠落を検出
  - param / return / initialized storeのname/model欠落を検出
  - unit testで `missing_required_field` を固定
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-7差分後）

- [x] **missing required fieldのCLI JSON出力を固定する**
  - `missing_required_field` を含むinvalid YAML root helperを追加
  - `brewprint validate --format json` のdiagnostics内codeを確認
  - validation failedのnon-zero errorを確認
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-8差分後）

- [x] **diagnostic code specific化 第1段を実装する**
  - duplicate系diagnosticを具体code化
  - flow系diagnosticを具体code化
  - warning-only validate testを `unsupported_flow_entry` 期待へ更新
  - flow diagnostic codeをunit testで固定
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-9差分後）

- [x] **transition diagnostic code specific化を実装する**
  - `unresolved_transition_state` / `unresolved_transition_event` を追加
  - `duplicate_transition` / `missing_transition_guard` を追加
  - transition系diagnosticを具体code化
  - transition diagnostic codeをunit testで固定
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-10差分後）

- [x] **view / scenario diagnostic code specific化を実装する**
  - API View / ER View / Sequence Scenario系diagnosticを具体code化
  - `duplicate_view` / `invalid_view_definition` / `duplicate_view_module` を追加
  - `unresolved_sequence_step` / `non_continuous_sequence` を追加
  - view / scenario diagnostic codeをunit testで固定
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-11差分後）

- [x] **diagnostics specを追加する**
  - `docs/spec/diagnostics.md` を追加
  - Diagnostic object / severity / text format / JSON formatを記述
  - diagnostic orderingを記述
  - 現在のdiagnostic code一覧をカテゴリ別に記述
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-12差分後）

- [x] **diagnostic fallback codeを定数化する**
  - `semantic_validation` fallback codeを定数化
  - `symbolTable.addDiagnostic` の直書き文字列を削除
  - 挙動変更なし
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-13差分後）

- [x] **M7総括docを追加する**
  - `docs/impl/go-m7-summary.md` を追加
  - M7で追加したCLI / diagnostics / validation / spec / 境界 / 次候補を整理
  - 次sessionの復帰用メモを記載
