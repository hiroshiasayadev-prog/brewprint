package resolve

import (
	"strings"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

func TestParseResolvePrimitiveTypeRef(t *testing.T) {
	ref, err := parseTypeRef("str", "shop")
	if err != nil {
		t.Fatalf("parseTypeRef returned error: %v", err)
	}
	if ref.Kind != semantic.TypeRefPrimitive || ref.Name != "str" || ref.Raw != "str" {
		t.Fatalf("ref = %#v, want primitive str", ref)
	}
}

func TestParseResolveNamedModelTypeRef(t *testing.T) {
	ref, err := parseTypeRef("user", "shop")
	if err != nil {
		t.Fatalf("parseTypeRef returned error: %v", err)
	}
	if ref.Kind != semantic.TypeRefNamedModel || ref.Model != "shop.model.user" {
		t.Fatalf("ref = %#v, want named model shop.model.user", ref)
	}

	qualified, err := parseTypeRef("catalog.product", "shop")
	if err != nil {
		t.Fatalf("parseTypeRef returned error: %v", err)
	}
	if qualified.Kind != semantic.TypeRefNamedModel || qualified.Model != "catalog.model.product" {
		t.Fatalf("qualified = %#v, want named model catalog.model.product", qualified)
	}
}

func TestParseResolveListTypeRef(t *testing.T) {
	ref, err := parseTypeRef("list<user>", "shop")
	if err != nil {
		t.Fatalf("parseTypeRef returned error: %v", err)
	}
	if ref.Kind != semantic.TypeRefList || ref.Elem == nil || ref.Elem.Model != "shop.model.user" {
		t.Fatalf("ref = %#v, want list<shop.model.user>", ref)
	}
}

func TestParseResolveDictTypeRef(t *testing.T) {
	ref, err := parseTypeRef("dict<config>", "shop")
	if err != nil {
		t.Fatalf("parseTypeRef returned error: %v", err)
	}
	if ref.Kind != semantic.TypeRefDict || ref.Value == nil || ref.Value.Model != "shop.model.config" {
		t.Fatalf("ref = %#v, want dict<shop.model.config>", ref)
	}
}

func TestParseResolveNestedTypeRef(t *testing.T) {
	ref, err := parseTypeRef("list<dict<user>>", "shop")
	if err != nil {
		t.Fatalf("parseTypeRef returned error: %v", err)
	}
	if ref.Kind != semantic.TypeRefList || ref.Elem == nil || ref.Elem.Kind != semantic.TypeRefDict || ref.Elem.Value == nil || ref.Elem.Value.Model != "shop.model.user" {
		t.Fatalf("ref = %#v, want list<dict<shop.model.user>>", ref)
	}
}

func TestBuildSetsSemanticTypeRefs(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		typeRefStructModelFile("shop/model/user.yaml", "user"),
		typeRefStructModelFile("shop/model/config.yaml", "config"),
		{
			ID:   "shop/model/user_list.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{{
				ID:      "user_list",
				Kind:    "list",
				Element: "list<user>",
			}}},
		},
		{
			ID:   "shop/model/config_map.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{{
				ID:    "config_map",
				Kind:  "dict",
				Value: "dict<config>",
			}}},
		},
		{
			ID:   "shop/task/run.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Tasks: []rawyaml.Task{{
				ID:   "run",
				Main: true,
				Params: []rawyaml.Param{
					{Name: "users", Model: "list<user>"},
					{Name: "configs", Model: "dict<config>"},
				},
				Returns: &rawyaml.Return{Name: "nested", Model: "list<dict<user>>"},
			}}},
		},
	}})
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	task := project.TasksByQID["shop.task.run"]
	if task == nil || len(task.Params) != 2 || task.Params[0].TypeRef == nil || task.Params[0].TypeRef.Kind != semantic.TypeRefList {
		t.Fatalf("task params TypeRef not set: %#v", task)
	}
	if task.Returns == nil || task.Returns.TypeRef == nil || task.Returns.TypeRef.Kind != semantic.TypeRefList || task.Returns.TypeRef.Elem == nil || task.Returns.TypeRef.Elem.Kind != semantic.TypeRefDict {
		t.Fatalf("task return TypeRef not set: %#v", task.Returns)
	}
	userList := project.ModelsByQID["shop.model.user_list"]
	if userList == nil || userList.ElementRef == nil || userList.ElementRef.Kind != semantic.TypeRefList {
		t.Fatalf("model ElementRef not set: %#v", userList)
	}
	configMap := project.ModelsByQID["shop.model.config_map"]
	if configMap == nil || configMap.ValueRef == nil || configMap.ValueRef.Kind != semantic.TypeRefDict {
		t.Fatalf("model ValueRef not set: %#v", configMap)
	}
}

