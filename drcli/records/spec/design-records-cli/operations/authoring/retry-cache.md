# Contract: DRCLI retry body cache

- **id**: `spec:drcli.design_records_cli.operations.authoring.retry_cache`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.operations.authoring`
- **contract_class**: interface

## What this is

Defines durable preservation and reuse of Markdown body content after a retryable authoring failure.
The cache is operational state, not a Design Record and not part of repository discovery or canonical record identity.

## Request

The retry-cache contract is entered in either of two ways.

| action | semantic input |
|---|---|
| Preserve received body | Exact body content already received by a create or update operation whose preparation or write attempt failed in a retryable way before successful completion. |
| Reuse cached body | One opaque body cache ID supplied as the body source for a later create or update invocation. |

Inline body and standard-input body have identical cache behavior once DRCLI has received the content.

A failure is retryable for body-cache purposes when the same body can be used again while the caller repairs non-body authoring inputs or retries a preparation or pre-write failure.
The cache preserves the exact received body; it does not normalize, summarize, regenerate, or reinterpret it.

A body cache entry contains only the preserved body content plus operational metadata needed to manage that entry.
It does not implicitly retain or select:

- repository root;
- target record identity;
- author fields;
- section selector;
- create or update operation choice;
- direct or deferred completion choice.

Using a body cache ID therefore never changes repository selection for the current invocation.

Cache state must survive:

- termination of the DRCLI process that created it;
- termination of the caller session;
- a later invocation from an independent caller session.

One shared DRCLI installation may retain body cache entries originating from work on multiple repositories.
Operational storage must prevent one entry from being confused with another; the storage engine and physical layout are not specified here.

A body cache ID is opaque.
Callers must not infer repository, path, timestamp, storage location, or body content from its spelling.

A retained entry remains reusable while valid under the installation's retention policy.
Unknown or expired entries cannot be used as body sources.
This contract does not choose a concrete retention duration.

## Response

When DRCLI preserves a received body after a retryable failure, the result includes:

- a body cache ID;
- enough validity information for the caller to know that the body can be retried;
- the original operation failure diagnostics separately from cache creation.

Presence of a body cache ID does not by itself mean that repository content was unmodified.
The enclosing authoring or write result remains authoritative for whether a write occurred.

When a later invocation supplies a valid body cache ID, DRCLI resolves it to the preserved body internally and applies the same body semantics as if that body had been supplied inline or through standard input.
The operation need not expose the cached body back to the caller.

Reuse does not consume the entry merely because one authoring attempt used it.
The entry may be reused while it remains valid.

## Errors

| condition | machine-readable classification | required effect |
|---|---|---|
| Body cache ID is unknown | `authoring_body_cache_unavailable` (`reason=unknown`), exit 3 | Reject it as a body source without repository modification. |
| Body cache entry is expired or otherwise no longer valid | `authoring_body_cache_unavailable` (`reason=expired` or `invalid`), exit 3 | Reject it as a body source without repository modification. |
| A cache ID is supplied together with another body source for the same body-bearing operation | CLI `usage_error`, exit 2; `authoring_invalid_request` (`reason=body_source_multiple`), exit 1 only if a type-valid logical request reaches semantic authoring | Reject the body-source combination. |
| Cache storage cannot preserve a body after a retryable failure | Original authoring `error` plus nested `error.context.cache_error` with `code=authoring_cache_preservation_failed`; enclosing error determines exit 1 or 3 | Report the authoring failure and cache-preservation failure distinctly in the same JSON document; do not invent a cache ID. |
| Cached body is valid but the new authoring request is otherwise invalid | The create/update `authoring_invalid_request` or `authoring_candidate_validation_failed`, exit 1, or CLI `usage_error`, exit 2 when applicable | Preserve normal authoring validation behavior; cache validity does not bypass request or candidate validation. |

When a retained cache entry cannot be read reliably because of I/O or permission failure (rather than a reliably established unknown/expired/invalid ID), use `authoring_operational_failure` (`reason=retained_state_io`), exit 3. When preservation itself fails after an original authoring error, the original error and cataloged nested `cache_error` are both reported in the same JSON document without fabricating `body_cache_id`.

## Boundary

This contract owns only opaque body preservation and reuse across independent invocations.
Create and update own when body content is required and how it contributes to the candidate record.
The shared write contract owns whether repository modification is permitted.

No temporary-file upload command, caller-session affinity, MCP session, storage engine, or fixed retention duration is defined here.

## Related specs

| ref | relation |
|---|---|
| spec:drcli.design_records_cli.operations.authoring | Parent authoring overview. |
| spec:drcli.design_records_cli.operations.authoring.create | Create body-source use. |
| spec:drcli.design_records_cli.operations.authoring.update | Update body-source use. |
| DRMCP-REQ-MCP-002 | Baseline for retry body preservation. |
