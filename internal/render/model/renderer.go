package model

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func RenderFile(project *semantic.Project, fileID semantic.FileID) (string, error) {
	model := publicModelForFile(project, fileID)
	if model == nil {
		return "", fmt.Errorf("public model not found for file %s", fileID)
	}
	return RenderModel(project, model)
}

func RenderModel(project *semantic.Project, model *semantic.Model) (string, error) {
	if model == nil {
		return "", fmt.Errorf("model is nil")
	}

	var b strings.Builder
	fmt.Fprintf(&b, "# %s\n\n", model.ID)
	if note := strings.TrimSpace(model.Note); note != "" {
		b.WriteString(note)
		b.WriteString("\n\n")
	}

	b.WriteString("## Public model\n\n")
	b.WriteString("| property | value |\n")
	b.WriteString("|---|---|\n")
	fmt.Fprintf(&b, "| kind | %s |\n", tableCell(model.Kind))
	b.WriteString("| visibility | public |\n")
	fmt.Fprintf(&b, "| source | %s |\n\n", tableCell("yaml/"+model.FileID.String()))

	writeKindSection(&b, model)
	writePrivateModelsSection(&b, project, model.FileID)
	return b.String(), nil
}

func publicModelForFile(project *semantic.Project, fileID semantic.FileID) *semantic.Model {
	if project == nil {
		return nil
	}
	if qid := project.MainNodeByFile[fileID]; qid != "" {
		if model := project.ModelsByQID[qid]; model != nil {
			return model
		}
	}

	var models []*semantic.Model
	for _, model := range project.ModelsByQID {
		if model.FileID == fileID && !model.FilePrivate {
			models = append(models, model)
		}
	}
	if len(models) != 1 {
		return nil
	}
	return models[0]
}

func writeKindSection(b *strings.Builder, model *semantic.Model) {
	switch model.Kind {
	case "struct":
		b.WriteString("### Fields\n\n")
		b.WriteString("| field | type | note |\n")
		b.WriteString("|---|---|---|\n")
		for _, field := range model.Fields {
			fmt.Fprintf(b, "| %s | %s | %s |\n", tableCell(field.Name), tableCell(field.Type), tableCell(field.Note))
		}
		b.WriteString("\n")
	case "enum":
		b.WriteString("### Values\n\n")
		b.WriteString("| value | note |\n")
		b.WriteString("|---|---|\n")
		for _, value := range model.Values {
			fmt.Fprintf(b, "| %s | %s |\n", tableCell(value), tableCell(""))
		}
		b.WriteString("\n")
	case "list":
		b.WriteString("### Element\n\n")
		b.WriteString("| property | value |\n")
		b.WriteString("|---|---|\n")
		fmt.Fprintf(b, "| element | %s |\n\n", tableCell(model.Element))
	case "dict":
		b.WriteString("### Value\n\n")
		b.WriteString("| property | value |\n")
		b.WriteString("|---|---|\n")
		fmt.Fprintf(b, "| value | %s |\n\n", tableCell(model.Value))
	case "tagged_union":
		b.WriteString("### Discriminator\n\n")
		b.WriteString("| property | value |\n")
		b.WriteString("|---|---|\n")
		fmt.Fprintf(b, "| discriminator | %s |\n\n", tableCell(model.Discriminator))
		b.WriteString("### Variants\n\n")
		for _, variant := range model.Variants {
			fmt.Fprintf(b, "#### `%s`\n\n", variant.Tag)
			if len(variant.Fields) == 0 {
				b.WriteString("No payload fields.\n\n")
			} else {
				b.WriteString("| field | type | note |\n")
				b.WriteString("|---|---|---|\n")
				for _, field := range variant.Fields {
					fmt.Fprintf(b, "| %s | %s | %s |\n", tableCell(field.Name), tableCell(field.Type), tableCell(field.Note))
				}
				b.WriteString("\n")
			}
		}
	}
}

func writePrivateModelsSection(b *strings.Builder, project *semantic.Project, fileID semantic.FileID) {
	models := privateModelsForFile(project, fileID)
	if len(models) == 0 {
		return
	}
	b.WriteString("## Private models\n\n")
	b.WriteString("File-private helper schemas defined in this model YAML file.\n")
	b.WriteString("Promote a helper model to a public model file when it needs to be reused from other YAML files.\n\n")
	b.WriteString("| model | kind | shape | note |\n")
	b.WriteString("|---|---|---|---|\n")
	for _, model := range models {
		fmt.Fprintf(b, "| %s | %s | %s | %s |\n",
			tableCell(model.ID),
			tableCell(model.Kind),
			tableCell(compactShape(model)),
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

func compactShape(model *semantic.Model) string {
	switch model.Kind {
	case "struct":
		if len(model.Fields) == 0 {
			return ""
		}
		parts := make([]string, 0, len(model.Fields))
		for _, field := range model.Fields {
			parts = append(parts, field.Name+": "+field.Type)
		}
		return strings.Join(parts, "<br/>")
	case "enum":
		return strings.Join(model.Values, "<br/>")
	case "list":
		if model.Element == "" {
			return ""
		}
		return "element: " + model.Element
	case "dict":
		if model.Value == "" {
			return ""
		}
		return "value: " + model.Value
	case "tagged_union":
		if model.Discriminator == "" && len(model.Variants) == 0 {
			return ""
		}
		parts := make([]string, 0, 1+len(model.Variants))
		parts = append(parts, "discriminator: "+model.Discriminator)
		for _, v := range model.Variants {
			parts = append(parts, "tag: "+v.Tag)
		}
		return strings.Join(parts, "<br/>")
	default:
		return ""
	}
}

func tableCell(value string) string {
	value = strings.TrimSpace(value)
	if value == "" || value == "-" {
		return "-"
	}
	value = strings.ReplaceAll(value, "\r\n", "\n")
	value = strings.ReplaceAll(value, "\n", "<br/>")
	value = strings.ReplaceAll(value, "|", `\|`)
	return value
}
