# 046: State / Sequence / Wireframe / Preview のrender出力配置

- **status**: accepted
- **date**: 2026-04-26
- **supersedes**:

## 背景

ADR-043で `renders/` の基本構造を定義し、ADR-045で `render_index.yaml` によるgroupingを定義した。
UC-001の移行レビューを通じて、ADR-043 / ADR-045だけでは以下の配置規則が不足していることが分かった。

1. **State Diagramのファイル名粒度が曖昧**
   - ADR-043のディレクトリ図では `state-{state-id}.md` と書かれている。
   - 一方、`docs/spec/views/state-diagram.md` は「1ファイル = 1FSM = 1枚の図」と規定している。
   - さらに、将来的には同一module内に複数FSMファイルを許可したい。
   - したがって、stateノード単位でもmodule単位でもなく、FSMファイル単位の安定IDが必要である。

2. **Sequence scenarioの所属groupが未定義**
   - sequence scenarioは `yaml/views/scenarios/` 配下にあり、module直下には属さない。
   - ただしADR-032により、sequence scenarioは必須フィールド `state_file` を持つ。
   - `state_file` が指すFSMファイルのmoduleから所属groupを推論できるが、このルールはまだADR-043 / ADR-045に明記されていない。

3. **Wireframe renderの配置が未定義**
   - ADR-029 / ADR-042 / `docs/spec/views/wireframe.md` により、wireframe DSLとHTML fragment renderは定義済みである。
   - しかしADR-043の `renders/` 構造には `wireframe-*` の出力先がない。
   - Wireframeはstateノードに属するため、State Diagramとは異なりstate単位でrenderされる。
   - 同一module内に複数FSMファイルを許可する場合、state IDだけでは同一group内で衝突しうるため、FSMファイル単位のIDを含める必要がある。

4. **Preview HTMLの位置づけが未定義**
   - UC-001には `docs/preview-wireframe.html` が存在する。
   - これは個別wireframe HTML fragmentとは異なり、複数wireframeをブラウザで確認するためのpreview harnessである。
   - 人間が見るためのHTMLであること自体は `renders/` から除外する理由にならない。Markdown renderも人間とLLMの双方が読む成果物である。
   - 分類基準は「人間向けかどうか」ではなく、「Go rendererがYAMLから決定論的に生成するcanonical / auxiliary render outputかどうか」とする必要がある。

## 決定

### 1. FSM IDの生成規則

State Diagram / Sequence / Wireframe が参照するFSMファイルには、`fsm-id` を導入する。

`fsm-id` は `yaml/` ルートからの相対パスをもとに、以下の規則で生成する。

1. 拡張子 `.yaml` / `.yml` を除く
2. 末尾のファイル名が `state` の場合、その `state` セグメントを除く
3. path separator（`/` または `\\`）を `-` に変換する

例:

```text
yaml/auth/state.yaml                → auth
yaml/order/state.yaml               → order
yaml/order/payment_state.yaml       → order-payment_state
yaml/order/admin/refund_state.yaml  → order-admin-refund_state
```

`state.yaml` を特別扱いするのは、現行の標準配置でmodule直下の `state.yaml` がそのmoduleの代表FSMを表すためである。
一方、`payment_state.yaml` のような明示ファイル名はFSM IDに残す。

### 2. State Diagramの出力ファイル名

State DiagramはFSMファイル単位でrenderする。

出力ファイル名は以下とする。

```text
renders/{group-id}/state-{fsm-id}.md
```

例:

```text
yaml/auth/state.yaml                → renders/auth/state-auth.md
yaml/order/state.yaml               → renders/commerce/state-order.md
yaml/inventory/state.yaml           → renders/catalog/state-inventory.md
yaml/order/payment_state.yaml       → renders/commerce/state-order-payment_state.md
yaml/order/admin/refund_state.yaml  → renders/commerce/state-order-admin-refund_state.md
```

出力先groupはFSMファイルのtop-level moduleから決定する。
明示group / 暗黙groupの解決はADR-045と同じ規則に従う。

ADR-043の `state-{state-id}.md` は、実運用上は `state-{fsm-id}.md` を意図していた表記として補正する。

### 3. Sequence scenarioの所属group

Sequence scenarioは、`state_file` が指すFSMファイルのtop-level moduleから所属groupを推論する。

`state_file` は `yaml/` ルートからの相対パスとして解釈する。

```yaml
as: sequence_diagram
id: checkout_flow
state_file: order/state.yaml
```

この場合、top-level moduleは `order` であり、`render_index.yaml` で `order` を含むgroupに出力する。

```text
yaml/views/scenarios/checkout_flow.yaml
  state_file: order/state.yaml
  order ∈ commerce
  → renders/commerce/seq-checkout_flow.md
```

`state_file` のtop-level moduleが明示groupに含まれない場合は、ADR-045のuncovered moduleルールに従い、暗黙group（group id = module名）に出力する。

`state_file` はADR-032で必須フィールドであるため、省略されているsequence scenarioはparser errorとする。

### 4. Wireframe renderの出力ファイル名

Wireframe renderは、wireframeを持つstateノード単位でHTML fragmentとして出力する。

出力先groupは、そのstateノードを定義しているFSMファイルのtop-level moduleから決定する。
明示group / 暗黙groupの解決はADR-045と同じ規則に従う。

出力ファイル名は以下とする。

```text
renders/{group-id}/wireframe-{fsm-id}-{state-id}.html
```

例:

```text
yaml/auth/state.yaml の login_screen              → renders/auth/wireframe-auth-login_screen.html
yaml/auth/state.yaml の loading                   → renders/auth/wireframe-auth-loading.html
yaml/order/state.yaml の cart                     → renders/commerce/wireframe-order-cart.html
yaml/order/state.yaml の checkout_screen          → renders/commerce/wireframe-order-checkout_screen.html
yaml/order/payment_state.yaml の processing       → renders/commerce/wireframe-order-payment_state-processing.html
```

この命名により、同一group内に複数FSMファイルが存在し、それぞれが同じstate IDを持つ場合でも衝突を避ける。

### 5. Preview HTMLの位置づけ

Preview HTMLは、個別diagram / fragmentのcanonical renderではなく、複数のrender結果をブラウザで確認するためのpreview harnessである。

ただし、Go rendererがYAMLおよびrender結果から決定論的に生成する成果物である場合は、`renders/` 配下に置く。

Preview harnessの出力先は以下とする。

```text
renders/_preview/wireframe.html
```

`_preview/` は通常groupではなく、preview用途の特殊ディレクトリである。
`_cross/` と同様、アンダースコアプレフィックスによりgroup IDと区別する。

group idに `_preview` を含むアンダースコア始まりのIDを使うことは、ADR-043 / ADR-045の方針と同じくvalidation errorとする。

既存の `docs/preview-wireframe.html` は、Go rendererで決定論的に生成できるpreview harnessとして扱う場合、`renders/_preview/wireframe.html` へ移行する。
手書きの試作HTMLとして残す場合は `docs/` 配下に置いてもよいが、その場合はcanonical render fixtureとはみなさない。

## 理由

### State DiagramをFSMファイル単位にする理由

State Diagramは個別stateノードではなくFSM全体を描く図である。
`docs/spec/views/state-diagram.md` も「1ファイル = 1FSM = 1枚の図」と規定している。
したがって、stateノード単位の `state-{state-id}.md` ではなく、FSMファイル単位の `state-{fsm-id}.md` とする。

### fsm-idをmodule-idではなくファイルパス由来にする理由

現行UC-001では `auth/state.yaml` / `order/state.yaml` / `inventory/state.yaml` のように1moduleにつき1FSMである。
しかし、将来的には同一module内に複数FSMファイルを許可したい。

`state-{module-id}.md` では、同一module内に複数FSMがある場合に出力名が衝突する。
`yaml/` ルート相対パス由来の `fsm-id` にすることで、現行の短いファイル名を維持しつつ、複数FSMにも対応できる。

### Sequence scenarioをstate_fileからgroup推論する理由

Sequence scenarioは、どのFSM transition列を辿るかを `state_file` で明示している。
この情報を使えば所属groupは決定的に決まる。

`render_index.yaml` に `groups[].scenarios` を追加する案も考えられるが、`state_file` とgroup設定の二重管理になり、scenarioの実体と配置設定が乖離する恐れがあるため採用しない。

### Wireframeにfsm-idを含める理由

Wireframeはstateノードの属性であり、「このstateの画面構造」を表す。
State DiagramはFSM全体を1枚で描くが、wireframeはstateごとに別のHTML fragmentとしてrenderする方が自然である。

ただし、同一group内に複数FSMファイルがあり、それぞれが同じstate IDを持つ可能性がある。
`wireframe-{state-id}.html` だけでは衝突しうるため、`wireframe-{fsm-id}-{state-id}.html` とする。

また、`docs/spec/views/wireframe.md` の出力契約はHTML fragmentであるため、拡張子は `.html` とする。

### Preview HTMLをrenders/_previewに置く理由

`renders/` は「人間向け」かどうかではなく、「YAMLからGo rendererが決定論的に生成するrender outputかどうか」で分類する。
Markdown renderも人間が読み、LLMも読む成果物であるため、人間向けであることは `renders/` から除外する理由にならない。

Preview HTMLは個別wireframe fragmentそのものではないが、Go rendererが決定論的に生成するpreview harnessとして扱える。
そのため、通常groupとは分けつつ `renders/_preview/` に置く。

## 影響

- ADR-043の `state-{state-id}.md` 表記は、本ADRに基づき `state-{fsm-id}.md` として補正する。
- UC-001の既存State Diagram renderは現状のまま以下に配置できる。
  - `renders/auth/state-auth.md`
  - `renders/commerce/state-order.md`
  - `renders/catalog/state-inventory.md`
- UC-001のsequence renderは以下に配置できる。
  - `renders/commerce/seq-checkout_flow.md`
  - `renders/commerce/seq-payment_webhook_flow.md`
- UC-001のwireframe renderは以下に配置できる。
  - `renders/auth/wireframe-auth-login_screen.html`
  - `renders/auth/wireframe-auth-loading.html`
  - `renders/commerce/wireframe-order-cart.html`
  - `renders/commerce/wireframe-order-checkout_screen.html`
- `docs/uc/001-ec-checkout-flow/docs/preview-wireframe.html` は、renderer生成のpreview harnessとして扱う場合、`renders/_preview/wireframe.html` に移動する。
- ADR-043およびdoc-policyのUC構造例は、本ADRの命名規則と `_preview/` 配置に追随して更新する必要がある。
- Go rendererはState / Sequence / Wireframe / Preview出力時に本ADRのgroup解決規則と命名規則に従う。

## Evidence
- commit: 881c92d
- impl commit: tbd
- 参考: 特になし
