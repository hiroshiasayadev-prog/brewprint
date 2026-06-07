package query

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"github.com/hiroshiasayadev-prog/brewprint/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/internal/source"
)

func TestQueryServiceUC001(t *testing.T) {
	service := newUC001Service(t)

	t.Run("ListObjects", func(t *testing.T) {
		orderTasks, err := service.ListObjects(ListObjectsRequest{Object: "node", Kind: "task", Module: "order"})
		if err != nil {
			t.Fatalf("ListObjects order tasks: %v", err)
		}
		checkout := assertHasObject(t, orderTasks.Objects, "node", "task", "order.task.checkout")
		if checkout.Module != "order" || checkout.Source["file"] != "order/task/checkout.yaml" {
			t.Fatalf("checkout object = %#v", checkout)
		}

		fields, err := service.ListObjects(ListObjectsRequest{Object: "field", Module: "order"})
		if err != nil {
			t.Fatalf("ListObjects order fields: %v", err)
		}
		assertHasObject(t, fields.Objects, "field", "field", "order.model.order.id")

		modelFieldAlias, err := service.ListObjects(ListObjectsRequest{Kind: "model_field", Module: "order"})
		if err != nil {
			t.Fatalf("ListObjects model_field alias: %v", err)
		}
		assertHasObject(t, modelFieldAlias.Objects, "field", "field", "order.model.order.id")

		views, err := service.ListObjects(ListObjectsRequest{Object: "view"})
		if err != nil {
			t.Fatalf("ListObjects views: %v", err)
		}
		assertHasObject(t, views.Objects, "view", "api_table", "ec_api")

		transitions, err := service.ListObjects(ListObjectsRequest{Object: "transition", File: "order/state.yaml"})
		if err != nil {
			t.Fatalf("ListObjects transitions: %v", err)
		}
		assertHasObject(t, transitions.Objects, "transition", "transition", "order/state.yaml#checkout_screen:checkout_submitted")
	})

	t.Run("InspectFile", func(t *testing.T) {
		stateFile, err := service.Inspect(InspectRequest{Selector: Selector{Object: "file", Kind: "state_file", ID: "order/state.yaml"}})
		if err != nil {
			t.Fatalf("Inspect state file: %v", err)
		}
		if stateFile.Object.Kind != "state_file" || stateFile.Signature["kind"] != "state_file" {
			t.Fatalf("state file inspect = %#v", stateFile)
		}
		if len(stateFile.Members["states"].([]ObjectRef)) != 5 || len(stateFile.Members["events"].([]ObjectRef)) != 4 || len(stateFile.Members["transitions"].([]TransitionRef)) != 5 {
			t.Fatalf("state file members = %#v", stateFile.Members)
		}

		checkoutFile, err := service.Inspect(InspectRequest{Selector: Selector{Object: "file", ID: "order/task/checkout.yaml"}})
		if err != nil {
			t.Fatalf("Inspect checkout file: %v", err)
		}
		if checkoutFile.Object.Kind != "node" || checkoutFile.Members["main_node"].(ObjectRef).ID != "order.task.checkout" {
			t.Fatalf("checkout file inspect = %#v", checkoutFile)
		}
		if _, ok := checkoutFile.Members["flow"]; !ok {
			t.Fatalf("checkout file flow missing: %#v", checkoutFile.Members)
		}
	})

	t.Run("GetSignature", func(t *testing.T) {
		login, err := service.GetSignature(GetSignatureRequest{Selector: Selector{ID: "auth.task.login"}})
		if err != nil {
			t.Fatalf("GetSignature login: %v", err)
		}
		if login.Object.Kind != "task" {
			t.Fatalf("login kind = %s, want task", login.Object.Kind)
		}
		if login.Signature["main"] != true {
			t.Fatalf("login main = %#v, want true", login.Signature["main"])
		}
		params := login.Signature["params"].([]ParamSignature)
		if len(params) != 1 || params[0].Model != "auth.model.login_form" {
			t.Fatalf("login params = %#v", params)
		}
		ret := login.Signature["returns"].(*ReturnSignature)
		if ret.Name != "auth_token" || ret.Model != "auth.model.token" || ret.Asset == nil {
			t.Fatalf("login returns = %#v", ret)
		}
		endpoint := login.Signature["endpoint"].(EndpointSignature)
		if endpoint.Method != "POST" || endpoint.LeafPath != "login" {
			t.Fatalf("login endpoint = %#v", endpoint)
		}

		model, err := service.GetSignature(GetSignatureRequest{Selector: Selector{ID: "auth.model.login_form"}})
		if err != nil {
			t.Fatalf("GetSignature model: %v", err)
		}
		fields := model.Signature["fields"].([]FieldSignature)
		if len(fields) == 0 {
			t.Fatalf("login_form fields is empty")
		}

		store, err := service.GetSignature(GetSignatureRequest{Selector: Selector{ID: "auth.store.user_db"}})
		if err != nil {
			t.Fatalf("GetSignature store: %v", err)
		}
		if store.Signature["store_kind"] != "db" || store.Signature["of"] != "auth.model.credential" {
			t.Fatalf("user_db signature = %#v", store.Signature)
		}

		join, err := service.GetSignature(GetSignatureRequest{Selector: Selector{Object: "node", Kind: "join", File: "order/task/checkout.yaml", LocalID: "finalize_checkout"}})
		if err != nil {
			t.Fatalf("GetSignature join: %v", err)
		}
		joinParams := join.Signature["params"].([]ParamSignature)
		if len(joinParams) != 2 {
			t.Fatalf("join params = %#v", joinParams)
		}

		state, err := service.GetSignature(GetSignatureRequest{Selector: Selector{ID: "order.state.checkout_screen"}})
		if err != nil {
			t.Fatalf("GetSignature state: %v", err)
		}
		if state.Signature["initial"] != false || state.Signature["final"] != false {
			t.Fatalf("state signature = %#v", state.Signature)
		}
		wireframe := state.Signature["wireframe"].(map[string]any)
		if wireframe["present"] != true {
			t.Fatalf("state wireframe signature = %#v", wireframe)
		}

		event, err := service.GetSignature(GetSignatureRequest{Selector: Selector{ID: "order.event.payment_webhook_received"}})
		if err != nil {
			t.Fatalf("GetSignature event: %v", err)
		}
		if event.Signature["source"] != "external" || event.Signature["actor"] != "stripe" {
			t.Fatalf("event signature = %#v", event.Signature)
		}
		payload := event.Signature["payload"].(map[string]any)
		if payload["model"] != "payment.model.payment_event" {
			t.Fatalf("event payload signature = %#v", payload)
		}

		scenario, err := service.GetSignature(GetSignatureRequest{Selector: Selector{Object: "view", Kind: "sequence_diagram", ID: "checkout_flow"}})
		if err != nil {
			t.Fatalf("GetSignature scenario: %v", err)
		}
		if scenario.Object.Object != "view" || scenario.Object.Kind != "sequence_diagram" || scenario.Object.ID != "checkout_flow" {
			t.Fatalf("scenario object = %#v", scenario.Object)
		}
		if scenario.Signature["id"] != "checkout_flow" || scenario.Signature["title"] != "チェックアウトフロー" || scenario.Signature["state_file"] != "order/state.yaml" {
			t.Fatalf("scenario signature = %#v", scenario.Signature)
		}

		transition, err := service.GetSignature(GetSignatureRequest{Selector: Selector{Object: "transition", ID: "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"}})
		if err != nil {
			t.Fatalf("GetSignature transition: %v", err)
		}
		if transition.Object.Object != "transition" || transition.Object.ID != "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']" {
			t.Fatalf("transition object = %#v", transition.Object)
		}
		if transition.Signature["from"] != "processing" || transition.Signature["on"] != "payment_webhook_received" || transition.Signature["to"] != "confirmed" || transition.Signature["guard"] != "payload.status == 'succeeded'" || transition.Signature["action"] != "payment.webhooks.task.process_payment" {
			t.Fatalf("transition signature = %#v", transition.Signature)
		}

		field, err := service.GetSignature(GetSignatureRequest{Selector: Selector{Object: "field", ID: "order.model.order", LocalID: "id"}})
		if err != nil {
			t.Fatalf("GetSignature field: %v", err)
		}
		if field.Object.Object != "field" || field.Object.ID != "order.model.order.id" || field.Object.LocalID != "id" {
			t.Fatalf("field object = %#v", field.Object)
		}
		if field.Signature["name"] != "id" || field.Signature["type"] != "str" || field.Signature["pk"] != true {
			t.Fatalf("field signature = %#v", field.Signature)
		}
	})

	t.Run("GetSource", func(t *testing.T) {
		login, err := service.GetSource(GetSourceRequest{Selector: Selector{ID: "auth.task.login"}})
		if err != nil {
			t.Fatalf("GetSource login: %v", err)
		}
		if login.Object.ID != "auth.task.login" || login.Source.File != "auth/task/login.yaml" || login.Snippet.Language != "yaml" {
			t.Fatalf("login source envelope = %#v", login)
		}
		if !strings.Contains(login.Snippet.Text, "id: login") || !strings.Contains(login.Snippet.Text, "type: task") || !strings.Contains(login.Snippet.Text, "returns:") {
			t.Fatalf("login source snippet = %q", login.Snippet.Text)
		}

		field, err := service.GetSource(GetSourceRequest{Selector: Selector{Object: "field", ID: "order.model.order", LocalID: "id"}})
		if err != nil {
			t.Fatalf("GetSource field: %v", err)
		}
		if field.Object.ID != "order.model.order.id" || !strings.Contains(field.Snippet.Text, "name: id") || !strings.Contains(field.Snippet.Text, "pk: true") {
			t.Fatalf("field source = %#v", field)
		}

		transition, err := service.GetSource(GetSourceRequest{Selector: Selector{Object: "transition", ID: "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"}})
		if err != nil {
			t.Fatalf("GetSource transition: %v", err)
		}
		if transition.Object.ID != "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']" || !strings.Contains(transition.Snippet.Text, "from: processing") || !strings.Contains(transition.Snippet.Text, "payload.status == 'succeeded'") {
			t.Fatalf("transition source = %#v", transition)
		}

		fileSource, err := service.GetSource(GetSourceRequest{Selector: Selector{Object: "file", ID: "views/api_table.yaml"}})
		if err != nil {
			t.Fatalf("GetSource file: %v", err)
		}
		if fileSource.Object.Object != "file" || fileSource.Source.File != "views/api_table.yaml" || !strings.Contains(fileSource.Snippet.Text, "as: api_table") {
			t.Fatalf("file source = %#v", fileSource)
		}
	})

	t.Run("GetReferences", func(t *testing.T) {
		loginRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "auth.task.login"}})
		if err != nil {
			t.Fatalf("GetReferences login: %v", err)
		}
		assertHasReference(t, loginRefs.References, "param_model", "out", "auth.task.login", "auth.model.login_form")
		assertHasReference(t, loginRefs.References, "return_model", "out", "auth.task.login", "auth.model.token")
		assertHasReference(t, loginRefs.References, "produces_asset", "out", "auth.task.login", "")
		assertHasReference(t, loginRefs.References, "reads", "out", "auth.task.login", "auth.store.user_db")
		assertHasReference(t, loginRefs.References, "writes", "out", "auth.task.login", "auth.store.session_store")

		storeRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "auth.store.user_db"}, Direction: "in"})
		if err != nil {
			t.Fatalf("GetReferences store: %v", err)
		}
		assertHasReference(t, storeRefs.References, "reads", "in", "auth.task.login", "auth.store.user_db")

		modelRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "auth.model.login_form"}, Direction: "in"})
		if err != nil {
			t.Fatalf("GetReferences model: %v", err)
		}
		assertHasReference(t, modelRefs.References, "param_model", "in", "auth.task.login", "auth.model.login_form")

		readRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "auth.task.login"}, Kinds: []string{"reads"}})
		if err != nil {
			t.Fatalf("GetReferences reads filter: %v", err)
		}
		if len(readRefs.References) != 2 {
			t.Fatalf("reads filtered refs len = %d, want 2: %#v", len(readRefs.References), readRefs.References)
		}
		for _, ref := range readRefs.References {
			if ref.Kind != "reads" {
				t.Fatalf("non-read ref in reads filter: %#v", ref)
			}
		}

		checkoutActionRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "order.task.checkout"}, Direction: "in", Kinds: []string{"transition_action"}})
		if err != nil {
			t.Fatalf("GetReferences checkout transition_action: %v", err)
		}
		checkoutAction := assertHasReference(t, checkoutActionRefs.References, "transition_action", "in", "order/state.yaml#checkout_screen:checkout_submitted", "order.task.checkout")
		if checkoutAction.From.StateFile != "order/state.yaml" || checkoutAction.From.FromState != "checkout_screen" || checkoutAction.From.On != "checkout_submitted" || checkoutAction.From.ToState != "processing" || checkoutAction.From.Action != "order.task.checkout" {
			t.Fatalf("checkout transition endpoint = %#v", checkoutAction.From)
		}

		paymentActionRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "payment.webhooks.task.process_payment"}, Direction: "in", Kinds: []string{"transition_action"}})
		if err != nil {
			t.Fatalf("GetReferences payment transition_action: %v", err)
		}
		paymentAction := assertHasReference(t, paymentActionRefs.References, "transition_action", "in", "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']", "payment.webhooks.task.process_payment")
		if paymentAction.From.Guard != "payload.status == 'succeeded'" || paymentAction.From.Action != "payment.webhooks.task.process_payment" {
			t.Fatalf("payment transition endpoint = %#v", paymentAction.From)
		}

		paymentEventRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "order.event.payment_webhook_received"}, Direction: "both"})
		if err != nil {
			t.Fatalf("GetReferences payment event: %v", err)
		}
		assertHasReference(t, paymentEventRefs.References, "event_payload", "out", "order.event.payment_webhook_received", "payment.model.payment_event")
		assertHasReference(t, paymentEventRefs.References, "event_actor", "out", "order.event.payment_webhook_received", "stripe")
		assertHasReference(t, paymentEventRefs.References, "transition_event", "in", "order/state.yaml#processing:payment_webhook_received[payload.status == 'failed']", "order.event.payment_webhook_received")
		assertHasReference(t, paymentEventRefs.References, "transition_event", "in", "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']", "order.event.payment_webhook_received")

		payloadModelRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "payment.model.payment_event"}, Direction: "in", Kinds: []string{"event_payload"}})
		if err != nil {
			t.Fatalf("GetReferences payment model event_payload: %v", err)
		}
		assertHasReference(t, payloadModelRefs.References, "event_payload", "in", "order.event.payment_webhook_received", "payment.model.payment_event")

		inventoryStoreRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "catalog.store.inventory_db"}, Direction: "in", Kinds: []string{"event_watches"}})
		if err != nil {
			t.Fatalf("GetReferences inventory store event_watches: %v", err)
		}
		assertHasReference(t, inventoryStoreRefs.References, "event_watches", "in", "inventory.event.inventory_changed", "catalog.store.inventory_db")

		processingStateRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "order.state.processing"}, Direction: "both", Kinds: []string{"transition_from", "transition_to"}})
		if err != nil {
			t.Fatalf("GetReferences processing state transitions: %v", err)
		}
		assertHasReference(t, processingStateRefs.References, "transition_to", "in", "order/state.yaml#checkout_screen:checkout_submitted", "order.state.processing")
		assertHasReference(t, processingStateRefs.References, "transition_from", "in", "order/state.yaml#processing:payment_webhook_received[payload.status == 'failed']", "order.state.processing")
		assertHasReference(t, processingStateRefs.References, "transition_from", "in", "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']", "order.state.processing")

		checkoutScenarioRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{Object: "view", Kind: "sequence_diagram", ID: "checkout_flow"}, Direction: "out"})
		if err != nil {
			t.Fatalf("GetReferences checkout scenario: %v", err)
		}
		assertHasReference(t, checkoutScenarioRefs.References, "scenario_state_file", "out", "checkout_flow", "order/state.yaml")
		assertHasReference(t, checkoutScenarioRefs.References, "scenario_step_transition", "out", "scenario_step:checkout_flow:1", "order/state.yaml#cart:view_checkout")
		assertHasReference(t, checkoutScenarioRefs.References, "scenario_step_transition", "out", "scenario_step:checkout_flow:2", "order/state.yaml#checkout_screen:checkout_submitted")

		transitionRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{Object: "transition", ID: "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"}, Direction: "both"})
		if err != nil {
			t.Fatalf("GetReferences transition: %v", err)
		}
		if transitionRefs.Object.Object != "transition" || transitionRefs.Object.ID != "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']" {
			t.Fatalf("transition object = %#v", transitionRefs.Object)
		}
		assertHasReference(t, transitionRefs.References, "transition_from", "out", "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']", "order.state.processing")
		assertHasReference(t, transitionRefs.References, "transition_event", "out", "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']", "order.event.payment_webhook_received")
		assertHasReference(t, transitionRefs.References, "transition_to", "out", "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']", "order.state.confirmed")
		assertHasReference(t, transitionRefs.References, "transition_action", "out", "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']", "payment.webhooks.task.process_payment")
		assertHasReference(t, transitionRefs.References, "scenario_step_transition", "in", "scenario_step:payment_webhook_flow:1", "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']")

		fieldRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{Object: "field", ID: "order.model.order", LocalID: "id"}, Direction: "both"})
		if err != nil {
			t.Fatalf("GetReferences field: %v", err)
		}
		if fieldRefs.Object.Object != "field" || fieldRefs.Object.ID != "order.model.order.id" || fieldRefs.Object.LocalID != "id" {
			t.Fatalf("field object = %#v", fieldRefs.Object)
		}
		assertHasReference(t, fieldRefs.References, "field_type", "out", "order.model.order.id", "str")
		assertHasReference(t, fieldRefs.References, "field_fk", "in", "payment.model.payment_event.order_id", "order.model.order.id")
		assertHasReference(t, fieldRefs.References, "field_fk", "in", "order.model.order_item.order_id", "order.model.order.id")

		stateFileRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{Object: "file", Kind: "state_file", ID: "order/state.yaml"}, Direction: "in"})
		if err != nil {
			t.Fatalf("GetReferences state file: %v", err)
		}
		if stateFileRefs.Object.Object != "file" || stateFileRefs.Object.Kind != "state_file" || stateFileRefs.Object.ID != "order/state.yaml" {
			t.Fatalf("state file object = %#v", stateFileRefs.Object)
		}
		assertHasReference(t, stateFileRefs.References, "scenario_state_file", "in", "checkout_flow", "order/state.yaml")
		assertHasReference(t, stateFileRefs.References, "scenario_state_file", "in", "payment_webhook_flow", "order/state.yaml")

		stateFileTransitionRefs, err := service.GetReferences(GetReferencesRequest{
			Selector:  Selector{Object: "file", Kind: "state_file", ID: "order/state.yaml"},
			Direction: "out",
			Kinds:     []string{"transition_from", "transition_event", "transition_to", "transition_action"},
		})
		if err != nil {
			t.Fatalf("GetReferences state file transition-owned refs: %v", err)
		}
		assertHasReference(t, stateFileTransitionRefs.References, "transition_from", "out", "order/state.yaml#checkout_screen:checkout_submitted", "order.state.checkout_screen")
		assertHasReference(t, stateFileTransitionRefs.References, "transition_event", "out", "order/state.yaml#checkout_screen:checkout_submitted", "order.event.checkout_submitted")
		assertHasReference(t, stateFileTransitionRefs.References, "transition_to", "out", "order/state.yaml#checkout_screen:checkout_submitted", "order.state.processing")
		assertHasReference(t, stateFileTransitionRefs.References, "transition_action", "out", "order/state.yaml#checkout_screen:checkout_submitted", "order.task.checkout")

		nodeFileRefs, err := service.GetReferences(GetReferencesRequest{
			Selector:  Selector{Object: "file", Kind: "node", ID: "auth/task/login.yaml"},
			Direction: "out",
			Kinds:     []string{"reads"},
		})
		if err != nil {
			t.Fatalf("GetReferences node file refs: %v", err)
		}
		assertHasReference(t, nodeFileRefs.References, "reads", "out", "auth.task.login", "auth.store.user_db")

		nodeFileFlowRefs, err := service.GetReferences(GetReferencesRequest{
			Selector:  Selector{Object: "file", Kind: "node", ID: "order/task/checkout.yaml"},
			Direction: "in",
			Kinds:     []string{"consumes_asset"},
		})
		if err != nil {
			t.Fatalf("GetReferences node file flow refs: %v", err)
		}
		if len(nodeFileFlowRefs.References) != 0 {
			t.Fatalf("node file refs included raw flow wiring: %#v", nodeFileFlowRefs.References)
		}
	})

	t.Run("GetReferenceTree", func(t *testing.T) {
		transitionSelector := Selector{Object: "transition", ID: "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"}
		tree, err := service.GetReferenceTree(GetReferenceTreeRequest{
			Selector:  transitionSelector,
			Direction: "out",
			Depth:     1,
			Kinds:     []string{"transition_from", "transition_event", "transition_to", "transition_action"},
		})
		if err != nil {
			t.Fatalf("GetReferenceTree transition: %v", err)
		}
		if tree.Root.ID != transitionSelector.ID || tree.Direction != "out" || tree.Depth != 1 || tree.Truncated {
			t.Fatalf("transition tree envelope = %#v", tree)
		}
		if len(tree.Nodes) != 5 || len(tree.Edges) != 4 {
			t.Fatalf("transition tree nodes=%d edges=%d tree=%#v", len(tree.Nodes), len(tree.Edges), tree)
		}
		assertHasReferenceTreeNode(t, tree.Nodes, "order.state.processing", 1, []string{"transition_from"})
		assertHasReferenceTreeNode(t, tree.Nodes, "order.event.payment_webhook_received", 1, []string{"transition_event"})
		assertHasReferenceTreeNode(t, tree.Nodes, "order.state.confirmed", 1, []string{"transition_to"})
		assertHasReferenceTreeNode(t, tree.Nodes, "payment.webhooks.task.process_payment", 1, []string{"transition_action"})
		assertHasReferenceTreeEdge(t, tree.Edges, "transition_action", "out", transitionSelector.ID, "payment.webhooks.task.process_payment", 1)

		rootOnly, err := service.GetReferenceTree(GetReferenceTreeRequest{Selector: transitionSelector, Direction: "out", Depth: 0})
		if err != nil {
			t.Fatalf("GetReferenceTree depth 0: %v", err)
		}
		if len(rootOnly.Nodes) != 1 || len(rootOnly.Edges) != 0 || rootOnly.Nodes[0].Depth != 0 || rootOnly.Nodes[0].Object.ID != transitionSelector.ID {
			t.Fatalf("depth 0 tree = %#v", rootOnly)
		}

		truncated, err := service.GetReferenceTree(GetReferenceTreeRequest{Selector: transitionSelector, Direction: "out", Depth: 1, MaxEdges: 1})
		if err != nil {
			t.Fatalf("GetReferenceTree truncation: %v", err)
		}
		if !truncated.Truncated || len(truncated.TruncatedReasons) != 1 || truncated.TruncatedReasons[0] != "max_edges" {
			t.Fatalf("truncated tree = %#v", truncated)
		}

		nodeFileTree, err := service.GetReferenceTree(GetReferenceTreeRequest{
			Selector:  Selector{Object: "file", Kind: "node", ID: "auth/task/login.yaml"},
			Direction: "out",
			Depth:     1,
			Kinds:     []string{"reads"},
		})
		if err != nil {
			t.Fatalf("GetReferenceTree node file: %v", err)
		}
		if nodeFileTree.Root.Object != "file" || nodeFileTree.Root.Kind != "node" {
			t.Fatalf("node file tree root = %#v", nodeFileTree.Root)
		}
		assertHasReferenceTreeEdge(t, nodeFileTree.Edges, "reads", "out", "auth.task.login", "auth.store.user_db", 1)
		assertHasReferenceTreeNode(t, nodeFileTree.Nodes, "auth.store.user_db", 1, []string{"reads"})

		if _, err := service.GetReferenceTree(GetReferenceTreeRequest{Selector: transitionSelector, Direction: "out", Depth: 5}); err == nil {
			t.Fatalf("GetReferenceTree depth 5 expected error")
		}
	})

	t.Run("AnalyzeImpactCore", func(t *testing.T) {
		taskRemove, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{ID: "payment.webhooks.task.process_payment"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeRemove},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact task remove: %v", err)
		}
		if taskRemove.Target.ID != "payment.webhooks.task.process_payment" || taskRemove.Change.Kind != AnalyzeImpactChangeRemove || taskRemove.Truncated {
			t.Fatalf("AnalyzeImpact task remove envelope = %#v", taskRemove)
		}
		removedAction := assertHasImpact(t, taskRemove.Impacts, "transition_action", "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']")
		if removedAction.Severity != "breaking" || removedAction.Fixability != "manual_review" || strings.Join(removedAction.Via, ",") != "transition_action" {
			t.Fatalf("task remove transition action impact = %#v", removedAction)
		}
		removedScenario := assertHasImpact(t, taskRemove.Impacts, "sequence_step_action", "payment_webhook_flow")
		if removedScenario.Severity != "breaking" || removedScenario.Fixability != "manual_review" || strings.Join(removedScenario.Via, ",") != "scenario_step_transition,transition_action" {
			t.Fatalf("task remove sequence impact = %#v", removedScenario)
		}
		removedRender := assertHasImpact(t, taskRemove.Impacts, "render_output", "commerce/dag-process_payment.md")
		if removedRender.Severity != "warning" || removedRender.Fixability == "mechanical" || !strings.Contains(removedRender.Reason, "commerce/dag-process_payment.md") || !strings.Contains(removedRender.RecommendedAction, "brewprint render") {
			t.Fatalf("task remove render output impact = %#v", removedRender)
		}
		if taskRemove.Summary.BySeverity["breaking"] == 0 || taskRemove.Summary.BySeverity["warning"] == 0 || taskRemove.Summary.ByKind["transition_action"] != 1 || taskRemove.Summary.ByKind["sequence_step_action"] != 1 || taskRemove.Summary.ByKind["render_output"] == 0 {
			t.Fatalf("task remove summary = %#v impacts=%#v", taskRemove.Summary, taskRemove.Impacts)
		}

		taskRename, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{ID: "payment.webhooks.task.process_payment"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeRename, NewID: "payment.webhooks.task.handle_payment"},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact task rename action/sequence: %v", err)
		}
		if taskRename.Target.ID != "payment.webhooks.task.process_payment" || taskRename.Change.Kind != AnalyzeImpactChangeRename || taskRename.Truncated {
			t.Fatalf("AnalyzeImpact task rename envelope = %#v", taskRename)
		}
		renamedAction := assertHasImpact(t, taskRename.Impacts, "transition_action", "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']")
		if renamedAction.Severity != "breaking" || renamedAction.Fixability != "suggested" || len(renamedAction.SuggestedFixes) == 0 {
			t.Fatalf("task rename transition action impact = %#v", renamedAction)
		}
		renamedSequence := assertHasImpact(t, taskRename.Impacts, "sequence_step_action", "payment_webhook_flow")
		if renamedSequence.Severity != "breaking" || renamedSequence.Fixability != "suggested" || len(renamedSequence.SuggestedFixes) == 0 {
			t.Fatalf("task rename sequence impact = %#v", renamedSequence)
		}
		if len(taskRename.Diagnostics) != 0 {
			t.Fatalf("task rename diagnostics = %#v", taskRename.Diagnostics)
		}

		flowRename, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{Object: "node", ID: "order/task/checkout.yaml#build_order"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeRename, NewID: "order/task/checkout.yaml#create_order"},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact task rename flow: %v", err)
		}
		flowImpact := assertHasImpact(t, flowRename.Impacts, "flow_step_task", "order.task.checkout")
		if flowImpact.Severity != "breaking" || flowImpact.Fixability != "suggested" || len(flowImpact.SuggestedFixes) == 0 {
			t.Fatalf("task rename flow impact = %#v", flowImpact)
		}
		flowRender := assertHasImpact(t, flowRename.Impacts, "render_output", "commerce/dag-checkout.md")
		if flowRender.Severity != "info" || flowRender.Fixability == "mechanical" || !strings.Contains(flowRender.Reason, "commerce/dag-checkout.md") {
			t.Fatalf("task rename flow render impact = %#v", flowRender)
		}
		if flowRename.Summary.ByKind["flow_step_task"] != 1 || flowRename.Summary.ByKind["render_output"] == 0 || flowRename.Summary.BySeverity["breaking"] != 1 || flowRename.Summary.BySeverity["info"] == 0 || flowRename.Summary.ByFixability["suggested"] != 1 {
			t.Fatalf("task rename flow summary = %#v impacts=%#v", flowRename.Summary, flowRename.Impacts)
		}

		contract, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{ID: "payment.webhooks.task.process_payment"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeContract, Note: "params changed"},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact task change_contract: %v", err)
		}
		assertHasImpact(t, contract.Impacts, "transition_action", "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']")
		assertHasImpact(t, contract.Impacts, "sequence_step_action", "payment_webhook_flow")
		if contract.Summary.BySeverity["warning"] == 0 || contract.Summary.ByKind["render_output"] == 0 || contract.Summary.ByFixability["manual_review"] == 0 {
			t.Fatalf("task change_contract summary = %#v impacts=%#v", contract.Summary, contract.Impacts)
		}

		transitionImpact, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{Object: "transition", ID: "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeTransitionTarget, NewTo: "order.state.failed"},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact transition target change: %v", err)
		}
		if transitionImpact.Target.ID != "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']" || transitionImpact.Truncated {
			t.Fatalf("AnalyzeImpact transition envelope = %#v", transitionImpact)
		}
		scenarioImpact := assertHasImpact(t, transitionImpact.Impacts, "transition_scenario_step", "payment_webhook_flow")
		if scenarioImpact.Severity != "warning" || scenarioImpact.Fixability != "manual_review" || strings.Join(scenarioImpact.Via, ",") != "scenario_step_transition" {
			t.Fatalf("scenario impact = %#v", scenarioImpact)
		}
		if scenarioImpact.Source == nil || scenarioImpact.Source.File != "views/scenarios/payment_webhook_flow.yaml" || scenarioImpact.Source.Line == 0 || scenarioImpact.Source.Column == 0 {
			t.Fatalf("scenario impact source = %#v", scenarioImpact.Source)
		}
		actionImpact := assertHasImpact(t, transitionImpact.Impacts, "transition_action_task", "payment.webhooks.task.process_payment")
		if actionImpact.Severity != "warning" || actionImpact.Fixability != "manual_review" || strings.Join(actionImpact.Via, ",") != "transition_action" {
			t.Fatalf("action impact = %#v", actionImpact)
		}
		if actionImpact.Source == nil || actionImpact.Source.File != "order/state.yaml" || actionImpact.Source.Line == 0 || actionImpact.Source.Column == 0 {
			t.Fatalf("action impact source = %#v", actionImpact.Source)
		}
		if len(transitionImpact.Diagnostics) != 0 {
			t.Fatalf("transition impact diagnostics = %#v", transitionImpact.Diagnostics)
		}
		transitionRender := assertHasImpact(t, transitionImpact.Impacts, "render_output", "commerce/state-order.md")
		if transitionRender.Severity != "info" || transitionRender.Fixability == "mechanical" || !strings.Contains(transitionRender.RecommendedAction, "brewprint render") {
			t.Fatalf("transition render output impact = %#v", transitionRender)
		}
		if transitionImpact.Summary.ByKind["transition_scenario_step"] != 1 || transitionImpact.Summary.ByKind["transition_action_task"] != 1 || transitionImpact.Summary.ByKind["render_output"] == 0 || transitionImpact.Summary.ByFixability["manual_review"] < 2 {
			t.Fatalf("transition impact summary = %#v", transitionImpact.Summary)
		}

		transitionInvalidTarget, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{Object: "transition", ID: "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeTransitionTarget, NewTo: "order.state.missing"},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact transition invalid target: %v", err)
		}
		resolutionImpact := assertHasImpact(t, transitionInvalidTarget.Impacts, "transition_target_resolution", "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']")
		if resolutionImpact.Severity != "breaking" || resolutionImpact.Fixability != "manual_review" || !strings.Contains(resolutionImpact.Reason, "new_to") {
			t.Fatalf("transition target resolution impact = %#v", resolutionImpact)
		}

		transitionRemove, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{Object: "transition", ID: "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeRemove},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact transition remove: %v", err)
		}
		removedScenarioImpact := assertHasImpact(t, transitionRemove.Impacts, "transition_scenario_step", "payment_webhook_flow")
		if removedScenarioImpact.Severity != "breaking" || removedScenarioImpact.Fixability != "manual_review" {
			t.Fatalf("transition remove scenario impact = %#v", removedScenarioImpact)
		}

		fieldRename, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{Object: "field", ID: "order.model.order.id"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeRename, NewID: "order.model.order.order_id"},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact field rename: %v", err)
		}
		if fieldRename.Target.ID != "order.model.order.id" || len(fieldRename.Impacts) == 0 {
			t.Fatalf("AnalyzeImpact field rename response = %#v", fieldRename)
		}
		fieldRender := assertHasImpact(t, fieldRename.Impacts, "render_output", "_cross/er.md")
		if fieldRender.Severity != "info" || fieldRender.Fixability == "mechanical" || !strings.Contains(fieldRender.Reason, "_cross/er.md") || !strings.Contains(fieldRender.RecommendedAction, "brewprint render") {
			t.Fatalf("field rename render output impact = %#v", fieldRender)
		}
		flowParamImpact := assertHasImpact(t, fieldRename.Impacts, "flow_param_field", "order.task.checkout")
		if flowParamImpact.Severity != "breaking" || flowParamImpact.Fixability != "unknown" || strings.Join(flowParamImpact.Via, ",") != "flow_param_field_resolution" {
			t.Fatalf("field rename flow param impact = %#v", flowParamImpact)
		}
		if fieldRename.Summary.BySeverity["breaking"] == 0 || fieldRename.Summary.BySeverity["info"] == 0 || fieldRename.Summary.ByFixability["unknown"] == 0 || fieldRename.Summary.ByKind["field_consumer"] == 0 || fieldRename.Summary.ByKind["render_output"] == 0 || fieldRename.Summary.ByKind["flow_param_field"] == 0 {
			t.Fatalf("AnalyzeImpact field rename summary = %#v impacts=%#v", fieldRename.Summary, fieldRename.Impacts)
		}
		foundFileOnlySource := false
		for _, impact := range fieldRename.Impacts {
			if impact.Source != nil && impact.Source.File != "" && (impact.Source.Line == 0 || impact.Source.Column == 0) {
				foundFileOnlySource = true
				break
			}
		}
		if !foundFileOnlySource {
			t.Fatalf("AnalyzeImpact field rename expected source.file without line/column: %#v", fieldRename.Impacts)
		}
		foundSourceDiagnostic := false
		for _, diagnostic := range fieldRename.Diagnostics {
			if diagnostic.Code == "source_location_unavailable" {
				foundSourceDiagnostic = true
				break
			}
		}
		if !foundSourceDiagnostic {
			t.Fatalf("AnalyzeImpact field rename diagnostics = %#v", fieldRename.Diagnostics)
		}

		fieldAliasRename, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{Kind: "model_field", ID: "order.model.order.id"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeRename, NewID: "order.model.order.order_id"},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact model_field alias rename: %v", err)
		}
		if fieldAliasRename.Target.Object != "field" || fieldAliasRename.Target.ID != "order.model.order.id" || !hasString(fieldAliasRename.Coverage.Analyzed, "flow_param_field_resolution") {
			t.Fatalf("AnalyzeImpact model_field alias response = target %#v coverage %#v", fieldAliasRename.Target, fieldAliasRename.Coverage)
		}

		fieldTypeChange, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{Object: "field", ID: "order.model.order", LocalID: "id"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeType, NewType: "int"},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact field change_type: %v", err)
		}
		fieldTypeRender := assertHasImpact(t, fieldTypeChange.Impacts, "render_output", "_cross/er.md")
		if fieldTypeRender.Severity != "info" || fieldTypeRender.Fixability == "mechanical" || !strings.Contains(fieldTypeRender.Reason, "_cross/er.md") {
			t.Fatalf("field change_type render output impact = %#v", fieldTypeRender)
		}
		if fieldTypeChange.Summary.ByKind["render_output"] == 0 || fieldTypeChange.Summary.BySeverity["info"] == 0 {
			t.Fatalf("AnalyzeImpact field change_type summary = %#v impacts=%#v", fieldTypeChange.Summary, fieldTypeChange.Impacts)
		}

		modelRename, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{ID: "order.model.order"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeRename, NewID: "order.model.order_v2"},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact model rename: %v", err)
		}
		modelRender := assertHasImpact(t, modelRename.Impacts, "render_output", "commerce/dag-checkout.md")
		if modelRender.Severity != "info" || modelRender.Fixability == "mechanical" || !strings.Contains(modelRender.Reason, "commerce/dag-checkout.md") {
			t.Fatalf("model rename render output impact = %#v", modelRender)
		}
		if modelRename.Summary.ByKind["render_output"] == 0 || modelRename.Summary.BySeverity["info"] == 0 {
			t.Fatalf("AnalyzeImpact model rename summary = %#v impacts=%#v", modelRename.Summary, modelRename.Impacts)
		}

		addTaskCollision, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{ID: "order.task.checkout"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeAdd, AddedID: "order.task.checkout"},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact add task collision: %v", err)
		}
		addTaskImpact := assertHasImpact(t, addTaskCollision.Impacts, "name_collision", "order.task.checkout")
		if addTaskImpact.Severity != "breaking" || addTaskImpact.Fixability != "manual_review" || strings.Join(addTaskImpact.Via, ",") != "name_collision" {
			t.Fatalf("AnalyzeImpact add task collision impact = %#v", addTaskImpact)
		}
		if addTaskCollision.Summary.ByKind["name_collision"] != 1 || addTaskCollision.Summary.BySeverity["breaking"] != 1 {
			t.Fatalf("AnalyzeImpact add task collision summary = %#v impacts=%#v", addTaskCollision.Summary, addTaskCollision.Impacts)
		}

		addFieldCollision, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{ID: "order.task.checkout"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeAdd, AddedID: "order.model.order.id"},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact add field collision: %v", err)
		}
		addFieldImpact := assertHasImpact(t, addFieldCollision.Impacts, "name_collision", "order.model.order.id")
		if addFieldImpact.Object.Object != "field" || addFieldImpact.Severity != "breaking" || addFieldImpact.Fixability != "manual_review" {
			t.Fatalf("AnalyzeImpact add field collision impact = %#v", addFieldImpact)
		}

		addNoCollision, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{ID: "order.task.checkout"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeAdd, AddedID: "order.task.missing_future_task"},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact add no collision: %v", err)
		}
		if len(addNoCollision.Impacts) != 0 || addNoCollision.Summary.BySeverity["breaking"] != 0 || addNoCollision.Summary.ByFixability["manual_review"] != 0 {
			t.Fatalf("AnalyzeImpact add no collision = %#v", addNoCollision)
		}
		if strings.Join(addNoCollision.Coverage.Analyzed, ",") != "name_collision" || !hasString(addNoCollision.Coverage.NotAnalyzed, "type_resolution") || !hasString(addNoCollision.Coverage.NotAnalyzed, "writer_coverage") {
			t.Fatalf("AnalyzeImpact add coverage = %#v", addNoCollision.Coverage)
		}

		unsupported, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{Object: "view", Kind: "api_table", ID: "ec_api"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeRemove},
		})
		if err != nil {
			t.Fatalf("AnalyzeImpact unsupported selector should not error: %v", err)
		}
		if len(unsupported.Impacts) != 0 || len(unsupported.Diagnostics) != 1 || unsupported.Diagnostics[0].Code != "unsupported_selector" {
			t.Fatalf("AnalyzeImpact unsupported response = %#v", unsupported)
		}

		if _, err := service.AnalyzeImpact(AnalyzeImpactRequest{
			Selector: Selector{ID: "auth.task.login"},
			Change:   AnalyzeImpactChange{Kind: AnalyzeImpactChangeRename},
		}); err == nil || !strings.Contains(err.Error(), "invalid change payload") {
			t.Fatalf("AnalyzeImpact missing new_id error = %v", err)
		}

		var invalidPayload AnalyzeImpactRequest
		if err := json.Unmarshal([]byte(`{"selector":{"id":"auth.task.login"},"change":{"kind":"remove","new_id":"auth.task.sign_in"}}`), &invalidPayload); err != nil {
			t.Fatalf("unmarshal invalid analyze impact payload: %v", err)
		}
		if _, err := service.AnalyzeImpact(invalidPayload); err == nil || !strings.Contains(err.Error(), "invalid change payload") {
			t.Fatalf("AnalyzeImpact extra payload error = %v", err)
		}

		fakeImpacts := []ImpactEntry{
			{Kind: "field_consumer", Severity: "breaking", Fixability: "manual_review"},
			{Kind: "render_output", Severity: "info", Fixability: "unknown"},
		}
		truncated, didTruncate, reasons := truncateImpacts(fakeImpacts, 1)
		assignImpactIDs(truncated)
		if !didTruncate || len(reasons) != 1 || reasons[0] != "max_impacts" || len(truncated) != 1 || truncated[0].ID != "impact-001" {
			t.Fatalf("AnalyzeImpact truncation helpers = impacts %#v truncated %v reasons %#v", truncated, didTruncate, reasons)
		}
		summary := summarizeImpacts(fakeImpacts)
		if summary.BySeverity["breaking"] != 1 || summary.BySeverity["info"] != 1 || summary.ByFixability["manual_review"] != 1 || summary.ByKind["field_consumer"] != 1 {
			t.Fatalf("AnalyzeImpact summary helper = %#v", summary)
		}
	})

	t.Run("M11Selectors", func(t *testing.T) {
		buildOrderSelector := Selector{Object: "node", ID: "order/task/checkout.yaml#build_order"}
		buildOrder, err := service.GetSignature(GetSignatureRequest{Selector: buildOrderSelector})
		if err != nil {
			t.Fatalf("GetSignature private sub task: %v", err)
		}
		if buildOrder.Object.ID != "order/task/checkout.yaml#build_order" || buildOrder.Object.Kind != "task" || buildOrder.Object.LocalID != "build_order" || buildOrder.Object.QualifiedID != "" {
			t.Fatalf("private sub task object = %#v", buildOrder.Object)
		}
		ret := buildOrder.Signature["returns"].(*ReturnSignature)
		if ret.Name != "draft_order" || ret.Asset == nil || ret.Asset.ID != "order/task/checkout.yaml#build_order#draft_order" {
			t.Fatalf("private sub task returns = %#v", ret)
		}

		fork, err := service.GetSignature(GetSignatureRequest{Selector: Selector{Object: "node", Kind: "fork", File: "order/task/checkout.yaml", LocalID: "parallel_processing"}})
		if err != nil {
			t.Fatalf("GetSignature private fork: %v", err)
		}
		if fork.Object.ID != "order/task/checkout.yaml#parallel_processing" || fork.Object.Kind != "fork" {
			t.Fatalf("private fork object = %#v", fork.Object)
		}

		assetSelector := Selector{Object: "asset", ID: "order/task/checkout.yaml#build_order#draft_order"}
		asset, err := service.GetSignature(GetSignatureRequest{Selector: assetSelector})
		if err != nil {
			t.Fatalf("GetSignature asset: %v", err)
		}
		if asset.Object.ID != "order/task/checkout.yaml#build_order#draft_order" || asset.Object.Object != "asset" || asset.Signature["producer"] != "order/task/checkout.yaml#build_order" || asset.Signature["model"] != "order.model.order" || asset.Signature["scope_file"] != "order/task/checkout.yaml" {
			t.Fatalf("asset signature = object %#v signature %#v", asset.Object, asset.Signature)
		}

		assetRefs, err := service.GetReferences(GetReferencesRequest{Selector: assetSelector, Direction: "out", Kinds: []string{"consumes_asset"}})
		if err != nil {
			t.Fatalf("GetReferences asset consumers: %v", err)
		}
		assertHasReference(t, assetRefs.References, "consumes_asset", "out", "order/task/checkout.yaml#build_order#draft_order", "order/task/checkout.yaml#reserve_inventory")
		assertHasReference(t, assetRefs.References, "consumes_asset", "out", "order/task/checkout.yaml#build_order#draft_order", "order/task/checkout.yaml#notify_payment_gateway")

		buildOrderRefs, err := service.GetReferences(GetReferencesRequest{Selector: buildOrderSelector, Direction: "out", Kinds: []string{"produces_asset"}})
		if err != nil {
			t.Fatalf("GetReferences private sub task: %v", err)
		}
		assertHasReference(t, buildOrderRefs.References, "produces_asset", "out", "order/task/checkout.yaml#build_order", "order/task/checkout.yaml#build_order#draft_order")

		fieldAlias, err := service.GetSignature(GetSignatureRequest{Selector: Selector{Object: "field", Kind: "model_field", ID: "order.model.order", LocalID: "id"}})
		if err != nil {
			t.Fatalf("GetSignature model_field alias: %v", err)
		}
		if fieldAlias.Object.Object != "field" || fieldAlias.Object.Kind != "field" || fieldAlias.Object.ID != "order.model.order.id" {
			t.Fatalf("model_field alias object = %#v", fieldAlias.Object)
		}

		fullFieldAlias, err := service.GetSignature(GetSignatureRequest{Selector: Selector{Object: "field", Kind: "model_field", ID: "order.model.order.id"}})
		if err != nil {
			t.Fatalf("GetSignature full model_field alias: %v", err)
		}
		if fullFieldAlias.Object.Object != "field" || fullFieldAlias.Object.ID != "order.model.order.id" {
			t.Fatalf("full model_field alias object = %#v", fullFieldAlias.Object)
		}

		inspected, err := service.Inspect(InspectRequest{Selector: buildOrderSelector})
		if err != nil {
			t.Fatalf("Inspect private sub task: %v", err)
		}
		if inspected.Object.ID != "order/task/checkout.yaml#build_order" || inspected.Source["file"] != "order/task/checkout.yaml" || len(inspected.References) == 0 {
			t.Fatalf("private sub task inspect = %#v", inspected)
		}
	})

	t.Run("ListEndpoints", func(t *testing.T) {
		endpoints, err := service.ListEndpoints(ListEndpointsRequest{APITableID: "ec_api"})
		if err != nil {
			t.Fatalf("ListEndpoints ec_api: %v", err)
		}
		if len(endpoints.Tables) != 1 {
			t.Fatalf("endpoint tables len = %d, want 1: %#v", len(endpoints.Tables), endpoints.Tables)
		}
		table := endpoints.Tables[0]
		if table.ID != "ec_api" || table.HTTPRootPath != "/api" {
			t.Fatalf("endpoint table = %#v", table)
		}
		if len(table.Sections) != 5 {
			t.Fatalf("endpoint sections len = %d, want 5: %#v", len(table.Sections), table.Sections)
		}
		assertHasEndpoint(t, table.Sections, "auth", "auth.task.login", "POST", "/api/login")
		assertHasEndpoint(t, table.Sections, "catalog", "catalog.task.get_items", "GET", "/api/catalog_items")
		assertHasEndpoint(t, table.Sections, "payment.webhooks", "payment.webhooks.task.process_payment", "POST", "/api/stripe")
	})

	t.Run("Inspect", func(t *testing.T) {
		login, err := service.Inspect(InspectRequest{Selector: Selector{ID: "auth.task.login"}})
		if err != nil {
			t.Fatalf("Inspect login: %v", err)
		}
		assets := login.Members["assets"].([]AssetRef)
		if len(assets) != 1 || assets[0].Name != "auth_token" {
			t.Fatalf("login assets = %#v", assets)
		}
		if len(login.References) == 0 {
			t.Fatalf("login references is empty")
		}

		checkout, err := service.Inspect(InspectRequest{Selector: Selector{ID: "order.task.checkout"}})
		if err != nil {
			t.Fatalf("Inspect checkout: %v", err)
		}
		if _, ok := checkout.Members["flow"]; !ok {
			t.Fatalf("checkout flow member missing: %#v", checkout.Members)
		}
		if _, ok := checkout.Members["sub_tasks"]; !ok {
			t.Fatalf("checkout sub_tasks member missing: %#v", checkout.Members)
		}
		checkoutActions := checkout.Members["action_transitions"].([]TransitionRef)
		if len(checkoutActions) != 1 {
			t.Fatalf("checkout action transitions = %#v", checkoutActions)
		}
		if checkoutActions[0].ID != "order/state.yaml#checkout_screen:checkout_submitted" || checkoutActions[0].Action != "order.task.checkout" {
			t.Fatalf("checkout action transition = %#v", checkoutActions[0])
		}

		payment, err := service.Inspect(InspectRequest{Selector: Selector{ID: "payment.webhooks.task.process_payment"}})
		if err != nil {
			t.Fatalf("Inspect payment process: %v", err)
		}
		paymentActions := payment.Members["action_transitions"].([]TransitionRef)
		if len(paymentActions) != 1 {
			t.Fatalf("payment action transitions = %#v", paymentActions)
		}
		if paymentActions[0].ID != "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']" {
			t.Fatalf("payment action transition = %#v", paymentActions[0])
		}

		model, err := service.Inspect(InspectRequest{Selector: Selector{ID: "auth.model.login_form"}})
		if err != nil {
			t.Fatalf("Inspect model: %v", err)
		}
		if _, ok := model.Members["fields"]; !ok {
			t.Fatalf("model fields member missing: %#v", model.Members)
		}

		store, err := service.Inspect(InspectRequest{Selector: Selector{ID: "auth.store.user_db"}})
		if err != nil {
			t.Fatalf("Inspect store: %v", err)
		}
		if _, ok := store.Members["model"]; !ok {
			t.Fatalf("store model member missing: %#v", store.Members)
		}
		assertHasReference(t, store.References, "reads", "in", "auth.task.login", "auth.store.user_db")

		state, err := service.Inspect(InspectRequest{Selector: Selector{ID: "order.state.processing"}})
		if err != nil {
			t.Fatalf("Inspect processing state: %v", err)
		}
		incoming := state.Members["incoming_transitions"].([]TransitionRef)
		outgoing := state.Members["outgoing_transitions"].([]TransitionRef)
		if len(incoming) != 1 || incoming[0].ID != "order/state.yaml#checkout_screen:checkout_submitted" {
			t.Fatalf("processing incoming transitions = %#v", incoming)
		}
		if len(outgoing) != 2 {
			t.Fatalf("processing outgoing transitions = %#v", outgoing)
		}
		if outgoing[0].ID != "order/state.yaml#processing:payment_webhook_received[payload.status == 'failed']" || outgoing[1].ID != "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']" {
			t.Fatalf("processing outgoing transitions = %#v", outgoing)
		}

		event, err := service.Inspect(InspectRequest{Selector: Selector{ID: "order.event.payment_webhook_received"}})
		if err != nil {
			t.Fatalf("Inspect payment event: %v", err)
		}
		triggering := event.Members["triggering_transitions"].([]TransitionRef)
		if len(triggering) != 2 {
			t.Fatalf("payment event triggering transitions = %#v", triggering)
		}
		hints := event.Members["sequence_hints"].(map[string]any)
		if hints["participant"] != "Actor" || hints["actor"] != "stripe" || hints["message_label_source"] != "METHOD path" {
			t.Fatalf("payment event sequence hints = %#v", hints)
		}

		checkoutScenario, err := service.Inspect(InspectRequest{Selector: Selector{Object: "view", Kind: "sequence_diagram", ID: "checkout_flow"}})
		if err != nil {
			t.Fatalf("Inspect checkout scenario: %v", err)
		}
		checkoutSteps := checkoutScenario.Members["steps"].([]ScenarioStepRef)
		if len(checkoutSteps) != 2 {
			t.Fatalf("checkout scenario steps = %#v", checkoutSteps)
		}
		if checkoutSteps[0].Index != 1 || checkoutSteps[0].Transition.ID != "order/state.yaml#cart:view_checkout" || checkoutSteps[0].Action != nil {
			t.Fatalf("checkout step 1 = %#v", checkoutSteps[0])
		}
		if checkoutSteps[1].Index != 2 || checkoutSteps[1].Transition.ID != "order/state.yaml#checkout_screen:checkout_submitted" || checkoutSteps[1].Action == nil || *checkoutSteps[1].Action != "order.task.checkout" {
			t.Fatalf("checkout step 2 = %#v", checkoutSteps[1])
		}
		assertHasReference(t, checkoutScenario.References, "scenario_state_file", "out", "checkout_flow", "order/state.yaml")

		paymentScenario, err := service.Inspect(InspectRequest{Selector: Selector{Object: "view", Kind: "sequence_diagram", ID: "payment_webhook_flow"}})
		if err != nil {
			t.Fatalf("Inspect payment scenario: %v", err)
		}
		paymentSteps := paymentScenario.Members["steps"].([]ScenarioStepRef)
		if len(paymentSteps) != 1 {
			t.Fatalf("payment scenario steps = %#v", paymentSteps)
		}
		if paymentSteps[0].Guard != "payload.status == 'succeeded'" || !paymentSteps[0].GuardExactMatch || paymentSteps[0].Transition.ID != "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']" || paymentSteps[0].Action == nil || *paymentSteps[0].Action != "payment.webhooks.task.process_payment" {
			t.Fatalf("payment scenario step = %#v", paymentSteps[0])
		}

		transition, err := service.Inspect(InspectRequest{Selector: Selector{Object: "transition", ID: "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"}})
		if err != nil {
			t.Fatalf("Inspect transition: %v", err)
		}
		if transition.Signature["guard"] != "payload.status == 'succeeded'" || transition.Signature["action"] != "payment.webhooks.task.process_payment" {
			t.Fatalf("transition signature = %#v", transition.Signature)
		}
		fromState := transition.Members["from_state"].(ObjectRef)
		eventRef := transition.Members["event"].(ObjectRef)
		toState := transition.Members["to_state"].(ObjectRef)
		actionTask := transition.Members["action_task"].(ObjectRef)
		if fromState.ID != "order.state.processing" || eventRef.ID != "order.event.payment_webhook_received" || toState.ID != "order.state.confirmed" || actionTask.ID != "payment.webhooks.task.process_payment" {
			t.Fatalf("transition members = %#v", transition.Members)
		}
		assertHasReference(t, transition.References, "scenario_step_transition", "in", "scenario_step:payment_webhook_flow:1", "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']")

		field, err := service.Inspect(InspectRequest{Selector: Selector{Object: "field", ID: "order.model.order", LocalID: "id"}})
		if err != nil {
			t.Fatalf("Inspect field: %v", err)
		}
		if field.Signature["name"] != "id" || field.Signature["type"] != "str" || field.Signature["pk"] != true {
			t.Fatalf("field signature = %#v", field.Signature)
		}
		fieldModel := field.Members["model"].(ObjectRef)
		if fieldModel.ID != "order.model.order" || field.Members["type"] != "str" {
			t.Fatalf("field members = %#v", field.Members)
		}
		assertHasReference(t, field.References, "field_type", "out", "order.model.order.id", "str")
		assertHasReference(t, field.References, "field_fk", "in", "order.model.order_item.order_id", "order.model.order.id")
		assertHasReference(t, field.References, "field_fk", "in", "payment.model.payment_event.order_id", "order.model.order.id")

		apiTable, err := service.Inspect(InspectRequest{Selector: Selector{Object: "view", Kind: "api_table", ID: "ec_api"}})
		if err != nil {
			t.Fatalf("Inspect API table: %v", err)
		}
		if apiTable.Object.Object != "view" || apiTable.Object.Kind != "api_table" || apiTable.Object.ID != "ec_api" {
			t.Fatalf("API table object = %#v", apiTable.Object)
		}
		if apiTable.Signature["http_root_path"] != "/api" {
			t.Fatalf("API table signature = %#v", apiTable.Signature)
		}
		apiModules := apiTable.Members["modules"].([]apiInspectModule)
		if len(apiModules) != 5 {
			t.Fatalf("API table modules = %#v", apiModules)
		}
		apiEndpoints := apiTable.Members["collected_endpoints"].([]apiInspectEndpoint)
		assertHasAPIInspectEndpoint(t, apiEndpoints, "auth.task.login", "POST", "/api/login")
		assertHasAPIInspectEndpoint(t, apiEndpoints, "payment.webhooks.task.process_payment", "POST", "/api/stripe")

		erView, err := service.Inspect(InspectRequest{Selector: Selector{Object: "view", Kind: "er_diagram", ID: "ec_er"}})
		if err != nil {
			t.Fatalf("Inspect ER view: %v", err)
		}
		if erView.Object.Object != "view" || erView.Object.Kind != "er_diagram" || erView.Object.ID != "ec_er" {
			t.Fatalf("ER view object = %#v", erView.Object)
		}
		includedModels := erView.Members["included_models"].([]ObjectRef)
		assertHasObjectRef(t, includedModels, "auth.model.credential")
		assertHasObjectRef(t, includedModels, "order.model.order")
		relations := erView.Members["fk_relations"].([]erInspectFKRelation)
		assertHasERInspectRelation(t, relations, "order.model.order", "user_id", "auth.model.credential", "username")
	})
}

