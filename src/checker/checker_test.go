package checker

import (
	// "errors"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindsNoDuplicateLines(t *testing.T) {
	m := make(map[string]QuizItem)

	m["abc"] = QuizItem{"abc", []int{0}}
	m["cdf"] = QuizItem{"cdf", []int{1}}
	m["joy"] = QuizItem{"joy", []int{2}}

	assert.Equal(t, m, Check("abc\ncdf\njoy"))
}
