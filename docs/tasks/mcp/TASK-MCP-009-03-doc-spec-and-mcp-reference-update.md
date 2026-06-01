# TASK-MCP-009-03: Doc, spec, and MCP reference update

- **id**: TASK-MCP-009-03
- **status**: done
- **date**: 2026-06-02
- **work_item**: WORK-MCP-009
- **source_requirement**: REQ-MCP-009
- **estimate**: 1d-2d
- **depends_on**:
  - TASK-MCP-009-02
- **outputs**:
  - Updated doc-policy / guides / legacy docs as needed
  - Updated Design Records MCP tools/schema spec if needed
  - Updated MCP implementation/tests if references or guide catalog behavior need changes

## Goal

TASK-MCP-009-02 の cleanup policy decision に従い、必要な docs / spec / MCP reference updates を適用する。

この task は authoring guidance canonicalization と legacy cleanup の反映を扱う。REQ-MCP-008 authoring transaction implementation には拡張しない。

## Work

- TASK-MCP-009-02 の per-file classification を確認し、更新対象と非対象を明確にする。
- `docs/doc-policy.md` の authoring guidance / startup / responsibility boundary wording が cleanup policy と矛盾する場合は更新する。
- `docs/guides/*.md` に canonical source として不足する内容がある場合は、対象 guide の責務範囲内で補う。
- Legacy docs を `delete` / `thin entrypoint` / `keep non-canonical` classification に従って更新する。
- `SPEC-design-records-mcp-tools` / `SPEC-design-records-mcp-schema` に guide catalog / retrieval / source path boundary の stale reference がある場合は更新する。
- MCP implementation / tests に guide catalog source path や removed legacy file への依存がある場合のみ、必要最小限で更新する。
- 更新対象が implementation / tests に及ぶ場合も、REQ-MCP-008 の proposal / accept / authoring transaction tool implementation には踏み込まない。

## Done condition

- TASK-MCP-009-02 の cleanup policy decision に従った docs / spec / MCP reference updates が完了している。
- Legacy docs が二重正本として読まれない状態になっているか、非正本として残す理由が明示されている。
- Design Records MCP authoring guidance retrieval の public contract が、guide ID を public reference とする境界と矛盾していない。
- Required implementation / tests update がある場合は適用されている。
- REQ-MCP-008 authoring transaction implementation に scope creep していない。

## Verification

- 更新した files と、TASK-MCP-009-02 classification との対応を Evidence に記録する。
- Stale reference check を実施し、legacy paths / canonical guide IDs / guide source path boundary の残存状況を Evidence に記録する。
- Spec または implementation を変更した場合は、対象 tests または未実施理由を Evidence に記録する。
- `validate_records` で `REQ-MCP-009`, `WORK-MCP-009`, and this task の metadata relation を確認する。

## Evidence

2026-06-02 update result:

Changed files:

- `docs/adr-authoring-guide.md`: rewrote as a legacy compatibility entrypoint. It explicitly says it is not canonical authoring guidance, points to guide IDs `adr-authoring` and `artifact-boundary`, and states that guide source paths are not public contract.
- `docs/spec-authoring-guide.md`: rewrote as a legacy compatibility entrypoint. It explicitly says it is not canonical authoring guidance, points to guide IDs `spec-authoring` and `artifact-boundary`, and removes the legacy cross-reference to `docs/adr-authoring-guide.md`.
- `docs/requirements/README.md`: thinned to a directory entrypoint. It keeps directory role / layout orientation and points to guide IDs `requirement-authoring` and `artifact-boundary`.
- `docs/work-items/README.md`: thinned to a directory entrypoint. It keeps directory role / layout orientation, preserves the legacy milestone boundary note, and points to guide IDs `work-item-authoring` and `artifact-boundary`.
- `docs/tasks/README.md`: thinned to a directory entrypoint. It keeps directory role / layout orientation, preserves legacy M-series / archive boundary notes, and points to guide IDs `task-authoring` and `artifact-boundary`.
- `docs/investigations/README.md`: thinned to a directory entrypoint. It keeps directory role / layout orientation, removes the stale-prone current investigations index, points to guide IDs `investigation-authoring` and `artifact-boundary`, and directs record discovery to `list_records(kind=investigation)`.
- `docs/TASKS.md`: updated the migration note so new task / work item authoring guidance points to guide IDs through Design Records MCP rather than README paths.
- `docs/spec/concepts/project-artifact-model/index.md`: updated current ownership boundary wording to use guide IDs and bumped `last_updated` to `2026-06-02`.
- `docs/spec/concepts/traceability/metadata-schema.md`: updated investigation authoring ownership wording to guide ID `investigation-authoring` and bumped `last_updated` to `2026-06-02`.

