# PRODUCT-REQ-SPEC-013: Inventory Brewprint design-governance terms

- **id**: PRODUCT-REQ-SPEC-013
- **status**: accepted
- **date**: 2026-07-04
- **source_refs**:
  - PRODUCT-REQ-SPEC-011
  - PRODUCT-TASK-SPEC-025-11
  - spec:product.design_records.authoring_standards.task_authoring

## Requirement

Build a machine-readable corpus of Brewprint design-governance terms and phrases from existing artifacts.

The corpus must preserve observed usage before defining, classifying, approving, replacing, or consolidating terms.

The result must provide evidence for later work on shared concepts, conflicting meanings, qualified terms, deprecated wording, and domain use cases when the corpus justifies that work.

### Gathering criterion

Extract a term or phrase when differing interpretations could change at least one of the following:

- artifact responsibility;
- ownership or authority;
- the authoritative or normative source;
- permitted write target or operation;
- workflow routing or Task graph behavior;
- lifecycle or completion meaning;
- acceptance, review, verification, or finding state;
- provenance, dependency, projection, or derived relation;
- responsibility, scope, or stop boundary;
- precedence when artifacts or statements conflict.

Extraction units may include:

- noun phrases;
- verb-object phrases;
- state phrases;
- relation or consequence statements.

Do not extract a phrase solely because it:

- uses a definition-shaped sentence such as `A is B`;
- names a specific implementation component;
- describes ordinary work without governance consequences;
- appears frequently;
- uses unusual English.

Each observation represents one combination of:

```text
term or phrase
+ observed meaning
+ source artifact
+ semantic consequence
```

Record separate observations when the same phrase has different meanings or consequences.
Repeated equivalent usage within one source artifact may be represented by one observation.

Do not normalize synonyms, merge meanings, judge correctness, or define the term during gathering.

### Observation schema

Store observations as JSON Lines using schema version `bp-wide-term-observation-v1`.
Each line contains one independent observation.

Required authored fields:

```json
{
  "schema_version": "bp-wide-term-observation-v1",
  "observation_id": "PRODUCT-TASK-SPEC-XXX-YY-O0001",
  "source_path": "product/records/spec/example.md",
  "source_location": "## Section heading",
  "term_or_phrase": "canonical Specification",
  "observed_meaning": "The Specification that owns the currently valid contract.",
  "semantic_consequence": "Other artifacts must not redefine the current contract locally.",
  "inclusion_trigger": "authority",
  "confidence": "clear"
}
```

Allowed `inclusion_trigger` values:

- `authority`;
- `ownership`;
- `artifact_responsibility`;
- `operation`;
- `routing`;
- `lifecycle`;
- `completion`;
- `relation`;
- `boundary`;
- `validation`;
- `other`.

Allowed `confidence` values:

- `clear`;
- `inferred`;
- `unclear`.

Optional authored field:

```json
{
  "source_excerpt": "One short source sentence containing the observed phrase."
}
```

The excerpt must remain short and must not replace `observed_meaning` or `semantic_consequence`.

Fields derived mechanically from the source file should not be manually duplicated in every observation. These may include:

- artifact kind;
- source artifact ID;
- app namespace;
- domain namespace;
- file name;
- batch ID.

Each gathering batch must record machine-readable coverage metadata:

```yaml
schema: bp-wide-term-batch-v1
batch_id: PRODUCT-TASK-SPEC-XXX-YY
scope:
  - product/records/spec/design-records/**
files_discovered: 0
files_scanned: 0
files_failed: 0
observation_count: 0
```

## Evidence

- PRODUCT-TASK-SPEC-025-11 reorganized 219 Task-corpus findings into `skills/task-vocabulary-reference/`, but that artifact is observational and does not define Brewprint-wide terms.
- Existing artifacts repeatedly use governance-sensitive phrases such as `canonical`, `accepted`, `normative`, `owner`, `boundary`, `projection`, `materialize`, `route`, `propagate`, and `synchronize`.
- The meaning and consequence of these phrases can affect artifact authority, write ownership, workflow routing, lifecycle, and completion.
- A free-form Markdown inventory would make later combination, validation, grouping, and tool-assisted analysis unreliable.
- The user accepted JSONL as the required observation format because corpus-scale extraction and later organization will require tool support.
- PRODUCT-TASK-SPEC-027-37 structurally validated all 32 accepted extraction batches.
- PRODUCT-INV-SPEC-011 concluded with 733 of 733 corrected-manifest sources accounted for.
- The concluded corpus contains 5,699 conforming observation records.
- The final coverage contains 0 `no_candidate`, 0 `failed`, 0 `missing`, and 0 `unreadable` results.
- Seven readable sources changed after manifest creation and remained valid under the accepted drift contract.
- The observations remain unclassified and unnormalized.
- PRODUCT-WORK-SPEC-027 owns the completion judgment for this Requirement.

## Required Outcome

- The corpus scope selected by the downstream Work Item is fully scanned.
- Every assigned source file has extracted observations or an explicit no-candidate result.
- Every observation conforms to `bp-wide-term-observation-v1`.
- Every gathering batch records `bp-wide-term-batch-v1` coverage metadata.
- Each observation preserves the observed phrase, meaning, semantic consequence, inclusion reason, source location, and confidence.
- Observations remain unclassified and unnormalized source evidence.
- The result is suitable for machine-assisted combination, validation, grouping, and later use-case extraction without reinterpreting free-form logs.

## Explicitly Excluded Scope

- Defining terms or selecting one preferred meaning.
- Designing a final vocabulary taxonomy or classification model.
- Approving, prohibiting, replacing, or consolidating wording.
- Merging synonyms or resolving conflicting meanings.
- Projecting terms into a Specification, skill, authoring guide, or validator.
- Bulk rewriting existing artifacts.
- Defining Task-type use cases or Task-type controlled vocabulary.
- Inventorying general application-architecture terminology without a design-governance consequence.

## Boundary

This Requirement owns corpus-wide candidate discovery, observed-use evidence, JSONL observation shape, and coverage accountability.

The downstream Work Item owns the exact corpus scope, batch partitioning, Task graph, output placement, direct JSONL observation authoring, coverage Evidence, and execution route.

A lightweight inventory authoring interface and structural validator may be implemented as a required support mechanism when direct JSONL authoring would otherwise risk malformed or incomplete evidence.

Aggregation, semantic-analysis tooling, classification, definition, preferred wording, retirement, and domain use-case extraction belong to later work only when the investigation result justifies them.
