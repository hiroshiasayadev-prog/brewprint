# DRCLI-TASK-CLI-002-14: Author selected-repository current-source discovery contracts

- **status**: done
- **date**: 2026-10-07
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: authoring
- **estimate**: 0.5d-1d
- **depends_on**:
  - DRCLI-TASK-CLI-002-13
- **outputs**:
  - spec:drcli.design_records_cli.current_record_model.current_source_corpus
  - spec:drcli.design_records_cli.current_record_model
  - spec:drcli.design_records_cli.cli.repository_selection_and_traversal

## Goal

Project the accepted selected-repository current-source discovery decision into the bounded DRCLI Specification surface needed to unblock integration.

## Work

Update only the direct affected consistency surface:

- current-record-model/current-source-corpus.md;
- current-record-model/index.md when topic/boundary wording requires it;
- cli/repository-selection-and-traversal.md;
- diagnostics only if T13 explicitly requires a new canonical diagnostic owner and the existing diagnostic catalog is the correct owner.

Preserve PRODUCT ownership of artifact placement and canonical identity.
Remove the stale copied DRMCP assumption that DRCLI requires preconfigured app-namespace/records-root associations when that contradicts T13.

Do not modify root DRCLI indexes; T09 still owns integration.
Do not modify DRMCP.

## Done condition

- Selected repository to current-source corpus construction is fully specified.
- App namespace resolution is explicit and non-path-canonical.
- Ignore behavior composes with root discovery.
- Ambiguous/unresolvable roots have explicit diagnostics.
- No mandatory manifest/registration is introduced.
- T09 can resume without new judgment.

## Verification

- Trace T13 decisions into exact Specification sections.
- Confirm no stale "DRCLI does not auto-discover configured current sources" contradiction remains.
- Confirm malformed sources remain observable under broad validation once a root is resolved.
- Confirm only the direct affected files plus this Task record changed.

## Evidence

- Input decision: DRCLI-TASK-CLI-002-13 D-001 through D-009, consumed without reopening any decision.
- Updated `spec:drcli.design_records_cli.current_record_model.current_source_corpus` to define automatic candidate `records`-root discovery, canonical-identity namespace evidence, exact-one-namespace resolution, empty/zero/multi-namespace handling, duplicate-app-root failure, ignore composition, and post-resolution inclusion of malformed and unadmitted sources.
- Updated `spec:drcli.design_records_cli.current_record_model` to remove explicit-configuration bootstrapping and project the selected-repository discovery progression.
- Updated `spec:drcli.design_records_cli.cli.repository_selection_and_traversal` to make ignore pruning precede root discovery, namespace evidence, and corpus construction, while preserving cwd/`--repo` selection and `--no-ignore` as an ignore-only override.
- Removed the incompatible copied DRMCP current-source-configuration coverage claim that required explicit preconfiguration. PRODUCT placement authority remains unchanged.
- Reused the existing context-free `configuration_failure` operation code. No diagnostic code, diagnostic context field, or diagnostic Specification was added or changed.
- No repository manifest, persistent registration, per-repository installation, or path-derived app identity was introduced.
- Verification confirmed that a repository with no candidate root has a valid empty current state, a resolved root retains malformed and unadmitted in-corpus sources for broad validation, and current-source resolution failures return no partial normal result.
- A focused DRCLI-tree scan found legacy configured-source wording in `spec:drcli.design_records_cli.current_record_model.current_record_scopes`. That file is outside T14's writer boundary; T09 explicitly owns bounded consistency edits across the completed DRCLI Specification tree, so this is a mechanical T09 integration repair and requires no new design judgment.
- This Task wrote only the three declared Specification outputs and this Task record. Root DRCLI indexes, DRMCP records, the parent Work Item, and Git state were not modified.
