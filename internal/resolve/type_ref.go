package resolve

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

type typeRefParseError struct {
	Raw string
	Pos int
	Msg string
}

func (e *typeRefParseError) Error() string {
	return fmt.Sprintf("invalid TypeRef %q at position %d: %s", e.Raw, e.Pos, e.Msg)
}

func parseTypeRef(raw string, module string) (*semantic.TypeRef, error) {
	parser := typeRefParser{raw: raw, module: module}
	ref, err := parser.parse()
	if err != nil {
		return nil, err
	}
	return ref, nil
}

const maxTypeRefContainerNestingDepth = 16

type typeRefParser struct {
	raw    string
	module string
	pos    int
}

func (p *typeRefParser) parse() (*semantic.TypeRef, error) {
	p.skipSpaces()
	if p.eof() {
		return nil, p.err("empty TypeRef")
	}
	ref, err := p.parseRef(0)
	if err != nil {
		return nil, err
	}
	p.skipSpaces()
	if !p.eof() {
		return nil, p.err("unexpected trailing input")
	}
	return ref, nil
}

func (p *typeRefParser) parseRef(containerDepth int) (*semantic.TypeRef, error) {
	p.skipSpaces()
	start := p.pos
	ident := p.readIdent()
	if ident == "" {
		return nil, p.err("expected type name")
	}
	if !isValidTypeRefIdent(ident) {
		return nil, p.errAt(start, "invalid type name")
	}
	p.skipSpaces()
	if p.consume('<') {
		if ident != "list" && ident != "dict" {
			return nil, p.errAt(start, "unsupported container kind: "+ident)
		}
		if containerDepth+1 > maxTypeRefContainerNestingDepth {
			return nil, p.errAt(start, fmt.Sprintf("container nesting depth exceeds limit %d", maxTypeRefContainerNestingDepth))
		}
		typeArg, err := p.parseRef(containerDepth + 1)
		if err != nil {
			return nil, err
		}
		p.skipSpaces()
		if !p.consume('>') {
			return nil, p.err("expected '>'")
		}
		if ident == "list" {
			return &semantic.TypeRef{Kind: semantic.TypeRefList, Raw: p.raw[start:p.pos], Name: "list", Elem: typeArg}, nil
		}
		return &semantic.TypeRef{Kind: semantic.TypeRefDict, Raw: p.raw[start:p.pos], Name: "dict", Value: typeArg}, nil
	}
	if ident == "list" || ident == "dict" {
		return nil, p.errAt(start, ident+" requires a type argument")
	}
	if isPrimitive(ident) {
		return &semantic.TypeRef{Kind: semantic.TypeRefPrimitive, Raw: ident, Name: ident}, nil
	}
	return &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Raw: ident, Name: ident, Model: resolveTypeRefModelQID(p.module, ident)}, nil
}

func resolveTypeRefModelQID(module, ref string) semantic.QualifiedID {
	if ref == "" {
		return ""
	}
	if isFullQID(ref, "model") {
		return semantic.QualifiedID(ref)
	}
	if strings.Contains(ref, ".") {
		parts := strings.Split(ref, ".")
		return qidFor(strings.Join(parts[:len(parts)-1], "."), "model", parts[len(parts)-1])
	}
	return resolveModelQID(module, ref)
}

func (p *typeRefParser) readIdent() string {
	start := p.pos
	for !p.eof() {
		r, size := utf8.DecodeRuneInString(p.raw[p.pos:])
		if r == utf8.RuneError && size == 0 {
			break
		}
		if !isTypeRefIdentRune(r) {
			break
		}
		p.pos += size
	}
	return strings.TrimSpace(p.raw[start:p.pos])
}

func isTypeRefIdentRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '.'
}

func isValidTypeRefIdent(ident string) bool {
	if ident == "" || strings.HasPrefix(ident, ".") || strings.HasSuffix(ident, ".") || strings.Contains(ident, "..") {
		return false
	}
	return true
}

func (p *typeRefParser) skipSpaces() {
	for !p.eof() {
		r, size := utf8.DecodeRuneInString(p.raw[p.pos:])
		if !unicode.IsSpace(r) {
			return
		}
		p.pos += size
	}
}

func (p *typeRefParser) consume(want byte) bool {
	if p.eof() || p.raw[p.pos] != want {
		return false
	}
	p.pos++
	return true
}

