package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseMakefile(t *testing.T) {

	dir := os.TempDir()
	filePath := filepath.Join(dir, "Makefile")

	t.Run("Run non existing file", func(t *testing.T) {
		_, err := ParseMakefile("Makefile")
		if err == nil {
			t.Errorf("error in makefile path")
		}
	})

	t.Run("Run existing file with right data", func(t *testing.T) {
		makefile := `
target:
	echo "Hello, World!"`

		file, err := os.Create(filePath)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		defer os.Remove(file.Name())

		_, err = file.WriteString(makefile)
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		_, err = ParseMakefile(file.Name())
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		expectedGraph := &Graph{
			Nodes: map[string]Node{
				"target": {
					dependencies: []string{},
					commands:     []string{`echo "Hello, World!"`},
				},
			},
		}

		graph, _ := ParseMakefile(file.Name())

		if !reflect.DeepEqual(graph.Nodes, expectedGraph.Nodes) {
			t.Errorf("Parsed graph %v does not match expected graph %v", graph, expectedGraph)
		}

	})

	t.Run("Run invalid format file", func(t *testing.T) {
		makefile := `
		:
			echo "Hello, World!"`

		file, err := os.Create(filePath)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		_, err = file.WriteString(makefile)
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		_, err = ParseMakefile(file.Name())
		if err != ErrorInvalidFormat {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("Run invalid format file", func(t *testing.T) {
		makefile := `
		: test
		   echo "Hello, World!"`

		file, _ := os.Create(filePath)
		_, err := file.WriteString(makefile)
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		_, err = ParseMakefile(file.Name())
		if err != ErrorInvalidFormat {
			t.Errorf("Error: %s", err)
		}

	})

	t.Run("Run invalid format file", func(t *testing.T) {
		makefile := `
		echo 'executing build`

		file, _ := os.Create(filePath)
		_, err := file.WriteString(makefile)
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		_, err = ParseMakefile(file.Name())
		if err != ErrorInvalidFormat {
			t.Errorf("Error: %s", err)
		}
	})
}

func TestCheckNoCommands(t *testing.T) {

	t.Run("Run target without commands", func(t *testing.T) {
		graph := &Graph{
			Nodes: map[string]Node{
				"build": {
					dependencies: []string{"test"},
					commands:     []string{},
				},
				"test": {
					dependencies: []string{},
					commands:     []string{"echo 'test'"},
				},
			},
		}

		want := fmt.Errorf("%w:%v", ErrorNoCommandFound, graph.Nodes["build"])

		err := graph.CheckNoCommands()
		if !reflect.DeepEqual(err, want) {
			t.Errorf("target hasn't commands should give error")
		}
	})

	t.Run("Run target commands", func(t *testing.T) {
		graph := &Graph{
			Nodes: map[string]Node{
				"build": {
					dependencies: []string{"test"},
					commands:     []string{"echo 'build'"},
				},
				"test": {
					dependencies: []string{},
					commands:     []string{"echo 'test'"},
				},
			},
		}

		err := graph.CheckNoCommands()
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

}
