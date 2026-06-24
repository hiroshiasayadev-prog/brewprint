# Reference: Artifact ID grammar

- **id**: `spec:product.design_records.namespace_model.artifact_id_grammar`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.design_records.namespace_model`

## What this is

Defines the canonical artifact ID grammar for new sequential records — ADR, investigation, requirement, work item, and task: public ID format, sequence format, allocation scopes, subdomain exclusion rule, and canonical-ref policy. Spec identity is governed by `spec:product.design_records.spec_format.spec_id_as_ref`.

## Grammar

### REQ / WORK / INV / ADR

```
<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<SEQUENCE>
```

Examples: `DRMCP-REQ-MCP-033`, `BPDSL-WORK-DATA-016`, `PRODUCT-REQ-NAMESPACE-003`, `DRMCP-ADR-MCP-001`

> Subdomains are not segments in the artifact ID. They are expressed as a `subdomain` field in record metadata. See `spec:product.design_records.namespace_model.subdomain_model`.

> New ADRs use the grammar above. Existing issued ADR IDs are governed by `spec:product.brewprint.compatibility.legacy_id_compatibility`.

### TASK

```
<APP_NAMESPACE>-TASK-<DOMAIN_NAMESPACE>-<WORK_SEQUENCE>-<TASK_SEQUENCE>
```

Examples: `BPDSL-TASK-DATA-016-01`, `PRODUCT-TASK-NAMESPACE-002-01`

## Sequence format

| artifact | sequence format | example |
|---|---|---|
| REQ / WORK / INV / ADR | 3-digit zero-padded | `001`, `016`, `099` |
| TASK WORK_SEQUENCE | Inherits the parent WORK's 3-digit number | `016` |
| TASK TASK_SEQUENCE | 2-digit zero-padded | `01`, `12` |

For REQ, WORK, INV, and ADR, sequence allocation is scoped by:

```
app namespace + artifact kind + domain namespace
```

TASK sequence allocation is scoped by its parent Work Item. The WORK_SEQUENCE segment inherits the parent Work Item's three-digit sequence number.

## Canonical reference

The complete public ID — `<APP>-<KIND>-<DOMAIN>-<SEQ>` for REQ/WORK/INV/ADR, or `<APP>-TASK-<DOMAIN>-<WS>-<TS>` for TASK — is the canonical record ID-as-ref.

Bare forms (`REQ-*`, `WORK-*`, `TASK-*`, etc.) are internal grammar fragments used in spec text for brevity. They are not valid canonical external references.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.namespace_model` | Parent overview. |
| `spec:product.brewprint.compatibility.legacy_id_compatibility` | Compatibility policy for issued V01-* legacy IDs. |
| `spec:product.design_records.namespace_model.subdomain_model` | Subdomain grouping; subdomains are not ID segments. |
| `spec:product.brewprint.compatibility.existing_artifacts` | Attribution policy for existing artifacts. |
| `spec:product.design_records.traceability.artifact_refs` | How canonical IDs function as semantic refs. |
