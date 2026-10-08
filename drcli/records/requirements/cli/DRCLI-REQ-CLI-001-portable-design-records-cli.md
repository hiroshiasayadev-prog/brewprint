# DRCLI-REQ-CLI-001: Portable Design Records CLI

- **status**: accepted
- **date**: 2026-10-07
- **source_refs**:
  - DRMCP-REQ-MCP-001
  - DRMCP-REQ-MCP-002
  - DRMCP-REQ-MCP-003
  - spec:drmcp.design_records_mcp.operations
  - spec:product.design_records.authoring_standards
  - spec:product.design_records.repository_layout

## Requirement

DRCLI must provide Design Records discovery, read, validation, guidance, and authoring through a portable command-line interface.

DRCLI is an independent application namespace.
DRCLI does not replace, deprecate, or supersede DRMCP.
DRCLI may reuse DRMCP requirements and specifications as design evidence where their semantics remain applicable.

The CLI must work as a copied self-contained distribution on Windows and Linux.
The same distribution model must also allow one shared installation to serve multiple repositories.

## Evidence

- DRMCP already defines useful contracts for record discovery, retrieval, search, reference resolution, validation, authoring guidance, authoring transactions, and body-cache retry behavior.
- MCP transport is not required for the intended DRCLI usage.
- CLI callers include humans, automation, AI agents, CI, and remote command executors.
- Shell-embedded Markdown can contain quotes, backslashes, substitutions, newlines, and other characters that make argument-only transport unreliable for long bodies.
- Repository-wide traversal can become unnecessarily expensive when ignored dependency or generated directories are scanned.

## Required Outcome

### Portable deployment

- A DRCLI distribution can be copied as one directory and used without an installer.
- The distribution does not require MCP registration, a daemon, or a separately installed language runtime.
- The distribution includes the portable Design Records standards package used for authoring guidance and semantic validation.
- Bundled standards are resolved from the DRCLI distribution, not from the caller's current working directory.
- One DRCLI distribution can be placed on PATH and reused across multiple repositories.
- Windows and Linux expose equivalent command semantics, validation semantics, diagnostics, and exit outcomes.

### Repository selection

- The selected repository defaults to the caller's current working directory.
- A caller can explicitly select another repository root.
- Repository selection does not depend on where the DRCLI executable is installed.

### Repository discovery

- DRCLI discovers Design Records from the selected repository.
- Discovery honors Git-compatible ignore rules from the selected repository by default.
- Ignored directories are not traversed during normal discovery.
- A caller can explicitly disable ignore handling when full traversal is required.
- DRCLI does not require repository-specific installation to discover records.

### Read and navigation

DRCLI provides command-line access to the current Design Records capabilities needed for:

- repository and namespace discovery;
- record listing;
- logical tree navigation;
- exact record retrieval;
- exact section retrieval;
- scoped literal and pattern search;
- canonical reference resolution;
- broad-scope validation;
- exact-record validation.

DRCLI may reuse the corresponding DRMCP operation semantics.
Transport-specific MCP request and response shapes are not DRCLI requirements.

### Authoring guidance

DRCLI exposes the portable authoring standards as guidance operations.

The guidance surface supports:

- listing authoring guides with canonical ID, title, and abstract;
- reading one exact guide by canonical ref;
- returning complete guide Markdown without summarization or normalization.

Guidance uses the portable design_records standards package as its source.
DRCLI does not maintain a separate legacy guide source or guide-specific semantic model.

### Authoring

DRCLI supports record creation and partial record update according to PRODUCT-owned Design Records authoring standards.

Normal successful create and update operations must not require:

- a separate body-upload command;
- a mandatory propose-then-accept round trip;
- an interactive confirmation prompt.

Deferred proposal workflows may be provided for callers that explicitly request them.

### Body input and retry

Authoring operations support both:

- inline body input through command arguments; and
- body input through standard input.

Long or syntax-sensitive Markdown can use standard input without shell escaping the body as a command argument.

When DRCLI has received a body and proposal or write preparation fails in a retryable way:

- DRCLI preserves the submitted body;
- DRCLI returns an opaque body_cache_id;
- a later invocation can retry using the body_cache_id without resending the body.

Body-cache state survives CLI process termination and caller session termination.
Cache IDs are safe to use from independent sessions.
A shared DRCLI installation can retain operational state for more than one repository.

Any retained proposal state exposed by DRCLI follows the same process-independent persistence principle.

### CLI help and usage diagnostics

- DRCLI provides root-level help and command-specific help without requiring repository access.
- Help describes available commands, required and optional operands, body-input modes, repository selection, output modes, and representative valid invocations.
- Invalid command usage returns a concise diagnostic that identifies the invalid input or missing requirement.
- Usage errors also show the correct invocation shape or a directly relevant valid example.
- Usage errors identify the exact help command that provides the complete contract for the failed operation.
- Unknown commands and options must not fail with only a parser message or opaque error code.
- Mutually exclusive inputs, missing required inputs, invalid selector combinations, and unsupported operation modes must report the accepted alternatives.
- Machine-readable diagnostics remain available without removing actionable human-readable usage guidance.
- Help and usage diagnostics do not require callers to know an MCP-style request schema.

### Validation and diagnostics

- DRCLI validates records against PRODUCT-owned Design Records semantics.
- DRCLI does not redefine artifact identity, layout, lifecycle, section, or reference semantics.
- Validation distinguishes contract violations, advisories, unadmitted sources, identity conflicts, selector failures, and execution failures where applicable.
- Validation and authoring diagnostics are available in machine-readable form.
- Output limiting or truncation must not change the semantic validation outcome.

## Explicitly Excluded Scope

- Replacing, removing, renaming, or deprecating DRMCP.
- MCP transport or MCP server compatibility.
- Choosing a database engine or concrete operational-state file layout.
- Fixing concrete command names, flag names, JSON schemas, or numeric exit codes in this Requirement.
- Requiring callers to infer valid usage from implementation errors or external documentation.
- UI behavior.
- Defining Design Records artifact semantics inside DRCLI.
- Requiring one DRCLI installation per repository.
- Requiring body content to be written to a temporary file before authoring.

## Boundary

PRODUCT owns:

- Design Records artifact semantics;
- namespace and ID semantics;
- repository layout semantics;
- authoring standards;
- spec format;
- lifecycle rules;
- canonical references;
- portable standards-package content.

DRCLI owns:

- command-line interaction;
- help, usage discovery, and actionable invalid-usage diagnostics;
- portable packaging;
- repository selection;
- discovery orchestration;
- read and navigation execution;
- guidance projection;
- validation execution and diagnostics;
- authoring request handling;
- body input transport;
- retry cache behavior;
- process-independent operational-state persistence.

DRMCP remains an independent application and design reference.
DRCLI may adopt applicable DRMCP behavior without creating a replacement or compatibility obligation.
