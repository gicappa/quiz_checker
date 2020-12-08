package checker

import (
	// "errors"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindsNoDuplicateLines(t *testing.T) {
	m := make(map[string]QuizItem)

	m["abc"] = QuizItem{"abc", nil}
	m["cdf"] = QuizItem{"cdf", nil}
	m["joy"] = QuizItem{"joy", nil}

	assert.Equal(t, m, Check("abc\ncdf\njoy"))
}
