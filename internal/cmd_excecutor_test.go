package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRunTarget(t *testing.T) {
	t.Run("Run non existing target", func(t *testing.T) {
		graph := &Graph{
			Nodes: map[string]Node{
				"build": {
					dependencies: []string{},
					commands:     []string{"echo 'build'"},
				},
			},
		}

		targetName := "main.o"
		want := fmt.Errorf("target %s not found", targetName)

		err := graph.RunTarget(targetName)
		if !reflect.DeepEqual(err, want) {
			t.Errorf("run non existing target %s %s", err, want)
		}
	})

	t.Run("Run existing target", func(t *testing.T) {
		graph := &Graph{
			Nodes: map[string]Node{
				"build": {
					dependencies: []string{},
					commands:     []string{"echo 'build'"},
				},
			},
		}
		err := graph.RunTarget("build")
		if err != nil {
			t.Errorf("RunTarget returned an error: %s", err)
		}
	})

	t.Run("Run target with wrong command", func(t *testing.T) {
		graph := &Graph{
			Nodes: map[string]Node{
				"build": {
					dependencies: []string{},
					commands:     []string{"echoo 'build'"},
				},
			},
		}

		err := graph.RunTarget("build")
		if err == nil {
			t.Errorf("command is wrong should give error")
		}
	})

	t.Run("Run target with right command", func(t *testing.T) {
		graph := &Graph{
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
		err := graph.RunTarget("build")
		if err != nil {
			t.Errorf("RunTarget returned an error: %s", err)
		}
	})

	t.Run("Run target that has dependency with wrong command", func(t *testing.T) {
		graph := &Graph{
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
		err := graph.RunTarget("build")
		if err == nil {
			t.Errorf("command is wrong should give error")
		}
	})

	t.Run("Run target with more than one dependency", func(t *testing.T) {
		graph := &Graph{
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
		err := graph.RunTarget("build")
		if err != nil {
			t.Errorf("RunTarget returned an error: %s", err)
		}
	})

}
