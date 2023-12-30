package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortByAAndLength(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		countA_i := strings.Count(words[i], "a")
		countA_j := strings.Count(words[j], "a")

		if countA_i == countA_j {
			return len(words[i]) > len(words[j])
		}

		return countA_i > countA_j
	})

	return words
}

func main() {
	input := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "anneanne", "kf", "zc", "lklklklklklklklkl", "l"}

	fmt.Println("Input:")
	fmt.Println(input)
	output := sortByAAndLength(input)
	fmt.Println("Output:")
	fmt.Println(output)
}
