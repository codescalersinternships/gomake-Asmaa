package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

// RunTarget executes the commands for the specified target
func RunTarget(graph *Graph, targetName string) error {

	target, found := graph.Nodes[targetName]
	if !found {
		return fmt.Errorf("target %s not found", targetName)
	}

	// Execute dependencies first
	for _, dep := range target.dependencies {
		err := RunTarget(graph, dep)
		if err != nil {
			return err
		}
	}

	// Execute the target's commands
	for _, command := range target.commands {
		if !strings.HasPrefix(command, "@") {
			fmt.Println(command)
		} else {
			command = command[1:]
		}
		parts := strings.Fields(command)

		name := parts[0]
		arg := parts[1:]

		cmd := exec.Command(name, arg...)
		output, err := cmd.CombinedOutput()

		if err != nil {
			return err
		}

		fmt.Print(string(output))
	}

	return nil
}
