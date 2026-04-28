package resolve

import (
	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func buildTransitions(raw *rawyaml.Project, project *semantic.Project, symbols *symbolTable) {
	for _, file := range raw.Files {
		if file.Kind != rawyaml.FileKindNode || file.NodeFile == nil || len(file.NodeFile.Transitions) == 0 {
			continue
		}
		fileID := semantic.FileID(file.ID)
		var transitions []semantic.Transition
		for _, rawTransition := range file.NodeFile.Transitions {
			transition := semantic.Transition{
				FileID: fileID,
				From:   rawTransition.From,
				On:     rawTransition.On,
				To:     rawTransition.To,
				Action: rawTransition.Action,
				Guard:  rawTransition.Guard,
				Note:   rawTransition.Note,
			}
			transition.FromState = resolveNodeQID(project, fileID, semantic.NodeKindState, rawTransition.From)
			transition.Event = resolveNodeQID(project, fileID, semantic.NodeKindEvent, rawTransition.On)
			transition.ToState = resolveNodeQID(project, fileID, semantic.NodeKindState, rawTransition.To)
			if rawTransition.Action != "" {
				transition.ActionTask = resolveNodeQID(project, fileID, semantic.NodeKindTask, rawTransition.Action)
				if transition.ActionTask == "" {
					transition.ActionTask = semantic.QualifiedID(rawTransition.Action)
				}
			}
			if transition.FromState == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedTransitionState, fileID, "unresolved transition from state: "+rawTransition.From)
			}
			if transition.Event == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedTransitionEvent, fileID, "unresolved transition event: "+rawTransition.On)
			}
			if transition.ToState == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedTransitionState, fileID, "unresolved transition to state: "+rawTransition.To)
			}
			transitions = append(transitions, transition)
		}
		project.TransitionsByFile[fileID] = transitions
		validateGuardBranches(fileID, transitions, symbols)
	}
}

func validateGuardBranches(fileID semantic.FileID, transitions []semantic.Transition, symbols *symbolTable) {
	byPair := map[string][]semantic.Transition{}
	seenKey := map[string]struct{}{}
	for _, transition := range transitions {
		pair := transition.From + "\x00" + transition.On
		byPair[pair] = append(byPair[pair], transition)
		key := pair + "\x00" + transition.Guard
		if _, exists := seenKey[key]; exists {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateTransition, fileID, "duplicate transition: "+transition.From+" on "+transition.On+" guard "+transition.Guard)
		}
		seenKey[key] = struct{}{}
	}
	for _, group := range byPair {
		if len(group) < 2 {
			continue
		}
		for _, transition := range group {
			if transition.Guard == "" {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticMissingTransitionGuard, fileID, "guard is required for branched transition: "+transition.From+" on "+transition.On)
			}
		}
	}
}
