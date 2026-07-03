# TRV-ADR-SPEC-004: Use repository-root-relative Task-path input

- **status**: accepted
- **date**: 2026-07-02
- **depends_on**:
  - TRV-ADR-SPEC-003
- **supersedes**: []
- **migrated_to_spec**: null

## Context

TRV requires one stable identity for the Task evaluated by each invocation.
The standalone application reads persisted Task records from a configured repository root.

Accepting bodies, public IDs, or absolute paths would create additional resolution contracts.
Unbounded paths could also read content outside the configured repository.

TRV-ADR-SPEC-003 owns the public `task_path` field.
This ADR owns the field meaning and repository safety boundary.

## Decision

Interpret `task_path` as one Task-file path relative to the configured repository root.

For each accepted input, TRV must:

- resolve the path within the configured repository root;
- read the persisted Task record;
- obtain the declared `task_type` from that Task record.

Reject absolute Task paths.
Reject input that escapes the configured repository root.

Do not accept these values as alternative primary inputs:

- Task body content;
- public record ID;
- absolute Task path.

Keep these mechanics outside this ADR:

- normalization algorithm;
- separator handling;
- case handling;
- symlink behavior;
- filesystem API;
- Go path types;
- exact error-category encoding.

## Rationale

A repository-root-relative path provides one narrow identity for the persisted Task.
The path avoids adding record-index or public-ID resolution responsibilities to the standalone application.

Reading the persisted Task preserves the Task record as the invocation source.
Deriving `task_type` from that record prevents caller-selected checklist behavior.

Root containment limits filesystem access to the configured repository boundary.
Rejecting absolute paths keeps deployment-specific locations outside the public contract.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Accept the full Task body. | The caller could evaluate content that is not the persisted Task record. |
| Accept a public record ID. | The standalone application would require record-index resolution ownership. |
| Accept absolute Task paths. | The public contract would expose deployment-specific locations and weaken root containment. |
| Infer `task_type` from the path or caller input. | Checklist selection would no longer use the persisted Task declaration. |
| Define the normalization and symlink algorithm here. | Those mechanics require implementation-ready W004 design. |

## Consequences

- T05 must define the current normative Task-input and containment contract.
- Callers must supply one repository-root-relative Task path through `task_path`.
- TRV reads the persisted Task and uses its declared `task_type`.
- Task bodies and public record IDs remain unsupported as primary inputs.
- Absolute paths and root-escaping paths fail before semantic evaluation.
- W004 must define normalization, separator, case, symlink, filesystem, type, and error-encoding mechanics.
- Current DRMCP record resolution remains outside W003.

## Evidence

- TRV-TASK-SPEC-001-02 D-003 selected one repository-root-relative Task path and root containment.
- TRV-TASK-SPEC-003-02 routed D-003 into this ADR boundary.
- TRV-ADR-SPEC-003 owns the public `task_path` field.
- `spec:product.responsibility_boundary_validator` requires checklist selection from the persisted Task's declared `task_type`.
