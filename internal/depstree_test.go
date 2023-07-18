package internal

import (
	"errors"
	"fmt"
	"testing"
)

func TestCheckCircularDependencies(t *testing.T) {

	t.Run("Graph without circulat dependency", func(t *testing.T) {
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
	})

	t.Run("Run target with non existing dependency", func(t *testing.T) {
		graph := &Graph{
			Nodes: map[string]Node{
				"target": {
					dependencies: []string{"test"},
					commands:     []string{`echo "Hello, World!"`},
				},
			},
		}

		dep := "test"
		node := "target"
		want := fmt.Errorf("%w dependency: %s not found for target: %s", ErrorDependencyNotFound, dep, node)

		err := graph.CheckCircularDependencies()
		if errors.Is(err, want) {
			t.Errorf("there is non existing dependency found for target")
		}
	})

	t.Run("Graph with circular dependency", func(t *testing.T) {
		graph := &Graph{
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

		node := "test"
		dep := "build"
		want := fmt.Errorf("%w found between: %s -> %s", ErrorCircularDependency, node, dep)

		err := graph.CheckCircularDependencies()
		if errors.Is(err, want) {
			t.Errorf("circular dependency should exist")
		}
	})
}
