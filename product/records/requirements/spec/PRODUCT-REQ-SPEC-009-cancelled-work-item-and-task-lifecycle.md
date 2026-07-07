# PRODUCT-REQ-SPEC-009: Cancelled Work Item and Task Lifecycle

- **id**: PRODUCT-REQ-SPEC-009
- **status**: accepted
- **date**: 2026-07-03
- **source_refs**: []

## Requirement

Work Item and Task lifecycles must support a terminal `cancelled` status.

The status must represent an intentional stop before the owned completion condition is satisfied.

Cancelling a Work Item must also cancel every owned Task whose current status is `not_started`, `in_progress`, or `blocked`.

## Evidence

- `blocked` represents a temporary stop caused by a dependency or external decision.
- `done` represents successful satisfaction of the owned completion condition.
- Neither status represents an intentionally abandoned Work Item or Task whose outcome remains incomplete.
- Work Items may become invalid or unnecessary before completion.
- Their unfinished Tasks need a terminal state that preserves the unsuccessful execution history.

## Required Outcome

- Add `cancelled` to the canonical Work Item status set.
- Add `cancelled` to the canonical Task status set.
- Define `cancelled` as terminal and distinct from successful completion.
- Permit cancellation before the owned completion condition is satisfied.
- When a Work Item becomes `cancelled`, change each owned `not_started`, `in_progress`, or `blocked` Task to `cancelled`.
- Preserve owned Tasks that are already `done`.

## Explicitly Excluded Scope

- Migrating existing Work Items or Tasks to `cancelled`.
- Automatically cancelling child Work Items or transitive descendants.
- Defining the framing workflow that may decide to cancel a Work Item.
- Defining concrete DRMCP, command, or implementation mechanics.

## Boundary

PRODUCT owns the canonical cancellation meaning and required Work Item-to-Task lifecycle propagation.

Subsequent design work owns transition rules, validation, Evidence requirements, synchronization procedure, tool behavior, and implementation.