func (p *typeRefParser) eof() bool {
	return p.pos >= len(p.raw)
}

func (p *typeRefParser) err(msg string) error {
	return p.errAt(p.pos, msg)
}

func (p *typeRefParser) errAt(pos int, msg string) error {
	return &typeRefParseError{Raw: p.raw, Pos: pos, Msg: msg}
}

func typeRefNamedModels(ref *semantic.TypeRef) []semantic.TypeRef {
	if ref == nil {
		return nil
	}
	switch ref.Kind {
	case semantic.TypeRefNamedModel:
		return []semantic.TypeRef{*ref}
	case semantic.TypeRefList:
		return typeRefNamedModels(ref.Elem)
	case semantic.TypeRefDict:
		return typeRefNamedModels(ref.Value)
	default:
		return nil
	}
}

func typeRefModelsExist(project *semantic.Project, ref *semantic.TypeRef) bool {
	for _, named := range typeRefNamedModels(ref) {
		if !modelExists(project, named.Model) {
			return false
		}
	}
	return true
}

func normalizeContainerTypeRef(project *semantic.Project, ref *semantic.TypeRef) *semantic.TypeRef {
	if project == nil || ref == nil || ref.Kind != semantic.TypeRefNamedModel {
		return ref
	}
	model := modelByQID(project, ref.Model)
	if model == nil {
		return ref
	}
	switch model.Kind {
	case "list":
		if model.ElementRef == nil {
			return ref
		}
		return &semantic.TypeRef{
			Kind: semantic.TypeRefList,
			Raw:  fmt.Sprintf("list<%s>", model.ElementRef.String()),
			Name: "list",
			Elem: model.ElementRef,
		}
	case "dict":
		if model.ValueRef == nil {
			return ref
		}
		return &semantic.TypeRef{
			Kind:  semantic.TypeRefDict,
			Raw:   fmt.Sprintf("dict<%s>", model.ValueRef.String()),
			Name:  "dict",
			Value: model.ValueRef,
		}
	default:
		return ref
	}
}

func resolveScopedTypeRefs(project *semantic.Project) {
	if project == nil {
		return
	}
	seenTasks := map[*semantic.Task]struct{}{}
	for _, task := range project.TasksByQID {
		if _, ok := seenTasks[task]; ok {
			continue
		}
		seenTasks[task] = struct{}{}
		resolveScopedParams(project, task.FileID, task.Params)
		if task.Returns != nil {
			resolveScopedReturn(project, task.FileID, task.Returns)
		}
	}
	seenBranches := map[*semantic.Branch]struct{}{}
	for _, branch := range project.BranchesByQID {
		if _, ok := seenBranches[branch]; ok {
			continue
		}
		seenBranches[branch] = struct{}{}
		resolveScopedParams(project, branch.FileID, branch.Params)
	}
	seenForks := map[*semantic.Fork]struct{}{}
	for _, fork := range project.ForksByQID {
		if _, ok := seenForks[fork]; ok {
			continue
		}
		seenForks[fork] = struct{}{}
		resolveScopedParams(project, fork.FileID, fork.Params)
	}
	seenJoins := map[*semantic.Join]struct{}{}
	for _, join := range project.JoinsByQID {
		if _, ok := seenJoins[join]; ok {
			continue
		}
		seenJoins[join] = struct{}{}
		resolveScopedParams(project, join.FileID, join.Params)
		if join.Returns != nil {
			resolveScopedReturn(project, join.FileID, join.Returns)
		}
	}
	for _, model := range allModels(project) {
		resolveScopedTypeRef(project, model.FileID, model.ElementRef)
		resolveScopedTypeRef(project, model.FileID, model.ValueRef)
		for i := range model.Fields {
			resolveScopedTypeRef(project, model.FileID, model.Fields[i].TypeRef)
		}
		for i := range model.Variants {
			for j := range model.Variants[i].Fields {
				resolveScopedTypeRef(project, model.FileID, model.Variants[i].Fields[j].TypeRef)
			}
		}
	}
}

func resolveScopedParams(project *semantic.Project, fileID semantic.FileID, params []semantic.Param) {
	for i := range params {
		resolveScopedTypeRef(project, fileID, params[i].TypeRef)
		if params[i].TypeRef != nil && params[i].TypeRef.Kind == semantic.TypeRefNamedModel {
			params[i].Model = params[i].TypeRef.Model
		}
	}
}

