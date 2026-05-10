# ADR Authoring Guide

この文書は、brewprint の ADR を起票・レビュー・更新するときの実践ルールをまとめる。

ADR のフォーマットそのものは `docs/doc-policy.md` を正とする。
この guide は、ADR を書くときの判断、責務境界、実例の扱い、アンチパターンを補足する。

> 由来: ADR-068

---

## 1. 目的

ADR は **Architecture Decision Record** であり、設計判断の記録である。

ADR には、その判断を理解するために必要な以下を書く。

- 背景
- 決定
- 理由
- 却下した代替案
- 影響範囲
- evidence

ADR は、現行仕様そのもの、作業チェックリスト、fixture migration 状態、実装引継ぎメモの代替文書ではない。

---

## 2. ADR に書くもの / 書かないもの

### ADR に書くもの

- なぜその判断が必要になったか
- 何を決めたか
- なぜその形を選んだか
- どの代替案を却下したか
- どの spec / implementation / UC / task に影響するか
- どの証拠・実例・過去ADRに基づくか

### ADR に抱え込まないもの

- 現在の完全な仕様本文
- 作業 checklist
- migration の細かい順序
- fixture の作業状態
- 実装手順の詳細
- renderer / validation / MCP tool などの実装引継ぎ本文

これらが必要な場合、ADR には影響範囲と参照先を書くに留める。
詳細は、それぞれの責務を持つ文書に置く。

---

## 3. 文書間の責務境界

brewprint の docs では、文書間の関係を以下の3概念で整理する。

- **所有**: その文書が一次情報として保持し、更新責任を持つ内容
- **参照**: 他文書が一次情報として持つ内容を、必要に応じてリンク・言及すること
- **影響範囲**: その決定により更新・実装・検証が必要になる領域

ADR は、spec / task file / UC docs / fixture / impl notes など、他責務 artifact の詳細を所有しない。
必要な場合は、影響範囲と参照先を示すに留める。

| 文書 | 所有するもの | 参照のみ示すもの |
|---|---|---|
| ADR | 設計判断、背景、理由、却下案、影響範囲 | 現行仕様本文、具体task、fixture migration詳細、実装手順 |
| spec | 現在の正しい仕様 | ADRの判断背景（由来注記） |
| task file | 作業項目、順序、完了条件、実装・migration checklist | 設計判断の長い理由、現行仕様本文 |
| UC docs / fixture | 実例、coverage、gap発見ログ、migration状態 | 汎用仕様そのもの、設計判断の最終根拠 |

この境界により、ADR は時間が経っても「なぜその判断をしたか」を保持し、現行仕様・作業状態・実例の変化はそれぞれの責務文書で更新できる。

---

## 4. 実例・実装・fixture を根拠にするとき

実例・実装・fixture から設計判断を抽出する ADR では、具体例を恒久仕様として扱わない。

ADR が判断根拠として扱うべきなのは、個々の YAML shape や一時的な実装状態そのものではなく、そこから抽出された設計上の性質である。

### 書くべきこと

- その実例が起票時点の evidence であること
- 実例が将来 migration / implementation / spec 反映で変わりうること
- ADR が根拠にする性質は何か
- fixture migration の範囲や順序は task file / UC task file で管理すること

### 推奨文例

```markdown
本文中の具体例は、ADR起票時点の実例であり、恒久仕様ではない。
本ADRが判断根拠として扱うのは、具体的な shape ではなく、そこから抽出された設計上の性質である。
具体的な fixture migration の範囲や順序は task file で管理する。
```

必要に応じて、対象UCや fixture の名前を入れる。

```markdown
本ADRで参照する UC-002 の YAML は、ADR起票時点の self-hosting 実例である。
これらの YAML は今後の spec反映・implementation・UC-002 migration によって更新されうるため、本文中の具体的な field 名や `any` / `str + note` の配置は恒久仕様ではない。
```

---

## 5. 観測事実と設計判断を分ける

ADR では、実例から見つかった観測事実と、それに基づく設計判断を分けて書く。

### 悪い例

```markdown
UC-002 の `analyze_impact_response.impacts` が `any` なので `impact_entry` を作る。
```

この書き方は、起票時点の fixture shape をそのまま仕様判断にしている。
fixture が変わると ADR の根拠も壊れて見える。

