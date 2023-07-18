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
		return
	}

	graph, err := internal.ParseMakefile(filePath)
	if err != nil {
		fmt.Println("Error parsing Makefile:", err)
		return
	}

	err = graph.CheckNoCommands()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = graph.CheckCircularDependencies()
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	err = graph.RunTarget(target)
	if err != nil {
		fmt.Println("Error running target:", err)
		os.Exit(1)
	}
}
