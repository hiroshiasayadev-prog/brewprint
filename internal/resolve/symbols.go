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
	s.diags = append(s.diags, semantic.Diagnostic{Severity: severity, FileID: fileID, Message: message})
}

func (s *symbolTable) addNode(node semantic.Node) {
	qid := node.GetQID()
	fileID := node.GetFileID()
	if _, exists := s.project.NodesByQID[qid]; exists {
		s.addDiagnostic(semantic.SeverityError, fileID, "duplicate node qid: "+qid.String())
		return
	}

	s.project.NodesByQID[qid] = node
	s.project.NodesByFile[fileID] = append(s.project.NodesByFile[fileID], node)
	if node.IsMain() {
		if _, exists := s.project.MainNodeByFile[fileID]; exists {
			s.addDiagnostic(semantic.SeverityError, fileID, "duplicate main node in file")
		} else {
			s.project.MainNodeByFile[fileID] = qid
		}
	}

	switch n := node.(type) {
	case *semantic.Task:
		s.project.TasksByQID[qid] = n
	case *semantic.Model:
		s.project.ModelsByQID[qid] = n
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
			s.addDiagnostic(semantic.SeverityError, fileID, "duplicate actor id: "+n.ID)
			return
		}
		s.project.ActorsByQID[qid] = n
	}
}
