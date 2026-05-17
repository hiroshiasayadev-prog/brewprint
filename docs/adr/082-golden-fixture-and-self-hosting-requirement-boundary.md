# 082: golden fixture と self-hosting requirement の責務境界

- **status**: accepted
- **date**: 2026-05-16
- **depends_on**: ADR-043, ADR-050, ADR-057, ADR-068, ADR-081, ADR-083
- **supersedes**:
- **migrated_to_spec**:

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

brewprint の `docs/uc/` は、当初「実例 YAML と render 結果を置く場所」として運用されてきた。
UC-001 は EC checkout flow の YAML と render 出力を持ち、renderer / validation の fixture として使われる。

一方、UC-002 self-hosting は、brewprint 自身を brewprint YAML で表現する試みとして作られた。
UC-002 には、MCP公開contractのblueprint化、内部レイヤーのblueprint化、spec gap発見ログ、editor / viewer 要件メモ、作業計画が含まれている。

しかし、この内容は「この YAML からこの render が生成されることを検証する golden fixture」の責務を超えている。
self-hosting は、brewprint 本体という project が満たすべき要求、内部設計、coverage、work item、task をまたぐ活動である。

ADR-081 により `docs/requirements/` は要求・不足・要望を捕捉する layer として定義される。
ADR-083 により、requirements / work-items / tasks / internal design / coverage / YAML の責務境界が整理された。
この前提では、self-hosting を `docs/requirements/self-hosting/` 一箇所に集約するのではなく、内容ごとに適切な artifact へ分解する必要がある。

本ADRは、`docs/uc/` を golden fixture corpus として再定義し、self-hosting 的内容を fixture から切り離す境界を定める。

## 決定

### 1. UC / fixture の責務を golden fixture に限定する

`docs/uc/` は、brewprint 仕様・実装を検証する golden fixture corpus として扱う。

UC / fixture の主目的は以下である。

```text
fixture directory の yaml/ を入力として render / validate し、期待値と一致することを検証する
```

UC / fixture が所有するものは以下に限定する。

- 入力 YAML
- 期待 render 出力
- 期待 diagnostics（必要な場合）
- fixture の意図を説明する短い README
- fixture-local coverage / note（どの観点を検証するかの局所説明）

UC / fixture は、requirements、長期的な spec gap backlog、editor / viewer 要件、internal design、work item、実装作業計画を所有しない。

fixture / golden は、brewprint 処理系や test harness の検証資産であり、一般 brewprint project の source-of-truth chain には含めない。

### 2. 期待 render 出力ディレクトリは新規 fixture では `render_expected/` とする

golden fixture として commit される render 出力は、生成物そのものではなく expected output である。

そのため、新規 fixture では期待 render 出力ディレクトリ名を `render_expected/` とする。

```text
docs/uc/NNN-example/
  yaml/
  render_expected/
```

render test harness は、`yaml/` を入力に一時ディレクトリへ actual render を生成し、expected output と比較する。

```text
brewprint render --yaml-root docs/uc/NNN-example/yaml --out <tmp>/render_actual --clean
compare <tmp>/render_actual with docs/uc/NNN-example/render_expected
```

`render_actual/` は repository に commit しない。

比較セマンティクス、すなわち byte-exact / line-exact / normalized whitespace / HTML DOM 比較などの詳細は test harness spec または internal design の責務とし、本ADRでは定義しない。
本ADRが決めるのは、fixture に commit する expected output の directory 命名と責務境界である。

### 3. UC-001 の既存 `renders/` は維持する

UC-001 は ADR-057 により v1.0.0-spec の canonical fixture として固定されている。
そのため、UC-001 の既存 `renders/` は v1.0.0-spec 凍結資産として維持する。

本ADRは、UC-001 の v1.0.0-spec snapshot を retroactive に rename しない。

