package main

import (
	"fmt"
	"strings"
	"flag"
)

// ReverseWords return the reverse words of a string sentence.
func ReverseWords(s string) string {
	var reverse []string

	// split the sentece in an array of string words
	words := strings.Split(s, " ");

	// save the words in reverse order
	for i := range(words) {
		reverse = append(reverse, words[len(words) - 1 - i])
	}

	// return the reverse words sentence as string
	return strings.Join(reverse, " ")
}

func main() {
	var sentence string

	// add a flag for receive the sentence by command line
	flag.StringVar(&sentence, "s", "", "input sentence")
	flag.Parse()

	// print a message and finish if the sentence is not defined
	if sentence == "" {
		fmt.Println("Please write a string sentence with the `-s` flag")
		return
	}

	// run the function
	reverse := ReverseWords(sentence)
	fmt.Println(reverse)
}
