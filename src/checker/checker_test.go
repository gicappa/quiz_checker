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
func TestFindsADuplicateLineNoCase(t *testing.T) {
	expected := make(map[string]QuizItem)

	expected["iam"] = QuizItem{"I am", []int{1, 3}}
	expected["youare"] = QuizItem{"You are", []int{2}}
	expected["heis"] = QuizItem{"He is", []int{4}}

	assert.Equal(t, expected, Check("I am\nYou are\nI am\nHe is"))
}

func TestDisregardsSpaceDifferences(t *testing.T) {
	expected := make(map[string]QuizItem)

	expected["myduplicate"] = QuizItem{"my Du pli cate", []int{1, 2}}
	expected["notaduplicate"] = QuizItem{"Not a duplicate", []int{3}}

	assert.Equal(t, expected, Check("my Du pli cate\nMy d upli Ca te\nNot a duplicate"))
}

func TestDisregardsSymbolsDifferences(t *testing.T) {
	expected := make(map[string]QuizItem)

	expected["myduplicate"] = QuizItem{"my, Du pli cate1", []int{1, 2}}
	expected["notaduplicate"] = QuizItem{"Not a duplicate", []int{3}}

	assert.Equal(t, expected, Check("my, Du pli cate1\nMy.. d; up';\"li Ca te1\nNot a duplicate"))
}
