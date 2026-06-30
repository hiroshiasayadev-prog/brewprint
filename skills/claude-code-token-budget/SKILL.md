# Claude Code token-budget prompt authoring

## Purpose

This skill governs the ChatGPT assistant when it creates ready-to-run prompts for Claude Code.

The goal is to reduce Claude Code token consumption without weakening implementation correctness, independent review, or Brewprint traceability.

This skill does not govern Claude Code directly. It governs how the ChatGPT assistant scopes and writes prompts that the user will give to Claude Code.

## Use this skill when

Use this skill when the user asks for any of the following:

- a Claude Code implementation prompt;
- a Claude Code review or re-review prompt;
- a finding-correction prompt;
- a closure-synchronization prompt;
- a handoff prompt for a new Claude Code session;
- a task graph intended for Haiku or Sonnet execution;
- a lower-token version of an existing Claude Code workflow.

Apply this skill by default when a task includes implementation, review, correction, and post-correction review.

## Required companion and execution-hub gate

Read `executor-ready-task-design.md` together with this file for implementation prompt and Task-graph design.

Before issuing an implementation or finding-correction prompt, evaluate the application trigger in `execution-hub-task-pattern.md`.

When the trigger is met:

- read `execution-hub-task-pattern.md`;
- persist the execution graph in Work Item and Task records;
- independently review the graph;
- synchronize the reviewed release;
- issue prompts only for executor Tasks released by that synchronization.

Do not treat chat-only slice labels, handoff prose, or a ready-to-run prompt as a substitute for the persistent execution graph.

Do not force the pattern onto a single self-contained executor Task with no shared writer, dependent output, separate aggregate owner, or release coordination need.

## Core principles

### Optimize prompt scope, not only session count

A small code change can still consume many tokens when each session re-reads broad authority, explores the repository, runs repeated commands, or performs an unnecessary global review.

Do not automatically solve high token use by splitting an already-scoped Task into more sessions. First reduce:

- authority read breadth;
- repository exploration;
- repeated full-file reads;
- repeated test execution;
- unbounded review scope;
- verbose command output;
- duplicated background in handoff prompts.

### Preserve required authority

Token reduction must not remove authority required to make the requested judgment.

Include the smallest complete authority set. Do not include every upstream record merely because it exists.

A complete authority set normally contains:

- the repository instruction source required for the session;
- the target Task;
- the directly governing accepted contract needed for the work;
- the exact implementation and test files in scope;
- the exact prior finding for a correction or closure review.

Include the parent Work Item, Requirement, ADR, or Specification only when the Task does not contain enough contract detail or the requested judgment explicitly depends on it.

Do not paste entire records into a prompt when Claude Code can read the exact paths locally. Provide paths and, when useful, the relevant section names.

### One prompt, one responsibility

Each prompt must have one primary responsibility:

- inventory;
- implementation;
- independent review;
- finding correction;
- finding closure review;
- closure synchronization.

Do not combine implementation and independent review in one prompt.

Do not combine finding correction and closure review in one prompt.

Do not ask an implementation session to update Design Record lifecycle evidence unless the user explicitly requests the combined operation.

### Independent review remains independent

An independent review must inspect the scoped implementation, tests, diff, and governing contract.

Do not let the reviewer accept the implementer's summary as proof.

Token reduction for review comes from a narrow review boundary, not from removing direct evidence.

### Reuse known state

Use facts already established in the conversation. Do not ask the user to restate:

- repository path;
- accepted decisions;
- current Task status;
- previous findings;
- completed verification;
- known changed files.

A handoff prompt should state only the current boundary, accepted facts needed by the next session, and the exact next action.

## Prompt construction procedure

### 0. Evaluate persistent graph requirements

Before classifying the prompt phase, determine whether `execution-hub-task-pattern.md` applies.

If it applies and the graph has not been persisted, reviewed, and released, stop prompt construction and return to scope-freeze or Design Record authoring.