`render_expected/` は、UC-002 整理後の新規 fixture または新しい spec snapshot 以降の fixture から適用する。
既存 fixture の `renders/` から `render_expected/` への移行可否・タイミングは、後続 migration task で判断する。

### 4. ADR-043 / project-layout の fixture 運用を refine する

ADR-043 および `docs/spec/project-layout.md` は、通常 project の render output として `renders/` を定義している。

本ADRはそれを全面的に supersede しない。
通常の user project / generated output directory と、golden fixture の expected output directory を区別する refinement を加える。

| context | directory |
|---|---|
| 通常 project の renderer output | `renders/` |
| golden fixture の expected output | `render_expected/` |
| test harness の actual output | temporary directory, not committed |

この分岐は、accepted 後に `docs/spec/project-layout.md` へ反映する。

### 5. self-hosting は fixture から分離し、artifact ごとに分解する

brewprint self-hosting は UC / golden fixture ではなく、brewprint 本体 project を brewprint で記述しながら、不足・内部設計・coverage・作業計画を発見する活動である。

self-hosting の主な内容は以下である。

- brewprint 自身の MCP contract を blueprint 化できること
- internal layer を blueprint 化できること
- spec gap を検出し、必要な言語表現力へ反映できること
- requirement / work item / coverage / internal design の traceability を実用できること
- fixture として検証すべき YAML / expected render を切り出せること

これらは render expected output の比較だけでは閉じない。
そのため、self-hosting の内容は `docs/requirements/self-hosting/` 一箇所へ集約しない。
ADR-083 の artifact placement decision rule に従い、以下のように分解する。

| 内容 | 移動先 / 扱い |
|---|---|
| 要求・不足・要望・spec gap 候補 | `docs/requirements/` |
| 複数 artifact にまたがる進捗・影響範囲 | `docs/work-items/` |
| 複数 module / data structure / tool interface にまたがる wiring route | `docs/internal-design/` |
| spec / internal design / YAML の対応関係 | `docs/coverage/` |
| 具体的な移行手順・編集順序・チェックリスト | `docs/tasks/` |
| 実装中の一時メモ・handoff | `docs/impl/` |
| 検証用の入力 YAML / expected render / expected diagnostics | `docs/uc/` |

UC として残す場合は、self-hosting の一部を検証するための限定 fixture として再構成する。

### 6. editor / viewer notes は self-hosting fixture とは分離する

既存 UC-002 の `editor-viewer-notes.md` は、self-hosting 作業中に発見された editor / viewer 要件を蓄積している。

ただし、editor / viewer は golden fixture の責務ではない。
また、brewprint 本体 language / renderer / MCP implementation の requirement と同一とも限らない。
将来の editor / viewer product または周辺 tooling の requirements として扱う方が自然である。

そのため、editor / viewer notes は UC fixture に残さない。
accepted 後の移行では、以下のいずれかに分離する。

- `docs/requirements/editor/`
- `docs/requirements/viewer/`
- editor / viewer 専用の notes / backlog 文書
- 後続 product / tooling project 側の requirements

具体的な移動先は、editor / viewer をどの project / product として扱うかを決める後続 task または ADR で確定する。

### 7. UC-002 は self-hosting workspace として再分類し、fixture 部分だけを残す

既存 `docs/uc/002-brewprint-self-hosting/` は、起票時点では UC directory に存在するが、実質的には self-hosting workspace である。

本ADR accepted 後、UC-002 の内容は以下の方針で整理する。

| 現在の内容 | 移動先 / 扱い |
|---|---|
| self-hosting 概要 | requirements / work-items / internal design への入口として再配置 |
| MCP公開contract blueprint 方針 | spec / internal design / YAML / coverage に分解 |
| internal layer blueprint 方針 | internal design / YAML / coverage に分解 |
| editor / viewer notes | self-hosting fixture とは別の editor / viewer requirement または notes |
| spec gap backlog | requirements に捕捉し、必要なら work item 化 |
| 作業計画・移行 checklist | tasks に移す |
| 入力 YAML / expected render / expected diagnostics | golden fixture として必要な範囲だけ `docs/uc/` に残す |

