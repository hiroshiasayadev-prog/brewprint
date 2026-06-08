# Milestone 5: 残りrendererを順に実装する

- **status**: closed
- **scope**: renderers
- **source**: migrated from docs/TASKS.md
- **last_updated**: 2026-04-30

---

## Tasks

- [x] **State Diagram renderer 第1段を実装する**
  - guard分岐の choice pseudostate（V01-ADR-035）
  - FSMファイル単位の `state-{fsm-id}.md`（V01-ADR-046）
  - UC-001 golden test 3本
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28）

- [x] **Sequence Diagram renderer を実装する**
  - event source別矢印ルール（V01-ADR-036）
  - actionなしtransition（V01-ADR-037）
  - sub task reads/writes traversal（V01-ADR-038）
  - step index prefix / DB片方向message（V01-ADR-041）
  - UC-001 golden test 2本
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28）

- [x] **ER renderer を実装する**
  - 横断view YAML（V01-ADR-039）
  - cross module FK
  - `store.kind: db` から model fields を辿る
  - UC-001 golden test 1本
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28）

- [x] **API Table renderer を実装する**
  - `as: api_table`
  - `http_root_path`
  - endpoint task の leaf path 合成（V01-ADR-028）
  - UC-001 golden test 1本
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28）

- [x] **Wireframe renderer を実装する**
  - HTML fragment
  - `main` container / `layout` object（V01-ADR-042）
  - fixed CSS profile
  - preview harness
  - UC-001 golden test fragment 4本 + preview 1本
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28）
