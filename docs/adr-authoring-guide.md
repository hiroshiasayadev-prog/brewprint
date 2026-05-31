# ADR Authoring Guide

> Authoring guide ID: `adr-authoring`
> Boundary guide ID: `artifact-boundary`

この文書は、brewprint の ADR を起票・レビュー・更新するときの実践ルールとフォーマットをまとめる。

ADR のフォーマット、metadata、Evidence の書き方はこの guide を正とする。
`docs/doc-policy.md` はセッション開始時の入口方針のみを持つ。

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
- 判断前の探索ログ、影響範囲調査、選択肢比較、未確定論点の蓄積
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
| ADR | 設計判断、背景、理由、却下案、影響範囲 | 現行仕様本文、具体task、fixture migration詳細、実装手順、調査ログ、trace mapping 詳細 |
| investigation | 調査結果、根拠、影響範囲の仮説、未確定点、選択肢、後続 artifact 候補 | 設計判断、現行仕様、要求そのもの、作業進捗、完了状態 |
| requirement | 要求、不足、要望、stable requirement ID | 実装手順、完了状態、個別 trace mapping |
| work item | source requirement を解消する到達点、作業フロー全体、横断進捗、影響範囲、task graph | 現行仕様本文、設計判断の長い理由、個別 task の status 正本 |
| task file | 短期に閉じられる具体作業、完了条件、個別 status、verification evidence | 設計判断の長い理由、現行仕様本文、work item 全体の到達点・task graph |
| spec | 現在の正しい仕様、semantic ref declaration | ADRの判断背景（由来注記）、具体実装手順 |
| internal design | spec を実装へ落とす internal wiring / route | 設計判断、要求、横断進捗、MVP の semantic endpoint contract |
| external relation / assurance artifact | concrete requirement が成立した場合に配置・責務を含めて新設判断する将来候補（MVP operational scope 外） | 現行 MVP の canonical reference contract、各 artifact が所有する意味本文、fixture-level validation |
| UC docs / fixture | 実例、固定入力、期待出力、fixture-local 検証補助 | 汎用仕様、project-level trace relation、横断 gap / migration 進捗 |

この境界により、ADR は時間が経っても「なぜその判断をしたか」を保持し、現行仕様・作業状態・実例の変化はそれぞれの責務文書で更新できる。

### ADR と semantic trace の境界

ADR は判断と影響範囲を記録するが、canonical semantic ref の declaration や resolver / validation contract を所有しない。
Semantic trace MVP では、現行 `spec:` semantic ref と investigation canonical reference の contract は spec が所有し、internal design は implementation-facing wiring route として存続するが active semantic endpoint ではない。External coverage artifact と realization relation は concrete requirement が生じるまで future decision とする。

ADR が traceability に影響する場合は、更新対象の spec / internal design / 将来新設を判断しうる external artifact を影響範囲として示し、trace detail を ADR 本文に二重管理しない。
設計判断から現行仕様を辿る場合は、spec の由来注記と canonical reference contract を用いる。

ADR-091 に基づき、requirement / work item / task 間の canonical relation は `REQ-*` / `WORK-*` / `TASK-*` ID-as-ref を用い、physical path は supported canonical relation として扱わない。Workflow artifact から他 artifact への参照規則は既存の canonical reference 方針に従い、ADR authoring guide で追加拡張しない。

### ADR と investigation の境界

ADR は、設計判断とその理由を所有する。
investigation は、判断前の調査結果、根拠、影響範囲の仮説、未確定点、選択肢、後続 artifact 候補を所有する。

複雑な変更で、判断前の探索ログ、影響範囲調査、選択肢比較、未確定論点を長く保存する必要がある場合、ADR に抱え込まず investigation を起票または参照する。

ADR が investigation を根拠にする場合、ADR 本文では調査内容を再掲しすぎず、設計判断に必要な要点と investigation への参照に留める。

investigation は ADR の起票前に必ず必要な gate ではない。
単純な設計判断や、既に根拠・影響範囲が明確な変更では、ADR 単独で判断を記録してよい。

investigation の directory / ID / metadata / status / lifecycle / authoring format は `docs/investigations/README.md` が所有する。
ADR authoring guide には investigation の完全 format を重複して書かない。

---

## 4. ADR フォーマット

### ファイル名規則

```text
docs/adr/
  NNN-タイトル.md   （例: 001-node-type-splitting.md）
```

NNNは3桁ゼロ埋めの連番にする。

### 本文テンプレート

