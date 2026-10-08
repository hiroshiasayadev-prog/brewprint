# DRCLI-TASK-CLI-002-02: Decide DRCLI external CLI contract

- **status**: done
- **date**: 2026-10-07
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: decision
- **estimate**: 1d
- **depends_on**:
  - DRCLI-TASK-CLI-002-01
- **outputs**:
  - DRCLI-TASK-CLI-002-02

## Goal

Fix the DRCLI-specific observable CLI contract that cannot be inherited directly from MCP transport.

## Work

Decide and record a terminal decision ledger for:

- logical executable and command-group structure;
- mapping from semantic operations to commands and subcommands;
- cwd-default repository selection and explicit repository override;
- Git-compatible ignore default and explicit full-traversal override;
- root help and command-specific help behavior;
- invalid-command, invalid-option, missing-input, mutually-exclusive-input, and invalid-selector usage guidance;
- inline body, stdin body, and body_cache_id input selection;
- human-readable and machine-readable output modes;
- stdout/stderr responsibility;
- exit-status categories and their meaning;
- portable distribution and bundled standards resolution;
- process-independent operational state behavior for body caches and optional proposals;
- multi-repository use from one shared installation;
- Windows/Linux observable parity.

Do not choose a database engine or implementation technology unless observable behavior cannot be specified without it.
Do not author Specification files.

For every decision, record ADR route as `not_required`, `covered`, or `required`.
If a new ADR is required, stop with the exact unresolved ADR boundary rather than authoring it here.

## Done condition

- Every listed external-contract decision is terminal.
- T08 can author the complete `cli/**` Specification without inventing behavior.
- Help/error behavior always makes valid usage discoverable without an MCP schema.
- No implementation technology is selected merely for convenience.

## Verification

- Trace every DRCLI-REQ-CLI-001 CLI-specific Required Outcome to one decision item.
- Confirm the decisions do not redefine PRODUCT-owned Design Records semantics.
- Confirm no Specification file was changed.

## Evidence

### Decision ledger

Workflow state: `decision_complete`. Current cursor: none. All owned decisions are terminal. Canonical projection target for every item is `DRCLI-TASK-CLI-002-08` under `drcli/records/spec/design-records-cli/cli/**`.

