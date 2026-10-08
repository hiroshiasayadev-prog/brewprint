# Overview: DRCLI authoring operations

- **id**: `spec:drcli.design_records_cli.operations.authoring`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations`

## What this is

Defines DRCLI authoring transaction behavior for Design Record create, partial update, write completion, retry body caching, and explicitly requested deferred proposals.
PRODUCT remains authoritative for record semantics, author-facing fields, identity, placement, lifecycle, section shape, and validation rules.

## Current contract

| concern | contract |
|---|---|
| Supported records | ADR, Requirement, Work Item, Task, Specification, and Investigation authoring use their PRODUCT-owned authoring standards. |
| Normal completion | A successful create or update completes its write in the same DRCLI invocation. Mandatory propose-then-accept is prohibited. |
| Deferred proposal | Proposal preparation is optional and occurs only when the caller explicitly requests deferred completion. |
| Body sources | An operation that consumes Markdown body content accepts it inline, from standard input, or from an existing opaque body cache ID, subject to the operation-specific body rules. |
| Retry body preservation | A retryable preparation failure after DRCLI has received body content preserves that exact body and returns an opaque body cache ID. |
| Process independence | Body-cache and retained-proposal state survives CLI process termination and caller-session termination. |
| Shared installation | One DRCLI installation may retain independent operational state for multiple repositories without allowing retained proposal state to select or modify the wrong repository. |
| Candidate validation | Create and update validate the complete affected candidate state before modification. Unrelated repository diagnostics do not become authoring blockers. |
| Write guard | Direct writes and deferred-proposal acceptance re-check stale targets, identity availability, required coupled changes, and required validation immediately before modification. |
| Interaction | Authoring is non-interactive by default. No confirmation prompt or temporary-file upload step is required. |

DRCLI consumes PRODUCT-owned semantics instead of copying them as independent DRCLI rules.
Generated H1 text, generated identity projection, generated metadata serialization, generated filenames, and generated placement are not caller-authored substitutes for PRODUCT authority.

USDM authoring is not added to this operation family by this Task. The PRODUCT USDM authoring guide explicitly leaves standalone tool contracts and repository-layout integration out of scope; DRCLI does not invent that missing integration here.

### Supported authority set

| record kind | PRODUCT authority |
|---|---|
| ADR | spec:product.design_records.authoring_standards.adr_authoring |
| Requirement | spec:product.design_records.authoring_standards.requirement_authoring |
| Work Item | spec:product.design_records.authoring_standards.work_item_authoring |
| Task | spec:product.design_records.authoring_standards.task_authoring |
| Specification | spec:product.design_records.authoring_standards.spec_authoring |
| Investigation | spec:product.design_records.authoring_standards.investigation_authoring |

### Body-source rules

Body transport does not change authoring semantics.

- Inline body and standard-input body are equivalent once DRCLI has received the bytes.
- An existing body cache ID supplies only the preserved body content. It does not supply repository selection, target identity, metadata fields, section selectors, or completion disposition.
- An operation that requires body content accepts exactly one body source.
- An operation that does not consume body content rejects an unrelated body source.
- No body source may be used to smuggle an H1, generated metadata, generated identity, or generated placement where the applicable PRODUCT authoring contract excludes those values.
- A body cache ID never changes the repository selected for the current invocation.

### Candidate affected set

The affected set contains every record whose persisted content would change if the authoring operation completes.

The set includes the requested target and every additional record that PRODUCT semantics require to change as part of the same authoring outcome.
DRCLI must not make PRODUCT-required coupled updates optional merely to keep the affected set single-record.

Candidate-state validation evaluates the complete affected set after all requested and required changes are applied conceptually.
Existing diagnostics outside that set do not change the candidate result or block the authoring write.

### DRMCP baseline adaptation

| DRMCP baseline behavior | DRCLI disposition |
|---|---|
| Proposal creation never writes | Retained for explicitly requested deferred proposals. |
| Every write requires propose then accept | Replaced. Normal create and update write in one invocation. |
| Candidate affected-set validation | Retained for direct and deferred paths. |
| Accept-time stale-target and collision checks | Retained as shared write guards for direct writes and proposal acceptance. |
| Opaque body cache for retry | Retained and extended across independent CLI invocations and caller sessions. |
| Retained proposal state | Retained only for explicit deferred callers and made process-independent. |
| MCP session state | Removed. No authoring behavior depends on an MCP session. |
| Body upload before authoring | Removed. Inline and standard-input body are first-class inputs. |
| Concrete retention duration, storage engine, flags, and rendered output shape | Not fixed by this contract. |

## Non-goals

- Defining PRODUCT-owned record fields, lifecycle values, section schemas, ID grammar, or placement rules.
- Choosing command names, command flags, exit numbers, rendered CLI output, or a persistence engine.
- Requiring a daemon, MCP registration, or caller session.
- Providing arbitrary multi-record transaction semantics beyond changes required by PRODUCT authority.
- Defining USDM repository integration or standalone USDM authoring operations.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Create operation | Contract | spec:drcli.design_records_cli.operations.authoring.create | Creates one PRODUCT-governed Design Record and normally writes it in the same invocation. |
| Update operation | Contract | spec:drcli.design_records_cli.operations.authoring.update | Applies one candidate partial update to an existing PRODUCT-governed Design Record. |
| Write completion | Contract | spec:drcli.design_records_cli.operations.authoring.write | Owns shared candidate validation, stale-write guards, and modification eligibility. |
| Retry body cache | Contract | spec:drcli.design_records_cli.operations.authoring.retry_cache | Preserves received body content for a later independent invocation after retryable failure. |
| Deferred proposal | Contract | spec:drcli.design_records_cli.operations.authoring.deferred_proposal | Owns optional retained candidate preparation, inspection, discard, and acceptance. |

## Requirement trace

| DRCLI-REQ-CLI-001 outcome | owning contract |
|---|---|
| Create and partial update follow PRODUCT authoring authority | create, update |
| Successful create and update do not require propose then accept | create, update, write |
| Deferred proposal remains available only when explicit | deferred_proposal |
| Inline body input | create, update |
| Standard-input body input | create, update |
| Existing body cache ID input | create, update, retry_cache |
| Retryable failure preserves a received body | retry_cache |
| Cache survives process and caller-session termination | retry_cache |
| Retained proposal state survives process and caller-session termination | deferred_proposal |
| One shared installation retains state for multiple repositories | retry_cache, deferred_proposal |
| Candidate validation precedes modification | write |
| Accept and direct write re-check stale targets and required validation | write, deferred_proposal |
| Non-interactive authoring | create, update |
| No temporary-file upload requirement | create, update, retry_cache |

## Related specs

| ref | relation |
|---|---|
| DRCLI-REQ-CLI-001 | Source DRCLI requirement. |
| DRMCP-REQ-MCP-002 | Authoring-transaction design baseline. |
| spec:product.design_records.authoring_standards | PRODUCT-owned authoring authority. |
| spec:product.design_records.repository_layout | PRODUCT-owned placement authority. |
