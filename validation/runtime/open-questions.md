# Runtime design open questions

## Status

This file lists unresolved decisions that must be settled from concrete capability and artifact consumers.
It is not a backlog of optional language features.

## Source-file identification

Confirmed boundary:

- every design-facing PuppyDSL source file uses the exact compound extension `.puppy.yaml`;
- runtime bootstrap source discovery considers only files ending with `.puppy.yaml`;
- `.yaml`, `.yml`, and `.puppy.yml` do not identify PuppyDSL source files;
- ordinary YAML files may coexist beneath the selected PuppyDSL application root as application data and are not interpreted as PuppyDSL declarations by location or root key;
- bound application manifests use the separate exact suffix `.manifest.yaml` and are processed only by the build-time manifest binding contract;
- the designated entry declaration is `main.puppy.yaml`.

## Build-time manifest binding

Confirmed boundary:

- puppygen recursively discovers exact `.manifest.yaml` files beneath its selected PuppyDSL application root;
- generated output trees and puppygen temporary or backup trees are excluded from manifest discovery;
- physical line one of every bound manifest is exactly `#bind: <bind-name>` with no preceding content;
- bind names use lower-snake spelling, identify generated manifest documents independently of source paths, and are unique within one application;
- a bound manifest contains exactly one YAML document whose root is a mapping;
- the generic binder validates only the source envelope, bind identity, YAML syntax, document cardinality, and mapping-root shape;
- manifest-specific schemas and semantic validation remain owned by each consumer;
- puppygen emits parsed mapping-root `yaml.Node` values beneath `<go-native-module>/puppygen/manifest`;
- each bind name receives one PascalCase generated `manifest.Document` variable and is also available through exact `ByBindName` lookup;
- generated manifest documents carry the bind name and application-root-relative source path and are immutable build data by contract;
- generation embeds the parsed node tree directly so the compiled application performs no manifest filesystem discovery or YAML reparsing;
- the manifest source-set fingerprint is distinct from the `.puppy.yaml` declaration source fingerprint;
- MCP publication and artifact-module registration both consume this generic boundary and own their own decoding, validation, reference resolution, and registration.

## Identifier and generated Go API

Confirmed boundary:

- type IDs use `[A-Z][A-Za-z0-9]*`;
- function namespace segments, ordinary function terminal segments, context IDs, event IDs, module IDs, provided-value names, signature inputs, locals, transparent-record fields, and enum members use lower-snake spelling;
- an implementation-oriented function terminal segment may begin with one leading underscore;
- hyphens are prohibited in declaration identifiers but remain valid in ordinary prose, quoted string data, canonical Design Record IDs, and artifact references;
- application-owned PuppyDSL module directory segments and source-file stems use lower-snake spelling;
- generated global Go types are physically emitted beneath `<puppydsl-root>/go_native/puppygen/types` and imported from `<go-native-module>/puppygen/types`;
- a qualified function `<namespace>.<function_name>` is imported from `<go-native-module>/puppygen/<namespace as path segments>` and exported as the PascalCase form of `<function_name>`;
- the designated entry function is the reserved unqualified ID `main` and is exported as `Main` from the root `<go-native-module>/puppygen` package;
- every declared function receives a typed caller facade regardless of implementation form;
- an `implementation.native` function additionally receives `Bind<FunctionName>` and a generated concrete Go implementation-function type in the same package;
- generated file splitting and private helper layout are implementation details;
- app-owned runtime composition uses an app-specific namespace such as `drmcp_runtime`, while `runtime` remains reserved for generic interpreter facilities.

## Declaration file envelope

Confirmed boundary:

