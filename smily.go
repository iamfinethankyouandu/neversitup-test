package main

import (
	"strings"
)

func CountSmilyFace(text []string) int {
	result := 0
	for _, v := range text {
		result += inSmilyList(strings.Split(v, ""))

	}

	return result
}

func inSmilyList(faces []string) int {
	countFace := len(faces)
	match := 0
	lists := []string{":", ")", "D", ";", "-", "~"}

	for _, face := range faces {
		for _, list := range lists {
			if face == list {
				match++
				break
			}
		}
	}

	if match == countFace {
		return 1
	}

	return 0
}
