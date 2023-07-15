package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// ErrorNoTarget
var ErrorNoTarget = errors.New("no target found")

// ErrorInvalidFormat
var ErrorInvalidFormat = errors.New("invalid format for makefile")

// ErrorNoFile
var ErrorNoFile = errors.New("file not found")

// ErrorNoCommandFound
var ErrorNoCommandFound = errors.New("commands not found for target")

// CheckMakeFile checks for error in openning file
func CheckMakeFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, ErrorNoFile
	}

	return file, nil
}

// ParseCommand parses command entered from the user
func ParseCommand(filePath string, target string) (string, string, error) {

	if len(target) == 0 {
		return "", "", ErrorNoTarget
	}
	return filePath, target, nil
}

// ParseMakefile parses the Makefile and returns the graph representation
func ParseMakefile(filePath string) (*Graph, error) {

	file, err := CheckMakeFile(filePath)
	if err != nil {
		return nil, err
	}

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
			if targetName == "" {
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
	return graph, nil
}

// CheckNoCommands checks if there is a traget that hasn't commands
func CheckNoCommands(graph *Graph) error {
	for _, node := range graph.Nodes {
		if len(node.Commands) == 0 {
			return fmt.Errorf("error %s:%s", ErrorNoCommandFound, node.Name)
		}
	}
	return nil
}
