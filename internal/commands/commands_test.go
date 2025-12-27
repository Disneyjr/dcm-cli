package commands

import (
	"testing"

	"github.com/Disneyjr/dcm/internal/workspace"
)

func TestResolveGroupServices(t *testing.T) {
	parallelTrue := true
	parallelFalse := false

	ws := &workspace.Workspace{
		Projects: map[string]workspace.Project{
			"p1": {Path: "./p1"},
			"p2": {Path: "./p2"},
			"p3": {Path: "./p3"},
		},
		Groups: map[string]workspace.Group{
			"base": {
				Services: []string{"p1", "p2"},
				Parallel: &parallelTrue,
			},
			"extended": {
				Extends:  "base",
				Services: []string{"p3"},
				Parallel: &parallelFalse,
			},
		},
	}

	// Test basic resolution
	services, parallel, err := resolveGroupServices(ws, "base", make(map[string]bool))
	if err != nil {
		t.Fatalf("Base resolution failed: %v", err)
	}
	if len(services) != 2 || services[0] != "p1" || services[1] != "p2" {
		t.Errorf("Unexpected services for base: %v", services)
	}
	if !parallel {
		t.Error("Base should be parallel")
	}

	// Test inheritance
	services, parallel, err = resolveGroupServices(ws, "extended", make(map[string]bool))
	if err != nil {
		t.Fatalf("Extended resolution failed: %v", err)
	}
	if len(services) != 3 || services[0] != "p1" || services[1] != "p2" || services[2] != "p3" {
		t.Errorf("Unexpected services for extended: %v", services)
	}
	if parallel {
		t.Error("Extended should be sequential (parallel=false)")
	}
}

func TestResolveGroupServicesCycle(t *testing.T) {
	ws := &workspace.Workspace{
		Groups: map[string]workspace.Group{
			"a": {Extends: "b", Services: []string{"p1"}},
			"b": {Extends: "a", Services: []string{"p2"}},
		},
	}

	_, _, err := resolveGroupServices(ws, "a", make(map[string]bool))
	if err == nil {
		t.Error("Expected error for cycle, got nil")
	}
}
