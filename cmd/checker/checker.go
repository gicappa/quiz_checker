package main

import (
	"fmt"
	"file"
	"os"
)

func main() {
	file_name := os.Args[1]
	fmt.Println(string(file.LoadQuiz(file_name)))
}