具体的なファイル移動、rename、既存リンク更新、fixture再生成は task file で扱う。
本ADRは境界を決めるものであり、移行作業 checklist は所有しない。

### 8. fixture coverage は fixture-local coverage として扱う

fixture は requirement 本文や system obligation を所有しない。
ただし、fixture が何を検証するかを説明する fixture-local coverage / note は持てる。

既存の `docs/uc/**/docs/coverage.md` は fixture-local coverage として扱う。
これは ADR-083 の `docs/coverage/` が定義する project-level trace coverage とは別概念である。

fixture-local coverage は、必要に応じて requirement ID や semantic ref を参照してよい。
ただし、project-level の spec / internal design / YAML 対応関係は `docs/coverage/` が所有する。

### 9. `docs/uc/` の名称変更は本ADRでは扱わない

`docs/uc/` という名前は「ユースケース」を連想させるため、golden fixture corpus としては曖昧である。

しかし、ディレクトリ rename は既存リンク・render placement・test harness・doc参照への影響が大きいため、本ADRでは扱わない。

本ADRで決めるのは、`docs/uc/` の責務を golden fixture corpus として再定義することである。
rename の要否と移行先名称は後続ADRまたは migration task で扱う。

## 理由

### なぜ UC を golden fixture に限定するか

UC / fixture は、brewprint の仕様・実装が期待どおり動くことを検証する入力と期待出力である。

ここに requirements、spec gap backlog、internal design、work item、task planning を置くと、fixture が検証資産なのか要求定義なのかが曖昧になる。
fixture は長期的な設計判断や要求の一次情報を所有せず、検証資産として単純に保つ方がよい。

### なぜ `render_expected/` か

`renders/` は renderer の生成物全般に見える。
しかし fixture に commit される render は、test harness が比較する expected output である。

`render_expected/` と命名することで、実行時生成される actual output と区別できる。
これにより、手編集禁止・比較対象・golden update の意味が明確になる。

### なぜ UC-001 は維持するか

UC-001 は v1.0.0-spec の canonical fixture として凍結済みである。
既存 snapshot を rename すると、過去の spec snapshot の意味が変わってしまう。

そのため、新しい expected output 命名は新規 fixture から適用し、UC-001 は既存構造を維持する。

### なぜ self-hosting を一箇所へ移さないか

self-hosting は、単なる render golden fixture ではない。
brewprint 本体を対象に、MCP contract、internal layer、spec gap、requirement traceability、coverage、implementation workflow を検証する活動である。

これらを `docs/requirements/self-hosting/` 一箇所に集めると、requirements が内部設計・進捗・coverage・task を抱えてしまう。
ADR-083 の責務境界に従い、self-hosting の内容は artifact ごとに分解する。

### なぜ editor / viewer notes を分離するか

editor / viewer は brewprint 本体の language / renderer / MCP 実装と密接に関係するが、同一 project requirement とは限らない。
将来の editor product / viewer product / tooling requirements として独立に扱う方が、self-hosting fixture の責務を濁らせない。

### なぜ UC-002 をすぐ削除しないか

既存の UC-002 には、起票時点の evidence、作業ログ、MCP contract YAML、spec gap 発見ログが含まれている。
これらを一度に削除・移動すると情報損失やリンク切れが起きやすい。

本ADRは責務境界を決め、具体的な migration は task file で段階的に行う。

## 却下した代替案

### 代替案A: UC-002 をそのまま UC として維持する

- 利点: 既存構成を変えずに済む
- 欠点: UC が requirement discovery / task planning / spec gap backlog / internal design を抱え続け、golden fixture としての責務が曖昧になる

→ 却下。self-hosting は fixture から分離し、artifact ごとに分解する。

### 代替案B: UC を requirement と golden fixture の両方として扱う

