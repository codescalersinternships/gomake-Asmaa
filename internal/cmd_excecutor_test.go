package internal

import (
	"testing"
)

func TestRunTarget(t *testing.T) {
	// Set up a sample graph with targets and commands
	graph := &Graph{
		Nodes: map[string]*Node{
			"build": {
				Name:         "build",
				Dependencies: []string{},
				Commands:     []Command{{"echo 'build'", false}},
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
		Nodes: map[string]*Node{
			"build": {
				Name:         "build",
				Dependencies: []string{},
				Commands:     []Command{{"echoo 'build'", false}},
			},
		},
	}

	// Invoke the RunTarget function with the "build" target
	err = RunTarget(graph, "build")
	if err == nil {
		t.Errorf("Error: %s", err)
	}

	graph = &Graph{
		Nodes: map[string]*Node{
			"build": {
				Name:         "build",
				Dependencies: []string{"test"},
				Commands:     []Command{{"echo 'build'", false}},
			},
			"test": {
				Name:         "test",
				Dependencies: []string{},
				Commands:     []Command{{"echo 'test'", false}},
			},
		},
	}
	err = RunTarget(graph, "build")
	if err != nil {
		t.Errorf("RunTarget returned an error: %s", err)
	}

	graph = &Graph{
		Nodes: map[string]*Node{
			"build": {
				Name:         "build",
				Dependencies: []string{"test"},
				Commands:     []Command{{"echo 'build'", false}},
			},
			"test": {
				Name:         "test",
				Dependencies: []string{},
				Commands:     []Command{{"echoo 'test'", false}},
			},
		},
	}
	err = RunTarget(graph, "build")
	if err == nil {
		t.Errorf("Error: %s", err)
	}
}
