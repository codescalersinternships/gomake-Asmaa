package main

import (
	"flag"
	"fmt"

	"github.com/codescalersinternships/gomake-Asmaa/internal"
)

func main() {

	makefilePath := flag.String("f", "Makefile", "Makefile path")

	targetFlag := flag.String("t", "", "target")
	flag.Parse()

	filePath, target, err := internal.ParseCommand(*makefilePath, *targetFlag)
	if err != nil {
		fmt.Println("Error parsing Makefile:", err)
		return
	}

	graph, err := internal.ParseMakefile(filePath)
	if err != nil {
		fmt.Println("Error parsing Makefile:", err)
		return
	}

	err = internal.CheckNoCommands(graph)
	if err != nil {
		fmt.Println("Error:", err)
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
