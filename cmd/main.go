package main

import (
	"flag"
	"log"

	"github.com/codescalersinternships/gomake-Asmaa/internal"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "f", "Makefile", "Makefile path")

	var target string
	flag.StringVar(&target, "t", "", "target")
	flag.Parse()

	app := internal.NewApp()
	err := app.Run(filePath, target)
	if err != nil {
		log.Fatalf("Error:%p", err)
	}
}
