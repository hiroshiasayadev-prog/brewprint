# V01-TASK-MCP-009-02: Cleanup policy and canonical boundary decision

- **id**: V01-TASK-MCP-009-02
- **status**: done
- **date**: 2026-06-02
- **work_item**: V01-WORK-MCP-009
- **source_requirement**: V01-REQ-MCP-009
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-MCP-009-01
- **outputs**:
  - Cleanup policy decision
  - Per-file classification: delete / thin entrypoint / keep non-canonical
  - Reviewer handoff or review result if needed

## Goal

V01-TASK-MCP-009-01 の inventory をもとに、authoring guidance canonical source と legacy docs の扱いを決める。

曖昧さが残る場合は、Opus / reviewer 向けの handoff または review result を evidence として残し、判断根拠を明確にする。

## Work

- V01-TASK-MCP-009-01 の inventory を確認し、各 legacy file の残存価値と stale risk を評価する。
- Authoring guidance の canonical source を `docs/guides/*.md` と Design Records MCP authoring guidance retrieval に寄せる前提が、doc-policy / spec / existing implementation と矛盾しないか確認する。
- 対象ファイルごとに以下のいずれかへ分類する。
  - delete
  - thin entrypoint
  - keep non-canonical
- `thin entrypoint` または `keep non-canonical` の場合は、正本ではないこと、入口として残す理由、更新責務を明確にする。
- `delete` の場合は、削除前に必要な参照元更新、MCP catalog / tests 影響、human entrypoint 影響を明確にする。
- Accepted ADR supersession が必要かを確認する。通常は scope 外とし、必要性が具体的に出た場合のみ別途扱う。
- Ambiguity が残る場合は reviewer handoff prompt または review findings を残す。

## Done condition

- Cleanup policy decision が記録されている。
- 対象 legacy file ごとの classification が `delete` / `thin entrypoint` / `keep non-canonical` のいずれかで記録されている。
- V01-TASK-MCP-009-03 が具体的に更新対象を選べる状態になっている。
- V01-REQ-MCP-008 authoring transaction implementation、V01-WORK-MCP-006 metadata validation strictness、accepted ADR supersession を不要に再オープンしていない。
- 判断に ambiguity が残る場合は、reviewer handoff または review result が evidence として残っている。

## Verification

- V01-TASK-MCP-009-01 の inventory と cleanup policy decision の対応を Evidence に記録する。
- `validate_records` で `V01-REQ-MCP-009`, `V01-WORK-MCP-009`, and this task の metadata relation を確認する。
- Reviewer handoff / review result を使った場合は、その unresolved issues を Evidence に記録する。

## Evidence

2026-06-02 cleanup policy decision:

Canonical source boundary:

- Canonical authoring guidance is the guide ID based surface exposed by Design Records MCP `list_authoring_guides` / `get_authoring_guidance`.
- Repo-local source for those guides is `docs/guides/*.md`, but guide source path is not a public response contract. Public references should use guide IDs such as `task-authoring` and `artifact-boundary`.
- `docs/doc-policy.md` remains the startup / safety policy and MCP entrypoint delegation doc. It does not own artifact-specific authoring rules.
- Legacy top-level authoring guides and artifact README files must not present themselves as canonical authoring guidance owners after cleanup.
- Design Records MCP specs remain public tool/schema contracts. They are not cleanup targets merely because they describe authoring guidance retrieval.
- Accepted ADR supersession is not required for this cleanup decision. No V01-WORK-MCP-006 metadata strictness decision and no V01-REQ-MCP-008 authoring transaction implementation scope is reopened.

V01-TASK-MCP-009-01 inventory input used:

- Legacy docs still present themselves as authoring guidance owners.
- Legacy-only material must be handled before deletion.
- Design Records MCP specs must not be treated as cleanup targets.

Per-file cleanup classification:

