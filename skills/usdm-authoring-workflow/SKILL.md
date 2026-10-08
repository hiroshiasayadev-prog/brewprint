# USDM authoring workflow

## Purpose

This skill governs the workflow for deriving MVP USDM requirement records from Product or app Specifications.

It prevents:

- requirement-like source text being skipped before review;
- concept prose being copied into USDM rows;
- final `RNNN` row IDs being assigned before candidate cleanup;
- implementation architecture names replacing Product-owned requirement language;
- user review packages becoming too large to review;
- USDM files being written before the row set is approved.

The repository USDM authoring standard remains authoritative for USDM format, row semantics, row-ID stability, and coverage metadata.

## Use this skill when

Use this workflow when a session will derive or update USDM requirement rows from source Specifications.

Use this workflow for:

- creating a new USDM requirement branch;
- extracting requirement candidates from source specs;
- converting requirement candidates into USDM rows;
- reviewing whether candidates are implementation-facing requirements;
- preparing a compact user review package before writing USDM files;
- updating existing USDM rows when source Specifications change.

## Do not use this skill when

Do not use this workflow for:

- defining the USDM artifact format itself;
- defining USDM coverage tool contracts;
- adding `usdm_covers` metadata to implementation Specifications;
- proving implementation correctness;
- making architecture component, use-case, parser, resolver, projection, or storage design decisions;
- creating Work Items or Tasks for USDM migration planning.

If row extraction exposes an undecided design choice, do not resolve it inside USDM authoring. Mark the candidate as `needs_decision`.

## Required source rules

Before using this workflow, read the current USDM authoring standard:

```text
product/records/spec/design-records/authoring-standards/usdm-authoring.md
```

Read the USDM artifact and coverage specs only for format, placement, ID, row, and tooling rules.

Do not use the USDM spec area itself as requirement source material unless the task explicitly derives requirements for USDM tooling or USDM artifacts.

## Inputs

A USDM authoring session supplies:

| input | meaning |
|---|---|
| target USDM branch | The USDM topic tree or branch being authored. |
| source specs | Product or app Specifications used as requirement source material. |
| excluded specs | Specs that may be read for rules or context but not as requirement source material. |
| existing USDM records | Existing rows and row IDs when updating a branch. |
| write permission state | Whether the user has approved file creation or update. |

## Topic split rule

Split USDM topics by Product-owned implementation concern.

Use this rule:

```text
source Specification meaning
  ×
coverage-friendly implementation responsibility
  =
USDM topic
```

Do not split topics by source file alone.
Do not split topics by future architecture components.
Do not use architecture component names unless the source Specification fixes those names.

A good topic boundary lets one or more implementation Specifications claim coverage without guessing the source intent.

## Normal flow

```text
Phase 1: requirement candidate inventory
  -> Phase 2: clarification and classification
  -> Phase 3: compact user review package
  -> user approval
  -> USDM file authoring
  -> USDM validation and coverage check
```

Do not jump from source specs directly to final USDM rows.

## Phase 1: requirement candidate inventory

Phase 1 owns broad extraction for omission prevention.

Read the target source specs and extract all requirement-like statements.

Include statements that look like:

- implementation constraints;
- format or placement rules;
- discovery rules;
- validation-relevant rules;
- path identity rules;
- relation or reference rules;
- ownership boundaries that affect implementation behavior;
- explicit non-goals that implementation must reject or avoid as Product-level behavior.

Do not assign final `RNNN` row IDs in Phase 1.
Use candidate IDs such as `C001`, `C002`, and so on.

Use this working table internally or in a detailed handoff:

| candidate id | source spec | source statement | candidate type | initial target topic | memo |
|---|---|---|---|---|---|

Candidate type values:

| value | meaning |
|---|---|
| `requirement_like` | Likely implementation-facing requirement. |
| `constraint_like` | Constraint that may become a requirement row. |
| `validation_like` | Validation or invalidity behavior. |
| `boundary_like` | Boundary that may imply support, rejection, scope, or exclusion behavior. |
| `concept_only` | Concept or vocabulary without direct implementation requirement. |
| `rationale_only` | Reasoning or historical explanation. |
| `unclear` | Ambiguous source meaning or scope. |

## Phase 2: clarification and classification

Phase 2 owns conversion from broad candidates into USDM-ready rows.

