# 057: brewprint v1 snapshot

- **status**: proposed
- **date**: 2026-05-02

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

M0〜M13まですべてclosedとなり、brewprintのコア機構が一通り揃った。

- 全レイヤー実装: source / rawyaml / semantic / resolve / render / query / mcp
- 主要MCPツール: project exploration / diagram element query / impact traversal / `analyze_impact`（hybrid v1 close）
- spec構成の安定: ADR-050のspec-first方針下で `overview` / `project-layout` / `file-types` / `naming` / `nodes` / `edges` / `diagnostics` / `mcp` / `views/*` が揃う
- UC-001（EC Checkout Flow）が一貫して通るgolden fixtureとして機能
- `go test ./...` パス、`gofmt` 整形済み

次の段階として brewprint self-hosting（brewprint自身のblueprintをbrewprintで設計）に着手するにあたり、その手前で「v1の凍結点」を明示する必要がある。
self-hostingで露呈する仕様gapを「v1からの差分」として追跡可能にするためである。

doc-policy.md §11で未解決とされていた「Release snapshots運用」も本ADRで解決する。

## 決定

### 1. brewprintにおける「version 1」の3レベル分離

brewprintは以下の3レベルでversionを使い分ける。

| レベル | 名前 | 意味 |
|---|---|---|
| 公開contract | **MCP v1** | MCPツール公開仕様。`docs/spec/mcp/versioning.md` で運用中 |
| 仕様+実装スナップショット | **brewprint v1.0.0-spec** | M0〜M13完了 + spec整備状態。本ADRで凍結 |
| 製品全体 | **brewprint v1** | 上記2つを内包する全体ブランド |

本ADRが定義するのは2番目（`v1.0.0-spec`）。
MCP v1は独立した公開contract軸であり、本ADRと独立して進化しうる。

### 2. brewprint v1.0.0-spec の判定基準

以下をすべて満たす状態を `v1.0.0-spec` とする。

1. **milestone**: M0〜M13がすべてclosed
2. **spec**: 主要specが揃っている
   - `docs/spec/overview.md`
   - `docs/spec/project-layout.md`
   - `docs/spec/file-types.md`
   - `docs/spec/naming.md`
   - `docs/spec/nodes.md`
   - `docs/spec/edges.md`
   - `docs/spec/diagnostics.md`
   - `docs/spec/mcp/*.md`
   - `docs/spec/views/*.md`
3. **MCP**: ADR-054 / ADR-055 / ADR-056 が定めるMCPツール群がhybrid v1として実装済み
4. **UC**: UC-001（EC Checkout Flow）がgolden fixtureとして通る
5. **品質**: `go test ./...` パス、`gofmt` 整形済み

本ADR起票時点で上記すべてを満たしているため、v1.0.0-spec成立条件は満たされている。

### 3. v1の参照UC

UC-001（EC Checkout Flow）を **v1の参照UC** として固定する。

- `docs/uc/001-ec-checkout-flow/` は v1.0.0-spec のcanonical fixtureとして扱う
- v1範囲のspec仕様は、UC-001で表現可能であることが暗黙の検証基準である
- self-hosting以降、新規UCを追加する場合もUC-001を破壊しない方向で進める

### 4. Release snapshots運用

doc-policy.md §11で未解決だった「Release snapshots運用」を以下のとおり定める。

#### gitタグ命名

- 仕様+実装スナップショット: `v{MAJOR}.{MINOR}.{PATCH}-spec`
  - 例: `v1.0.0-spec`
- 公開contractのバージョンは別軸（MCP v1など）で扱う。本タグ規則は仕様+実装スナップショット専用

#### 凍結対象

タグを切る時点で、以下が同時にスナップショット対象となる。

- `docs/adr/*` — 起票済みADR一式
- `docs/spec/**` — 現行仕様
- `docs/uc/**` — UC fixture（特にUC-001のYAML + renders）
- Go実装ツリー全体（ビルド可能・テスト通過状態）

#### 運用頻度