- every declaration file begins with exactly one `# responsibility:` source comment followed by exactly one `# excludes:` source comment before the declaration root;
- both required comments contain non-empty single-line statements;
- exactly one empty line separates the required header from the declaration root;
- the responsibility statement defines why the declarations belong in the same file;
- the exclusion statement identifies adjacent concerns that must remain outside the file;
- the required header is source-level design authority rather than decoded YAML data;
- raw-source envelope validation occurs before ordinary YAML declaration decoding;
- the file header owns the declaration-set boundary, while each declaration's required `detail` owns that declaration's individual semantic responsibility;
- missing, empty, reordered, malformed, or misplaced required header comments are declaration-file errors;
- additional source comments may appear only after the required header.

## Type declaration syntax

Pending:

- how primitive, named-scalar, enum, transparent-record, union, optional, list, and dictionary Go representations prove conformance to their declared type contracts;
- exact generated operation names and source format for opaque native representation binding.

Confirmed boundary:

- type declarations are grouped beneath a `types` root mapping, and each declaration identity is its mapping key;
- a named scalar is distinct from its primitive representation and from other named scalars, with no implicit conversion;
- named scalars are nongeneric, expose no fields, and do not define a closed value set;
- a named scalar declaration body contains exactly one `scalar: <primitive>` type body in addition to required declaration metadata;
- enum declaration bodies contain one `enum` member list;
- one enum declaration creates one type, and all of its members are values of that type;
- enum literals use the qualified form `EnumType.member`;
- enums are distinct from strings and other enum types, with no implicit conversion;
- enums are nongeneric, expose no fields, and require a nonempty duplicate-free member list;
- a declared enum may be consumed through exhaustive language-level `switch`, whose cases name every qualified member exactly once;
- an enum switch has no `default`, rejects members from another enum, and retains the target's enum type inside every case without creating member-specific types;
- opaque type declaration bodies contain exactly `opaque: true` in addition to required declaration metadata;
- opaque types are nongeneric, expose no fields, cannot be constructed by design-facing YAML, and admit no implicit conversion;
- every opaque TypeID requires exactly one application-owned compile-time Go representation binding before final generation and activation;
- missing, duplicate, unknown, non-opaque, `any`, or `interface{}` opaque representation bindings are generation failures;
- puppygen emits one distinct nominal Go wrapper per opaque TypeID rather than a raw representation alias or `any`-backed carrier;
- two opaque TypeIDs may use the same native representation while remaining non-interchangeable nominal wrapper types;
- generated callers and native binders use the TypeID-specific wrapper rather than the raw native representation;
- generated typed construction and representation access preserve compile-time checking without reflection or dynamic casts;
- opaque representation binding occurs before final wrapper and function generation through a staging API that does not depend on the final generated wrappers;
- the application-owned binding selects an opaque type's hidden representation while the YAML declaration owns its design-facing identity;
- union declaration bodies contain a nonempty list of distinct separately declared variant types;
- a union value has one interpreter-owned active variant type, but no separately declared tag or payload field;
- a value of one directly listed union variant is assignment-compatible where that union is the explicitly known target type, while union-to-variant downcast, unrelated-union conversion, transitive variant assignment, and union literal construction remain prohibited;
- the active variant type is not projectable data, and arbitrary pattern matching, `isinstance`, and runtime type assertions remain prohibited;
- a boolean predicate does not narrow a local's static type;
- a declared union may be eliminated through exhaustive language-level `switch`, whose cases name every declared variant type exactly once;
- inside each union case, the switch target local keeps its name and is statically narrowed to the named variant type;
- `optional<T>` is eliminated through the same exhaustive `switch` construct with exactly one `present` case and one `absent` case;
- inside optional `present`, the target keeps its name and narrows to `T`; inside `absent`, the target is unavailable in value positions because no `T` value exists;
- outside an optional switch, the target retains `optional<T>`;
- optional case words are control syntax rather than values or enum members, and the language defines no `none`, `some`, `null`, or other optional literal or constructor;
- `switch` has no `default` case and never exposes a union variant or optional presence state as string or enum data;
- declared runtime boundaries such as structure-specific module capability slots may also narrow a union after bootstrap proves the structure-to-variant and slot-signature mapping;
- transparent-record declaration bodies use optional `generic` and required `record` fields;
- the generic parameter order is declared in the type declaration key `Name<T, ...>`;
- declarations without type parameters omit both `<...>` and `generic`;
- unconstrained generic parameters are omitted from `generic`;
- constrained parameters map directly to the constraint name, for example `T: equatable`;
- the initial supported generic constraint vocabulary contains only `equatable`;
- `dict<K, V>` is a design-facing immutable generic composite;
- dictionary key type `K` is restricted to `string`, `integer`, `boolean`, a named scalar backed by one of those primitives, or a declared enum;
- union, list, optional, dictionary, transparent-record, opaque, `void`, and `never` key types are prohibited;
- dictionary key equality preserves declared type identity and value identity, with no implicit primitive-to-named-scalar conversion;
- dictionary entry order has no semantic meaning and duplicate keys never overwrite an earlier entry;
- `dict<K, V>` provides required-key bracket lookup through a statically typed binding-origin access chain, and a missing runtime key is an execution failure rather than an implicit `optional<V>` result;
- bounded dictionary iteration binds one `K` key and one `V` value and uses canonical key order: lexical order for strings, numeric order for integers, `false` then `true` for booleans, primitive-representation order for named scalars, and declaration order for enums;
- the initial profile provides no dictionary insertion, replacement, deletion, or other mutation syntax;
- the dictionary key restriction is built into `dict<K, V>` and does not add a user-authored generic constraint;
- a transparent-record declaration graph may be recursive only when every type-reference cycle crosses at least one `list<T>` or `dict<K, V>` constructor;
- direct self-reference, mutual recursion through direct record fields, and recursion whose only boundary is `optional<T>` are prohibited;
- PuppyDSL exposes no pointer, nullable-reference, or address-identity type; native pointer graphs and shared node identity require an opaque type and declared operations;
- `record` maps each exposed field name to one declared type expression;
- design-facing YAML may construct one concrete transparent-record value with `Type{field: value, ...}` authored as a folded scalar;
- a transparent-record constructor supplies every declared field exactly once, accepts only literals or visible bindings as field values, and rejects inline calls, projections, nested constructors, and arbitrary expressions;
- constructor field values require exact type identity except for direct variant-to-declared-union assignment compatibility;
- the YAML type declaration is authoritative and native Go structs do not expose undeclared fields;
- generic records such as `SetComparison<T>` are declared in YAML and native implementations must conform to that declaration;
- generic types and generic function signatures are required;
- `never` is a built-in control-result type for a function call with no normally continuing path;
- `never` is not constructible and is prohibited as an input, transparent-record field, union variant, generic type argument, context-provided value, or event payload;
- no `any`, dynamic object, class inheritance, YAML reflection, or implicit conversion other than direct variant-to-declared-union assignment compatibility.

