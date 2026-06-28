# PRODUCT-ADR-SPEC-003: Define operational checks and warning behavior for portable package generation

- **status**: accepted
- **date**: 2026-06-26
- **depends_on**:
  - PRODUCT-ADR-SPEC-002
- **supersedes**: []
- **migrated_to_spec**: null

## Context

`PRODUCT-ADR-SPEC-002` selects a full rebuild through a PRODUCT-owned Python command.

The generator builds a temporary package before replacing `bin/design-records/`.
T06 needs one contract for checks, warnings, failures, cleanup, and verification.

Semantic source defects must not block package generation.
Operational failures must prevent publication of an incomplete package.

## Decision

### Check phases

The generator performs checks in this order:

| phase | required behavior |
|---|---|
| Whole-tree copy | Copy the complete source file set into a unique temporary tree. |
| Copy check | Compare source and temporary-tree relative file paths. |
| Canonical rewrite | Apply only the accepted `spec:product.design_records` prefix rewrite. |
| Rewrite check | Confirm that no other source content changed. |
| Warning scan | Emit semantic warnings without changing generated content. |
| Publication | Replace the destination only after all operational pre-publication checks pass. |
| Publication confirmation | Confirm that the destination exists as an enumerable directory. |

The generator does not repeat the complete pre-publication scan after replacement.

### Operational success

Generation succeeds only when all required operational steps complete:

- the fixed source directory is readable;
- the temporary tree is created;
- whole-tree copy completes;
- copy completeness passes;
- canonical rewrite execution completes;
- rewrite integrity passes;
- warning scans complete operationally;
- destination replacement completes;
- the published destination exists and can be enumerated.

Semantic warning count does not affect success.

### Copy completeness

The source and generated relative file-path sets must match exactly.

The check includes Markdown and non-Markdown files.
The check detects missing and unexpected files.
Empty-directory preservation is not required.

A full rebuild and exact file-set comparison exclude stale destination files.

### Rewrite integrity

The accepted mapping is:

```text
spec:product.design_records
→ spec:design_records
```

The mapping also applies to canonical refs below that prefix.

The rewrite check confirms these rules:

- only recognized canonical refs under the accepted source prefix may change;
- content outside the accepted rewrite remains unchanged;
- rewrite execution failure is an operational failure;
- a remaining canonical source-prefix ref is a semantic warning;
- the producer does not perform additional semantic repair.

### Warning classes

The generator emits these independent semantic warning classes:

| warning class | condition |
|---|---|
| duplicate canonical ref | Multiple generated files declare the same package canonical ref. |
| unresolved internal ref | A package ref under `spec:design_records` does not resolve in the generated package. |
| external canonical ref | A canonical `spec:` ref points outside `spec:design_records`. |
| unrewritten source ref | A canonical `spec:product.design_records` ref remains after rewrite. |
| source-authoring boundary finding | Content appears app-local, wiring-specific, migration-specific, or project-tracking-specific. |

Each applicable class may emit for the same file.
Semantic warnings never block publication.
The generator does not repair, remove, filter, or reinterpret warning-producing content.

Duplicate, unresolved, external, and unrewritten checks are deterministic structural checks.
Source-authoring boundary checks are best-effort rules based on explicit reviewable patterns.
They do not use LLM inference.
A clean boundary scan does not prove semantic app independence.

Failure to detect a semantic issue is not an operational failure.
Failure to complete an intended scan because of an operational error is an operational failure.

### Exit behavior

| outcome | exit code |
|---|---:|
| Successful generation without semantic warnings | `0` |
| Successful generation with semantic warnings | `0` |
| Operational generation failure | `1` |

Warnings and errors are emitted to standard error.
Normal generation summaries may use standard output.
The first release defines no additional numeric exit-code taxonomy.

### Cleanup and recovery

