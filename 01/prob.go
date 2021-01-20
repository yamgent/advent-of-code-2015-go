package main

import (
	"aoc2015/utils"
	"fmt"
	"strings"
)

func runPart1(input string) string {
	input = strings.TrimSpace(input)

	floor := 0
	for _, v := range input {
		if v == '(' {
			floor++
		} else {
			floor--
		}
	}

	return fmt.Sprint(floor)
}

func runPart2(input string) string {
	floor := 0
	for i, v := range input {
		if v == '(' {
			floor++
		} else {
			floor--
		}

		if floor < 0 {
			return fmt.Sprint(i + 1)
		}
	}
	return "Impossible"
}

func main() {
	input := utils.ReadFromFile("input.txt")

	fmt.Println(runPart1(input))
	fmt.Println(runPart2(input))
}