The prompt must not introduce an essential file, symbol, test, dependency, writer, model route, or verification owner that is absent from the accepted Task graph.

### 1. Classify the phase

Choose exactly one phase.

| phase | purpose |
|---|---|
| inventory | Locate the minimum implementation or review inputs when the exact scope is not yet known. |
| implementation | Produce the scoped code or test change. |
| independent review | Judge the scoped change against the governing contract. |
| finding correction | Correct specified findings only. |
| finding closure review | Verify specified findings and direct regressions only. |
| closure synchronization | Update Task or Work Item lifecycle and Evidence after accepted review. |

Do not create an inventory phase when the target files and contract boundary are already known.

### 2. Build the minimum context pack

List exact paths under `Read first`.

Prefer this order:

1. repository instruction source required for the agent;
2. target Task;
3. direct governing contract only when needed;
4. exact production files;
5. exact test or fixture files;
6. exact previous finding or accepted implementation summary.

Do not use repository-wide discovery to locate records with already-known paths or IDs.

### 3. Define the execution boundary

Every prompt must state:

- files allowed to change;
- files that must not change;
- commands that must run;
- whether staging or commit is prohibited;
- whether Design Record updates are prohibited;
- whether production implementation is prohibited;
- the exact review boundary;
- the stop condition.

Do not add generic prohibitions that are unrelated to the task.

### 4. Bound exploration

Use scoped search only when the exact symbol or dependency is unknown.

Prompt rules should normally require:

- no repository-wide traversal;
- no broad recursive grep;
- no reading unrelated Tasks or Work Items;
- no unrelated refactor or cleanup;
- no repository-wide clean-status claim from scoped evidence;
- no repeated full-file read unless the file changed or a specific unresolved question requires it.

Do not prohibit all search when a narrow symbol search is necessary. Name the permitted directory, file type, or symbol.

### 5. Bound command execution

Prefer the narrowest verification that can establish the requested result.

Use this order when applicable:

1. focused test for the changed package or behavior;
2. scoped static or fixture validation;
3. broader suite only when the Task or closure gate requires it.

Prompt rules should normally require:

- do not rerun an unchanged failing command without a concrete diagnostic reason or code change;
- do not paste full successful logs when exit status and concise evidence are sufficient;
- preserve the exact failing excerpt when a failure blocks completion;
- use scoped Git status, diff, and whitespace checks when the task boundary is scoped.

### 6. Define concise output

Request only information needed for the next decision.

Default output shape:

1. result or verdict;
2. changed or reviewed files;
3. verification and exit status;
4. findings or blocker;
5. next gate readiness.

Do not request a narrative replay of every command or file read.

## Phase-specific rules

### Inventory

Use inventory only when implementation ownership, file scope, or verification commands are genuinely unknown.

Inventory must be read-only.

Expected output:

- exact files required;
- exact governing contract refs;
- implementation symbols or seams;
- focused verification commands;
- unresolved blockers.

The inventory must not redesign the Task or produce implementation.

### Implementation

The implementation prompt must:

- name the exact behavior to implement;
- name allowed production and test files;
- preserve accepted boundaries;
- prohibit independent review and lifecycle closure;
- require focused verification;
- stop on a genuine contract decision not already accepted.

Do not ask the implementer to prove repository-wide correctness.

Do not request broad architectural commentary unless the implementation becomes blocked by an actual architecture conflict.

### Independent review

The review prompt must:

- be read-only;
- identify the exact Task and changed-file boundary;
- identify the governing contract and Done conditions;
- inspect the scoped diff and relevant tests;
- distinguish blocking, major, minor, and advisory observations;
- report all material findings without expanding into unrelated design review;
- prohibit corrections, staging, commit, and lifecycle updates.

Do not impose an arbitrary finding-count limit. Scope the review instead.

### Finding correction

The correction prompt must quote or reproduce the exact finding IDs and required corrections.

The correction prompt must:

- change only files needed to close those findings;
- preserve already accepted behavior;
- avoid unrelated cleanup;
- run focused regression verification;
- avoid lifecycle synchronization;
- avoid claiming that the findings are closed before independent review.

