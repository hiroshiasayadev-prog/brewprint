# Contract: Validation policy

- **id**: `spec:product.concepts.spec_format.validation_policy`
- **status**: accepted
- **date**: 2026-06-11
- **parent**: `spec:product.concepts.spec_format`
- **contract_class**: `format`

## What this is

This spec defines initial validation severity and ownership for the spec-format contract. It also preserves the parser-aware rule that YAML front matter and fenced code blocks are ignored when counting real headings.

This policy identifies durable DRMCP implementation targets and the temporary PRODUCT-level validation bridge. It does not require patching the current DRMCP implementation.

## Current contract

Validation diagnostics use severity appropriate to the migration phase and owning follow-up work. New or migrated specs are held to the accepted format; existing specs may receive warnings until migration work explicitly updates them.

This policy identifies durable DRMCP implementation targets and the temporary PRODUCT-level validation bridge. Current DRMCP patching remains out of scope.

## Rules

Validation must be parser-aware: YAML front matter and fenced code blocks are ignored when counting real headings.

The DRMCP-owned rows below identify the durable implementation target for a later DRMCP redesign/reimplementation phase. Before that phase, PRODUCT-WORK-SPEC-006 may provide temporary PRODUCT-level validation tooling for migration; that temporary tooling must not patch the current DRMCP implementation.

## Validation rules

| validation rule | initial severity | owner |
|---|---|---|
| missing real ATX H1 | error | DRMCP-WORK-SPEC-001 |
| multiple real ATX H1 headings | error | DRMCP-WORK-SPEC-001 |
| H1 does not match `# <SpecKind>: <Title>` | error for migrated/new specs; warning during inventory | DRMCP-WORK-SPEC-001 |
| missing H1-adjacent `id` / `status` / `date` / `parent` | error for new/migrated specs; warning during inventory | DRMCP-WORK-SPEC-001 |
| missing or invalid H1-adjacent `contract_class` on a `Contract` spec | error for new/migrated specs; warning during inventory | DRMCP-WORK-SPEC-001 |
| H1-adjacent `id` does not match path-derived canonical spec ref | error for new/migrated specs; warning during inventory, migration, or transient working states | DRMCP-WORK-SPEC-001 |
| any YAML front matter | warning for existing specs; error for new/migrated specs under this format | DRMCP-WORK-SPEC-001 |
| hidden front matter `depends_on` / source refs | warning for existing specs; error for new/migrated specs under this format | DRMCP-WORK-SPEC-001 |
| hidden front matter topic refs | warning for existing specs; error for new/migrated specs under this format | DRMCP-WORK-SPEC-001 |
| front matter `design_record.kind` | warning for existing specs; error for new/migrated specs under this format | DRMCP-WORK-SPEC-001 |
| missing `## What this is` | warning for existing specs; error for new/migrated specs | DRMCP-WORK-SPEC-001 |
| invalid accepted spec kind | error | DRMCP-WORK-SPEC-001 |
| missing required section by kind | warning or error based on migration phase | DRMCP-WORK-SPEC-001 |
| invalid `## Topics` table columns, including canonical `file` instead of `ref` | error for migrated/new Index or Overview+Topics specs | DRMCP-WORK-SPEC-002 |
| unresolved child `ref` | error for migrated/new Index or Overview+Topics specs | DRMCP-WORK-SPEC-002 |
| duplicate parent declaration | error | DRMCP-WORK-SPEC-002 |
| parent grammar violation | error | DRMCP-WORK-SPEC-002 |
| topic cycle | deferred until graph validation contract | DRMCP-WORK-SPEC-002 |

## Errors

| condition | severity |
|---|---|
| New or migrated spec violates a rule marked error | Error. |
| Existing unmigrated spec violates a migration-warning rule | Warning until PRODUCT-WORK-SPEC-005 migrates it. |
| Inventory, migration, or transient working state has a path-derived canonical ref mismatch | Warning until the migration step explicitly declares the spec migrated. |
| Temporary PRODUCT tooling patches current DRMCP implementation | Out of scope for PRODUCT-WORK-SPEC-006 and this contract. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.concepts.spec_format` | Parent Index for this contract. |
| `spec:product.concepts.spec_format.document_shape` | Defines H1, metadata, and required-section checks. |
| `spec:product.concepts.spec_format.topics_table` | Defines topic table checks. |
| `spec:product.concepts.spec_format.spec_id_as_ref` | Defines ID derivation and parent grammar checks. |
