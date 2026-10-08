# Build-time manifest binding model

## Status

Working contract confirmed on 2026-07-18.

This document defines the generic boundary that discovers application-owned YAML manifests during PuppyDSL generation and exposes each parsed manifest as a compiled Go value.
Manifest-specific schemas, validation, registration, and runtime behavior remain owned by their consuming subsystems.

## Objective

A PuppyDSL application may own YAML manifests used by MCP publication, artifact-module registration, runtime configuration, or another application concern.
Those manifests must be available to the compiled application without runtime filesystem discovery or runtime YAML parsing.

The build-time flow is:

```text
PuppyDSL application root
  -> discover bound manifest files
  -> validate the bind header
  -> parse each YAML document
  -> generate package puppygen/manifest
  -> compile generated yaml.Node values into the application binary
```

This mechanism binds YAML documents into generated Go data.
It does not interpret the document's application-specific schema.

## Discovery boundary

`puppygen` receives one PuppyDSL application root through its existing `-root` input.
Manifest discovery recursively scans beneath that application root.

A bound manifest file uses the exact lowercase compound extension:

```text
.manifest.yaml
```

The shorter `.yaml` extension, `.yml`, `.manifest.yml`, and other suffixes do not identify a bound manifest.
A bound manifest is not a `.puppy.yaml` declaration file and is not loaded into the PuppyDSL declaration registry.

Generated output trees and puppygen temporary or backup trees are excluded from discovery so generated files cannot become new build inputs.
Discovery order is the lexical order of application-root-relative paths and is used only to make generation and diagnostics deterministic.
Manifest ordering has no application semantic meaning unless a consuming manifest schema defines one inside the YAML document.

## Bind header

The first physical line of every bound manifest is exactly:

```text
#bind: <bind-name>
```

No blank line, comment, byte-order mark, or other content may precede the bind header.
`<bind-name>` uses the lower-snake identifier form:

```text
[a-z][a-z0-9]*(?:_[a-z0-9]+)*
```

The bind name is the manifest identity used by generated Go access.
The file name and source path are diagnostic and build-input evidence; they do not define manifest identity.

One PuppyDSL application must not contain two bound manifests with the same bind name.
The generator also rejects bind names whose generated exported Go symbols collide.

The bind header is source-envelope metadata and is not included in the parsed YAML mapping exposed to consumers.

## YAML document contract

After validating the bind header, puppygen parses the complete file as one YAML document.

The document must contain exactly one root mapping.
YAML multi-document streams, an empty document, or a scalar or sequence root are invalid.
The generic manifest binder validates only:

- the bind header and bind-name grammar;
- bind-name uniqueness and generated-symbol uniqueness;
- YAML syntax;
- one-document cardinality;
- mapping-root shape.

The binder does not validate manifest-specific keys, values, required fields, references, or semantic invariants.
For example, MCP tool names and FunctionIDs belong to the MCP manifest decoder, while artifact capability slots and module invariants belong to the artifact-module builder.

## Generated Go API

Bound manifests are generated into:

```text
<go-native-module>/puppygen/manifest
```

The generated package provides this stable conceptual API:

```go
package manifest

import "gopkg.in/yaml.v3"

type Document struct {
    BindName string
    Source   string
    Root     yaml.Node
}

var McpTools = Document{
    BindName: "mcp_tools",
    Source:   "manifests/mcp/tools.manifest.yaml",
    Root:     /* parsed mapping node */,
}

func ByBindName(name string) (Document, bool)

const Fingerprint = "..."
```

Each lower-snake bind name maps to one exported PascalCase package variable.
`Source` is the slash-normalized path relative to the selected PuppyDSL application root.
`Root` is the parsed YAML mapping node, not the source-envelope comment and not a runtime filesystem reference.
Generated `Document` values are immutable application build data by contract; consumers must not mutate the node tree.

`ByBindName` resolves the same generated documents by their exact bind names.
The generated variable and lookup API expose identical document values.

Generated file names, file counts, literal rendering strategy, and private lookup tables remain generator implementation details.

## Build-time parsing and binary inclusion

Puppygen parses every bound manifest before emitting Go source.
The generated Go source reconstructs the parsed `yaml.Node` tree directly.
The compiled application therefore requires neither the original manifest files nor a runtime YAML unmarshal step.

A malformed manifest fails puppygen generation before Go compilation.
A manifest-specific schema error may be detected by a later generated validator, build coordinator, or application bootstrap according to the consuming subsystem's contract.

## Fingerprint

The generated `manifest.Fingerprint` identifies the exact sorted bound-manifest source set.
It is computed from each application-root-relative source path and exact file content, including the bind header.

The manifest fingerprint remains separate from the existing PuppyDSL declaration source fingerprint because `.manifest.yaml` files and `.puppy.yaml` files have different grammars and consumers.
A build coordinator may derive a combined application build-input fingerprint from both values without changing either source-set identity.

## Consumer boundary

A consumer receives one generic `manifest.Document` and then owns all content-specific processing:

```text
manifest.Document.Root
  -> consumer-specific decode
  -> consumer-specific validation
  -> consumer-specific reference resolution
  -> immutable consumer registration or configuration
```

Examples:

```text
MCP tool manifest
  -> public tool declaration validation
  -> FunctionID resolution
  -> MCP registration

artifact module manifest
  -> module schema validation
  -> capability FunctionID resolution
  -> ArtifactModuleRegistration
```

The generic manifest binder must not depend on MCP, artifact modules, repository configuration, or any other manifest consumer.
A consuming subsystem must not rediscover the original manifest file at runtime when the generated document is available.

## Generation failures

Puppygen rejects generation when:

- a file ending in `.manifest.yaml` lacks the bind header on physical line one;
- the bind header is malformed or has an empty or invalid lower-snake bind name;
- two files declare the same bind name;
- two bind names produce colliding exported Go symbols;
- the YAML is malformed;
- the file contains no document or more than one document;
- the YAML document root is not a mapping;
- generated Go source cannot represent or compile the parsed node tree.

Manifest-specific schema failures are outside this generic failure set.

## Initial exclusions

The initial binding model does not provide:

- runtime manifest filesystem discovery;
- runtime YAML reparsing of generated manifest data;
- manifest hot reload;
- manifest mutation;
- schema inference from YAML content;
- one universal manifest schema;
- automatic FunctionID, capability-slot, or transport registration inside puppygen core;
- dynamic installation of new manifests into an already compiled process.
