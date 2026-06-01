# 093: Design Records MCP authoring transaction model

- **status**: accepted
- **date**: 2026-06-01
- **depends_on**: ADR-076, ADR-077, ADR-092
- **supersedes**:
- **migrated_to_spec**: 2026-06-02

> This ADR is a snapshot of the decision at the time it was written.
> Refer to spec for the current behavior.

## 背景

Design Records MCP currently provides artifact-oriented read, navigation, validation, reference resolution, next ADR suggestion, and authoring guidance retrieval.

The current public surface is intentionally read-only. It can list, retrieve, validate, and resolve design records and workflow artifacts by stable IDs or guide IDs, but it cannot create or update records.

This creates a dogfooding gap. AI assistants can read and navigate records through Design Records MCP, but must switch back to filesystem path editing when users ask to create or update requirements, work items, tasks, ADRs, or existing spec sections.

That split makes authoring path-oriented even when the user intent is artifact-oriented, such as `REQ-MCP-new`, `WORK-MCP-008`, or `TASK-MCP-008-01`.

REQ-MCP-008 captures the need to add authoring transaction support. TASK-MCP-008-01 confirmed that the absence of write capability is not a bug in the current MCP; it is a deliberate public contract gap in the read-only MVP.

## 決定

Design Records MCP will add artifact-oriented authoring write capability using a propose -> accept transaction model.

The authoring transaction model has the following decisions.

1. Authoring writes are not immediate writes. A proposal operation prepares the change and returns a proposal ID, resolved target identity, previewable diff, validation result, expiry information, and note. Proposal creation must not modify repository files.
2. A separate accept operation applies a previously created proposal. Accept responses must distinguish whether repository files were written.
3. Proposal records must carry enough base-state information to detect stale application. Accept must re-check the target file state, target kind, and resolved target ID availability immediately before writing. If the proposal is stale, the target changed, or the resolved ID is no longer available, accept must not modify repository files and must return retry/re-propose guidance.
4. Write inputs use artifact identity as the primary public contract. Tool inputs are based on record kind, ID, domain, section name, and structured authoring fields. Physical file paths may be returned for transparency but are not primary authoring inputs.
5. Create operations may accept a `new` placeholder in the ID position. The MCP resolves the final ID on the server side using the appropriate sequence. Update operations reject `new` IDs.
6. Task creation requires enough parent context to avoid ambiguous numbering. The concrete request schema belongs to the tools spec, but the model requires parent-aware task ID resolution.
7. Workflow artifact creation must define how required reciprocal workflow metadata is kept valid. In particular, task creation must not rely on task ID shape alone to infer the parent relation. Workflow artifact create proposals may include required reciprocal metadata updates in the same proposal when those updates are necessary to keep the created artifact valid under workflow relation validation, or they may return explicit required follow-up updates before acceptance.
8. Required reciprocal metadata updates in a workflow artifact create proposal are not considered a general-purpose multi-record atomic transaction in this ADR. The MVP still excludes arbitrary bundling of unrelated records into one atomic transaction with rollback semantics.
9. Operations that require a large Markdown body accept exactly one body source: `body` or `body_cache_id`. Supplying both is invalid and must not create a proposal or cache entry. Operations that do not require a large Markdown body, such as template-driven create operations derived from structured fields, may omit both. The set of body-required operations is owned by the tools spec.
10. Proposal and body caches use a 3 day retention period. Proposal and cache responses must include expiry information. Expired or unknown proposal/body cache IDs must not write files and must return machine-readable diagnostics.
11. If a large submitted body cannot be persisted into a proposal because proposal/write preparation fails, the MCP should preserve the submitted body when possible and return a `body_cache_id` with retry guidance.
12. MVP update operations are set-only. They replace the kind-specific metadata block as a whole or replace a named Markdown section as a whole. Add/remove convenience operations are out of scope for the MVP.
13. Metadata replacement covers the kind-specific metadata block, including YAML front matter for spec records and bullet metadata blocks for decision / requirement / work item / task records. The exact serialization belongs to the tools spec, but replacement must validate required recognized fields and report missing-field diagnostics.
14. Named section replacement is valid only when the section selector resolves to exactly one Markdown section. Zero-match or multi-match selectors must not write files and should return candidate headings when possible.
15. Write failure and validation failure are separate response states. Accept results must indicate written/not-written semantics, and validation failures should include repair guidance when possible.
16. The MVP does not provide automatic rollback after accepted writes. If post-write validation fails, the response must report that files were written and guide the caller to create a repair proposal.

The MVP authoring surface covers:

- creating requirement records;
- creating work item records;
- creating task records;
- creating decision records;
- replacing kind-specific metadata blocks as a whole;
- replacing named Markdown sections as a whole;
- getting, accepting, and discarding proposed writes.

The exact tool names, request/response schemas, proposal ID format, proposal lifecycle state names, diagnostics, section matching algorithm, metadata serialization, and storage mechanics are owned by the Design Records MCP tools spec and implementation, not by this ADR.

## 理由

Immediate write is unsafe for AI-assisted authoring. It makes it easy to corrupt Markdown, write to the wrong file, or lose a large generated body when an operation fails midway. A proposal-first model makes the intended diff reviewable before repository files are modified.

Artifact-oriented inputs preserve the same abstraction already used by read-side Design Records MCP tools. Users and assistants usually reason in terms of `REQ-*`, `WORK-*`, `TASK-*`, `ADR-*`, `SPEC-*`, guide IDs, and section names. Requiring physical paths as primary inputs would preserve the current path-oriented authoring gap.

Server-side `new` resolution prevents assistants from guessing record numbers and reduces collision risk. It also keeps numbering rules close to the index and validation logic that already know existing records. Accept-time revalidation is still required because another agent may create the resolved ID or modify the target file after proposal creation.

