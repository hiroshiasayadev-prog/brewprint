# Milestone 13: MCP analyze_impact を実装する

- **status**: closed
- **scope**: MCP / QueryService
- **source**: split from M12 analyze_impact design
- **last_updated**: 2026-05-01

---

## Closeout review status

Docs closeout draft was created at `docs/impl/go-m13-summary.md`.

M13 is closed as a hybrid v1 close.

Closeout basis:

- Final M13 gap-fix changes were applied: shared mechanical judgement gate, v1-minimal flow param field impact collector, transition new target/action resolution check, and add coverage correction.
- `gofmt -w (Get-ChildItem -Recurse -Filter *.go -File | ForEach-Object FullName)` completed successfully.
- `go test ./...` passed after the final gap-fix changes.
- Remaining limitations are documented in `docs/spec/mcp/tools/analyze-impact.md` and `docs/impl/go-m13-summary.md`.

---

## Context

M12で `analyze_impact` の設計は完了済み。

- ADR-056で tool design を accepted
- `docs/spec/mcp/tools/analyze-impact.md` に外部I/O契約を定義済み
- `docs/spec/mcp/overview.md` / `docs/spec/mcp/versioning.md` は設計反映済み想定

M13では、ADR-056 / spec に従って QueryService / MCP wrapper / tests を実装する。

`analyze_impact` は `get_references` / `get_reference_tree` / `inspect` / source 補助情報 / render output file mapping を材料に、change kind ごとの意味づけ済み impact list を返す上位toolとして扱う。

---

## Non-goals

M13では以下を実装しない。

- model 間の structural compatibility / subtyping 判定
- nullable / optional / required の互換性判定
- semantic contract / behavioral compatibility 判定
- render output md 内の presentation detail 差分
- wireframe element binding 影響分析
- Raw YAML AST の外部公開
- `get_reference_tree` の新規 reference kind 追加
- suggested_fixes の自動適用 refactor tool

これらは `coverage.not_analyzed` / `assumptions` / `diagnostics` で明示する。

---

## Tasks

- [x] **実装前に spec 整合性を確認する**
  - `docs/spec/mcp/tools/analyze-impact.md` を実装の正として読む
  - `docs/spec/mcp/overview.md` の tool overview / selection guidance に `analyze_impact` が含まれていることを確認する
  - `docs/spec/mcp/versioning.md` から future candidate としての `analyze_impact` が外れていることを確認する
  - `docs/spec/mcp/schema.md` / `errors.md` に共通型・error code 追加が必要か確認する
  - 足りないspec項目があれば、実装前に spec を更新する

- [x] **QueryService に `AnalyzeImpact` の外部契約型を追加する**
  - request:
    - `selector`
    - `change` discriminated object
    - `scope_modules` optional
    - `max_impacts` optional, default `200`
  - response:
    - `target`
    - `change`
    - `summary.by_severity`
    - `summary.by_fixability`
    - `summary.by_kind`
    - `impacts[]`
    - `coverage`
    - `assumptions`
    - `truncated`
    - `truncated_reasons`
    - `diagnostics`
  - enum:
    - severity: `breaking` / `warning` / `info`
    - fixability: `mechanical` / `suggested` / `manual_review` / `unknown`
    - suggested fix confidence: `high` / `medium` / `low`
  - unsupported selector は tool error ではなく、空 impacts + `unsupported_selector` diagnostic + coverage で返す

- [x] **`change` discriminated object の validation を実装する**
  - `rename`: `new_id` 必須
  - `remove`: 追加payload不要
  - `change_type`: `new_type` 必須
  - `change_contract`: `note` optional
  - `change_transition_target`: `new_to` / `new_action` の少なくとも一方を許容
  - `add`: `added_id` 必須
  - kind と payload の不正組み合わせを validation error にする

- [x] **AnalyzeImpact の orchestration を実装する**
  - selector を ObjectRef に resolve する
  - `scope_modules` による対象絞り込みを適用する
  - change kind / target kind から必要な collector を選ぶ
  - impact id を deterministic に採番する（例: `impact-001`）
  - `max_impacts` 到達時に `truncated=true` / `truncated_reasons=["max_impacts"]` を返す
  - summary を impacts から集計する
  - `coverage.analyzed` / `coverage.not_analyzed` を常に返す
  - `assumptions` を常に返す

- [x] **reference based impact collector を実装する**
  - `get_references` 相当の direct references を材料にする
  - `get_reference_tree` 相当の bounded graph traversal を材料にする
  - field / model / task / transition / state / event / actor / store の supported selector を扱う
  - `via` は最短到達経路の reference kind 配列として返す
  - direct reference と traversal 結果の重複 impact を dedupe する

- [x] **flow wiring impact collector を実装する**
  - `inspect(task).members.flow.entries` 相当の内部情報を読む
  - `flow_step_task_resolution` を分析する
    - task rename / remove / change_contract が flow step に与える影響を返す
  - `flow_param_field_resolution` を分析する
    - flow params の source/target 解決を確認する
    - node_return / main_param / foreach_item / implicit_join を扱う
  - type check は v1 では型 signature identity 比較に限定する
    - primitive literal 一致
    - model id 一致
    - model ↔ primitive / 別model は不一致
  - source location を impact に inline で含める