```markdown
# NNN: タイトル

- **status**: proposed / accepted / superseded
- **date**: YYYY-MM-DD
- **supersedes**: （該当する場合、旧ADR番号）
- **migrated_to_spec**: YYYY-MM-DD（spec移管済みの場合のみ）

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

なぜこの決定が必要だったか。

## 決定

何を決めたか。起票時点での仕様案・規定を含めてよい。
仕様詳細が別specに移管済み・新設予定の場合はリンクで参照させる。

> 仕様詳細: [docs/spec/xxx.md](../spec/xxx.md) §N （該当する場合）

## 理由

なぜそう決めたか。却下した代替案も書く。

## 影響

この決定が他の仕様・実装に与える影響。

## Evidence
- commit: <ADR起票時のcommit hash>
- impl commit: <実装反映時のcommit hash。未着手なら "tbd"。doc運用ADR等で該当しない場合は "該当なし">
- 参考: <"dagsterのassets参考" / "Goのinterface慣習" 程度の軽い記述。なければ省略>
```

冒頭の「起票時点のスナップショット」注記は新規ADRに入れることを推奨する。
ADR単体で読まれた場合に、現行仕様と誤読されることを防ぐためである。

### status の基準

- `proposed` — 議論中・まだ覆りうる
- `accepted` — 確定。変更する場合は新しいADRでsupersedesする
- `superseded` — 旧ADRの場合。新ADR番号をsupersedesに記載する

### Evidence の書き方

- **commit / impl commit**: git logから拾う。ADR起票と実装反映が同コミットなら1行でOK
- **artifact reference**: ADR / investigation / spec record を根拠として示す場合は、可能な限り canonical ID-as-ref（例: `ADR-088`, `INV-DOCS-003`）を用いる。Physical path は補助的な所在説明に限り、canonical reference として扱わない
- **参考**: OSS名・言語慣習名のレベル。URL/取得日は書かない
  - ✅ `dagsterのsoftware-defined assets参考` / `Goのinterface慣習` / `特になし`
  - ❌ `https://docs.dagster.io/... (retrieved 2026-04-19)` ← 不要

---

## 5. 実例・実装・fixture を根拠にするとき

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

## 6. 観測事実と設計判断を分ける

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

## 7. 影響範囲と後続作業の書き方

ADR の「影響」には、後続作業が発生する事実や影響範囲を書く。
ただし、具体的な checklist、順序、完了条件は task file に置く。

### 良い例

```markdown
### Data layer expressiveness work への影響

本ADRにより、対応する work item では enum model の spec / implementation / UC-002 fixture migration が必要になる。
具体的な作業項目、順序、完了条件は work item 配下の短期 task file で追跡する。
```

既存 `docs/tasks/m*.md` を参照する場合、それは ADR-091 以前の legacy milestone-shaped work record の所在を示す場合に限り、canonical workflow relation として扱わない。新形式に `milestone` artifact または relation を追加しない。

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

## 8. spec との境界

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

## 9. ADR metadata の機械可読性

ADR 冒頭の metadata block は Design Records MCP によって機械的に読まれるため、人間向け注釈を field value に混ぜない。

特に以下を守る。

- `status` は `proposed` / `accepted` / `superseded` のいずれかにする
- `status` に `accepted（一部superseded）` のような注釈を入れない
- `supersedes` は comma-separated `ADR-NNN` record ID list にする
- 置換対象がない場合、`supersedes` は空欄にする
- `なし` は書かない
- `013` のような裸番号は `ADR-013` と書く
- `ADR-010（一部）` のような注釈付き値は書かない
- 注釈や補足説明は metadata block ではなく本文に書く

悪い例:

```markdown
- **status**: accepted（一部superseded）
- **supersedes**: なし
- **supersedes**: 013
- **supersedes**: ADR-010（一部）
```

良い例:

```markdown
- **status**: accepted
- **supersedes**:
- **supersedes**: ADR-013
- **supersedes**: ADR-010
```

部分的な置き換えや補足関係がある場合は、metadata では ID のみを書き、詳細は本文で説明する。

---

## 10. reference examples

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

## 11. アンチパターン

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

## 12. 更新方針

この guide は docs 運用補助文書であり、ADR と異なり必要に応じて更新してよい。

- 新しい実例が出たら reference examples を追加する
- 良い文例・悪い文例は改善してよい
- doc-policy.md と矛盾する場合は doc-policy.md を優先し、この guide を修正する
- guide の方針を大きく変える場合は、必要に応じて ADR を起票する
