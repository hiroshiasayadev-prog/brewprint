# 081: project requirements layer と semantic traceability

- **status**: accepted
- **date**: 2026-05-16
- **depends_on**: ADR-050, ADR-068, ADR-083
- **supersedes**:
- **migrated_to_spec**:

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

brewprint は ADR-050 により spec-first のドキュメント運用へ移行した。
現行仕様は `docs/spec/**` に置き、ADR は設計判断の背景・理由・却下案を記録する。
ADR-068 では ADR / spec / task / UC docs の責務境界を整理した。

その後、仕様変更・self-hosting・fixture 整理・implementation 作業の途中で、追加要件、spec gap、改善要望、後回し判断が継続的に発見されることが分かった。
これらを会話ログ、task file、impl summary、UC notes に散らすと、後から「なぜこの変更をするのか」「何が未判断なのか」「どの work item の起点なのか」を辿れなくなる。

一方で、requirements を「対象 system が満たすべき system obligation spec」として扱うと、spec-first の責務が崩れる。
現行仕様の唯一の正は spec であり、requirements は仕様本文・実装設計・進捗・coverage を所有しない。

ADR-083 により、requirements / work-items / tasks / internal design / coverage / YAML の責務境界が整理された。
本ADRはその前提に立ち、`docs/requirements/` を「要求・不足・要望・spec gap 候補を捕捉する layer」として再定義する。

## 決定

### 1. brewprint project の artifact として `docs/requirements/` を導入する

brewprint project layout に、要求・不足・要望を置く場所として `docs/requirements/` を導入する。

requirements は、以下のような「何が必要か」を所有する。

- 追加要件
- user need
- self-hosting / fixture 作成中に見つかった不足
- editor / viewer / workflow 上の要望
- spec gap 候補
- 採用判断待ちの要求
- deferred / rejected された要求の履歴
- 純粋なリファクタリングや bugfix の起点となる「あるべき状態と現状の乖離」

requirements は、決定・現行仕様本文・内部設計・coverage・具体作業手順を所有しない。

| artifact | 所有するもの |
|---|---|
| ADR | spec 変更や project policy 変更の設計判断、背景、理由、却下案 |
| spec | 現行仕様の唯一の正 |
| requirements | 要求・不足・要望・spec gap 候補 |
| work-items | requirement を完了させるための横断進捗と影響範囲 |
| tasks | 具体的な作業手順・チェックリスト |
| internal design | spec semantics と YAML model を target implementation へ写像する wiring route |
| coverage | spec / internal design / YAML の semantic ref 間の対応関係 |
| YAML | 対象 design model / target implementation への primary DSL source |

### 2. requirements は source of truth ではない

requirements は、仕様本文の source of truth ではない。
requirements は YAML の source of truth でも、target implementation の source of truth でもない。

現行仕様は spec が所有する。
YAML authoring は spec に従う。
internal design は spec semantics と YAML model を target implementation へ写像する wiring route を記録する。
coverage は spec / internal design / YAML の対応関係を管理する。
work item は requirement を完了させるための進捗と影響範囲を管理する。

requirements は「この project で何が必要か、何が不足しているか、何を採用・延期・却下したか」を残すための artifact である。

採用済み requirement であっても、それだけでは現行仕様にならない。
現行仕様として外部 contract にする場合は、spec へ反映する。
実装や YAML の更新が必要な場合は、work item を起票して追跡する。

### 3. work item は source requirement を必ず持つ

ADR-083 に従い、work item は source requirement を必ず持つ。

requirement は work item の起点であり、work item は requirement を完了させるための横断進捗である。
requirement が「何が必要か」を持ち、work item が「どの artifact をどこまで更新するか」を持つ。

純粋なリファクタリング、bugfix、internal design 整理のような ADR を要さない変更も、実装側 requirement として requirements layer に捕捉し、そこから work item を引く。
work item が requirement なしで自走することは認めない。