func TestPrivateSubNodeAssetIdentityBoundaryQueries(t *testing.T) {
	checkoutFile := privateReturningHelperQueryFile("shop/task/checkout.yaml", "checkout")
	refundFile := privateReturningHelperQueryFile("shop/task/refund.yaml", "refund")
	checkoutFile.NodeFile.Models = nil
	refundFile.NodeFile.Models = nil
	project, diagnostics := resolve.Build(&rawyaml.Project{Files: []rawyaml.File{{
		ID:   "shop/model/receipt.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{Models: []rawyaml.Model{{
			ID:     "receipt",
			Kind:   "struct",
			Fields: []rawyaml.ModelField{{Name: "id", Type: "str"}},
		}}},
	}, checkoutFile, refundFile}})
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	service := NewService(project)

	if _, err := service.GetSignature(GetSignatureRequest{Selector: Selector{ID: "shop.task.helper"}}); err == nil {
		t.Fatalf("public-shaped private sub task selector unexpectedly resolved")
	}
	if _, err := service.GetSignature(GetSignatureRequest{Selector: Selector{Object: "asset", ID: "shop.task.helper#result"}}); err == nil {
		t.Fatalf("public-shaped private asset selector unexpectedly resolved")
	}

	checkoutAssetID := "shop/task/checkout.yaml#helper#result"
	refundAssetID := "shop/task/refund.yaml#helper#result"
	checkoutAsset, err := service.GetSignature(GetSignatureRequest{Selector: Selector{Object: "asset", ID: checkoutAssetID}})
	if err != nil {
		t.Fatalf("GetSignature checkout asset: %v", err)
	}
	if checkoutAsset.Object.ID != checkoutAssetID || checkoutAsset.Signature["producer"] != "shop/task/checkout.yaml#helper" {
		t.Fatalf("checkout asset signature = object %#v signature %#v", checkoutAsset.Object, checkoutAsset.Signature)
	}
	refundAsset, err := service.GetSignature(GetSignatureRequest{Selector: Selector{Object: "asset", ID: refundAssetID}})
	if err != nil {
		t.Fatalf("GetSignature refund asset: %v", err)
	}
	if refundAsset.Object.ID != refundAssetID || refundAsset.Signature["producer"] != "shop/task/refund.yaml#helper" {
		t.Fatalf("refund asset signature = object %#v signature %#v", refundAsset.Object, refundAsset.Signature)
	}

	checkoutRefs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{Object: "asset", ID: checkoutAssetID}, Direction: "out", Kinds: []string{"consumes_asset"}})
	if err != nil {
		t.Fatalf("GetReferences checkout asset: %v", err)
	}
	assertHasReference(t, checkoutRefs.References, "consumes_asset", "out", checkoutAssetID, "shop/task/checkout.yaml#consume")
	for _, ref := range checkoutRefs.References {
		if ref.To.ID == "shop/task/refund.yaml#consume" {
			t.Fatalf("checkout asset references leaked to refund consumer: %#v", checkoutRefs.References)
		}
	}

	refundSource, err := service.GetSource(GetSourceRequest{Selector: Selector{Object: "asset", ID: refundAssetID}})
	if err != nil {
		t.Fatalf("GetSource refund asset: %v", err)
	}
	if refundSource.Object.ID != refundAssetID || refundSource.Source.File != "shop/task/refund.yaml" {
		t.Fatalf("refund asset source = %#v", refundSource)
	}
}

