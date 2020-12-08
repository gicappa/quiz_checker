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

	for i := 1; scanner.Scan(); i++ {
		item := getOrCreateQuizItem(items, scanner.Text())
		item.lineNumbers = append(item.lineNumbers, i)
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
	return strings.ReplaceAll(strings.ToLower(str), " ", "")
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
