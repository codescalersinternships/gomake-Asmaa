package internal

import (
	"errors"
	"fmt"
)

// Node represents a node in the graph
type Node struct {
	dependencies []string
	commands     []string
	visited      bool
	inStack      bool
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

	visitedNode.visited = true
	visitedNode.inStack = true

	for _, dep := range visitedNode.dependencies {
		depNode := graph.Nodes[dep]
		if depNode == nil {
			return fmt.Errorf("%w dependency: %s not found for target: %s", ErrorDependencyNotFound, dep, node)
		}

		if depNode.inStack {
			return fmt.Errorf("%w found between: %s -> %s", ErrorCircularDependency, node, dep)
		}

		if !depNode.visited {
			err := DFS(graph, dep)
			if err != nil {
				return err
			}
		}
	}

	visitedNode.inStack = false
	return nil
}

// CheckCircularDependencies checks for circular dependencies
func CheckCircularDependencies(graph *Graph) error {

	for node := range graph.Nodes {
		visitedNode := graph.Nodes[node]
		if !visitedNode.visited {
			err := DFS(graph, node)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
