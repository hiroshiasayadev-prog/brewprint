package dag

import (
	"strings"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func TestRenderDAGPrivateModelsSection(t *testing.T) {
	project, diagnostics := resolve.Build(&rawyaml.Project{Files: []rawyaml.File{{
		ID:   "auth/task/login.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{
			Tasks: []rawyaml.Task{{
				ID:      "login",
				Main:    true,
				Params:  []rawyaml.Param{{Name: "form", Model: "login_form"}},
				Returns: &rawyaml.Return{Name: "token", Model: "login_token"},
			}},
			Models: []rawyaml.Model{
				{
					ID:     "login_form",
					Kind:   "struct",
					Fields: []rawyaml.ModelField{{Name: "username", Type: "str"}, {Name: "credentials", Type: "credentials"}},
					Note:   "request shape",
				},
				{
					ID:      "credentials",
					Kind:    "list",
					Element: "str",
				},
				{
					ID:     "login_token",
					Kind:   "struct",
					Fields: []rawyaml.ModelField{{Name: "access_token", Type: "str"}},
				},
			},
		},
	}}})
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == semantic.SeverityError && diagnostic.Code != "invalid_private_model_reference" {
			t.Fatalf("semantic diagnostic: %s: %s", diagnostic.FileID, diagnostic.Message)
		}
	}

	actual, err := RenderFile(project, "auth/task/login.yaml")
	if err != nil {
		t.Fatalf("RenderFile returned error: %v", err)
	}
	if !strings.Contains(actual, "## Private models") {
		t.Fatalf("render missing Private models section:\n%s", actual)
	}
	if !strings.Contains(actual, "| login_form | struct | login.param:form | username: str<br/>credentials: credentials | request shape |") {
		t.Fatalf("render missing login_form private model row:\n%s", actual)
	}
	if !strings.Contains(actual, "| credentials | list | login_form.credentials | list<str> | — |") {
		t.Fatalf("render missing nested helper used-by row:\n%s", actual)
	}
	mermaid := mermaidBlock(actual)
	// login_form and login_token appear as ADR-074 TypeRef hints on param/returns labels — that is correct.
	// credentials is a nested helper model not directly referenced as a param/returns TypeRef — must not appear.
	if strings.Contains(mermaid, "credentials") {
		t.Fatalf("nested helper model 'credentials' leaked into Mermaid body:\n%s", mermaid)
	}
	if !strings.Contains(mermaid, "form([form: login_form])") {
		t.Fatalf("missing expected TypeRef hint 'login_form' on param 'form':\n%s", mermaid)
	}
	if !strings.Contains(mermaid, "token([token: login_token])") {
		t.Fatalf("missing expected TypeRef hint 'login_token' on returns 'token':\n%s", mermaid)
	}
}

func mermaidBlock(markdown string) string {
	start := strings.Index(markdown, "```mermaid")
	if start < 0 {
		return ""
	}
	rest := markdown[start+len("```mermaid"):]
	end := strings.Index(rest, "```")
	if end < 0 {
		return rest
	}
	return rest[:end]
}
