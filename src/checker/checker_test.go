package checker

import (
	// "errors"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindsNoDuplicateLines(t *testing.T) {
	expected := make(map[string]QuizItem)

	expected["abc"] = QuizItem{"abc", []int{1}}
	expected["cdf"] = QuizItem{"cdf", []int{2}}
	expected["joy"] = QuizItem{"joy", []int{3}}

	assert.Equal(t, expected, Check("abc\ncdf\njoy"))
}
func TestFindsADuplicateLine(t *testing.T) {
	expected := make(map[string]QuizItem)

	expected["abc"] = QuizItem{"abc", []int{1}}
	expected["cdf"] = QuizItem{"cdf", []int{2, 4}}
	expected["joy"] = QuizItem{"joy", []int{3}}

	assert.Equal(t, expected, Check("abc\ncdf\njoy\ncdf"))
}
