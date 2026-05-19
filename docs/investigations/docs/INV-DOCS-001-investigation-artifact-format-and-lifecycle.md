# INV-DOCS-001: investigation artifact format and lifecycle

- **status**: concluded
- **date**: 2026-05-19
- **trigger**: ADR-085
- **scope**: investigation artifact の directory / ID / format / lifecycle / authoring boundary を調査する
- **non_scope**: Design Records MCP implementation, full semantic trace schema, requirement / work item / coverage schema の確定, investigation format の正式決定
- **source_refs**:
  - ADR-050
  - ADR-068
  - ADR-081
  - ADR-083
  - ADR-084
  - ADR-085
- **follow_up_candidates**:
  - docs/investigations/README.md
  - docs/doc-policy.md
  - docs/adr-authoring-guide.md
  - docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md
  - docs/tasks/m18-semantic-traceability-foundation.md

> この investigation は、ADR-085 により導入された `docs/investigations/` の初回調査 artifact である。
> 本文中の format / status / ID / lifecycle は、調査対象であり、起票時点では確定仕様ではない。
> 採用判断は、後続 ADR / README / doc-policy / authoring guide / task file が所有する。

## 調査スコープ

この investigation は、brewprint における investigation artifact の最小運用を定義するための判断材料を集める。

調査対象は以下である。

- `docs/investigations/` の directory 運用
- investigation ID 規則
- file name 規則
- front matter の required / optional field
- status 語彙
- 調査 scope / non-scope の明示方法
- 起点 artifact / trigger の記録方法
- 後続 artifact 候補の記録方法
- investigation から別 investigation を起票する場合の扱い
- requirement / work item / ADR / task / spec / internal design / coverage との責務境界
- docs/doc-policy.md / docs/adr-authoring-guide.md へ反映する最小内容
- 将来的に Design Records MCP または別 MCP index 対象にするか

## 非スコープ

この investigation では、以下は確定しない。

- Design Records MCP への `kind: investigation` 追加
- investigation 用 MCP tool contract
- semantic trace MVP の active prefix 拡張
- requirement schema の確定
- work item schema の確定
- coverage mapping schema の確定
- internal design format の確定
- investigation format の正式決定
- investigation をすべての変更の必須 gate にする運用

## 背景

ADR-085 は、複雑な変更における調査結果、根拠、影響範囲、未確定点、選択肢、後続 artifact 候補を保存する artifact layer として `docs/investigations/` を導入した。

ADR-085 では、investigation は決定、現行仕様、要求そのもの、横断進捗、完了状態、具体作業手順を所有しないと定めた。
また、investigation は requirement / work item / task / ADR / spec / internal design / coverage / 別 investigation の起票・更新前に必ず必要な gate ではなく、複雑な変更で調査結果を後続 artifact へ保存する必要がある場合にのみ使うと定めた。

一方で、ADR-085 は investigation の詳細 format、status 語彙、ID 規則、MCP index 方針、lifecycle を確定しない。
そのため、初回 investigation として本 artifact を起票し、最小運用の判断材料を調査する。

## 調査したもの

起票時点で参照対象とする既存 artifact は以下である。

