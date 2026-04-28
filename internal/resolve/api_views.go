package resolve

import (
	"fmt"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func buildAPIViews(raw *rawyaml.Project, project *semantic.Project, symbols *symbolTable) {
	for _, file := range raw.Files {
		if file.Kind != rawyaml.FileKindView || file.APIView == nil {
			continue
		}
		view := buildAPIView(file, symbols)
		if view == nil {
			continue
		}
		if _, exists := project.APIViewsByID[view.ID]; exists {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateView, semantic.FileID(file.ID), "duplicate API view id: "+view.ID)
			continue
		}
		project.APIViewsByID[view.ID] = view
	}
}

func buildAPIView(file rawyaml.File, symbols *symbolTable) *semantic.APIView {
	raw := file.APIView
	fileID := semantic.FileID(file.ID)
	if raw.ID == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidViewDefinition, fileID, "API view id is required")
		return nil
	}
	if raw.HTTPRootPath == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidViewDefinition, fileID, "API view http_root_path is required: "+raw.ID)
		return nil
	}
	if len(raw.Modules) == 0 {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidViewDefinition, fileID, "API view modules must not be empty: "+raw.ID)
		return nil
	}

	view := &semantic.APIView{
		FileID:       fileID,
		ID:           raw.ID,
		Note:         raw.Note,
		HTTPRootPath: raw.HTTPRootPath,
	}
	seen := map[string]struct{}{}
	for i, rawModule := range raw.Modules {
		if rawModule.Module == "" {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidViewDefinition, fileID, fmt.Sprintf("API view module is required at index %d: %s", i, raw.ID))
			continue
		}
		if _, exists := seen[rawModule.Module]; exists {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateViewModule, fileID, "duplicate API view module: "+rawModule.Module)
			continue
		}
		seen[rawModule.Module] = struct{}{}
		view.Modules = append(view.Modules, semantic.APIViewModule{Module: rawModule.Module, IncludeSubmodules: rawModule.IncludeSubmodules})
	}
	return view
}
