# DRCLI-TASK-CLI-002-12: Route current-source discovery gap

- **status**: done
- **date**: 2026-10-07
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - DRCLI-TASK-CLI-002-09
- **outputs**:
  - DRCLI-WORK-CLI-002
  - DRCLI-TASK-CLI-002-13
  - DRCLI-TASK-CLI-002-14
  - DRCLI-TASK-CLI-002-09

## Goal

Route the exact new decision and authoring work required by the T09 current-source discovery blocker.

## Work

- Preserve T09's blocked Evidence.
- Create one decision Task for selected-repository current-source discovery.
- Create one authoring Task to project that decision into the exact affected DRCLI Specifications.
- Add the authoring Task as a prerequisite for T09 resumption.
- Preserve T10 and T11 after T09.

## Done condition

- The blocker has one decision owner and one bounded Specification writer.
- T09 cannot resume before the new Specification authoring completes.
- No unrelated Task graph or Specification content changes.

## Verification

- Confirm T13 owns only the new discovery judgment.
- Confirm T14 owns only the affected current-source/CLI consistency surface.
- Confirm T09 retains its original completion judgment.

## Evidence

- T09 stopped correctly on one materially unresolved design gap: the selected repository did not define how DRCLI obtains app-namespace/records-root associations.
- The user previously fixed the external behavior that DRCLI starts from cwd or --repo and traverses beneath that repository while honoring Git-compatible ignore rules.
- A repository-local mandatory manifest or per-repository registration would conflict with the accepted copy-and-run/shared-installation outcome.
- T13 and T14 were materialized as the minimal reconvergence route.
