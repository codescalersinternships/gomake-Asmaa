package main

import (
	"fmt"

	"github.com/codescalersinternships/gomake-Asmaa/internal"
)

func main() {

	filePath, target, err := internal.ParseCommand()
	if err != nil {
		fmt.Println("Error parsing Makefile:", err)
		return
	}

	file, err := internal.CheckMakeFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	graph, err := internal.ParseMakefile(file)
	if err != nil {
		fmt.Println("Error parsing Makefile:", err)
		return
	}

	err = internal.CheckCircularDependencies(graph)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	err = internal.RunTarget(graph, target)
	if err != nil {
		fmt.Println("Error running target:", err)
		return
	}
}
