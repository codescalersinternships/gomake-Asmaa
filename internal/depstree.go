package internal

import (
	"fmt"
	"errors"
)

// Node represents a node in the graph
type Node struct {
	Name         string
	Dependencies []string
	Commands     []Command
	Visited      bool
	InStack      bool
}

type Command struct {
	command string
	prefix bool
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
func DFS(graph *Graph, node *Node, buildOrder *[]string) error {
	node.Visited = true
	node.InStack = true

	for _, dep := range node.Dependencies {
		depNode := graph.Nodes[dep]
		if depNode == nil {
			return fmt.Errorf("error %s dependency %s not found for target %s", ErrorDependencyNotFound, dep, node.Name)
		}

		if depNode.InStack {
			return fmt.Errorf("error %s found between: %s -> %s",ErrorCircularDependency, node.Name, depNode.Name)
		}

		if !depNode.Visited {
			err := DFS(graph, depNode, buildOrder)
			if err != nil {
				return err
			}
			continue
		} 
		if depNode.InStack {
			return fmt.Errorf("error %s found between: %s -> %s",ErrorCircularDependency, node.Name, depNode.Name)
		}
	}

	node.InStack = false
	*buildOrder = append(*buildOrder, node.Name)

	return nil
}


// CheckCircularDependencies checks for circular dependencies
func CheckCircularDependencies(graph *Graph) (error) {
	buildOrder := make([]string, 0)

	for _, node := range graph.Nodes {
		if !node.Visited {
			err := DFS(graph, node, &buildOrder)
			if err != nil {
				return err
			}
		}
	}

	return nil
}