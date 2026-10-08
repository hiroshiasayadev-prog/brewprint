# DRCLI-TASK-CLI-002-03: Author DRCLI artifact and current-record model specs

- **status**: done
- **date**: 2026-10-07
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: authoring
- **estimate**: 1d-2d
- **depends_on**:
  - DRCLI-TASK-CLI-002-01
- **outputs**:
  - spec:drcli.design_records_cli.artifacts
  - spec:drcli.design_records_cli.current_record_model

## Goal

Create the DRCLI artifact-contract and current-record-model Specification areas from the current DRMCP baseline.

## Work

Write only:

- `drcli/records/spec/design-records-cli/artifacts/**`
- `drcli/records/spec/design-records-cli/current-record-model/**`

Use `drmcp/records/spec/design-records-mcp/artifacts/**` and `current-record-model/**` as the primary baseline.

Rules:

- preserve PRODUCT-owned identity, source, metadata, structure, admission, conflict, and addressability semantics;
- change DRMCP/MCP transport wording only where DRCLI ownership requires it;
- use `spec:drcli.design_records_cli...` refs for DRCLI-owned contracts;
- keep PRODUCT semantic refs as external authority;
- do not write operation, diagnostic, CLI, guidance, authoring, or root-index files;
- stop if adapting the baseline would require a new semantic choice beyond DRCLI-REQ-CLI-001.

## Done condition

- Both assigned Specification areas are complete and internally navigable.
- No MCP transport dependency remains in DRCLI-owned normative behavior.
- No PRODUCT semantics are duplicated as independent DRCLI authority.
- No file outside the assigned path families changed.

## Verification

- Compare the resulting file set against the DRMCP baseline.
- Check refs, parents, H1-adjacent metadata, and topic indexes.
- Search the assigned output for unintended `spec:drmcp.design_records_mcp` ownership refs and classify intentional external evidence separately.

## Evidence

- Inputs: `DRCLI-REQ-CLI-001`, completed coordination Task `DRCLI-TASK-CLI-002-01`, and the current DRMCP `artifacts/**` plus `current-record-model/**` Specification baselines.
- Created 42 DRCLI Specification files across the two assigned path families; the relative file set matches the DRMCP baseline 42-for-42.
- Replaced DRMCP-owned self references with `spec:drcli.design_records_cli...` refs while preserving PRODUCT `spec:product...` authority refs and baseline USDM evidence refs.
- Applied current Brewprint Specification shape without changing product semantics: artifact topic entry files use `Index`; the shared template area is an `Overview`; required Contract/format H2 sections are present; stale baseline parent/topic metadata was normalized for internal navigation.
- Verification found no `spec:drmcp.design_records_mcp`, `DRMCP`, or normative `MCP` text in the assigned outputs. Eleven remaining `mcp` strings occur only inside retained `usdm:drmcp.mcp_capabilities...` H1 metadata evidence refs in four current-record-model files.
- All 42 files have path-matching DRCLI spec IDs, H1-adjacent `status`, `date`, and `parent` metadata, and required document-shape sections for their declared kind.
- Parent/topic navigation checks passed for all children inside the assigned areas. One forward DRCLI ref to `spec:drcli.design_records_cli.operations.discovery_and_listing.sequential_record_listing` remains intentionally unresolved until sibling Task T04 authors its owned operation area.
- PRODUCT spec-ref sets and USDM ref sets were preserved from the DRMCP baseline for every corresponding file.
- No DRMCP file, sibling lane output, root DRCLI spec index, or parent Work Item lifecycle was modified.
- No independent review occurred; integrated review remains owned by `DRCLI-TASK-CLI-002-10`.
