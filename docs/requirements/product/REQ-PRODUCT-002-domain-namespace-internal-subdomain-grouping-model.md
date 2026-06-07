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

- DRMCP の `MCP` domain は REQ-MCP-001 〜 REQ-MCP-032 を含み、`AUTHORING` / `SCHEMA` / `TOOLS` 等の概念的グルーピングが識別できる
- REQ-PRODUCT-001 が定義する namespace model では app / domain の 2 軸のみが扱われ、subdomain は明示的にスコープ外である
- ID grammar への subdomain 組み込みは破壊的変更になるため、ID とは分離して定義する必要がある

## Required Outcome

- domain namespace 内の subdomain を定義するモデルが存在する
- subdomain は artifact ID に強制エンコードされない形（ラベル・タグ・カタログ注釈等）での表現を許容する
- DRMCP `MCP` domain を例として、subdomain 候補一覧が示される

## Explicitly Excluded Scope

- artifact ID への subdomain エンコーディング（破壊的変更になるため別途 ADR が必要）
- 全 domain への subdomain 適用の強制

## Boundary

このREQは subdomain モデルの定義を所有する。namespace catalog や v2 ID grammar の変更は所有しない。