## Function declaration syntax

Pending:

- exact runtime-consumer payload-binding syntax for consumers that inspect non-void event payloads;

Confirmed boundary:

- function declarations are grouped beneath a `functions` root mapping, and each stable function ID is its mapping key;
- generic functions declare their ordered type parameters in the declaration key `id<T, ...>`;
- optional `generic` maps only constrained parameters directly to supported constraint names; omitted parameters are unconstrained;
- nongeneric functions omit both `<...>` and `generic`;
- every function declares its named inputs and one output through the signature grammar `(<name>: <type>, ...) -> <type>`;
- every `signature` field uses a YAML folded block scalar with `>-`, including signatures whose content fits on one line;
- every function signature declares exactly one output type; ordinary functions return one value, `void` completes normally without an information-bearing value, and `never` has no normally continuing result path;
- multiple related values use one named transparent record output;
- native, steps-defined, and stub implementations expose the same declared signature and context contract;
- native functions use `implementation: { native: <binding-id> }`;
- a native binding ID names an activated registry entry rather than a Go package or source symbol;
- native bindings must resolve during bootstrap and exactly match the declared function signature;
- every declared function generates a typed Go caller facade from its FunctionID and signature;
- every native function additionally generates `Bind<FunctionName>` in the same function-namespace package;
- opaque parameters and results in those APIs use exact TypeID-specific nominal wrappers, including when several opaque TypeIDs share one native representation;
- final function API generation is prohibited until every referenced opaque TypeID has exactly one valid representation binding;
- passing an incompatible Go function or the wrong opaque wrapper type to the generated binder is a Go compile error;
- generated declaration fingerprints and activated binding descriptors are rechecked before activation so stale generated code or stale native binaries are rejected before execution;
- `native`, `steps`, and `stub` implementation forms are mutually exclusive;
- `implementation.stub` accepts exactly `native`, `steps`, or `undecided` and does not accept booleans or arbitrary strings;
- `stub: native` records that the final form is native while its binding and Go implementation remain unauthored;
- `stub: steps` records that the final form is steps-defined while its PuppyDSL body remains unauthored;
- `stub: undecided` records that native-versus-steps ownership remains an unresolved design decision;
- `stub: native` and `stub: steps` pass authoring validation without warning;
- `stub: undecided` passes authoring validation with one unresolved-implementation-direction warning;
- every stub is non-executable and prohibited by design-closure and activation validation regardless of its intent value;
- every called function and referenced type must be declared even during authoring; undefined declarations are validation errors rather than permitted placeholders;
- a steps-defined implementation places `steps` and any required `return` under `implementation`;
- an ordinary step may use either a plain YAML scalar or `>-` folded scalar for its expression text, and both forms create the same one-time immutable local binding;
- `>-` changes only YAML authoring and decoding, not assignment, mutability, initialization, or result typing semantics;
- `steps` is a nonempty source-ordered mapping from step keys to calls, binding-origin access chains, dictionary constructors, transparent-record constructors, `if`, exhaustive `switch` over declared unions, enums, or optional values, or bounded `for`;
- a step whose static result type is neither `void` nor `never` uses an ordinary key and creates one immutable local binding;
- a step whose static result type is `void` or `never` uses `_<completion-name>`, creates no local binding, and is not represented as a value asset in a derived dataflow graph;
- the exact key `_` is invalid, and underscore-prefixed step keys are invalid for information-bearing results;
- discard step keys cannot be referenced or returned;
- steps-defined functions whose output is neither `void` nor `never` require `return` naming one input, context value, or already bound local assignment-compatible with the declared output type; exact identity is required except for direct variant-to-declared-union assignment;
- `return` does not accept inline calls, projections, or other expressions;
- `void` functions omit `return` and complete successfully after their final step;
- native, steps-defined, and authoring-only stub functions may declare `-> never` under their ordinary implementation-profile rules;
- a steps-defined `never` function omits `return`, and every reachable path must terminate through another `never` call;
- `runtime.throw(message: string) -> never` is the initial native boundary that initiates ordinary unhandled execution-failure propagation, and PuppyDSL provides no catch or recovery construct;
- a `never` call uses a terminal discard step such as `_failed`, establishes no local, and prohibits any later reachable step in that lexical path;
- a `void` call uses a continuing discard step such as `_validated` or `_served` and establishes no local;
- ordinary lexical bindings are immutable; reassignment and shadowing are prohibited;
- native binding ID may differ from the design-facing function ID;
- context declarations are grouped beneath a `contexts` root mapping, and each context ID is its mapping key;
- each context declaration body contains a `provides` mapping when it owns provided values;
- each provided entry declares `provider: <function-call>` and may declare `destructor: <function-call>`;
- provided-value type is inferred from the provider signature and no separate `type` field is declared;
- provider output must be a cacheable value type; `void` and `never` are prohibited;
- provider arguments may reference only same-context or active-ancestor provided values and may not contain nested calls or expressions;
- destructor arguments may reference only `self`, and destructor output must be `void`;
- zero-input provider and destructor calls use explicit `()`;
- `opens` on a function is an ordered list of context IDs opened for the function's complete invocation;
- the first listed context is outermost, contexts close in reverse order, and omission of `opens` means the function opens no context;
- a context ID names a declaration, while context scope means the runtime interval during which one instance is active; there is no separate scope ID;
- context opening and closing inside steps, conditions, or iterations are prohibited;
- reopening an already active context ID is prohibited;
- provider references determine construction dependencies and must form an acyclic graph;
- an earlier `opens` entry may not depend on a context opened later in the same list;
- event declarations are grouped beneath an `events` root mapping, and each stable event ID is its mapping key rather than an arbitrary resource-specific consistency field;
- every event declares exactly one payload type; events without event data use `payload: void`, and `never` is prohibited as a payload;
- a consumer that reacts only to event identity omits payload binding; `invalidated_by` and `restart_on` contain event IDs only;
- long-lived contexts may declare `invalidated_by` event IDs; invalidation marks the current generation stale without mutating operations already using it;
- an operation pins its required long-lived context generations before body execution, and all later requirements in that operation resolve from those pinned generations;
- restartable entry functions may declare bounded `restart_on` event IDs with `max_restarts`; restart applies to the complete invocation only;
- runtime execution failures use the internal envelope `code`, `function`, `message`, and failure-specific `context`;
- restart exhaustion uses `code: operation_restart_exhausted`; `function` is the restartable entry function ID, and `context` contains exactly `max_restarts` and `restarts_performed`;
- mapping that runtime failure into DRMCP or another consumer's external diagnostic envelope is a separate operation and adapter concern;
- context-provided dependencies are declared through `requires`, are injectable, and are queryable;
- bootstrap validation must reject every entry call path with an unsatisfied direct or transitive context requirement;
- bootstrap validation must reject unknown or duplicate `opens` entries, invalid open order, unavailable provider ancestors, and direct or indirect provider dependency cycles;
- steps execute in source order;
- forward references are rejected as unbound-local references, so no separate local-step cycle rule is required;
- YAML-defined direct and mutual recursion are prohibited by bootstrap call-graph cycle validation;
- native Go implementations may use recursion internally;
- the prohibition is a runtime/bootstrap policy and could be relaxed later without changing ordinary YAML call syntax;
- generic per-node fuel is not used to bound ordinary calls or collection processing;
- no explicit visibility field is required initially; a leading `_` segment marks implementation-oriented functions by convention.

