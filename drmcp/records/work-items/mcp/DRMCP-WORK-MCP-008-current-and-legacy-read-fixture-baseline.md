# DRMCP-WORK-MCP-008: Current and legacy read fixture baseline

- **id**: DRMCP-WORK-MCP-008
- **status**: in_progress
- **date**: 2026-06-28
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-REQ-MCP-001
  - DRMCP-WORK-MCP-001
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-004
  - DRMCP-WORK-MCP-005
  - DRMCP-WORK-MCP-006
  - DRMCP-WORK-MCP-007
  - DRMCP-TASK-MCP-001-08
  - PRODUCT-WORK-SPEC-014
  - spec:product.brewprint.compatibility.legacy_id_compatibility
  - spec:product.design_records.spec_format
  - spec:product.design_records.traceability
- **tasks**:
  - DRMCP-TASK-MCP-008-01
  - DRMCP-TASK-MCP-008-02
  - DRMCP-TASK-MCP-008-03

## Goal

Create the accepted fixture baseline for current-format reads and configured legacy archive fallback.

Keep current and legacy fixture roots separate while proving that neither scope leaks into the other.

## Boundary

This Work Item owns:

- a fixture coverage matrix for every relevant outcome in `DRMCP-REQ-MCP-001`;
- a current-record fixture root;
- a separate legacy archive fixture root;
- explicit current-only and configured-legacy test configurations;
- current app-aware sequential record fixtures;
- current H1-adjacent spec metadata and path-derived `spec:` ref fixtures;
- current cross-namespace relation fixtures;
- accepted V01 sequential-family fixtures;
- configured current-to-legacy relation fixtures;
- rejected identity, source-format, duplicate, root-configuration, and fallback cases;
- fixture manifests and fixture-local structural checks;
- review of current and legacy scope isolation;
- scoped validation and independent review for fixture quality.

This Work Item does not own:

- identity, format, relation, query, resolver, validation, or diagnostic design;
- PRODUCT compatibility policy changes;
- production parser, index, tool, resolver, or validator implementation;
- production behavior assertions owned by implementation Work Items;
- contract changes made only to match a fixture;
- authoring transaction fixtures.

Fixture-local checks cover structure, placement, metadata, expected classification, and intentional duplicates.
Implementation Work Items own runtime behavior assertions against these fixtures.

## Impact Scope

| ref or area | impact |
|---|---|
| `DRMCP-REQ-MCP-001` | Supplies the required fixture classes and accepted outcomes. |
| `DRMCP-ADR-MCP-001` | Governs current format, accepted legacy families, separate indexes, and archive isolation. |
| `DRMCP-WORK-MCP-003` | Supplies current discovery, spec parsing, duplicate, and index-separation cases. |
| `DRMCP-WORK-MCP-004` | Supplies listing and exact-retrieval cases. |
| `DRMCP-WORK-MCP-005` | Supplies current-first resolution and configured fallback cases. |
| `DRMCP-WORK-MCP-006` | Supplies validation, diagnostic, and path-exposure cases. |
| `DRMCP-WORK-MCP-007` | Supplies the accepted validation-work disposition and prevents stale fixture scope. |
| `PRODUCT-WORK-SPEC-014` | Supplies the accepted `V01-SPEC-*` rejection boundary. |
| DRMCP fixture directories | Receive separate current and legacy roots plus manifests. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Coverage matrix | W003-W007 and `PRODUCT-WORK-SPEC-014` | Map each accepted and rejected requirement outcome to one or more fixture cases. |
| B. Current fixture baseline | Phase A | Create current sequential, current spec, cross-namespace, and current-only configuration fixtures. |
| C. Legacy fixture baseline | Phase A | Create configured legacy roots, accepted V01 records, and current-to-legacy relation fixtures. |
| D. Rejected and isolation cases | Phases B-C | Create invalid identity, source-format, root, duplicate, disabled-fallback, and leakage cases. |
| E. Structural verification and review | Phases B-D | Verify manifests and separation, run independent review, correct, and close. |

Implementation begins only after this fixture baseline is accepted.

## Task Candidates

| candidate | scope | dependency |
|---|---|---|
| T01 | Build the requirement-to-fixture coverage matrix from W003-W007. | Accepted contract boundaries. |
| T02 | Create current-root sequential, spec, cross-namespace, and current-only configuration fixtures. | T01. |
| T03 | Create configured legacy-root, accepted V01, and current-to-legacy fixtures. | T01. |
| T04 | Create rejected identity, invalid source-format, root-error, duplicate, and disabled-fallback fixtures. | T02-T03. |
| T05 | Verify manifests and non-leakage structure, validate, review, correct, and close. | T04. |

## Completion Condition

This Work Item is complete when all of the following are true:

- every required fixture outcome has a traceable matrix row;
- current fixtures cover app-aware sequential IDs and path-derived current spec refs;
- current fixtures cover cross-namespace relations and operation without legacy roots;
- legacy fixtures cover approved V01 sequential families under configured fallback;
- fixtures reject `V01-SPEC-*`, app-prefixless IDs, physical paths, fuzzy repair, and YAML-front-matter current specs;
- fixtures cover duplicate current identity and invalid, duplicate, or overlapping roots;
- fixtures cover disabled fallback and unresolved current and legacy references;
- current and legacy roots are physically and logically separate;
- fixture structure can detect listing, validation, indexing, or authoring leakage;
- production implementation is not changed by this Work Item;
- independent review reports no blocking or major findings;
- `DRMCP-REQ-MCP-001` lists this Work Item in `work_items`;
- final evidence records the matrix, files, structural checks, review verdict, and residual limitations.

## Evidence

- `DRMCP-ADR-MCP-001`: Accepted fixture sequence and current/legacy isolation direction.
- `DRMCP-REQ-MCP-001`: Source Requirement and required fixture classes.
- `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-007`: Upstream contract and disposition owners.
- `PRODUCT-WORK-SPEC-014`: Compatibility prerequisite.
- `DRMCP-TASK-MCP-001-08`: Hub lifecycle gate for this Work Item.
- `DRMCP-TASK-MCP-008-01` records the requirement-to-fixture matrix, bounded existing-test inventory, accepted package-local fixture candidate, manifest schema, ownership split, and review prompt.
- T01 independent review returned `PASS` with no blocking, major, or minor finding.
- External T01 scoped lifecycle and targeted whitespace checks completed with the corrected all-untracked verification shape.
- `DRMCP-TASK-MCP-008-02` opened on 2026-06-28 and accepted `drmcp/src/internal/designrecords/testdata/read-baseline/` as the package-local fixture root.
- T02 materialized C01 through C14 and C17 in one authoritative manifest.
- T02 created separate PRODUCT and DRMCP current roots, app-aware sequential fixtures using active app-domain assignments, path-derived current specs, one valid Topics parent-child arrangement, and fixture-local current-only configuration declarations.
- No legacy archive, rejection, invalid-root, duplicate, overlap, unresolved, empty-`legacy_roots`, or leakage fixture was created in T02.
- Production implementation and existing Go tests remain unchanged by T02.
- Initial external scoped fixture and lifecycle verification returned `fixture_shape=OK` and `lifecycle_shape=OK`; targeted whitespace checks also returned the expected tracked exit `0` and untracked exit `1` results.
- The first T02 independent review returned `NEEDS REVISION` with F-MAJ-01, F-MAJ-02, and F-MIN-01.
- The correction moved the PRODUCT ADR from the inactive `PRODUCT` / `MCP` assignment to `PRODUCT-ADR-SPEC-901`, synchronized every relation and exact input, made manifest-internal paths fixture-root-relative, and represented `runtime_owner` consistently as a list with both accepted C07 owners.
- The initial fixture and lifecycle verification predates the corrected bytes and is not treated as corrected-fixture evidence.
- Corrected scoped `git.inspect_worktree` checks passed for all three workflow files and all eleven fixture files; only LF-to-CRLF conversion warnings were reported, and repository-wide cleanliness was not asserted.
- Corrected external verification returned `fixture_shape=OK` and `lifecycle_shape=OK` after the three findings were addressed.
- Independent scoped re-review returned `PASS`; F-MAJ-01, F-MAJ-02, and F-MIN-01 are closed with no blocking, major, or minor regression.
- `DRMCP-TASK-MCP-008-02` is `done`.
- `DRMCP-TASK-MCP-008-03` materialized L01 through L09 in the authoritative manifest with one separate configured legacy root, the five approved V01 families, one exact retrieval source, one current-to-legacy relation, and one current-miss-to-unique-legacy-source arrangement.
- T03 did not create L10 through L13, R01 through R24, `V01-SPEC-*`, invalid or overlapping roots, disabled fallback, duplicate identity, leakage fixtures, production implementation changes, or existing Go test changes.
- T03 independent review returned `PASS` with no blocking, major, or minor finding.
- The review accepted legacy lexical mapping, root separation, L01 through L09 coverage, relation and fallback boundaries, manifest non-regression, lifecycle synchronization, and scoped Git evidence.
- The review-time Python verifier remained `NOT RUN`; no review-time execution result is inferred.
- After closure synchronization, external verification returned `fixture_shape=OK` and `lifecycle_shape=OK` against the final T03 state.
- `DRMCP-TASK-MCP-008-03` is `done`.
- W008 remains `in_progress` because T04 through T05 are not complete.
- Runtime behavior, repository-local tests, and repository-wide clean status are not inferred from fixture structure.
