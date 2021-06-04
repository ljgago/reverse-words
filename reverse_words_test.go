package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var sentenceTests = []struct {
	sentence        string // input
	reverseSentence string // expected result
}{
	{"today is the first day of the rest of my life", "life my of rest the of day first the is today"},
	{"Hello, how are you?", "you? are how Hello,"},
	{"That's one small step for man, one giant leap for mankind", "mankind for leap giant one man, for step small one That's"},
}

func TestReverseWords(t *testing.T) {
	for _, s := range(sentenceTests) {
		assert.Equal(t, s.reverseSentence, ReverseWords(s.sentence), "the sentence should be equal")
	}
}
