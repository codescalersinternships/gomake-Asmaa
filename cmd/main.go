package cmd

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/codescalersinternships/gomake-Asmaa/internal"
)

func main() {
	filePath, target, err := internal.ParseCommand()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	file, err := CheckMakeFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	graph, err := Make(file, target)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(graph)
	// checkCircular()
	// TargetExcecution(file, target)

}

//////////////////////////////////////////////////////////
///////////////////////// GOMAKE /////////////////////////
//////////////////////////////////////////////////////////

// TargetNotFound
var TargetNotFound = errors.New("target not found")

func ParseCommand() (string, string, error) {
	filePath := flag.String("f", "Makefile", "make file path")

	target := flag.String("t", "", "make file path")
	flag.Parse()

	if len(*target) == 0 {
		return "", "", TargetNotFound
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

func Make(file *os.File, target string) (*Graph, error) {
	graph := &Graph{
		Nodes: make(map[string]*Node),
	}
	var currentTarget *Node

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")

			targetName := strings.TrimSpace(parts[0])
			dependencies := strings.Split(strings.TrimSpace(parts[1]), " ")
			currentTarget = &Node{
				Name:         targetName,
				Dependencies: dependencies,
				Commands:     make([]string, 0),
			}
			graph.Nodes[targetName] = currentTarget
		} else if currentTarget != nil {
			Commands := strings.Fields(line)
			currentTarget.Commands = append(currentTarget.Commands, Commands...)
		}
	}
	return graph, scanner.Err()
}

//////////////////////////////////////////////////////////
//////////////////////// depstree ////////////////////////
//////////////////////////////////////////////////////////

// Node represents a node in the graph
type Node struct {
	Name         string
	Dependencies []string
	Commands     []string
}

// Graph represents the entire graph
type Graph struct {
	Nodes map[string]*Node
}
