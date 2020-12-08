package checker

import (
	"bufio"
	"strings"
)

// QuizItem contains the question and the
// line numbers in which is present
type QuizItem struct {
	quizTest    string
	lineNumbers []int
}

// Check the lines of a file for duplicates
func Check(lines string) map[string]QuizItem {
	items := make(map[string]QuizItem)

	scanner := bufio.NewScanner(strings.NewReader(lines))

	for i := 0; scanner.Scan(); i++ {
		items[scanner.Text()] = QuizItem{scanner.Text(), []int{i}}
	}

	return items
}

// Map<String, QuizItem> lines = new HashMap<>();

// StringTokenizer tokenizer = new StringTokenizer(quizText, "\n");

// for (Integer i = 1; tokenizer.hasMoreTokens(); i++) {
// 		String quizItemText = tokenizer.nextToken();

// 		QuizItem quizItem = lines.get(hash(quizItemText));

// 		if (quizItem == null) {
// 				quizItem = new QuizItem(quizItemText);
// 		}

// 		quizItem.addLineNumber(i);

// 		lines.put(quizItem.hash(), quizItem);
// }

// return lines;