- 利点: 文書数を増やさずに済む
- 欠点: requirement と test fixture の一次情報が混ざり、変更時の責務が分からなくなる

→ 却下。requirements と fixture は分ける。

### 代替案C: self-hosting を `docs/requirements/self-hosting/` に一括集約する

- 利点: UC-002 からの移動先が分かりやすい
- 欠点: requirements が内部設計・coverage・work item・task を抱え、ADR-083 の artifact boundary と衝突する

→ 却下。self-hosting の内容は requirements / work-items / internal design / coverage / tasks / fixtures に分解する。

### 代替案D: すべての既存 fixture を即座に `render_expected/` へ rename する

- 利点: 命名が一貫する
- 欠点: UC-001 の v1.0.0-spec 凍結資産を retroactive に変えてしまう

→ 却下。UC-001 は維持し、新規 fixture から `render_expected/` を使う。

### 代替案E: `renders/` のまま新規 fixture も運用する

- 利点: 既存ディレクトリ名を維持できる
- 欠点: expected output と actual generated output の区別が弱い

→ 却下。新規 fixture では `render_expected/` を使う。

### 代替案F: ただちに `docs/uc/` を rename する

- 利点: 名前と責務が一致する
- 欠点: 既存リンク・render placement・test harness・doc参照への影響が大きい

→ 本ADRでは採用しない。責務再定義を先に行い、rename は後続で検討する。

## 影響

### ADR-081 への影響

requirements は fixture や self-hosting workspace を所有しない。
self-hosting 由来の要求・不足・要望・spec gap 候補は requirements に捕捉するが、進捗は work item、内部設計は internal design、対応関係は coverage、作業手順は tasks が所有する。

### ADR-083 との関係

本ADRは ADR-083 の artifact boundary を前提とする。
fixture / golden は一般 project の source-of-truth chain には含めない。
ただし work item は fixture / golden impact を追跡してよい。

### docs への影響

- `docs/doc-policy.md` の UC 運用を golden fixture corpus として更新する必要がある
- `docs/adr-authoring-guide.md` の責務表から UC 所有の gap 発見ログ / migration 状態を外し、requirements / work-items / tasks / internal design / coverage に振り直す必要がある
- UC-002 の requirement 的内容を artifact ごとに分解する migration task が必要になる
- editor / viewer notes の行き先を self-hosting fixture とは別に決める必要がある

### project layout / render spec への影響

`docs/spec/project-layout.md` は、通常 project の `renders/` と fixture の `render_expected/` を区別する方針を反映する必要がある。

ADR-043 / project-layout の `renders/` 方針は通常 project の render output として維持し、fixture expected output は refinement として追加する。

### task への影響

移行作業は task file に分けて管理する。
少なくとも以下が必要になる。

- UC-002 から requirements / work-items / internal design / coverage / tasks への内容分解
- fixture として残す YAML / expected render / expected diagnostics の選別
- 新規 fixture の `render_expected/` 運用定義
- 既存リンク・coverage・test harness の更新

### M15 / v1.1 への影響

M15 v1.1 では UC-002 が大きく書き換わる見込みであるため、本ADRの責務整理は同時に進めやすい。
ただし、本ADR自体は M15 の language feature 実装をブロックしない。

### Design Records MCP / requirements traceability への影響

ADR-081 の requirement traceability と ADR-083 の semantic ref 方針と連携し、fixture-local coverage と project-level coverage を区別する必要がある。
将来 MCP で requirement / fixture coverage を index する場合、本ADRの境界が前提になる。

## Evidence

- commit: a80ec7b
- impl commit: tbd
- 参考: ADR-043 project root layout and render output, ADR-050 spec-first documentation policy, ADR-057 brewprint v1 snapshot, ADR-068 ADR authoring guide, ADR-081 project requirements layer と semantic traceability, ADR-083 project artifact boundary と YAML as primary implementation source
