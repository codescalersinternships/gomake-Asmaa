package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// ErrorNoTarget if user didn't enter target while running command
var ErrorNoTarget = errors.New("no target found")

// ErrorInvalidFormat if there is an error in Makefile format
var ErrorInvalidFormat = errors.New("invalid format for makefile")

// ErrorNoCommandFound if target hassn't commands
var ErrorNoCommandFound = errors.New("commands not found for target")

func newGrap() *Graph {
	return &Graph{
		Nodes: make(map[string]Node),
	}
}

// ParseMakefile parses the Makefile and returns the graph representation
func ParseMakefile(filePath string) (*Graph, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	graph := newGrap()
	var currentTarget Node
	targetName := ""

	for scanner.Scan() {
		line := scanner.Text()
		// Skip comments and empty lines
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		// Found a command for the current target
		if strings.HasPrefix(line, "\t") && targetName != "" {
			command := strings.TrimPrefix(line, "\t")
			currentTarget.commands = append(currentTarget.commands, command)
			graph.Nodes[targetName] = currentTarget
			continue
		}

		// Found a new target
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")

			targetName = strings.TrimSpace(parts[0])
			if targetName == "" {
				return nil, ErrorInvalidFormat
			}

			currentTarget = Node{
				dependencies: make([]string, 0),
				commands:     make([]string, 0),
			}

			// Found a dependency for the current target
			dependencies := strings.Fields(parts[1])
			currentTarget.dependencies = append(currentTarget.dependencies, dependencies...)
			graph.Nodes[targetName] = currentTarget
			continue
		}
		return nil, ErrorInvalidFormat
	}
	return graph, nil
}

// CheckNoCommands checks if there is a traget that hasn't commands
func (graph *Graph) CheckNoCommands() error {
	for _, node := range graph.Nodes {
		if len(node.commands) == 0 {
			return fmt.Errorf("%w:%v", ErrorNoCommandFound, node)
		}
	}
	return nil
}
