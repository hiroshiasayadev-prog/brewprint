package resolve

import (
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

// taggedUnionModelFile builds a rawyaml.File with a single tagged union model
// in the model directory.
func taggedUnionModelFile(fileID string, model rawyaml.Model) rawyaml.File {
	return rawyaml.File{
		ID:   fileID,
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{model}},
	}
}

func minimalTaggedUnionModel(id string) rawyaml.Model {
	return rawyaml.Model{
		ID:            id,
		Kind:          "tagged_union",
		Discriminator: "kind",
		Variants: []rawyaml.ModelVariant{
			{Tag: "rename", Fields: []rawyaml.ModelField{{Name: "new_id", Type: "str"}}},
			{Tag: "remove", Fields: []rawyaml.ModelField{}},
		},
	}
}

// --- Valid cases ---

func TestBuildValidTaggedUnionModel(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Main:          true,
			Kind:          "tagged_union",
			Discriminator: "kind",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{{Name: "new_id", Type: "str"}}},
				{Tag: "remove", Fields: []rawyaml.ModelField{}},
				{Tag: "change_type", Fields: []rawyaml.ModelField{{Name: "new_type", Type: "str"}}},
			},
		}),
	}})
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	model := project.ModelsByQID["shop.model.change"]
	if model == nil || model.Kind != "tagged_union" {
		t.Fatalf("tagged union model not built: %#v", model)
	}
	if model.Discriminator != "kind" {
		t.Fatalf("discriminator = %q, want %q", model.Discriminator, "kind")
	}
	if len(model.Variants) != 3 {
		t.Fatalf("variants len = %d, want 3", len(model.Variants))
	}
	if model.Variants[0].Tag != "rename" || len(model.Variants[0].Fields) != 1 {
		t.Fatalf("variant[0] = %#v", model.Variants[0])
	}
	if model.Variants[1].Tag != "remove" || model.Variants[1].Fields == nil || len(model.Variants[1].Fields) != 0 {
		t.Fatalf("variant[1] (payload-less) = %#v", model.Variants[1])
	}
}

func TestTaggedUnionVariantFieldTypeRefResolvestoPrimitive(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", minimalTaggedUnionModel("change")),
	}})
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	model := project.ModelsByQID["shop.model.change"]
	if model == nil {
		t.Fatalf("model not found")
	}
	field := model.Variants[0].Fields[0]
	if field.TypeRef == nil || field.TypeRef.Kind != semantic.TypeRefPrimitive || field.TypeRef.Name != "str" {
		t.Fatalf("variant field TypeRef = %#v, want primitive str", field.TypeRef)
	}
}

func TestTaggedUnionVariantFieldTypeRefResolvesToNamedModel(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		typeRefStructModelFile("shop/model/user.yaml", "user"),
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Variants: []rawyaml.ModelVariant{
				{Tag: "update", Fields: []rawyaml.ModelField{{Name: "payload", Type: "user"}}},
			},
		}),
	}})
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	model := project.ModelsByQID["shop.model.change"]
	if model == nil {
		t.Fatalf("model not found")
	}
	field := model.Variants[0].Fields[0]
	if field.TypeRef == nil || field.TypeRef.Kind != semantic.TypeRefNamedModel || field.TypeRef.Model != "shop.model.user" {
		t.Fatalf("variant field TypeRef = %#v, want named model shop.model.user", field.TypeRef)
	}
}

func TestTaggedUnionUsedAsStructFieldTypeRef(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", minimalTaggedUnionModel("change")),
		{
			ID:   "shop/model/request.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{{
				ID:   "request",
				Kind: "struct",
				Fields: []rawyaml.ModelField{
					{Name: "selector", Type: "str"},
					{Name: "change", Type: "change"},
				},
			}}},
		},
	}})
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	req := project.ModelsByQID["shop.model.request"]
	if req == nil {
		t.Fatalf("request model not found")
	}
	changeField := req.Fields[1]
	if changeField.TypeRef == nil || changeField.TypeRef.Kind != semantic.TypeRefNamedModel || changeField.TypeRef.Model != "shop.model.change" {
		t.Fatalf("change field TypeRef = %#v, want named model shop.model.change", changeField.TypeRef)
	}
}