タグは **メジャーな仕様マイルストーン** に合わせて切る。毎リリース・毎milestoneでは切らない。
v1.0.0-spec が最初のスナップショット。次のタグはself-hostingを経たspec改訂が一定量蓄積してから検討する。

#### タグ実行

ADR-057がacceptedとなり、commitが完了した時点でユーザーが以下を実行する。

```bash
git tag -a v1.0.0-spec -m "brewprint v1.0.0-spec snapshot"
git push origin v1.0.0-spec
```

### 5. DISCLAIMER.mdの新設方針

brewprintは個人プロジェクトでありながら将来プロダクトになりうる。
法務的な主張（業務時間外開発、会社リソース不使用、公知技術の組合せ）を明示する文書を別ファイル化する。

- ファイル: `DISCLAIMER.md`（プロジェクトルート）
- README.mdからリンクで参照する
- 文面はユーザーが起草する。本ADRでは方針のみ定める

本ADRがacceptedとなった時点で `DISCLAIMER.md` 起草はopen taskとなる。
v1.0.0-spec タグを切る前に最低限の文面を入れることを推奨するが、ブロッカーとはしない。

### 6. v1後検討事項の引き継ぎ

doc-policy.md §11の未解決事項のうち、本ADRで解決していないものは以下:

- ADR-010の複数論点混在（CA強制 / ディレクトリ構造 / model-asset分離）の分割要否
  - v1後に検討する。漸進移行ルール（ADR-050 §7）に従い、ADR-010に触れたタイミングでspec移管を進める

doc-policy.md §11は本ADR反映と同commitで更新する。

## 理由

**v1スナップショットを切る理由**

- self-hosting着手前に凍結点を明示することで、self-hostingが露呈する仕様gapを「v1からの差分」として追跡できる
- brewprintは別会話のClaudeも含めた共同作業前提のプロジェクト。「v1完成」という共通認識をdoc化することで、認識ずれを防ぐ
- doc-policy.md §11で未解決だった「Release snapshots運用」の最初の実例として、節目凍結の運用パターンを確立する

**3レベル分離の理由**

- `docs/spec/mcp/versioning.md` がすでに「MCP v1」の運用を始めている
- brewprint全体の「v1」と区別しないと、版数の意味が混乱する
- 公開contract（MCP）と内部仕様+実装は独立して進化しうるので、軸を分ける

**`-spec` サフィックスの理由**

- 純粋な `v1.0.0` は将来「製品v1」のために残す
- `-spec` は仕様+実装スナップショットを表し、製品リリース版とは別軸であることを示す
- セマンティックバージョニングと両立可能

**却下した代替案**

- 純粋な `v1.0.0` タグ — 将来の製品リリースとの版数衝突を招く
- spec / 実装で別タグ — brewprintはspec-first方針下で両者が密結合。分離する実益がない
- ADR-057でDISCLAIMER文面まで起草 — 法務的主張はユーザー固有の事情を含むため、Claudeが起案するのは不適切

## 影響

### 既存docへの影響

- `docs/doc-policy.md` §11の更新（本ADRと同commitで実施）
  - 「Release snapshots運用」をv1.0.0-spec運用として確定
  - 「DISCLAIMER的文書の要否」を「DISCLAIMER.md新設方針確定（文面起草はopen）」に更新
  - 「ADR-010の複数論点混在」はv1後検討として残す

### TASKS.mdへの影響

- 「検討中」セクションの brewprint self-hosting を milestone化（M14）する準備に入れる
- `DISCLAIMER.md` 起草をopen taskとして追加

### 後続作業への影響

- self-hosting milestone（M14）はv1.0.0-spec タグ確定後に着手するのが望ましい
- v1後の仕様改訂はsupersedingやspec修正で進め、いずれ次の `v{N}.0.0-spec` タグで凍結する

## Evidence
- commit: tbd（本ADR起票時のcommit hash）
- impl commit: 該当なし（ドキュメント運用ADR）
- 参考: セマンティックバージョニング、Goプロジェクトのtag運用慣習
