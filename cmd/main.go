package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/codescalersinternships/gomake-Asmaa/internal"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "f", "Makefile", "Makefile path")

	var target string
	flag.StringVar(&target, "t", "", "target")
	flag.Parse()

	if len(target) == 0 {
		fmt.Println("Error parsing Makefile:", internal.ErrorNoTarget)
		os.Exit(1)
	}

	graph, err := internal.ParseMakefile(filePath)
	if err != nil {
		fmt.Println("Error parsing Makefile:", err)
		os.Exit(2)
	}

	err = graph.CheckNoCommands()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(3)
	}

	err = graph.CheckCircularDependencies()
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(4)
	}

	err = graph.RunTarget(target)
	if err != nil {
		fmt.Println("Error running target:", err)
		os.Exit(5)
	}
}