## Invocation grammar

Pending:

- whether future style enforcement should require named binding for particular multi-argument signatures;
- whether internal normalized dependency nodes remain visible for debugging.

Confirmed direction:

- calls use a parenthesized function-call form;
- dictionary construction uses `dict<K, V>{[key]: value, ...}` inside one folded YAML scalar;
- dictionary keys and values accept only literals or visible bindings, nested expressions are prohibited, explicit type arguments type empty dictionaries, and bracketed keys distinguish key expressions from record fields;
- duplicate dictionary keys are rejected rather than overwritten, including execution failure when equality becomes known only from runtime binding values;
- a value-producing access chain begins with one visible binding and may alternate declared transparent-record field projection and required-key dictionary lookup, for example `inventory.apps[app_namespace].artifact_roots[artifact_kind]`;
- dictionary lookup keys are primitive literals, enum literals, or visible bindings whose type exactly matches `K`; inline calls, projections, constructors, arithmetic, and other nested expressions are prohibited in the key position;
- a missing key in required bracket lookup is an execution failure and bracket access is unavailable for lists and other non-dictionary values;
- optional dictionary lookup may be exposed by an ordinary declared function returning `optional<V>` and does not change required bracket-lookup semantics;
- transparent-record construction uses `Type{field: value, ...}` inside one folded YAML scalar;
- constructor targets must resolve to concrete transparent-record types, and constructor fields use the same literal and visible-binding token rules as call arguments;
- the YAML parser reads each call as one scalar containing invocation text;
- a dedicated invocation lexer and parser interpret the function target, arguments, literals, and binding references inside that text;
- quoted invocation tokens are string literals;
- unquoted invocation identifiers and qualified identifiers are binding references;
- unquoted decimal integer tokens and the exact tokens `true` / `false` are primitive literals;
- type-qualified invocation members are enum literals when their qualifier resolves to a declared enum;
- the invocation grammar defines no null or optional literal and rejects YAML-like null spellings, `none`, `some(...)`, and equivalent forms inside call text;
- each call uses either positional or named binding, and the two modes cannot be mixed within one call;
- short calls may use YAML plain scalars and long calls use YAML folded scalars; scalar style does not change invocation-token meaning;