func TestTaskFilePrivateHelperModelsDoNotLeakAsPublicQueryTargets(t *testing.T) {
	project, diagnostics := resolve.Build(&rawyaml.Project{Files: []rawyaml.File{{
		ID:   "auth/task/login.yaml",
		Kind: rawyaml.FileKindNode,
		NodeFile: &rawyaml.NodeFile{
			Tasks: []rawyaml.Task{{
				ID:      "login",
				Main:    true,
				Returns: &rawyaml.Return{Name: "token", Model: "login_token"},
			}},
			Models: []rawyaml.Model{
				{ID: "login_form", Kind: "struct", Fields: []rawyaml.ModelField{{Name: "username", Type: "str"}}},
				{ID: "login_token", Kind: "struct", Fields: []rawyaml.ModelField{{Name: "access_token", Type: "str"}, {Name: "form", Type: "login_form"}}},
			},
		},
	}}})
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics = %#v, want none", diagnostics)
	}
	service := NewService(project)

	signature, err := service.GetSignature(GetSignatureRequest{Selector: Selector{ID: "auth.task.login"}})
	if err != nil {
		t.Fatalf("GetSignature task: %v", err)
	}
	params := signature.Signature["params"].([]ParamSignature)
	if len(params) != 0 {
		t.Fatalf("params = %#v, want none", params)
	}
	ret := signature.Signature["returns"].(*ReturnSignature)
	if ret.Model != "login_token" || strings.Contains(ret.Model, "#") || ret.Asset.Model != "login_token" || strings.Contains(ret.Asset.Model, "#") {
		t.Fatalf("private helper return model leaked synthetic id: %#v", ret)
	}

	refs, err := service.GetReferences(GetReferencesRequest{Selector: Selector{ID: "auth.task.login"}, Direction: "out"})
	if err != nil {
		t.Fatalf("GetReferences task: %v", err)
	}
	for _, ref := range refs.References {
		if ref.Kind != "param_model" && ref.Kind != "return_model" {
			continue
		}
		if strings.Contains(ref.To.ID, "#") || strings.Contains(ref.To.QualifiedID, "#") || ref.To.Kind == "model" {
			t.Fatalf("private helper reference leaked as public model target: %#v", ref)
		}
	}

	listed, err := service.ListObjects(ListObjectsRequest{Object: "node", Kind: "model"})
	if err != nil {
		t.Fatalf("ListObjects models: %v", err)
	}
	for _, object := range listed.Objects {
		if object.ID == "auth/task/login.yaml#login_form" || object.ID == "auth/task/login.yaml#login_token" || object.Label == "login_form" || object.Label == "login_token" {
			t.Fatalf("private helper model leaked into list_objects: %#v", listed.Objects)
		}
	}
}

