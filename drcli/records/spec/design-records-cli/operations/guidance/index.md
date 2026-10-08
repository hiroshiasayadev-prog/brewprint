# Overview: Authoring guidance operations

- **id**: `spec:drcli.design_records_cli.operations.guidance`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations`

## What this is

Defines DRCLI operations that expose the portable Design Records authoring standards as guidance without creating a separate guidance model.

## Current contract

| concern | contract |
|---|---|
| Source | The portable `design_records` standards package supplied with DRCLI. |
| Semantic authority | PRODUCT-owned Design Records Specs remain authoritative. The package is their copied and canonical-ref-rewritten distribution. |
| Record treatment | Guidance entries are normal current Specs from the package. |
| Guidance root | `spec:design_records.authoring_standards`. |
| List scope | Child Specs under `spec:design_records.authoring_standards.*`; the root is excluded. |
| Detail scope | One exact canonical child ref under `spec:design_records.authoring_standards.*`. |
| Lookup behavior | Exact canonical refs only. No basename, filename, path, title, alias, fuzzy, or inferred lookup. |
| Legacy model | No legacy guide directory, guide-specific index, or separate semantic model is authoritative. |

The guidance operations do not define how CLI commands or options are spelled.
CLI mapping is owned by the DRCLI CLI contract.

## Boundary

DRCLI owns guidance projection and operation outcomes.

PRODUCT owns the authoring-standard semantics and portable package content.
The package carries those semantics under the fixed `design_records` namespace.

This area does not define:

- package production;
- package-root resolution;
- current-record discovery outside the guidance scope;
- CLI command names or option names;
- diagnostic serialization;
- authoring transaction behavior.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Guidance list operation | Contract | `spec:drcli.design_records_cli.operations.guidance.list` | Lists current authoring-standard child Specs as canonical ID, title, and abstract projections. |
| Guidance exact-read operation | Contract | `spec:drcli.design_records_cli.operations.guidance.exact_read` | Reads one exact authoring-standard child Spec and returns its complete Markdown source verbatim. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.authoring_standards` | PRODUCT source authority for authoring-standard semantics before package ref rewriting. |