この制約は、明示的な手間を増やす代わりに、暗黙的な手間を減らすためのものである。
後続作業者（将来の自分・別の人間・AI agent）が、変更理由を会話ログから再構築しなくてよい状態を保つ。

### 4. requirement status は要求そのものの扱いを表す

requirement は lifecycle status を持つ。

| status | 意味 |
|---|---|
| `captured` | 要求・不足・要望として捕捉済み |
| `decision_needed` | 採用可否や設計判断が必要 |
| `accepted` | 要求として採用する |
| `deferred` | 要求としては認識するが後回し |
| `rejected` | 採用しない |

requirement.status は、要求そのものの扱いを表す。
work item の layer 別進捗とは混ぜない。

例えば、requirement が `accepted` であっても、対応する work item が `yaml_pending` や `implementation_pending` のままである状態は正常である。
ある要求が project 上で実現されたとみなせるのは、対応する work item が必要な反映・検証を完了した時点である。

requirement.status が `deferred` / `rejected` に変わっても、過去の work item status を機械的に巻き戻さない。
work item status は、その work item 自体の作業状態を表す。

requirement.status の変更に伴って、過去の work item、YAML、internal design、coverage、tasks の整理・撤回・無効化が必要になる場合、その整理自体を別 requirement として捕捉し、必要に応じて work item を起票する。

### 5. requirement ID は ADR 番号と結合しない

requirement ID は ADR 番号を含めない。
requirement は ADR より長生きする可能性があり、後続 ADR / spec 更新によって意味が refinement されうるためである。

推奨形式は domain-scoped sequence とする。

```text
REQ-DAG-001
REQ-MCP-001
REQ-RESOLVE-001
REQ-SELFHOST-001
```

ADR 由来を requirement metadata に直接記録する field は設けない。
ADR との関係を知りたい場合は、requirement が参照する spec semantic ref から spec の由来注記や Design Records MCP を辿る。

### 6. requirement は semantic ref で関連 artifact を参照する

requirement は、必要に応じて関連する spec / YAML / internal design / coverage / work item を参照できる。

ただし、その primary key は physical path ではなく semantic ref とする。
これは ADR-083 の trace layer common principle に従う。

初期 schema では、少なくとも以下を候補とする。

```yaml
id: REQ-DAG-001
title: DAG asset TypeRef hint を導入したい
status: accepted
source_refs:
  specs:
    - spec:dag.asset-node-type-hint
related_refs:
  yaml:
    - uc:001.yaml.checkout-flow
  internal_design:
    - internal-design:resolver.type-ref-hint
work_items:
  - WORK-DAG-001
notes:
  - この requirement は例であり、正式 schema は後続 spec / ADR で定義する。
```

本ADRでは requirement schema の完全定義は行わない。
正式な field 名、semantic ref 命名規則、redirect / superseded mapping、MCP response への露出は後続 spec / ADR で定義する。

### 7. section anchor は補助情報として扱う

spec document 内の特定 section を参照したい場合、stable semantic anchor を補助情報として使える。

ただし、Markdown 見出しに Pandoc / kramdown 形式の `{#...}` を直接書く方式は採用しない。
GitHub Flavored Markdown / VS Code preview / 既存 Design Records MCP parser との互換性が弱く、見出しにも literal text として表示されうるためである。

section-level traceability が必要な場合、spec front matter に `sections` mapping を持たせる方針を第一候補とする。

例:

```yaml
---
scope: docs/spec/views/dag/asset-node-type-hint.md
status: confirmed
sections:
  spec:dag.asset-node-type-hint: Asset node TypeRef hint
---
```

本ADRでは `sections` metadata の完全 schema は定義しない。
正式な schema、anchor ID 命名規則、Design Records MCP response への露出は後続 spec / ADR で扱う。

### 8. section anchor は append-only に扱う

section anchor を発行する場合、以下の lifecycle rule に従う。

