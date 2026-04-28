# Go M8 render CLI / 一括render pipeline 総括

- **status**: complete
- **last_updated**: 2026-04-29
- **repo**: `C:\Users\imved\projects\brewprint`
- **verified**: `go test ./internal/render/dag` / `go test ./...` 通過済み（2026-04-29、M8-8差分後）

---

## 1. 次セッションで最初に読むもの

次セッションでは以下の順で読むと復帰しやすい。

1. `docs/prompt_chappy.md`
2. `docs/doc-policy.md`
3. `docs/impl/go-m7-summary.md`
4. `docs/impl/go-m8-summary.md`（このファイル）
5. `docs/TASKS.md` の末尾

必要に応じて追加で読むファイル:

- `cmd/brewprint/main.go`
- `cmd/brewprint/main_test.go`
- `internal/render/project/renderer.go`
- `internal/render/project/renderer_test.go`
- `internal/render/dag/renderer_test.go`
- `internal/render/placement/resolver.go`
- `docs/uc/001-ec-checkout-flow/renders/index.md`
- `docs/uc/001-ec-checkout-flow/renders/commerce/index.md`
- `docs/uc/001-ec-checkout-flow/renders/catalog/index.md`

---

## 2. M8で完了したこと

M8では `brewprint render` CLI と一括render pipelineを開始した。

### M8-1: render CLI 第1段

- `brewprint render --yaml-root <path> --out <path>` を追加
- semantic diagnosticsを通過したprojectだけrenderする
- `internal/render/project` packageを追加
- 既存rendererを束ねて以下を一括出力する
  - DAG
  - State
  - Sequence
  - ER
  - API
  - Wireframe
  - Preview
- `render_index.yaml` に基づき以下へ配置する
  - `renders/{group}/...`
  - `renders/_cross/...`
  - `renders/_preview/...`
- master `index.md` / group `index.md` を生成する
- CLI testではUC-001を `t.TempDir()` に出力し、生成件数と主要出力パスを確認する
- renderer本文golden一致は各 `internal/render/*` package testへ委譲し、CLI testは一括生成・配置確認に限定した

### M8-2: 一括render manifest test

- `internal/render/project/renderer_test.go` を追加
- UC-001の一括render出力23件をmanifestとして固定
- 全render fileのcontentが空でないことを確認
- `Write` がネストした出力ディレクトリを作成して書き込めることを確認

### M8-3: UC-001 render fixture追随 第1段

M8一括render manifestで不足していたDAG fixture 3本を追加した。

- `docs/uc/001-ec-checkout-flow/renders/commerce/dag-add_to_cart.md`
- `docs/uc/001-ec-checkout-flow/renders/catalog/dag-get_items.md`
- `docs/uc/001-ec-checkout-flow/renders/commerce/dag-process_payment.md`

また、master indexのDAG件数を現行一括render出力に合わせて更新した。

- `docs/uc/001-ec-checkout-flow/renders/index.md`

ただし、group indexの完全追随は後続差分扱いにしている。

### M8-4: DAG renderer golden test拡張

`internal/render/dag/renderer_test.go` のgolden対象に追加DAG 3本を加えた。

- `cart/task/add_to_cart.yaml`
- `catalog/task/get_items.yaml`
- `payment/webhooks/task/process_payment.yaml`

`dag-process_payment.md` は実レンダー出力に合わせて修正済み。

### M8-5: UC-001 group index fixture追随

M8-3で後続差分扱いにしていた group index を project renderer 形式へ追随した。

- `docs/uc/001-ec-checkout-flow/renders/auth/index.md`
- `docs/uc/001-ec-checkout-flow/renders/catalog/index.md`
- `docs/uc/001-ec-checkout-flow/renders/commerce/index.md`

反映内容:

- `commerce/index.md` に `dag-add_to_cart.md` / `dag-process_payment.md` を追加
- `catalog/index.md` に `dag-get_items.md` を追加
- 3つの group index を `kind | title | path` の project renderer 形式へ統一
- `internal/render/project/renderer_test.go` に group index fixture 一致テストを追加

注意:

- master `renders/index.md` は今回の fixture 一致テスト対象から外した。
- master index はタイトル生成や列アラインメントなど、project renderer index生成仕様整理の論点が残っているため。
- `go test ./...` はユーザー実行で通過済み（2026-04-29、M8-5差分後）。

### M8-6: render CLI `--clean` 追加

既存 `renders/` を再生成しやすくするため、`brewprint render` に `--clean` を追加した。

```powershell
brewprint render --yaml-root docs/uc/001-ec-checkout-flow/yaml --out docs/uc/001-ec-checkout-flow/renders --clean
```

