package resolve

import (
	"fmt"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func buildScenarios(raw *rawyaml.Project, project *semantic.Project, symbols *symbolTable) {
	for _, file := range raw.Files {
		if file.Kind != rawyaml.FileKindView || file.SequenceScenario == nil {
			continue
		}
		scenario := buildScenario(file, project, symbols)
		if scenario == nil {
			continue
		}
		if _, exists := project.ScenariosByID[scenario.ID]; exists {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateView, semantic.FileID(file.ID), "duplicate sequence scenario id: "+scenario.ID)
			continue
		}
		project.ScenariosByID[scenario.ID] = scenario
	}
}

func buildScenario(file rawyaml.File, project *semantic.Project, symbols *symbolTable) *semantic.SequenceScenario {
	raw := file.SequenceScenario
	fileID := semantic.FileID(file.ID)
	if raw.ID == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidViewDefinition, fileID, "sequence scenario id is required")
		return nil
	}
	if raw.StateFile == "" {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidViewDefinition, fileID, "sequence scenario state_file is required: "+raw.ID)
		return nil
	}
	if len(raw.Steps) == 0 {
		symbols.addDiagnosticCode(semantic.SeverityError, diagnosticInvalidViewDefinition, fileID, "sequence scenario steps must not be empty: "+raw.ID)
		return nil
	}

	scenario := &semantic.SequenceScenario{
		FileID:    fileID,
		ID:        raw.ID,
		Title:     raw.Title,
		StateFile: semantic.FileID(raw.StateFile),
	}
	for i, rawStep := range raw.Steps {
		transition, ok := resolveScenarioTransition(project, scenario.StateFile, rawStep)
		if !ok {
			symbols.addDiagnosticCode(semantic.SeverityError, diagnosticUnresolvedSequenceStep, fileID, fmt.Sprintf("unresolved sequence step %d: %s via %s", i+1, rawStep.FromState, rawStep.Via))
			continue
		}
		if len(scenario.Steps) > 0 {
			prev := scenario.Steps[len(scenario.Steps)-1].Transition
			if prev.To != rawStep.FromState {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticNonContinuousSequence, fileID, fmt.Sprintf("sequence step %d is not continuous: previous to=%s, current from_state=%s", i+1, prev.To, rawStep.FromState))
			}
		}
		scenario.Steps = append(scenario.Steps, semantic.SequenceStep{
			FromState:  rawStep.FromState,
			Via:        rawStep.Via,
			Guard:      rawStep.Guard,
			Transition: transition,
		})
	}
	return scenario
}

func resolveScenarioTransition(project *semantic.Project, stateFile semantic.FileID, step rawyaml.SequenceStep) (semantic.Transition, bool) {
	fromState := resolveNodeQID(project, stateFile, semantic.NodeKindState, step.FromState)
	event := resolveNodeQID(project, stateFile, semantic.NodeKindEvent, step.Via)
	if fromState == "" || event == "" {
		return semantic.Transition{}, false
	}
	eventKey := semantic.TransitionEventKey{StateFile: stateFile, FromState: fromState, Event: event}
	candidates := project.TransitionsByStateEvent[eventKey]
	if len(candidates) == 0 {
		return semantic.Transition{}, false
	}
	if len(candidates) == 1 {
		transition := candidates[0].Transition
		if step.Guard == "" || transition.Guard == step.Guard {
			return transition, true
		}
		return semantic.Transition{}, false
	}
	if step.Guard == "" {
		return semantic.Transition{}, false
	}
	key := semantic.TransitionKey{StateFile: stateFile, FromState: fromState, Event: event, Guard: step.Guard}
	ref, ok := project.TransitionsByStateEventGuard[key]
	if !ok {
		return semantic.Transition{}, false
	}
	return ref.Transition, true
}
