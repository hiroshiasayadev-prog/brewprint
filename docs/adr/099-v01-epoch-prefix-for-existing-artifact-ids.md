# 099: 既存 artifact ID への V01 epoch prefix 付与

- **status**: accepted
- **date**: 2026-06-08
- **depends_on**: ADR-097, ADR-098
- **supersedes**: ADR-098（domain token 基準の配布方針）
- **migrated_to_spec**: 

## 背景

ADR-098 は既存 artifact を domain token 基準で app namespace に配布する方針（`REQ-MCP-*` → `DRMCP-REQ-MCP-*` 等）を採用した。しかし以下の問題が判明した。

- domain token が app namespace を一意に示さないケースが存在する（例: `WORK-MCP-004` は MCP domain だが関心の実体は BPDSL の semantic identity / state-machine identity）
- ADR は時系列の全社記録であり、app namespace への配布が構造的に不可能

ADR-098 の mapping table は「機械的に実行できる」という前提が成立しないため、supersede する。

## 決定

既存の全 artifact ID に `V01-` を prefix として付与する。`V01` は v2 grammar 採用以前の epoch を表す合成 app namespace token である。

```
V01-REQ-MCP-001       (旧: REQ-MCP-001)
V01-WORK-MCP-004      (旧: WORK-MCP-004)
V01-ADR-097           (旧: ADR-097)
V01-INV-DATA-001      (旧: INV-DATA-001)
V01-REQ-PRODUCT-003   (旧: REQ-PRODUCT-003)
```

変換はスクリプトで自動化し、以下をアトミックに更新する:

- ファイル名（ID prefix 部分に `V01-` を追加）
- record 本文の `id:` フィールド
- 全ファイル横断の cross-reference（`source_refs` / `work_items` / `depends_on` / `impact_refs` 等）

新規 artifact は v2 grammar（`<APP>-<KIND>-<DOMAIN>-<INDEX>`）を適用し、`V01-` は使用しない。

## 理由

**判断ゼロの機械的変換**: `V01-` prefix はすべての既存 artifact に一律に付与する。domain token と app namespace の対応を判断する必要がない。ADR も例外なく処理できる。

**epoch の明示**: `V01-` は「v2 grammar 採用以前に作成された artifact」という事実を ID 自体に刻む。UI / MCP は `V01-` prefix の有無だけで「旧世代か否か」を判定できる。

**UI / MCP の分岐が単純**: 
- `V01-` prefix あり → 旧世代 artifact。domain token（`MCP` / `DATA` 等）でグルーピング
- `V01-` prefix なし → v2 artifact。app namespace（`DRMCP` / `BPDSL` 等）でグルーピング

旧世代の artifact の実際の app 帰属は MCP の logical mapping（namespace-model spec）で解決すればよく、ID 自体に encode する必要はない。

## ADR への適用

ADR は `ADR-NNN` 形式を維持するという従来の方針（namespace-model spec）を本 ADR が上書きする。`V01-ADR-NNN` 形式に統一することで、他の artifact と同じ変換規則で処理できる。

## 却下した代替案

**domain token 基準の配布（ADR-098）**: `WORK-MCP-004` のように domain token が複数 app にまたがるケースで判断不能になる。機械的変換の前提が成立しない。

**全 PRODUCT prefix（`PRODUCT-REQ-MCP-001` 等）**: `PRODUCT` は cross-app 関心を表す実在の app namespace であり、旧世代 artifact の格納先として使うと `PRODUCT-` prefix が「旧世代」と「真の cross-app concern」の両方を指すことになり曖昧になる。

**既存 ID 変更なし（ADR-096 復活）**: UI / MCP が「prefix なし = 旧世代」と検出する必要があり、実質的に V01 approach と同じ分岐が必要。ただし ID 自体に epoch 情報が含まれないため可読性が劣る。

## 影響

- `docs/` 全 Markdown ファイルの ID 文字列が `V01-` prefix 付きに更新される
- Go コードの ID フォーマット validator を `V01-` prefix に対応させる必要がある（WORK-PRODUCT-003 で対応）
- namespace-model spec の v2 grammar セクション・mapping rule セクションを本決定に合わせて更新する
- ADR-096 の「既存 artifact ID 変更なし」決定は本 ADR が引き継いで上書きする（ADR-098 経由）

## Evidence

- ADR-096: 既存 artifact の PRODUCT namespace 所有と per-app migration 非実施
- ADR-097: app namespace-first repository directory layout の採用
- ADR-098: 既存 artifact ID の v2 grammar への一括移行（本 ADR が supersede）
- INV-PRODUCT-001: migration impact inventory（WORK-MCP-004 等の判断困難ケースの根拠）
- REQ-MCP-013: Design Records MCP record browser UI
