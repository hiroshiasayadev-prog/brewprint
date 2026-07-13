# Reference: Record discovery paths

- **id**: `spec:product.design_records.repository_layout.record_discovery_paths`
- **status**: draft
- **date**: 2026-07-13
- **parent**: `spec:product.design_records.repository_layout`

## What this is

Defines app-independent path-pattern conventions for locating record files within a `records_root` by record kind.
It does not define DRMCP configuration representation, parser behavior, admission results, or diagnostics.

## Current contract

| kind | discovery path pattern | standard file-name form |
|---|---|---|
| `decision` | `<records_root>/adr/*/*.md` | `<APP>-ADR-<DOMAIN>-<SEQUENCE>-<slug>.md` |
| `spec` | `<records_root>/spec/**/*.md` | Path-derived Specification rules apply. |
| `investigation` | `<records_root>/investigations/*/*.md` | `<APP>-INV-<DOMAIN>-<SEQUENCE>-<slug>.md` |
| `requirement` | `<records_root>/requirements/*/*.md` | `<APP>-REQ-<DOMAIN>-<SEQUENCE>-<slug>.md` |
| `work_item` | `<records_root>/work-items/*/*.md` | `<APP>-WORK-<DOMAIN>-<SEQUENCE>-<slug>.md` |
| `task` | `<records_root>/tasks/*/*.md` | `<APP>-TASK-<DOMAIN>-<WORK_SEQUENCE>-<TASK_SEQUENCE>-<slug>.md` |

Sequential ADR sources use the domain-subdirectory pattern.
Flat ADR sources directly under `<records_root>/adr/` are outside the current sequential discovery contract.

Sequential discovery uses physical artifact-kind and domain placement plus the `.md` extension.
It does not filter a source by the public-ID text in its file name.

## Rules

- Path patterns describe repository placement, not tool indexing behavior.
- Path patterns use a `records_root` supplied by the caller or implementation context.
- Spec discovery uses the topic tree under `<records_root>/spec/`.
- Sequential record discovery uses the kind and domain subdirectories defined by the repository layout model.
- A sequential file-name public-ID prefix must match the H1 public ID as a repository-conformance rule.
- A file-name mismatch does not exclude the source from discovery and does not supply or replace canonical identity.
- Tool-specific inclusion filters are outside this contract, but they must not use sequential file-name identity text to exclude an otherwise in-scope Markdown source.
- App namespace is supplied by implementation context and is not derived from the physical `records_root` path.

## DRMCP boundary

The following implementation-specific concerns remain outside PRODUCT normative text.

| implementation-specific concern | app-local owner |
|---|---|
| Configured app namespace and records-root association. | `spec:drmcp.design_records_mcp.current_record_model.current_source_corpus`. |
| Artifact source-root, sequential domain-depth, and current-source corpus construction. | `spec:drmcp.design_records_mcp.current_record_model.current_source_corpus`. |
| H1 identity agreement, candidate formation, and nonidentity conformance boundary. | `spec:drmcp.design_records_mcp.current_record_model.artifact_candidate_and_admission`. |
| Concrete validation findings and diagnostics. | DRMCP validation Specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.repository_layout` | Parent repository-layout overview. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Sequential H1 identity and physical-context agreement. |

## Sources

- V01-ADR-076 section bootstrap policy.
- V01-ADR-092 section 1.
- PRODUCT-ADR-SPEC-019 establishes H1-only sequential identity and file-name conformance projection.
