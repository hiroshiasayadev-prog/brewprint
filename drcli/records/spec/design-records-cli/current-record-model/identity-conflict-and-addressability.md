# Concept: Identity conflict and addressability

- **id**: `spec:drcli.design_records_cli.current_record_model.identity_conflict_and_addressability`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.current_record_model`
- **usdm_covers**:
  - usdm:drmcp.mcp_capabilities.validation.validation_scope_and_subjects#R014,R017-R018

## What this is

Defines repository-wide identity claim aggregation, identity conflicts, unique addressability, and operation-dependent selectability.

## Concept model

| concept | definition |
|---|---|
| Identity claim | One current artifact candidate associated with one canonical current identity. |
| Claim set | Every identity claim for one canonical current identity across the resolved current source set. |
| Identity conflict | A claim set containing more than one current artifact candidate. |
| Uniquely addressable current record | The admitted current record produced by a claim set containing exactly one candidate. |
| Addressable-record set | Every uniquely addressable current record in the complete current-record state. |
| Selectable current record | A uniquely addressable current record included by one selected scope and any operation-specific eligibility rule. |

## Rules

### Repository-wide claim aggregation

DRCLI aggregates current artifact candidates by canonical identity across every resolved current source.
Claim aggregation is repository-wide and does not change with an operation's selected scope.

| claim-set size | repository-wide state |
|---|---|
| `0` | No current record exists for the canonical identity. |
| `1` | The candidate becomes one admitted and uniquely addressable current record. |
| `2 or more` | An identity conflict exists and no current record is admitted for that identity. |

DRCLI does not select a winner from an identity conflict.
Source order, filesystem order, records-root discovery order, and content quality do not establish precedence.

### Addressability

One canonical current identity addresses at most one current record.
An exact canonical ref resolves only when the identity has one uniquely addressable current record.
A malformed, unresolved, or conflicted ref does not resolve to a current record.
DRCLI does not repair, complete, normalize, or infer an exact ref during addressability evaluation.

Nonidentity artifact-contract violations do not remove repository-wide addressability.
A nonconforming admitted record remains in the addressable-record set.

### Selectability

Selectability is derived after repository-wide addressability.
A record is selectable for an operation only when all applicable conditions hold:

1. The record is uniquely addressable.
2. The record belongs to the selected current-record scope.
3. The record satisfies any additional eligibility rule owned by the operation Specification.

Scope exclusion does not change repository-wide addressability.
Operation-specific ineligibility does not create an identity conflict or remove the admitted record.

### Partial admission

Admission is evaluated independently for each canonical identity.
An unadmitted source does not invalidate unrelated uniquely addressable records.
An identity conflict does not invalidate records under other canonical identities.

A configuration or execution failure may prevent construction of a trustworthy complete current-record state.
Such a failure belongs to operation failure handling rather than normal partial admission.

### Conflict and failure provenance

DRCLI retains the repository-root-relative path of each identity-conflict member.
Broad validation identifies a conflict by canonical identity and every member path.
Broad validation identifies each unadmitted source by its repository-root-relative path.

A repository-relative or absolute source path is not a detailed record selector.
A source must become one uniquely addressable current record before exact-ref validation can select it.

## Boundary

| concern | owner |
|---|---|
| Candidate identity claims | `spec:drcli.design_records_cli.current_record_model.artifact_candidate_and_admission`. |
| Repository-wide claim aggregation | This Specification. |
| Identity conflict and no-winner rule | This Specification. |
| Uniquely addressable current record | This Specification. |
| Scope-dependent selectability | This Specification and the selected scope Specification. |
| Operation-specific eligibility | The consuming operation Specification. |
| Conflict and unadmitted-source result projection | Validation result Specifications. |
| Canonical reference grammar | Product artifact-reference Specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.current_record_model.artifact_candidate_and_admission` | Supplies one identity claim per current artifact candidate. |
| `spec:drcli.design_records_cli.current_record_model.current_record_scopes` | Determines scope inclusion after unique addressability. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for canonical artifact refs. |