- [x] **sequence step impact collector を実装する**
  - sequence diagram inspect 相当の resolved steps を読む
  - `sequence_step_task_resolution` を分析する
  - task rename / remove / change_contract / transition action 変更が scenario step に与える影響を返す
  - sequence step の source location が取れない場合は `fixability=unknown` または diagnostic で明示する

- [x] **render output file impact collector を実装する**
  - 変更対象 object を含む render group / output file path を特定する
  - render output impact は file 粒度のみ返す
  - severity は原則 `info`
  - recommended_action は `brewprint render` 再実行を促す文言にする
  - md 内の表示差分は `coverage.not_analyzed` に `render_presentation_details` として残す

- [x] **change kind 別 handler を実装する**
  - `rename`
    - 参照解決不能になる箇所を `severity=breaking` として返す
    - source token 置換が一意に決まる場合のみ `fixability=mechanical`
    - `replace_reference` suggested fix を返す
  - `remove`
    - incoming references / flow step / sequence step / transition action を breaking として返す
    - 原則 `fixability=manual_review`
  - `change_type`
    - field type / task param / returns / flow_param の型 signature identity を確認する
    - structural compatibility は見ず、coverage.not_analyzed に残す
  - `change_contract`
    - task params / returns の shape 変更影響を返す
    - flow wiring / transition action / sequence step への影響を返す
    - semantic contract compatibility は見ず、coverage.not_analyzed に残す
  - `change_transition_target`
    - transition `to` / `action` の参照解決を確認する
    - state machine の意味的妥当性は見ず、manual_review に寄せる
  - `add`
    - name collision を確認する
    - type resolution を確認する
    - writer coverage を可能な範囲で確認する
    - impact analysis というより consistency check として summary / coverage を返す

- [x] **mechanical judgement gate を実装する**
  - `fixability=mechanical` の5要件を共通関数で判定する
    1. 置換対象 source location が一意に特定できる
    2. 置換前 token が source 上で一意
    3. 置換後 token が明確に1つに定まる
    4. 置換後の reference 解決先が変わらない
    5. YAML 構造を変えない単純 token 置換である
  - 1つでも欠けたら最低でも `suggested` に下げる
  - 不確実性が高い場合は `manual_review` または `unknown` にする
  - `suggested_fixes[].confidence` を必ず返す

- [x] **SourceLocation / source fallback を実装する**
  - impacts[].source は file / line / column を inline 必須で返す
  - source range が取れない場合の扱いを決める
    - impact 自体は返す
    - `fixability=unknown` または diagnostic を付ける
  - `source_preview` は v1 必須ではないため、実装する場合のみ optional として返す
  - `get_source` の snippet 取得責務とは分離する

- [x] **MCP wrapper に `analyze_impact` tool を追加する**
  - tool registry に `analyze_impact` を追加する
  - input JSON schema を discriminated object 形式に合わせる
  - output shape を spec に合わせる
  - validation error / diagnostics の返し方を既存toolと揃える
  - stdio JSON-RPC 経由の smoke test を追加する

- [x] **QueryService unit tests を追加する**
  - rename field: breaking + mechanical + replace_reference
  - remove task: transition_action / flow_step / sequence_step への breaking impact
  - change_type field: primitive mismatch / model id mismatch / matching case
  - change_contract task: flow_param / sequence step 影響
  - change_transition_target: new_to / new_action の resolution
  - add field: name_collision / type_resolution / writer_coverage
  - unsupported selector: empty impacts + diagnostic + coverage
  - `max_impacts` truncation
  - `scope_modules` filtering
  - `coverage.not_analyzed` が常に返ること
  - `summary` が impacts から正しく集計されること

- [x] **MCP wrapper tests を追加する**
  - valid request / response roundtrip
  - invalid change payload の validation error
  - unsupported selector response
  - `max_impacts` default 適用
  - diagnostics が JSON-RPC response に含まれること

- [x] **fixture / UC を必要最小限追加する**
  - flow step / flow param / foreach / implicit_join を含む fixture
  - sequence diagram step を含む fixture
  - transition action を含む fixture
  - render output file mapping を確認できる fixture
  - 既存UCで足りる場合は重複 fixture を増やさない

- [x] **実装後 docs を更新する**
  - [x] 実装で spec と異なる判断が出た場合は spec を修正する
  - [x] `docs/impl/go-m13-summary.md` を作成する
  - [x] `gofmt ./...` / `go test ./...` の通過結果を `docs/impl/go-m13-summary.md` に追記する
  - [x] M13完了時に `docs/tasks/m13-mcp-analyze-impact-implementation.md` を closed にする
  - [x] M13完了時に `docs/TASKS.md` の M13 status を closed にする
  - [ ] 必要なら ADR-056 の `impl commit` を更新する

