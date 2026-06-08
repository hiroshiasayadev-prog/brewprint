# V01-ADR-095: YAML DSL と Design Records MCP の結合境界

- **status**: accepted
- **date**: 2026-06-07
- **depends_on**: 
- **supersedes**: 
- **migrated_to_spec**: 

## 背景

V01-REQ-MCP-013（Design Records MCP record browser UI）の実装検討を経て、UI をどの app namespace に置くかという問いが生じた。その過程で、YAML DSL と Design Records MCP がそれぞれ異なる関心を持つ独立したシステムであることが明確になった。

Design Records は要件・決定・調査・作業記録といった設計文書の管理を関心とする。YAML DSL はそれらの設計から生まれる実装の構造定義を関心とする。

両者は現状 brewprint という単一リポジトリに同居しているが、結合の度合いが明示的に定義されていなかった。明示しないまま UI 実装や namespace 分割を進めると、不必要な密結合が生まれる。

## 決定

YAML DSL と Design Records の結合は ID 参照のみとする。依存方向は DSL → Design Records の単方向とする。

- YAML DSL は Design Records の artifact ID を参照できる
- Design Records は YAML DSL の構造を直接参照しない
- impl は DSL を介して Design Records と間接的につながる。impl → Design Records の直接参照は持たない
- tests 等の要件に関わる artifact は Design Records 側に属する

## 理由

両システムが扱うドメインは独立している。Design Records はメタ層（設計情報の管理）、YAML DSL は対象層（実装構造の定義）である。共有すべき情報は ID のみであり、スキーマ・バリデーションロジック・リリースサイクルを共有する必要がない。

密結合にした場合、Design Records の ID スキーマ変更が DSL スキーマ変更・impl 変更に波及し、両システムのリリースサイクルが同期を強制される。

この境界を明確にすることで、V01-REQ-PRODUCT-001 が定義する app namespace において Design Records MCP（DRMCP）と Brewprint DSL（BPDSL）を別 app namespace として扱う根拠となる。

## 却下した代替案

**双方向参照を許容する案**: Design Records が DSL 構造を直接参照するユースケースを想定した場合、スキーマ結合が発生し、独立したリリース・バージョニングができなくなる。具体的なユースケースが確認されていないため却下。

**同一 namespace に置く案**: V01-REQ-MCP-013 の検討で明確になった通り、UI は DRMCP の消費者であり一部ではない。同様に YAML DSL も DRMCP の消費者である。同一 namespace に置くと所有権の境界が曖昧になる。

## 影響

- V01-REQ-PRODUCT-001 の app namespace 列挙において、DRMCP と BPDSL を別 app namespace として扱うことができる
- impl は Design Records artifact ID を直接保持しない。DSL レコードを介した間接参照とする
- YAML DSL の self-hosting 実現後も、この境界は変更を要しない

## Evidence

- V01-REQ-MCP-013: Design Records MCP record browser UI（UI の namespace 検討の起点）
- V01-REQ-PRODUCT-001: App and domain namespace model（app namespace 分離の要件）
- commit: tbd
