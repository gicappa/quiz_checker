package checker

import (
	"bufio"
	"regexp"
	"strings"
)

// QuizItem contains the question and the
// line numbers in which is present
type QuizItem struct {
	QuizTest    string
	LineNumbers []int
}

// Check the lines of a file for duplicates
func Check(lines string) map[string]QuizItem {
	items := make(map[string]QuizItem)

	scanner := bufio.NewScanner(strings.NewReader(lines))

	for i := 1; scanner.Scan(); i++ {
		item := getOrCreateQuizItem(items, scanner.Text())
		item.LineNumbers = append(item.LineNumbers, i)
		items[hash(scanner.Text())] = item
	}

	return items
}

func getOrCreateQuizItem(items map[string]QuizItem, question string) QuizItem {

	if _, ok := items[hash(question)]; ok {
		return items[hash(question)]
	}

	return QuizItem{question, []int{}}
}

func hash(str string) string {
	return regexp.MustCompile(`[^a-z\\d]`).ReplaceAllString(strings.ToLower(str), ``)
}
