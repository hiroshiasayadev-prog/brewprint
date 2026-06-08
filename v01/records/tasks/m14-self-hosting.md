# Milestone 14: brewprint self-hosting

- **status**: paused
- **scope**: UC-002 / spec / docs
- **source**: V01-ADR-057でv1.0.0-spec凍結後の次段階。`docs/TASKS.md` 検討中セクションから昇格
- **last_updated**: 2026-05-31

---

## Context

V01-ADR-057でbrewprint v1.0.0-spec が凍結された（gitタグ `v1.0.0-spec` 発行済み）。
次の段階として、brewprint自身のblueprintをbrewprintで設計するself-hostingに着手する。

このmilestoneは2つの目的を兼ねる:

1. **brewprint v1の実用検証** — v1.0.0-spec の表現力・MCP coverageが、brewprint自身という実用規模のソフトウェアを表現するに足るかを実証する
2. **editor/viewer要件抽出** — brewprintをblueprint化する作業中に浮かび上がるeditor/viewerの要件を、UC-002配下にメモとして蓄積する

UC-002は v1.0.0-spec の追加canonical fixtureとして機能させる。
v1.0.0-spec は凍結済みであり、self-hostingで露呈した仕様gapは原則v2向け改訂候補として記録する。

---

## v1.0.0-spec を最初の凍結基準とする

検証は v1.0.0-spec のspecだけで brewprint をどこまで表現できるかを基準に行う。
v1範囲のspec仕様で表現できない箇所はspec gapとして `TASKS-UC-002.md` に記録し、V01-ADR-057で確定したRelease snapshots運用に従って次の `v{N}.0.0-spec` に向けた改訂候補として扱う。

---

## Non-goals

M14では以下を実装しない / 着手しない。

- brewprint v2.0.0-spec の凍結
- editor / viewer の実装
- editor / viewer の `docs/spec/editor-requirements.md` 等への正式仕様化（spec昇格はM14後）
- UC-001（EC Checkout Flow）の破壊的変更
- v1.0.0-spec のspec / ADR / Go実装ツリーの遡及修正（V01-ADR-050 / V01-ADR-057に準ずる）
- V01-ADR-010の複数論点混在分割（v1後検討事項として継続。V01-ADR-057 §6）
- DISCLAIMER.mdの起草（M14と独立。ユーザー起草）

self-hosting途中でv1範囲内のspec誤記レベルの修正が必要になった場合は、V01-ADR-050 §7の漸進移行ルールに従って個別判断する。

---

## 検証アプローチ: MCP公開contract優先 → レイヤー単位

self-hostingの作業順序は以下の2段構えで進める。

### Phase A: MCP公開contract をblueprint化する

brewprintの対外公開仕様の核であるMCP toolsから先にblueprint化する。
V01-ADR-054 / V01-ADR-055 / V01-ADR-056で定義され、`docs/spec/mcp/**` で仕様化されているtool群を、brewprint YAMLで表現する。

これによりMCP coverageが実用に耐えるかを最優先で検証する。

### Phase B: 内部レイヤーをblueprint化する

Phase A完了後、内部レイヤーをblueprint化する。
レイヤー境界はV01-ADR-047で確定しており、現行Go実装と対応する。

レイヤー順は以下を想定する（self-hosting途中で順序を見直してよい）:

1. source layer
2. rawyaml layer
3. semantic layer
4. resolve layer
5. query layer
6. render layer
7. CLI（`brewprint render` / `brewprint validate`）

Phase Bの過程で全体俯瞰のER（model間の依存関係）が副次的に固まることを期待する。

---

## spec gap発見時の処理ルール

self-hostingで露呈した仕様gapは以下のルールで処理する。

- **追跡場所**: `docs/uc/002-brewprint-self-hosting/TASKS-UC-002.md` に発見順に記録する
- **累積整理**: gapが一定量蓄積した時点で `docs/spec/v2-gap-list.md`（仮）等の累積ファイル化を検討する。M14開始時点では作らない
- **ADR起票トリガー**:
  - v1範囲内で即修正可能なspec誤記レベル: V01-ADR-050 §7の漸進移行ルールに従い個別ADR起票またはspec修正
  - v2に向けた構造的変更: 専用ADRを起票し、status `proposed` でacceptedまで議論する
- **v1.0.0-spec の凍結を破壊しない**: 凍結対象（ADR / spec / UC-001 / Go実装ツリー）への変更は、M14の都合で安易に行わない

---

## editor / viewer 要件の蓄積場所

self-hosting中に浮かび上がるeditor / viewer要件は以下に蓄積する。

- **蓄積場所**: `docs/uc/002-brewprint-self-hosting/editor-viewer-notes.md`（UC-002起票時に新設）
- **粒度**: 探索物としてのメモレベルで可（仕様化前の段階）
- **spec昇格**: 要件が一定量蓄積し、editor/viewerが実装対象として認識された段階で `docs/spec/editor-requirements.md` 等に整理して移管する。M14内では昇格しない
- **大きな方針決定**: editor/viewerの設計判断（別バイナリ / MCP越し / AST直アクセス等）が出た場合は専用ADRを起票する

---

## Tasks

### UC-002骨格整備

- [x] **UC-002 ディレクトリ骨格を作成する**
  - `docs/uc/002-brewprint-self-hosting/` を新設する
  - `docs/uc/002-brewprint-self-hosting/yaml/` を新設する
  - `docs/uc/002-brewprint-self-hosting/renders/` を新設する（renderer出力先。手書きしない）