func TestResolvedTransitionIndexesUC001(t *testing.T) {
	project := loadUC001Project(t)

	key := semantic.TransitionKey{
		StateFile: semantic.FileID("order/state.yaml"),
		FromState: semantic.QualifiedID("order.state.processing"),
		Event:     semantic.QualifiedID("order.event.payment_webhook_received"),
		Guard:     "payload.status == 'succeeded'",
	}
	ref, ok := project.TransitionsByStateEventGuard[key]
	if !ok {
		t.Fatalf("transition index missing key %#v", key)
	}
	if ref.Transition.ToState != semantic.QualifiedID("order.state.confirmed") {
		t.Fatalf("indexed transition to = %s, want order.state.confirmed", ref.Transition.ToState)
	}

	eventKey := semantic.TransitionEventKey{
		StateFile: semantic.FileID("order/state.yaml"),
		FromState: semantic.QualifiedID("order.state.processing"),
		Event:     semantic.QualifiedID("order.event.payment_webhook_received"),
	}
	if got := len(project.TransitionsByStateEvent[eventKey]); got != 2 {
		t.Fatalf("transition event candidates len = %d, want 2", got)
	}

	actions := project.ActionsByTask[semantic.QualifiedID("payment.webhooks.task.process_payment")]
	if len(actions) != 1 {
		t.Fatalf("actionsByTask len = %d, want 1: %#v", len(actions), actions)
	}
	if actions[0].Transition.Guard != "payload.status == 'succeeded'" {
		t.Fatalf("action transition guard = %q", actions[0].Transition.Guard)
	}
}

