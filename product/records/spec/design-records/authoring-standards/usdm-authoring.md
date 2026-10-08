# Reference: USDM authoring

- **id**: `spec:product.design_records.authoring_standards.usdm_authoring`
- **status**: draft
- **date**: 2026-09-30
- **parent**: `spec:product.design_records.authoring_standards`

## What this is

Authoring rules for MVP USDM records and `usdm_covers` coverage metadata.

This guide explains how to split requirement topics, write requirement rows, assign stable row IDs, and declare implementation Specification coverage.

## Non-goals

- USDM format validation rules.
- Standalone USDM tool request and response contracts.
- Artifact-model or repository-layout integration.
- Migration of existing records into USDM.
- DRMCP Domain component design.
- Implementation progress tracking.

## Rules

### When to create USDM records

| condition | authoring route |
|---|---|
| A Product or app Specification contains implementation constraints that must be covered by implementation specs. | Create or update a USDM `requirement` record with a canonical `spec:` source. |
| An accepted upstream requirement has no corresponding Specification source. | Create or update a USDM `requirement` record with `source: literal`. |
| A topic needs navigation or grouping without requirement rows. | Create or update a USDM `index` record. |
| A need or gap is not yet accepted as an upstream requirement. | Use a Requirement artifact, not USDM. |
| A design choice is not decided. | Use a decision Task or ADR route, not USDM. |
| A current normative rule must be stated. | Use a Specification, not USDM. |
| Work progress or completion must be tracked. | Use a Work Item or Task, not USDM. |

### Topic splitting

| rule | level |
|---|---|
| Split USDM topics by implementation concern, not by source file alone. | MUST |
| Keep one USDM `requirement` record focused on one coherent implementation concern. | MUST |
| Use multiple `## Requirements: <title>` sections when one topic derives rows from multiple sources or distinct requirement groups. | MUST |
| Use separate USDM records when parser, validation, resolution, projection, or storage concerns need independent coverage. | SHOULD |
| Do not force a 1:1 relation between source specs and USDM records. | MUST |

Examples, not exhaustive:

| source situation | USDM split |
|---|---|
| One spec defines both parse shape and validation behavior. | Separate parse and validation topics when implementation specs cover them separately. |
| Multiple specs define one identity concept. | One identity topic may contain multiple Requirements sections with distinct titles and canonical `spec:` sources. |
| Accepted upstream requirements exist before a corresponding spec. | Use one or more coherently titled sections with `source: literal`. |
| A source spec only provides rationale. | Do not create a USDM row unless the source states an implementation requirement. |

### Requirement row writing

| rule | level |
|---|---|
| Write one implementation requirement per row. | MUST |
| Use precise nouns from the source spec. | MUST |
| Keep design choices out of the row unless the source spec already fixes them. | MUST |
| Keep public diagnostic wording out of the row unless the source spec requires exact wording. | MUST |
| Keep task instructions and progress out of the row. | MUST |
| Use `notes` only for clarification or temporary limitations. | SHOULD |
| State the concrete implementation-facing requirement or constraint. | MUST |
| Make clear what the implementation must support, reject, preserve, derive, scope, expose, or treat as invalid. | MUST |
| Do not only restate a concept, ownership area, design intent, or document boundary. | MUST |

A row should be coverable by at least one implementation Specification file.

A row should not require reading an entire conversation to understand the required implementation outcome.

### Requirement sections

| rule | level |
|---|---|
| Use `## Requirements: <title>` for every coherent requirement group. | MUST |
| Put exactly one `> source: <source>` field immediately after the H2. | MUST |
| Use either `literal` or a canonical `spec:` ref as the source value. | MUST |
| Use `literal` only when the rows are direct upstream requirements and no corresponding Specification is their source. | MUST |
| Put rows under the source that directly provides the requirements. | MUST |
| Keep each title unique within one USDM record. | MUST |
| Name the requirement content, concern, or boundary. | MUST |
| Do not use `literal` or a canonical `spec:` ref as the title. | MUST |
| Do not use physical file names, conversation names, or session names as titles. | MUST |
| Do not use titles that claim normative authority, such as `Canonical requirements` or `Official contract`. | MUST |
| Use `notes` when a row condenses wording from multiple nearby clauses in the same Specification. | MAY |

When one Specification feeds multiple implementation concerns, each USDM topic may contain distinct titled sections with the same canonical `spec:` source.

When a corresponding Specification later becomes the direct source for a literal section, replace `literal` with the canonical `spec:` ref. Preserve every existing row ID during source replacement.

### Row hierarchy and ID stability

| rule | level |
|---|---|
| Use `RNNN` for top-level rows. | MUST |
| Use `<parent row ID>-NN` for a direct child. | MUST |
| Derive a non-top-level row's parent by removing the final `-NN` segment. | MUST |
| Create the immediate parent row before authoring a child row. | MUST |
| Treat the hierarchy as pure decomposition only. | MUST |
| Do not encode optional, OR, conditional, weighted, or cardinality relations in the ID. | MUST |
| Start a newly authored child sibling set at `01`. | MUST |
| Append new row IDs after the highest existing ID in the same sibling set when possible. | SHOULD |
| Do not reuse a removed sibling ID for a different requirement. | MUST |
| Do not renumber existing rows after implementation specs have started covering them. | MUST |
| Row ID gaps are allowed when rows are removed after review. | MUST |
| Do not reuse a removed row ID for a different requirement. | MUST |