| ID | Status | Decision | Reason | ADR route |
|---|---|---|---|---|
| D-001 | decided | The logical executable name is `drcli`. The root command is partitioned into `scope`, `record`, `tree`, `reference`, `validate`, `guide`, and `proposal` command groups. | One stable executable and noun-oriented groups make the operation surface discoverable without MCP request-schema knowledge. | not_required |
| D-002 | decided | Semantic operations map as follows: `scope list` → record-scope discovery; `record list` → sequential-record listing; `tree children` → direct tree-child listing; `tree inspect` → bounded tree overview; `record get` → exact record retrieval; `record sections` → exact H2 section retrieval; `record search` → lexical record search; `reference resolve` → canonical reference resolution; `validate scope` → broad-scope validation; `validate records` → exact-record validation; `guide list` / `guide get` → authoring-guidance list/detail; `record create` / `record update` → normal one-invocation authoring; `proposal create` / `proposal update` / `proposal get` / `proposal accept` / `proposal discard` → explicit deferred-proposal workflow. | The mapping preserves the applicable DRMCP operation boundaries while giving direct authoring a CLI-native path that does not require propose-then-accept. | not_required |
| D-003 | decided | Every repository-dependent invocation selects exactly one repository. The default repository root is the process current working directory exactly; DRCLI does not search ancestors. Global `--repo <path>` explicitly replaces that default. | This is the direct CLI projection of the accepted repository-selection outcome and keeps installation location unrelated to target-repository selection. | not_required |
| D-004 | decided | Repository traversal honors Git-compatible ignore rules by default. `--no-ignore` is the explicit full-traversal override on commands whose operation builds a repository record view. The override disables ignore filtering only; it does not broaden PRODUCT-owned artifact or identity semantics. | This directly satisfies the discovery requirement without inventing another hidden-directory policy. | not_required |
| D-005 | decided | `drcli` with no command, `drcli --help`, and `drcli help` show root help and succeed. `drcli <group> --help`, `drcli <group> <command> --help`, and `drcli help <group> <command>` show scoped help and succeed. Help paths require no repository, bundled-standards, or operational-state access. | Help must remain usable before a repository, standards package, or retained state can be opened. | not_required |
| D-006 | decided | Unknown commands list valid siblings and the exact help command. Unknown options identify the option and valid option surface. Missing input identifies the missing operand/option and shows the accepted invocation shape. Mutually exclusive inputs name the conflicting inputs and accepted alternatives. Malformed or invalid selector combinations show accepted selector shapes plus the exact scoped help command. A well-formed but unresolved selector is an operation result, not a usage error. | This makes valid usage actionable instead of exposing only parser text or opaque codes. | not_required |
| D-007 | decided | Body-capable authoring commands use exactly three explicit body sources: `--body <markdown>`, `--body-stdin`, or `--body-cache-id <id>`. They are mutually exclusive. Stdin is read as body only when `--body-stdin` is present; DRCLI never guesses body mode from whether stdin is piped. Supplying no body source is valid only when the selected authoring operation does not require body content. | Explicit source selection avoids shell-escaping ambiguity and prevents accidental stdin consumption while retaining cache retry. | not_required |
| D-008 | decided | Operation output defaults to `--format human`; `--format json` selects machine-readable output. DRCLI does not switch output mode implicitly based on TTY detection. Explicit help remains human help text. Machine-readable diagnostics retain the same actionable usage information as human diagnostics. | Deterministic mode selection is required for humans, CI, agents, and remote executors on both supported OS families. | not_required |
| D-009 | decided | stdout carries the primary operation result whenever a trustworthy result exists, including partial or negative semantic results, and carries explicitly requested help. stderr carries usage failures and failures that prevent a trustworthy operation result. In JSON mode, stdout contains one machine-readable result document when operation execution produced a trustworthy result; usage/execution/internal failures emit one machine-readable diagnostic document on stderr. No extra prose is mixed into JSON result streams. | This gives pipelines one stable data stream while preserving diagnostics and nonzero status signaling. | not_required |
| D-010 | decided | Exit status categories are stable across platforms: `0 success`; `1 contract_result_failure` for a valid invocation that produced a trustworthy but unsatisfied semantic result, including validation error findings or partial selector failure; `2 usage_error`; `3 execution_failure` for unavailable repository/package/state, I/O, permission, or other inability to produce a trustworthy result; `4 internal_failure` for unexpected invariant/software failure. Warnings/advisories alone do not make status nonzero. | Separate usage, semantic, operational, and internal failures without conflating validation findings with runtime failure. | not_required |
| D-011 | decided | Distribution is one copyable directory per supported OS with logical command `drcli` (`drcli.exe` may be the Windows filename), no installer, daemon, MCP registration, or separately installed language runtime. The bundled Design Records standards tree is resolved relative to the DRCLI executable/distribution root, never from cwd or the selected repository. Missing or unusable bundled standards fail package-dependent operations as `execution_failure`; help remains available. | This is the accepted portable-package model projected to the CLI without selecting an implementation language or package loader. | not_required |
| D-012 | decided | Body caches and retained proposals are process- and caller-session-independent opaque state. No database engine, file layout, numeric retention duration, or cleanup mechanism is fixed here. State lookup by opaque ID must not depend on cwd. A body cache represents the preserved submitted body and can be reused by a later body-capable invocation against that invocation's selected repository. A retained proposal remains bound to the repository selected when it was created and must not be silently retargeted by a later cwd or `--repo` value. | The Requirement fixes persistence and independent-session reuse but intentionally leaves storage technology and timeout policy open. Repository binding is required for proposals because they retain repository-specific prepared state; body caches need no additional repository semantics. | not_required |
| D-013 | decided | One shared DRCLI installation serves any number of repositories without registration or per-repository installation. Each invocation independently selects one repository by D-003; no persistent “current repository” is created. Opaque retained-state IDs from different repositories must coexist without collision or ambiguous lookup, and proposal repository binding follows D-012. | This directly closes the shared-installation requirement while preserving per-invocation repository selection. | not_required |
| D-014 | decided | Windows and Linux use the same logical commands, option names, help/diagnostic meaning, JSON contract, body-source rules, and numeric exit statuses. Text streams are UTF-8. Only platform-native executable filename and accepted filesystem path spelling may differ; path spelling must not change Design Records semantics or result meaning. | This fixes observable parity while allowing unavoidable host filesystem syntax differences. | not_required |

### Required Outcome trace

| DRCLI-REQ-CLI-001 outcome | Decision coverage |
|---|---|
| Portable deployment | D-011, D-013, D-014 |
| Repository selection | D-003, D-013 |
| Repository discovery and ignore behavior | D-002, D-004 |
| Read and navigation | D-001, D-002, D-008, D-009, D-010 |
| Authoring guidance | D-002, D-005, D-011 |
| Normal authoring and optional deferred proposals | D-002, D-007, D-012 |
| Body input and retry | D-007, D-012, D-013 |
| CLI help and usage diagnostics | D-005, D-006, D-008, D-009, D-010 |
| Validation and diagnostics | D-002, D-006, D-008, D-009, D-010, D-014 |

### Verification result

- Every CLI-specific Required Outcome in `DRCLI-REQ-CLI-001` is traced to at least one terminal decision above.
- The command mapping reuses DRMCP semantic operation boundaries; it does not redefine PRODUCT-owned identity, layout, lifecycle, reference, authoring, or validation semantics.
- ADR routing is terminal for every decision. No item introduces an independent architecture, ownership, persistence-technology, or compatibility choice requiring a new ADR; all routes are `not_required`.
- The PRODUCT Task-responsibility validator has no operational TRV implementation available to invoke: current `TRV-WORK-SPEC-001` records validator delivery as suspended under `TRV-ADR-SPEC-006`. No semantic validator result is synthesized; the PRODUCT validator contract states that execution failure does not establish semantic non-compliance and does not itself enforce workflow continuation or release.
- No DRCLI Specification file was authored or modified by this Task.
- The parent Work Item lifecycle remains unchanged.

