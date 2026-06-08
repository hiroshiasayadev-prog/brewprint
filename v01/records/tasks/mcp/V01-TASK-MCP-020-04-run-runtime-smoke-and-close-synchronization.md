# V01-TASK-MCP-020-04: Run runtime smoke and close synchronization

- **id**: V01-TASK-MCP-020-04
- **status**: done
- **date**: 2026-06-05
- **work_item**: V01-WORK-MCP-020
- **source_requirement**: V01-REQ-MCP-020
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-020-03
- **outputs**:
  - runtime smoke evidence
  - V01-REQ-MCP-020 / V01-WORK-MCP-020 / task close synchronization

## Goal

Run runtime smoke for the new metadata field replacement operation and synchronize close status for the related workflow artifacts.

## Work

- Run the relevant Go test set after implementation is complete.
- Run an MCP runtime smoke that updates only a task `status` field through the new field-level metadata replacement operation.
- Confirm unspecified metadata fields remain preserved after the update proposal / accept flow.
- Record evidence in this task.
- Synchronize final statuses for `TASK-MCP-020-*`, `V01-WORK-MCP-020`, and `V01-REQ-MCP-020` when acceptance criteria are met.

## Done condition

- Runtime smoke passes.
- Test commands and results are recorded.
- Related workflow artifacts are status-synchronized.
- No unexpected repository-wide validation errors are introduced by this work.

## Verification

- Run targeted tests and MCP runtime smoke.
- Run Design Records validation for the affected workflow artifacts.

## Evidence
Runtime smoke PASS.

Smoke target:

- Temp root copied from the current docs tree.
- MCP server launched with Python `subprocess.Popen` using `go run ./cmd/design-records-mcp --root <temp-root>`.
- `propose_record_update` was called with:
  - id: `V01-TASK-MCP-020-04`
  - kind: `task`
  - update.type: `metadata_fields_replace`
  - metadata.status: `doing`

Result:

- initialize: success
- proposal_created: true
- proposal_id: `pw_000001`
- proposal validation: ok
- `accept_proposed_write`: written true
- accept validation: ok
- `get_record` confirmed `V01-TASK-MCP-020-04.status == "doing"` in the temp root.
- Existing metadata fields were preserved:
  - `work_item == V01-WORK-MCP-020`
  - `source_requirement == V01-REQ-MCP-020`
  - `depends_on == [V01-TASK-MCP-020-03]`
  - `outputs` preserved both runtime smoke and close synchronization entries.

Implementation tests:

- `go test ./internal/designrecords -run TestAuthoringMetadata -v`: PASS
- `go test ./internal/designrecords ./internal/designrecordsmcp`: PASS

Close synchronization:

- `V01-TASK-MCP-020-01`: done
- `V01-TASK-MCP-020-02`: done
- `V01-TASK-MCP-020-03`: done
- `V01-TASK-MCP-020-04`: done
- `V01-WORK-MCP-020`: done
- `V01-REQ-MCP-020`: accepted