```yaml
result: path.to.function(args)
parse_outcome: module.reference.parse(base_ref)
```

is preferred over a verbose authoring surface based on `id`, `use`, and `with`.

## Artifact module capability slots

Pending:

- final required slot inventory;
- whether reference canonical formatting is artifact-bound or generic over artifact declarations;
- whether fact extraction is one slot or several narrower slots;
- whether validation execution is one module slot or discovered from artifact-owned contracts.

Confirmed boundary:

- artifact module manifests are bound `.manifest.yaml` documents supplied through the generated `puppygen/manifest` package rather than `.puppy.yaml` declaration roots;
- the artifact-module consumer decodes a `modules` root mapping, and each stable module ID is its mapping key;
- the manifest bind name identifies the application build document and does not replace the module IDs inside it;
- every module declaration includes required non-empty `detail` design authority;
- module manifests register ordinary functions into standardized slots;
- the module itself does not introduce a new language-level interface concept.

## Base reference parsing

Confirmed boundary:

- exact supplied reference text is never trimmed, repaired, completed, normalized, inferred, or fuzzy-matched;
- sequential `<APP_NAMESPACE>` and `<DOMAIN_NAMESPACE>` tokens use `[A-Z][A-Z0-9_]*` in the working contract;
- `-` is reserved as the sequential canonical-ID segment separator;
- shared sequential routing splits the exact input on `-`, reads token `0` as app namespace, token `1` as artifact-kind discriminator, token `2` as domain namespace, and preserves later tokens as `artifact_segments: list<string>`;
- `SequentialBaseReference` exposes exactly `raw`, `app_namespace`, `artifact_kind`, `domain_namespace`, and source-ordered `artifact_segments`;
- the selected artifact module, not shared routing, assigns artifact-specific names to `artifact_segments` and validates their count, width, role, and meaning;
- tree-form `spec:` refs retain the Product-owned lowercase segment grammar `[a-z0-9][a-z0-9_]*`;
- `TreeBaseReference` exposes exactly `raw`, `record_kind`, `app_namespace`, and source-ordered `path_segments` after the app namespace;
- tree `path_segments` are canonical identity segments and do not encode whether the last segment is represented physically by a directory `index.md` or a non-index Markdown file;
- physical directory, index-file, and leaf-file mapping belongs to source-derived identity construction or current-tree resolution rather than exact-reference parsing;
- the sequential namespace-token rule remains a working decision that must later be routed into Product authority;
- shared parsing owns routing evidence only and does not prove complete canonical grammar validity;
- `artifact.base.reference.parse(raw: string) -> SharedRoutingOutcome` is the shared parser contract;
- `SharedRoutingOutcome` is the union of `RoutedBaseReference` and `UnrouteableReference`;
- `RoutedBaseReference` exposes exactly `reference: BaseReference` so routing success/failure is separated from later tree/sequential handling;
- `UnrouteableReference` exposes exactly `raw: string`; its variant identity means the input matches neither the sequential nor tree routing envelope;
- shared routing declares no separate failure-reason type or field;
- `BaseReference` variants are `TreeBaseReference` and `SequentialBaseReference`, each preserving `raw` and structure-specific routing evidence;
- an `UnrouteableReference` does not proceed to module lookup;
- module-bound dispatch validates the selected module structure, narrows the union to the matching variant type, and invokes that module's `reference.parse` capability;
- the artifact-specific parser owns complete canonical grammar confirmation and returns `ArtifactReferenceParseOutcome`;
- `ArtifactReferenceParseOutcome` is the union of opaque `ArtifactReference` and transparent `InvalidArtifactReference`;
- native `ArtifactReference` is structure-specific while remaining one design-facing opaque type;
- native `SequentialArtifactReference` contains the original `SequentialBaseReference` plus artifact-specific `map<string, string>` fields whose keys name canonical tail segments;
- a successful sequential parser maps every `artifact_segments` entry to exactly one module-declared field in source order, with no missing or additional field;
- native `TreeArtifactReference` contains only the original `TreeBaseReference`, because that base already represents the complete canonical tree identity;
- the native sequential map is not a design-facing `dict<K, V>` value and YAML cannot inspect or construct it;
- no duplicate top-level artifact kind or structure field is required because the retained base discriminator supports module lookup;
- `InvalidArtifactReference` exposes exactly `raw: string` and declares no finer failure-reason field;
- only the `ArtifactReference` variant may proceed to canonical formatting or current-state resolution;
- both shared-routing failure and artifact-specific canonical-parse failure map to the consuming operation's malformed or invalid-selector classification;
- unresolved and conflicted outcomes require current repository state and occur only after successful canonical parsing.

