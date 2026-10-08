# DRCLI-TASK-CLI-003-16: Package portable Windows and Linux distributions

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: implementation
- **estimate**: 2d
- **depends_on**:
  - DRCLI-TASK-CLI-003-12
  - DRCLI-TASK-CLI-003-13
  - DRCLI-TASK-CLI-003-15
- **outputs**:
  - drcli/scripts/package.ps1
  - drcli/scripts/package.sh
  - drcli/tests/package/package_test.go
  - drcli/README.md
  - drcli/dist/windows/drcli.exe
  - drcli/dist/linux/drcli
  - drcli/dist/windows/design-records/
  - drcli/dist/linux/design-records/

## Goal

Implement package portable windows and linux distributions in its sole-owned test/package target.

## Work

### Read

- `DRCLI-TASK-CLI-003-16` (this Task).
- `drcli/records/spec/design-records-cli/cli/runtime-and-distribution.md`
- `bin/design-records/authoring-standards/index.md`
- `scripts/verify.bat`
- `go.mod`

### Change

- `drcli/scripts/package.ps1`
- `drcli/scripts/package.sh`
- `drcli/tests/package/package_test.go`
- `drcli/README.md`
- `drcli/dist/windows/drcli.exe`
- `drcli/dist/linux/drcli`
- `drcli/dist/windows/design-records/`
- `drcli/dist/linux/design-records/`

### Implement

- Fixed interface: package.ps1/package.sh run go -C drcli build with GOOS windows/linux and GOARCH amd64, copy accepted bin/design-records tree to executable-relative design-records/, verify SHA256 and manifest; never mutate source generator or global bin.
- Behavior: Deliver two standalone copyable directories with no installer, daemon, required MCP transport or installed language runtime; both contain executable and portable standards with rewritten design_records refs. Runtime uses executable-relative package path rather than cwd or --repo; allow a shared installation for multiple selected repos and separate retained state; do not embed or overwrite user repositories. Packaging must not depend on source-path execution after copying.
- Model routing: Codex/Sonnet for platform integration and cross-process isolation.

### Test

- On each supported OS, from a different cwd, root help works with missing package/repo/state; packaged guide list/get read actual bundled specs; cross-compiled artifacts exist; no accidental paths outside drcli/dist; runtime version and command/help exit codes agree.
- Focused command: `go -C drcli test ./tests/package -count=1`, expected exit 0.

### Stop

- BLOCKED if build conventions, package tree or required predecessor APIs are unavailable, a source outside Change must change, or tests fail. Name exact dependency.

### Prohibited operations

- No change to accepted Specification, PRODUCT, DRMCP, historical review, another Task's code; no staging or commits. Review and verification never repair code.

### Output

- Declared changed paths, focused test evidence, package checksums, downstream verification handoff.

## Implementation contract

| target | required change | acceptance criterion | verification |
|---|---|---|---|
| `drcli/scripts/package.ps1` | Produce the exact build or test target for this Task. | Deliver two standalone copyable directories with no installer, daemon, required MCP transport or installed language runtime; both contain executable and portable standards with rewritten design_records refs. | `go -C drcli test ./tests/package -count=1` |
| `drcli/scripts/package.sh` | Produce the exact build or test target for this Task. | Deliver two standalone copyable directories with no installer, daemon, required MCP transport or installed language runtime; both contain executable and portable standards with rewritten design_records refs. | `go -C drcli test ./tests/package -count=1` |
| `drcli/tests/package/package_test.go` | Produce the exact build or test target for this Task. | Deliver two standalone copyable directories with no installer, daemon, required MCP transport or installed language runtime; both contain executable and portable standards with rewritten design_records refs. | `go -C drcli test ./tests/package -count=1` |
| `drcli/README.md` | Produce the exact build or test target for this Task. | Deliver two standalone copyable directories with no installer, daemon, required MCP transport or installed language runtime; both contain executable and portable standards with rewritten design_records refs. | `go -C drcli test ./tests/package -count=1` |
| `drcli/dist/windows/drcli.exe` | Produce an executable-relative packaged output. | On each supported OS, from a different cwd, root help works with missing package/repo/state. | `go -C drcli test ./tests/package -count=1` |
| `drcli/dist/linux/drcli` | Produce an executable-relative packaged output. | On each supported OS, from a different cwd, root help works with missing package/repo/state. | `go -C drcli test ./tests/package -count=1` |
| `drcli/dist/windows/design-records/` | Produce an executable-relative packaged output. | On each supported OS, from a different cwd, root help works with missing package/repo/state. | `go -C drcli test ./tests/package -count=1` |
| `drcli/dist/linux/design-records/` | Produce an executable-relative packaged output. | On each supported OS, from a different cwd, root help works with missing package/repo/state. | `go -C drcli test ./tests/package -count=1` |

## Done condition

- All contract targets satisfy the stated behaviors and `go -C drcli test ./tests/package -count=1` exits 0.

## Verification

- Inspect exact package outputs and run `go -C drcli test ./tests/package -count=1`; T18 owns cross-platform integration.

## Evidence

TBD
