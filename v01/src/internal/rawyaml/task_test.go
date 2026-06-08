package rawyaml

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestReturnSourceYAML(t *testing.T) {
	var task Task
	input := []byte(`id: run
returns:
  name: result
  model: order
  source: validated_items
`)
	if err := yaml.Unmarshal(input, &task); err != nil {
		t.Fatalf("yaml.Unmarshal returned error: %v", err)
	}
	if task.Returns == nil {
		t.Fatalf("task.Returns = nil, want return")
	}
	if task.Returns.Source != "validated_items" {
		t.Fatalf("task.Returns.Source = %q, want validated_items", task.Returns.Source)
	}
}