func TestInvalidTypeRefDiagnostics(t *testing.T) {
	tooDeep := strings.Repeat("list<", maxTypeRefContainerNestingDepth+1) + "user" + strings.Repeat(">", maxTypeRefContainerNestingDepth+1)
	cases := []string{"list<", "dict<>", "map<user>", tooDeep}
	for _, raw := range cases {
		t.Run(raw, func(t *testing.T) {
			_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{typeRefModelFile("shop/model/user.yaml", raw)}})
			assertDiagnosticCode(t, diagnostics, diagnosticInvalidTypeRef)
			if countDiagnosticCode(diagnostics, diagnosticUnresolvedFieldType) != 0 {
				t.Fatalf("got unresolved_field_type with invalid_type_ref: %#v", diagnostics)
			}
		})
	}
}

func TestOpaqueContainerTypeRefWarnsButBuilds(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		typeRefStructModelFile("shop/model/user.yaml", "user"),
		{
			ID:   "shop/task/run.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Tasks: []rawyaml.Task{{
				ID:      "run",
				Returns: &rawyaml.Return{Name: "payload", Model: "list<dict<any>>"},
			}}},
		},
	}})
	if project == nil {
		t.Fatalf("project is nil")
	}
	assertDiagnostic(t, diagnostics, diagnosticOpaqueTypeRef, semantic.SeverityWarning)
	if countDiagnosticCode(diagnostics, diagnosticInvalidTypeRef) != 0 {
		t.Fatalf("got invalid_type_ref for opaque but syntactically valid TypeRef: %#v", diagnostics)
	}
}

func TestBareAnyDoesNotWarnAsOpaqueContainer(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "shop/task/run.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Tasks: []rawyaml.Task{{
				ID:      "run",
				Returns: &rawyaml.Return{Name: "payload", Model: "any"},
			}}},
		},
	}})
	if countDiagnosticCode(diagnostics, diagnosticOpaqueTypeRef) != 0 {
		t.Fatalf("bare any should not emit opaque_type_ref: %#v", diagnostics)
	}
}

func TestFieldListMissingModelUsesUnresolvedFieldType(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{typeRefModelFile("shop/model/user.yaml", "list<missing_model>")}})
	assertDiagnosticCode(t, diagnostics, diagnosticUnresolvedFieldType)
	if countDiagnosticCode(diagnostics, diagnosticInvalidTypeRef) != 0 {
		t.Fatalf("got invalid_type_ref for syntactically valid missing model: %#v", diagnostics)
	}
}

func TestParamListMissingModelUsesUnresolvedModel(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "shop/task/run.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Tasks: []rawyaml.Task{{
				ID:     "run",
				Params: []rawyaml.Param{{Name: "users", Model: "list<missing_model>"}},
			}}},
		},
	}})
	assertDiagnosticCode(t, diagnostics, diagnosticUnresolvedModel)
	if countDiagnosticCode(diagnostics, diagnosticInvalidTypeRef) != 0 {
		t.Fatalf("got invalid_type_ref for syntactically valid missing model: %#v", diagnostics)
	}
}

func TestReturnPrimitiveTypeRefIsValid(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "shop/task/run.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Tasks: []rawyaml.Task{{
				ID:      "run",
				Main:    true,
				Returns: &rawyaml.Return{Name: "message", Model: "str"},
			}}},
		},
	}})
	if countDiagnosticCode(diagnostics, diagnosticUnresolvedModel) != 0 || countDiagnosticCode(diagnostics, diagnosticInvalidTypeRef) != 0 {
		t.Fatalf("primitive return should be valid, diagnostics: %#v", diagnostics)
	}
	task := project.TasksByQID["shop.task.run"]
	if task == nil || task.Returns == nil || task.Returns.TypeRef == nil || task.Returns.TypeRef.Kind != semantic.TypeRefPrimitive {
		t.Fatalf("task return TypeRef = %#v", task)
	}
}

func TestExistingUC001EquivalentStillBuilds(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		typeRefStructModelFile("shop/model/user.yaml", "user"),
		{
			ID:   "shop/task/login.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Tasks: []rawyaml.Task{{
				ID:     "login",
				Main:   true,
				Params: []rawyaml.Param{{Name: "user", Model: "user"}},
				Returns: &rawyaml.Return{
					Name:  "user",
					Model: "user",
				},
			}}},
		},
	}})
	if len(diagnostics) != 0 {
		t.Fatalf("UC-001 equivalent diagnostics = %#v, want none", diagnostics)
	}
}