Examples, not exhaustive:

```text
R001
R001-01
R001-02
R001-02-01
R002
```

Stable hierarchical row IDs protect `usdm_covers` references and preserve decomposition identity.

Keep one decomposition tree inside one USDM requirement record. Do not use a row in another USDM record or app namespace as a parent or child. When an implementation Specification in another app owns responsibility for a row, reference the target row by its full USDM requirement ID in `usdm_covers` instead.

### Coverage authoring

| rule | level |
|---|---|
| Add `usdm_covers` to an implementation Specification after the spec intentionally covers the exact row. | MUST |
| List full USDM requirement IDs or compact row-list expressions anchored to one USDM record ID. | MUST |
| Hierarchical rows may appear directly or in same-record comma lists. | MUST |
| Existing `RNNN-RNNN` range shorthand remains top-level only. | MUST |
| Do not list bare USDM record IDs without row fragments. | MUST |
| Keep coverage file-level during the MVP. | MUST |
| Do not add direct child coverage only to hide a parent coverage failure. | MUST |
| Do not claim direct descendant coverage merely because an ancestor is directly covered. | MUST |
| Do not claim coverage from overview prose unless the overview owns the implementation contract. | SHOULD |
| Remove or correct dangling coverage when the referenced row no longer exists. | MUST |

Direct coverage means the Specification claims to cover that exact requirement row.
Effective coverage may also be derived when every direct child of an uncovered non-leaf is effectively covered.

A directly covered parent remains effectively covered even when descendants are uncovered.
Those descendants remain visible as non-blocking refinement warnings.
The warning behavior does not define an explicit optional relation.

Coverage does not prove implementation correctness.

Examples:

```markdown
- **usdm_covers**:
  - `usdm:product.foo.bar#R001`
  - `usdm:product.foo.bar#R001-01`
  - `usdm:product.foo.bar#R001-01,R001-02`
  - `usdm:product.foo.bar#R001-R005`
```

Use full IDs for isolated rows. Use compact comma lists for hierarchical siblings. Use range syntax only for top-level rows.

### Review checklist

Before marking USDM authoring complete, check these items:

- Every `requirement` record has at least one `## Requirements: <title>` table.
- Every Requirements section has one immediate source field.
- Every source is `literal` or a canonical `spec:` ref.
- Every Requirements title is non-empty and unique within the record.
- Every row has a unique hierarchical row-local ID.
- Every non-top-level row has its immediate parent in the same USDM requirement record.
- No parent-child edge crosses a USDM requirement-record or app-namespace boundary.
- Every row is phrased as an implementation requirement.
- Existing row IDs were not renumbered after title changes, source replacement, or section restructuring.
- Row ID gaps are intentional deletions, not accidental renumbering.
- Covering implementation specs use full USDM requirement IDs or compact same-record row lists.
- Hierarchical rows use direct IDs or comma lists; existing range shorthand remains top-level only.
- Blocking uncovered rows have no directly covered ancestor.
- Uncovered descendants beneath a directly covered ancestor are retained as non-blocking refinement warnings.
- No row is treated as explicitly optional unless a later accepted contract introduces that semantic.

## Authoring interface requirements

### Create

The author supplies:

- app namespace;
- USDM kind;
- topic path;
- title;
- parent USDM ID, `root`, or `-`;
- `## What this is` body;
- requirement section titles, source values, and tables when the kind is `requirement`.

The author does not supply:

- design decisions;
- implementation package names unless the source spec requires them;
- public response wording unless the source spec requires exact wording;
- task status or progress notes.

### Update

A partial update supplies only changed metadata fields, requirement sections, source fields, requirement rows, or coverage metadata.

Rules:

- Omitted metadata fields remain unchanged.
- Update `date` only when the USDM record meaning changes.
- Preserve existing row IDs when changing wording for clarity.
- Preserve existing row IDs when changing a section title or source.
- Preserve existing row IDs when moving rows between sections or splitting or merging sections.
- Create a new row when the required implementation outcome changes materially.
- Do not silently renumber rows to make tables look cleaner.
- Do not fill row ID gaps by reusing removed row IDs.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.authoring_standards` | Parent Index. |
| `spec:product.design_records.authoring_standards.writing_standard` | Design record prose rules. |
| `spec:product.design_records.authoring_standards.artifact_boundary` | Artifact selection guidance. |
| `spec:product.design_records.usdm` | USDM overview. |
| `spec:product.design_records.usdm.artifact_format` | USDM record format rules. |
| `spec:product.design_records.usdm.coverage_format` | Coverage metadata rules. |
| `spec:product.design_records.usdm.coverage_tools` | Standalone USDM tool contracts. |
| PRODUCT-REQ-SPEC-015 | Original MVP source requirement. |
| PRODUCT-REQ-SPEC-016 | Hierarchical decomposition and derived-coverage source requirement. |
| PRODUCT-ADR-SPEC-020 | Hierarchical row and effective-coverage decision. |
| PRODUCT-WORK-SPEC-029 | Original MVP source Work Item. |
| PRODUCT-WORK-SPEC-031 | Hierarchical decomposition design Work Item. |
