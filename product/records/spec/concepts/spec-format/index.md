# Index: Spec format

- **id**: `spec:product.concepts.spec_format`
- **status**: accepted
- **date**: 2026-06-11
- **parent**: `root`

## What this is

This Index is the navigation-first entry point for the PRODUCT-level spec format contract. It points to focused child specs that define document shape, topic tables, spec ID-as-ref behavior, validation policy, and follow-up ownership boundaries.

The detailed contract body intentionally lives in child specs. This file owns topic organization for `spec:product.concepts.spec_format` and should stay small enough for agents and humans to choose the right child spec before reviewing implementation or migration work.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Spec format overview | Overview | `spec:product.concepts.spec_format.overview` | Purpose, front matter policy summary, scope, traceability policy, and non-goals. |
| Spec document shape | Contract | `spec:product.concepts.spec_format.document_shape` | Accepted spec kinds, H1 format, H1-adjacent metadata, contract classes, and required section matrix. |
| Topics table | Contract | `spec:product.concepts.spec_format.topics_table` | `## Topics` table columns, ref-first child targets, authoritative parent declarations, and duplicate-parent rules. |
| Spec ID-as-ref | Concept | `spec:product.concepts.spec_format.spec_id_as_ref` | Path-derived canonical spec refs, parent reference grammar, and examples. |
| Validation policy | Contract | `spec:product.concepts.spec_format.validation_policy` | Parser-aware validation, migration warning/error policy, and temporary tooling boundary. |
| Follow-up boundary | Concept | `spec:product.concepts.spec_format.follow_up_boundary` | Follow-up ownership across PRODUCT and DRMCP work, including traceability and non-scope boundaries. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.concepts.namespace_model` | Defines app and domain namespace model used by path-derived spec IDs. |
| `spec:product.concepts.repository_layout` | Defines namespace-first repository layout used by the format. |
| `spec:drmcp.design_records_mcp.overview` | Defines DRMCP record discovery and validation scope that later implementation work must align with. |
