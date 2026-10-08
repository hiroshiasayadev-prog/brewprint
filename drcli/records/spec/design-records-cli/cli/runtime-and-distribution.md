# Contract: DRCLI runtime and distribution

- **id**: `spec:drcli.design_records_cli.cli.runtime_and_distribution`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.cli`
- **contract_class**: `interface`

## What this is

Defines the portable DRCLI distribution boundary, bundled standards lookup, process-independent retained state, shared-installation behavior, and Windows/Linux observable parity.

## Non-goals

- Choosing a programming language, database engine, storage layout, package loader, or cleanup mechanism.
- Fixing a numeric retention duration for body caches or retained proposals.
- Requiring an installer, daemon, MCP registration, or per-repository installation.

## Request

DRCLI is distributed as one copyable directory for each supported operating system.
The logical command name is `drcli`.
The Windows executable filename may be `drcli.exe`.
The distribution uses no installer, daemon, or MCP registration and requires no separately installed language runtime.

The bundled Design Records standards tree is resolved relative to the DRCLI executable or distribution root.
Bundled standards lookup never resolves from the process current working directory or selected repository.

One shared DRCLI installation serves any number of repositories.
Repositories require no registration and no per-repository DRCLI installation.
Each repository-dependent invocation independently selects one repository under the repository-selection contract.
DRCLI creates no persistent current-repository setting.

Body caches and retained proposals are opaque operational state.
Operational state survives termination of the creating DRCLI process and caller session.
Opaque state lookup does not depend on the current working directory.

A body cache preserves the submitted body for later reuse.
A later body-capable invocation applies that body to the repository selected by that later invocation.
A body cache does not retain a hidden current repository.

A retained proposal remains bound to the repository selected when the proposal was created.
A later current working directory or `--repo` value must not silently retarget the proposal.
Opaque retained-state IDs originating from different repositories coexist without collision or ambiguous lookup.

Windows and Linux expose the same logical commands, option names, help meaning, diagnostic meaning, JSON contract, body-source rules, and numeric exit statuses.
All CLI text streams use UTF-8.
Only the platform-native executable filename and accepted filesystem path spelling may differ.
Filesystem path spelling never changes Design Records semantics or result meaning.

## Response

A package-dependent operation resolves bundled standards from the installed distribution and uses the same package regardless of selected repository.
A retained-state lookup resolves the opaque ID independently of process lifetime, caller session, and current working directory.
A shared installation keeps retained state unambiguous while repository selection remains per invocation.

No observable contract requires a daemon, background service, MCP session, database product, or language runtime.

## Errors

| condition | required behavior |
|---|---|
| Bundled standards required by an operation are missing or unusable | Fail that operation as `execution_failure` with status `3`. |
| Operational-state storage required by an invocation is unavailable or unreadable | Fail as `execution_failure` with status `3`. |
| A retained proposal is invoked from a different selected repository | Never retarget it; preserve the deferred-proposal contract's non-writing repository-binding outcome. |
| Help is requested while bundled standards, repository access, or retained state is unavailable | Render help successfully without accessing those resources. |

An unknown or expired opaque state ID remains subject to its authoring operation contract.
This runtime contract does not invent a retention period or reinterpret semantic state outcomes.

## Boundary

This contract owns deployment and observable runtime independence only.
Repository selection belongs to `spec:drcli.design_records_cli.cli.repository_selection_and_traversal`.
Body-cache semantics and proposal lifecycle remain owned by their authoring operation Specifications.

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.cli.repository_selection_and_traversal` | Owns per-invocation repository selection. |
| `spec:drcli.design_records_cli.cli.output_and_exit_status` | Owns execution-failure stream and status behavior. |
| `spec:drcli.design_records_cli.operations.authoring.retry_cache` | Owns body-cache preservation and reuse semantics. |
| `spec:drcli.design_records_cli.operations.authoring.deferred_proposal` | Owns proposal repository binding and lifecycle semantics. |
| `spec:drcli.design_records_cli.operations.guidance` | Consumes the portable bundled authoring-standards package. |