func newUC001Service(t *testing.T) *Service {
	t.Helper()
	return NewService(loadUC001Project(t))
}

func privateReturningHelperQueryFile(fileID string, mainID string) rawyaml.File {
	return rawyaml.File{
		ID:      fileID,
		Kind:    rawyaml.FileKindNode,
		Content: privateReturningHelperQueryContent(mainID),
		NodeFile: &rawyaml.NodeFile{
			Models: []rawyaml.Model{{
				ID:     "receipt",
				Kind:   "struct",
				Fields: []rawyaml.ModelField{{Name: "id", Type: "str"}},
			}},
			Tasks: []rawyaml.Task{
				{ID: mainID, Main: true},
				{ID: "helper", Returns: &rawyaml.Return{Name: "result", Model: "receipt"}},
				{ID: "consume", Params: []rawyaml.Param{{Name: "input", Model: "receipt"}}},
			},
			Flow: []rawyaml.FlowEntry{
				{Step: "helper"},
				{Step: "consume", Params: map[string]string{"input": "helper"}},
			},
		},
	}
}

func privateReturningHelperQueryContent(mainID string) string {
	return "nodes:\n" +
		"  - id: " + mainID + "\n" +
		"    type: task\n" +
		"    main: true\n" +
		"  - id: helper\n" +
		"    type: task\n" +
		"    returns:\n" +
		"      name: result\n" +
		"      model: receipt\n" +
		"  - id: consume\n" +
		"    type: task\n" +
		"    params:\n" +
		"      - name: input\n" +
		"        model: receipt\n" +
		"flow:\n" +
		"  - step: helper\n" +
		"  - step: consume\n" +
		"    params:\n" +
		"      input: helper\n"
}

