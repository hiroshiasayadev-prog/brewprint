package resolve

import (
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func TestPrivateSubTaskSameLocalIDAcrossFilesOK(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		privateSubTaskFile("shop/task/checkout.yaml", "checkout", "helper"),
		privateSubTaskFile("shop/task/refund.yaml", "refund", "helper"),
	}})

	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	if _, ok := project.NodesByQID["shop.task.helper"]; ok {
		t.Fatalf("private sub task was registered in NodesByQID")
	}
	if project.TasksByQID["shop/task/checkout.yaml#helper"] == nil {
		t.Fatalf("checkout helper missing from private task index")
	}
	if project.TasksByQID["shop/task/refund.yaml#helper"] == nil {
		t.Fatalf("refund helper missing from private task index")
	}
}

func TestPrivateSubTaskReturnAssetIdentityUsesInternalQID(t *testing.T) {
	checkoutFile := privateReturningSubTaskFile("shop/task/checkout.yaml", "checkout", "helper")
	refundFile := privateReturningSubTaskFile("shop/task/refund.yaml", "refund", "helper")
	checkoutFile.NodeFile.Models = nil
	refundFile.NodeFile.Models = nil
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		typeRefStructModelFile("shop/model/receipt.yaml", "receipt"),
		checkoutFile,
		refundFile,
	}})

	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	checkout := project.TasksByQID["shop/task/checkout.yaml#helper"]
	refund := project.TasksByQID["shop/task/refund.yaml#helper"]
	if checkout == nil || checkout.Returns == nil || checkout.Returns.Asset == nil {
		t.Fatalf("checkout helper return asset missing")
	}
	if refund == nil || refund.Returns == nil || refund.Returns.Asset == nil {
		t.Fatalf("refund helper return asset missing")
	}
	if checkout.Returns.Asset.ProducedBy != "shop/task/checkout.yaml#helper" {
		t.Fatalf("checkout asset producer = %q, want internal qid", checkout.Returns.Asset.ProducedBy)
	}
	if refund.Returns.Asset.ProducedBy != "shop/task/refund.yaml#helper" {
		t.Fatalf("refund asset producer = %q, want internal qid", refund.Returns.Asset.ProducedBy)
	}
	checkoutAssetID := semantic.AssetID(checkout.Returns.Asset.ProducedBy, checkout.Returns.Asset.Name)
	refundAssetID := semantic.AssetID(refund.Returns.Asset.ProducedBy, refund.Returns.Asset.Name)
	if checkoutAssetID == refundAssetID {
		t.Fatalf("private return asset IDs collided: %q", checkoutAssetID)
	}
	if semantic.AssetObjectKey(checkout.Returns.Asset) == semantic.AssetObjectKey(refund.Returns.Asset) {
		t.Fatalf("private return asset object keys collided: %q", semantic.AssetObjectKey(checkout.Returns.Asset))
	}
}

func TestPrivateSubTaskPublicAliasNotRegisteredAsTask(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		privateSubTaskFile("shop/task/checkout.yaml", "checkout", "helper"),
		privateSubTaskFile("shop/task/refund.yaml", "refund", "helper"),
	}})

	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	if project.NodesByID["shop.task.helper"] != nil {
		t.Fatalf("private sub task public alias was registered in NodesByID")
	}
	if project.TasksByQID["shop.task.helper"] != nil {
		t.Fatalf("private sub task public alias was registered in TasksByQID")
	}
}

func TestTransitionActionFullQIDDoesNotUsePrivatePublicAlias(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		privateSubTaskFile("shop/task/checkout.yaml", "checkout", "helper"),
		{
			ID:   "shop/state.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{
				States: []rawyaml.State{
					{ID: "idle"},
					{ID: "done"},
				},
				Events: []rawyaml.Event{{ID: "submitted"}},
				Transitions: []rawyaml.Transition{{
					From:   "idle",
					On:     "submitted",
					To:     "done",
					Action: "shop.task.helper",
				}},
			},
		},
	}})

	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	if len(project.ActionsByTask["shop.task.helper"]) != 0 {
		t.Fatalf("private public-shaped alias was indexed as a transition action")
	}
	if len(project.ActionsByTask["shop/task/checkout.yaml#helper"]) != 0 {
		t.Fatalf("full public-shaped transition action resolved to private task")
	}
}

