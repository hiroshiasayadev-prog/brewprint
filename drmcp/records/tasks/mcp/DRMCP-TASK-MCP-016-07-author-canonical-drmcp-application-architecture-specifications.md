# DRMCP-TASK-MCP-016-07: Author canonical DRMCP application-architecture Specifications

- **id**: DRMCP-TASK-MCP-016-07
- **status**: done
- **date**: 2026-07-04
- **work_item**: DRMCP-WORK-MCP-016
- **task_type**: authoring
- **estimate**: 1.5d
- **depends_on**:
  - DRMCP-TASK-MCP-016-06
- **outputs**:
  - spec:drmcp.application_architecture
  - spec:drmcp.application_architecture.application_boundary_and_components
  - spec:drmcp.application_architecture.dependency_and_responsibility
  - spec:drmcp.application_architecture.runtime_and_state
  - spec:drmcp.application_architecture.failure_and_evolution

## Goal

Author the canonical DRMCP application-architecture topic tree from the accepted T03 decisions and the final ADR set.

## Work

- Create `drmcp/records/spec/application-architecture/index.md` as the Overview and topic map.
- Create `application-boundary-and-components.md` for application scope, actors, the six components, and the deferred authoring seam.
- Create `dependency-and-responsibility.md` for inward dependencies, owned and excluded responsibilities, standards placement, Guidance placement, ports, and forbidden edges.
- Create `runtime-and-state.md` for request-scoped snapshots, Current-versus-Legacy separation, Application internal collaboration, configuration, lifecycle, and deferred mutable authoring state.
- Create `failure-and-evolution.md` for request rejection, expected semantic outcomes, execution failure, downstream-local refinements, and architecture-return triggers.
- Project every applicable D-001 through D-017 decision into exactly one primary view and add direct cross-view references where needed.
- Reference the exact ADR set produced by T06.
- Preserve package, interface, type, algorithm, adapter API, library, and implementation choices as downstream details.
- Keep the D-012 and D-013 responsibility example non-binding and clearly outside the active architecture contract.
- Do not correct the existing Guidance source Specifications in this Task.

## Done condition

- The topic tree contains one Overview and exactly four focused architecture views.
- The four views cover every current T03 architecture decision applicable to W016.
- The six-component graph and inward dependency direction are normative and coherent.
- Runtime state keeps Current and Legacy separate and request-scoped.
- PRODUCT standards and Guidance responsibilities match T03.
- Failure semantics and architecture-return triggers are normative.
- Deferred authoring internals remain deferred and non-binding.
- No exact module contract or implementation design is introduced.
- The Specification set is ready for integrated review with T06 and T08 outputs.

## Verification

- PASS: All five output refs match their path-derived Specification identities.
- PASS: Every Specification uses one accepted H1 kind, visible metadata, the required parent, and no YAML front matter.
- PASS: `index.md` is an `Overview` and authoritatively declares exactly four child views.
- PASS: The `## Topics` table uses only `title`, `kind`, `ref`, and `summary`.
- PASS: Every child `parent` equals `spec:drmcp.application_architecture`.
- PASS: D-001 through D-017 have exactly one primary view.
- PASS: D-012 and D-013 remain deferred and outside the active architecture contract.
- PASS: The current contract matches `DRMCP-ADR-MCP-007`, `DRMCP-ADR-MCP-008`, and `DRMCP-ADR-MCP-009`.
- PASS: `DRMCP-ADR-MCP-002` is referenced only as superseded historical authority.
- PASS: The six-component graph and inward dependency direction are consistent across views.
- PASS: Current Records and Legacy Archive remain separate source capabilities, separate ports, and separate request state.
- PASS: Guidance uses its own source capability and does not use a record snapshot.
- PASS: No exact module contract, package design, interface design, algorithm, adapter API, or implementation design was introduced.
- PASS: This Task did not modify Guidance operation Specifications.
- PASS: No Work Item graph, review verdict, finding repair, closure state, production source, test, or fixture changed.
- PASS: Scoped Git diff inspection was complete and untruncated for all six writable files.
- PASS: Scoped Git whitespace inspection reported no findings.
- PASS: No file is staged.

## Evidence

- Startup: Read repository-root `prompt_chappy.md` before every other repository file.
- Startup exclusion: Did not read `CLAUDE.md` or `AGENTS.md`.
- Access mode: DRMCP is non-operational, so all Design Record reads and writes used filesystem fallback.
- Created: `spec:drmcp.application_architecture` at `drmcp/records/spec/application-architecture/index.md`.
- Created: `spec:drmcp.application_architecture.application_boundary_and_components`.
- Created: `spec:drmcp.application_architecture.dependency_and_responsibility`.
- Created: `spec:drmcp.application_architecture.runtime_and_state`.
- Created: `spec:drmcp.application_architecture.failure_and_evolution`.
- Primary projection: D-016 -> Overview.
- Primary projection: D-001, D-002, D-003, D-005, and D-006 -> application boundary and components.
- Primary projection: D-007, D-010, and D-011 -> dependency and responsibility.
- Primary projection: D-004, D-008, D-009, D-012, D-013, and D-017 -> runtime and state.
- Primary projection: D-014 and D-015 -> failure and evolution.
- ADR projection: `DRMCP-ADR-MCP-007` supplies the whole-application component model.
- ADR projection: `DRMCP-ADR-MCP-008` supplies inward dependency and responsibility ownership.
- ADR projection: `DRMCP-ADR-MCP-009` supplies request-scoped record state and application lifecycle.
- Superseded authority: `DRMCP-ADR-MCP-002` was not used as current authority.
- Deferred treatment: D-012 proposal/body-cache design remains deferred. Its cross-layer example is labeled non-binding and outside the active architecture contract.
- Deferred treatment: D-013 write transaction and post-write validation design remains deferred.
- Concurrent boundary: Guidance source-path correction remains owned by T08. This Task did not write the three Guidance Specifications.
- Excluded design: No module contract, package layout, interface, type, function, method, constructor, algorithm, adapter API, authoring transaction, or implementation detail was introduced.
- Scoped diff: Inspected the five new Specifications and this Task without truncation. No writable-boundary violation was found.
- Whitespace: Scoped Git whitespace inspection returned `pass` with no findings. LF-to-CRLF notices were advisory only.
- Semantic validator: The standalone responsibility validator was unavailable because DRMCP is non-operational and no standalone invocation tool is exposed in this session.
- Git operations: No stage or commit was performed.
