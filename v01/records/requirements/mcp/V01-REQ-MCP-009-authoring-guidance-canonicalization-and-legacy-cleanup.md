# V01-REQ-MCP-009: Authoring guidance canonicalization and legacy cleanup が必要

- **id**: V01-REQ-MCP-009
- **status**: accepted
- **date**: 2026-06-01
- **source_refs**:
  - V01-REQ-MCP-005
  - SPEC-design-records-mcp-tools
  - SPEC-design-records-mcp-schema
  - docs/doc-policy.md
  - docs/guides/adr-authoring.md
  - docs/guides/spec-authoring.md
  - docs/guides/requirement-authoring.md
  - docs/guides/work-item-authoring.md
  - docs/guides/task-authoring.md
  - docs/guides/investigation-authoring.md
  - docs/guides/artifact-boundary.md
  - docs/adr-authoring-guide.md
  - docs/spec-authoring-guide.md
  - docs/requirements/README.md
  - docs/work-items/README.md
  - docs/tasks/README.md
  - docs/investigations/README.md
- **work_items**:
  - V01-WORK-MCP-009

## Requirement

Authoring guidance の正本を `docs/guides/*.md` と Design Records MCP の `list_authoring_guides` / `get_authoring_guidance` に寄せ、旧 top-level authoring guide と artifact README の重複・二重正本化を解消する必要がある。

特に、`docs/adr-authoring-guide.md`、`docs/spec-authoring-guide.md`、`docs/requirements/README.md`、`docs/work-items/README.md`、`docs/tasks/README.md`、`docs/investigations/README.md` は、既に `docs/guides/*-authoring.md` または `docs/guides/artifact-boundary.md` へ抽出された内容と重複している可能性がある。

今後の workflow metadata strictness や authoring rule の判断で、旧 README / 旧 guide と新 guides のどちらを正とするかが曖昧になることを避けるため、canonical source を明確化し、旧文書を削除・薄型入口化・維持のいずれかに分類する必要がある。

## Evidence

V01-REQ-MCP-005 により、project authoring guidance retrieval support として `list_authoring_guides` / `get_authoring_guidance` が導入され、guide ID 経由で authoring guidance を取得する方向が定義された。

一方で、旧 top-level guide と artifact README は引き続き repository に残っており、Codex / assistant review では `docs/guides/*` と README の両方を読む運用になりやすい。これは authoring guidance の正本が二重化しているように見える。

V01-WORK-MCP-006 / V01-TASK-MCP-006-01 の metadata validation gap review でも、workflow artifact の required metadata wording を確認する際に、新しい `docs/guides/*-authoring.md` と旧 README の両方が参照された。内容が大きく変わらないとしても、今後の更新時に差分・stale docs・誤読を生む risk がある。

2026-06-02 close evidence:

- V01-WORK-MCP-009 completed the inventory, cleanup policy decision, doc/spec/reference update, and close verification flow for this requirement.
- Canonical authoring guidance boundary is guide ID based retrieval through Design Records MCP `list_authoring_guides` / `get_authoring_guidance`; repo-local `docs/guides/*.md` remains the source, but source path is not the public contract.
- Legacy top-level guides were retained only as non-canonical compatibility entrypoints. Artifact README files were thinned to directory entrypoints and point to guide IDs instead of owning authoring rules.
- `validate_records` for V01-REQ-MCP-009, V01-WORK-MCP-009, and V01-TASK-MCP-009-01..04 returned `ok: true`, `diagnostics: null`.
- MCP authoring guidance retrieval was verified for guide IDs `adr-authoring`, `spec-authoring`, `requirement-authoring`, `work-item-authoring`, `task-authoring`, `investigation-authoring`, and `artifact-boundary`.
- Targeted stale wording search found no current-facing wording that makes legacy guides / README files canonical authoring guidance owners.
- Plain `git diff --check` for tracked modified files returned exit code 0. The untracked closeout files were checked with `git diff --no-index --check -- NUL <file>` and showed only line-ending normalization warnings, with no whitespace errors.
- V01-REQ-MCP-008 authoring transaction support, V01-WORK-MCP-006 metadata validation strictness, and accepted ADR supersession were not reopened.

## Required Outcome

- Authoring guidance の canonical source を `docs/guides/*.md` と MCP authoring guidance retrieval に寄せるか判断されている。
- 旧 top-level authoring guide と artifact README について、削除、薄型入口化、維持の分類が完了している。
- 削除または薄型化する場合、参照元、MCP implementation、tests、doc-policy、spec への影響が確認されている。
- Human entrypoint として最低限残すべき導線がある場合、その責務が明確になっている。
- 旧文書を残す場合は、残す理由と正本ではないことが明示されている。

## Explicitly Excluded Scope

- V01-WORK-MCP-006 の workflow artifact metadata validation strictness 判断に混ぜない。
- V01-TASK-MCP-006-02 の missing / empty / date diagnostic contract decision をブロックしない。
- 旧 guide / README を即時削除すること自体をこの requirement では決定しない。
- Accepted ADR をこの requirement だけで supersede しない。
- Authoring transaction support の実装は V01-REQ-MCP-008 の scope とし、本 requirement では扱わない。

## Boundary

この requirement は、authoring guidance の正本整理と legacy guide / README cleanup の必要性を捕捉する。

実際にどのファイルを削除・薄型化・維持するか、どの spec / doc-policy / tests を更新するかは、後続 work item と task で inventory review を行って判断する。

本 requirement は workflow artifact の metadata strictness contract そのものを所有しない。V01-WORK-MCP-006 は引き続き Design Records MCP の validation strictness 判断と実現を所有する。
