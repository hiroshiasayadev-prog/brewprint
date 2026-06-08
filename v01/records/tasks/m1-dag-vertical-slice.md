# Milestone 1: UC-001のDAG 1本を縦に通す

- **status**: closed
- **scope**: DAG renderer / ResolvedProject vertical slice
- **source**: migrated from docs/TASKS.md
- **last_updated**: 2026-04-30

---

## Tasks

- [x] **Go project skeleton を作る**
  - package構成
  - UC-001 fixture 読み込み
  - YAML decode の入口
  - golden test の入口

- [x] **YAML loader + file classifier を実装する**
  - node file / view file / `render_index.yaml` を分類する
  - view file は `as:` で判定する
  - node file は `nodes:` を入口にする

- [x] **Raw YAML structs を実装する**
  - M1範囲として task / model / store / actor / params / returns / reads / writes / initializes を実装
  - view file / render_index.yaml は分類のみで、decodeは後続milestoneに回す
  - validationは薄くてよい

- [x] **Symbol table / QualifiedID parser を実装する**
  - V01-ADR-027 sentinel方式
  - actor global（V01-ADR-031）
  - module nesting
  - 同モジュール内ID直書き
  - クロスモジュールフルパス

- [x] **ResolvedProject build の最小版を実装する**
  - main node by file
  - model / store / task / actor index
  - implicit asset
  - file-private initialized store
  - reads/writes index
  - nodesByQID / nodesByFile
  - DAG renderer が必要な範囲だけ先に実装する

- [x] **DAG renderer vertical slice を通す**
  - UC-001の `auth.task.login` を Raw → Resolved → Render まで通す
  - `docs/uc/001-ec-checkout-flow/renders/auth/dag-login.md` とgolden一致させる
  - params / returns boundary
  - implicit asset
  - store read/write edge label（V01-ADR-044）
  - Tasks詳細セクション
  - `go test ./...` 通過（2026-04-27）
