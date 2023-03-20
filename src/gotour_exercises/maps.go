package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	counts := make(map[string]int)

	for _, word := range words {
		counts[word] += 1
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}