# V01-ADR-096: 既存 artifact の PRODUCT namespace 所有と per-app migration 非実施

- **status**: accepted
- **date**: 2026-06-07
- **depends_on**: V01-ADR-095
- **supersedes**: 
- **migrated_to_spec**: 

## 背景

V01-REQ-PRODUCT-001 により、app namespace と domain namespace を分離した v2 artifact ID 文法の定義が必要になった。v2 文法のターゲット ID は `<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<SEQUENCE>` であり、例として `DRMCP-REQ-DATA-013` のような形式を想定する。

brewprint には現在 `REQ-MCP-*` / `WORK-DATA-*` / `TASK-DATA-*-*` など大量の artifact が存在する。これらは app namespace の概念が存在しなかった時代、brewprint が事実上単一 app として運用されていた時期に起票されたものである。

V01-ADR-095 により DRMCP / BPDSL が別 app namespace として扱われることが決まったが、既存 artifact を各 app namespace に振り分ける migration を行うかどうかを決定する必要がある。

## 決定

既存の全 artifact は `PRODUCT` namespace が所有するものとして扱う。per-app namespace への振り分け migration は実施しない。

新規に起票される artifact は、帰属 app namespace が確定している場合は `<APP_NAMESPACE>-...` 形式を使い、cross-app または帰属が不明な場合は `PRODUCT` namespace を使う。

## 理由

既存 artifact の数は膨大であり、per-app への振り分け migration のコストが現実的でない。

既存 artifact の多くは brewprint がまだ単一 app であった時代の設計判断・要件・作業記録であり、後付けで app namespace を割り当てることで意味が変わるリスクがある。`PRODUCT` は cross-app または repo-wide な関心を表す namespace であり、単一 app 時代の artifact をそのまま `PRODUCT` に帰属させることは意味上も整合する。

ID の安定性は設計記録の信頼性に直結するため、参照の断絶を生む migration は最小化する。

## 却下した代替案

**per-app migration を実施する案**: 各 artifact を `DRMCP` / `BPDSL` 等の app namespace に振り分けることで v2 文法に完全準拠できる。しかし migration 量が膨大であり、振り分け判断が曖昧な artifact も多い。また ID 変更による参照断絶リスクが高い。

## 影響

- 既存 `REQ-MCP-*` / `WORK-DATA-*` / `TASK-DATA-*-*` 等の ID は変更しない
- `PRODUCT` は単一 app 時代の遺産的 namespace として機能する
- V01-REQ-PRODUCT-001 の namespace catalog 定義において、PRODUCT namespace の所有対象を「既存 artifact 全体 + cross-app concerns」として記述する

## Evidence

- V01-REQ-PRODUCT-001: App and domain namespace model（namespace 移行の前提要件）
- V01-ADR-095: YAML DSL と Design Records MCP の結合境界（app namespace 境界の根拠）
- commit: tbd
