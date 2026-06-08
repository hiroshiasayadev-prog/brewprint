package wireframe

import (
	"fmt"
	"html"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

func RenderState(project *semantic.Project, stateQID semantic.QualifiedID) (string, error) {
	if project == nil {
		return "", fmt.Errorf("project is nil")
	}
	state := project.StatesByQID[stateQID]
	if state == nil {
		return "", fmt.Errorf("state not found: %s", stateQID)
	}
	if state.Wireframe == nil {
		return "", fmt.Errorf("state has no wireframe: %s", stateQID)
	}
	return renderElement(*state.Wireframe, 0), nil
}

func RenderPreview(project *semantic.Project, title string) string {
	if title == "" {
		title = "Wireframe Preview"
	}
	states := statesWithWireframe(project)

	var b strings.Builder
	b.WriteString("<!doctype html>\n")
	b.WriteString("<html lang=\"ja\">\n")
	b.WriteString("<head>\n")
	b.WriteString("  <meta charset=\"utf-8\">\n")
	b.WriteString("  <title>" + escape(title) + "</title>\n")
	b.WriteString("  <link rel=\"stylesheet\" href=\"../../../../spec/views/wireframe.css\">\n")
	b.WriteString("  <link rel=\"stylesheet\" href=\"../../../../spec/views/wireframe.preview.css\">\n")
	b.WriteString("</head>\n")
	b.WriteString("<body>\n")
	b.WriteString("  <main class=\"wf-preview-page\">\n")
	b.WriteString("    <h1>" + escape(title) + "</h1>\n\n")
	for i, state := range states {
		if i > 0 {
			b.WriteString("\n")
		}
		b.WriteString("    <section>\n")
		b.WriteString("      <h2>" + escape(stateLabel(state)) + "</h2>\n")
		b.WriteString("      <div class=\"wf-preview-frame\">\n")
		fragment := strings.TrimRight(renderElement(*state.Wireframe, 0), "\n")
		for _, line := range strings.Split(fragment, "\n") {
			b.WriteString("        " + line + "\n")
		}
		b.WriteString("      </div>\n")
		b.WriteString("    </section>\n")
	}
	b.WriteString("  </main>\n")
	b.WriteString("</body>\n")
	b.WriteString("</html>\n")
	return b.String()
}

func statesWithWireframe(project *semantic.Project) []*semantic.State {
	if project == nil {
		return nil
	}
	var states []*semantic.State
	for _, state := range project.StatesByQID {
		if state.Wireframe != nil {
			states = append(states, state)
		}
	}
	return orderStatesByProjectTraversal(project, states)
}

func orderStatesByProjectTraversal(project *semantic.Project, states []*semantic.State) []*semantic.State {
	wanted := map[semantic.QualifiedID]*semantic.State{}
	for _, state := range states {
		wanted[state.QID] = state
	}

	fileIDs := make([]semantic.FileID, 0, len(project.NodesByFile))
	for fileID := range project.NodesByFile {
		fileIDs = append(fileIDs, fileID)
	}
	sort.Slice(fileIDs, func(i, j int) bool { return fileIDs[i] < fileIDs[j] })

	out := make([]*semantic.State, 0, len(states))
	seen := map[semantic.QualifiedID]struct{}{}
	for _, fileID := range fileIDs {
		for _, node := range project.NodesByFile[fileID] {
			state, ok := node.(*semantic.State)
			if !ok || state.Wireframe == nil {
				continue
			}
			if _, ok := wanted[state.QID]; !ok {
				continue
			}
			out = append(out, state)
			seen[state.QID] = struct{}{}
		}
	}
	for _, state := range states {
		if _, ok := seen[state.QID]; ok {
			continue
		}
		out = append(out, state)
	}
	return out
}

func stateLabel(state *semantic.State) string {
	module := moduleForFile(state.FileID)
	if module == "" {
		return state.ID
	}
	return module + "." + state.ID
}

func renderElement(element semantic.WireframeElement, depth int) string {
	var b strings.Builder
	renderElementInto(&b, element, depth)
	return b.String()
}

func renderElementInto(b *strings.Builder, element semantic.WireframeElement, depth int) {
	indent := strings.Repeat("  ", depth)
	switch element.Type {
	case "col", "row", "grid":
		b.WriteString(indent + "<div" + attrs(element, "wf-"+element.Type) + ">\n")
		writeChildren(b, element.Children, depth+1)
		b.WriteString(indent + "</div>\n")
	case "card":
		b.WriteString(indent + "<section" + attrs(element, "wf-card") + ">\n")
		writeChildren(b, element.Children, depth+1)
		b.WriteString(indent + "</section>\n")
	case "sidebar":
		b.WriteString(indent + "<aside" + attrs(element, "wf-sidebar") + ">\n")
		writeChildren(b, element.Children, depth+1)
		b.WriteString(indent + "</aside>\n")
	case "header":
		b.WriteString(indent + "<header" + attrs(element, "wf-header") + ">\n")
		writeChildren(b, element.Children, depth+1)
		b.WriteString(indent + "</header>\n")
	case "footer":
		b.WriteString(indent + "<footer" + attrs(element, "wf-footer") + ">\n")
		writeChildren(b, element.Children, depth+1)
		b.WriteString(indent + "</footer>\n")
	case "main":
		b.WriteString(indent + "<main" + attrs(element, "wf-main") + ">\n")
		writeChildren(b, element.Children, depth+1)
		b.WriteString(indent + "</main>\n")
	case "text":
		b.WriteString(indent + "<span" + attrs(element, "wf-text") + ">" + escape(element.Label) + "</span>\n")
	case "badge":
		b.WriteString(indent + "<span" + attrs(element, "wf-badge") + ">" + escape(element.Label) + "</span>\n")
	case "icon":
		b.WriteString(indent + "<span" + attrs(element, "wf-icon") + ">[icon]</span>\n")
	case "image":
		b.WriteString(indent + "<div" + attrs(element, "wf-image") + ">[image]</div>\n")
	case "divider":
		b.WriteString(indent + "<hr" + attrs(element, "wf-divider") + " />\n")
	case "button":
		disabled := ""
		if element.Disabled {
			disabled = " disabled"
		}
		b.WriteString(indent + "<button" + attrs(element, "wf-button") + disabled + ">" + escape(element.Label) + "</button>\n")
	case "input":
		renderField(b, element, depth, "text")
	case "password":
		renderField(b, element, depth, "password")
	case "select":
		b.WriteString(indent + "<div" + attrs(element, "wf-field") + ">\n")
		b.WriteString(indent + "  <label>" + escape(element.Label) + "</label>\n")
		b.WriteString(indent + "  <select")
		if element.Disabled {
			b.WriteString(" disabled")
		}
		b.WriteString("></select>\n")
		b.WriteString(indent + "</div>\n")
	case "checkbox":
		renderCheckable(b, element, depth, "checkbox", "wf-checkbox")
	case "radio":
		renderCheckable(b, element, depth, "radio", "wf-radio")
	default:
		b.WriteString(indent + "<div" + attrs(element, "wf-unknown") + "></div>\n")
	}
}

func writeChildren(b *strings.Builder, children []semantic.WireframeElement, depth int) {
	for _, child := range children {
		renderElementInto(b, child, depth)
	}
}

func renderField(b *strings.Builder, element semantic.WireframeElement, depth int, inputType string) {
	indent := strings.Repeat("  ", depth)
	b.WriteString(indent + "<div" + attrs(element, "wf-field") + ">\n")
	b.WriteString(indent + "  <label>" + escape(element.Label) + "</label>\n")
	b.WriteString(indent + "  <input type=\"" + inputType + "\"")
	if element.Placeholder != "" {
		b.WriteString(" placeholder=\"" + escape(element.Placeholder) + "\"")
	}
	if element.Disabled {
		b.WriteString(" disabled")
	}
	b.WriteString(" />\n")
	b.WriteString(indent + "</div>\n")
}

func renderCheckable(b *strings.Builder, element semantic.WireframeElement, depth int, inputType, class string) {
	indent := strings.Repeat("  ", depth)
	b.WriteString(indent + "<label" + attrs(element, class) + "><input type=\"" + inputType + "\"")
	if element.Disabled {
		b.WriteString(" disabled")
	}
	b.WriteString(" /> " + escape(element.Label) + "</label>\n")
}

func attrs(element semantic.WireframeElement, class string) string {
	parts := []string{"class=\"" + class + "\""}
	if element.ID != "" {
		parts = append(parts, "data-wf-id=\""+escape(element.ID)+"\"")
	}
	if element.Fires != "" {
		parts = append(parts, "data-wf-fires=\""+escape(element.Fires)+"\"")
	}
	style := styleAttr(element)
	if style != "" {
		parts = append(parts, "style=\""+escape(style)+"\"")
	}
	return " " + strings.Join(parts, " ")
}

func styleAttr(element semantic.WireframeElement) string {
	var styles []string
	if element.Type == "grid" && element.Cols > 0 {
		styles = append(styles, fmt.Sprintf("grid-template-columns: repeat(%d, 1fr)", element.Cols))
	}
	if element.Span > 0 {
		styles = append(styles, fmt.Sprintf("grid-column: span %d", element.Span))
	}
	if element.Layout != nil {
		styles = append(styles, layoutStyles(*element.Layout)...)
	}
	if len(styles) == 0 {
		return ""
	}
	return strings.Join(styles, "; ") + ";"
}

func layoutStyles(layout semantic.WireframeLayout) []string {
	var styles []string
	if value := sizeStyle("width", layout.Width); value != "" {
		styles = append(styles, value)
	}
	if value := sizeStyle("height", layout.Height); value != "" {
		styles = append(styles, value)
		if n, ok := numberValue(layout.Height); ok {
			styles = append(styles, fmt.Sprintf("min-height: %dpx", n))
		}
	}
	if layout.MinWidth > 0 {
		styles = append(styles, fmt.Sprintf("min-width: %dpx", layout.MinWidth))
	}
	if layout.MinHeight > 0 {
		styles = append(styles, fmt.Sprintf("min-height: %dpx", layout.MinHeight))
	}
	if layout.Grow {
		styles = append(styles, "flex: 1 1 0%", "min-width: 0", "min-height: 0")
	}
	if layout.Gap > 0 {
		styles = append(styles, fmt.Sprintf("gap: %dpx", layout.Gap))
	}
	if padding := paddingStyle(layout.Padding); padding != "" {
		styles = append(styles, padding)
	}
	if align := alignStyle(layout.Align); align != "" {
		styles = append(styles, align)
	}
	if justify := justifyStyle(layout.Justify); justify != "" {
		styles = append(styles, justify)
	}
	if scroll := scrollStyle(layout.Scroll); scroll != "" {
		styles = append(styles, scroll)
	}
	return styles
}

func sizeStyle(name string, value any) string {
	n, ok := numberValue(value)
	if ok {
		return fmt.Sprintf("%s: %dpx", name, n)
	}
	s, ok := value.(string)
	if !ok || s == "" {
		return ""
	}
	switch s {
	case "fill":
		return name + ": 100%"
	case "fit":
		if name == "height" {
			return name + ": auto"
		}
		return name + ": fit-content"
	default:
		return ""
	}
}

func numberValue(value any) (int, bool) {
	switch v := value.(type) {
	case int:
		return v, true
	case int64:
		return int(v), true
	case float64:
		return int(v), true
	case float32:
		return int(v), true
	default:
		return 0, false
	}
}

func paddingStyle(value any) string {
	if n, ok := numberValue(value); ok {
		return fmt.Sprintf("padding: %dpx", n)
	}
	m, ok := value.(map[string]any)
	if !ok {
		return ""
	}
	top, right, bottom, left := 0, 0, 0, 0
	if x, ok := numberValue(m["x"]); ok {
		right, left = x, x
	}
	if y, ok := numberValue(m["y"]); ok {
		top, bottom = y, y
	}
	if v, ok := numberValue(m["top"]); ok {
		top = v
	}
	if v, ok := numberValue(m["right"]); ok {
		right = v
	}
	if v, ok := numberValue(m["bottom"]); ok {
		bottom = v
	}
	if v, ok := numberValue(m["left"]); ok {
		left = v
	}
	return fmt.Sprintf("padding: %dpx %dpx %dpx %dpx", top, right, bottom, left)
}

func alignStyle(value string) string {
	switch value {
	case "start":
		return "align-items: flex-start"
	case "center":
		return "align-items: center"
	case "end":
		return "align-items: flex-end"
	case "stretch":
		return "align-items: stretch"
	default:
		return ""
	}
}

func justifyStyle(value string) string {
	switch value {
	case "start":
		return "justify-content: flex-start"
	case "center":
		return "justify-content: center"
	case "end":
		return "justify-content: flex-end"
	case "between":
		return "justify-content: space-between"
	default:
		return ""
	}
}

func scrollStyle(value string) string {
	switch value {
	case "none":
		return "overflow: visible"
	case "x":
		return "overflow-x: auto"
	case "y":
		return "overflow-y: auto"
	case "both":
		return "overflow: auto"
	default:
		return ""
	}
}

func moduleForFile(fileID semantic.FileID) string {
	parts := strings.Split(fileID.String(), "/")
	if len(parts) == 0 {
		return ""
	}
	for i, part := range parts[:len(parts)-1] {
		if part == "model" || part == "store" || part == "task" || part == "state" || part == "event" || part == "branch" || part == "fork" || part == "join" {
			return strings.Join(parts[:i], ".")
		}
	}
	if len(parts) == 1 {
		return ""
	}
	return strings.Join(parts[:len(parts)-1], ".")
}

func escape(s string) string {
	return html.EscapeString(s)
}