func TestTaggedUnionTypeRefCompatibilityNominal(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change_a.yaml", rawyaml.Model{
			ID: "change_a", Kind: "tagged_union", Discriminator: "kind",
			Variants: []rawyaml.ModelVariant{{Tag: "x", Fields: []rawyaml.ModelField{}}},
		}),
		taggedUnionModelFile("shop/model/change_b.yaml", rawyaml.Model{
			ID: "change_b", Kind: "tagged_union", Discriminator: "kind",
			Variants: []rawyaml.ModelVariant{{Tag: "x", Fields: []rawyaml.ModelField{}}},
		}),
	}})
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}

	changeA := mustParseTypeRefForTest(t, "change_a")
	changeB := mustParseTypeRefForTest(t, "change_b")
	str := mustParseTypeRefForTest(t, "str")
	any := mustParseTypeRefForTest(t, "any")

	if !typeRefsCompatible(project, changeA, mustParseTypeRefForTest(t, "change_a")) {
		t.Fatalf("same tagged union should be compatible")
	}
	if typeRefsCompatible(project, changeA, changeB) {
		t.Fatalf("different tagged union models should not be compatible")
	}
	if typeRefsCompatible(project, changeA, str) || typeRefsCompatible(project, str, changeA) {
		t.Fatalf("tagged union and str should not be compatible")
	}
	if !typeRefsCompatible(project, changeA, any) || !typeRefsCompatible(project, any, changeA) {
		t.Fatalf("any should remain compatible with tagged union")
	}
}

// --- Invalid cases ---

func TestTaggedUnionMissingDiscriminator(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:   "change",
			Kind: "tagged_union",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidTaggedUnionModel)
}

func TestTaggedUnionEmptyDiscriminator(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidTaggedUnionModel)
}

func TestTaggedUnionDotPathDiscriminator(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "object.kind",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidTaggedUnionModel)
}

func TestTaggedUnionMissingVariants(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidTaggedUnionModel)
}

func TestTaggedUnionEmptyVariants(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Variants:      []rawyaml.ModelVariant{},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidTaggedUnionModel)
}

func TestTaggedUnionMissingVariantTag(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Variants: []rawyaml.ModelVariant{
				{Tag: "", Fields: []rawyaml.ModelField{}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidTaggedUnionModel)
}

func TestTaggedUnionDuplicateVariantTag(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{}},
				{Tag: "rename", Fields: []rawyaml.ModelField{}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticDuplicateVariantTag)
	assertNoDiagnosticCode(t, diagnostics, diagnosticInvalidTaggedUnionModel)
}

func TestTaggedUnionMissingVariantFields(t *testing.T) {
	// Fields == nil means the "fields" key was absent from the YAML variant.
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename"}, // Fields is nil (key absent)
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidVariantField)
}

func TestTaggedUnionDiscriminatorFieldRepeatedInVariant(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{
					{Name: "kind", Type: "str"},
					{Name: "new_id", Type: "str"},
				}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidVariantField)
}

func TestTaggedUnionVariantFieldDisallowedPK(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{
					{Name: "new_id", Type: "str", PK: true},
				}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidVariantField)
}

func TestTaggedUnionVariantFieldDisallowedFK(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{
					{Name: "new_id", Type: "str", FK: "other.id"},
				}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidVariantField)
}

func TestTaggedUnionVariantFieldDisallowedUnique(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{
					{Name: "new_id", Type: "str", Unique: true},
				}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidVariantField)
}

func TestTaggedUnionDuplicateVariantPayloadFieldName(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{
					{Name: "new_id", Type: "str"},
					{Name: "new_id", Type: "str"},
				}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticDuplicateModelField)
}

func TestTaggedUnionVariantFieldInvalidTypeRef(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{
					{Name: "new_id", Type: "list<"},
				}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidTypeRef)
	assertNoDiagnosticCode(t, diagnostics, diagnosticUnresolvedFieldType)
}

func TestTaggedUnionVariantFieldUnresolvedTypeRef(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{
					{Name: "new_id", Type: "missing_model"},
				}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticUnresolvedFieldType)
	assertNoDiagnosticCode(t, diagnostics, diagnosticInvalidTypeRef)
}

func TestTaggedUnionDisallowedFieldsOnModel(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Fields:        []rawyaml.ModelField{{Name: "id", Type: "str"}},
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidTaggedUnionModel)
}

func TestTaggedUnionDisallowedElementOnModel(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Element:       "str",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidTaggedUnionModel)
}

func TestTaggedUnionDisallowedValueOnModel(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Value:         "str",
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidTaggedUnionModel)
}

func TestTaggedUnionDisallowedValuesOnModel(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		taggedUnionModelFile("shop/model/change.yaml", rawyaml.Model{
			ID:            "change",
			Kind:          "tagged_union",
			Discriminator: "kind",
			Values:        []string{"a", "b"},
			Variants: []rawyaml.ModelVariant{
				{Tag: "rename", Fields: []rawyaml.ModelField{}},
			},
		}),
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticInvalidTaggedUnionModel)
}
