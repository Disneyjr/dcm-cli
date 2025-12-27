package workspace

import (
	"os"
	"testing"
)

func TestLoadWorkspace(t *testing.T) {
	content := `{
  "version": "1.0",
  "projects": {
    "test": { "path": "./test", "description": "Test Project" }
  },
  "groups": {
    "test-group": { "services": ["test"] }
  }
}`
	err := os.WriteFile("workspace.json", []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create test workspace.json: %v", err)
	}
	defer os.Remove("workspace.json")

	ws := NewWorkspace()
	err = LoadWorkspace(ws)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if ws.Version != "1.0" {
		t.Errorf("expected version 1.0, got %s", ws.Version)
	}

	if _, ok := ws.Projects["test"]; !ok {
		t.Errorf("expected project 'test' to exist")
	}
}