## Domain outcome handling

Pending:

- whether a reusable generic result union is eventually required or concrete domain unions remain sufficient.

Confirmed boundary:

- expected outcomes are typed values;
- closed expected outcomes may be declared as unions and consumed through exhaustive variant `switch`;
- expected presence or absence uses `optional<T>` and exhaustive `present` / `absent` switch rather than an untyped null convention;
- shared routing uses the concrete domain union `SharedRoutingOutcome = RoutedBaseReference | UnrouteableReference`;
- artifact-specific parsing uses the concrete domain union `ArtifactReferenceParseOutcome = ArtifactReference | InvalidArtifactReference`;
- neither parse outcome carries a separate reason enum; each invalid variant preserves only the exact raw input;
- outcome-specific native predicate and extraction functions are not required merely to separate declared success and failure variants or optional presence states;
- steps-defined YAML describes expected domain branching but does not expose `try`, `catch`, `raise`, `except`, or execution-failure branching;
- detailed retry, fallback, error translation, and compensating cleanup belong inside semantic native Go functions when required;
- unhandled execution failure propagates automatically through callers to the selected entry boundary;
- execution failure is not an ordinary YAML-visible result branch;
- active union variant types remain interpreter-owned and are available only for exhaustive case selection, never as projectable data;
- optional presence state is likewise interpreter-owned control state; only `present` exposes the contained `T`, while `absent` exposes no value;
- arbitrary pattern matching, `isinstance`, reflection, switching over open-ended scalar values, and runtime type assertions remain prohibited; exhaustive switching over declared enums and optional values is permitted.

