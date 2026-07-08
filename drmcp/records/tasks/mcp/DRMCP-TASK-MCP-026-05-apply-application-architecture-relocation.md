# DRMCP-TASK-MCP-026-05: Apply application-architecture relocation

- **id**: DRMCP-TASK-MCP-026-05
- **status**: done
- **date**: 2026-07-08
- **work_item**: DRMCP-WORK-MCP-026
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-026-04
- **outputs**:
  - `spec:drmcp.implementation.application_architecture`

## Goal

Apply the reviewed migration script to relocate the application-architecture Specification tree and synchronize spec-scope canonical refs.

## Work

Apply-mode migration is allowed only after T04 returns `PASS`.

The migration must:

- move the five application-architecture Specification files into `drmcp/records/spec/implementation/application-architecture/`;
- update canonical refs under `drmcp/records/spec/` from `spec:drmcp.application_architecture` to `spec:drmcp.implementation.application_architecture`;
- update root and child `id` metadata;
- update child `parent` metadata;
- update topic tables, related-spec rows, and prose refs identified by T01;
- remove the old application-architecture tree without leaving a compatibility stub;
- preserve architecture semantics.

Do not redesign `spec:drmcp.implementation` root content in this Task unless validation blocks relocation and a new decision owner accepts the root update.

## Done condition

- The reviewed script is applied successfully.
- Old application-architecture files no longer remain at the old path.
- New application-architecture files exist under `drmcp/records/spec/implementation/application-architecture/`.
- No old `spec:drmcp.application_architecture` refs remain under `drmcp/records/spec/`, except if explicitly recorded as a blocked validation exception.
- No compatibility stub remains under the old tree.
- No ADR, Requirement, Work Item, Task, or Investigation history cleanup is performed.

## Verification

- T04 returned `PASS` and allowed T05 apply.
- The reviewed migration script was applied by the user on 2026-07-08.
- Apply output reported `result: applied`.
- Apply output reported five moved files.
- Apply output reported 12 rewritten files.
- User ran `Select-String -Path "drmcp/records/spec/**/*.md" -Pattern "spec:drmcp.application_architecture"` after apply.
- The user-provided command output contained no old-ref matches.
- No ADR creation, stage, commit, or push was performed by this Task.

## Evidence

Applied command:

```text
python drmcp/scripts/relocate_application_architecture_specs.py --repo-root C:\Users\imved\projects\brewprint --apply
```

Apply output:

```text
mode: apply
result: applied
moved_files: 5
rewritten_files: 12
```

Post-apply old-ref check command:

```text
Select-String -Path "drmcp/records/spec/**/*.md" -Pattern "spec:drmcp.application_architecture"
```

Post-apply old-ref check result:

- No matches were reported in the user-provided output.

Observed scoped git status from the user-provided output:

- The five old application-architecture files are deleted at the old path.
- The new application-architecture directory exists under `drmcp/records/spec/implementation/application-architecture/`.
- The expected spec-scope in-place rewrite files are modified.
- The output also contained unrelated modified and untracked records outside the T05 migration scope.
- Those unrelated records were not attributed to T05.

T05 applied the reviewed migration.
T06 must review the relocated Specification tree and synchronized refs before W026 closure.