Classify every candidate against the USDM authoring standard.

Use these dispositions:

| disposition | meaning |
|---|---|
| `row` | Candidate should become a USDM requirement row. |
| `drop_concept_only` | Candidate only defines a concept, ownership area, or document boundary. |
| `drop_rationale_only` | Candidate only gives rationale, evidence, or history. |
| `merge` | Candidate duplicates or overlaps another row. |
| `split` | Candidate contains multiple implementation requirements. |
| `needs_clarification` | Candidate is unclear or too broad to row safely. |
| `needs_decision` | Candidate requires a design decision before USDM authoring. |
| `postpone` | Candidate likely belongs to another USDM branch. |

A valid USDM row states one concrete implementation-facing requirement or constraint.

The row must make clear what the implementation must:

- support;
- reject;
- preserve;
- derive;
- scope;
- expose;
- treat as invalid.

Do not only restate a concept, ownership area, design intent, or document boundary.

Use precise nouns from the source Specification.
Do not add architecture component names unless the source Specification requires those names.
Do not include public diagnostic wording unless the source Specification requires exact wording.
Do not include task status, work progress, or implementation-package names.

Use this working table internally or in a detailed handoff:

| candidate id | disposition | target topic | proposed row text or action | reason |
|---|---|---|---|---|

## Phase 3: compact user review package

Phase 3 owns review compression.

For user review, output only the fields needed for judgment.

Group rows by USDM topic or source spec topic.

Use this table shape:

| id | requirement | verdict | note |
|---|---|---|---|

Use these verdict values:

| verdict | meaning |
|---|---|
| `keep` | Row can be written as-is. |
| `revise` | Requirement is valid, but wording should change. |
| `drop` | Candidate should not become a USDM row. |
| `split` | Candidate contains multiple requirements. |
| `merge` | Candidate should be merged into another row. |
| `pending` | User judgment, source clarification, or design decision is needed. |

Do not include the full candidate inventory, classification matrix, source statement dumps, or long rationale in the user review package unless the user asks for them.

Keep source evidence and detailed classification in working notes, not in the compact review table.

## Row ID assignment

Assign final `RNNN` row IDs only after Phase 2 classification.

For new records:

- number rows sequentially from `R001` within each USDM requirement record;
- do not include dropped candidates in row numbering;
- do not leave numbering gaps to preserve candidate IDs.

For updates:

- preserve existing row IDs when the requirement meaning remains the same;
- append new rows after existing rows when practical;
- do not renumber existing rows after coverage metadata may exist;
- do not reuse a removed row ID for a different requirement.

## Write gate

Do not create or update USDM files until the user approves the compact review package or explicitly authorizes writing.

When writing is approved:

- create or update only the target USDM branch;
- use the MVP USDM artifact format;
- use canonical `spec:` refs in `## Requirements: <spec ref>` headings;
- keep one Markdown table under each source-spec requirements section;
- keep notes short and limited to clarification or temporary limitations.

## Validation gate

After USDM files are written, validate with the available USDM tools.

Preferred checks:

```text
@usdm-mcp.validate_usdm
@usdm-mcp.check_usdm_coverage
```

When MCP is unavailable, use the standalone USDM tools if the session has command execution.

Before USDM records exist, zero records and zero requirements is a valid empty state.
After USDM records exist, uncovered requirements are expected until implementation Specifications intentionally declare `usdm_covers`.

Do not add `usdm_covers` only to hide uncovered rows.

## Output discipline

For planning or handoff prompts, include:

- target USDM branch;
- source specs;
- excluded specs;
- provisional topic tree;
- expected review table shape;
- write gate;
- validation gate;
- prohibitions.

For user review, output only:

- topic heading;
- `id`;
- `requirement`;
- `verdict`;
- `note`.

For final completion after writing, report:

- files created or updated;
- validation command or MCP operation used;
- validation result;
- coverage result;
- expected uncovered state, if applicable.

## Prohibitions

Do not:

- create Work Items or Tasks as part of row extraction;
- restart or continue unrelated cancelled work;
- make Domain component input/output decisions;
- use USDM format specs as source requirements for non-USDM topics;
- add `usdm_covers` during requirement row authoring;
- stage, commit, or push unless explicitly requested;
- run broad repository-wide search when target specs are known;
- claim coverage proves implementation correctness.
