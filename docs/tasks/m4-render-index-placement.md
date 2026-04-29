# Milestone 4: render_index / output placement を実装する

- **status**: closed
- **scope**: render_index / output placement
- **source**: migrated from docs/TASKS.md
- **last_updated**: 2026-04-30

---

## Tasks

- [x] **render_index.yaml validation 第1段を実装する**
  - ADR-045に準拠
  - group id命名規則
  - module重複禁止
  - uncovered module の warning
  - nested module の親group所属
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28）

- [x] **renders/ 出力配置 第1段を実装する**
  - ADR-043 / ADR-046に準拠
  - `renders/{group}/...`
  - `_cross/`
  - `_preview/`
  - master index / group index skeleton
