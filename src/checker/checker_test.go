package checker

import (
	// "errors"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindsNoDuplicateLines(t *testing.T) {
	expected := make(map[string]QuizItem)

	expected["abc"] = QuizItem{"abc", []int{0}}
	expected["cdf"] = QuizItem{"cdf", []int{1}}
	expected["joy"] = QuizItem{"joy", []int{2}}

	assert.Equal(t, expected, Check("abc\ncdf\njoy"))
}
