package resolve

import (
	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func Build(raw *rawyaml.Project) (*semantic.Project, []semantic.Diagnostic) {
	project := semantic.NewProject()
	symbols := newSymbolTable(project)

	for _, file := range raw.Files {
		fileID := semantic.FileID(file.ID)
		project.SourceFilesByID[fileID] = semantic.SourceFile{
			ID:      fileID,
			Kind:    string(file.Kind),
			ViewAs:  file.ViewAs,
			Content: file.Content,
		}
		if file.RenderIndex != nil {
			for _, group := range file.RenderIndex.Groups {
				project.RenderGroups = append(project.RenderGroups, semantic.RenderGroup{
					ID:      group.ID,
					Label:   group.Label,
					Modules: group.Modules,
				})
			}
		}
	}

	for _, file := range raw.Files {
		if file.Kind != rawyaml.FileKindNode || file.NodeFile == nil {
			continue
		}
		fileID := semantic.FileID(file.ID)
		module := moduleForFileID(fileID)

		for _, model := range file.NodeFile.Models {
			symbols.addNode(buildModel(fileID, module, model))
		}
		for _, store := range file.NodeFile.Stores {
			symbols.addNode(buildStore(fileID, module, store))
		}
		for _, actor := range file.NodeFile.Actors {
			symbols.addNode(buildActor(fileID, actor))
		}
		for _, state := range file.NodeFile.States {
			symbols.addNode(buildState(fileID, module, state))
		}
		for _, event := range file.NodeFile.Events {
			symbols.addNode(buildEvent(fileID, module, event))
		}
		for _, task := range file.NodeFile.Tasks {
			symbols.addNode(buildTask(fileID, module, task))
		}
		for _, branch := range file.NodeFile.Branches {
			symbols.addNode(buildBranch(fileID, module, branch))
		}
		for _, fork := range file.NodeFile.Forks {
			symbols.addNode(buildFork(fileID, module, fork))
		}
		for _, join := range file.NodeFile.Joins {
			symbols.addNode(buildJoin(fileID, module, join))
		}
	}

	buildInitializedStores(project, symbols)
	resolveTaskStoreAccess(project, symbols)
	buildFlows(raw, project, symbols)
	buildTransitions(raw, project, symbols)
	buildScenarios(raw, project, symbols)
	buildERViews(raw, project, symbols)
	buildAPIViews(raw, project, symbols)
	validateProject(project, symbols)
	buildReferences(project)
	return project, sortedDiagnostics(symbols.diags)
}

func buildTask(fileID semantic.FileID, module string, raw rawyaml.Task) *semantic.Task {
	qid := qidFor(module, "task", raw.ID)
	task := &semantic.Task{
		BaseNode: semantic.BaseNode{
			QID:    qid,
			FileID: fileID,
			ID:     raw.ID,
			Kind:   semantic.NodeKindTask,
			Main:   raw.Main,
			Note:   raw.Note,
		},
		Endpoint: raw.Endpoint,
		Method:   raw.Method,
		Path:     raw.Path,
	}

	for _, param := range raw.Params {
		task.Params = append(task.Params, semantic.Param{
			Name:      param.Name,
			Model:     resolveModelQID(module, param.Model),
			ModelName: param.Model,
			Note:      param.Note,
		})
	}
	if raw.Returns != nil {
		modelQID := resolveModelQID(module, raw.Returns.Model)
		task.Returns = &semantic.Return{
			Name:      raw.Returns.Name,
			Model:     modelQID,
			ModelName: raw.Returns.Model,
			Asset: &semantic.Asset{
				Name:       raw.Returns.Name,
				Model:      modelQID,
				ModelName:  raw.Returns.Model,
				ProducedBy: qid,
				FileID:     fileID,
			},
		}
	}
	for _, init := range raw.Initializes {
		task.Initializes = append(task.Initializes, semantic.InitializedStore{
			Name:  init.Name,
			Model: resolveModelQID(module, init.Model),
			Note:  init.Note,
		})
	}
	for _, read := range raw.Reads {
		task.Reads = append(task.Reads, semantic.StoreRef{Name: read})
	}
	for _, write := range raw.Writes {
		task.Writes = append(task.Writes, semantic.StoreRef{Name: write})
	}
	return task
}

func buildBranch(fileID semantic.FileID, module string, raw rawyaml.ControlNode) *semantic.Branch {
	branch := &semantic.Branch{
		BaseNode: semantic.BaseNode{
			QID:    qidFor(module, "branch", raw.ID),
			FileID: fileID,
			ID:     raw.ID,
			Kind:   semantic.NodeKindBranch,
			Main:   raw.Main,
			Note:   raw.Note,
		},
	}
	branch.Params = buildParams(module, raw.Params)
	return branch
}

func buildFork(fileID semantic.FileID, module string, raw rawyaml.ControlNode) *semantic.Fork {
	fork := &semantic.Fork{
		BaseNode: semantic.BaseNode{
			QID:    qidFor(module, "fork", raw.ID),
			FileID: fileID,
			ID:     raw.ID,
			Kind:   semantic.NodeKindFork,
			Main:   raw.Main,
			Note:   raw.Note,
		},
	}
	fork.Params = buildParams(module, raw.Params)
	return fork
}

func buildJoin(fileID semantic.FileID, module string, raw rawyaml.ControlNode) *semantic.Join {
	qid := qidFor(module, "join", raw.ID)
	join := &semantic.Join{
		BaseNode: semantic.BaseNode{
			QID:    qid,
			FileID: fileID,
			ID:     raw.ID,
			Kind:   semantic.NodeKindJoin,
			Main:   raw.Main,
			Note:   raw.Note,
		},
	}
	join.Params = buildParams(module, raw.Params)
	if raw.Returns != nil {
		modelQID := resolveModelQID(module, raw.Returns.Model)
		join.Returns = &semantic.Return{
			Name:      raw.Returns.Name,
			Model:     modelQID,
			ModelName: raw.Returns.Model,
			Asset: &semantic.Asset{
				Name:       raw.Returns.Name,
				Model:      modelQID,
				ModelName:  raw.Returns.Model,
				ProducedBy: qid,
				FileID:     fileID,
			},
		}
	}
	return join
}

func buildParams(module string, raw []rawyaml.Param) []semantic.Param {
	params := make([]semantic.Param, 0, len(raw))
	for _, param := range raw {
		params = append(params, semantic.Param{
			Name:      param.Name,
			Model:     resolveModelQID(module, param.Model),
			ModelName: param.Model,
			Note:      param.Note,
		})
	}
	return params
}

func buildModel(fileID semantic.FileID, module string, raw rawyaml.Model) *semantic.Model {
	model := &semantic.Model{
		BaseNode: semantic.BaseNode{
			QID:    qidFor(module, "model", raw.ID),
			FileID: fileID,
			ID:     raw.ID,
			Kind:   semantic.NodeKindModel,
			Note:   raw.Note,
		},
		Kind:    raw.Kind,
		Element: raw.Element,
		Value:   raw.Value,
	}
	for _, field := range raw.Fields {
		model.Fields = append(model.Fields, semantic.ModelField{
			Name:   field.Name,
			Type:   field.Type,
			PK:     field.PK,
			FK:     field.FK,
			Unique: field.Unique,
			Note:   field.Note,
		})
	}
	return model
}

func buildStore(fileID semantic.FileID, module string, raw rawyaml.Store) *semantic.Store {
	return &semantic.Store{
		BaseNode: semantic.BaseNode{
			QID:    qidFor(module, "store", raw.ID),
			FileID: fileID,
			ID:     raw.ID,
			Kind:   semantic.NodeKindStore,
			Note:   raw.Note,
		},
		StoreKind: raw.Kind,
		Of:        resolveModelQID(module, raw.Of),
		OfName:    raw.Of,
	}
}

func buildActor(fileID semantic.FileID, raw rawyaml.Actor) *semantic.Actor {
	return &semantic.Actor{
		BaseNode: semantic.BaseNode{
			QID:    actorQID(raw.ID),
			FileID: fileID,
			ID:     raw.ID,
			Kind:   semantic.NodeKindActor,
			Note:   raw.Note,
		},
	}
}

func buildState(fileID semantic.FileID, module string, raw rawyaml.State) *semantic.State {
	return &semantic.State{
		BaseNode: semantic.BaseNode{
			QID:    qidFor(module, "state", raw.ID),
			FileID: fileID,
			ID:     raw.ID,
			Kind:   semantic.NodeKindState,
			Note:   raw.Note,
		},
		Initial:   raw.Initial,
		Final:     raw.Final,
		Wireframe: buildWireframe(raw.Wireframe),
	}
}

func buildWireframe(raw *rawyaml.WireframeElement) *semantic.WireframeElement {
	if raw == nil {
		return nil
	}
	out := &semantic.WireframeElement{
		Type:        raw.Type,
		ID:          raw.ID,
		Label:       raw.Label,
		Cols:        raw.Cols,
		Fires:       raw.Fires,
		Disabled:    raw.Disabled,
		Placeholder: raw.Placeholder,
		Span:        raw.Span,
		Layout:      buildWireframeLayout(raw.Layout),
	}
	for _, child := range raw.Children {
		built := buildWireframe(&child)
		if built != nil {
			out.Children = append(out.Children, *built)
		}
	}
	return out
}

func buildWireframeLayout(raw *rawyaml.WireframeLayout) *semantic.WireframeLayout {
	if raw == nil {
		return nil
	}
	return &semantic.WireframeLayout{
		Width:     raw.Width,
		Height:    raw.Height,
		MinWidth:  raw.MinWidth,
		MinHeight: raw.MinHeight,
		Grow:      raw.Grow,
		Gap:       raw.Gap,
		Padding:   raw.Padding,
		Align:     raw.Align,
		Justify:   raw.Justify,
		Scroll:    raw.Scroll,
	}
}

func buildEvent(fileID semantic.FileID, module string, raw rawyaml.Event) *semantic.Event {
	event := &semantic.Event{
		BaseNode: semantic.BaseNode{
			QID:    qidFor(module, "event", raw.ID),
			FileID: fileID,
			ID:     raw.ID,
			Kind:   semantic.NodeKindEvent,
			Note:   raw.Note,
		},
		Source:      raw.Source,
		Actor:       raw.Actor,
		WatchesName: raw.Watches,
	}
	if raw.Payload != nil {
		event.PayloadName = raw.Payload.Model
		event.PayloadModel = resolveModelQID(module, raw.Payload.Model)
	}
	if raw.Watches != "" {
		event.Watches = resolveSameModuleStoreQID(module, raw.Watches)
	}
	return event
}

func buildInitializedStores(project *semantic.Project, symbols *symbolTable) {
	for _, task := range project.TasksByQID {
		for i := range task.Initializes {
			init := &task.Initializes[i]
			store := &semantic.Store{
				BaseNode: semantic.BaseNode{
					QID:    localStoreQID(task.QID, init.Name),
					FileID: task.FileID,
					ID:     init.Name,
					Kind:   semantic.NodeKindStore,
					Note:   init.Note,
				},
				StoreKind:   "collection",
				Of:          init.Model,
				OfName:      shortName(init.Model),
				FilePrivate: true,
				LocalName:   init.Name,
			}
			init.Store = store
			if project.StoresByFileLocal[task.FileID] == nil {
				project.StoresByFileLocal[task.FileID] = map[string]*semantic.Store{}
			}
			if _, exists := project.StoresByFileLocal[task.FileID][init.Name]; exists {
				symbols.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateInitializedStore, task.FileID, "duplicate initialized store: "+init.Name)
				continue
			}
			project.StoresByFileLocal[task.FileID][init.Name] = store
		}
	}
}

