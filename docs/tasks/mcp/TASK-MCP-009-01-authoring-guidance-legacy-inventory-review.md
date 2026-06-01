# TASK-MCP-009-01: Authoring guidance legacy inventory review

- **id**: TASK-MCP-009-01
- **status**: done
- **date**: 2026-06-02
- **work_item**: WORK-MCP-009
- **source_requirement**: REQ-MCP-009
- **estimate**: 0.5d-1d
- **depends_on**:
- **outputs**:
  - Legacy guide / README inventory
  - Canonical overlap summary
  - Cleanup classification input

## Goal

`docs/guides/*.md` と旧 top-level authoring guides / artifact README の重複・差分・入口責務を整理し、cleanup 判断の材料を作る。

この task では、どの legacy file を削除・薄型化・維持するかを決定しない。削除や thinning も行わない。

## Work

- `docs/guides/*.md` の authoring guidance と artifact boundary guide の現在内容を確認する。
- 以下の legacy / entrypoint docs を確認し、canonical guide との overlap、差分、残存 entrypoint 責務を整理する。
  - `docs/adr-authoring-guide.md`
  - `docs/spec-authoring-guide.md`
  - `docs/requirements/README.md`
  - `docs/work-items/README.md`
  - `docs/tasks/README.md`
  - `docs/investigations/README.md`
- `docs/doc-policy.md` と Design Records MCP authoring guidance retrieval contract の入口説明を確認し、legacy docs と矛盾して見える箇所を抽出する。
- Legacy docs ごとに、後続 TASK-MCP-009-02 が分類判断できるように以下を記録する。
  - canonical guide と重複している内容
  - canonical guide に未抽出の内容
  - human entrypoint として残す可能性がある内容
  - stale / misleading /二重正本に見える wording
- Inventory は調査結果として扱い、cleanup decision は後続 task に渡す。

## Done condition

- 対象となる current guide / legacy guide / README の inventory が完了している。
- Canonical overlap と差分が、後続判断に使える粒度で整理されている。
- Cleanup classification input が作成され、TASK-MCP-009-02 が delete / thin entrypoint / keep non-canonical を判断できる状態になっている。
- この task で legacy docs の削除、薄型化、正本変更を行っていない。

## Verification

- 読んだ files / authoring guidance IDs を Evidence に記録する。
- `validate_records` で `REQ-MCP-009`, `WORK-MCP-009`, and this task の metadata relation を確認する。
- Markdown 変更を行わない task のため、format / implementation tests は原則不要とする。

## Evidence

2026-06-02 review result:

- Reviewed current startup / policy / guide context:
  - `AGENTS.md`
  - `docs/prompt_chappy.md`
  - `docs/doc-policy.md`
  - `docs/guides/adr-authoring.md`
  - `docs/guides/spec-authoring.md`
  - `docs/guides/requirement-authoring.md`
  - `docs/guides/work-item-authoring.md`
  - `docs/guides/task-authoring.md`
  - `docs/guides/investigation-authoring.md`
  - `docs/guides/artifact-boundary.md`
- Reviewed legacy / entrypoint docs:
  - `docs/adr-authoring-guide.md`
  - `docs/spec-authoring-guide.md`
  - `docs/requirements/README.md`
  - `docs/work-items/README.md`
  - `docs/tasks/README.md`
  - `docs/investigations/README.md`
- Reviewed MCP public contract context:
  - `docs/spec/design-records-mcp/tools.md`
  - `docs/spec/design-records-mcp/schema.md`
- Inventory finding: legacy docs still present themselves as authoring guidance owners. Examples include `docs/adr-authoring-guide.md` saying ADR format / metadata / Evidence rules are owned by that guide, and requirements / work-items README wording saying those README files own authoring guidance. This conflicts with the newer `docs/doc-policy.md` and MCP authoring guidance retrieval boundary, where guide IDs and `docs/guides/*.md` are the authoring guidance source.
- Inventory finding: legacy-only material remains before deletion can be safe. `docs/adr-authoring-guide.md` has detailed reference examples and anti-pattern sections that are more specific than `docs/guides/adr-authoring.md`. `docs/spec-authoring-guide.md` has exact front matter and origin-note examples that are more specific than `docs/guides/spec-authoring.md`. `docs/tasks/README.md` contains legacy M-series / archive boundary notes. `docs/investigations/README.md` still has directory / index role material.
- Inventory finding: `docs/spec/design-records-mcp/tools.md` and `docs/spec/design-records-mcp/schema.md` are Design Records MCP public tool/schema contracts, not legacy cleanup targets. Their authoring-guidance sections support guide ID based retrieval and explicitly keep guide source file path out of the public response contract.
- Cleanup input for TASK-MCP-009-02: classify the legacy docs so they no longer present themselves as canonical authoring guidance; move or intentionally drop legacy-only material before deleting files; keep MCP specs as public contract docs rather than cleanup targets.
- Verification: `validate_records` for `kind=task`, `id_range=TASK-MCP-009-01..TASK-MCP-009-02` returned `ok: true` and `diagnostics: null`.
- This task performed inventory / review only. No legacy doc deletion, thinning, source-of-truth change, MCP implementation change, git add, or commit was performed.