| file | classification | decision / reason | V01-TASK-MCP-009-03 handling |
|---|---|---|---|
| `docs/adr-authoring-guide.md` | delete | Legacy top-level guide overlaps `docs/guides/adr-authoring.md` and currently presents itself as the ADR authoring source. It is a delete candidate, but only after legacy-only reference examples and anti-pattern details are either intentionally dropped or moved into `docs/guides/adr-authoring.md` / `docs/guides/artifact-boundary.md` as appropriate. | Compare legacy reference examples / anti-patterns with canonical guides. Move only material that should remain canonical; record intentionally dropped material in evidence; then delete or leave a temporary non-canonical pointer only if deletion is blocked by stale references. |
| `docs/spec-authoring-guide.md` | delete | Legacy top-level guide overlaps `docs/guides/spec-authoring.md` and currently points to another legacy authoring guide. It is a delete candidate, but exact front matter and origin-note examples are legacy-only material that need an explicit keep/drop decision before deletion. | Move exact front matter / origin-note examples into `docs/guides/spec-authoring.md` if still useful. Remove legacy cross-reference to `docs/adr-authoring-guide.md`; then delete or leave a temporary non-canonical pointer only if deletion is blocked by stale references. |
| `docs/requirements/README.md` | thin entrypoint | Requirement authoring rules are canonical in `requirement-authoring` and boundary details in `artifact-boundary`. The README can remain as a human directory entrypoint, but not as authoring guidance owner. | Replace authoring-rule body with a thin entrypoint that points to guide ID `requirement-authoring` and `artifact-boundary`, plus the directory purpose if useful. |
| `docs/work-items/README.md` | thin entrypoint | Work item authoring rules are canonical in `work-item-authoring` and boundary details in `artifact-boundary`. Current wording says this README owns authoring guidance, which should be removed. | Replace authoring-rule body with a thin entrypoint that points to guide ID `work-item-authoring` and `artifact-boundary`, preserving only directory-level orientation. |
| `docs/tasks/README.md` | thin entrypoint | Task authoring rules are canonical in `task-authoring` and boundary details in `artifact-boundary`. The README may still be useful for task directory orientation and legacy M-series / archive boundary notes. | Thin to guide ID `task-authoring` / `artifact-boundary` entrypoint. Preserve useful legacy M-series / archive boundary notes or move them to `docs/doc-policy.md` if they are startup/archive policy rather than task authoring guidance. |
| `docs/investigations/README.md` | thin entrypoint | Investigation authoring rules are canonical in `investigation-authoring` and boundary details in `artifact-boundary`. The README still has possible directory / index role value. | Thin to guide ID `investigation-authoring` / `artifact-boundary` entrypoint. Preserve directory layout / current investigations index role only if it remains useful and clearly non-canonical for authoring rules. |
| `docs/doc-policy.md` | keep non-canonical | This is not legacy authoring guidance. It owns startup, safety, MCP entrypoint delegation, and archive notes. Current wording already says artifact-specific rules belong to authoring guidance tools and guide source path is not public contract. | Keep. Update only if stale references remain after legacy cleanup, or if archive notes from `docs/tasks/README.md` need to be centralized here. Do not add artifact-specific authoring rules. |
| `docs/spec/design-records-mcp/tools.md` | keep non-canonical | This is MCP public tool contract. It defines `list_authoring_guides` / `get_authoring_guidance`, read-only behavior, guide ID lookup, and response shape. It is not an authoring guide cleanup target. | Keep. Update only if stale wording is found during V01-TASK-MCP-009-03, especially any wording that implies guide source path is public contract. Do not thin or delete. |
| `docs/spec/design-records-mcp/schema.md` | keep non-canonical | This is MCP schema/source model contract. It defines authoring guidance as a read-only guidance source separate from Design Records record kinds and says guide source path is not public response contract. It is not an authoring guide cleanup target. | Keep. Update only if stale wording is found during V01-TASK-MCP-009-03. Do not thin or delete. |

Handling of legacy-only material:

- Legacy-only material must not be silently lost. V01-TASK-MCP-009-03 should review each unique section and either move it into the appropriate canonical guide or record that it was intentionally dropped because it is obsolete, too example-specific, or outside the canonical guide scope.
- `docs/adr-authoring-guide.md`: review `reference examples` and `アンチパターン` sections. Preserve only reusable authoring guidance, not stale task history.
- `docs/spec-authoring-guide.md`: review exact front matter and origin-note examples. Preserve only examples that remain current for spec authors.
- `docs/tasks/README.md`: preserve legacy M-series / archive boundary notes if they are still useful; otherwise record intentional drop or move to startup/archive policy.
- `docs/investigations/README.md`: preserve directory / index role only as non-canonical entrypoint material.

What V01-TASK-MCP-009-03 should edit:

- Update canonical guides only where legacy-only material is intentionally retained:
  - likely `docs/guides/adr-authoring.md`
  - likely `docs/guides/spec-authoring.md`
  - possibly `docs/guides/task-authoring.md`
  - possibly `docs/guides/investigation-authoring.md`
  - `docs/guides/artifact-boundary.md` only for responsibility-boundary material
- Delete `docs/adr-authoring-guide.md` and `docs/spec-authoring-guide.md` after retained material and references are handled.
- Thin `docs/requirements/README.md`, `docs/work-items/README.md`, `docs/tasks/README.md`, and `docs/investigations/README.md` so they point to guide IDs and no longer claim authoring guidance ownership.
- Keep `docs/doc-policy.md`, `docs/spec/design-records-mcp/tools.md`, and `docs/spec/design-records-mcp/schema.md`; update them only for stale references created or revealed by cleanup.
- Run a stale reference check for legacy paths, guide IDs, and wording that treats `docs/guides/*.md` source paths as public contract.
- Do not change MCP implementation unless a concrete dependency on removed legacy files is found.

Verification:

- `validate_records` for `kind=task`, `id_range=V01-TASK-MCP-009-01..V01-TASK-MCP-009-02` returned `ok: true` and `diagnostics: null`.

No cleanup edits, file deletions, MCP implementation changes, git add, or commit were performed in this task.