- 一度発行した anchor ID は別概念に再利用しない
- section rename では anchor ID を維持する
- section move では anchor ID を維持する
- section split では、旧 anchor を最も近い後継 section に残し、新しい概念には新 anchor を発行する
- section merge では、複数 anchor が同一 section / document を指してよい
- anchor の廃止が必要な場合は、redirect / superseded mapping を後続 schema で扱う

これにより、requirement trace の安定性を保つ。

### 9. requirements は lazy に起票する

すべての spec document / section に requirement を必ず作成する必要はない。

requirement artifact は、以下のような場面で lazy に起票する。

- work item の起点が必要になった
- 追加要件や spec gap が見つかった
- 複数 artifact にまたがる影響範囲を追跡する必要がある
- self-hosting / editor / viewer / fixture 作成 / implementation 作業中に不足が発見された
- deferred / rejected された判断を履歴として残したい

単純な仕様や、一時的に traceability が不要な仕様にまで requirement を強制しない。
一方で、追加要件が発見された場合は、後回しにして notes や task に散らすのではなく、requirement として早期に捕捉する。

### 10. requirements は implementation-facing specification を所有しない

requirements は implementation-facing specification を所有しない。

複数 module / data structure / tool interface にまたがる wiring route、component boundary、resolver order、cache / index ownership、phase 分担などは `docs/internal-design/` が所有する。
requirements は、それらの internal design を必要に応じて semantic ref で参照してよい。

requirements layer の中に implementation document を置かない。
実装作業中の一時メモや handoff は `docs/impl/` が所有し、長期的に維持すべき内部設計判断は internal design が所有する。

### 11. requirements は fixture / golden の責務を所有しない

requirements は fixture / golden を所有しない。

fixture / golden は、brewprint 処理系や test harness の検証資産である。
fixture の責務境界は ADR-082 が所有する。

requirement は、必要に応じて検証に関係する fixture semantic ref を参照してよい。
ただし、fixture が requirement 本文を所有したり、requirement が expected render output を所有したりしてはならない。

### 12. Design Records MCP の将来拡張として `kind: requirement` を予約する

Design Records MCP または関連 MCP tool で、将来的に `kind: requirement` を扱う余地を予約する。

ただし、本ADRは MVP 実装を変更するものではない。
将来拡張では、以下を query 可能にする。

- requirement を list / get / validate する
- requirement から source spec semantic ref を辿る
- requirement から related work item を辿る
- spec semantic ref から related requirement を辿る
- requirement から fixture / YAML / internal design semantic ref を辿る
- orphan requirement / orphan work item を検出する

MCP の外部 contract は semantic query interface とする。
requirements file の物理構造は内部表現であり、外部 contract にはしない。

## 理由

### なぜ requirements layer が必要か

spec は現行仕様の唯一の正を持つが、追加要件、spec gap、改善要望、延期判断、却下判断の履歴までは所有しない。
これらを会話ログや task file に散らすと、後から変更理由を辿れなくなる。

requirements layer を導入することで、「何が必要か」を長期的に残し、work item の起点を安定させる。

### なぜ requirements を system obligation spec にしないか

requirements が system obligation spec になると、spec と二重管理になる。
また、requirements から YAML / implementation へ直接流れるように見えると、ADR-083 の artifact boundary と衝突する。

現行仕様は spec が所有する。
requirements は要求・不足・要望を所有する。
この責務分離により、spec-first を維持しながら、未整理の要求を失わずに済む。

### なぜ work item と分けるか

requirements に layer 別進捗を持たせると、requirements が task list 化する。
一方で、task file は作業順序・チェックリストであり、長期的な影響範囲 tracking には弱い。

そのため、requirements と tasks の間に work item を置く。
requirements は「何が必要か」、work item は「どの artifact をどこまで更新するか」、tasks は「次に何を実行するか」を所有する。

### なぜ semantic ref を使うか

physical path、Markdown heading、directory layout、file split / merge は変わりうる。
trace の primary key を physical path にすると、rename や split のたびに対応関係が壊れる。