| outcome | required behavior |
|---|---|
| Success | Remove the temporary tree and completed backup. |
| Failure before replacement | Remove the temporary tree and leave the existing destination unchanged. |
| Replacement failure with successful rollback | Restore the prior destination and remove unnecessary temporary artifacts. |
| Replacement failure with failed rollback | Preserve recovery-capable backup or remnants and report their paths. |
| Cleanup-only failure after successful publication | Keep generation successful and emit a non-blocking cleanup warning. |

Rollback failure is an operational failure.
Recovery-capable artifacts must not be deleted only to satisfy cleanup.
Each invocation uses non-reused temporary and backup paths.

### Concurrent execution

Concurrent generation against the same repository destination is unsupported in the first release.

The generator does not implement a lock or waiting protocol.
Callers must invoke the generator serially.
A filesystem collision becomes an ordinary operational failure.
T06 and T07 do not require concurrency tests.

### Independence

The generator uses only these fixed inputs:

- `product/records/spec/design-records/`;
- `bin/design-records/`;
- the accepted canonical prefix mapping.

The generator resolves the repository root from the command file location.
It does not use the process working directory for source or destination resolution.

The generator does not query:

- the current app namespace set;
- DRMCP runtime state;
- a host app registry;
- Brewprint process state.

Dedicated registry mocks and repository-external working-directory tests are not required.
T06 verifies the dependency boundary through implementation inspection and normal generation tests.

### T06 verification minimum

T06 must verify these cases:

1. All source files are copied with unchanged relative paths.
2. Only the accepted canonical prefix rewrite changes content.
3. Semantic findings emit warnings and retain exit code `0`.
4. Copy, rewrite, or replacement failure returns exit code `1`.
5. Replacement failure preserves or restores the previous destination when possible.
6. Successful generation cleans temporary and backup artifacts.
7. The generator does not depend on DRMCP or an app registry.
8. Repository-root resolution uses the command file location, not the process working directory.

## Rationale

Pre-publication checks prevent an incomplete temporary tree from becoming current.
A minimal post-publication check confirms that replacement produced a usable directory.

Exact file-set comparison is sufficient for whole-tree copy completeness.
It avoids dependency-graph interpretation during generation.

Structural warning classes produce reviewable output without assigning semantic correction to the producer.
Best-effort boundary scanning avoids an impossible guarantee over natural-language meaning.

Simple exit codes integrate directly with `scripts/verify.bat`.
Unsupported concurrency avoids lock lifecycle and stale-lock recovery in the first release.

## Rejected alternatives

| alternative | reason rejected |
|---|---|
| Repeat all checks after destination replacement | Duplicates work and introduces conflicting pre- and post-publication results. |
| Fail generation when semantic warnings exist | Contradicts the accepted warning boundary. |
| Check only Markdown files | Can silently omit future non-Markdown package resources. |
| Validate copy completeness by file count | Missing and unexpected files can offset each other. |
| Require complete semantic detection of app-local content | Natural-language semantic completeness is not operationally testable. |
| Use a warning-specific nonzero exit code | Common build tooling would treat warnings as failures. |
| Add a repository lock | No supported concurrent-generation use case justifies lock lifecycle complexity. |
| Require repository-external execution tests | Implementation dependency inspection is sufficient for the first release. |

## Consequences

- T06 implements this contract without changing the public CLI.
- T06 may choose internal functions and diagnostic formatting.
- T06 must preserve the warning-versus-failure boundary.
- T07 may add broader fixtures without redefining operational success.
- Cleanup warnings are non-blocking only after a valid destination is published.
- Concurrent generation remains unsupported until a separate decision adds coordination.
- Source authoring remains responsible for semantic correction.

## Evidence

- `PRODUCT-ADR-SPEC-002` selects the rebuild and replacement model.
- `PRODUCT-TASK-SPEC-013-02` defines package refs and semantic warning classes.
- `PRODUCT-TASK-SPEC-013-03` defines non-persistent warning handling.
- `PRODUCT-TASK-SPEC-013-04` assigns checks, exit behavior, cleanup, and concurrency to T05.
- `PRODUCT-WORK-SPEC-013` assigns operational checks and portability verification to T05.
- The user accepted each T05 design decision on 2026-06-26.
