package resolve

import "github.com/hiroshiasayadev-prog/brewprint/internal/semantic"

type symbolTable struct {
	project *semantic.Project
	diags   []semantic.Diagnostic
}

func newSymbolTable(project *semantic.Project) *symbolTable {
	return &symbolTable{project: project}
}

func (s *symbolTable) addDiagnostic(severity semantic.Severity, fileID semantic.FileID, message string) {
	s.addDiagnosticCode(severity, diagnosticSemanticValidation, fileID, message)
}

func (s *symbolTable) addDiagnosticCode(severity semantic.Severity, code string, fileID semantic.FileID, message string) {
	s.diags = append(s.diags, semantic.Diagnostic{Severity: severity, Code: code, FileID: fileID, Message: message})
}

func (s *symbolTable) addNode(node semantic.Node) {
	qid := node.GetQID()
	fileID := node.GetFileID()
	if isFilePrivateSubNode(node) {
		if hasFilePrivateSubNodeID(s.project.NodesByFile[fileID], node.GetID()) {
			s.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateSubNode, fileID, "duplicate sub node local id: "+node.GetID())
		}
		s.project.NodesByFile[fileID] = append(s.project.NodesByFile[fileID], node)
		if _, exists := s.project.NodesByID[qid]; !exists {
			s.project.NodesByID[qid] = node
			s.addNodeByKindID(node, qid)
		}
		return
	}

	if _, exists := s.project.NodesByQID[qid]; exists {
		s.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateNode, fileID, "duplicate node qid: "+qid.String())
		return
	}

	s.project.NodesByQID[qid] = node
	s.project.NodesByID[qid] = node
	s.project.NodesByFile[fileID] = append(s.project.NodesByFile[fileID], node)
	if node.IsMain() {
		if _, exists := s.project.MainNodeByFile[fileID]; exists {
			s.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateMainNode, fileID, "duplicate main node in file")
		} else {
			s.project.MainNodeByFile[fileID] = qid
		}
	}

	s.addNodeByKindID(node, qid)
}

func isFilePrivateSubNode(node semantic.Node) bool {
	if node == nil || node.IsMain() {
		return false
	}
	switch node.GetKind() {
	case semantic.NodeKindTask, semantic.NodeKindBranch, semantic.NodeKindFork, semantic.NodeKindJoin:
		return true
	default:
		return false
	}
}

func hasFilePrivateSubNodeID(nodes []semantic.Node, id string) bool {
	for _, node := range nodes {
		if isFilePrivateSubNode(node) && node.GetID() == id {
			return true
		}
	}
	return false
}

func (s *symbolTable) addNodeByKindID(node semantic.Node, qid semantic.QualifiedID) {
	fileID := node.GetFileID()
	switch n := node.(type) {
	case *semantic.Task:
		s.project.TasksByQID[qid] = n
	case *semantic.Model:
		s.project.ModelsByQID[qid] = n
	case *semantic.State:
		s.project.StatesByQID[qid] = n
	case *semantic.Event:
		s.project.EventsByQID[qid] = n
	case *semantic.Store:
		s.project.StoresByQID[qid] = n
	case *semantic.Branch:
		s.project.BranchesByQID[qid] = n
	case *semantic.Fork:
		s.project.ForksByQID[qid] = n
	case *semantic.Join:
		s.project.JoinsByQID[qid] = n
	case *semantic.Actor:
		if _, exists := s.project.ActorsByQID[qid]; exists {
			s.addDiagnosticCode(semantic.SeverityError, diagnosticDuplicateActor, fileID, "duplicate actor id: "+n.ID)
			return
		}
		s.project.ActorsByQID[qid] = n
	}
}
