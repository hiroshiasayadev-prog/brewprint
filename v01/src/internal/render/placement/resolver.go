package placement

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

type Severity string

const (
	SeverityError   Severity = "error"
	SeverityWarning Severity = "warning"
)

type Diagnostic struct {
	Severity Severity
	Message  string
}

type Group struct {
	ID       string
	Label    string
	Modules  []string
	Implicit bool
}

type Resolver struct {
	Groups        []Group
	ModuleToGroup map[string]string
}

var groupIDPattern = regexp.MustCompile(`^[a-z0-9_]+$`)

func NewResolver(raw *rawyaml.Project, project *semantic.Project) (*Resolver, []Diagnostic) {
	renderIndex := findRenderIndex(raw)
	modules := topLevelModules(project)
	groups, mapping, diags := resolveGroups(renderIndex, modules)
	return &Resolver{Groups: groups, ModuleToGroup: mapping}, diags
}

func findRenderIndex(raw *rawyaml.Project) *rawyaml.RenderIndex {
	if raw == nil {
		return nil
	}
	for _, file := range raw.Files {
		if file.Kind == rawyaml.FileKindRenderIndex && file.RenderIndex != nil {
			return file.RenderIndex
		}
	}
	return nil
}

func topLevelModules(project *semantic.Project) []string {
	seen := map[string]struct{}{}
	if project != nil {
		for _, node := range project.NodesByQID {
			module := topLevelModuleForFile(node.GetFileID())
			if module != "" {
				seen[module] = struct{}{}
			}
		}
	}
	modules := make([]string, 0, len(seen))
	for module := range seen {
		modules = append(modules, module)
	}
	sort.Strings(modules)
	return modules
}

func resolveGroups(index *rawyaml.RenderIndex, modules []string) ([]Group, map[string]string, []Diagnostic) {
	var diags []Diagnostic
	groups := []Group{}
	moduleToGroup := map[string]string{}
	knownModules := map[string]struct{}{}
	for _, module := range modules {
		knownModules[module] = struct{}{}
	}

	if index == nil || len(index.Groups) == 0 {
		if index != nil {
			diags = append(diags, Diagnostic{Severity: SeverityError, Message: "render_index.groups must not be empty"})
		}
		for _, module := range modules {
			if strings.HasPrefix(module, "_") {
				diags = append(diags, Diagnostic{Severity: SeverityWarning, Message: "module starts with reserved underscore and is skipped: " + module})
				continue
			}
			groups = append(groups, Group{ID: module, Label: module, Modules: []string{module}, Implicit: true})
			moduleToGroup[module] = module
		}
		return groups, moduleToGroup, diags
	}

	for _, rawGroup := range index.Groups {
		if rawGroup.ID == "" {
			diags = append(diags, Diagnostic{Severity: SeverityError, Message: "render_index group id is required"})
			continue
		}
		if strings.HasPrefix(rawGroup.ID, "_") {
			diags = append(diags, Diagnostic{Severity: SeverityError, Message: "render_index group id must not start with underscore: " + rawGroup.ID})
		}
		if !groupIDPattern.MatchString(rawGroup.ID) {
			diags = append(diags, Diagnostic{Severity: SeverityError, Message: "render_index group id must match [a-z0-9_]+: " + rawGroup.ID})
		}
		if len(rawGroup.Modules) == 0 {
			diags = append(diags, Diagnostic{Severity: SeverityError, Message: "render_index group modules must not be empty: " + rawGroup.ID})
		}

		group := Group{ID: rawGroup.ID, Label: rawGroup.Label, Modules: append([]string(nil), rawGroup.Modules...)}
		if group.Label == "" {
			group.Label = group.ID
		}
		groups = append(groups, group)

		for _, module := range rawGroup.Modules {
			if strings.HasPrefix(module, "_") {
				diags = append(diags, Diagnostic{Severity: SeverityError, Message: "render_index module must not start with underscore: " + module})
			}
			if strings.ContainsAny(module, "/\\.") {
				diags = append(diags, Diagnostic{Severity: SeverityError, Message: "render_index module must be top-level module: " + module})
			}
			if existing, exists := moduleToGroup[module]; exists {
				diags = append(diags, Diagnostic{Severity: SeverityError, Message: fmt.Sprintf("module %s belongs to multiple groups: %s, %s", module, existing, rawGroup.ID)})
				continue
			}
			moduleToGroup[module] = rawGroup.ID
		}
	}

	var uncovered []string
	for _, module := range modules {
		if _, ok := moduleToGroup[module]; ok {
			continue
		}
		if _, mentioned := knownModules[module]; !mentioned {
			continue
		}
		if strings.HasPrefix(module, "_") {
			diags = append(diags, Diagnostic{Severity: SeverityWarning, Message: "module starts with reserved underscore and is skipped: " + module})
			continue
		}
		uncovered = append(uncovered, module)
	}
	sort.Strings(uncovered)
	for _, module := range uncovered {
		diags = append(diags, Diagnostic{Severity: SeverityWarning, Message: "module is not covered by render_index and will use implicit group: " + module})
		groups = append(groups, Group{ID: module, Label: module, Modules: []string{module}, Implicit: true})
		moduleToGroup[module] = module
	}

	return groups, moduleToGroup, diags
}

func (r *Resolver) GroupForFile(fileID semantic.FileID) (string, bool) {
	if r == nil {
		return "", false
	}
	module := topLevelModuleForFile(fileID)
	group, ok := r.ModuleToGroup[module]
	return group, ok
}

func (r *Resolver) DAGPath(task *semantic.Task) (string, error) {
	if task == nil {
		return "", fmt.Errorf("task is nil")
	}
	group, ok := r.GroupForFile(task.FileID)
	if !ok || group == "" {
		return "", fmt.Errorf("no render group for task %s", task.QID)
	}
	return group + "/dag-" + task.ID + ".md", nil
}

func topLevelModuleForFile(fileID semantic.FileID) string {
	parts := strings.Split(fileID.String(), "/")
	if len(parts) == 0 {
		return ""
	}
	if parts[0] == "actors.yaml" || parts[0] == "views" || parts[0] == "render_index.yaml" {
		return ""
	}
	return parts[0]
}