挙動:

- validation / placement / render file build が成功した後、書き込み直前に `--out` を削除して作り直す
- `--clean` 未指定時は従来どおり既存ファイル削除なしで上書きする
- `--out` が空・`.`・filesystem root・`..` を含む場合は拒否する
- `--out` が `--yaml-root` を含む場合は YAML 削除事故防止のため拒否する

追加テスト:

- `cmd/brewprint`: stale file が `--clean` で削除されること
- `cmd/brewprint`: `--out` が `--yaml-root` を含む場合に拒否されること
- `internal/render/project`: `CleanOutRoot` が既存ディレクトリを削除して作り直すこと
- `internal/render/project`: unsafe path を拒否すること

注意:

- `go test ./...` はユーザー実行で通過済み（2026-04-29、M8-6差分後）。

### M8-7: UC-001 render fixture再生成コマンド明記

`docs/uc/001-ec-checkout-flow/README.md` に canonical fixture 再生成コマンドを明記した。

```powershell
brewprint render --yaml-root docs/uc/001-ec-checkout-flow/yaml --out docs/uc/001-ec-checkout-flow/renders --clean
go test ./...
```

また、README内の `renders/` ツリーにM8で追加したDAG fixtureを反映した。

- `commerce/dag-add_to_cart.md`
- `commerce/dag-process_payment.md`
- `catalog/dag-get_items.md`

### M8-8: project renderer index生成仕様固定

実装先行だった project renderer の index 生成仕様を fixture と test で固定した。

反映内容:

- master index title は project directory名を人間向けに整形して生成する
  - 例: `001-ec-checkout-flow` → `EC Checkout Flow`
- group index title は `render_index.yaml` の `label` を使う
  - `auth` → `認証`
  - `commerce` → `商取引`
  - `catalog` → `カタログ`
- master `renders/index.md` も project renderer fixture一致テスト対象に戻した
- preview title も project名整形へ追随した
  - `EC Checkout Flow Wireframe Preview`
- `humanizeProjectDir` の unit test を追加した

更新ファイル:

- `internal/render/project/renderer.go`
- `internal/render/project/renderer_test.go`
- `internal/render/wireframe/renderer_test.go`
- `docs/uc/001-ec-checkout-flow/renders/index.md`
- `docs/uc/001-ec-checkout-flow/renders/auth/index.md`
- `docs/uc/001-ec-checkout-flow/renders/commerce/index.md`
- `docs/uc/001-ec-checkout-flow/renders/catalog/index.md`
- `docs/uc/001-ec-checkout-flow/renders/_preview/wireframe.html`

注意:

- `go test ./...` はユーザー実行で通過済み（2026-04-29、M8-8差分後）。

---

## 3. 重要ファイル

### CLI

- `cmd/brewprint/main.go`
- `cmd/brewprint/main_test.go`

### 一括render

- `internal/render/project/renderer.go`
- `internal/render/project/renderer_test.go`

### 既存renderer / placement

- `internal/render/dag/renderer.go`
- `internal/render/dag/renderer_test.go`
- `internal/render/state/renderer.go`
- `internal/render/sequence/renderer.go`
- `internal/render/er/renderer.go`
- `internal/render/api/renderer.go`
- `internal/render/wireframe/renderer.go`
- `internal/render/placement/resolver.go`
- `internal/render/placement/index.go`

### UC-001 fixture

- `docs/uc/001-ec-checkout-flow/renders/index.md`
- `docs/uc/001-ec-checkout-flow/renders/auth/index.md`
- `docs/uc/001-ec-checkout-flow/renders/commerce/index.md`
- `docs/uc/001-ec-checkout-flow/renders/catalog/index.md`
- `docs/uc/001-ec-checkout-flow/renders/commerce/dag-add_to_cart.md`
- `docs/uc/001-ec-checkout-flow/renders/catalog/dag-get_items.md`
- `docs/uc/001-ec-checkout-flow/renders/commerce/dag-process_payment.md`

### タスク管理

- `docs/TASKS.md`

---

## 4. CLI仕様

```sh
brewprint render --yaml-root <path> --out <path> [--clean]
```

例:

```powershell
brewprint render --yaml-root docs/uc/001-ec-checkout-flow/yaml --out docs/uc/001-ec-checkout-flow/renders
brewprint render --yaml-root docs/uc/001-ec-checkout-flow/yaml --out docs/uc/001-ec-checkout-flow/renders --clean
```

正常時:

```text
rendered 23 file(s)
```

挙動:

- YAML load errorはerror
- semantic error diagnosticがある場合はrenderせずerror
- semantic warning diagnosticはstderrへ出し、renderは継続
- render_index placement errorがある場合はrenderせずerror
- render_index placement warningはstderrへ出し、renderは継続
- `--clean` 指定時は、render file build成功後かつwrite前に `--out` を削除して作り直す
- `--clean` 指定時でも、validation / placement / render file buildが失敗した場合は既存 `--out` を削除しない

---

## 5. 現在の一括render manifest

UC-001の期待出力23件:

```text
_cross/api.md
_cross/er.md
_preview/wireframe.html
auth/dag-login.md
auth/index.md
auth/state-auth.md
auth/wireframe-auth-loading.html
auth/wireframe-auth-login_screen.html
catalog/dag-get_items.md
catalog/index.md
catalog/state-inventory.md
commerce/dag-add_to_cart.md
commerce/dag-checkout.md
commerce/dag-process_order.md
commerce/dag-process_payment.md
commerce/dag-validate_cart.md
commerce/index.md
commerce/seq-checkout_flow.md
commerce/seq-payment_webhook_flow.md
commerce/state-order.md
commerce/wireframe-order-cart.html
commerce/wireframe-order-checkout_screen.html
index.md
```

---

## 6. 境界

維持する境界:

```text
source  -> rawyaml
resolve -> rawyaml, semantic
query   -> semantic
mcp     -> query
renderer -> semantic
```

禁止:

```text
internal/mcp -> rawyaml
internal/mcp内でYAML load / resolve
internal/mcp内でrenderer呼び出し
renderer内でRaw YAML structsを直接読むこと
renderer内でname resolution / semantic validationを再実装すること
```

M8の `internal/render/project` は `rawyaml.Project` と `semantic.Project` を受け取る。
これは `render_index.yaml` のgroup設定を読むために現状必要。
ただし、個別renderer（DAG / State / Sequence / ER / API / Wireframe）は引き続き `semantic.Project` を入力にする。

---

## 7. 残っていること / 次候補

### 完了: UC-001 group index完全追随

`docs/uc/001-ec-checkout-flow/renders/index.md` はM8-3で件数を更新済み。
M8-5で group index は project renderer 形式へ追随済み。

反映済み:

- `docs/uc/001-ec-checkout-flow/renders/commerce/index.md`
  - `dag-add_to_cart.md` を追加
  - `dag-process_payment.md` を追加
  - project renderer形式へ統一
- `docs/uc/001-ec-checkout-flow/renders/catalog/index.md`
  - `dag-get_items.md` を追加
  - project renderer形式へ統一
- `docs/uc/001-ec-checkout-flow/renders/auth/index.md`
  - project renderer形式へ統一
- group index fixture一致テストを `internal/render/project/renderer_test.go` に追加済み。
- M8-5差分後の `go test ./...` はユーザー実行で通過済み。

### 完了: render CLIで既存rendersを再生成する運用の整理

M8-6で `--clean` 実装と安全チェックは追加済み。

追加済み:

- `--clean` を追加して出力前に対象ディレクトリを削除する
- 削除対象を安全な `--out` に限定するチェックを入れる

完了:

- `go test ./...` で M8-6 差分を確認済み
- 既存render fixture再生成コマンドをREADMEへ明記済み

### 完了: project renderer index生成の仕様整理

M8-8で project renderer の index 生成仕様は現行 fixture と test で固定済み。

固定した仕様:

- master index titleは project directory名を人間向けに整形する
- group index titleは `render_index.yaml` の `label` を使う
- group index列は `kind | title | path`
- master index / group index は project renderer fixture一致テストで固定する

---

## 8. 直近検証

最後にユーザーが実行して通過したコマンド:

```powershell
go test ./...
```

結果:

- 全package: OK
- `internal/render/project`: OK
- `internal/render/wireframe`: OK

M8-8差分後に通過済み。

---

## 9. commit状態メモ

この環境からgit状態は確認していないため、次セッション冒頭で必ず確認すること。

```powershell
git status
```

会話上の情報:

- M8-1 / M8-2 はcommit済みの流れ
- M8-3について、ユーザーは「commitは打った」と発言
- M8-4〜M8-8については、まとめてcommit対象にするのが自然

次セッションでは、まず `git status` で未commit差分を確認する。

---

## 10. 次セッション開始時の推奨手順

```powershell
cd C:\Users\imved\projects\brewprint
git status
go test ./...
```

その後、状況に応じて以下へ進む。

1. M8-4〜M8-8差分をcommitする
2. 次milestone候補を選ぶ
   - state / event / scenario 系reverse lookup index
   - validation強化の続き
   - render CLI / docs の仕上げ