func TestTypeRefsCompatibleOK(t *testing.T) {
	project := typeRefCompatibilityProject(t)
	cases := []struct {
		name   string
		source string
		target string
	}{
		{name: "same primitive", source: "str", target: "str"},
		{name: "same named model", source: "user", target: "user"},
		{name: "any source", source: "any", target: "user"},
		{name: "any target", source: "user", target: "any"},
		{name: "same list", source: "list<user>", target: "list<user>"},
		{name: "list any source", source: "list<any>", target: "list<user>"},
		{name: "list any target", source: "list<user>", target: "list<any>"},
		{name: "same dict", source: "dict<config>", target: "dict<config>"},
		{name: "dict any source", source: "dict<any>", target: "dict<config>"},
		{name: "dict any target", source: "dict<config>", target: "dict<any>"},
		{name: "named list source", source: "user_list", target: "list<user>"},
		{name: "named list target", source: "list<user>", target: "user_list"},
		{name: "named dict source", source: "config_map", target: "dict<config>"},
		{name: "named dict target", source: "dict<config>", target: "config_map"},
		{name: "nested containers", source: "list<dict<user>>", target: "list<dict<user>>"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			source := mustParseTypeRefForTest(t, tc.source)
			target := mustParseTypeRefForTest(t, tc.target)
			if !typeRefsCompatible(project, source, target) {
				t.Fatalf("typeRefsCompatible(%s -> %s) = false, want true", tc.source, tc.target)
			}
		})
	}
}

func TestTypeRefsCompatibleNG(t *testing.T) {
	project := typeRefCompatibilityProject(t)
	cases := []struct {
		name   string
		source string
		target string
	}{
		{name: "different primitive", source: "str", target: "int"},
		{name: "primitive to named model", source: "str", target: "user"},
		{name: "different named model", source: "user", target: "order"},
		{name: "different list element", source: "list<user>", target: "list<order>"},
		{name: "different dict value", source: "dict<config>", target: "dict<user>"},
		{name: "list to dict", source: "list<user>", target: "dict<user>"},
		{name: "named list different element", source: "user_list", target: "list<order>"},
		{name: "named dict different value", source: "config_map", target: "dict<user>"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			source := mustParseTypeRefForTest(t, tc.source)
			target := mustParseTypeRefForTest(t, tc.target)
			if typeRefsCompatible(project, source, target) {
				t.Fatalf("typeRefsCompatible(%s -> %s) = true, want false", tc.source, tc.target)
			}
		})
	}
}

func TestNormalizeContainerTypeRefUsesProjectModelsByQID(t *testing.T) {
	project := typeRefCompatibilityProject(t)

	userList := normalizeContainerTypeRef(project, mustParseTypeRefForTest(t, "user_list"))
	if userList.Kind != semantic.TypeRefList || userList.Elem == nil || userList.Elem.Model != "shop.model.user" {
		t.Fatalf("normalized user_list = %#v, want list<shop.model.user>", userList)
	}

	configMap := normalizeContainerTypeRef(project, mustParseTypeRefForTest(t, "config_map"))
	if configMap.Kind != semantic.TypeRefDict || configMap.Value == nil || configMap.Value.Model != "shop.model.config" {
		t.Fatalf("normalized config_map = %#v, want dict<shop.model.config>", configMap)
	}

	user := normalizeContainerTypeRef(project, mustParseTypeRefForTest(t, "user"))
	if user.Kind != semantic.TypeRefNamedModel || user.Model != "shop.model.user" {
		t.Fatalf("normalized user = %#v, want nominal named model", user)
	}
}

func typeRefModelFile(fileID string, fieldType string) rawyaml.File {
	return rawyaml.File{
		ID:   fileID,
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{{
			ID:   "user",
			Kind: "struct",
			Fields: []rawyaml.ModelField{{
				Name: "profile",
				Type: fieldType,
			}},
		}}},
	}
}

func typeRefStructModelFile(fileID string, id string) rawyaml.File {
	return rawyaml.File{
		ID:   fileID,
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{{
			ID:   id,
			Kind: "struct",
			Fields: []rawyaml.ModelField{{
				Name: "id",
				Type: "str",
			}},
		}}},
	}
}

func typeRefCompatibilityProject(t *testing.T) *semantic.Project {
	t.Helper()
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		typeRefStructModelFile("shop/model/user.yaml", "user"),
		typeRefStructModelFile("shop/model/order.yaml", "order"),
		typeRefStructModelFile("shop/model/config.yaml", "config"),
		{
			ID:   "shop/model/user_list.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{{
				ID:      "user_list",
				Kind:    "list",
				Element: "user",
			}}},
		},
		{
			ID:   "shop/model/config_map.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{{
				ID:    "config_map",
				Kind:  "dict",
				Value: "config",
			}}},
		},
	}})
	if len(diagnostics) != 0 {
		t.Fatalf("typeRefCompatibilityProject diagnostics = %#v, want none", diagnostics)
	}
	return project
}

func mustParseTypeRefForTest(t *testing.T, raw string) *semantic.TypeRef {
	t.Helper()
	ref, err := parseTypeRef(raw, "shop")
	if err != nil {
		t.Fatalf("parseTypeRef(%q) returned error: %v", raw, err)
	}
	return ref
}
