package dag

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

func writePrivateModelsSection(b *strings.Builder, project *semantic.Project, fileID semantic.FileID) {
	models := privateModelsForFile(project, fileID)
	if len(models) == 0 {
		return
	}
	b.WriteString("\n## Private models\n\n")
	b.WriteString("| model | kind | used by | shape | note |\n")
	b.WriteString("|---|---|---|---|---|\n")
	for _, model := range models {
		fmt.Fprintf(b, "| %s | %s | %s | %s | %s |\n",
			model.ID,
			tableCell(model.Kind),
			tableCell(strings.Join(privateModelUsedBy(project, fileID, model.QID), "<br/>")),
			tableCell(privateModelShape(model)),
			tableCell(model.Note),
		)
	}
	b.WriteString("\n")
}

func privateModelsForFile(project *semantic.Project, fileID semantic.FileID) []*semantic.Model {
	if project == nil {
		return nil
	}
	byName := project.PrivateModelsByFile[fileID]
	models := make([]*semantic.Model, 0, len(byName))
	for _, model := range byName {
		models = append(models, model)
	}
	sort.Slice(models, func(i, j int) bool { return models[i].ID < models[j].ID })
	return models
}

func privateModelUsedBy(project *semantic.Project, fileID semantic.FileID, target semantic.QualifiedID) []string {
	if project == nil || target == "" {
		return nil
	}
	var out []string
	for _, node := range project.NodesByFile[fileID] {
		switch n := node.(type) {
		case *semantic.Task:
			out = append(out, paramModelUses(n.ID, n.Params, target)...)
			if n.Returns != nil && typeRefReferencesModel(n.Returns.TypeRef, target) {
				out = append(out, n.ID+".returns")
			}
		case *semantic.Branch:
			out = append(out, paramModelUses(n.ID, n.Params, target)...)
		case *semantic.Fork:
			out = append(out, paramModelUses(n.ID, n.Params, target)...)
		case *semantic.Join:
			out = append(out, paramModelUses(n.ID, n.Params, target)...)
			if n.Returns != nil && typeRefReferencesModel(n.Returns.TypeRef, target) {
				out = append(out, n.ID+".returns")
			}
		case *semantic.Model:
			for _, field := range n.Fields {
				if typeRefReferencesModel(field.TypeRef, target) {
					out = append(out, n.ID+"."+field.Name)
				}
			}
			if typeRefReferencesModel(n.ElementRef, target) {
				out = append(out, n.ID+".element")
			}
			if typeRefReferencesModel(n.ValueRef, target) {
				out = append(out, n.ID+".value")
			}
		}
	}
	if len(out) == 0 {
		return []string{"-"}
	}
	return out
}

func paramModelUses(parentID string, params []semantic.Param, target semantic.QualifiedID) []string {
	var out []string
	for _, param := range params {
		if typeRefReferencesModel(param.TypeRef, target) {
			out = append(out, parentID+".param:"+param.Name)
		}
	}
	return out
}

func typeRefReferencesModel(ref *semantic.TypeRef, target semantic.QualifiedID) bool {
	if ref == nil || target == "" {
		return false
	}
	switch ref.Kind {
	case semantic.TypeRefNamedModel:
		return ref.Model == target
	case semantic.TypeRefList:
		return typeRefReferencesModel(ref.Elem, target)
	case semantic.TypeRefDict:
		return typeRefReferencesModel(ref.Value, target)
	default:
		return false
	}
}

func privateModelShape(model *semantic.Model) string {
	switch model.Kind {
	case "struct":
		if len(model.Fields) == 0 {
			return "-"
		}
		parts := make([]string, 0, len(model.Fields))
		for _, field := range model.Fields {
			parts = append(parts, field.Name+": "+field.Type)
		}
		return strings.Join(parts, "<br/>")
	case "list":
		return "list<" + model.Element + ">"
	case "dict":
		return "dict<" + model.Value + ">"
	case "enum":
		if len(model.Values) == 0 {
			return "-"
		}
		return strings.Join(model.Values, "<br/>")
	case "tagged_union":
		if model.Discriminator == "" {
			return "-"
		}
		parts := make([]string, 0, 1+len(model.Variants))
		parts = append(parts, "discriminator: "+model.Discriminator)
		for _, v := range model.Variants {
			parts = append(parts, "tag: "+v.Tag)
		}
		return strings.Join(parts, "<br/>")
	default:
		return "-"
	}
}

func tableCell(value string) string {
	value = strings.TrimSpace(value)
	if value == "" || value == "-" {
		return "—"
	}
	value = strings.ReplaceAll(value, "\r\n", "\n")
	value = strings.ReplaceAll(value, "\n", "<br/>")
	value = strings.ReplaceAll(value, "|", `\|`)
	return value
}