func resolveScopedReturn(project *semantic.Project, fileID semantic.FileID, ret *semantic.Return) {
	resolveScopedTypeRef(project, fileID, ret.TypeRef)
	if ret.TypeRef != nil && ret.TypeRef.Kind == semantic.TypeRefNamedModel {
		ret.Model = ret.TypeRef.Model
		if ret.Asset != nil {
			ret.Asset.Model = ret.TypeRef.Model
		}
	}
}

func resolveScopedTypeRef(project *semantic.Project, fileID semantic.FileID, ref *semantic.TypeRef) {
	if ref == nil {
		return
	}
	switch ref.Kind {
	case semantic.TypeRefNamedModel:
		if isBareTypeRefName(ref.Name) {
			if model := privateModelByLocalID(project, fileID, ref.Name); model != nil {
				ref.Model = model.QID
				return
			}
		}
		ref.Model = resolveTypeRefModelQID(moduleForFileID(fileID), ref.Name)
	case semantic.TypeRefList:
		resolveScopedTypeRef(project, fileID, ref.Elem)
	case semantic.TypeRefDict:
		resolveScopedTypeRef(project, fileID, ref.Value)
	}
}

func isBareTypeRefName(name string) bool {
	return name != "" && !strings.Contains(name, ".")
}

func modelByQID(project *semantic.Project, qid semantic.QualifiedID) *semantic.Model {
	if project == nil || qid == "" {
		return nil
	}
	if model := project.ModelsByQID[qid]; model != nil {
		return model
	}
	return privateModelByQID(project, qid)
}

func privateModelByQID(project *semantic.Project, qid semantic.QualifiedID) *semantic.Model {
	if project == nil || qid == "" {
		return nil
	}
	for _, byName := range project.PrivateModelsByFile {
		for _, model := range byName {
			if model.QID == qid {
				return model
			}
		}
	}
	return nil
}

func privateModelByLocalID(project *semantic.Project, fileID semantic.FileID, id string) *semantic.Model {
	if project == nil || id == "" {
		return nil
	}
	return project.PrivateModelsByFile[fileID][id]
}

func allModels(project *semantic.Project) []*semantic.Model {
	if project == nil {
		return nil
	}
	models := make([]*semantic.Model, 0, len(project.ModelsByQID))
	for _, model := range project.ModelsByQID {
		models = append(models, model)
	}
	for _, byName := range project.PrivateModelsByFile {
		for _, model := range byName {
			models = append(models, model)
		}
	}
	return models
}

func typeRefsCompatible(project *semantic.Project, source *semantic.TypeRef, target *semantic.TypeRef) bool {
	if source == nil || target == nil {
		return false
	}
	source = normalizeContainerTypeRef(project, source)
	target = normalizeContainerTypeRef(project, target)
	if typeRefIsAny(source) || typeRefIsAny(target) {
		return true
	}
	if source.Kind != target.Kind {
		return false
	}
	switch source.Kind {
	case semantic.TypeRefPrimitive:
		return source.Name == target.Name
	case semantic.TypeRefNamedModel:
		return source.Model != "" && source.Model == target.Model
	case semantic.TypeRefList:
		return typeRefsCompatible(project, source.Elem, target.Elem)
	case semantic.TypeRefDict:
		return typeRefsCompatible(project, source.Value, target.Value)
	default:
		return false
	}
}

func typeRefIsAny(ref *semantic.TypeRef) bool {
	return ref != nil && ref.Kind == semantic.TypeRefPrimitive && ref.Name == "any"
}

func typeRefHasOpaqueAnyContainer(ref *semantic.TypeRef) bool {
	if ref == nil {
		return false
	}
	switch ref.Kind {
	case semantic.TypeRefList:
		return typeRefContainsAny(ref.Elem)
	case semantic.TypeRefDict:
		return typeRefContainsAny(ref.Value)
	default:
		return false
	}
}

func typeRefContainsAny(ref *semantic.TypeRef) bool {
	if ref == nil {
		return false
	}
	if typeRefIsAny(ref) {
		return true
	}
	switch ref.Kind {
	case semantic.TypeRefList:
		return typeRefContainsAny(ref.Elem)
	case semantic.TypeRefDict:
		return typeRefContainsAny(ref.Value)
	default:
		return false
	}
}
