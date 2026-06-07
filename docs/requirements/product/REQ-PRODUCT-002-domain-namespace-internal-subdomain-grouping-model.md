# REQ-PRODUCT-002: Domain namespace internal subdomain grouping model

- **id**: REQ-PRODUCT-002
- **status**: captured
- **date**: 2026-06-07
- **source_refs**:
  - REQ-PRODUCT-001
- **work_items**:

## Requirement

単一の domain namespace が大量の artifact を抱えるにつれ（例: DRMCP の `MCP` domain が REQ×32 超）、domain 内に概念的なグルーピングが必要になることが示唆されている。artifact ID に subdomain を直接エンコードするかどうかとは独立に、subdomain モデルの定義が必要である。

## Evidence
- DRMCP の `MCP` domain は REQ-MCP-001 〜 REQ-MCP-032 を含み、以下の subdomain が識別できる

  | subdomain | 対象 REQ | 件数 |
  |---|---|---|
  | `TOOLS` | 001, 002, 003, 007 | 4 |
  | `SCHEMA` | 004, 006, 026 | 3 |
  | `AUTHORING` | 005, 008, 009, 010〜012, 014〜025, 027〜032 | 24 |
  | `UI` | 013 | 1 |

  - `AUTHORING` には authoring guidance 取得・管理（REQ-MCP-005, 009）を統合した。guidance は authoring workflow を支援する位置づけであり、spec レベルで分離されていれば subdomain としての統合は許容される
  - `AUTHORING` が全体の約 75% を占めるが、subdomain を細分化する必要はない。spec モジュール内で concern を分離すれば十分

- REQ-PRODUCT-001 が定義する namespace model では app / domain の 2 軸のみが扱われ、subdomain は明示的にスコープ外である
- ID grammar への subdomain 組み込みは破壊的変更になるため、ID とは分離して定義する必要がある

## Required Outcome
- domain namespace 内の subdomain を定義するモデルが存在する
- subdomain は artifact ID に強制エンコードされない形での表現を許容する
  - 具体的には YAML metadata の `subdomain:` フィールド（key-value label 形式）
  - v2 artifact ID grammar（`<APP>-<KIND>-<DOMAIN>-<INDEX>`）にも subdomain セグメントを追加しない
- index はドメインごとにフラットなシーケンスを維持する。subdomain ごとにリセットしない
- subdomain の有効値は事前定義カタログを持たず、既存 records から動的に導出する
  - 「domain 内に存在する `subdomain` フィールドの値の集合」がカタログを構成する
- propose 系ツールが `subdomain` フィールドに新規値を検出した場合、同 domain 内の既存 subdomain 値を列挙して author に提示する（write-time advisory）
  - block はしない。類似値アルゴリズムは使用しない。author が最終判断する
- DRMCP `MCP` domain を例として、subdomain 候補一覧が示される

## Explicitly Excluded Scope
- artifact ID への subdomain エンコーディング（v2 grammar 含め、破壊的変更になるため別途 ADR が必要）
- subdomain ごとの index リセット・サブシーケンス管理
- 全 domain への subdomain 適用の強制
- subdomain 有効値の事前定義カタログおよびそのガバナンスプロセス
- 類似値マッチング・embedding 比較などのアルゴリズム的提案

## Boundary

このREQは subdomain モデルの定義を所有する。namespace catalog や v2 ID grammar の変更は所有しない。