- ADR-050: spec-first documentation policy
- ADR-068: ADR authoring guide
- ADR-081: project requirements layer と semantic traceability
- ADR-083: project artifact boundary と YAML as primary implementation source
- ADR-084: semantic trace MVP scope と artifact boundary
- ADR-085: investigation artifact boundary
- docs/doc-policy.md
- docs/adr-authoring-guide.md
- docs/tasks/m18-semantic-traceability-foundation.md
- docs/spec/concepts/traceability/*

## 調査項目ごとの確認結果

### Q1: directory 運用をどうするか

#### 確認対象

`docs/investigations/` 直下運用、domain subdir、時系列 subdir のどれを初期候補にするか。

#### 観測事実

- ADR / spec / task は、現在それぞれ責務別 directory に置かれている。
- investigation は初回導入直後であり、実例数がまだない。
- 初期から domain subdir を切ると、investigation の置き場所判断自体が追加負荷になる。

#### 候補

- A. `docs/investigations/` 直下に全 investigation を置く
- B. `docs/investigations/<domain>/` に分ける
- C. `docs/investigations/YYYY/` のように時系列で分ける

#### 判断に必要な観点

- file 数が増えたときの探索性
- 起票時の判断負荷
- semantic ID と physical path の結合度
- 将来 MCP index 対象にする場合の扱いやすさ

#### 後続判断先

- `docs/investigations/README.md`
- 必要なら後続 ADR

### Q2: ID 規則をどうするか

#### 確認対象

investigation ID を domain-scoped sequence にするか、単純連番にするか。

#### 観測事実

- ADR-081 は requirement ID の候補として `REQ-<DOMAIN>-NNN` を示している。
- ADR-084 は requirement / work item / coverage mapping ID の形式候補を domain-scoped sequence として予約している。
- ADR 番号と requirement ID は結合しない方針である。

#### 候補

- A. `INV-<DOMAIN>-NNN`
- B. `INV-NNN`
- C. ADR と同様に numeric only

例:

```text
INV-DOCS-001
INV-TRACE-001
INV-MCP-001
INV-YAML-001
```

#### 判断に必要な観点

- requirement / work item / coverage ID との一貫性
- domain をまたぐ investigation の扱い
- ID rename を避けられるか
- ADR 番号と混同しないか

#### 後続判断先

- `docs/investigations/README.md`
- 必要なら traceability spec / MCP index 方針

### Q3: file name 規則をどうするか

#### 確認対象

file name に ID を大文字で残すか、小文字化するか、連番だけにするか。

#### 観測事実

- ADR file は `NNN-slug.md` 形式である。
- investigation は ADR 番号体系とは別の ID を持つ候補である。
- requirement / work item / coverage も将来 domain-scoped ID を持つ可能性がある。

#### 候補

- A. `INV-DOCS-001-investigation-artifact-format-and-lifecycle.md`
- B. `inv-docs-001-investigation-artifact-format-and-lifecycle.md`
- C. `001-investigation-artifact-format-and-lifecycle.md`

#### 判断に必要な観点

- 目視で ID を確認しやすいか
- Windows / Git / tooling 上の扱いやすさ
- 既存 ADR file naming との違いを許容するか
- MCP index が path ではなく metadata / ID を見る前提にできるか

#### 後続判断先

- `docs/investigations/README.md`

### Q4: metadata / front matter の形式をどうするか

#### 確認対象

Markdown 冒頭の bullet metadata に何を置くか。YAML front matter 化するか。

#### 観測事実

- ADR は Markdown 冒頭に bullet metadata を置いている。
- spec は front matter を使うものがある。
- investigation は初期導入時点では Design Records MCP の index 対象ではない。
- machine-readable 化を急ぐと、実例前に schema を固定しやすい。

#### 候補

required field 候補:

- `status`
- `date`
- `trigger`
- `scope`
- `non_scope`
- `source_refs`
- `follow_up_candidates`

optional field 候補:

- `supersedes`
- `related_requirements`
- `related_work_items`
- `related_adrs`
- `related_specs`
- `related_internal_design`
- `related_coverage`
- `follow_up_results`

#### 判断に必要な観点

- 人間レビューしやすいか
- MCP index 対象化を後で妨げないか
- 必須 field が多すぎて軽量性を失わないか
- `scope` / `non_scope` を metadata と本文のどちらに置くべきか

#### 後続判断先

- `docs/investigations/README.md`
- 将来的に Design Records MCP / 別 MCP interface

### Q5: status 語彙をどうするか

#### 確認対象

investigation の lifecycle status をどう表すか。

#### 観測事実

- ADR status は `proposed` / `accepted` / `superseded` である。
- requirement status は要求の扱いを表す。
- work item status は layer 別進捗を表す。
- investigation status が採用判断に見えると、ADR の責務と衝突する。

#### 候補

- `investigating`
- `concluded`
- `superseded`
- `archived`
- `proposed`

#### 判断に必要な観点

- status が「決定」を意味してしまわないか
- 調査完了と後続 artifact 反映完了を混同しないか
- `concluded` と `archived` の違いが必要か
- `proposed` が ADR status と紛らわしくないか

#### 後続判断先

- `docs/investigations/README.md`

### Q6: scope / non-scope をどう書くか

#### 確認対象

investigation が何でも置き場にならないよう、調査範囲をどう明示するか。

#### 観測事実

- ADR-085 は investigation が必須 gate ではないことを明示している。
- investigation は後続 artifact の起点になりうるが、後続 artifact の責務を所有しない。
- scope が曖昧だと、調査中に requirement / work item / ADR / task の内容まで抱え込みやすい。

#### 候補

- metadata に短い `scope` / `non_scope` を置く
- 本文に `## 調査スコープ` / `## 非スコープ` を置く
- 起点 artifact / 後続 artifact / 判断未確定点を scope section に含める

#### 判断に必要な観点

- 起票時に書ける粒度か
- 長い scope を metadata に入れすぎないか
- 調査中に scope change が起きた場合の扱い
- 別 investigation への切り出し基準

#### 後続判断先

- `docs/investigations/README.md`
- 必要なら `docs/doc-policy.md`

### Q7: 起点 artifact と後続 artifact をどう記録するか

#### 確認対象

trigger / source_refs / follow_up_candidates / follow_up_results をどう分けるか。

#### 観測事実

- investigation は後続 artifact の起点になりうる。
- ただし、後続 artifact の採用判断や完了状態は ownership 外である。
- 調査根拠として読む artifact と、調査結果から生まれる artifact は意味が違う。

#### 候補

- `trigger`: なぜこの investigation が起票されたか
- `source_refs`: 調査根拠として読む artifact
- `follow_up_candidates`: 調査結果から起票・更新されうる artifact
- `follow_up_results`: 実際に作成・更新された artifact

#### 判断に必要な観点

- candidate と result を分けるべきか
- result を入れると進捗管理に見えないか
- semantic ref 化する時期
- physical path と semantic ref の混在をどう扱うか

#### 後続判断先

- `docs/investigations/README.md`
- requirements / work-items schema は後続判断待ち

### Q8: investigation から別 investigation を起票できるか

#### 確認対象

調査中に別領域の調査が必要になった場合の扱い。

#### 観測事実

- ADR-085 は別 investigation が後続 artifact になりうることを認めている。
- investigation は必須 gate ではないため、再帰的に常時要求すると運用が重くなる。
- scope が膨らみすぎると、調査 artifact が巨大化する。

#### 候補

- 別 investigation を許可する
- 元 investigation の scope を拡張して扱う
- 別 investigation 起票を原則禁止する

#### 判断に必要な観点

- scope creep を防げるか
- 後続調査の起点を追えるか
- investigation の無限再帰を防げるか
- follow-up candidate / result にどう記録するか

#### 後続判断先

- `docs/investigations/README.md`

### Q9: README / doc-policy / authoring guide の分担をどうするか

#### 確認対象

investigation format と責務境界をどの artifact に反映するか。

#### 観測事実

- docs/doc-policy.md はセッション開始時の入口方針である。
- docs/adr-authoring-guide.md は ADR の書き方を所有する。
- investigation の完全 format を doc-policy / adr-authoring-guide に入れると、それぞれの責務が膨らむ。
- ADR には探索ログや選択肢比較を抱え込まないための境界説明が必要である。

#### 候補

- `docs/investigations/README.md`: investigation format / lifecycle の実務ガイド
- `docs/doc-policy.md`: docs layer としての最小責務説明
- `docs/adr-authoring-guide.md`: ADR と investigation の境界説明

#### 判断に必要な観点

- セッション開始時に必要な情報か
- ADR authoring 時だけ必要な情報か
- investigation authoring 時に必要な詳細か
- 同じルールを複数文書に重複させないか

#### 後続判断先

- `docs/investigations/README.md`
- `docs/doc-policy.md`
- `docs/adr-authoring-guide.md`

### Q10: MCP index 対象にするか

#### 確認対象

Design Records MCP または別 MCP interface で investigation を扱うか。

#### 観測事実

- ADR-085 は、Design Records MCP への investigation 追加を必須変更とはしていない。
- ADR-084 は semantic trace MVP の active prefix を spec / internal-design / coverage に限定している。
- investigation は現時点で現行仕様や trace edge の source of truth ではない。

#### 候補

- MVP では MCP index 対象にしない
- Design Records MCP に `kind: investigation` を追加する
- 別 MCP interface として扱う

#### 判断に必要な観点

- read-only list / get / validate の必要性
- traceability MVP scope を広げすぎないか
- Design Records MCP の責務に合うか
- 別 MCP interface の方が自然か

#### 後続判断先

- 後続 ADR
- Design Records MCP spec / implementation task

## 横断的な観測事実

### investigation は後続 artifact の起点になりうる

investigation は、調査結果から requirement / work item / task / ADR / spec update / internal design / coverage mapping / 別 investigation を生む可能性がある。

ただし investigation は、これらの artifact の採用判断、現行仕様、進捗、完了状態を所有しない。

### investigation は必須 gate ではない

ADR-085 は、investigation をすべての変更の前提として要求しない。

単純な typo 修正、既に影響範囲が明確な作業、既存 task の細分化、軽微な requirement 捕捉などでは investigation を起票しない方がよい。

### investigation と work item は「影響範囲」の意味が異なる

investigation が記録する影響範囲は、調査上の仮説・根拠・不確実性・選択肢を含む分析結果である。

work item が所有する影響範囲は、layer 別の処理状態・完了管理・横断進捗である。

そのため、investigation は work item の代替ではない。

## 後続判断に渡す候補

この section は決定を所有しない。
後続 artifact が判断するための候補を整理する。

### directory / ID / file name

- `docs/investigations/<domain>/` 運用
- `INV-<DOMAIN>-NNN` ID
- `INV-<DOMAIN>-NNN-<slug>.md` file name

### metadata / body structure

- metadata に `status` / `date` / `trigger` / `scope` / `non_scope` / `source_refs` / `follow_up_candidates` を置く
- 本文に `## 調査スコープ` / `## 非スコープ` / `## 調査項目ごとの確認結果` / `## 横断的な観測事実` / `## 後続 artifact 候補` / `## 未確定点` を置く
- `follow_up_results` は実際に作成・更新された artifact を後から記録する任意 field として検討する

### status vocabulary

- `investigating`
- `concluded`
- `superseded`

`proposed` は ADR status と意味が近いため、採用する場合は混同しない説明が必要である。

### docs 反映方針

- `docs/investigations/README.md` に format / lifecycle の実務ガイドを置く候補
- `docs/doc-policy.md` には layer の存在と責務だけを最小反映する候補
- `docs/adr-authoring-guide.md` には ADR が探索ログを抱え込まないための境界だけを反映する候補

### MCP 方針

- MVP では MCP index 対象にしない候補
- 将来的に list / get / validate が必要になったら、`kind: investigation` または別 MCP interface を検討する候補

## 後続 artifact 候補

### 必須候補

- `docs/investigations/README.md`
  - investigation format / lifecycle の実務ガイドを置く候補。
- `docs/tasks/m18-semantic-traceability-foundation.md`
  - M18 の scope / references / task に investigation artifact を反映する候補。

### 任意候補

- `docs/doc-policy.md`
  - `docs/investigations/` の存在と責務を最小反映する候補。
- `docs/adr-authoring-guide.md`
  - ADR と investigation の境界を反映する候補。
- `docs/adr/083-project-artifact-boundary-and-yaml-as-implementation-source.md`
  - artifact placement decision rule に investigation 分岐を追加する候補。

### 後続判断待ち

- `docs/spec/concepts/traceability/*`
  - investigation を semantic trace schema に含めるかは未決。
- Design Records MCP
  - `kind: investigation` または別 MCP interface は未決。
- requirements / work-items schema
  - investigation 参照 field を持つかは未決。

## 未確定点

- investigation status の正式語彙
- `follow_up_results` field を持つか
- `source_refs` / `follow_up_candidates` を semantic ref 化するか
- README へ昇格する前に追加実例が必要か
- ADR-083 を直接更新するか、ADR-085 を後続 refinement として扱うか
- Design Records MCP で investigation を扱うか
- investigation format を ADR で確定するか、README で運用ルールとして確定するか