func TestPrivateSubTaskDuplicateWithinFileDiagnostic(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{{
		ID:   "shop/task/checkout.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{
			Tasks: []rawyaml.Task{
				{ID: "checkout", Main: true},
				{ID: "helper"},
				{ID: "helper"},
			},
			Flow: []rawyaml.FlowEntry{{Step: "helper"}},
		},
	}}})

	assertDiagnosticCode(t, diagnostics, diagnosticDuplicateSubNode)
	assertNoDiagnosticCode(t, diagnostics, diagnosticDuplicateNode)
	assertNoDiagnosticCode(t, diagnostics, diagnosticUnresolvedFlowTask)
}

func TestMainTaskSameModuleDuplicateNode(t *testing.T) {
	_, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "shop/task/checkout.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Tasks: []rawyaml.Task{
				{ID: "run", Main: true},
			}},
		},
		{
			ID:   "shop/task/run.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Tasks: []rawyaml.Task{
				{ID: "run", Main: true},
			}},
		},
	}})

	assertDiagnosticCode(t, diagnostics, diagnosticDuplicateNode)
	assertNoDiagnosticCode(t, diagnostics, diagnosticDuplicateSubNode)
}

func TestLocalFlowStepPrefersSameFilePrivateSubTask(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "shop/task/use_shared.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{
				Tasks: []rawyaml.Task{
					{ID: "use_shared", Main: true},
					{ID: "shared"},
				},
				Flow: []rawyaml.FlowEntry{{Step: "shared"}},
			},
		},
		{
			ID:   "shop/task/shared.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Tasks: []rawyaml.Task{
				{ID: "shared", Main: true},
			}},
		},
	}})

	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	flow := project.FlowByFile["shop/task/use_shared.yaml"]
	if len(flow) != 1 || flow[0].Step.Task != "shop/task/use_shared.yaml#shared" {
		t.Fatalf("flow step resolved to %q, want same-file private sub task", flow[0].Step.Task)
	}
}

func TestBareFlowStepFallsBackToSameModuleMainTask(t *testing.T) {
	project, diagnostics := Build(&rawyaml.Project{Files: []rawyaml.File{
		{
			ID:   "shop/task/use_shared.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{
				Tasks: []rawyaml.Task{
					{ID: "use_shared", Main: true},
				},
				Flow: []rawyaml.FlowEntry{{Step: "shared"}},
			},
		},
		{
			ID:   "shop/task/shared.yaml",
			Kind: rawyaml.FileKindNode,
			NodeFile: &rawyaml.NodeFile{Tasks: []rawyaml.Task{
				{ID: "shared", Main: true},
			}},
		},
	}})

	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	flow := project.FlowByFile["shop/task/use_shared.yaml"]
	if len(flow) != 1 || flow[0].Step.Task != "shop.task.shared" {
		t.Fatalf("flow step resolved to %q, want same-module main task", flow[0].Step.Task)
	}
}

func privateSubTaskFile(fileID string, mainID string, helperID string) rawyaml.File {
	return rawyaml.File{
		ID:   fileID,
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{
			Tasks: []rawyaml.Task{
				{ID: mainID, Main: true},
				{ID: helperID},
			},
			Flow: []rawyaml.FlowEntry{{Step: helperID}},
		},
	}
}

func privateReturningSubTaskFile(fileID string, mainID string, helperID string) rawyaml.File {
	file := privateSubTaskFile(fileID, mainID, helperID)
	file.NodeFile.Models = []rawyaml.Model{{
		ID:     "receipt",
		Kind:   "struct",
		Fields: []rawyaml.ModelField{{Name: "id", Type: "str"}},
	}}
	file.NodeFile.Tasks[1].Returns = &rawyaml.Return{Name: "result", Model: "receipt"}
	return file
}
