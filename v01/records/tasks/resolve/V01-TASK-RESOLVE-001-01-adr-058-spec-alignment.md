# V01-TASK-RESOLVE-001-01: V01-ADR-058 の file-private sub node scope を spec に反映する

- **id**: V01-TASK-RESOLVE-001-01
- **status**: done
- **date**: 2026-05-31
- **work_item**: V01-WORK-RESOLVE-001
- **source_requirement**: V01-REQ-RESOLVE-001
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - `docs/spec/nodes.md` の sub node file-private scope 明確化
  - `docs/spec/naming.md` の public QualifiedID / file-private local ID 境界明確化
  - `docs/spec/diagnostics.md` の duplicate_node 対象明確化

## Goal

V01-ADR-058 の accepted 判断を現行 spec に反映し、実装者が file-private sub node を project-wide public QualifiedID uniqueness の対象と誤読しない状態にする。

## Work

- `V01-ADR-058` の以下の判断を spec に反映する。
  - sub node は project-wide QualifiedID 一意性制約の対象外
  - sub node ID は同一 file 内で一意
  - 別 file にある同名 sub node local ID は衝突しない
  - sub node は external YAML / public MCP selector 向けの public QualifiedID を持たない
  - flow / reads / writes 等の bare ID は同一 file の sub node を優先して解決する
- `docs/spec/nodes.md` の file structure / common field 周辺を更新する。
- `docs/spec/naming.md` の main node / sub node identity 境界を更新する。
- `docs/spec/diagnostics.md` の `duplicate_node` 説明を main/public node collision と file-local duplicate の扱いに整理する。
- V01-ADR-078 の MCP synthetic ID 方針と混同しないよう、public QualifiedID と MCP query layer synthetic ID の違いを必要最小限で明示する。

## Done condition

- V01-ADR-058 §4 の spec update intent が、対象 spec に反映されている。
- public QualifiedID、file-private local ID、MCP synthetic ID の責務が混同されていない。
- 後続 task が spec gap を理由に resolver 修正を保留しなくてよい状態になっている。

## Verification

- `docs/spec/nodes.md` / `docs/spec/naming.md` / `docs/spec/diagnostics.md` を読み戻す。
- V01-ADR-058 と矛盾しないことを確認する。
- V01-ADR-078 の `<semantic-anchor-id>#<local-id>` 方針を、public QualifiedID として誤記していないことを確認する。

## Evidence

- `docs/spec/nodes.md` に、sub node が file-private local ID であり、同一 file 内でのみ一意、別 file 同名 sub node とは衝突しないことを明記した。
- `docs/spec/naming.md` に、public QualifiedID は main node の identity であり、sub node は public QualifiedID を持たないことを明記した。
- `docs/spec/naming.md` に、V01-ADR-078 の `<semantic-anchor-id>#<local-id>` は MCP query layer 用 synthetic ID であり、YAML authoring の public QualifiedID ではないことを明記した。
- `docs/spec/diagnostics.md` の `duplicate_node` を public node QualifiedID collision に限定し、同一 file 内 sub node local ID 重複用に `duplicate_sub_node` を追加した。
- Codex review result: `OK with minor fixes`。MCP synthetic ID wording と `edges.md` stale bare node wording の補強指摘を反映した。
- `docs/spec/naming.md` に、MCP schema / ObjectRef migration details は本 spec section 外であり、この resolver work では変更しないことを明記した。
- `docs/spec/edges.md` に、bare node/source resolution は `naming.md` §4 に従い、same-file file-private sub node / source first、same-module main node fallback であることを cross-reference として追記した。
- Design Records MCP validation: `validate_records(kind="spec")` returned `ok: true` before and after minor review fixes.
