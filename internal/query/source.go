package query

import (
	"fmt"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

const sourceRangeUnavailable = "source_range_unavailable"

type sourceBlock struct {
	startLine int
	endLine   int
	column    int
	text      string
}

func (s *Service) GetSource(req GetSourceRequest) (GetSourceResponse, error) {
	if req.Fallback != "" && req.Fallback != "file" && req.Fallback != "error" {
		return GetSourceResponse{}, fmt.Errorf("unsupported fallback: %s", req.Fallback)
	}

	object, fileID, block, err := s.sourceForSelector(req.Selector)
	if err != nil {
		return GetSourceResponse{}, err
	}
	if block.text != "" {
		return sourceResponse(object, fileID, block, "", nil), nil
	}
	if req.Fallback == "error" {
		return GetSourceResponse{}, fmt.Errorf("source range not found: %s", object.ID)
	}
	fileBlock, err := s.fullFileSource(fileID)
	if err != nil {
		return GetSourceResponse{}, err
	}
	diags := []semantic.Diagnostic{{
		Severity: semantic.SeverityWarning,
		Code:     sourceRangeUnavailable,
		FileID:   fileID,
		Message:  "source range is unavailable; returned whole file",
	}}
	return sourceResponse(object, fileID, fileBlock, "file", diags), nil
}

func (s *Service) sourceForSelector(selector Selector) (ObjectRef, semantic.FileID, sourceBlock, error) {
	if s.isFileSelector(selector) {
		fileID, err := s.fileBySelector(selector)
		if err != nil {
			return ObjectRef{}, "", sourceBlock{}, err
		}
		kind := s.fileKind(fileID)
		block, err := s.fullFileSource(fileID)
		return fileObjectRef(fileID, kind), fileID, block, err
	}
	if s.isAPIViewSelector(selector) {
		view, err := s.apiViewBySelector(selector)
		if err != nil {
			return ObjectRef{}, "", sourceBlock{}, err
		}
		return apiViewObjectRef(view), view.FileID, s.findViewSource(view.FileID, view.ID), nil
	}
	if s.isERViewSelector(selector) {
		view, err := s.erViewBySelector(selector)
		if err != nil {
			return ObjectRef{}, "", sourceBlock{}, err
		}
		return erViewObjectRef(view), view.FileID, s.findViewSource(view.FileID, view.ID), nil
	}
	if s.isScenarioSelector(selector) {
		scenario, err := s.scenarioBySelector(selector)
		if err != nil {
			return ObjectRef{}, "", sourceBlock{}, err
		}
		return scenarioObjectRef(scenario), scenario.FileID, s.findViewSource(scenario.FileID, scenario.ID), nil
	}
	if s.isTransitionSelector(selector) {
		transition, err := s.transitionBySelector(selector)
		if err != nil {
			return ObjectRef{}, "", sourceBlock{}, err
		}
		return transitionObjectRef(transition), transition.FileID, s.findTransitionSource(transition), nil
	}
	if s.isFieldSelector(selector) {
		model, field, err := s.modelFieldBySelector(selector)
		if err != nil {
			return ObjectRef{}, "", sourceBlock{}, err
		}
		return fieldObjectRef(model, field), model.FileID, s.findFieldSource(model, field), nil
	}
	if s.isAssetSelector(selector) {
		asset, err := s.assetBySelector(selector)
		if err != nil {
			return ObjectRef{}, "", sourceBlock{}, err
		}
		return assetObjectRef(asset), asset.FileID, s.findAssetSource(asset), nil
	}
	if s.isPrivateSubNodeSelector(selector) {
		node, err := s.privateSubNodeBySelector(selector)
		if err != nil {
			return ObjectRef{}, "", sourceBlock{}, err
		}
		return objectRef(node), node.GetFileID(), s.findNodeSource(node), nil
	}

	node, err := s.nodeByID(selector.ID)
	if err != nil {
		return ObjectRef{}, "", sourceBlock{}, err
	}
	return objectRef(node), node.GetFileID(), s.findNodeSource(node), nil
}

func (s *Service) fullFileSource(fileID semantic.FileID) (sourceBlock, error) {
	content, ok := s.fileContent(fileID)
	if !ok {
		return sourceBlock{}, fmt.Errorf("source file not found: %s", fileID)
	}
	lines := splitSourceLines(content)
	return makeBlock(lines, 0, len(lines)), nil
}

func (s *Service) findNodeSource(node semantic.Node) sourceBlock {
	content, ok := s.fileContent(node.GetFileID())
	if !ok {
		return sourceBlock{}
	}
	lines := splitSourceLines(content)
	return findTopLevelSequenceItem(lines, "nodes", map[string]string{
		"id":   node.GetID(),
		"type": string(node.GetKind()),
	})
}

func (s *Service) findFieldSource(model *semantic.Model, field semantic.ModelField) sourceBlock {
	nodeBlock := s.findNodeSource(model)
	if nodeBlock.text == "" {
		return sourceBlock{}
	}
	lines := splitSourceLines(s.project.SourceFilesByID[model.FileID].Content)
	fieldBlock := findNestedSequenceItem(lines, nodeBlock.startLine-1, nodeBlock.endLine, "fields", map[string]string{"name": field.Name})
	if fieldBlock.text != "" {
		return fieldBlock
	}
	return nodeBlock
}

func (s *Service) findAssetSource(asset *semantic.Asset) sourceBlock {
	if asset == nil {
		return sourceBlock{}
	}
	node := s.project.NodesByID[asset.ProducedBy]
	if node == nil {
		for _, candidate := range s.project.TasksByQID {
			if candidate.Returns != nil && candidate.Returns.Asset == asset {
				node = candidate
				break
			}
		}
	}
	if node == nil {
		for _, candidate := range s.project.JoinsByQID {
			if candidate.Returns != nil && candidate.Returns.Asset == asset {
				node = candidate
				break
			}
		}
	}
	if node == nil {
		return sourceBlock{}
	}
	nodeBlock := s.findNodeSource(node)
	if nodeBlock.text == "" {
		return sourceBlock{}
	}
	lines := splitSourceLines(s.project.SourceFilesByID[asset.FileID].Content)
	returns := findNestedMappingBlock(lines, nodeBlock.startLine-1, nodeBlock.endLine, "returns")
	if returns.text != "" && blockHasKeyValue(lines, returns.startLine-1, returns.endLine, "name", asset.Name) {
		return returns
	}
	return nodeBlock
}

func (s *Service) findTransitionSource(transition semantic.Transition) sourceBlock {
	content, ok := s.fileContent(transition.FileID)
	if !ok {
		return sourceBlock{}
	}
	lines := splitSourceLines(content)
	start, end := topLevelSectionRange(lines, "transitions")
	if start < 0 {
		return sourceBlock{}
	}
	match := map[string]string{"from": transition.From, "on": transition.On, "to": transition.To}
	if transition.Guard != "" {
		match["guard"] = transition.Guard
	}
	for i := start; i < end; i++ {
		trimmed := strings.TrimSpace(lines[i])
		if !strings.HasPrefix(trimmed, "- ") && trimmed != "-" {
			continue
		}
		itemEnd := nextSequenceItemOrSectionEnd(lines, i+1, end, indentOf(lines[i]))
		if !blockMatches(lines, i, itemEnd, match) {
			i = itemEnd - 1
			continue
		}
		if transition.Guard == "" && blockHasKey(lines, i, itemEnd, "guard") {
			i = itemEnd - 1
			continue
		}
		return makeBlock(lines, i, itemEnd)
	}
	return sourceBlock{}
}

func (s *Service) findViewSource(fileID semantic.FileID, id string) sourceBlock {
	content, ok := s.fileContent(fileID)
	if !ok {
		return sourceBlock{}
	}
	lines := splitSourceLines(content)
	if blockHasKeyValue(lines, 0, len(lines), "id", id) {
		return makeBlock(lines, 0, len(lines))
	}
	return sourceBlock{}
}

func (s *Service) fileContent(fileID semantic.FileID) (string, bool) {
	if err := s.requireProject(); err != nil {
		return "", false
	}
	file, ok := s.project.SourceFilesByID[fileID]
	return file.Content, ok
}

func sourceResponse(object ObjectRef, fileID semantic.FileID, block sourceBlock, fallback string, diagnostics []semantic.Diagnostic) GetSourceResponse {
	loc := SourceLocation{File: fileID.String()}
	if block.startLine > 0 {
		loc.Line = block.startLine
		loc.Column = block.column
		loc.EndLine = block.endLine
		loc.EndColumn = 1
	}
	if diagnostics == nil {
		diagnostics = []semantic.Diagnostic{}
	}
	return GetSourceResponse{
		Object:      object,
		Source:      loc,
		Snippet:     SourceSnippet{Language: "yaml", Text: block.text},
		Fallback:    fallback,
		Diagnostics: diagnostics,
	}
}

func splitSourceLines(content string) []string {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")
	return strings.Split(content, "\n")
}

func findTopLevelSequenceItem(lines []string, section string, match map[string]string) sourceBlock {
	start, end := topLevelSectionRange(lines, section)
	if start < 0 {
		return sourceBlock{}
	}
	return findSequenceItemInRange(lines, start, end, match)
}

func findNestedSequenceItem(lines []string, start, end int, section string, match map[string]string) sourceBlock {
	sectionStart, sectionEnd := nestedSectionRange(lines, start, end, section)
	if sectionStart < 0 {
		return sourceBlock{}
	}
	return findSequenceItemInRange(lines, sectionStart, sectionEnd, match)
}

func findSequenceItemInRange(lines []string, start, end int, match map[string]string) sourceBlock {
	for i := start; i < end; i++ {
		trimmed := strings.TrimSpace(lines[i])
		if !strings.HasPrefix(trimmed, "- ") && trimmed != "-" {
			continue
		}
		itemEnd := nextSequenceItemOrSectionEnd(lines, i+1, end, indentOf(lines[i]))
		if blockMatches(lines, i, itemEnd, match) {
			return makeBlock(lines, i, itemEnd)
		}
		i = itemEnd - 1
	}
	return sourceBlock{}
}

func findNestedMappingBlock(lines []string, start, end int, key string) sourceBlock {
	for i := start; i < end; i++ {
		trimmed := strings.TrimSpace(lines[i])
		if !strings.HasPrefix(trimmed, key+":") {
			continue
		}
		keyIndent := indentOf(lines[i])
		blockEnd := i + 1
		for blockEnd < end {
			if strings.TrimSpace(lines[blockEnd]) != "" && indentOf(lines[blockEnd]) <= keyIndent {
				break
			}
			blockEnd++
		}
		return makeBlock(lines, i, blockEnd)
	}
	return sourceBlock{}
}

func topLevelSectionRange(lines []string, section string) (int, int) {
	for i, line := range lines {
		if indentOf(line) == 0 && strings.TrimSpace(line) == section+":" {
			return i + 1, nextTopLevelSection(lines, i+1)
		}
	}
	return -1, -1
}

func nestedSectionRange(lines []string, start, end int, section string) (int, int) {
	for i := start; i < end; i++ {
		if strings.TrimSpace(lines[i]) != section+":" {
			continue
		}
		sectionIndent := indentOf(lines[i])
		sectionEnd := i + 1
		for sectionEnd < end {
			if strings.TrimSpace(lines[sectionEnd]) != "" && indentOf(lines[sectionEnd]) <= sectionIndent {
				break
			}
			sectionEnd++
		}
		return i + 1, sectionEnd
	}
	return -1, -1
}

func nextTopLevelSection(lines []string, start int) int {
	for i := start; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		if indentOf(lines[i]) == 0 && strings.HasSuffix(trimmed, ":") {
			return i
		}
	}
	return len(lines)
}

