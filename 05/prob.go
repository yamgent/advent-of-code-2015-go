package main

import (
	"aoc2015/utils"
	"fmt"
	"strings"
)

func runPart1(input string) string {
	req1 := func(line string) bool {
		vowels := 0
		for _, v := range line {
			switch v {
			case 'a':
				vowels++
			case 'e':
				vowels++
			case 'i':
				vowels++
			case 'o':
				vowels++
			case 'u':
				vowels++
			}
		}
		return vowels >= 3
	}
	req2 := func(line string) bool {
		for i := 0; i < len(line)-1; i++ {
			if line[i] == line[i+1] {
				return true
			}
		}
		return false
	}
	req3 := func(line string) bool {
		forbidden := [...]string{"ab", "cd", "pq", "xy"}
		for _, v := range forbidden {
			if strings.Index(line, v) != -1 {
				return false
			}
		}
		return true
	}

	total := 0
	lines := strings.Split(input, "\n")
	for _, v := range lines {
		if req1(v) && req2(v) && req3(v) {
			total++
		}
	}

	return fmt.Sprint(total)
}

func runPart2(input string) string {
	req1 := func(line string) bool {
		seen := make(map[string]int)

		for i := 0; i < len(line)-1; i++ {
			substr := line[i : i+2]
			lastpos, ok := seen[substr]
			if !ok {
				seen[substr] = i
			} else if lastpos+1 < i {
				return true
			}
		}

		return false
	}
	req2 := func(line string) bool {
		for i := 0; i < len(line)-2; i++ {
			if line[i] == line[i+2] {
				return true
			}
		}
		return false
	}

	total := 0
	lines := strings.Split(input, "\n")
	for _, v := range lines {
		if req1(v) && req2(v) {
			total++
		}
	}

	return fmt.Sprint(total)
}

func main() {
	input := utils.ReadFromFile("input.txt")

	fmt.Println(runPart1(input))
	fmt.Println(runPart2(input))
}
