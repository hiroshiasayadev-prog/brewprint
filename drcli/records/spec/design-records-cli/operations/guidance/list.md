# Contract: Guidance list operation

- **id**: `spec:drcli.design_records_cli.operations.guidance.list`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations.guidance`
- **contract_class**: `interface`

## What this is

Lists authoring-standard child Specs from the portable `design_records` package as a compact guidance projection.

## Request

The operation has no guidance-specific selector input.

The operation always applies this fixed scope:

| scope | value |
|---|---|
| app namespace | `design_records` |
| record kind | `spec` |
| included canonical refs | `spec:design_records.authoring_standards.*` |
| excluded root | `spec:design_records.authoring_standards` |

Repository selection, output formatting, and concrete CLI syntax are outside this contract.

## Response

The result is an ordered sequence of guidance entries.

Each entry contains:

| field | required | meaning |
|---|---:|---|
| `id` | yes | Exact canonical package Spec ref. |
| `title` | yes | First H1 text from the current package Spec source. |
| `abstract` | yes | Exact body of the `## What this is` section from the current package Spec source. |

Entries are ordered by canonical `id` in ASCII lexical order.

Every addressable current Spec in the fixed child subtree is a catalog candidate.
A candidate is projectable only when its canonical identity, first H1, `## What this is` body, and complete source read are available.

A valid operation with no in-scope child Specs returns an empty sequence.

Normal guidance output does not add physical source paths or internal index state.

## Errors

| condition | required outcome |
|---|---|
| An in-scope canonical identity is conflicted. | The guidance catalog is unavailable. Do not select a duplicate winner. |
| An in-scope candidate is unreadable. | The guidance catalog is unavailable. Do not return a partial catalog. |
| An in-scope candidate cannot supply the required title or abstract. | The guidance catalog is unavailable. Do not omit the candidate and return a partial catalog. |

Current source validity and identity conflicts remain governed by the normal current Spec model.
This operation does not fall back to physical-path enumeration or a legacy guide source.

## Boundary

The operation projects package Specs only.
It does not reinterpret, summarize, repair, or independently own their authoring semantics.

Concrete diagnostic codes and CLI rendering are outside this contract.

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.operations.guidance` | Parent guidance-operation boundary. |
| `spec:product.design_records.authoring_standards` | PRODUCT source authority for the projected authoring standards. |
