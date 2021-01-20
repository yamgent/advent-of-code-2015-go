package main

import (
	"aoc2015/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func runPart1(input string) string {
	lines := strings.Split(input, "\n")

	total := 0
	for _, v := range lines {
		if v == "" {
			break
		}

		lengths := strings.Split(v, "x")
		l, _ := strconv.Atoi(lengths[0])
		w, _ := strconv.Atoi(lengths[1])
		h, _ := strconv.Atoi(lengths[2])

		areas := []int{l * w, l * h, w * h}

		slack := areas[0]
		for _, a := range areas {
			if a < slack {
				slack = a
			}
		}

		total += 2*(areas[0]+areas[1]+areas[2]) + slack
	}

	return fmt.Sprint(total)
}

func runPart2(input string) string {
	lines := strings.Split(input, "\n")

	total := 0
	for _, v := range lines {
		if v == "" {
			break
		}

		lengthsStr := strings.Split(v, "x")
		lengths := []int{}
		for _, v := range lengthsStr {
			num, _ := strconv.Atoi(v)
			lengths = append(lengths, num)
		}

		sort.Ints(lengths)
		total += 2*(lengths[0]+lengths[1]) + (lengths[0] * lengths[1] * lengths[2])
	}

	return fmt.Sprint(total)
}

func main() {
	input := utils.ReadFromFile("input.txt")

	fmt.Println(runPart1(input))
	fmt.Println(runPart2(input))
}
