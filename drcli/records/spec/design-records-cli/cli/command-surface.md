# Contract: DRCLI command surface

- **id**: `spec:drcli.design_records_cli.cli.command_surface`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.cli`
- **contract_class**: `interface`

## What this is

Defines the logical DRCLI executable, command hierarchy, and exact mapping from CLI commands to semantic operation contracts.
The mapped operation Specifications remain authoritative for request, result, and Design Records semantics.

## Non-goals

- Redefining any semantic operation owned under `spec:drcli.design_records_cli.operations`.
- Defining operation-specific option names that were not fixed by the accepted CLI decision ledger.
- Choosing parser, framework, language, or implementation technology.

## Request

The logical executable name is `drcli`.
The root command contains exactly these command groups:

- `scope`
- `record`
- `tree`
- `reference`
- `validate`
- `guide`
- `proposal`

The following command paths map to the stated semantic operation boundary.

| command | semantic operation or action | authoritative operation spec |
|---|---|---|
| `drcli scope list` | record-scope discovery | `spec:drcli.design_records_cli.operations.discovery_and_listing.record_scope_discovery` |
| `drcli record list` | sequential-record listing | `spec:drcli.design_records_cli.operations.discovery_and_listing.sequential_record_listing` |
| `drcli tree children` | direct tree-child listing | `spec:drcli.design_records_cli.operations.discovery_and_listing.tree_child_listing` |
| `drcli tree inspect` | bounded tree overview | `spec:drcli.design_records_cli.operations.discovery_and_listing.tree_overview` |
| `drcli record get` | exact record retrieval | `spec:drcli.design_records_cli.operations.retrieval.exact_record_retrieval` |
| `drcli record sections` | exact H2 section retrieval | `spec:drcli.design_records_cli.operations.retrieval.h2_section_retrieval` |
| `drcli record search` | lexical record search | `spec:drcli.design_records_cli.operations.search.lexical_record_search` |
| `drcli reference resolve` | canonical reference resolution | `spec:drcli.design_records_cli.operations.reference_resolution` |
| `drcli validate scope` | broad-scope validation | `spec:drcli.design_records_cli.operations.validation.scope_validation` |
| `drcli validate records` | exact-record validation | `spec:drcli.design_records_cli.operations.validation.exact_record_validation` |
| `drcli guide list` | authoring-guidance list | `spec:drcli.design_records_cli.operations.guidance.list` |
| `drcli guide get` | authoring-guidance exact read | `spec:drcli.design_records_cli.operations.guidance.exact_read` |
| `drcli record create` | normal create authoring | `spec:drcli.design_records_cli.operations.authoring.create` |
| `drcli record update` | normal partial-update authoring | `spec:drcli.design_records_cli.operations.authoring.update` |
| `drcli proposal create` | deferred create preparation | `spec:drcli.design_records_cli.operations.authoring.deferred_proposal` |
| `drcli proposal update` | deferred update preparation | `spec:drcli.design_records_cli.operations.authoring.deferred_proposal` |
| `drcli proposal get` | retained-proposal inspection | `spec:drcli.design_records_cli.operations.authoring.deferred_proposal` |
| `drcli proposal accept` | retained-proposal acceptance | `spec:drcli.design_records_cli.operations.authoring.deferred_proposal` |
| `drcli proposal discard` | retained-proposal discard | `spec:drcli.design_records_cli.operations.authoring.deferred_proposal` |

Command-specific help owns discovery of the concrete operands and options accepted by each command.
Sibling CLI contracts own repository selection, body-source flags, output mode, and usage handling.

## Response

A valid command dispatches exactly one mapped semantic operation or retained-proposal action.
CLI dispatch does not alter the mapped operation's semantic request or result meaning.

## Errors

| condition | behavior |
|---|---|
| Unknown command or command group | Use the CLI usage-diagnostic contract; do not guess another command. |
| Known command with invalid CLI shape | Use the CLI usage-diagnostic contract; do not invoke the semantic operation. |
| Valid command whose semantic request fails | Preserve the mapped operation result or failure and apply the output/status contract. |

## Boundary

This contract owns command names and command-to-operation dispatch only.
Semantic selectors, result fields, validation rules, authoring rules, and diagnostic codes remain owned by their operation or PRODUCT Specifications.

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.cli.help_and_usage_diagnostics` | Owns command discovery and CLI usage failures. |
| `spec:drcli.design_records_cli.cli.repository_selection_and_traversal` | Owns repository and traversal options. |
| `spec:drcli.design_records_cli.cli.authoring_body_input` | Owns body-source flags for body-capable authoring commands. |
| `spec:drcli.design_records_cli.cli.output_and_exit_status` | Owns result rendering, streams, and process status. |