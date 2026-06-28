# PRODUCT-ADR-SPEC-002: Generate portable Design Records package through a PRODUCT-owned rebuild command

- **status**: accepted
- **date**: 2026-06-26
- **depends_on**:
  - PRODUCT-ADR-SPEC-001
- **supersedes**: []
- **migrated_to_spec**: null

## Context

PRODUCT owns the Design Records semantics distributed to DRMCP.

The portable package copies `product/records/spec/design-records/`.
The producer rewrites canonical refs into the `spec:design_records` namespace.

The generation mechanism must preserve PRODUCT authority.
The mechanism must also avoid stale or partially generated package trees.

The repository currently builds `bin/design-records-mcp.exe`.
The bundled package therefore resolves to the sibling path `bin/design-records/`.

## Decision

Use one PRODUCT-owned Python command as the public generation entrypoint:

```text
product/src/tools/generate_design_records_package.py
```

The command uses these fixed paths:

| concern | path or rule |
|---|---|
| authoritative source | `product/records/spec/design-records/` |
| repository-local executable | `bin/design-records-mcp.exe` |
| generated package | `bin/design-records/` |
| root resolution | Resolve the repository root from the command file location. |
| working directory | Do not depend on the process working directory. |

The public command accepts no source or destination override.
Internal test seams may supply temporary roots or destinations.

Each invocation performs a full rebuild:

1. Generate a complete temporary tree from the authoritative source.
2. Copy the whole source tree.
3. Apply only the accepted canonical spec-ref prefix rewrite.
4. Complete the operational checks defined by follow-up work.
5. Move the existing destination to a backup path.
6. Move the completed temporary tree to `bin/design-records/`.
7. Delete the backup after successful replacement.
8. Restore the backup when replacement fails and restoration remains possible.

The generated package is a derived artifact.
Do not commit `bin/design-records/` as a second maintained source tree.

`scripts/verify.bat` must invoke the public generator during the repository-local distribution build.
An operational generation failure fails the build.
Semantic warnings do not fail the build.

Manual edits under `bin/design-records/` have no semantic authority.
The next generation may discard those edits.
The generator never synchronizes generated edits back to PRODUCT source files.

## Rationale

A full rebuild removes stale files.
It also avoids incremental synchronization state.

Temporary generation prevents an incomplete tree from becoming the current package.

PRODUCT-owned placement keeps generation policy with the semantic owner.

Fixed public paths prevent alternate trees from appearing authoritative.

Build integration prevents stale package distribution.

An internal path seam supports isolated tests without widening the public command contract.

## Rejected alternatives

| alternative | reason rejected |
|---|---|
| Watcher, hook, or incremental synchronization | Adds synchronization state and deletion-miss risk. |
| Direct generation into the current destination | Can leave a missing or partial package after failure. |
| Shared generator under `tools/` | No current cross-owner reuse requirement justifies shared ownership. |
| Commit the generated package tree | Creates a second maintained tree and allows source-package drift. |
| Public source and destination arguments | Allows unsupported inputs or layouts to resemble the supported package. |
| Manual generation before the distribution build | Allows execution omissions and stale package distribution. |
| Move the executable into `bin/drmcp/` | Expands this decision into an unrelated distribution-layout migration. |

## Consequences

- T05 defines checks, warning emission, failures, exit behavior, cleanup, and concurrent-run handling.
- T06 implements the generator and adds its invocation to `scripts/verify.bat`.
- Automated tests use temporary paths through internal functions.
- Integration verification writes the real `bin/design-records/` destination.
- Consumers treat `bin/design-records/` as generated input, not semantic authority.
- Future shared tooling or configurable public paths require a separate decision.
- A manifest, registry, or package-versioning model also requires a separate decision.

## Evidence

- `PRODUCT-ADR-SPEC-001` establishes PRODUCT semantic ownership.
- `DRMCP-ADR-MCP-001` defines the package as a distribution of PRODUCT-owned semantics.
- `PRODUCT-TASK-SPEC-013-01` fixes the whole-tree source boundary.
- `PRODUCT-TASK-SPEC-013-02` fixes the package interface and canonical ref rewrite.
- `PRODUCT-TASK-SPEC-013-03` fixes the non-blocking warning-handling boundary.
- `PRODUCT-WORK-SPEC-013` assigns generation and tooling placement to T04.
- Current repository layout places the executable at `bin/design-records-mcp.exe`.
- Current `scripts/verify.bat` builds that executable and does not yet invoke the generator.
