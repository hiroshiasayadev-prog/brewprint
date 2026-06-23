# Reference: v2 artifact ID grammar

- **id**: `spec:product.concepts.namespace_model.v2_grammar`
- **status**: draft
- **date**: 2026-06-22
- **parent**: `spec:product.concepts.namespace_model`

## What this is

Defines the v2 artifact ID grammar, sequence format, and mapping rule from existing domain-first IDs to namespace-aware IDs.

## Grammar

The ID format for new artifacts is as follows.

**REQ / WORK / INV / ADR:**

```
<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<SEQUENCE>
```

Examples: `DRMCP-REQ-MCP-033`, `BPDSL-WORK-DATA-016`, `PRODUCT-REQ-NAMESPACE-003`, `DRMCP-ADR-MCP-001`

> **Note**: The domain namespace tokens shown here (`MCP`, `DATA`, `NAMESPACE` etc.) are examples. Subdomains are not included as segments in the v2 ID grammar. For subdomain grouping definitions, see `spec:product.concepts.namespace_model.subdomain_model`.

> **ADR note**: Existing ADRs retain `V01-ADR-NNN` per V01-ADR-099. The prior decision to maintain a global `ADR-NNN` sequence is superseded by V01-ADR-099, which adopts unified v2 grammar across all artifact kinds (including ADR) to remove form-branching logic in UI / MCP.

**TASK:**

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

```text
app namespace + artifact kind + domain namespace
```

TASK sequence allocation is scoped by its parent Work Item.

## Mapping rule from existing IDs

Logical mapping rules from existing domain-first IDs to v2 namespace-aware IDs.

Per V01-ADR-096, existing artifact IDs are not changed. These rules are used as reference rules for grouping, display, and attribution resolution in UI / MCP.

| existing domain prefix | app namespace | mapping rule | example |
|---|---|---|---|
| `MCP` | `DRMCP` | Prepend `DRMCP-` | `V01-REQ-MCP-013` → `DRMCP-REQ-MCP-013` |
| `DATA` | `BPDSL` | Prepend `BPDSL-` | `V01-WORK-DATA-009` → `BPDSL-WORK-DATA-009` |
| `RESOLVE` | `BPDSL` | Prepend `BPDSL-` | `V01-REQ-RESOLVE-001` → `BPDSL-REQ-RESOLVE-001` |

### PRODUCT prefix handling

Current `REQ-PRODUCT-NNN` / `WORK-PRODUCT-NNN` / `TASK-PRODUCT-NNN-NN` are a transitional format that uses `PRODUCT` as a domain identifier.

In the full v2 form, the domain namespace is explicit:

| current form | v2 full form | notes |
|---|---|---|
| `V01-REQ-PRODUCT-001` | `PRODUCT-REQ-NAMESPACE-001` | Namespace model requirement |
| `V01-WORK-PRODUCT-001` | `PRODUCT-WORK-NAMESPACE-001` | Namespace model work |
| `V01-TASK-PRODUCT-001-01` | `PRODUCT-TASK-NAMESPACE-001-01` | Namespace model task |

Existing `REQ-PRODUCT-*` / `WORK-PRODUCT-*` / `TASK-PRODUCT-*` are not migrated. The decision to migrate new PRODUCT-level artifacts to full v2 form will be made once a machine-readable namespace registry is in place and MCP can natively resolve v2 IDs. Until then, the current form (using `PRODUCT` as a domain prefix) continues.
