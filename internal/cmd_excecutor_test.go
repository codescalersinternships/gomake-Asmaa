package internal

import (
	"testing"
)

func TestRunTarget(t *testing.T) {
	// Set up a sample graph with targets and commands
	graph := &Graph{
		Nodes: map[string]Node{
			"build": {
				dependencies: []string{},
				commands:     []string{"echo 'build'"},
			},
		},
	}

	// Invoke the RunTarget function with the "build" target
	err := RunTarget(graph, "main.o")
	if err == nil {
		t.Errorf("RunTarget returned an error: %s", err)
	}

	// Invoke the RunTarget function with the "build" target
	err = RunTarget(graph, "build")
	if err != nil {
		t.Errorf("RunTarget returned an error: %s", err)
	}

	graph = &Graph{
		Nodes: map[string]Node{
			"build": {
				dependencies: []string{},
				commands:     []string{"echoo 'build'"},
			},
		},
	}

	// Invoke the RunTarget function with the "build" target
	err = RunTarget(graph, "build")
	if err == nil {
		t.Errorf("Error: %s", err)
	}

	graph = &Graph{
		Nodes: map[string]Node{
			"build": {
				dependencies: []string{"test"},
				commands:     []string{"@echo 'build'"},
			},
			"test": {
				dependencies: []string{},
				commands:     []string{"echo 'test'"},
			},
		},
	}
	err = RunTarget(graph, "build")
	if err != nil {
		t.Errorf("RunTarget returned an error: %s", err)
	}

	graph = &Graph{
		Nodes: map[string]Node{
			"build": {
				dependencies: []string{"test"},
				commands:     []string{"echo 'build'"},
			},
			"test": {
				dependencies: []string{},
				commands:     []string{"echoo 'test'"},
			},
		},
	}
	err = RunTarget(graph, "build")
	if err == nil {
		t.Errorf("Error: %s", err)
	}
}
