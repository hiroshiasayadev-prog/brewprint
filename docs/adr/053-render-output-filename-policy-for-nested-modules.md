# 053: nested module の render output filename は local ID + collision error とする

- **status**: accepted
- **date**: 2026-04-30
- **depends on**: ADR-043, ADR-045, ADR-046, ADR-050

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-043 / ADR-046 により、brewprint の render 出力は `renders/{group-id}/...`、`renders/_cross/...`、`renders/_preview/...` に配置されることが決まっている。
また、ADR-045 により、`render_index.yaml` の group 定義では最上位 module を指定し、nested module は親 module の group に含める方針が決まっている。

一方で、nested module 内の render 出力ファイル名をどうするかは未確定だった。

例:

```text
yaml/payment/webhooks/task/process_payment.yaml
```

この task の DAG render を、同一 group 直下に短い名前で出すのか、module path をファイル名やディレクトリ構造に反映するのかを決める必要があった。

候補は大きく3つある。

1. local ID filename を維持する
   - 例: `renders/commerce/dag-process_payment.md`
2. module path を filename に含める
   - 例: `renders/commerce/dag-payment-webhooks-process_payment.md`
3. group 内に module path の subdirectory を作る
   - 例: `renders/commerce/payment/webhooks/dag-process_payment.md`

ドキュメント生成ツール一般では source tree を出力URLや出力pathへ反映する設計が多いが、brewprint の `renders/` は source markdown のHTML化結果ではなく、semantic object から導出した view 成果物である。
さらに `render_index.yaml` により、source module tree と異なる group 配置を許容している。

そのため、現時点で source path preserving を採用すべきかは dogfood 前に判断しづらい。

## 決定

v1では、render 出力ファイル名に nested module path を含めない。
render 出力ファイル名は semantic object の local ID から決定する。

基本形式:

```text
dag-{task-id}.md
state-{fsm-id}.md
seq-{scenario-id}.md
wireframe-{fsm-id}-{state-id}.html
```

例:

```text
yaml/payment/webhooks/task/process_payment.yaml
→ renders/commerce/dag-process_payment.md
```

ただし、同一 group 内で複数の render output が同一 relative path に解決される場合、renderer / placement validation は error として停止する。

```text
renders/commerce/dag-process_payment.md
```

に対して複数の source object が対応する場合、silent overwrite は禁止する。

この collision error は render placement 層のエラーとして扱う。
semantic validation diagnostic に載せるか、render CLI の placement error として返すかは実装側で決めてよいが、少なくとも `brewprint render` は失敗し、既存ファイルを黙って上書きしてはならない。

## 理由

### local ID filename を維持する理由

- v1 の既存 fixture と実装を大きく変えずに済む
- 出力ファイル名が短く、人間が読みやすい
- `render_index.yaml` の group が利用者向けの整理単位であり、source module path を必ずしも出力pathに反映する必要がない
- nested module の path-preserving 方式は、dogfood 後に実需を見て判断した方がよい

### collision error を必須にする理由

local ID filename を維持すると、同一 group 内に同名 task / state / scenario が存在する場合に output path が衝突しうる。
このとき silent overwrite を許すと、render 出力が欠落しても気づけない。

そのため、衝突を検出した時点で render を止める。

これは以下の3原則を満たすためである。

1. 出力pathは安定していること
2. source構造または設定から決定的に導出できること
3. 衝突時に黙って上書きしないこと

### path-preserving を v1 で採用しない理由

module path を filename や directory に反映すると collision には強くなるが、以下の欠点がある。

- 既存 render fixture / index / link への影響が広い
- 出力pathが長くなる
- `render_index.yaml` で group 再配置する思想と、source module path を強く保持する思想が衝突する可能性がある
- 今のUC-001だけでは path-preserving が本当に読みやすいか判断しきれない

将来、実プロジェクトで同一group内の衝突が頻発する場合は、新ADRで path-preserving 方式へ変更してよい。

## 影響

- `docs/spec/project-layout.md` の nested module render 出力ファイル名未確定記述を置き換える
- renderer / placement validation で output path collision を検出する実装タスクが必要
- collision error の表示形式は実装時に決める。ただし silent overwrite 禁止は仕様として固定する
- 将来 path-preserving 方式へ移行する場合、本ADRを supersede する新ADRを起票する

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: ドキュメント生成ツールの安定出力path慣習