## Conditional execution

Pending:

- how unevaluable prerequisites suppress dependent validation findings.

Confirmed boundary:

- `if` is authored as `<result-step>: { if: { condition, then, optional else } }`, using an ordinary result local for an information-bearing result and a discard key for `void` or `never`;
- each branch contains ordered `steps` and an optional `return`;
- `if` is for boolean design-rule branching, while declared unions, enums, and optional values use exhaustive `switch`;
- an `if` condition is one previously produced boolean local;
- when no branch returns a value and at least one path continues normally, the `if` result is `void` and omitted `else` is a no-op;
- when every possible branch terminates through a `never` call, the complete `if` result is `never`;
- when a continuing branch returns `U`, every other explicitly continuing branch must return the same `U`;
- inside bounded `for`, a non-void `if` with omitted `else` treats the false path as implicit `continue` of the nearest enclosing iteration;
- outside bounded iteration, a non-void `if` with omitted `else` is invalid;
- a branch ending in explicit or implicit `continue` need not return;
- a branch ending in a `never` call has no continuing result and does not participate in branch result-type agreement;
- no missing branch result is implicitly converted to `optional<U>`;
- `then` and `else` use independent lexical child scopes;
- sibling branches may reuse a local name, but no branch local may shadow an active ancestor binding;
- inline comparison, arithmetic, arbitrary field traversal beyond statically typed binding-origin access chains, boolean composition, and type-test expressions are prohibited.

