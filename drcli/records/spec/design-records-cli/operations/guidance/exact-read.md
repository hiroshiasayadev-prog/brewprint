# Contract: Guidance exact-read operation

- **id**: `spec:drcli.design_records_cli.operations.guidance.exact_read`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations.guidance`
- **contract_class**: `interface`

## What this is

Reads one exact authoring-standard child Spec from the portable `design_records` package and returns its complete Markdown source.

## Request

The operation requires one semantic input:

| field | required | type | meaning |
|---|---:|---|---|
| `id` | yes | string | Exact canonical package Spec ref inside the guidance child subtree. |

Accepted IDs are exact refs matching:

```text
spec:design_records.authoring_standards.*
```

The root `spec:design_records.authoring_standards` is not a guidance detail target.

The supplied `id` is evaluated exactly as provided.

The operation performs no:

- whitespace trimming;
- case normalization;
- basename lookup;
- filename-stem lookup;
- physical-path lookup;
- title lookup;
- alias lookup;
- fuzzy lookup;
- inferred candidate lookup;
- reference-resolution fallback.

Concrete CLI argument spelling is outside this contract.

## Response

A successful result contains:

| field | required | meaning |
|---|---:|---|
| `id` | yes | Exact canonical package Spec ref. |
| `title` | yes | First H1 text from the current package Spec source. |
| `content` | yes | Complete Markdown source, verbatim and untruncated. |

The operation does not format, summarize, normalize, or truncate `content`.

Normal guidance output does not add physical source paths, parsed metadata fields, or internal index state.
Metadata already present inside the Markdown remains present because `content` is verbatim.

## Errors

| condition | required outcome |
|---|---|
| The exact canonical ref is absent. | Report guidance not found. |
| The supplied ref is outside the accepted guidance child subtree. | Report guidance not found. |
| The in-scope canonical identity is conflicted. | Report guidance unavailable. Do not select a duplicate winner. |
| The in-scope source is unreadable. | Report guidance unavailable. |
| The in-scope source cannot supply its first H1 or complete Markdown content. | Report guidance unavailable. |
| The semantic `id` input is missing or not a string. | Report invalid operation input. |

Concrete diagnostic codes and CLI rendering are outside this contract.

## Boundary

The operation returns the current package Spec source.
It does not reinterpret, repair, or independently own the guide semantics.

The operation does not consult a legacy guide directory or a guide-specific lookup map.

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.operations.guidance` | Parent guidance-operation boundary. |
| `spec:product.design_records.authoring_standards` | PRODUCT source authority for the returned authoring standards. |
