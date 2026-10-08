# Reference: Authoring transaction model

- **id**: `spec:drmcp.design_records_mcp.tools.authoring_transaction_model`
- **status**: draft
- **date**: 2026-06-30
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`

## What this is

Defines shared concepts for the 5 authoring transaction tools: proposal lifecycle, common response fields, body source and body cache, proposal validation affected record set, and diff_mode.

For the schema-level proposal model and body cache model, see `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema`.

## Current contract

### Authoring write flow

Authoring writes follow a propose → accept flow:

1. `propose_record_create` or `propose_record_update` creates a write candidate and returns a proposal ID, resolved target, diff response object, validation result, expiry, diagnostics, and a note.
2. Proposal creation MUST NOT modify repository files.
3. `get_proposed_write` retrieves the retained proposal content and lifecycle state.
4. `discard_proposed_write` discards the proposal; acceptance is prevented.
5. `accept_proposed_write` receives the proposal ID and, if all accept-time checks pass, writes to repository files.

Physical filesystem path is not the primary input for any authoring tool. Tool input uses record kind, record ID or `new` placeholder, domain, parent context, section selector, structured authoring fields, and body source when needed.

### Proposal lifecycle

Proposal state: `proposed` / `accepted` / `discarded`.

Expired proposals are not returned as retained proposals; they produce a `proposal_expired` diagnostic. Proposal retention is 3 days. All proposal responses include `expires_at`.

### Accept-time checks

`accept_proposed_write` re-verifies the following before writing:

- Proposal is not unknown / expired / discarded / already accepted.
- Target file state has not changed since proposal creation.
- Target kind matches the resolved target at proposal creation time.
- For create proposals: resolved ID is still available.
- For update proposals: target ID still resolves to the same record.
- Pre-write validation has no error-severity diagnostics.
- Every affected target and source-backed diagnostic can construct the required portable location before writing begins.

When any check fails, `accept_proposed_write` returns `written: false` and MUST NOT modify repository files.

Current-format post-write validation integration is deferred to `DRMCP-REQ-MCP-002`.

This Specification does not require the current authoring transaction to use the W011 snapshot architecture. It does not treat retained proposal state as an integrated substitute for current-format persisted-source validation.

Proposal lifecycle, pre-write validation, write eligibility, and write-result semantics remain owned by this Specification. `spec:drmcp.implementation` does not own current authoring-transaction validation behavior.

### Common authoring response fields

Proposal responses return at least:

| field | required | meaning |
|---|---:|---|
| `proposal_id` | yes | Proposal lookup key. Opaque string; callers must not interpret its structure. |
| `state` | yes | `proposed` |
| `operation` | yes | `create` / `update` |
| `target_kind` | yes | Resolved target record kind |
| `target` | yes | Requested / resolved target identity object |
| `expires_at` | yes | Proposal expiry timestamp |
| `retention_days` | yes | `3` |
| `diff` | yes | Diff response object (shape depends on `diff_mode`) |
| `validation` | yes | Validation result object |
| `diagnostics` | yes | Request / proposal diagnostic list |
| `note` | yes | Note explicitly stating that no repository files have been written and that `accept_proposed_write` is needed to apply the diff |

`target` returns at least:

| field | required | meaning |
|---|---:|---|
| `requested_id` | yes | ID supplied by the caller; may include `new` placeholder |
| `resolved_id` | yes | Final record ID resolved by MCP from the index |
| `kind` | yes | Target record kind |
| `domain` | no | Domain for domain-scoped workflow records |
| `parent_id` | no | Parent record ID used for parent-aware resolution (required for task create) |
| `path` | yes | Resolved normalized repository-relative path. Transparency output only — not the canonical request primary input. |

`diff` object fields vary by `diff_mode`:

| field | present in | meaning |
|---|---:|---|
| `format` | `patch`, `summary` | MVP: `unified` |
| `text` | `patch` only | Previewable unified diff text |
| `files` | `patch`, `summary` | Changed file summary list |
| `omitted` | `none` only | `true`; explicitly indicates that diff information was intentionally omitted |

`diff.files[]` entries include `path`, `change` (`create` / `modify`), and optionally `record_id` / `record_kind`. Workflow reciprocal metadata update proposals may have multiple entries in `files[]`.

### Authoring path representation

`target.path`, every `diff.files[].path`, repository path operands inside `diff.text`, and every successful `accept_proposed_write.files_written[].path` use one normalized repository-relative spelling.

Path rules:

- `/` is the only exposed separator;
- leading slash, trailing slash, duplicate separator, empty segment, `.`, and `..` are prohibited;
- Windows drive-qualified, UNC, device, URI, and other absolute forms are prohibited;
- repository spelling is preserved without case folding or locale-dependent normalization;
- the path must remain canonically within `repository_root` and the applicable current `records_root`;
- matching target, diff-summary, patch, and write-confirmation entries use exactly the same repository-relative path.

Git-style unified diff syntax remains distinct from path values:

- modify patches use `diff --git a/<path> b/<path>`, `--- a/<path>`, and `+++ b/<path>`;
- create patches may use `--- /dev/null` and `+++ b/<path>`;
- `a/` and `b/` are diff side prefixes, not part of the repository-relative path;
- `/dev/null` is a diff sentinel, not a repository or physical path;
- host absolute paths and backslash-separated Windows paths are prohibited in `diff.text`.

These scalar authoring paths are explicit transaction transparency, patch, and write-confirmation outputs. They are not diagnostic `location` objects and are never accepted as primary authoring request inputs.

`validation` returns at least `ok` and `diagnostics`. `ok` is `true` when there are no error diagnostics.

Validation failure and write failure are distinct states and must not be conflated.

Proposal-local validation (`validation.diagnostics`) describes only diagnostics for the affected record set in the candidate state. Unrelated existing repository diagnostics must not appear in proposal-local blocking diagnostics. If broader repository health is exposed alongside a proposal, it must use a field separate from `validation` (e.g. `repository_health`) and must not affect `validation.ok` or accept-time write eligibility unless the diagnosed record is in the affected record set.

Standard note text (exact wording is an implementation detail but must not weaken this meaning):

```text
No repository files have been written. Call accept_proposed_write with this proposal_id to apply the diff.
```

### Body source and body cache

Operations requiring Markdown body input accept exactly one of `body` or `body_cache_id`. Supplying both is invalid and must not create a proposal or body cache entry. Supplying neither is valid only for operations that do not require body input.

For `propose_record_create`:
- `fields` is required in all create modes.
- `body` and `body_cache_id` are optional section-only content sources, valid only when combined with `fields`.
- `body` supplies content sections only (starting at the first section heading like `## Goal`). It must not include an H1, metadata block, YAML metadata, `metadata id`, or a guessed resolved ID.
- Full-record body create without `fields` is invalid. When `body` is submitted without `fields`, the request is rejected; the response should include a new `body_cache` entry for retry with `fields + body_cache_id`.

