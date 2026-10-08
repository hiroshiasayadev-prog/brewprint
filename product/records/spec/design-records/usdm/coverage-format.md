# Contract: USDM coverage format

- **id**: `spec:product.design_records.usdm.coverage_format`
- **status**: draft
- **date**: 2026-09-30
- **parent**: `spec:product.design_records.usdm`
- **contract_class**: `format`

## What this is

This contract defines MVP coverage metadata from implementation Specifications to USDM requirement rows.

The contract lets static tools distinguish direct coverage, derived coverage, blocking uncovered requirements, and non-blocking refinement warnings.

## Current contract

Implementation Specifications may declare H1-adjacent `usdm_covers` metadata.

Each `usdm_covers` item is either a full USDM requirement ID or a compact row-list expression anchored to one USDM requirement record.

Coverage is file-level in the MVP. A listed row receives direct coverage from the Specification. Effective coverage may also be derived from complete child coverage.

## Rules

### Coverage metadata marker

The coverage marker is optional H1-adjacent metadata on Specification records.

```markdown
- **usdm_covers**:
  - usdm:<app_namespace>.<path.to.topic>#R001
  - usdm:<app_namespace>.<path.to.topic>#R001-01
  - usdm:<app_namespace>.<path.to.topic>#R001-01,R001-02
  - usdm:<app_namespace>.<path.to.topic>#R001-R005
```

| rule | level |
|---|---|
| `usdm_covers` values must be full USDM requirement IDs or compact row-list expressions anchored to one USDM record ID. | MUST |
| A row token may use the hierarchical `RNNN(-NN)*` grammar. | MUST |
| Compact comma lists may contain hierarchical row tokens from the same USDM record. | MUST |
| Existing ascending range syntax remains limited to top-level `RNNN-RNNN` tokens. | MUST |
| Hierarchical range shorthand is not defined by this contract. | MUST |
| `usdm_covers` must not contain USDM record IDs without row fragments. | MUST |
| Duplicate expanded requirement IDs inside one Specification are invalid. | MUST |
| Coverage order has no semantic meaning. | MUST |
| Coverage is file-level during the MVP. | MUST |
| Section-level coverage is deferred. | MUST |

### Direct coverage relation

| source | target | meaning |
|---|---|---|
| implementation Specification | USDM requirement row | The Specification directly claims to cover the implementation requirement. |

A direct coverage relation does not prove implementation correctness.

A direct coverage relation does not prove that the Specification wording is complete.

A direct coverage relation exists to prevent requirement omissions during authoring and review.

### Effective coverage

Coverage checks evaluate each row recursively.

| row state | effective coverage |
|---|---|
| The row has direct coverage. | Covered. Descendant coverage does not affect this row's covered state. |
| The row has no direct coverage, has children, and every direct child is effectively covered. | Derived covered. |
| The row is a leaf and has no direct coverage. | Uncovered. |
| The row has no direct coverage and at least one direct child is not effectively covered. | Uncovered. |

Direct coverage of a parent does not create direct coverage relations for descendants.

A parent without direct coverage can become derived covered through complete child coverage.
The rule is recursive, so complete leaf coverage can satisfy every uncovered ancestor in the decomposition path.

### Uncovered requirement classification

| condition | classification | coverage result |
|---|---|---|
| The row is effectively uncovered and no directly covered ancestor exists. | Blocking uncovered requirement. | Coverage check fails. |
| The row is effectively uncovered and at least one directly covered ancestor exists. | Refinement warning. | Coverage check remains non-blocking. |

A refinement warning keeps the uncovered descendant visible.
The warning does not define the descendant as semantically optional and does not add an `optional` relation type.

Static coverage checks must report both blocking uncovered requirement IDs and refinement warnings.

### Dangling coverage

A coverage entry is dangling when a Specification lists or expands to a full USDM requirement ID that does not exist in discovered USDM requirement rows.

Static coverage checks must report dangling coverage entries.

### Coverage scan scope

| item | MVP rule |
|---|---|
| USDM requirement source | `<app>/records/usdm/` |
| covering Specification source | `<app>/records/spec/` |
| coverage metadata | H1-adjacent `usdm_covers` marker. |
| cross-app coverage | Allowed when the full USDM requirement ID names the target app namespace. Cross-app coverage does not create cross-app row hierarchy. |

The MVP tools may start with explicit paths or repository root discovery. The tool contract defines request details.

Requirement decomposition remains local to the target USDM requirement record. Cross-app implementation responsibility is represented only by the direct `usdm_covers` relation in this contract; it does not alter parent-child structure.

## Validation rules

| condition | severity |
|---|---|
| `usdm_covers` item is not a full USDM requirement ID or valid compact row-list expression | Error. |
| `usdm_covers` item points to or expands to a missing USDM requirement row | Error for dangling coverage checks. |
| Effectively uncovered row with no directly covered ancestor | Report as a blocking uncovered requirement, not as malformed USDM. |
| Effectively uncovered row with a directly covered ancestor | Report as a non-blocking refinement warning. |
| Duplicate expanded `usdm_covers` requirement ID in one Specification | Error. |
| `usdm_covers` appears outside H1-adjacent metadata | Error for MVP tools. |

## Errors

| condition | handling |
|---|---|
| Covering Specification cannot be parsed enough to inspect H1-adjacent metadata. | Report a coverage scan error. |
| USDM requirement discovery fails for a configured app namespace. | Report a coverage scan error. |
| `usdm_covers` contains both valid and invalid entries. | Preserve valid expanded entries and report invalid entries. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.usdm` | Parent overview. |
| `spec:product.design_records.usdm.artifact_format` | Defines full USDM requirement IDs. |
| `spec:product.design_records.usdm.coverage_tools` | Defines tool behavior for uncovered and dangling coverage. |
| `spec:product.design_records.spec_format.document_shape` | Defines H1-adjacent metadata shape for Specifications. |
| PRODUCT-REQ-SPEC-015 | Original MVP source requirement. |
| PRODUCT-REQ-SPEC-016 | Hierarchical decomposition and derived-coverage source requirement. |
| PRODUCT-ADR-SPEC-020 | Direct, derived, and warning coverage decision. |
