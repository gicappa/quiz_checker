package quiz_checker

import (
	"fmt"
	"github.com/gicappa/quiz_checker/src/file"
	"os"
)

func main() {
	file_name := os.Args[1]
	fmt.Println(string(file.LoadQuiz(file_name)))
}
