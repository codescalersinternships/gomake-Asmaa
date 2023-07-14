package internal

import (
	"errors"
	"flag"
	"os"
	"strings"
	"bufio"
)

// ErrorNoTarget
var ErrorNoTarget = errors.New("no target found")

// ErrorInvalidFormat
var ErrorInvalidFormat = errors.New("invalid format for makefile")

func ParseCommand() (string, string, error) {
	filePath := flag.String("f", "Makefile", "make file path")

	target := flag.String("t", "", "make file path")
	flag.Parse()

	if len(*target) == 0 {
		return "", "", ErrorNoTarget
	}
	return *filePath, *target, nil
}

func CheckMakeFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// ParseMakefile parses the Makefile and returns the graph representation
func ParseMakefile(file *os.File) (*Graph, error) {

	scanner := bufio.NewScanner(file)

	graph := &Graph{
		Nodes: make(map[string]*Node),
	}
	var currentTarget *Node

	for scanner.Scan() {
		line := scanner.Text()
		// Skip comments and empty lines
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue 
		}

		// Found a new target
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")

			targetName := strings.TrimSpace(parts[0])
			if(targetName == "") {
				return nil, ErrorInvalidFormat
			}

			currentTarget = &Node{
				Name:         targetName,
				Dependencies: make([]string, 0),
				Commands:     make([]Command, 0),
			}

			graph.Nodes[targetName] = currentTarget
			// Found a dependency for the current target
			dependencies := strings.Fields(parts[1])
			currentTarget.Dependencies = append(currentTarget.Dependencies, dependencies...)

			continue
		} 

		// Found a command for the current target
		if strings.HasPrefix(line, "\t") && currentTarget != nil {
			command := strings.TrimPrefix(line, "\t")
			prefix := false
			if strings.HasPrefix(command, "@") {
				prefix = true
				command = command[1:]
			}
			currentTarget.Commands = append(currentTarget.Commands, Command{command, prefix})
			continue
		}
		return nil, ErrorInvalidFormat
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}