When findings are independent, create separate prompts only when they can be executed safely in parallel. Do not split mechanically.

### Finding closure review

The closure review prompt must focus on:

- whether each specified finding is closed;
- whether the correction introduced a direct regression;
- whether required focused verification exists and passed.

Do not reopen the full design unless the correction exposes a blocking contradiction.

The reviewer may report a new material finding only when it is caused by, or directly revealed by, the correction under review.

### Closure synchronization

The closure prompt must:

- update only the target Task, parent Work Item, or explicitly named records;
- preserve accepted contracts;
- record review disposition and verification evidence;
- prohibit production code, test, fixture, ADR, Requirement, or Specification changes unless explicitly in scope;
- prohibit staging or commit unless explicitly requested.

Do not rerun implementation work during closure synchronization.

## Model routing guidance

Model choice does not replace scope control.

Default routing:

| work | default model class |
|---|---|
| mechanical inventory with exact boundaries | lower-cost model |
| fixture or repetitive test materialization | lower-cost model when the contract is explicit |
| straightforward finding closure review | lower-cost model |
| lifecycle and Evidence synchronization | lower-cost model |
| production implementation with semantic choices | stronger coding model |
| ambiguous failure investigation | stronger coding model |
| architecture or ownership boundary review | stronger reasoning model |
| correction requiring contract interpretation | stronger reasoning or coding model |

Do not route work to a lower-cost model merely because only one file changes. Contract density and ambiguity matter more than file count.

## Prompt template

Use the following shape and remove sections that are not relevant.

```text
C:\Users\imved\projects\brewprint

Execute <TASK-ID or exact operation>.

Responsibility:
<inventory | implementation | independent review | finding correction | finding closure review | closure synchronization>

Read first:
1. <required repository instruction source>
2. <target Task>
3. <direct governing contract, only if required>
4. <exact implementation/test files>

Current accepted state:
- <only facts needed for this phase>

Required work:
- <specific operation>
- <specific behavior or finding>

Allowed changes:
- <exact paths, or none>

Do not:
- <task-relevant prohibitions only>
- perform repository-wide traversal
- perform unrelated cleanup or redesign
- stage or commit unless explicitly requested

Verification:
- <focused command>
- <broader command only when required by the gate>

Stop condition:
- Stop with BLOCKED only when a required contract decision or required input is unavailable.
- Name the exact missing decision, file, or evidence.
- Do not broaden repository exploration to guess the answer.

Output:
1. Result or verdict
2. Changed or reviewed files
3. Verification with exit status
4. Findings or blocker
5. Gate readiness
```

## Token-budget checks before returning a prompt

Before presenting a Claude Code prompt, verify:

- The prompt has one responsibility.
- The prompt does not repeat the full project history.
- Every required read has a concrete reason.
- No known path is replaced by a search instruction.
- Repository-wide traversal is unnecessary and prohibited.
- The allowed write set is explicit.
- Verification is focused before broad.
- Successful command logs are requested as concise evidence.
- Review independence is preserved.
- The prompt does not rely on `/compact` as the main token-control mechanism.
- An already-scoped Task is not split further without a concrete dependency or parallelism reason.
- The execution-hub trigger was evaluated.
- When the trigger applies, the Work Item and Task graph is persisted, independently reviewed, and released.
- The prompt introduces no essential contract absent from the released executor Task.

## Failure pattern correction

When the user reports excessive Claude Code consumption, do not immediately recommend more sessions or repeated compaction.

First inspect the workflow for:

- broad startup reads;
- duplicated authority in every phase;
- repository exploration despite known paths;
- repeated review of accepted boundaries;
- repeated test runs without changes;
- full log retention;
- implementation sessions performing self-review;
- closure sessions redoing implementation analysis;
- re-review prompts that accidentally request a full independent review.

Rewrite the next prompt to remove the identified source of consumption while preserving the required gate.
