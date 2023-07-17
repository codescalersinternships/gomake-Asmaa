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

	err := RunTarget(graph, "main.o")
	if err == nil {
		t.Errorf("No target has name main.o")
	}

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

	err = RunTarget(graph, "build")
	if err == nil {
		t.Errorf("error executing command")
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
		t.Errorf("error executing command")
	}
}
