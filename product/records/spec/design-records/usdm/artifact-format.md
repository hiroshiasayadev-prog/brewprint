# Contract: USDM artifact format

- **id**: `spec:product.design_records.usdm.artifact_format`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:product.design_records.usdm`
- **contract_class**: `format`

## What this is

This contract defines the MVP format for USDM records.

The format covers USDM placement, record kinds, IDs, metadata, required sections, requirement tables, row IDs, and full requirement IDs.

## Current contract

A USDM record is valid when it is placed under the MVP USDM directory, uses an accepted USDM kind, carries required visible metadata, and satisfies the required section shape for that kind.

USDM records are independent auxiliary artifacts in the MVP. They are not fully integrated into the canonical artifact model or repository layout yet.

## Rules

### Placement

| rule | level |
|---|---|
| MVP USDM records live under `<app>/records/usdm/`. | MUST |
| USDM tools must discover USDM records under each app namespace's `records/usdm/` tree. | MUST |
| Physical paths are repository locations, not canonical USDM IDs. | MUST |
| Full artifact-model and repository-layout integration is deferred. | MUST |

### USDM kind set

| kind | status | intent |
|---|---|---|
| `index` | accepted | Navigation or grouping record for a USDM topic area. |
| `requirement` | accepted | Requirement table record containing USDM requirement rows. |

No other USDM kind is accepted in the MVP.

### H1 and metadata

Each USDM record must use exactly one real ATX H1 outside fenced code blocks.

The H1 format is:

```markdown
# USDM <kind>: <Title>
```

The visible metadata block must immediately follow H1.

| marker | required | applies to | meaning |
|---|---:|---|---|
| `- **id**:` | yes | all USDM records | Canonical USDM record ID. |
| `- **status**:` | yes | all USDM records | USDM record lifecycle state. |
| `- **date**:` | yes | all USDM records | Creation or latest substantive update date. |
| `- **kind**:` | yes | all USDM records | Accepted USDM kind. |
| `- **parent**:` | yes | all USDM records | `root`, `-`, or a parent USDM record ID. |

### USDM record ID grammar

| item | grammar |
|---|---|
| USDM record ID | `usdm:<app_namespace>.<path.to.topic>` |
| app namespace segment | Existing app namespace. |
| topic path segment | Dot-separated lowercase topic path. |
| allowed segment characters | Lowercase ASCII letters, digits, and underscore. |

The metadata `id` value must match the record's canonical USDM ID.

The topic path should describe the requirement topic, not the physical filename.

### Required section matrix

| section or field | `index` | `requirement` |
|---|---:|---:|
| `## What this is` | required | required |
| `## Requirements: <title>` | prohibited | one or more required |
| Immediate `> source: <source>` field | prohibited | required once per Requirements section |

USDM `index` records intentionally remain thin in the MVP. Related requirement discovery belongs to tools, not duplicated index prose.

### Requirement section shape

Each requirement section uses this shape:

```markdown
## Requirements: <title>
> source: <source>

| id | requirement |
|---|---|
| R001 | ... |
```

The source field must immediately follow the H2 without an intervening blank line.

The source value uses one of these forms:

| source form | meaning |
|---|---|
| `literal` | The rows are direct upstream requirements and no corresponding Specification is the source. |
| canonical `spec:` ref | The referenced Specification directly provides the requirements represented by the rows. |

The title identifies the coherent requirement group for human readers. The title is not a source identity or row identity.

Title rules:

| rule | level |
|---|---|
| The title must be non-empty and unique within one USDM record. | MUST |
| The title must identify the requirement content, concern, or boundary. | MUST |
| The title must not equal `literal` or a canonical `spec:` ref. | MUST |
| The title must not be a physical file name, conversation name, or session name. | MUST |
| The title must not claim normative authority through names such as `Canonical requirements`, `Official contract`, or equivalent authority-bearing labels. | MUST |
| A title change must not change existing row IDs. | MUST |

MVP static validation enforces title presence, uniqueness, and separation from source values. The other content-based naming restrictions remain authoring and review rules.

### Requirement table shape

Each `## Requirements: <title>` section must contain one Markdown table after its source field.

The table columns are:

| column | required | meaning |
|---|---:|---|
| `id` | yes | Row-local requirement ID. |
| `requirement` | yes | Requirement statement. |
| `notes` | no | Clarification, source nuance, or temporary limitation. |

A USDM `requirement` record may contain multiple `## Requirements: <title>` sections.

Multiple requirement sections allow one USDM topic to derive rows from multiple sources. They also allow one source to feed separate parse, validation, resolution, or projection topics.

### Requirement row ID grammar

| item | grammar |
|---|---|
| row-local requirement ID | `RNNN` |
| first row ID | `R001` for newly authored records. |
| numbering scope | One USDM `requirement` record. |
| numbering rule | Strictly increasing when initially authored; gaps are allowed after reviewed row removal. |
| full requirement ID | `<usdm record id>#<row-local requirement ID>` |

Example:

| component | value |
|---|---|
| USDM record ID | `usdm:drmcp.design_records.identity` |
| row-local requirement ID | `R003` |
| full requirement ID | `usdm:drmcp.design_records.identity#R003` |

### Requirement row semantics

Each row states one implementation requirement provided by the containing section's source.

For `source: literal`, the row itself records the direct upstream requirement. For a canonical `spec:` source, the row normalizes a requirement directly provided by that Specification.

Rows should be small enough to map to implementation Specification coverage.

Rows must not contain component design, implementation package names, public diagnostic wording, or task instructions unless the source requires those details.

When a corresponding Specification later becomes the direct source for a literal section, the source field may be replaced with its canonical `spec:` ref. Source replacement, title changes, section moves, section splits, and section merges must preserve existing row IDs.

## Validation rules

| condition | severity |
|---|---|
| USDM record outside `<app>/records/usdm/` during MVP discovery | Error. |
| H1 kind outside `USDM index` or `USDM requirement` | Error. |
| Missing required metadata | Error. |
| Metadata `kind` not equal to the H1 USDM kind | Error. |
| Metadata `id` not matching `usdm:<app_namespace>.<path.to.topic>` grammar | Error. |
| Requirement record without a `## Requirements: <title>` section | Error. |
| Requirements title is empty, duplicated within one record, or equal to a source value | Error. |
| Requirement section source field is missing or does not immediately follow the H2 | Error. |
| Requirement section source is neither `literal` nor a canonical `spec:` ref | Error. |
| Requirement section contains more than one source field | Error. |
| Requirement table missing `id` or `requirement` column | Error. |
| Duplicate full requirement ID | Error. |
| Malformed row-local requirement ID | Error. |
| Row-local ID gaps after reviewed row removal | Allowed. |
| `index` record containing `## Requirements: <title>` | Error. |

## Errors

| condition | handling |
|---|---|
| Source file cannot be read as text. | Report a source read diagnostic. |
| Parser cannot find a real H1. | Report a format diagnostic and continue scanning other records. |
| Record identity cannot be determined. | Exclude the record from requirement ID construction and report a format diagnostic. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.usdm` | Parent overview. |
| `spec:product.design_records.usdm.coverage_format` | Defines coverage metadata from implementation Specifications to USDM requirement IDs. |
| `spec:product.design_records.usdm.coverage_tools` | Defines standalone tool behavior over this format. |
| PRODUCT-REQ-SPEC-015 | Source requirement. |