func loadUC001Project(t *testing.T) *semantic.Project {
	t.Helper()
	yamlRoot := filepath.FromSlash("../../docs/uc/001-ec-checkout-flow/yaml")
	loader := source.Loader{}
	raw, err := loader.Load(yamlRoot)
	if err != nil {
		t.Fatalf("load yaml root: %v", err)
	}
	project, diagnostics := resolve.Build(raw)
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == semantic.SeverityError {
			t.Fatalf("semantic diagnostic: %s: %s", diagnostic.FileID, diagnostic.Message)
		}
	}
	return project
}

func hasString(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func assertHasObject(t *testing.T, refs []ObjectRef, object, kind, id string) ObjectRef {
	t.Helper()
	for _, ref := range refs {
		if ref.Object == object && ref.Kind == kind && ref.ID == id {
			return ref
		}
	}
	t.Fatalf("object not found object=%s kind=%s id=%s in %#v", object, kind, id, refs)
	return ObjectRef{}
}

func assertHasEndpoint(t *testing.T, sections []EndpointSection, module, task, method, path string) {
	t.Helper()
	for _, section := range sections {
		if section.Module != module {
			continue
		}
		for _, endpoint := range section.Endpoints {
			if endpoint.Task == task && endpoint.Method == method && endpoint.Path == path {
				return
			}
		}
	}
	t.Fatalf("endpoint not found module=%s task=%s method=%s path=%s in %#v", module, task, method, path, sections)
}

func assertHasReference(t *testing.T, refs []Reference, kind, direction, fromID, toID string) Reference {
	t.Helper()
	for _, ref := range refs {
		if ref.Kind != kind || ref.Direction != direction {
			continue
		}
		if fromID != "" && ref.From.ID != fromID {
			continue
		}
		if toID != "" && ref.To.ID != toID {
			continue
		}
		return ref
	}
	t.Fatalf("reference not found kind=%s direction=%s from=%s to=%s in %#v", kind, direction, fromID, toID, refs)
	return Reference{}
}

func assertHasImpact(t *testing.T, impacts []ImpactEntry, kind, objectID string) ImpactEntry {
	t.Helper()
	for _, impact := range impacts {
		if impact.Kind == kind && impact.Object.ID == objectID {
			return impact
		}
	}
	t.Fatalf("impact not found kind=%s object=%s in %#v", kind, objectID, impacts)
	return ImpactEntry{}
}

func assertHasReferenceTreeNode(t *testing.T, nodes []ReferenceTreeNode, id string, depth int, via []string) ReferenceTreeNode {
	t.Helper()
	for _, node := range nodes {
		if node.Object.ID != id || node.Depth != depth || strings.Join(node.Via, ",") != strings.Join(via, ",") {
			continue
		}
		return node
	}
	t.Fatalf("reference tree node not found id=%s depth=%d via=%v in %#v", id, depth, via, nodes)
	return ReferenceTreeNode{}
}

func assertHasReferenceTreeEdge(t *testing.T, edges []ReferenceTreeEdge, kind, direction, fromID, toID string, depth int) ReferenceTreeEdge {
	t.Helper()
	for _, edge := range edges {
		if edge.Kind == kind && edge.Direction == direction && edge.From.ID == fromID && edge.To.ID == toID && edge.Depth == depth {
			return edge
		}
	}
	t.Fatalf("reference tree edge not found kind=%s direction=%s from=%s to=%s depth=%d in %#v", kind, direction, fromID, toID, depth, edges)
	return ReferenceTreeEdge{}
}

func assertHasAPIInspectEndpoint(t *testing.T, endpoints []apiInspectEndpoint, task, method, path string) {
	t.Helper()
	for _, endpoint := range endpoints {
		if endpoint.Task == task && endpoint.Method == method && endpoint.Path == path {
			return
		}
	}
	t.Fatalf("API inspect endpoint not found task=%s method=%s path=%s in %#v", task, method, path, endpoints)
}

func assertHasObjectRef(t *testing.T, refs []ObjectRef, id string) {
	t.Helper()
	for _, ref := range refs {
		if ref.ID == id {
			return
		}
	}
	t.Fatalf("object ref not found id=%s in %#v", id, refs)
}

func assertHasERInspectRelation(t *testing.T, relations []erInspectFKRelation, fromModel, fromField, toModel, toField string) {
	t.Helper()
	for _, relation := range relations {
		if relation.FromModel == fromModel && relation.FromField == fromField && relation.ToModel == toModel && relation.ToField == toField {
			return
		}
	}
	t.Fatalf("ER inspect relation not found from=%s.%s to=%s.%s in %#v", fromModel, fromField, toModel, toField, relations)
}