### 良い例

```markdown
UC-002 では、有限語彙を持つ値集合が `str + note` に閉じており、machine-readable constraint として扱えない。
そのため、有限語彙を表す model kind が必要である。
```

この書き方では、fixture の具体 shape ではなく、そこから抽出された性質を判断根拠にしている。

---

## 6. 影響範囲と後続作業の書き方

ADR の「影響」には、後続作業が発生する事実や影響範囲を書く。
ただし、具体的な checklist、順序、完了条件は task file に置く。

### 良い例

```markdown
### M15 への影響

本ADRにより、M15 Phase C では enum model の spec / implementation / UC-002 fixture migration が必要になる。
具体的な作業項目、順序、完了条件は `docs/tasks/m15-data-layer-expressiveness.md` および UC-002 側の task file で追跡する。
```

### 悪い例

```markdown
### M15 への影響

- [ ] parser に enum kind を追加する
- [ ] semantic model に Values を追加する
- [ ] diagnostics に invalid_enum_model を追加する
- [ ] UC-002 YAML を更新する
```

これは task file の責務であり、ADR 本文に置くと作業状態の変化で stale になりやすい。

---

## 7. spec との境界

ADR は起票時点の決定を記録する。
spec は現在の正しい仕様を記録する。

ADR に仕様案や具体例を書いてよいが、それは起票時点のスナップショットである。
設計が accepted になり spec に反映されたら、現行仕様は spec を参照する。

### ADR に書いてよいもの

- 決定の核
- 判断理由
- 却下した代替案
- 設計判断を説明するための具体例
- 起票時点の仕様案

### spec に置くべきもの

- 現在のスキーマ
- validation rule
- diagnostic code
- field definition
- MCP response schema
- renderer rule

---

## 8. reference examples

### ADR-067: 実例由来の設計判断

ADR-067 は、UC-002 self-hosting の実例から enum model 導入判断を抽出した例である。

特に参考になる点:

- UC-002 の実例を evidence として扱った
- 実例の YAML shape が将来 migration で変わりうることを明記した
- ADR が根拠にするのは具体 shape ではなく、有限語彙が `str + note` に閉じていたという設計上の性質だと整理した
- 後続 task list を ADR 本文から外し、M15 / UC-002 task file で追跡する方針にした

この例は、authoring guide 初版で参照しうる初期 example の一つである。
将来より適切な例が増えた場合、このセクションは追加・差し替えしてよい。

---

## 9. アンチパターン

### 1. fixture shape をそのまま仕様判断にする

```markdown
この YAML では field が `any` なので、新しい model を作る。
```

問題:

- 起票時点の fixture shape に判断が依存する
- migration 後に ADR が古く見える
- 抽出された設計上の性質が分からない

改善:

```markdown
この実例では、閉じた値集合が machine-readable constraint として表現できない。
そのため、有限語彙を表す model が必要である。
```

### 2. ADR に作業 checklist を置く

```markdown
- [ ] spec を更新する
- [ ] parser を直す
- [ ] tests を追加する
```

問題:

- 作業状態が変わるたびに ADR が stale になる
- task file と責務が重複する

改善:

```markdown
本ADRは spec / implementation / tests に影響する。
具体的な作業項目と完了条件は `docs/tasks/...` で追跡する。
```

### 3. ADR に現行仕様本文を長く抱え込む

問題:

- spec と二重管理になる
- 後続 spec 更新とずれやすい
- ADR が「現在の仕様」と誤読される

改善:

- ADR には決定の核と理由を書く
- 現行仕様は spec に置く
- ADR から spec へ参照を張る

### 4. 実装手順を ADR に書きすぎる

問題:

- 実装変更により ADR が stale に見える
- impl docs / code review / handoff memo と責務が重複する

改善:

- ADR には実装への影響範囲を書く
- 実装手順や引継ぎは `docs/impl/` または code 側で管理する

---

## 10. 更新方針

この guide は docs 運用補助文書であり、ADR と異なり必要に応じて更新してよい。

- 新しい実例が出たら reference examples を追加する
- 良い文例・悪い文例は改善してよい
- doc-policy.md と矛盾する場合は doc-policy.md を優先し、この guide を修正する
- guide の方針を大きく変える場合は、必要に応じて ADR を起票する