func resolveTaskStoreAccess(project *semantic.Project, symbols *symbolTable) {
	for _, task := range project.TasksByQID {
		module := moduleForFileID(task.FileID)
		for i := range task.Reads {
			ref := resolveStoreRef(project, task.FileID, module, task.Reads[i].Name)
			task.Reads[i] = ref
			if ref.Store == "" {
				continue
			}
			project.TasksReadingStore[ref.Store] = append(project.TasksReadingStore[ref.Store], task.QID)
		}
		for i := range task.Writes {
			ref := resolveStoreRef(project, task.FileID, module, task.Writes[i].Name)
			task.Writes[i] = ref
			if ref.Store == "" {
				continue
			}
			project.TasksWritingStore[ref.Store] = append(project.TasksWritingStore[ref.Store], task.QID)
		}
	}
}

func resolveStoreRef(project *semantic.Project, fileID semantic.FileID, module string, name string) semantic.StoreRef {
	if localByName := project.StoresByFileLocal[fileID]; localByName != nil {
		if store := localByName[name]; store != nil {
			return semantic.StoreRef{Name: name, Store: store.QID, FilePrivate: true}
		}
	}

	qid := resolveSameModuleStoreQID(module, name)
	if _, ok := project.StoresByQID[qid]; ok {
		return semantic.StoreRef{Name: name, Store: qid}
	}
	if isFullQID(name, "store") {
		qid = semantic.QualifiedID(name)
		if _, ok := project.StoresByQID[qid]; ok {
			return semantic.StoreRef{Name: name, Store: qid}
		}
	}
	return semantic.StoreRef{Name: name}
}
