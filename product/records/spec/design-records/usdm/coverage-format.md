# Contract: USDM coverage format

- **id**: `spec:product.design_records.usdm.coverage_format`
- **status**: draft
- **date**: 2026-07-08
- **parent**: `spec:product.design_records.usdm`
- **contract_class**: `format`

## What this is

This contract defines MVP coverage metadata from implementation Specifications to USDM requirement rows.

The contract lets static tools find USDM requirements that no implementation Specification claims to cover.

## Current contract

Implementation Specifications may declare H1-adjacent `usdm_covers` metadata.

Each `usdm_covers` item is a full USDM requirement ID.

Coverage is file-level in the MVP. A covering Specification claims that the Specification covers the listed requirement somewhere in the file.

## Rules

### Coverage metadata marker

The coverage marker is optional H1-adjacent metadata on Specification records.

```markdown
- **usdm_covers**:
  - usdm:<app_namespace>.<path.to.topic>#R001
```

| rule | level |
|---|---|
| `usdm_covers` values must be full USDM requirement IDs. | MUST |
| `usdm_covers` must not contain USDM record IDs without row fragments. | MUST |
| Duplicate entries inside one Specification are invalid. | MUST |
| Coverage order has no semantic meaning. | MUST |
| Coverage is file-level during the MVP. | MUST |
| Section-level coverage is deferred. | MUST |

### Coverage relation

| source | target | meaning |
|---|---|---|
| implementation Specification | USDM requirement row | The Specification claims to cover the implementation requirement. |

A coverage relation does not prove implementation correctness.

A coverage relation does not prove that the Specification wording is complete.

A coverage relation exists to prevent requirement omissions during authoring and review.

### Uncovered requirement

A USDM requirement row is uncovered when no implementation Specification lists the full USDM requirement ID in `usdm_covers`.

Static coverage checks must report uncovered requirement IDs.

### Dangling coverage

A coverage entry is dangling when a Specification lists a full USDM requirement ID that does not exist in discovered USDM requirement rows.

Static coverage checks must report dangling coverage entries.

### Coverage scan scope

| item | MVP rule |
|---|---|
| USDM requirement source | `<app>/records/usdm/` |
| covering Specification source | `<app>/records/spec/` |
| coverage metadata | H1-adjacent `usdm_covers` marker. |
| cross-app coverage | Allowed only when the full USDM requirement ID names the target app namespace. |

The MVP tools may start with explicit paths or repository root discovery. The tool contract defines request details.

## Validation rules

| condition | severity |
|---|---|
| `usdm_covers` item is not a full USDM requirement ID | Error. |
| `usdm_covers` item points to a missing USDM requirement row | Error for dangling coverage checks. |
| USDM requirement row is not covered by any Specification | Report as uncovered requirement, not as malformed USDM. |
| Duplicate `usdm_covers` item in one Specification | Error. |
| `usdm_covers` appears outside H1-adjacent metadata | Error for MVP tools. |

## Errors

| condition | handling |
|---|---|
| Covering Specification cannot be parsed enough to inspect H1-adjacent metadata. | Report a coverage scan error. |
| USDM requirement discovery fails for a configured app namespace. | Report a coverage scan error. |
| `usdm_covers` contains both valid and invalid entries. | Preserve valid entries and report invalid entries. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.usdm` | Parent overview. |
| `spec:product.design_records.usdm.artifact_format` | Defines full USDM requirement IDs. |
| `spec:product.design_records.usdm.coverage_tools` | Defines tool behavior for uncovered and dangling coverage. |
| `spec:product.design_records.spec_format.document_shape` | Defines H1-adjacent metadata shape for Specifications. |
| PRODUCT-REQ-SPEC-015 | Source requirement. |
