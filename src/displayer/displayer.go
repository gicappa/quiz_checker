package displayer

import (
	"fmt"

	"github.com/gicappa/quiz_checker/src/checker"
)

// Display the result of the checks
func Display(items map[string]checker.QuizItem) {
	for _, item := range items {
		if len(item.LineNumbers) > 1 {
			fmt.Println(item.QuizTest, item.LineNumbers)
		}
	}
}
