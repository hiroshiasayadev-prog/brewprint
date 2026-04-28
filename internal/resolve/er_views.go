package resolve

import (
	"fmt"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func buildERViews(raw *rawyaml.Project, project *semantic.Project, symbols *symbolTable) {
	for _, file := range raw.Files {
		if file.Kind != rawyaml.FileKindView || file.ERView == nil {
			continue
		}
		view := buildERView(file, symbols)
		if view == nil {
			continue
		}
		if _, exists := project.ERViewsByID[view.ID]; exists {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateView, semantic.FileID(file.ID), "duplicate ER view id: "+view.ID)
			continue
		}
		project.ERViewsByID[view.ID] = view
	}
}

func buildERView(file rawyaml.File, symbols *symbolTable) *semantic.ERView {
	raw := file.ERView
	fileID := semantic.FileID(file.ID)
	if raw.ID == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidViewDefinition, fileID, "ER view id is required")
		return nil
	}
	if len(raw.Modules) == 0 {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidViewDefinition, fileID, "ER view modules must not be empty: "+raw.ID)
		return nil
	}

	view := &semantic.ERView{
		FileID: fileID,
		ID:     raw.ID,
		Note:   raw.Note,
	}
	seen := map[string]struct{}{}
	for i, rawModule := range raw.Modules {
		if rawModule.Module == "" {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidViewDefinition, fileID, fmt.Sprintf("ER view module is required at index %d: %s", i, raw.ID))
			continue
		}
		if _, exists := seen[rawModule.Module]; exists {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateViewModule, fileID, "duplicate ER view module: "+rawModule.Module)
			continue
		}
		seen[rawModule.Module] = struct{}{}
		view.Modules = append(view.Modules, semantic.ERViewModule{Module: rawModule.Module})
	}
	return view
}
