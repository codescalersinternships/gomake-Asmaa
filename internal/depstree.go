package internal

import (
	"errors"
	"fmt"
)

// Node represents a node in the graph
type Node struct {
	Dependencies []string
	Commands     []Command
	Visited      bool
	InStack      bool
}

type Command struct {
	command string
	prefix  bool
}

// Graph represents the entire graph
type Graph struct {
	Nodes map[string]*Node
}

// ErrorDependencyNotFound
var ErrorDependencyNotFound = errors.New("dependency not found")

// ErrorCircularDependency
var ErrorCircularDependency = errors.New("circular dependency")

// DFS performs a depth-first search to detect cycles
func DFS(graph *Graph, node string) error {

	visitedNode := graph.Nodes[node]

	visitedNode.Visited = true
	visitedNode.InStack = true

	for _, dep := range visitedNode.Dependencies {
		depNode := graph.Nodes[dep]
		if depNode == nil {
			return fmt.Errorf("error %s dependency %s not found for target %s", ErrorDependencyNotFound, dep, node)
		}

		if depNode.InStack {
			return fmt.Errorf("error %s found between: %s -> %s", ErrorCircularDependency, node, dep)
		}

		if !depNode.Visited {
			err := DFS(graph, dep)
			if err != nil {
				return err
			}
		}
	}

	visitedNode.InStack = false
	return nil
}

// CheckCircularDependencies checks for circular dependencies
func CheckCircularDependencies(graph *Graph) error {

	for node := range graph.Nodes {
		visitedNode := graph.Nodes[node]
		if !visitedNode.Visited {
			err := DFS(graph, node)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