MVP body source rules by operation case:

| operation case | body requirement |
|---|---|
| `propose_record_create` structured metadata only | `fields` present; `body` / `body_cache_id` omitted |
| `propose_record_create` structured metadata plus content sections | `fields` and section-only `body` present; `body_cache_id` omitted |
| `propose_record_create` retry with cached content sections | `fields` and `body_cache_id` present; `body` omitted |
| `propose_record_create` body-only / cache-only create | invalid; `fields` is required |
| `propose_record_update` `metadata_block_replace` | `body` / `body_cache_id` must be omitted |
| `propose_record_update` `metadata_fields_replace` | `body` / `body_cache_id` must be omitted |
| `propose_record_update` `named_section_replace` | exactly one of `body` / `body_cache_id` |

Body cache retention is 3 days. Cache responses include `expires_at`. Cache entries remain reusable within the retention period, including after being used to create a proposal. Expired entries must not be used to create proposals.

When a large `body` is submitted and proposal/write preparation fails before the body can be persisted, the MCP should preserve the submitted body and return a retryable body cache entry:

```json
{
  "proposal_created": false,
  "body_cache": {
    "body_cache_id": "bc_opaque",
    "expires_at": "2026-06-05T00:00:00Z",
    "retention_days": 3
  },
  "diagnostics": [
    {
      "category": "proposal_preparation_failed",
      "severity": "error",
      "message": "proposal could not be prepared; retry with body_cache_id"
    }
  ]
}
```

### Proposal validation affected record set

The proposal-local affected record set is the set of records whose content is created or modified by the proposal.

For `propose_record_create`:
- The proposed target record.
- Any related records actually modified for required reciprocal workflow metadata under `reciprocal_update_mode`.

For `propose_record_update`:
- The target record being updated.

Proposal-time validation runs against the candidate repository state (proposal diff applied to the current state) but the returned diagnostics are filtered to the affected record set. Accept-time pre-write validation uses the same model after staleness, target-change, and ID-collision guards. Diagnostics in `validation.diagnostics` must be reproducible by running `validate_records` against the same affected record set in the same candidate or accepted/materialized state.

A source-backed proposal-local diagnostic requires the current portable `location` defined by `spec:drmcp.design_records_mcp.schema.diagnostics`, including an unmaterialized create target with a deterministic destination. If any required direct, conflict-member, or conflict-candidate location cannot be constructed, proposal preparation fails, no retained proposal is created, and no repository file is written.

Before acceptance begins writing, every affected path and required source-backed diagnostic location must be constructible. Failure returns `written: false` with an empty `files_written` list and starts no write. An implementation failure detected after files were actually modified must not report `written: false` or erase the actual written-file state.

> Source: V01-REQ-MCP-012, V01-TASK-MCP-011-01

### diff_mode

`propose_record_create` and `propose_record_update` accept an optional `diff_mode` parameter.

`diff_mode` controls the `diff` field shape in the proposal response. Proposal internal retention, accept behavior, validation, and diagnostics are independent of `diff_mode`.

| value | meaning |
|---|---|
| `summary` | Default. Returns `diff.files`; omits `diff.text`. |
| `patch` | Returns the full unified diff including `diff.text`. |
| `none` | Omits diff information; returns `diff: {"omitted": true}`. |

Default when omitted: `summary`.

In `summary` mode, `diff.files[]` entries (`path`, `change`, `record_id`, `record_kind`) provide a concise per-file summary. `diff.text` is not included.

Invalid `diff_mode` values produce `invalid_request` and no proposal is created.

`get_proposed_write` does not support `diff_mode`; it always returns the `patch`-equivalent response (full `diff.text` included).

> Source: V01-REQ-MCP-031