- [x] **UC-002 README.md を起票する**
  - doc-policy.md §uc運用のフォーマットに従う
  - 概要 / ファイル構成 / TODO/spec gap セクションを含む
  - UC-001のREADME.mdを参考形式とする
  - render fixture再生成コマンドを記載する

- [x] **UC-002 render_index.yaml を起票する**
  - Phase A / Phase B の進行に合わせて段階的にgroupを追加する想定
  - 初期はMCPレイヤーのgroupだけ定義する
  - V01-ADR-045 / V01-ADR-046に従う

- [x] **TASKS-UC-002.md を起票する**
  - spec gap発見ログの追跡に使う
  - 形式はUC-001のTASKS-UC-001.mdに揃える
  - Phase A の YAML が入った後に `docs/coverage.md` を起票する方針を記録する

- [x] **editor-viewer-notes.md を起票する**
  - 起票時は空のスケルトン + このM14への参照リンクのみで可
  - self-hosting作業中に追記していく

### Phase A: MCP公開contract のblueprint化

- [x] **MCP toolsの blueprint表現方針を決める**
  - 各MCPツール（`get_references` / `get_reference_tree` / `inspect` / `analyze_impact` 等）は `task` として表現する
  - MCP toolはHTTP endpointではないため、tool taskに `endpoint: true` は付けない
  - request / response型は `model`、MCP server / client は `actor`、ResolvedProjectは `store kind: context` として表現する
  - 方針は `docs/uc/002-brewprint-self-hosting/docs/phase-a-mcp-contract.md` に記録済み
  - 並列作業分割は `docs/uc/002-brewprint-self-hosting/docs/phase-a-work-split.md` に記録済み
  - 現時点では新規ADR不要と判断

- [x] **MCP tools の actor / model 定義をblueprint化する**
  - MCPツールの request / response型をbrewprint modelで表現済み
  - MCP server / client の actor を `yaml/actors.yaml` に定義済み
  - ResolvedProject context store を `yaml/mcp/store/resolved_project_store.yaml` に定義済み
  - `docs/spec/mcp/schema.md` / `errors.md` / `tools/*.md` を参照し、optional / enum / union / discriminated object等は `note` で暫定表現
  - v1 model表現力gapは `docs/uc/002-brewprint-self-hosting/TASKS-UC-002.md` に記録済み

- [x] **MCP tools の task / flow をblueprint化する**
  - 8 MCP tool分のtask YAMLを `yaml/mcp/task/*.yaml` に配置済み
  - 各ツール呼出しのflowは `validate_request -> query_service -> build_response` としてDAGで表現済み
  - QueryServiceとの境界は file-local `query_service` sub task + `reads: [resolved_project_store]` + `note` で表現済み
  - QueryServiceを独立node種別として捏造しない方針を維持

- [ ] **Phase A render を生成・確認する**
  - Phase A YAMLは配置済み
  - `docs/uc/002-brewprint-self-hosting/docs/coverage.md` は現状に追随済み
  - `brewprint render` / `go test ./...` は未実行
  - render結果に対するgolden test相当の確認は未完了
  - render出力上の表現力gapがあれば spec gap として記録する

### Phase B: 内部レイヤーのblueprint化

- [ ] **source layer をblueprint化する**
- [ ] **rawyaml layer をblueprint化する**
- [ ] **semantic layer をblueprint化する**
- [ ] **resolve layer をblueprint化する**
- [ ] **query layer をblueprint化する**
- [ ] **render layer をblueprint化する**
- [ ] **CLI（`brewprint render` / `brewprint validate`）をblueprint化する**

各レイヤーtaskの詳細サブtaskは、当該レイヤー着手時に本ファイルへ追記する形で展開する。
レイヤー単位で完了した時点でcommitを切る運用とする。

### 全体俯瞰

- [ ] **UC-002 全体ER をまとめる**
  - Phase B完了後、UC-002全体のER（model間依存）を確認する
  - v1範囲のER renderer表現力で足りるかを検証する

- [ ] **MCP coverage 実用検証レポートをまとめる**
  - Phase A / Phase B を通じてMCP toolsがself-hosting作業に十分だったかを `docs/uc/002-brewprint-self-hosting/` 配下にまとめる
  - 不足したMCP機能を `TASKS-UC-002.md` に記録する

### M14クローズ作業

- [ ] **TASKS-UC-002.md のspec gap一覧をレビューする**
  - v1範囲内修正 / v2向け構造変更 / 棄却 のいずれかにラベル付けする
  - v1範囲内修正は適宜個別ADR起票またはspec修正として処理する
  - v2向け構造変更は次の `v{N}.0.0-spec` 候補として残す

- [ ] **editor-viewer-notes.md のspec昇格要否を判断する**
  - 要件が一定量蓄積した場合、`docs/spec/editor-requirements.md` への昇格をM15以降のtaskとして提案する
  - M14内では昇格しない

- [ ] **M14完了時に docs を更新する**
  - 本ファイル `docs/tasks/m14-self-hosting.md` を closed にする
  - `docs/TASKS.md` のM14 status を closed にする
  - 必要に応じて `docs/impl/uc-002-self-hosting-summary.md` 等の引継ぎメモを起票する