Files intentionally unchanged:

- `docs/doc-policy.md`: current wording already says authoring guidance is retrieved through `list_authoring_guides` / `get_authoring_guidance`, guide IDs are public references, and guide source paths are not public contract.
- `docs/spec/design-records-mcp/tools.md`: unchanged because it owns the MCP public tool contract and already keeps guide source path out of the response contract.
- `docs/spec/design-records-mcp/schema.md`: unchanged because it owns the MCP source model contract and already treats `docs/guides/*.md` as internal source for guide ID based retrieval.
- MCP implementation / tests: unchanged. No hardcoded dependency on removed legacy files was found or needed because the legacy files were retained as non-canonical compatibility entrypoints.

Legacy-only material handling:

- `docs/adr-authoring-guide.md` detailed reference example / anti-pattern expansions were intentionally dropped from the legacy entrypoint. The reusable rule content remains covered by guide IDs `adr-authoring` and `artifact-boundary`; retaining long examples here would keep the file looking canonical.
- `docs/spec-authoring-guide.md` detailed front matter / origin-note examples were intentionally dropped from the legacy entrypoint for the same reason. The canonical update rule remains guide ID `spec-authoring`.
- `docs/tasks/README.md` preserved the legacy M-series / archive boundary as directory-entrypoint material, not authoring rules.
- `docs/investigations/README.md` preserved the directory role, not format, lifecycle guidance, or an investigation record list.

Search / MCP guidance checks:

- `list_authoring_guides` returned the expected guide IDs: `adr-authoring`, `artifact-boundary`, `investigation-authoring`, `requirement-authoring`, `spec-authoring`, `task-authoring`, and `work-item-authoring`.
- `get_authoring_guidance` for `adr-authoring` and `spec-authoring` returned Markdown content through guide ID lookup.
- Targeted stale wording search over the edited legacy entrypoints, `docs/doc-policy.md`, Design Records MCP tools/schema specs, `docs/TASKS.md`, and the two concept specs found no remaining wording that says a legacy guide / README owns canonical authoring guidance. Remaining hits were expected non-canonical statements, compatibility notes, or unrelated source-of-truth terminology.
- Broader legacy path search still finds historical ADR / investigation / task evidence and guide migration-note references to old paths. Those are historical provenance or already-closed task evidence, not current authoring guidance entrypoints.

Validation:

- Additional cleanup requested on 2026-06-02 removed the manually maintained `docs/investigations/README.md` current investigations list. Investigation list discovery is now documented as Design Records MCP `list_records(kind=investigation)`.
- `list_records(kind=investigation, limit=5)` returned investigation records through Design Records MCP, confirming the intended discovery path.
- Rerun `validate_records()` after the additional cleanup returned `ok: false` due to existing repository-wide diagnostics outside this cleanup scope: info-level investigation follow-up candidate diagnostics in `INV-DATA-001` / `INV-DATA-002`, and missing required metadata errors in existing `TASK-MCP-005-01` through `TASK-MCP-005-03`.
- Targeted rerun `validate_records(kind=task, id_range=TASK-MCP-009-03..TASK-MCP-009-03)` returned `ok: true`, `diagnostics: null`.
- Targeted rerun `validate_records(kind=investigation)` returned `ok: true` with the existing info-level `INV-DATA-001` / `INV-DATA-002` follow-up candidate diagnostics.
- Rerun `git diff --check` after the additional cleanup returned exit code 0. Git reported only existing line-ending normalization warnings (`LF will be replaced by CRLF`) for touched files.
- `validate_records(kind=task, id_range=TASK-MCP-009-03..TASK-MCP-009-03)` returned `ok: true`, `diagnostics: null`.
- `validate_records(kind=requirement, id_range=REQ-MCP-009..REQ-MCP-009)` returned `ok: true`, `diagnostics: null`.
- `validate_records(kind=work_item, id_range=WORK-MCP-009..WORK-MCP-009)` returned `ok: true`, `diagnostics: null`.
- `git diff --check` on modified files returned exit code 0. Git reported only existing line-ending normalization warnings (`LF will be replaced by CRLF`) for touched Markdown files.

Scope control:

- Did not implement MCP behavior.
- Did not change Go code.
- Did not reopen REQ-MCP-008 or WORK-MCP-006.
- Did not close WORK-MCP-009 / REQ-MCP-009.
- Did not git add or commit.
