package generator

import (
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type program struct {
	Root                string
	ModuleRoot          string
	ModulePath          string
	OutputImportPath    string
	NativeImportPath    string
	SourceFingerprint   string
	ManifestFingerprint string
	Types               []typeDeclaration
	Functions           []functionDeclaration
	Manifests           []manifestDeclaration
}

type manifestDeclaration struct {
	BindName string
	GoName   string
	Source   string
	Root     *yaml.Node
}

type typeKind string

const (
	typeKindScalar typeKind = "scalar"
	typeKindEnum   typeKind = "enum"
	typeKindRecord typeKind = "record"
	typeKindUnion  typeKind = "union"
	typeKindOpaque typeKind = "opaque"
)

type typeDeclaration struct {
	ID       string
	Detail   string
	Kind     typeKind
	Scalar   string
	Members  []string
	Fields   []recordField
	Variants []string
	Source   string
}

type recordField struct {
	Name string
	Type typeExpression
}

type functionDeclaration struct {
	ID              string
	Namespace       []string
	Terminal        string
	GoName          string
	Detail          string
	SignatureSource string
	Signature       functionSignature
	Implementation  implementationKind
	NativeBindingID string
	Source          string
}

type implementationKind string

const (
	implementationNative implementationKind = "native"
	implementationSteps  implementationKind = "steps"
	implementationStub   implementationKind = "stub"
)

type functionSignature struct {
	Inputs []functionInput
	Output typeExpression
}

type functionInput struct {
	Name string
	Type typeExpression
}

type typeExpression struct {
	Name      string
	Arguments []typeExpression
}

func (declaration functionDeclaration) packageRelativePath() string {
	if len(declaration.Namespace) == 0 {
		return ""
	}
	parts := append([]string(nil), declaration.Namespace...)
	return filepath.Join(parts...)
}
