# V01-TASK-DRMCP-001-02: namespace prefix 導出ルールを spec に記述する

- **id**: V01-TASK-DRMCP-001-02
- **status**: done
- **date**: 2026-06-09
- **work_item**: V01-WORK-DRMCP-001
- **source_requirement**: (TBD)
- **estimate**: 1d
- **depends_on**:
  - V01-TASK-DRMCP-001-01
- **outputs**:
  - drmcp/records/spec/design-records-mcp/overview.md
  - drmcp/records/spec/design-records-mcp/schema.md
  - drmcp/records/spec/design-records-mcp/tools.md

## Goal

drmcp/records/spec/design-records-mcp/ に namespace prefix 導出ルールを記述し、実装の仕様根拠とする。

## Work

- overview.md: "Record scanning と namespace prefix" セクションを追加（導出式 `strings.ToUpper(appNamespaceDir) + "-"`、kind 別 prefix 適用箇所テーブル、multi-root は MVP 外）
- schema.md: ADR H1 例を新フォーマット `# V01-ADR-076:` に更新、"ID normalization model" セクション追加、discovery テーブルを `<records_root>/` に更新
- tools.md: depends_on / response example のパスを v01/records 形式に更新、id_range endpoint を public ID 形式に更新

## Done condition

- namespace_prefix 導出ルールが spec に記述されている
- spec レビューがユーザーに提示されている

## Verification

- drmcp/records/spec/design-records-mcp/overview.md に "Record scanning と namespace prefix" セクションが存在する

## Evidence

- 2026-06-09: overview.md / schema.md / tools.md を更新し namespace prefix 導出ルールと ID normalization model を記述した。spec レビューをユーザーに提示した。
