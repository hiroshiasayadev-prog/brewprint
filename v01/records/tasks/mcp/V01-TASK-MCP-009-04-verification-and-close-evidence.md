# V01-TASK-MCP-009-04: Verification and close evidence

- **id**: V01-TASK-MCP-009-04
- **status**: done
- **date**: 2026-06-02
- **work_item**: V01-WORK-MCP-009
- **source_requirement**: V01-REQ-MCP-009
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-009-03
- **outputs**:
  - validate_records result
  - stale reference check result
  - MCP authoring guidance retrieval verification
  - V01-WORK-MCP-009 / V01-REQ-MCP-009 close evidence if complete

## Goal

Metadata relation、stale reference、MCP authoring guidance retrieval behavior を確認し、close 可能なら V01-WORK-MCP-009 / V01-REQ-MCP-009 の status 同期まで行う。

## Work

- `validate_records` で `V01-REQ-MCP-009`, `V01-WORK-MCP-009`, and `V01-TASK-MCP-009-01` through `V01-TASK-MCP-009-04` を確認する。
- Stale reference check を実施し、legacy docs の削除・薄型化・非正本維持に伴う壊れた参照や misleading reference が残っていないか確認する。
- Design Records MCP の `list_authoring_guides` / `get_authoring_guidance` で canonical guide retrieval が期待どおり機能することを確認する。
- V01-TASK-MCP-009-03 で implementation / tests を変更した場合は、対象テスト結果を確認する。
- Completion Condition を満たす場合は、V01-WORK-MCP-009 と V01-REQ-MCP-009 に close evidence を残し、status を同期する。
- Completion Condition を満たさない場合は、未解決事項と blocker を evidence として残す。

## Done condition

- `V01-REQ-MCP-009`, `V01-WORK-MCP-009`, and `V01-TASK-MCP-009-01` through `V01-TASK-MCP-009-04` の metadata relation validation result が記録されている。
- Stale reference check result が記録されている。
- MCP authoring guidance retrieval verification result が記録されている。
- V01-WORK-MCP-009 / V01-REQ-MCP-009 を close できるかが evidence に基づいて判断されている。
- Close する場合は status / evidence が同期されている。Close しない場合は unresolved issue が明確になっている。

## Verification

- `validate_records` の対象と結果を Evidence に記録する。
- Stale reference check の command / MCP result / manual review result を Evidence に記録する。
- MCP authoring guidance retrieval verification の対象 guide ID と結果を Evidence に記録する。
- 必要な tests を実行した場合は結果を Evidence に記録する。未実施の場合は理由を記録する。

## Evidence

2026-06-02 verification and close result:

Validation results:

- `validate_records(kind=requirement, id_range=V01-REQ-MCP-009..V01-REQ-MCP-009)` returned `ok: true`, `diagnostics: null`.
- `validate_records(kind=work_item, id_range=V01-WORK-MCP-009..V01-WORK-MCP-009)` returned `ok: true`, `diagnostics: null`.
- `validate_records(kind=task, id_range=V01-TASK-MCP-009-01..V01-TASK-MCP-009-04)` returned `ok: true`, `diagnostics: null`.

MCP authoring guidance retrieval results:

- `list_authoring_guides` returned canonical guide IDs `adr-authoring`, `artifact-boundary`, `investigation-authoring`, `requirement-authoring`, `spec-authoring`, `task-authoring`, and `work-item-authoring`.
- `get_authoring_guidance` succeeded for `adr-authoring`, `spec-authoring`, `requirement-authoring`, `work-item-authoring`, `task-authoring`, `investigation-authoring`, and `artifact-boundary`.
- Public response shape uses guide IDs. `list_authoring_guides` returned `id`, `title`, and `abstract` entries with no source path field; `get_authoring_guidance` returned `id`, `title`, and `content`. Guide content may include maintainer migration notes mentioning extracted-from paths, but those notes explicitly state they are not part of the public guide retrieval contract.

Stale wording search result:

- Targeted search covered `docs/adr-authoring-guide.md`, `docs/spec-authoring-guide.md`, `docs/requirements/README.md`, `docs/work-items/README.md`, `docs/tasks/README.md`, `docs/investigations/README.md`, `docs/doc-policy.md`, `docs/TASKS.md`, and `docs/spec/concepts/**/*.md`.
- Search terms covered legacy / canonical authoring wording, legacy guide / README paths, source path / public contract wording, and ownership wording around authoring guidance.
- No current-facing wording was found that makes legacy top-level guides or artifact README files canonical authoring guidance owners. Remaining hits are expected non-canonical compatibility statements, directory/archive orientation, MCP public-contract wording, or unrelated canonical reference terminology in traceability specs.

`git diff --check` result:

- Final plain `git diff --check` for tracked modified files returned exit code 0.
- The closeout records are untracked in the current worktree, so they were checked separately with `git diff --no-index --check -- NUL <file>` for V01-TASK-MCP-009-04, V01-WORK-MCP-009, and V01-REQ-MCP-009. Those commands returned exit code 1 because `--no-index` compares `NUL` to a real file, but output contained only line-ending normalization warnings and no whitespace errors.
- Git printed existing line-ending normalization warnings for modified Markdown / Go files (`LF will be replaced by CRLF`), but no whitespace errors.

Files modified during close:

- `docs/tasks/mcp/TASK-MCP-009-04-verification-and-close-evidence.md`
- `docs/work-items/mcp/WORK-MCP-009-authoring-guidance-canonicalization-and-legacy-cleanup.md`
- `docs/requirements/mcp/REQ-MCP-009-authoring-guidance-canonicalization-and-legacy-cleanup.md`

Close decision:

- Completion conditions for V01-TASK-MCP-009-04 are satisfied.
- V01-WORK-MCP-009 completion conditions are satisfied by V01-TASK-MCP-009-01 inventory, V01-TASK-MCP-009-02 cleanup policy, V01-TASK-MCP-009-03 cleanup/reference updates, and this verification.
- V01-REQ-MCP-009 required outcomes are satisfied. Nearby completed MCP requirements use status `accepted`, so V01-REQ-MCP-009 was closed as `accepted` rather than inventing a new requirement status.
- No unresolved blocker remains for V01-REQ-MCP-009 / V01-WORK-MCP-009.
