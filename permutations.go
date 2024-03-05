package main

import (
	"math/rand"
	"sort"
	"strings"
)

func Permutations(input string) []string {
	result := shufflings(input, []string{})
	sort.Strings(result)

	return result
}

func shufflings(text string, result []string) []string {
	texts := strings.Split(text, "")
	length := len(texts)
	if length <= 1 {
		result = append(result, text)
		return result
	}

	for i := 0; i < length; i++ {
		j := rand.Intn(length)
		texts[i], texts[j] = string(texts[j]), string(texts[i])
		if !checkDup(strings.Join(texts, ""), result) {
			result = append(result, strings.Join(texts, ""))
		}
		if len(result) == permutationsWithRepetition(text) {
			break
		}
	}

	if len(result) != permutationsWithRepetition(text) {
		result = shufflings(text, result)
	}

	return result
}

func checkDup(text string, texts []string) bool {
	for i := 0; i < len(texts); i++ {
		if text == texts[i] {
			return true
		}
	}
	return false
}

func sumMultiply(n int) int {
	sum := 1
	for i := 0; i < n; i++ {
		sum *= i + 1
	}
	return sum
}

// research
func permutationsWithRepetition(str string) int {
	n := len(str)

	frequencies := make(map[rune]int)
	for _, char := range str {
		frequencies[char]++
	}

	permutations := sumMultiply(n)
	for _, count := range frequencies {
		permutations /= sumMultiply(count)
	}

	return permutations
}
