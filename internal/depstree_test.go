package internal

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCheckCircularDependencies(t *testing.T) {
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

	graph, err := ParseMakefile(file.Name())
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	err = CheckCircularDependencies(graph)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	makefile = `target: test
	echo "Hello, World!"`
	_, err = file.WriteString(makefile)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	graph, err = ParseMakefile(file.Name())
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	err = CheckCircularDependencies(graph)
	if err == nil {
		t.Errorf("Error: %s", err)
	}

	makefile = `
build: test
	echo 'executing buildaaaa'
	@echo 'cmd2'

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

	err = CheckCircularDependencies(graph)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
}
