# PRODUCT-INV-SPEC-011: Brewprint design-governance term inventory

- **status**: concluded
- **date**: 2026-07-04
- **trigger**: PRODUCT-REQ-SPEC-013
- **scope**: Inventory observed Brewprint design-governance terms and phrases across the selected active and targeted legacy corpus.
- **non_scope**: Term classification, definition, normalization, preferred wording, semantic aggregation, use-case extraction, and source-artifact rewriting.
- **source_refs**:
  - PRODUCT-REQ-SPEC-013
- **follow_up_candidates**:
  - Machine-assisted aggregation and semantic analysis over the PRODUCT-INV-SPEC-011 observation corpus.
- **related_requirements**:
  - PRODUCT-REQ-SPEC-013
- **related_work_items**:
  - PRODUCT-WORK-SPEC-027

## Investigation scope

This Investigation owns one bounded research question:

> Which governance-sensitive terms and phrases are observed within the selected physical path set under the PRODUCT-REQ-SPEC-013 gathering criterion?

The included physical path set is:

```text
product/records/**/*.md
drmcp/records/**/*.md
bpdsl/records/**/*.md
trv/records/**/*.md
skills/**/*.md
v01/records/adr/**/*.md
v01/records/investigations/**/*.md
```

The scope is defined by these repository paths.
This Investigation does not define what qualifies as a Design Record.

## Out of scope

- Any other path under `v01/records/**`.
- Implementation source.
- Scripts.
- Generated outputs.
- Build artifacts.
- Term classification or taxonomy design.
- Term definition or correctness judgment.
- Synonym normalization or meaning merge.
- Preferred, canonical, deprecated, or prohibited wording.
- Semantic aggregation or clustering.
- Use-case extraction.
- Source-artifact rewriting.

## Background

PRODUCT-REQ-SPEC-013 requires a machine-readable corpus of observed Brewprint design-governance usage.

Differing interpretations can change artifact responsibility, authority, ownership, write targets, workflow routing, relations, lifecycle, or completion.
The corpus must preserve observed usage before later work defines or organizes vocabulary.

PRODUCT-TASK-SPEC-027-01 selected the exact physical corpus and Evidence placement.
PRODUCT-WORK-SPEC-027 assigns this Investigation durable ownership of the research method and later factual conclusion.

## What was investigated

The Investigation gathered evidence through 32 writer-disjoint extraction batches.
The term-inventory MCP structurally validated every batch after extraction.
The following method contract governed every batch.

### Gathering criterion

Extract a term or phrase when differing interpretations could change at least one governance consequence, including:

- artifact responsibility;
- authority or ownership;
- authoritative or normative source selection;
- permitted write target or operation;
- workflow routing or Task graph behavior;
- relation or provenance meaning;
- lifecycle or completion meaning;
- acceptance, review, verification, or finding state;
- responsibility, scope, stop boundary, or conflict precedence.

Eligible extraction units include:

- noun phrases;
- verb-object phrases;
- state phrases;
- relation or consequence statements.

Frequency alone does not satisfy the gathering criterion.
Gathering must not normalize synonyms, merge meanings, judge correctness, or define terms.

Each observation represents one combination:

```text
term_or_phrase
+ observed_meaning
+ source artifact
+ semantic_consequence
```

Separate observations are required when the same phrase has different observed meanings or consequences.
Repeated equivalent usage within one source artifact may remain one observation.

### Collection schemas

| schema | purpose |
|---|---|
| `bp-wide-term-observation-v1` | One independent observed-use record. |
| `bp-wide-term-batch-v1` | One batch coverage record with per-source disposition. |

### Machine-readable Evidence contract

Machine-readable Evidence belongs under:

```text
product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/
  manifest.json
  batches/<batch_id>.json
  results/<batch_id>.observations.jsonl
  results/<batch_id>.coverage.json
```

The Evidence contract uses these rules:

- Manifest hashes and modification metadata are best-effort snapshots.
- Source changes after manifest creation are expected.
- A `changed` snapshot status does not by itself fail a batch.
- Missing or unreadable sources must appear explicitly in batch coverage.
- Each extraction batch has one independent writer.
- Each batch writer owns only its observations JSONL and coverage JSON.
- This Work Item does not create a merged master JSONL.
- This Investigation does not perform semantic aggregation.

