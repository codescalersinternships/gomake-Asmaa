package internal

import (
	"errors"
	"flag"
)

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
