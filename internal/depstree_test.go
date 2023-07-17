package internal

import (
	"testing"
)

func TestCheckCircularDependencies(t *testing.T) {

	graph := &Graph{
		Nodes: map[string]Node{
			"target": {
				dependencies: []string{},
				commands:     []string{`echo "Hello, World!"`},
			},
		},
	}
	err := graph.CheckCircularDependencies()
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	graph = &Graph{
		Nodes: map[string]Node{
			"target": {
				dependencies: []string{"test"},
				commands:     []string{`echo "Hello, World!"`},
			},
		},
	}

	err = graph.CheckCircularDependencies()
	if err == nil {
		t.Errorf("expect found error")
	}

	graph = &Graph{
		Nodes: map[string]Node{
			"build": {
				dependencies: []string{"test"},
				commands:     []string{`echo "build"`, `@echo 'cmd2'`},
			},
			"test": {
				dependencies: []string{"build"},
				commands:     []string{`echo "test"`},
			},
		},
	}

	err = graph.CheckCircularDependencies()
	if err == nil {
		t.Errorf("Error: %s", err)
	}
}
