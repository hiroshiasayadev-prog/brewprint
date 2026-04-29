# Milestone 8: render CLI / 一括render pipeline を実装する

- **status**: open
- **scope**: render CLI / render pipeline
- **source**: migrated from docs/TASKS.md
- **last_updated**: 2026-04-30

---

## Tasks

- [x] **render CLI 第1段を実装する**
  - `brewprint render --yaml-root <path> --out <path>` を追加
  - validate相当のsemantic diagnosticsを通過したprojectだけrenderする
  - `internal/render/project` に一括render orchestratorを追加
  - 既存rendererを束ねて、DAG / State / Sequence / ER / API / Wireframe / Previewを出力する
  - `render_index.yaml` に基づき `renders/{group}/...` / `_cross/` / `_preview/` へ配置する
  - master `index.md` / group `index.md` を生成する
  - CLI testではUC-001を `t.TempDir()` に出力し、主要出力パスと生成件数を固定する
  - renderer本文のgolden一致は各 `internal/render/*` package testに委譲し、CLI testは一括生成・配置確認に限定する
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-29、M8-1差分後）

- [x] **一括render manifest testを追加する**
  - `internal/render/project/renderer_test.go` を追加
  - UC-001の一括render出力23件をmanifestとして固定
  - 全render fileのcontentが空でないことを確認する
  - `Write` がネストした出力ディレクトリを作成して書き込めることを固定
  - CLI testではなく一括render本体側で出力パス契約を担保する
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-29、M8-2差分後）

- [x] **UC-001 render fixtureを一括render出力へ追随する 第1段**
  - M8一括render manifestで不足していたDAG fixture 3本を追加
    - `renders/commerce/dag-add_to_cart.md`
    - `renders/catalog/dag-get_items.md`
    - `renders/commerce/dag-process_payment.md`
  - master `renders/index.md` のDAG件数を現行一括render出力に合わせて更新
  - group indexの完全追随は後続差分で扱う
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-29、M8-3差分後）

- [x] **DAG renderer golden testを追加fixtureへ拡張する**
  - `internal/render/dag/renderer_test.go` のgolden対象に追加DAG 3本を追加
    - `cart/task/add_to_cart.yaml`
    - `catalog/task/get_items.yaml`
    - `payment/webhooks/task/process_payment.yaml`
  - `dag-process_payment.md` を実レンダー出力に合わせて更新
  - `go test ./internal/render/dag` / `go test ./...` 通過済み（2026-04-29、M8-4差分後）

- [x] **UC-001 group index fixtureを一括render出力へ追随する**
  - `renders/commerce/index.md` に追加DAG 2本を反映
    - `dag-add_to_cart.md`
    - `dag-process_payment.md`
  - `renders/catalog/index.md` に `dag-get_items.md` を反映
  - `renders/auth/index.md` / `renders/catalog/index.md` / `renders/commerce/index.md` を project renderer 形式へ統一
  - `internal/render/project/renderer_test.go` に group index fixture 一致テストを追加
  - `go test ./...` 通過済み（2026-04-29、M8-5差分後）

- [x] **render CLI `--clean` を追加する**
  - `brewprint render --yaml-root <path> --out <path> --clean` を追加
  - render / validation / placement 成功後、書き込み直前に `--out` を削除して作り直す
  - `--out` が空・`.`・filesystem root・`..` を含む場合は拒否する
  - `--out` が `--yaml-root` を含む場合は YAML 削除事故防止のため拒否する
  - stale file 削除のCLI testを追加
  - unsafe clean path のunit testを追加
  - `go test ./...` 通過済み（2026-04-29、M8-6差分後）

- [x] **UC-001 render fixture再生成コマンドをREADMEへ明記する**
  - `docs/uc/001-ec-checkout-flow/README.md` に `brewprint render --yaml-root ... --out ... --clean` を追記
  - READMEの `renders/` ツリーに M8で追加したDAG fixtureを反映
  - docsのみの変更のため追加テストなし

- [x] **project renderer index生成仕様を固定する**
  - master index titleを project directory名から人間向けに整形する
  - group index titleに `render_index.yaml` の `label` を使う
  - master `index.md` も project renderer fixture一致テスト対象に戻す
  - preview titleも project名整形に追随
  - `go test ./...` 通過済み（2026-04-29、M8-8差分後）

- [ ] **render output path collision validationを実装する**
  - ADR-053 / `docs/spec/project-layout.md` に準拠
  - v1では nested module path を出力ファイル名に含めず、local ID filenameを維持する
  - 同一group内で複数render outputが同じrelative pathに解決された場合はerrorにする
  - silent overwriteは禁止する
  - collision対象source objectをerror messageに含める
  - project renderer / placement周辺のunit testで固定する
