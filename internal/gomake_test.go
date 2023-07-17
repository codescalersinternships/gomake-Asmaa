package internal

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseMakefile(t *testing.T) {

	makefile := `
target:
	echo "Hello, World!"`

	_, err := ParseMakefile("Makefile")
	if err == nil {
		t.Errorf("error in makefile path")
	}

	dir := os.TempDir()
	filePath := filepath.Join(dir, "Makefile")
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

	graph, err := ParseMakefile(file.Name())
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(graph.Nodes, expectedGraph.Nodes) {
		t.Errorf("Parsed graph %v does not match expected graph %v", graph, expectedGraph)
	}

	makefile = `
:
	echo "Hello, World!"`

	file, err = os.Create(filePath)
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

	makefile = `
 : test
	echo "Hello, World!"`

	file, err = os.Create(filePath)
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

	_, err = ParseMakefile("test")
	if err == nil {
		t.Errorf("Error: %s", err)
	}

	makefile = `
	echo 'executing build`

	file, err = os.Create(filePath)
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

}

func TestCheckNoCommands(t *testing.T) {
	makefile := `
	build: test
	
	test: build
		echo 'testtttttttttttt'`
	dir := os.TempDir()
	filePath := filepath.Join(dir, "Makefile")
	file, err := os.Create(filePath)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString(makefile)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	graph, err := ParseMakefile(file.Name())
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	err = CheckNoCommands(graph)
	if err == nil {
		t.Errorf("Error: %s", err)
	}

	makefile = `
	build: test
		echo 'build'
	test: build
		echo 'testtttttttttttt'`

	_, err = file.WriteString(makefile)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	graph, err = ParseMakefile(file.Name())
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	err = CheckNoCommands(graph)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