ADR-083 の trace layer common principle に従い、requirements も semantic ref を primary key として関連 artifact を参照する。
これにより、file layout の変更と trace の安定性を分離できる。

### なぜ lazy 起票か

すべての spec に requirement を義務付けると、運用コストが大きくなりすぎる。
requirements は、work item の起点や影響範囲追跡が必要になった箇所から段階的に起票する。

ただし、追加要件そのものは後回しにしない。
要求・不足・要望は早期に捕捉し、採用・延期・却下を後続判断に委ねる。

## 却下した代替案

### 代替案A: requirements を system obligation spec として扱う

- 利点: 実装が満たすべき要求を明確に書ける
- 欠点: spec と二重管理になり、現行仕様の source of truth が曖昧になる

→ 却下。requirements は仕様本文ではなく、要求・不足・要望を所有する。

### 代替案B: requirements から YAML / implementation / fixture へ直接 trace する

- 利点: 要求から成果物を直接辿れる
- 欠点: requirements が source of truth に見える。ADR-083 の YAML as primary DSL source / coverage / work item の責務境界と衝突する

→ 却下。requirements は work item の起点となり、YAML / internal design / spec の対応は coverage が所有する。

### 代替案C: task file だけで追加要件を管理する

- 利点: 既存の作業運用に乗せられる
- 欠点: task は作業手順であり、要求・不足・要望の履歴を長期的に持つ artifact ではない

→ 却下。task は work item を具体作業へ分解する。requirement の代替にはしない。

### 代替案D: impl summary だけで要求の履歴を残す

- 利点: 実装後の状態に近い
- 欠点: 実装前の採用判断・延期判断・影響範囲調査に使えない

→ 却下。impl summary は実装メモであり、requirement 本体は別に持つ。

### 代替案E: 巨大 spec file + section anchor で traceability を維持する

- 利点: file 数が少ない
- 欠点: 人間レビュー可能性が下がる。section anchor の過剰運用で spec が busy になる

→ 却下。spec は人間レビュー可能な粒度へ分割し、traceability は semantic ref / MCP / metadata / index が担う。

### 代替案F: MCP にだけ任せ、文書 artifact を作らない

- 利点: query tool で動的に扱える
- 欠点: MCP tool の入力となる一次情報が存在しない。自然言語本文から推定すると不安定になる

→ 却下。まず `docs/requirements/` を一次情報として持ち、MCP はそれを index / query する。

## 影響

### ADR-083 との関係

本ADRは ADR-083 の artifact boundary を前提とする。
requirements は source of truth ではなく、要求・不足・要望を所有する artifact である。
work item、internal design、coverage、tasks、YAML との責務境界は ADR-083 に従う。

### ADR-082 への影響

fixture / golden は requirements を所有しない。
requirements は fixture semantic ref を参照してよいが、fixture の expected output や test harness 責務は ADR-082 側で扱う。

self-hosting 由来の不足や要望は requirements に捕捉できる。
ただし self-hosting 全体を `docs/requirements/self-hosting/` に一括集約するのではなく、要求・横断進捗・内部設計・coverage・task に分解する。

### docs への影響

- `docs/requirements/` を新設する
- requirement artifact の最小 schema を後続 spec / ADR で定義する必要がある
- `docs/doc-policy.md` に requirements layer、status、lazy 起票、semantic ref 方針を反映する必要がある
- `docs/adr-authoring-guide.md` の責務表に requirements / work-items / coverage / internal design を反映する必要がある

### Design Records MCP への影響

MVP 直後の必須変更ではない。
ただし、将来的に `kind: requirement` の index / query / validate を扱う候補が生まれる。
requirement から source spec semantic ref、work item、fixture、YAML、internal design を query できるようになると、追加要件の impact analysis に使える。

## Evidence

- commit: a80ec7b
- impl commit: tbd
- 参考: ADR-050 spec-first documentation policy, ADR-068 ADR authoring guide, ADR-083 project artifact boundary と YAML as primary implementation source
