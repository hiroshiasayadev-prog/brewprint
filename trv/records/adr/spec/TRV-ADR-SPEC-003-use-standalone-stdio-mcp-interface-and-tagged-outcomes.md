# TRV-ADR-SPEC-003: Use standalone stdio MCP interface and tagged outcomes

- **status**: accepted
- **date**: 2026-07-02
- **depends_on**:
  - TRV-ADR-SPEC-001
  - PRODUCT-ADR-SPEC-017
- **supersedes**: []
- **migrated_to_spec**: null

## Context

TRV requires one external interface for AI-agent callers.
The first delivery is a standalone application rather than current DRMCP integration.

The external contract must distinguish completed semantic evaluation from failures that prevent evaluation.
Per-Task failures must remain visible without terminating the MCP server process.

The MCP adapter projects application outcomes.
The adapter must not redefine PRODUCT semantics or application-owned outcome construction.

## Decision

Expose TRV as a standalone MCP server over stdio.

Use these application identities:

- server name: `task-responsibility-validator`;
- tool name: `validate_task_responsibility`.

Use one public input envelope.
The envelope contains one required string field named `task_path` and no additional properties.

Use `outcome` as the external outcome discriminator.
Use exactly three external outcome classes:

| outcome | projected content |
|---|---|
| `semantic_evaluation` | Criterion results, derived `overall_compliant`, and `human_action_required`. |
| `structural_precondition_failure` | One failure category and one concise message. |
| `execution_failure` | One failure category, one concise message, and attempt count when model invocation began. |

Return every Per-Task failure as a structured tool result.
A Per-Task failure must not terminate the MCP server process.

Use normal MCP initialization negotiation for the protocol version.
Do not fix the protocol version as TRV application identity.

The MCP adapter must project the application outcome without recalculation, synthesis, completion, or reinterpretation.

## Rationale

A dedicated server and tool identity make the standalone responsibility explicit.
Stdio preserves a narrow MCP application boundary without adding a TRV-owned HTTP server.

One tagged envelope keeps semantic, structural, and execution outcomes distinguishable.
Structured failure results preserve server availability across independent Task invocations.

Application-owned outcomes already contain the validation meaning.
Transport projection must preserve that meaning rather than create a second evaluator.

Normal protocol negotiation avoids coupling application identity to one MCP protocol revision.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Integrate the tool into current DRMCP. | Current DRMCP integration requires separate future authority. |
| Expose a TRV-owned HTTP server. | HTTP ownership is outside the accepted standalone application boundary. |
| Return one untagged result shape. | Callers could not reliably distinguish completed evaluation from structural or execution failure. |
| Terminate the server after a Per-Task failure. | One invalid or failed Task would disrupt unrelated later invocations. |
| Let the MCP adapter derive or repair semantic results. | Transport code would duplicate application and PRODUCT authority. |
| Fix one MCP protocol version in TRV identity. | Protocol negotiation already owns version agreement. |

## Consequences

- T05 must define the current normative request, response, and caller-visible contract.
- The public `task_path` field receives its meaning and safety boundary from TRV-ADR-SPEC-004.
- PRODUCT-ADR-SPEC-017 remains the authority for caller-owned human acceptance and rejection.
- The MCP adapter remains a projection boundary around application-owned outcomes.
- Per-Task structural and execution failures remain non-fatal to the server process.
- Exact JSON Schema syntax, Go types, serialization mechanics, and MCP library selection remain W004-owned.
- Retry, timeout, prompt, provider HTTP, package, symbol, and command design remain outside this ADR.

## Evidence

- TRV-TASK-SPEC-001-02 D-002 selected a dedicated stdio MCP interface.
- TRV-TASK-SPEC-001-02 D-009 selected the server identity, tool identity, input envelope, and tagged outcomes.
- TRV-TASK-SPEC-003-02 routed D-002 and D-009 into this ADR boundary.
- TRV-ADR-SPEC-001 requires the MCP adapter to project application outcomes without changing their meaning.
- PRODUCT-ADR-SPEC-017 defines caller-owned human judgment and exception handling.
