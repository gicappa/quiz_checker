package quizchecker

import (
	"fmt"
	"os"

	"github.com/gicappa/quiz_checker/src/checker"
	"github.com/gicappa/quiz_checker/src/file"
)

func main() {
	fmt.Println("------------------")
	contentFile := file.Load(parseFileName())
	response := checker.Check(contentFile)
	fmt.Println(response)
}

func parseFileName() string {

	if len(os.Args) != 1 {
		fmt.Printf("usage: %s <filename>", os.Args[0])
		os.Exit(1)
	}

	return os.Args[1]
}
