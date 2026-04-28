package state

import (
	"fmt"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func RenderFile(project *semantic.Project, fileID semantic.FileID) (string, error) {
	if project == nil {
		return "", fmt.Errorf("project is nil")
	}
	states := statesForFile(project, fileID)
	events := eventsForFile(project, fileID)
	transitions := append([]semantic.Transition(nil), project.TransitionsByFile[fileID]...)
	if len(states) == 0 && len(events) == 0 && len(transitions) == 0 {
		return "", fmt.Errorf("state diagram file not found or empty: %s", fileID)
	}

	var b strings.Builder
	b.WriteString("# " + fsmID(fileID) + "\n\n")
	b.WriteString("```mermaid\n")
	b.WriteString("stateDiagram-v2\n")

	choices := choiceIDs(transitions)
	if len(choices) > 0 {
		for _, choice := range choices {
			b.WriteString("  state " + choice + " <<choice>>\n")
		}
		b.WriteString("\n")
	}

	for _, state := range states {
		if state.Initial {
			b.WriteString("  [*] --> " + state.ID + "\n")
		}
	}
	for _, state := range states {
		if state.Final {
			b.WriteString("  " + state.ID + " --> [*]\n")
		}
	}
	if hasInitialOrFinal(states) {
		b.WriteString("\n")
	}

	for _, line := range transitionLines(transitions) {
		if line == "" {
			b.WriteString("\n")
			continue
		}
		b.WriteString("  " + line + "\n")
	}
	b.WriteString("```\n\n")

	b.WriteString("## States\n\n")
	b.WriteString("| state | note |\n")
	b.WriteString("|---|---|\n")
	for _, state := range states {
		b.WriteString("| " + state.ID + " | " + tableText(state.Note) + " |\n")
	}
	b.WriteString("\n")

	b.WriteString("## Events\n\n")
	b.WriteString("| event | source | actor | note |\n")
	b.WriteString("|---|---|---|---|\n")
	for _, event := range events {
		actor := event.Actor
		if actor == "" {
			actor = "—"
		}
		b.WriteString("| " + event.ID + " | " + event.Source + " | " + actor + " | " + tableText(event.Note) + " |\n")
	}
	b.WriteString("\n")

	return b.String(), nil
}

func statesForFile(project *semantic.Project, fileID semantic.FileID) []*semantic.State {
	var states []*semantic.State
	for _, node := range project.NodesByFile[fileID] {
		if state, ok := node.(*semantic.State); ok {
			states = append(states, state)
		}
	}
	return states
}

func eventsForFile(project *semantic.Project, fileID semantic.FileID) []*semantic.Event {
	var events []*semantic.Event
	for _, node := range project.NodesByFile[fileID] {
		if event, ok := node.(*semantic.Event); ok {
			events = append(events, event)
		}
	}
	return events
}

func hasInitialOrFinal(states []*semantic.State) bool {
	for _, state := range states {
		if state.Initial || state.Final {
			return true
		}
	}
	return false
}

func choiceIDs(transitions []semantic.Transition) []string {
	byPair := transitionsByPair(transitions)
	pairs := transitionPairsInOrder(transitions)
	var choices []string
	for _, key := range pairs {
		group := byPair[key]
		if len(group) < 2 {
			continue
		}
		parts := strings.Split(key, "\x00")
		choices = append(choices, choiceID(parts[0], parts[1]))
	}
	return choices
}

func transitionLines(transitions []semantic.Transition) []string {
	byPair := transitionsByPair(transitions)
	pairs := transitionPairsInOrder(transitions)

	var lines []string
	for i, pair := range pairs {
		group := byPair[pair]
		if len(group) < 2 {
			transition := group[0]
			lines = append(lines, transition.From+" --> "+transition.To+" : "+directLabel(transition))
			continue
		}
		from := group[0].From
		on := group[0].On
		choice := choiceID(from, on)
		lines = append(lines, from+" --> "+choice+" : "+on)
		for _, transition := range group {
			lines = append(lines, choice+" --> "+transition.To+" : "+choiceLabel(transition))
		}
		if i < len(pairs)-1 {
			lines = append(lines, "")
		}
	}
	return lines
}

func transitionsByPair(transitions []semantic.Transition) map[string][]semantic.Transition {
	byPair := map[string][]semantic.Transition{}
	for _, transition := range transitions {
		key := transition.From + "\x00" + transition.On
		byPair[key] = append(byPair[key], transition)
	}
	return byPair
}

func transitionPairsInOrder(transitions []semantic.Transition) []string {
	seen := map[string]struct{}{}
	var pairs []string
	for _, transition := range transitions {
		key := transition.From + "\x00" + transition.On
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		pairs = append(pairs, key)
	}
	return pairs
}

func directLabel(transition semantic.Transition) string {
	label := transition.On
	if transition.Guard != "" {
		label += " [" + transition.Guard + "]"
	}
	if showAction(transition.Action) {
		label += " / " + transition.Action
	}
	return label
}

func choiceLabel(transition semantic.Transition) string {
	label := "[" + transition.Guard + "]"
	if showAction(transition.Action) {
		label += " / " + transition.Action
	}
	return label
}

func showAction(action string) bool {
	return strings.Contains(action, ".")
}

func choiceID(from, on string) string {
	return "_choice_" + sanitizeID(from) + "_" + sanitizeID(on)
}

func sanitizeID(id string) string {
	id = strings.ReplaceAll(id, ".", "_")
	id = strings.ReplaceAll(id, "-", "_")
	id = strings.ReplaceAll(id, "/", "_")
	return id
}

func fsmID(fileID semantic.FileID) string {
	path := strings.TrimSuffix(fileID.String(), ".yaml")
	path = strings.TrimSuffix(path, ".yml")
	if strings.HasSuffix(path, "/state") {
		path = strings.TrimSuffix(path, "/state")
	}
	return strings.ReplaceAll(path, "/", "-")
}

func tableText(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "—"
	}
	lines := cleanedLines(s)
	if hasBulletLines(lines) {
		s = flattenBulletNote(lines)
	} else {
		s = strings.Join(lines, "")
	}
	s = strings.ReplaceAll(s, "|", "\\|")
	return s
}

func cleanedLines(s string) []string {
	rawLines := strings.Split(s, "\n")
	lines := make([]string, 0, len(rawLines))
	for _, line := range rawLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return lines
}

func hasBulletLines(lines []string) bool {
	for _, line := range lines {
		if strings.HasPrefix(line, "- ") {
			return true
		}
	}
	return false
}

func flattenBulletNote(lines []string) string {
	var out strings.Builder
	var bullets []string
	flushBullets := func() {
		if len(bullets) == 0 {
			return
		}
		out.WriteString(strings.Join(bullets, "と "))
		bullets = nil
	}
	for _, line := range lines {
		if strings.HasPrefix(line, "- ") {
			bullets = append(bullets, strings.TrimSpace(strings.TrimPrefix(line, "- ")))
			continue
		}
		flushBullets()
		out.WriteString(line)
	}
	flushBullets()
	return out.String()
}
