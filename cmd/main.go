package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/codescalersinternships/gomake-Asmaa/internal"
)

func main() {

	filePath := flag.String("f", "Makefile", "Makefile path")

	target := flag.String("t", "", "target")
	flag.Parse()

	if len(*target) == 0 {
		fmt.Println("Error parsing Makefile:", internal.ErrorNoTarget)
		return
	}

	graph, err := internal.ParseMakefile(*filePath)
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

	err = internal.RunTarget(graph, *target)
	if err != nil {
		fmt.Println("Error running target:", err)
		os.Exit(1)
	}
}