func nextSequenceItemOrSectionEnd(lines []string, start, end, itemIndent int) int {
	for i := start; i < end; i++ {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed == "" {
			continue
		}
		indent := indentOf(lines[i])
		if indent <= itemIndent && strings.HasPrefix(trimmed, "- ") {
			return i
		}
		if indent < itemIndent {
			return i
		}
	}
	return end
}

func blockMatches(lines []string, start, end int, match map[string]string) bool {
	for key, want := range match {
		if !blockHasKeyValue(lines, start, end, key, want) {
			return false
		}
	}
	return true
}

func blockHasKeyValue(lines []string, start, end int, key, want string) bool {
	for i := start; i < end; i++ {
		k, v, ok := yamlKeyValue(lines[i])
		if !ok || k != key {
			continue
		}
		if want == "" {
			return v == ""
		}
		if v == want {
			return true
		}
	}
	return false
}

func blockHasKey(lines []string, start, end int, key string) bool {
	for i := start; i < end; i++ {
		k, _, ok := yamlKeyValue(lines[i])
		if ok && k == key {
			return true
		}
	}
	return false
}

func yamlKeyValue(line string) (string, string, bool) {
	trimmed := strings.TrimSpace(line)
	if strings.HasPrefix(trimmed, "-") {
		trimmed = strings.TrimSpace(strings.TrimPrefix(trimmed, "-"))
	}
	idx := strings.Index(trimmed, ":")
	if idx <= 0 {
		return "", "", false
	}
	key := strings.TrimSpace(trimmed[:idx])
	value := strings.TrimSpace(trimmed[idx+1:])
	value = stripYAMLComment(value)
	value = strings.Trim(value, "'")
	value = strings.Trim(value, "\"")
	return key, value, true
}

func stripYAMLComment(value string) string {
	inSingle := false
	inDouble := false
	for i, r := range value {
		switch r {
		case '\'':
			if !inDouble {
				inSingle = !inSingle
			}
		case '"':
			if !inSingle {
				inDouble = !inDouble
			}
		case '#':
			if !inSingle && !inDouble {
				return strings.TrimSpace(value[:i])
			}
		}
	}
	return strings.TrimSpace(value)
}

func makeBlock(lines []string, start, end int) sourceBlock {
	if start < 0 || start >= len(lines) || end < start {
		return sourceBlock{}
	}
	if end > len(lines) {
		end = len(lines)
	}
	text := strings.Join(lines[start:end], "\n")
	return sourceBlock{
		startLine: start + 1,
		endLine:   end,
		column:    indentOf(lines[start]) + 1,
		text:      strings.TrimRight(text, "\n"),
	}
}

func indentOf(line string) int {
	return len(line) - len(strings.TrimLeft(line, " \t"))
}
