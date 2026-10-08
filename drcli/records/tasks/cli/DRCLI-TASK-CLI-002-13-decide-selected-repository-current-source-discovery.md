# DRCLI-TASK-CLI-002-13: Decide selected-repository current-source discovery

- **status**: done
- **date**: 2026-10-07
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - DRCLI-TASK-CLI-002-12
- **outputs**:
  - DRCLI-TASK-CLI-002-13

## Goal

Fix how one selected repository yields the app-namespace and records-root associations consumed by the DRCLI current-record model.

## Work

Decide a terminal contract consistent with all accepted constraints:

- repository context is cwd by default or explicit --repo;
- normal operation traverses beneath that selected repository;
- Git-compatible ignore rules prune traversal by default;
- --no-ignore disables only ignore pruning;
- no mandatory per-repository installation, registration, or manifest may be required;
- one shared DRCLI installation must work across multiple repositories;
- app namespace must not become canonical merely from a physical directory name;
- malformed sources must remain reportable after a source root is resolved;
- ambiguous or unresolvable source-root namespace must fail explicitly rather than choose silently.

The preferred resolution, unless contradicted by canonical authority, is:

1. discover candidate Design Records roots beneath the selected repository from repository-layout structure rather than app-directory naming;
2. determine one root's app namespace from syntactically extractable canonical record identities within that root;
3. require exactly one distinct app namespace for that root;
4. once resolved, associate every in-corpus source under that root with the resolved app namespace, including malformed/unadmitted sources;
5. report zero-namespace and multi-namespace roots as explicit current-source discovery diagnostics;
6. do not require a repository manifest or persistent registration.

Do not author Specification files.

## Done condition

- The complete source-root discovery and app-namespace resolution behavior is terminal.
- Empty, malformed-only, mixed-namespace, ignored, and valid roots have explicit disposition.
- T14 can author exact contracts without inventing policy.

## Verification

- Confirm no path-name-derived canonical app identity is introduced.
- Confirm no mandatory repository config/registration is introduced.
- Confirm broad validation can still report malformed/unadmitted sources for a resolved root.
- Confirm behavior remains compatible with cwd/--repo and ignore semantics.

## Evidence

### Decision ledger

Workflow state: `decision_complete`. Current cursor: none. Canonical projection target is DRCLI-TASK-CLI-002-14.

| ID | Status | Decision | Reason | ADR route |
|---|---|---|---|---|
| D-001 | decided | Discover a candidate records root from traversal-eligible directories named exactly `records` at or beneath the selected repository. Do not derive app identity from the parent directory name. After a candidate root is found, do not treat nested `records` directories inside that root as additional app roots. | PRODUCT defines `records/` as the app-independent placement root while prohibiting path-derived app identity. | not_required |
| D-002 | decided | Resolve namespace evidence only from syntactically extractable canonical record identities in traversal-eligible, in-corpus sources. Sequential sources contribute the app segment from a valid H1 public-ID grammar. Spec sources contribute the app segment from a syntactically valid H1-adjacent `spec:<app>...` ID, mapped to the corresponding app-namespace token. File names and directory names never contribute namespace evidence. | PRODUCT makes sequential H1 identity authoritative and exposes the app segment in both sequential IDs and visible spec refs. | not_required |
| D-003 | decided | One candidate root resolves only when namespace evidence yields exactly one distinct app namespace. Zero namespaces is unresolvable. More than one is ambiguous. DRCLI never chooses a namespace winner. | This preserves non-path canonical identity and makes ambiguity explicit. | not_required |
| D-004 | decided | After a root resolves, associate every traversal-eligible source in that root's PRODUCT-defined current artifact corpora with the resolved app namespace, including malformed and unadmitted sources that contributed no namespace evidence. | Namespace resolution must not erase broad-validation visibility of malformed sources. | not_required |
| D-005 | decided | A malformed-only root still resolves when at least one source contains syntactically valid namespace evidence and every such identity names the same app namespace. If no source yields namespace evidence, the root is unresolvable. If evidence names multiple app namespaces, the root is ambiguous. | Admission and nonidentity conformance remain later phases; namespace extraction is only the bootstrap needed to establish expected app context. | not_required |
| D-006 | decided | Git-compatible ignore pruning applies before root discovery, namespace-evidence collection, and source-corpus construction. A fully ignored root is not discovered and produces no discovery diagnostic. `--no-ignore` only removes that pruning and then applies the same discovery rules. | This preserves T02 cwd/--repo and ignore semantics without broadening artifact semantics. | not_required |
| D-007 | decided | A selected repository with no discovered candidate root produces a valid empty current state. A discovered empty `records/` root yields zero namespace evidence and is unresolvable. Distinct resolved roots may coexist only when they resolve to distinct app namespaces. Two resolved roots for one app namespace are a configuration failure; DRCLI does not merge or choose between them. | This preserves one-root-per-app and one-app-per-root uniqueness without manifests or registration. | not_required |
| D-008 | decided | Unresolvable, ambiguous, or duplicate-app-root conditions make selected-repository current-state construction fail request-wide using the existing `configuration_failure` semantic operation classification. The diagnostic message must identify which cause occurred and the affected repository-relative records root; the multi-namespace cause must also identify the conflicting namespaces. No new diagnostic code, context schema, or owner is introduced. | Existing operations already define context-free `configuration_failure` and no partial normal result when current state cannot be constructed reliably. Cause-specific message content keeps the failure explicit without expanding every operation contract. | not_required |
| D-009 | decided | DRCLI requires no repository manifest, persistent registration, per-repository installation, or app-directory naming convention beyond locating the standard `records/` placement root. One shared installation discovers each invocation's selected repository independently. | This preserves DRCLI-REQ-CLI-001 and T02 D-003/D-013. | not_required |

### Edge disposition

| root state | disposition |
|---|---|
| Valid | Exactly one namespace resolves; create one app-namespace/records-root association and retain the full in-corpus source set. |
| Empty discovered `records/` root | Zero namespace evidence; `configuration_failure` as unresolvable. |
| Malformed-only | Resolve only when syntactically valid identity evidence still yields exactly one namespace; otherwise apply zero- or multi-namespace failure. |
| Mixed namespace | `configuration_failure` as ambiguous; no association and no winner. |
| Ignored | Omitted from discovery under default ignore handling; reconsidered normally under `--no-ignore`. |
| Unresolvable | `configuration_failure`; no association. |
| Duplicate app across roots | `configuration_failure`; roots are not merged. |

### Verification result

- T12 is `done`; T13 owns only the new current-source discovery judgment and T14 remains the canonical Specification projection owner.
- PRODUCT repository-layout authority permits automatic `records/` discovery and forbids deriving app namespace from the physical records-root path.
- PRODUCT identity authority provides syntactically extractable app segments from sequential H1 IDs and visible spec IDs; no mandatory repository manifest or registration authority was found.
- Broad validation remains able to report malformed and unadmitted in-corpus sources whenever their root resolves.
- cwd/`--repo`, Git-compatible ignore, `--no-ignore`, shared-installation, and no-registration decisions from T02 remain unchanged.
- The PRODUCT Task-responsibility validator has no operational implementation to invoke: current TRV design work remains blocked before production implementation. No semantic validator result was synthesized.
- No Specification file, DRMCP file, parent Work Item lifecycle, or Git state was modified by this Task.
