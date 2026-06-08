# V01-ADR-091: Workflow artifact の work item / task 責務分離と legacy milestone 移行

- **status**: accepted
- **date**: 2026-05-26
- **depends_on**: V01-ADR-081, V01-ADR-083, V01-ADR-089, V01-ADR-090
- **supersedes**:
- **migrated_to_spec**:

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

V01-ADR-081 / V01-ADR-083 は、requirement が必要性を、work item が requirement を解消するための横断進捗を、task が具体作業を所有する artifact boundary を導入した。

一方、従来の `docs/tasks/m*.md` は、現在の整理でいう work item 相当の長期計画・phase 構成・具体 checklist・完了記録を、`milestone` と呼んで一文書にまとめていた。特に M15 のように複数の設計判断・仕様更新・実装・検証を包含する単位は、着手して短期に閉じる task としては大きすぎる。

本判断では、作業単位が過大であることは単なる文書上の不便ではなく、着手意欲と進捗の可視性を損ない、product development の継続性を阻害する execution risk として扱う。

V01-REQ-MCP-002 / V01-WORK-MCP-002 の dogfooding では、batch retrieval capability を以下の六つの短期 task に分解して実行した。

- 利用 evidence と調査要否の整理
- 採用判断と ADR 要否の確定
- public contract / spec 更新
- implementation
- runtime verification
- evidence / status / dogfooding review

この分割により、V01-ADR-090 の review で見つかった public schema の未確定点を implementation 前に解消し、work item 全体を `done` まで進められた。Mermaid flow と work item status による工程把握も有効だった。

また、workflow artifact の参照 identity については、physical path を恒常参照として保持すると file relocation や archive migration による stale reference を招く。既存 canonical reference 方針と同じく、workflow artifact 自身の identity には record ID-as-ref を用いる必要がある。

## 決定

### 1. Workflow の基本単位は `requirement -> work item -> task` とする

Requirement は「何が必要か」を所有する。

Work item は、source requirement を解消するために必要な作業フロー全体を所有する。Work item は implementation だけを束ねるものではなく、必要に応じて investigation、ADR 判断、spec 更新、internal design、YAML 更新、implementation、fixture、verification、close / evidence 同期の task を束ねる。

Task は、work item を前進させる具体的な実行単位を所有する。Investigation artifact、ADR、spec update、implementation、verification 等は、必要な場合に task の成果物として作成・更新される。

### 2. Task は短期に閉じられる粒度へ分割する

新規 task は、原則として着手後 `0.5d` から `3d` 程度で完了判定できる単位へ分割する。

3日を明らかに超える見込みの作業、または複数の独立した判断・成果物・検証境界を含む作業は、原則として複数 task または複数 work item へ分割する。

Task は少なくとも以下を持つ。

- parent work item
- source requirement
- status
- estimate
- dependency task ID（存在する場合）
- outputs
- goal
- done condition
- verification / evidence

この粒度基準は厳密な工数見積制度ではなく、着手可能性と達成感を損なう過大 task を早期に検出する guard とする。

### 3. Work item は task graph と全体進捗を所有する

Work item は、配下 task の canonical ID list と、task 間の順序・分岐・blocker・並列可能性を示す task flow を所有する。

Task flow の標準表現は Mermaid `flowchart` とする。Task flow は execution order と分岐を説明するための view であり、個々の task の status の正本ではない。

Work item は、要求解消全体がどの処理段階にあるかを示す status を持つ。例えば `decision_pending`、`design_spec_pending`、`implementation_pending`、`verification_pending`、`done` は work item 全体の状態を表す。

### 4. Task status の正本は task artifact とし、work item に checkbox を複製しない

各 task の完了状態の正本は、その task artifact の `status` field とする。

Work item 本文に、配下 task の完了状態を表す手動更新 checkbox または status copy を必須項目として持たせない。Work item の status、task graph、個々の task status がそれぞれ別の責務を持つ状態を維持し、同一の完了情報を複数文書で手更新しない。

将来、MCP が requirement / work item / task を record として取得・集約可能になった場合、checkbox 相当の進捗一覧は task status から導出する projection として提供できる。手書きの source of truth として導入しない。

### 5. Work item が従来 milestone の役割を引き取り、milestone を新しい artifact relation にしない

新形式では、requirement の解消に向けた到達点、作業フロー、配下 task の順序・分岐・close 条件は work item が所有する。従来 `milestone` と呼んでいた実行計画の役割に対応する新しい workflow artifact は `WORK-*` である。

`milestone` を新しい artifact layer、canonical identity、metadata field、または work item 間を束ねる relation として導入しない。必要に応じて task list や Mermaid 図の人間向けタイトル、あるいは旧 M-series 記録を説明する歴史的ラベルとして用いるに留める。

従来の `docs/tasks/m*.md` は、起票時点では legacy milestone-shaped work record として残存する。既存文書の archive 移動、open legacy record の `WORK-*` / `TASK-*` への分解、`docs/TASKS.md` の再構成は、本判断に追従する migration work として別途扱う。

特に M15 は大きな open legacy record であり、本ADRの accepted 化のみを理由に機械的に変更・分割しない。移行対象として明示的に計画してから扱う。

### 6. Workflow artifact の恒常参照は ID-as-ref のみ support する

Requirement / work item / task の identity と、それらの相互 relation では canonical artifact ID-as-ref を用いる。

| relation | canonical reference form |
|---|---|
| requirement -> work item | `WORK-*` |
| work item -> source requirement | `REQ-*` |
| work item -> task | `TASK-*` |
| task -> work item | `WORK-*` |
| task -> source requirement | `REQ-*` |
| task -> dependency task | `TASK-*` |

