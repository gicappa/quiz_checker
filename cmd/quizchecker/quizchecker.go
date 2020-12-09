package main

import (
	"fmt"
	"os"

	"github.com/gicappa/quiz_checker/src/displayer"

	"github.com/gicappa/quiz_checker/src/checker"
	"github.com/gicappa/quiz_checker/src/file"
)

func main() {
	contentFile := file.Load(parseFileName())
	response := checker.Check(contentFile)
	displayer.Display(response)
}

func parseFileName() string {

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <filename>", os.Args[0])
		os.Exit(1)
	}

	return os.Args[1]
}