## Findings

### Structural validation

| measure | result |
|---|---:|
| Accepted batch IDs | 32 |
| Structurally valid batches | 32 |
| Structural diagnostics | 0 |
| Corrected manifest sources | 733 |
| Assigned sources | 733 |
| Disposed sources | 733 |
| Hidden, duplicate, or extra sources | 0 |

The original T37 text referred to the first 736-source manifest.
T04 superseded that snapshot before T37 execution because generated workflow records had entered the scope.
The corrected manifest is the extraction authority and contains 733 sources.

### Result counts

| result | count |
|---|---:|
| Observation records | 5,699 |
| Sources with `observations` disposition | 733 |
| Sources with `no_candidate` disposition | 0 |
| Sources with `failed` disposition | 0 |
| `changed` source snapshots | 7 |
| `missing` source snapshots | 0 |
| `unreadable` source snapshots | 0 |
| `unchanged` source snapshots | 726 |

### Batch-family totals

| batch family | batches | sources | observation records | changed snapshots |
|---|---:|---:|---:|---:|
| PT01-PT08 | 8 | 206 | 1,334 | 2 |
| PN01-PN04 | 4 | 117 | 1,165 | 1 |
| DT01-DT05 | 5 | 80 | 795 | 1 |
| DN01-DN03 | 3 | 72 | 435 | 3 |
| TR01-TR02 | 2 | 69 | 591 | 0 |
| BP01-BP02 | 2 | 37 | 279 | 0 |
| SK01 | 1 | 40 | 401 | 0 |
| LA01-LA05 | 5 | 99 | 509 | 0 |
| LI01-LI02 | 2 | 13 | 190 | 0 |
| **Total** | **32** | **733** | **5,699** | **7** |

### Changed snapshots

The following sources changed after manifest creation:

- `product/records/tasks/spec/PRODUCT-TASK-SPEC-027-01-decide-term-inventory-execution.md`;
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-027-02-implement-term-inventory-mcp.md`;
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-027-inventory-brewprint-design-governance-terms.md`;
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-016-08-correct-authoring-guidance-source-specifications.md`;
- `drmcp/records/spec/design-records-mcp/schema/authoring-guidance-source.md`;
- `drmcp/records/spec/design-records-mcp/tools/get-authoring-guidance.md`;
- `drmcp/records/spec/design-records-mcp/tools/list-authoring-guides.md`.

Every changed source remained readable and produced one or more observations.
Changed snapshots were non-blocking under the accepted Evidence contract.

## Cross-cutting observations

- Every corrected-manifest source produced at least one observation.
- No source required a `no_candidate` or `failed` disposition.
- Structural validation established schema, ownership, coverage, and count consistency.
- Structural validation did not evaluate term meaning or semantic correctness.
- Raw Evidence remains in 32 batch-owned JSONL files and 32 coverage files.
- No merged master JSONL, semantic cluster, vocabulary classification, or use-case extraction was produced.

## Follow-up judgment candidates

Candidate: Decide whether to start separate machine-assisted aggregation and semantic-analysis work over the 5,699 observations.

Any follow-up judgment must remain separate from this concluded raw-evidence Investigation.
The follow-up may evaluate shared concepts, conflicting meanings, qualified terms, deprecated wording, and domain use cases.
This Investigation does not claim that any specific semantic category exists.

## Recommendation

A separate aggregation and semantic-analysis Work Item appears appropriate because the corpus contains 5,699 independent observations.
The follow-up should preserve the 32 batch-owned files as immutable source Evidence.

## Follow-up artifact candidates

- A new Requirement and Work Item for machine-assisted aggregation and semantic analysis of PRODUCT-INV-SPEC-011 Evidence.
- A separate Investigation for conflicting meanings and domain use cases after aggregation establishes candidate groups.

## Open questions

No unresolved structural coverage question remains.
Semantic equivalence, conflict, qualification, preferred wording, and domain use cases remain unanswered because they were outside this Investigation.
