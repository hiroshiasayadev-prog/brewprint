# USDM requirement: Record content retrieval

- **id**: `usdm:drmcp.mcp_capabilities.retrieval.record_content_retrieval`
- **status**: draft
- **date**: 2026-07-12
- **kind**: requirement
- **parent**: `usdm:drmcp.mcp_capabilities.retrieval`

## What this is

Product requirements for exact retrieval of current Design Records, selectable record-content projections, H2 section retrieval, and bounded source-content output.

## Requirements: Exact record retrieval
> source: literal

| id | requirement | notes |
|---|---|---|
| R001 | DRMCP must retrieve a current record selected by its exact canonical ref. |  |
| R002 | DRMCP must support retrieval of multiple retrieval selectors in one request. | The concrete request shape and maximum selector count are defined by downstream Specifications. |
| R003 | DRMCP must preserve every successful retrieval result when another supplied selector cannot be retrieved. |  |
| R004 | DRMCP must make the outcome of each selector that cannot be retrieved identifiable to the caller. | Exact outcome and diagnostic representation are defined by downstream Specifications. |
| R005 | DRMCP must return successful retrieval results in a deterministic order corresponding to the supplied selectors. |  |
| R006 | DRMCP must not repair, complete, normalize, or infer a supplied record ref. | Search and reference-resolution behavior belong to separate capabilities. |
| R007 | DRMCP must not expose physical source paths in normal retrieval results. |  |

## Requirements: Record content projection
> source: literal

| id | requirement | notes |
|---|---|---|
| R008 | DRMCP must allow retrieval output to be selected as record metadata, an H2 heading list, or complete source content. | The concrete projection-selection request shape is defined by downstream Specifications. |
| R009 | DRMCP must return the canonical ref with each successfully retrieved record result. |  |
| R010 | DRMCP must make the parsed metadata available for a successfully retrieved current record when the metadata projection is selected. | Artifact Specifications define the metadata fields and value formats. |
| R011 | DRMCP must return real H2 headings in source order when the H2 heading-list projection is selected. | H1 and H3-or-deeper headings are not part of this projection. |
| R012 | DRMCP must return complete source content without summarization, normalization, reformatting, or truncation when the complete-content projection is selected and the content fits within the applied output limit. |  |

## Requirements: H2 section retrieval
> source: literal

| id | requirement | notes |
|---|---|---|
| R013 | DRMCP must support selecting an H2 section by appending `#` and the H2 title to a record ref. | The selector form is `<record-ref>#<H2-title>`. This row does not determine whether the form is a canonical section ref outside retrieval. |
| R014 | DRMCP must match the H2 title in a section selector exactly against source H2 title text. | No title repair, completion, case repair, or fuzzy matching. |
| R015 | DRMCP must return the selected H2 heading and the source content through the line before the next H2 heading or through the end of the source. |  |
| R016 | DRMCP must include H3 and deeper subsections contained within the selected H2 section. |  |
| R017 | DRMCP must return selected H2 section content without summarization, normalization, reformatting, or truncation when the section fits within the applied output limit. |  |
| R018 | When no H2 title matches the selector, DRMCP must return an outcome from which the available H2 headings can be determined. | Exact outcome representation is defined by downstream Specifications. |
| R019 | When multiple H2 headings have the selected title, DRMCP must report the duplication as a warning and return the first matching H2 section in source order. | Authoring validation may prevent this source state in a later capability. |

## Requirements: Retrieval output limits
> source: literal

| id | requirement | notes |
|---|---|---|
| R020 | DRMCP must impose an upper bound on the number of retrieval selectors accepted by one request. | The numeric bound is defined by downstream Specifications. |
| R021 | DRMCP must impose an upper bound on the aggregate source-content output returned by one request. | Source-content output includes complete record content and H2 section content. |
| R022 | DRMCP must allow the caller to specify the source-content output limit applied to a request. | The measurement unit and request field are defined by downstream Specifications. |
| R023 | DRMCP must apply a default source-content output limit when the caller does not specify one. | The numeric default is defined by downstream Specifications. |
| R024 | DRMCP must enforce a server-side maximum for caller-specified source-content output limits. | The numeric maximum is defined by downstream Specifications. |
| R025 | DRMCP must either return the complete requested source content for one retrieval selector or omit that selector's source content. | Applies independently to complete record content and H2 section content. |
| R026 | DRMCP must not return partially truncated source content when a selector's requested content exceeds the remaining output limit. |  |
| R027 | DRMCP must make it identifiable when source content was omitted because of the applied output limit. |  |
| R028 | When one selector's source content cannot be returned within the output limit, DRMCP must preserve other selector results whose requested source content can be returned completely within the limit. |  |
| R029 | Reaching the source-content output limit must not by itself prevent DRMCP from returning requested metadata or H2 heading-list projections. |  |
