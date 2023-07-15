package internal

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestParseCommand(t *testing.T) {
	args1 := "file"
	args2 := "target"
	filePath, target, err := ParseCommand(args1, args2)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if filePath != "file" {
		t.Errorf("Expected file path: file, but got: %s", filePath)
	}

	if target != "target" {
		t.Errorf("Expected target name: target, but got: %s", target)
	}

	args2 = ""

	_, _, err = ParseCommand(args1, args2)

	if err != ErrorNoTarget {
		t.Errorf("error target: %s", err)
	}
}

func TestCheckMakeFile(t *testing.T) {
	_, err := CheckMakeFile("test")
	if err == nil {
		t.Errorf("Error: %s", err)
	}

	_, err = CheckMakeFile("../Makefile")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestParseMakefile(t *testing.T) {
	makefile := `
target:
	echo "Hello, World!"`

	file, err := ioutil.TempFile("", "Makefile")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString(makefile)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	expectedGraph := &Graph{
		Nodes: map[string]*Node{
			"target": {
				Dependencies: []string{},
				Commands:     []Command{{"echo \"Hello, World!\"", false}},
			},
		},
	}

	graph, err := ParseMakefile(file.Name())
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(graph, expectedGraph) {
		t.Errorf("Parsed graph %v does not match expected graph %v", graph, expectedGraph)
	}

	makefile = `
:
	echo "Hello, World!"`

	file, err = ioutil.TempFile("", "Makefile")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	defer os.Remove(file.Name())

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

	file, err = ioutil.TempFile("", "Makefile")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	defer os.Remove(file.Name())

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

	file, err = ioutil.TempFile("", "Makefile")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	defer os.Remove(file.Name())

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
	file, err := ioutil.TempFile("", "Makefile")
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
