# Contract: DRCLI repository selection and traversal

- **id**: `spec:drcli.design_records_cli.cli.repository_selection_and_traversal`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.cli`
- **contract_class**: `interface`

## What this is

Defines per-invocation repository selection and traversal eligibility used by automatic current-source discovery.

## Non-goals

- Defining PRODUCT-owned artifact admission, identity, or source semantics.
- Creating a persistent current-repository setting.
- Selecting ignore-library or filesystem-walking implementation technology.

## Request

Every repository-dependent invocation selects exactly one repository root.

| input state | selected repository |
|---|---|
| `--repo <path>` is absent | The process current working directory exactly. |
| `--repo <path>` is present | The supplied path replaces the current-working-directory default. |

DRCLI never searches parent directories to locate a repository.
The installation location never selects the repository.
A later invocation performs repository selection again and does not inherit a prior selection.

Repository-dependent semantic operations that construct current-record state traverse only beneath the selected repository.
Such commands honor Git-compatible ignore rules by default.
Ignore pruning occurs before candidate records-root discovery, namespace-evidence collection, and source-corpus construction.
A pruned directory or file is not traversal-eligible.

Such commands accept `--no-ignore` as the explicit full-traversal override.
`--no-ignore` disables ignore filtering only.
`--no-ignore` does not broaden PRODUCT-owned artifact, current-source, identity, or admission semantics.
A fully ignored candidate records root is absent from discovery under default traversal and produces no current-source discovery diagnostic.

Command-specific help identifies whether `--no-ignore` applies to that command.
Help invocations do not require repository selection or repository access.

## Response

The selected repository root is the sole repository context supplied to the mapped repository-dependent operation.
Ignore filtering affects traversal eligibility only and never changes canonical Design Records semantics.

Current-state construction consumes the traversal-eligible repository view.
It automatically discovers candidate `records` roots and resolves their app namespaces under `spec:drcli.design_records_cli.current_record_model.current_source_corpus`.
The CLI supplies no preconfigured app-namespace and records-root association.
DRCLI requires no repository manifest, persistent registration, per-repository installation, or app-directory naming convention.

## Errors

| condition | classification |
|---|---|
| `--repo` lacks its required path value | CLI `usage_error`. |
| A selected repository cannot be accessed or used for the requested operation | `execution_failure`. |
| Current-source construction finds an unresolvable, ambiguous, or duplicate-app candidate records root | Request-wide semantic `configuration_failure` owned by `spec:drcli.design_records_cli.current_record_model.current_source_corpus`. |
| `--no-ignore` is supplied to a command whose help does not expose that option | CLI `usage_error`. |

Usage failures follow `spec:drcli.design_records_cli.cli.help_and_usage_diagnostics`.
Execution failures follow `spec:drcli.design_records_cli.cli.output_and_exit_status`.

## Boundary

This contract owns repository-root selection and traversal eligibility, including ignore-filter override behavior.
Candidate records-root discovery, namespace resolution, current-source failure triggers, and artifact admission remain owned by `spec:drcli.design_records_cli.current_record_model` and related artifact contracts.

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.current_record_model.current_source_corpus` | Owns candidate records-root discovery, namespace resolution, and current-source construction. |
| `spec:drcli.design_records_cli.current_record_model` | Owns current-record admission and addressability semantics after current-source construction. |
| `spec:drcli.design_records_cli.cli.help_and_usage_diagnostics` | Owns option discoverability and usage failures. |
| `spec:drcli.design_records_cli.cli.runtime_and_distribution` | Keeps installation and retained state independent from repository selection. |