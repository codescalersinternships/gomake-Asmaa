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
	Nodes map[string]Node
}

// ErrorDependencyNotFound
var ErrorDependencyNotFound = errors.New("dependency not found")

// ErrorCircularDependency
var ErrorCircularDependency = errors.New("circular dependency")

// DFS performs a depth-first search to detect cycles
func (graph *Graph) DFS(node string, visited map[string]bool, inStack map[string]bool) error {

	visitedNode := graph.Nodes[node]

	visited[node] = true
	inStack[node] = true

	for _, dep := range visitedNode.dependencies {
		_, found := graph.Nodes[dep]
		if !found {
			return fmt.Errorf("%w dependency: %s not found for target: %s", ErrorDependencyNotFound, dep, node)
		}

		if inStack[dep] {
			return fmt.Errorf("%w found between: %s -> %s", ErrorCircularDependency, node, dep)
		}

		if !visited[dep] {
			err := graph.DFS(dep, visited, inStack)
			if err != nil {
				return err
			}
		}
	}

	inStack[node] = false
	return nil
}

// CheckCircularDependencies checks for circular dependencies
func (graph *Graph) CheckCircularDependencies() error {

	visited := map[string]bool{}
	inStack := map[string]bool{}

	for node := range graph.Nodes {
		if !visited[node] {
			err := graph.DFS(node, visited, inStack)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