Workflow artifact 間の canonical relation として physical path は support しない。File path は実装内部の所在であり、恒常 identity / relation として metadata または authoring format に要求しない。

本ADRが新たに確定する canonical relation は、workflow artifact 間の `REQ-*` / `WORK-*` / `TASK-*` ID-as-ref に限る。Workflow artifact から spec 等の他 artifact への参照規則は既存の canonical reference 方針に従い、本ADRでは追加拡張しない。`req:` / `work:` / `task:` の semantic prefix は導入しない。Workflow artifact 本体の取得・resolve・validation には `REQ-*` / `WORK-*` / `TASK-*` ID-as-ref を用いる。

### 7. MCP support は後続 requirement で扱う

本ADRは artifact identity と authoring boundary を決定するが、Design Records MCP が `REQ-*` / `WORK-*` / `TASK-*` を record kind、`get_record(s)` 対象、`resolve_reference` 対象、validation 対象として公開する contract を直接決定しない。

この MCP support は `V01-REQ-MCP-003` の対象として、上記 ID-as-ref と physical path 非対応の方針を前提に設計する。

## 理由

### なぜ work item が investigation / ADR / spec / implementation を束ねるか

Requirement を解消するには、implementation より前に調査や判断が必要な場合がある。Work item を implementation 用に限定すると、判断が確定する前の進捗・依存・close 条件を束ねる artifact が失われる。

V01-WORK-MCP-002 では、独立 investigation は不要と判断しつつ、V01-ADR-090 の起票、spec contract 固定、implementation、runtime verification、close evidence を同一 flow として管理できた。この実例は、work item が requirement resolution flow 全体を所有することの有効性を示す。

### なぜ task を短期粒度にするか

巨大な作業文書は、論理的には網羅的でも、次に閉じるべき成果が見えにくくなり、着手負荷を高める。着手できない計画は、仕様が正しくても product を前進させない。

短期 task は、判断前に implementation へ進むことを防ぎ、完了・block・次行動を局所的に判定できる。V01-WORK-MCP-002 では、contract refinement と implementation が分離されていたため、未確定 schema を implementation 前に直せた。

### なぜ checkbox を正本にしないか

Task artifact の `status` と work item の checkbox が同じ完了状態を手動で持つと、一方だけ更新された場合に整合性が崩れる。Visual progress は有用だが、source of truth を増やす理由にはならない。

Status aggregation が必要なら、将来の MCP support により task status から導出すべきである。

### なぜ physical path を support しないか

Workflow artifact の identity は document location ではなく、requirement / work / task としての stable identity である。Physical path を relation として保持すると、directory 再編、legacy M-series record の archive migration、命名整理で relation が壊れる。

既存の canonical reference foundation と揃え、workflow artifact 本体および workflow artifact 間 relation は ID-as-ref で参照することで、location と identity を分離できる。他 artifact への参照規則は既存方針に委ね、本ADRで新たに定義しない。

## 却下した代替案

### 代替案A: Milestone 文書をそのまま task として維持する

却下する。M15 のように複数判断・複数 phase・複数成果物を含む文書は、着手して短期に閉じる具体 task として過大である。

### 代替案B: Work item を implementation phase のみに限定する

却下する。Investigation、判断、spec 更新を必要とする requirement で、実装前の進捗と依存を管理できない。

### 代替案C: Work item に task 完了 checkbox を手動管理する

却下する。Task artifact の `status` と二重管理になり、stale progress state を生む。将来必要なら tool-derived projection として扱う。

### 代替案D: Workflow artifact の relation に physical path を用いる

却下する。Location change により relation が stale になる。`REQ-*` / `WORK-*` / `TASK-*` ID-as-ref を canonical relation とする。

### 代替案E: Workflow artifact 用 semantic prefix を導入する

却下する。Record 本体には既に stable ID があり、`req:` / `work:` / `task:` を導入すると identity を二重化する。Section-level semantic addressing の concrete requirement が生じるまでは導入しない。

## 影響

### Project artifact model / authoring guidance への影響

`docs/spec/concepts/project-artifact-model/index.md` は、task が milestone completion を所有するという現行記述を修正し、work item が requirement resolution flow、到達点、task graph を所有することを反映する必要がある。Milestone は新しい artifact layer または relation として追加しない。

`docs/work-items/README.md` は、Mermaid task flow、canonical task ID list、work item status と task status の責務分離、physical path 非対応を authoring guidance に反映する必要がある。

Task authoring guidance は、短期粒度、metadata、ID-as-ref relation、status 正本の方針を保持する入口文書が必要となる。

### Existing milestone-shaped records への影響

従来の `docs/tasks/m*.md` と `docs/TASKS.md` は、work item 相当の計画を task file / milestone index として保持していた legacy layout として migration 対象となる。ただし本ADRは、その即時 archive 移動や open legacy record の分解を完了条件としない。

### V01-REQ-MCP-003 への影響

Requirement / work item / task の MCP support を設計する際、canonical reference input と relation validation は `REQ-*` / `WORK-*` / `TASK-*` ID-as-ref を前提とし、physical path を supported canonical relation に含めない。

## Evidence

- commit: tbd
- impl commit: 該当なし
- 参考: V01-ADR-081, V01-ADR-083, V01-ADR-089, V01-ADR-090, V01-REQ-MCP-002, V01-WORK-MCP-002, V01-TASK-MCP-002-06
