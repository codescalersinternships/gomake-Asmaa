package internal

import (
	"errors"
	"fmt"
)

// Node represents a node in the graph
type Node struct {
	dependencies []string
	commands     []string
}

// Graph represents the entire graph
type Graph struct {
	Nodes   map[string]Node
	visited map[string]bool
	inStack map[string]bool
}

// ErrorDependencyNotFound
var ErrorDependencyNotFound = errors.New("dependency not found")

// ErrorCircularDependency
var ErrorCircularDependency = errors.New("circular dependency")

// DFS performs a depth-first search to detect cycles
func (graph *Graph) DFS(node string) error {

	visitedNode := graph.Nodes[node]

	graph.visited[node] = true
	graph.inStack[node] = true

	for _, dep := range visitedNode.dependencies {
		_, found := graph.Nodes[dep]
		if !found {
			return fmt.Errorf("%w dependency: %s not found for target: %s", ErrorDependencyNotFound, dep, node)
		}

		if graph.inStack[dep] {
			return fmt.Errorf("%w found between: %s -> %s", ErrorCircularDependency, node, dep)
		}

		if !graph.visited[dep] {
			err := graph.DFS(dep)
			if err != nil {
				return err
			}
		}
	}

	graph.inStack[node] = false
	return nil
}

// CheckCircularDependencies checks for circular dependencies
func CheckCircularDependencies(graph *Graph) error {

	for node := range graph.Nodes {
		if !graph.visited[node] {
			err := graph.DFS(node)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
