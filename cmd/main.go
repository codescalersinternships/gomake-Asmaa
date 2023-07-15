
package main

import (
	"fmt"
	"flag"

	"github.com/codescalersinternships/gomake-Asmaa/internal"
)

func main() {

	arg1 := flag.String("f", "Makefile", "make file path")

	arg2 := flag.String("t", "", "make file path")
	flag.Parse()

	filePath, target, err := internal.ParseCommand(*arg1, *arg2)
	if err != nil {
		fmt.Println("Error parsing Makefile:", err)
		return
	}

	graph, err := internal.ParseMakefile(filePath)
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