The `body` / `body_cache_id` split protects large generated Markdown bodies. A retryable cache ID is more reliable than asking an assistant to regenerate a long body, because regenerated content can drift from the originally reviewed text.

A 3 day cache retention period is long enough for normal review and retry loops in an always-running local MCP server, while avoiding indefinite retention of stale proposed content.

Set-only MVP updates are intentionally conservative. Whole metadata block replacement and whole named-section replacement are easier to reason about, review, diff, and validate than add/remove relation helpers or partial Markdown AST edits. Add/remove can be introduced later once the transaction boundary is proven.

Workflow artifact creation needs special care because workflow relation validation includes reciprocal metadata. A task create flow that writes only the task file can leave the parent work item invalid. The proposal layer therefore must either include required reciprocal metadata updates in the proposed impact or make required follow-up updates explicit before acceptance.

Named-section replacement must fail closed when the section selector is ambiguous. Replacing the first matching heading is not safe enough for agent authoring because evidence and task sections are frequently edited and may contain nested or repeated headings.

Separating write failure from validation failure avoids ambiguous failure states. A caller must know whether repository files were modified before deciding whether to retry, repair, or discard.

## 却下した代替案

### Immediate write tools

Rejected for the MVP. Immediate writes reduce ceremony but make AI-assisted authoring too easy to apply without review. They also do not address the need to preserve failed large bodies and report clear written/not-written state.

### Filesystem path as primary input

Rejected. Path-first authoring would leave the core problem unresolved: read/navigation would remain artifact-oriented while write/update would remain filesystem-oriented.

### Add/remove relation operations in the MVP

Rejected for the MVP. Add/remove operations require current-state reads, merge behavior, ordering rules, duplicate handling, and conflict semantics. Whole-set replacement is simpler and safer for the first authoring transaction contract.

### Partial Markdown AST editing in the MVP

Rejected for the MVP. It is more powerful but introduces substantial ambiguity around section matching, formatting preservation, nested headings, and conflict behavior. Named-section whole replacement is sufficient for the first phase, but only when the section selector resolves to exactly one target section.

### Spec skeleton creation in the MVP

Rejected for the MVP. Existing spec records may be updated by metadata block or named-section replacement, but new spec document placement cannot be derived safely from `SPEC-new` or `SPEC-<slug>` alone. Spec placement needs a separate domain tree / placement discovery capability before artifact-oriented spec creation can be made safe; this follow-up is captured by REQ-MCP-010.

### Indefinite proposal/body cache retention

Rejected. Proposed content can contain stale or sensitive design text. A bounded 3 day retention window is enough for normal retry/review workflows without creating an unbounded local cache.

## 影響

The Design Records MCP tools spec must define the authoring transaction public contract. It should cover proposal creation/update, proposal retrieval, proposal discard, proposal accept, body source validation, cache behavior, ID resolution, base-state / staleness detection, diff response, validation result, and write outcome semantics.

The tools spec should define proposal lifecycle states and diagnostics for accepted, discarded, expired, stale, and unknown proposal/body cache IDs. Proposal and body cache responses should include expiry information.

The implementation must add proposal storage, body cache storage, 3 day retention handling, artifact-to-path resolution for supported record kinds, create/update rendering, diff generation, accept-time writing, and post-write validation/repair feedback.

The following categories describe the test surface implied by this decision. Concrete test cases, fixture shape, and pass criteria are owned by the implementation task and tests, not by this ADR.

Tests must cover at least:

- proposal creation does not write files;
- accept writes the proposed change;
- stale proposals, changed targets, and ID collisions are rejected before writing;
- discard prevents later accept;
- expired or unknown proposal/body cache IDs return diagnostics without writing;
- `new` works for create and is rejected for update;
- task creation requires parent-aware context;
- workflow artifact creation handles or reports required reciprocal metadata updates;
- `body` and `body_cache_id` are mutually exclusive;
- body cache retry works for cached content;
- validation failure and write failure responses are distinguishable;
- set-only metadata block / section updates behave deterministically;
- ambiguous or missing section selectors fail without writing;
- existing read-only tools continue to work.

The MVP does not include:

- generic filesystem write tools;
- path-first authoring APIs;
- add/remove relation convenience operations;
- partial Markdown AST editing;
- general-purpose multi-record atomic transactions with rollback semantics;
- automatic close cascades;
- automatic rollback of accepted writes after post-write validation failure;
- formatter integration;
- indefinite proposal/body cache retention;
- force-accepting invalid proposals.

## Evidence

- source requirement: REQ-MCP-008
- work item: WORK-MCP-008
- evidence task: TASK-MCP-008-01
- ADR drafting task: TASK-MCP-008-02
- ADR review task: TASK-MCP-008-03
- related guidance capability: REQ-MCP-005 / WORK-MCP-005
- related workflow navigation capability: REQ-MCP-007 / WORK-MCP-007
- current tools spec: SPEC-design-records-mcp-tools
- decision ID suggestion: `suggest_next_record` returned `ADR-093` for this title.
- implementation / agent-usability review: Codex review reported `Needs revision before Opus review`; major findings for staleness guard, reciprocal workflow metadata, section selector ambiguity, and metadata terminology were incorporated on 2026-06-01.
- independent ADR review: Opus review reported `OK with minor fixes before accept`; findings for general-purpose multi-record atomic transaction wording, `depends_on`, section order, body-source omission, and test ownership wording were incorporated before acceptance on 2026-06-01.
- spec reflection review: review reported `Needs revision before implementation`; spec skeleton creation / `SPEC-new` was excluded from MVP and spec placement discovery was captured as REQ-MCP-010.
- impl commit: tbd
