package er

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

type entity struct {
	Model *semantic.Model
	Store *semantic.Store
}

type relation struct {
	From   string
	To     string
	Unique bool
}

func RenderView(project *semantic.Project, viewID string) (string, error) {
	if project == nil {
		return "", fmt.Errorf("project is nil")
	}
	view := project.ERViewsByID[viewID]
	if view == nil {
		return "", fmt.Errorf("ER view not found: %s", viewID)
	}
	return render(project, view), nil
}

func render(project *semantic.Project, view *semantic.ERView) string {
	entities := collectEntities(project, view)
	includedModels := map[semantic.QualifiedID]struct{}{}
	for _, entity := range entities {
		includedModels[entity.Model.QID] = struct{}{}
	}
	relations := collectRelations(entities, includedModels)

	var b strings.Builder
	b.WriteString("# " + view.ID + "\n\n")
	if strings.TrimSpace(view.Note) != "" {
		b.WriteString(strings.TrimSpace(view.Note) + "\n\n")
	}
	b.WriteString("```mermaid\n")
	b.WriteString("erDiagram\n")
	for i, entity := range entities {
		if i > 0 {
			b.WriteString("\n")
		}
		writeEntity(&b, entity.Model)
	}
	if len(relations) > 0 {
		b.WriteString("\n")
		for _, relation := range relations {
			cardinality := "}o--||"
			if relation.Unique {
				cardinality = "|o--||"
			}
			b.WriteString("  " + relation.From + " " + cardinality + " " + relation.To + " : \"\"\n")
		}
	}
	b.WriteString("```\n")
	return b.String()
}

func collectEntities(project *semantic.Project, view *semantic.ERView) []entity {
	seenModels := map[semantic.QualifiedID]struct{}{}
	var entities []entity
	for _, viewModule := range view.Modules {
		stores := dbStoresForModule(project, viewModule.Module)
		for _, store := range stores {
			model := project.ModelsByQID[store.Of]
			if model == nil || model.Kind != "struct" {
				continue
			}
			if _, exists := seenModels[model.QID]; exists {
				continue
			}
			seenModels[model.QID] = struct{}{}
			entities = append(entities, entity{Model: model, Store: store})
		}
	}
	return entities
}

func dbStoresForModule(project *semantic.Project, module string) []*semantic.Store {
	var stores []*semantic.Store
	for _, store := range project.StoresByQID {
		if store.StoreKind != "db" || moduleForFile(store.FileID) != module {
			continue
		}
		stores = append(stores, store)
	}
	sort.Slice(stores, func(i, j int) bool {
		if stores[i].Of != stores[j].Of {
			return stores[i].Of < stores[j].Of
		}
		return stores[i].QID < stores[j].QID
	})
	return stores
}

func writeEntity(b *strings.Builder, model *semantic.Model) {
	b.WriteString("  " + model.ID + " {\n")
	for _, field := range model.Fields {
		line := "    " + erType(field) + " " + field.Name
		flags := fieldFlags(field)
		if flags != "" {
			line += " " + flags
		}
		b.WriteString(line + "\n")
	}
	b.WriteString("  }\n")
}

func erType(field semantic.ModelField) string {
	if field.FK != "" {
		return "string"
	}
	switch field.Type {
	case "str":
		return "string"
	case "int", "float", "bytes", "datetime", "any":
		return field.Type
	case "bool":
		return "boolean"
	default:
		return "json"
	}
}

func fieldFlags(field semantic.ModelField) string {
	var flags []string
	if field.PK {
		flags = append(flags, "PK")
	}
	if field.FK != "" {
		flags = append(flags, "FK")
	}
	return strings.Join(flags, ", ")
}

func collectRelations(entities []entity, includedModels map[semantic.QualifiedID]struct{}) []relation {
	seen := map[string]struct{}{}
	var relations []relation
	for _, entity := range entities {
		for _, field := range entity.Model.Fields {
			if field.FK == "" {
				continue
			}
			targetModel := modelQIDFromFK(field.FK)
			if _, ok := includedModels[targetModel]; !ok {
				continue
			}
			to := shortModelName(targetModel)
			key := entity.Model.ID + "\x00" + to + "\x00" + fmt.Sprint(field.Unique)
			if _, exists := seen[key]; exists {
				continue
			}
			seen[key] = struct{}{}
			relations = append(relations, relation{From: entity.Model.ID, To: to, Unique: field.Unique})
		}
	}
	return relations
}

func modelQIDFromFK(fk string) semantic.QualifiedID {
	parts := strings.Split(fk, ".")
	for i, part := range parts {
		if part == "model" && i+1 < len(parts) {
			return semantic.QualifiedID(strings.Join(parts[:i+2], "."))
		}
	}
	if len(parts) >= 2 {
		return semantic.QualifiedID(strings.Join(parts[:len(parts)-1], "."))
	}
	return ""
}

func shortModelName(qid semantic.QualifiedID) string {
	parts := strings.Split(qid.String(), ".")
	if len(parts) == 0 {
		return qid.String()
	}
	return parts[len(parts)-1]
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