## Bounded iteration

Pending:

- finding-generation integration details.

Confirmed boundary:

- list iteration uses `<result-step>: { for: { item, in, steps, optional return } }`;
- dictionary iteration uses `<result-step>: { for: { key, value, in, steps, optional return } }`;
- exactly one binding form is permitted: `item` for `list<T>`, or both `key` and `value` for `dict<K, V>`;
- the enclosing result step is always required; it is an ordinary local when `return` produces `list<U>` and a discard key when omitted `return` produces `void`;
- `in` references one previously bound immutable `list<T>` or `dict<K, V>` local and is resolved once when iteration begins;
- inline calls, access chains, constructors, or other expressions are not permitted in `in`;
- iteration operates over finite immutable list or dictionary values;
- list iteration preserves list order, while dictionary iteration uses the canonical key order defined by the value model;
- the key binding has type `K` and the value binding has type `V` during dictionary iteration;
- absence of `return` yields `void`;
- `return` of `U` mechanically yields `list<U>`; direct `void` steps cannot supply `U` because discard steps create no returnable local, though another declared function may return a complete composite value such as `list<void>`;
- `U` is determined statically from the returned local, so empty or fully skipped list or dictionary iteration still yields a typed empty `list<U>`;
- returned lists are not flattened and optional values are not filtered implicitly;
- explicit `continue` is allowed only as the terminal action of an `if` branch or switch case inside the bounded iteration and contributes no output element;
- omitted `else` on a non-void `if` inside bounded iteration is an implicit `continue` and likewise contributes no output element;
- a `never` call inside an iteration terminates the complete runtime invocation rather than contributing or continuing an iteration;
- nested `continue` targets only the nearest enclosing `for`; labelled or outer-targeted continue is not included;
- `break` is not included initially;
- no general unbounded loop is introduced;
- waiting or long-lived loops such as event loops and filesystem watchers belong to native Go or the runtime host;
- `if`, exhaustive `switch` over declared unions, enums, or optional values, and bounded `for` may nest in any combination;
- every `if` branch, `for` iteration body, and switch case may read enclosing bindings but does not export its own locals; list item bindings and dictionary key/value bindings are likewise iteration-local;
- only an ordinary local bound to a non-`void`, non-`never` control-flow result becomes visible in the enclosing scope; `void` and `never` control-flow results use discard keys and establish no local;
- a union switch narrows its target local, under the same name, to the selected variant type inside the case; an enum switch retains the declared enum type while fixing the selected member only within that case; an optional switch narrows the target to `T` in `present` and makes it unavailable as a value in `absent`;
- nested control flow applies the same lexical-scope rules recursively; no generic fuel counter is attached to control-flow nodes.

## Runtime scope

Pending:

- whether this runtime remains validation-specific or becomes a general Brewprint design-execution runtime;
- canonical repository placement and namespace if generalized.

Current rule:

- this working set remains under `validation/`;
- no general runtime claim is made until another concrete use case requires the same contracts.
