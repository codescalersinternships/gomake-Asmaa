package internal

import (
	"testing"
)

func TestRunTarget(t *testing.T) {
	graph := &Graph{
		Nodes: map[string]Node{
			"build": {
				dependencies: []string{},
				commands:     []string{"echo 'build'"},
			},
		},
	}

	err := graph.RunTarget("main.o")
	if err == nil {
		t.Errorf("No target has name main.o")
	}

	err = graph.RunTarget("build")
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

	err = graph.RunTarget("build")
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
	err = graph.RunTarget("build")
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
	err = graph.RunTarget("build")
	if err == nil {
		t.Errorf("error executing command")
	}

	graph = &Graph{
		Nodes: map[string]Node{
			"build": {
				dependencies: []string{"test", "publish"},
				commands:     []string{"echo 'build'"},
			},
			"test": {
				dependencies: []string{"execute"},
				commands:     []string{"echo 'test'"},
			},
			"publish": {
				dependencies: []string{},
				commands:     []string{"echo 'publish'"},
			},
			"execute": {
				dependencies: []string{},
				commands:     []string{"echo 'execute'"},
			},
		},
	}
	err = graph.RunTarget("build")
	if err != nil {
		t.Errorf("RunTarget returned an error: %s", err)
	}
}
