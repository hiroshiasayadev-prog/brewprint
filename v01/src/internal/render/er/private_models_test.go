package er

import (
	"strings"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

func TestRenderERExcludesTaskFileHelperModels(t *testing.T) {
	project, diagnostics := resolve.Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "auth/model/user.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{{
				ID:     "user",
				Kind:   "struct",
				Fields: []rawyaml.ModelField{{Name: "id", Type: "str", PK: true}},
			}}},
		},
		{
			ID:   "auth/store/user_db.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Stores: []rawyaml.Store{{
				ID:   "user_db",
				Kind: "db",
				Of:   "user",
			}}},
		},
		{
			ID:   "auth/task/login.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{
				Tasks: []rawyaml.Task{{
					ID:      "login",
					Main:    true,
					Returns: &rawyaml.Return{Name: "token", Model: "login_token"},
				}},
				Models: []rawyaml.Model{{
					ID:     "login_token",
					Kind:   "struct",
					Fields: []rawyaml.ModelField{{Name: "access_token", Type: "str"}},
				}},
			},
		},
		{
			ID:   "views/er.yaml",
			Kind: rawyaml.FileKindView,
			ERView: &rawyaml.ERView{
				ID:      "auth_er",
				Modules: []rawyaml.ERViewModule{{Module: "auth"}},
			},
		},
	}})
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == semantic.SeverityError {
			t.Fatalf("semantic diagnostic: %s: %s", diagnostic.FileID, diagnostic.Message)
		}
	}

	actual, err := RenderView(project, "auth_er")
	if err != nil {
		t.Fatalf("RenderView returned error: %v", err)
	}
	if !strings.Contains(actual, "  user {") {
		t.Fatalf("ER render missing public model:\n%s", actual)
	}
	if strings.Contains(actual, "login_token") || strings.Contains(actual, "access_token") {
		t.Fatalf("ER render included task-file helper model:\n%s", actual)
	}
}
