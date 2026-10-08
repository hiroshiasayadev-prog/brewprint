# Contract: DRCLI authoring body input

- **id**: `spec:drcli.design_records_cli.cli.authoring_body_input`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.cli`
- **contract_class**: `interface`

## What this is

Defines the CLI mapping for explicit Markdown body sources used by body-capable create, update, and deferred-proposal preparation commands.
Semantic body requirements remain owned by the authoring operation Specifications.

## Non-goals

- Defining when PRODUCT authoring semantics require body content.
- Defining body-cache retention, storage, or validity policy.
- Inferring body input from shell redirection or TTY state.

## Request

A body-capable authoring command accepts exactly these explicit body-source forms:

| CLI input | body source supplied to the semantic operation |
|---|---|
| `--body <markdown>` | The supplied inline Markdown body. |
| `--body-stdin` | The Markdown body read from standard input. |
| `--body-cache-id <id>` | The preserved body resolved from the opaque cache ID. |

The three body-source forms are mutually exclusive for one body input.
DRCLI reads standard input as body content only when `--body-stdin` is present.
Piped or redirected standard input does not implicitly select body mode.

Supplying no body source is valid only when the selected semantic authoring operation does not require body content.
A body cache ID supplies body content only.
A body cache ID does not supply repository selection, target identity, metadata fields, section selectors, or completion disposition.

The body-source mapping applies to `record create` and `record update` where their operation contract consumes body content.
The same mapping applies to `proposal create` and `proposal update` because they prepare the corresponding deferred authoring candidate.

## Response

A valid body-source selection yields one body value for the mapped semantic authoring request.
Inline and standard-input bodies are equivalent after DRCLI receives the body bytes.
A valid body cache ID yields the preserved body without changing any other CLI or semantic input.

## Errors

| condition | CLI behavior |
|---|---|
| More than one body-source form is supplied for one body input | `usage_error`; name the conflicting forms, show the three accepted alternatives, and give scoped help. |
| Body content is required by the selected operation but no body source is supplied | `usage_error`; identify the missing body input, show the three accepted forms, and give scoped help. |
| A body-source option is supplied where the selected command does not consume body content | `usage_error`; identify the unsupported option and show the valid option surface. |
| `--body` or `--body-cache-id` lacks its required value | `usage_error`; identify the missing value and show the accepted invocation shape. |

A syntactically valid but unknown, expired, or unusable body cache ID is not repaired as CLI usage.
Its outcome follows the authoring retry-cache and runtime contracts.

## Boundary

This contract owns only CLI body-source selection and mapping.
Create, update, retry-cache, and deferred-proposal Specifications own semantic body use and retained-state outcomes.

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.operations.authoring.create` | Owns create body requirements. |
| `spec:drcli.design_records_cli.operations.authoring.update` | Owns update body requirements. |
| `spec:drcli.design_records_cli.operations.authoring.retry_cache` | Owns opaque preserved-body reuse. |
| `spec:drcli.design_records_cli.operations.authoring.deferred_proposal` | Owns deferred authoring preparation. |
| `spec:drcli.design_records_cli.cli.help_and_usage_diagnostics` | Owns actionable invalid-usage rendering. |