package query

import (
	"fmt"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func (s *Service) GetSignature(req GetSignatureRequest) (GetSignatureResponse, error) {
	if s.isScenarioSelector(req.Selector) {
		scenario, err := s.scenarioBySelector(req.Selector)
		if err != nil {
			return GetSignatureResponse{}, err
		}
		return GetSignatureResponse{
			Object:      scenarioObjectRef(scenario),
			Signature:   signatureForScenario(scenario),
			Diagnostics: []semantic.Diagnostic{},
		}, nil
	}

	node, err := s.nodeByID(req.Selector.ID)
	if err != nil {
		return GetSignatureResponse{}, err
	}
	sig, doc, err := s.signatureForNode(node)
	if err != nil {
		return GetSignatureResponse{}, err
	}
	return GetSignatureResponse{
		Object:      objectRef(node),
		Signature:   sig,
		Doc:         doc,
		Diagnostics: []semantic.Diagnostic{},
	}, nil
}

func signatureForScenario(scenario *semantic.SequenceScenario) Signature {
	return Signature{
		"id":         scenario.ID,
		"title":      scenario.Title,
		"state_file": scenario.StateFile.String(),
	}
}

func (s *Service) signatureForNode(node semantic.Node) (Signature, string, error) {
	switch n := node.(type) {
	case *semantic.Task:
		sig := Signature{
			"main":   n.Main,
			"params": paramSignatures(n.Params),
			"reads":  storeIDs(n.Reads),
			"writes": storeIDs(n.Writes),
		}
		if ret := returnSignature(n.Returns); ret != nil {
			sig["returns"] = ret
		}
		if n.Endpoint {
			sig["endpoint"] = EndpointSignature{Method: n.Method, LeafPath: n.Path}
		}
		return sig, n.Note, nil
	case *semantic.Model:
		return Signature{
			"model_kind": n.Kind,
			"fields":     fieldSignatures(n.Fields),
		}, n.Note, nil
	case *semantic.Store:
		return Signature{
			"store_kind": n.StoreKind,
			"of":         n.Of.String(),
		}, n.Note, nil
	case *semantic.State:
		return Signature{
			"initial":   n.Initial,
			"final":     n.Final,
			"wireframe": map[string]any{"present": n.Wireframe != nil},
		}, n.Note, nil
	case *semantic.Event:
		sig := Signature{"source": n.Source}
		if n.Actor != "" {
			sig["actor"] = n.Actor
		}
		if n.PayloadModel != "" || n.PayloadName != "" {
			sig["payload"] = map[string]any{"model": n.PayloadModel.String()}
		}
		if n.Watches != "" || n.WatchesName != "" {
			sig["watches"] = n.Watches.String()
		}
		return sig, n.Note, nil
	case *semantic.Branch:
		return Signature{"params": paramSignatures(n.Params)}, n.Note, nil
	case *semantic.Fork:
		return Signature{"params": paramSignatures(n.Params)}, n.Note, nil
	case *semantic.Join:
		sig := Signature{"params": paramSignatures(n.Params)}
		if ret := returnSignature(n.Returns); ret != nil {
			sig["returns"] = ret
		}
		return sig, n.Note, nil
	default:
		return nil, "", fmt.Errorf("unsupported signature kind: %s", node.GetKind())
	}
}
