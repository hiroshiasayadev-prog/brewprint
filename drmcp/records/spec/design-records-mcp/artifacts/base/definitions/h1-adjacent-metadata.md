# Reference: H1-adjacent metadata

- **id**: `spec:drmcp.design_records_mcp.artifacts.base.definitions.h1_adjacent_metadata`
- **status**: draft
- **date**: 2026-07-10
- **parent**: `spec:drmcp.design_records_mcp.artifacts.base`

## What this is

Defines the shared field notation and value forms used by artifact-specific H1-adjacent metadata Specifications.

This Specification does not define artifact-specific fields, values, or metadata placement.

## Current contract

Each H1-adjacent metadata field uses this visible marker form:

```text
- **<field_name>**: <value>
```

| element | rule |
|---|---|
| field marker | The marker starts with `- `. |
| field name | The field name is enclosed by `**`. |
| separator | An ASCII colon follows the closing `**`. |
| value | The field value follows the separator. |

## Value forms

Artifact-specific metadata Specifications may assign one of these forms to a field.

| form | source notation |
|---|---|
| `scalar` | `- **<field_name>**: <value>` |
| `inline_list` | `- **<field_name>**: <value_1>, <value_2>` |
| `indented_list` | A field marker followed by directly indented child list items. |

An artifact-specific metadata Specification defines which form applies to each field.
This base definition does not define value syntax or allowed values within a form.

## Requirement values

Artifact-specific metadata Specifications assign one requirement value to each field.

| requirement | rule |
|---|---|
| `mandatory` | The field appears exactly once. |
| `optional` | The field appears zero or one time. |
| `reference` | The Specification named by the field's `value format` defines whether the field is mandatory, optional, or prohibited. |

A `reference` field must name a canonical Specification ref in `value format`.
The referenced Specification owns the field's presence conditions and may also define field-specific format constraints.
A `mandatory` or `optional` field normally declares its value format directly.
It may name a canonical Specification ref only when the field accepts a complex format spanning multiple value types that cannot be stated concisely in the field table.

## Value types

Artifact-specific metadata Specifications assign one value type to each field.

| value type | meaning |
|---|---|
| `string` | A plain string value that is not interpreted as an artifact reference. |
| `ref` | A Product-defined canonical artifact reference. |
| `ref_or_literal` | A Product-defined canonical artifact reference or one field-specific literal defined by the field's `value format`. |

## Value format declarations

The `value format` cell must state the accepted value domain in a concise, implementation-checkable form.

- For `ref` and `ref_or_literal`, list every permitted artifact kind explicitly by its Product-defined record kind name.
- State additional syntax, cardinality, literal, or allowed-value constraints directly and concisely.
- Do not include conceptual meaning, lifecycle rationale, ownership explanation, or other authoring guidance.
- Do not replace an explicit artifact-kind list or concise constraint with a Specification reference.
- A canonical Specification ref may define the value format only when the field accepts a complex format spanning multiple value types that cannot be stated concisely in the field table.

## Form and type rules

- A `scalar` field contains one value of its declared value type.
- Each element of an `inline_list` or `indented_list` uses the field's declared value type.
- A `string` value is not treated as an artifact reference even when its text resembles an artifact ID or canonical ref.
- A `ref` value is interpreted as a canonical artifact reference.
- A `ref_or_literal` value uses `ref` semantics when it matches an allowed canonical reference.
- Otherwise, a `ref_or_literal` value must equal one literal permitted by the field's `value format`.
- The artifact-specific metadata Specification defines any permitted artifact kind, reference class, value syntax, or allowed value constraint.
- This Specification does not define how a `ref` or `ref_or_literal` value contributes to parsed-record relations, indexes, trees, or graphs.

## Boundary

| concern | owner |
|---|---|
| Shared field marker notation | This Specification. |
| Shared value-form names | This Specification. |
| Shared requirement values and reference behavior | This Specification. |
| Shared value-type names and form/type behavior | This Specification. |
| Metadata placement relative to H1 and H2 sections | The artifact source Specification. |
| Artifact-specific field inventory | The artifact H1-adjacent metadata Specification. |
| Field requiredness and value format | The artifact H1-adjacent metadata Specification. |
| Field meaning and allowed values | Product authority consumed by the artifact Specification. |
| Parsing, normalization, and validation behavior | Other DRMCP Specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.spec_format.document_shape` | Product authority for Specification H1-adjacent metadata shape. |
| `spec:product.design_records.authoring_standards` | Product authority for workflow-record authoring rules. |
