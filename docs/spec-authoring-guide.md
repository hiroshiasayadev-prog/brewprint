# Spec Authoring Guide

> Authoring guide ID: `spec-authoring`
> Boundary guide ID: `artifact-boundary`

この文書は、brewprint の spec を作成・更新するときの実践ルールとフォーマットをまとめる。

spec は現行仕様の唯一の正である。
`docs/doc-policy.md` はセッション開始時の入口方針のみを持つ。
ADR の起票・更新・フォーマットは `docs/adr-authoring-guide.md` を参照する。

---

## 1. 目的

spec は、brewprint の現在有効な仕様を記録する文書である。

spec には、実装・LLM・読者が現在の振る舞いを判断するために必要な以下を書く。

- YAML schema
- field definition
- validation rule
- diagnostic code
- name resolution rule
- renderer rule
- MCP request / response schema
- project layout / file classification rule
- concept spec における artifact responsibility / canonical reference boundary

spec は、設計判断の長い背景、却下した代替案、作業 checklist、fixture migration 状態の代替文書ではない。
それらは ADR / task file / UC docs / impl notes に置く。

---

## 2. spec に書くもの / 書かないもの

### spec に書くもの

- 現在の正しい仕様
- machine-readable に近いレベルの schema / rule / diagnostic
- 実装やMCP toolから参照できる粒度の定義
- 仕様の適用範囲
- 仕様上の例外やunsupported扱い
- 局所的な由来注記

### spec に抱え込まないもの

- 設計判断の長い背景
- 却下した代替案
- 議論の経緯
- task checklist
- fixture migration の作業状態
- 実装引継ぎの詳細

これらが必要な場合、spec には参照先を書くに留める。

### concept spec の入口

複数 artifact layer にまたがる責務境界や、tool-independent な concept boundary は `docs/spec/concepts/` が所有する。

- [`docs/spec/concepts/project-artifact-model/index.md`](spec/concepts/project-artifact-model/index.md): artifact layer、source of truth、trace / MCP の位置付け
- [`docs/spec/concepts/traceability/index.md`](spec/concepts/traceability/index.md): canonical reference resolution foundation と future trace scope の境界

Semantic trace MVP では、active semantic ref は `spec:` のみであり、`internal-design:` / external coverage artifact / semantic realization relation は future decision とする。詳細は ADR-088 と traceability concept spec を参照する。

---

## 3. spec フォーマット

### Front Matter

各specファイルの先頭に以下を置く。

```markdown
---
scope: docs/spec/ファイル名.md
status: confirmed / draft / wip
last_updated: YYYY-MM-DD
summary: >
  このdocが何を定義するかを3行以内で。
depends_on:
  - docs/adr/NNN-xxx.md   # 関連する設計判断（複数可）
---
```

### field の意味

- `scope`: このspecが定義するファイルパス。
- `status`: specの確度。
  - `confirmed`: 現行仕様として確定済み。
  - `draft`: 仕様案。まだ変更されうる。
  - `wip`: 作成途中。レビューや追記が必要。
- `last_updated`: 最終更新日。
- `summary`: このspecが何を定義するかの短い説明。
- `depends_on`: 関連するADR。spec全体の出自や主要判断を示す。

---

## 4. セクション末尾の由来注記

特定の決定に基づくセクションには、本文中末尾に由来注記を入れる。

```markdown
## ファイル分類ルール

brewprint YAMLは以下の3種別に分類される...

- ノード定義ファイル: `nodes:` キーをトップレベルに持つ
- View定義ファイル: `as:` キーをトップレベルに持つ
- render_index.yaml: ファイル名で識別する

> 由来: ADR-030 §決定, ADR-043 §2
```

Front Matter の `depends_on` はspec全体の出自を示す。
セクション末尾の由来注記は、局所的な仕様判断の出自を示す。

---

## 5. ADR との境界

spec は現在の正しい仕様を記録する。
ADR は起票時点の設計判断を記録する。

ADR に仕様案や具体例が書かれていても、それは起票時点のスナップショットである。
設計が accepted になり spec に反映されたら、現行仕様は spec を参照する。

### spec に置くべきもの

- 現在のスキーマ
- validation rule
- diagnostic code
- field definition
- MCP response schema
- renderer rule

### ADR に置くべきもの

- 決定の核
- 判断理由
- 却下した代替案
- 設計判断を説明するための具体例
- 起票時点の仕様案

---

## 6. spec 更新時の注意

- spec が破綻していれば即修正する。ADRと違い、specは遡及修正してよい。
- spec修正の経緯はcommit履歴で辿る。
- 関連ADRがある場合は `depends_on` または由来注記を更新する。
- 既存ADRの仕様詳細をspecへ移した場合、ADR側には概要とspec参照を残す。
- 仕様変更によりUCやgolden fixtureに影響する場合、該当UC / task fileへ影響を記録する。
- 実装が未追従の場合、spec本文で曖昧にせず、task fileやimpl noteで未実装状態を追跡する。

---

## 7. UC / fixture との境界

UC docs / fixture は実例と期待render結果を所有する。
spec は個別fixtureの作業状態を所有しない。

UCから一般仕様を抽出した場合、specには抽出された汎用ルールを書く。
個別UCのmigration順序や作業チェックリストは UC task file に置く。
Gap / evidence / sign-off / external relation artifact が必要になった場合は、ADR-088 に従って配置と責務を再判断する。MVP では external artifact 用 directory や authoring entrance を設けない。

---

## 8. 更新方針

この guide は docs 運用補助文書であり、必要に応じて更新してよい。

- spec format を変える場合はこの guide を更新する。
- doc-policy.md と矛盾する場合は、入口方針として doc-policy.md を確認し、どちらを正にするか明示して修正する。
- guide の方針を大きく変える場合は、必要に応じて ADR を起票する。
