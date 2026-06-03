package dag

import (
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func TestAssetTypeHint(t *testing.T) {
	cases := []struct {
		name string
		ref  *semantic.TypeRef
		want string
	}{
		{"nil ref", nil, ""},
		{"primitive str", &semantic.TypeRef{Kind: semantic.TypeRefPrimitive, Name: "str"}, "str"},
		{"primitive any", &semantic.TypeRef{Kind: semantic.TypeRefPrimitive, Name: "any"}, "any"},
		{"primitive int", &semantic.TypeRef{Kind: semantic.TypeRefPrimitive, Name: "int"}, "int"},
		{"named model bare", &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Name: "login_form", Model: "auth.model.login_form"}, "login_form"},
		{"named model full qid", &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Name: "payment.model.payment_event", Model: "payment.model.payment_event"}, "payment_event"},
		{"named list model", &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Name: "cart_item_list", Model: "cart.model.cart_item_list"}, "cart_item_list"},
		{"inline list", &semantic.TypeRef{Kind: semantic.TypeRefList, Name: "list"}, "list"},
		{"inline dict", &semantic.TypeRef{Kind: semantic.TypeRefDict, Name: "dict"}, "dict"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := assetTypeHint(tc.ref)
			if got != tc.want {
				t.Errorf("assetTypeHint = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestNamedModelLocalID(t *testing.T) {
	cases := []struct {
		name string
		ref  *semantic.TypeRef
		want string
	}{
		{"nil", nil, ""},
		{"bare name", &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Name: "token", Model: "auth.model.token"}, "token"},
		{"full public qid", &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Name: "payment.model.payment_event", Model: "payment.model.payment_event"}, "payment_event"},
		{"private model qid", &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Name: "login_form", Model: "auth/task/login.yaml#login_form"}, "login_form"},
		{"no dots or hash", &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Name: "mymodel", Model: "mymodel"}, "mymodel"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := namedModelLocalID(tc.ref)
			if got != tc.want {
				t.Errorf("namedModelLocalID = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestCalcAssetHint_Ambiguity(t *testing.T) {
	orderModel1 := &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Name: "user_response", Model: "auth.model.user_response"}
	orderModel2 := &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Name: "user_response", Model: "payment.model.user_response"}
	clearModel := &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Name: "order", Model: "order.model.order"}
	primitiveRef := &semantic.TypeRef{Kind: semantic.TypeRefPrimitive, Name: "str"}

	ambiguous := computeAmbiguousHints([]*semantic.TypeRef{orderModel1, orderModel2, clearModel, primitiveRef})

	// Ambiguous: two QIDs map to local ID "user_response"
	if hint := calcAssetHint(orderModel1, ambiguous); hint != "" {
		t.Errorf("ambiguous named model should have no hint, got %q", hint)
	}
	// Not ambiguous: single QID for "order"
	if hint := calcAssetHint(clearModel, ambiguous); hint != "order" {
		t.Errorf("unambiguous named model should have hint 'order', got %q", hint)
	}
	// Primitive always has hint regardless of ambiguity map
	if hint := calcAssetHint(primitiveRef, ambiguous); hint != "str" {
		t.Errorf("primitive should have hint 'str', got %q", hint)
	}
	// nil ref
	if hint := calcAssetHint(nil, ambiguous); hint != "" {
		t.Errorf("nil ref should have no hint, got %q", hint)
	}
}

func TestComputeAmbiguousHints(t *testing.T) {
	// Same local ID, same QID (not ambiguous — repeated references to the same model)
	order1 := &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Name: "order", Model: "order.model.order"}
	order2 := &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Name: "order", Model: "order.model.order"}
	ambiguous := computeAmbiguousHints([]*semantic.TypeRef{order1, order2})
	if ambiguous["order"] {
		t.Errorf("same-QID refs should not be ambiguous")
	}

	// Same local ID, different QIDs (ambiguous)
	resp1 := &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Name: "response", Model: "auth.model.response"}
	resp2 := &semantic.TypeRef{Kind: semantic.TypeRefNamedModel, Name: "response", Model: "payment.model.response"}
	ambiguous = computeAmbiguousHints([]*semantic.TypeRef{resp1, resp2})
	if !ambiguous["response"] {
		t.Errorf("different-QID refs with same local ID should be ambiguous")
	}

	// Inline list — not a named model, should not appear in ambiguous map
	listRef := &semantic.TypeRef{Kind: semantic.TypeRefList, Name: "list"}
	ambiguous = computeAmbiguousHints([]*semantic.TypeRef{listRef})
	if ambiguous["list"] {
		t.Errorf("inline list should not appear in ambiguous hints map")
	}
}
