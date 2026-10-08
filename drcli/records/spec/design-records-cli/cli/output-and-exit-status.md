# Contract: DRCLI output and exit status

- **id**: `spec:drcli.design_records_cli.cli.output_and_exit_status`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.cli`
- **contract_class**: `interface`

## What this is

Defines explicit output modes, stdout and stderr responsibility, JSON stream discipline, and stable numeric process statuses.

## Non-goals

- Redefining semantic result fields or operation diagnostic objects.
- Selecting terminal styling, pager behavior, logging, or serialization implementation technology.
- Switching modes implicitly from TTY detection.

## Request

Operation output defaults to `--format human`.
`--format json` explicitly selects machine-readable operation output.
DRCLI does not change output mode based on whether a stream is attached to a TTY.
Explicit help always uses human-readable help text.

## Response

| outcome | stdout | stderr |
|---|---|---|
| Trustworthy successful operation result | Primary result. | No failure diagnostic. |
| Trustworthy negative or partial semantic result | Primary result. | No separate failure document required by this contract. |
| Explicit help | Human-readable help text. | No failure diagnostic. |
| CLI usage failure | No primary operation result. | One actionable usage diagnostic. |
| Failure that prevents a trustworthy operation result | No primary operation result. | One execution or internal failure diagnostic. |

A trustworthy result remains on stdout even when its semantic outcome makes the process status nonzero.
Warnings and advisories that accompany a trustworthy result remain part of that result contract.

In JSON mode, stdout contains exactly one machine-readable result document when a trustworthy operation result exists.
In JSON mode, usage, execution, and internal failures emit exactly one machine-readable diagnostic document on stderr.
JSON result streams contain no additional prose.
Machine-readable usage diagnostics retain the actionable information required by the help and usage contract.

For authoring create, update, shared write completion, retry cache, and deferred-proposal actions, every negative outcome is mapped by the canonical authoring codes and reason/guard discriminators in `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_catalog`. A trustworthy semantic rejection (including candidate PRODUCT-validation findings, collision, stale target, mismatch, or known discarded/accepted proposal state) uses one top-level `error` result document on stdout in JSON mode and status `1`. Unavailable retained cache/proposal state, package/repository resources, preparation or persistence I/O, and write-completion failure without a trustworthy normal result use one top-level `error` diagnostic document on stderr and status `3`. The known identity and lifecycle state of an invalid proposal is not a successful write. Unexpected invariant/software faults use status `4`; syntactic CLI usage violations use the actionable usage document on stderr and status `2`. The optional nested `cache_error` or returned `body_cache_id` remains in the same authoring error's context and never adds a second stdout/stderr document.

A failure after modification must accurately communicate observed `write_state` in the operational failure context and must not be reported as a verified non-write. Successful direct writes, valid update no-ops, successful deferred preparation, inspection, discard, and acceptance retain their existing status `0` and normal result contracts. PRODUCT-owned validation finding `code` and `contract_ref` remain unchanged beneath `authoring_candidate_validation_failed.context.affected_findings`.

Numeric exit statuses are stable across supported operating systems.

| status | category | meaning |
|---:|---|---|
| `0` | success | The invocation succeeded. Warnings or advisories alone do not change this status. |
| `1` | contract_result_failure | A valid invocation produced a trustworthy but unsatisfied semantic result. Validation error findings and partial selector failure are included. |
| `2` | usage_error | CLI command, option, required input, exclusivity, or selector-shape usage is invalid. |
| `3` | execution_failure | Repository, bundled package, retained state, I/O, permission, or another operational condition prevents a trustworthy result. |
| `4` | internal_failure | An unexpected software or invariant failure prevents a trustworthy result. |

## Errors

| condition | required handling |
|---|---|
| Human or JSON mode cannot produce a trustworthy operation result because required runtime resources are unavailable | stderr and status `3`. |
| An unexpected invariant or software failure occurs | stderr and status `4`. |
| A valid semantic operation returns an unsatisfied but trustworthy result | stdout and status `1`. |
| Only warnings or advisories are present | Preserve stdout result and status `0`. |

CLI usage failures are rendered under `spec:drcli.design_records_cli.cli.help_and_usage_diagnostics` and use status `2`.
Semantic operation Specifications decide whether a valid result is satisfied, partial, or negative.

## Boundary

This contract owns output-mode selection, process streams, JSON single-document framing, and numeric status categories.
Mapped operation Specifications own normal result structure and semantic success criteria.
`spec:drcli.design_records_cli.diagnostics` owns semantic operation diagnostic shape before CLI serialization.

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.cli.help_and_usage_diagnostics` | Defines actionable usage-diagnostic content. |
| `spec:drcli.design_records_cli.diagnostics` | Owns transport-independent semantic operation diagnostics. |
| `spec:drcli.design_records_cli.operations.validation` | Supplies validation outcomes that can yield status `1`. |