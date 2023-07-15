package internal

import (
	"os"
	"reflect"
	"testing"
	"io/ioutil"
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
	makefile := `target:
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
				Name:         "target",
				Dependencies: []string{},
				Commands:    []Command{Command{"echo \"Hello, World!\"", false}},
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
}
