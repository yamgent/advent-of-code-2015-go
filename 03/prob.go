package main

import (
	"aoc2015/utils"
	"fmt"
)

type point struct {
	X int
	Y int
}

func runPart1(input string) string {
	visited := map[point]int{
		{0, 0}: 1,
	}

	currentPos := point{0, 0}

	for _, v := range input {
		switch v {
		case '>':
			currentPos.X++
		case '<':
			currentPos.X--
		case '^':
			currentPos.Y++
		case 'v':
			currentPos.Y--
		}

		visited[currentPos] = 1
	}

	return fmt.Sprint(len(visited))
}

func runPart2(input string) string {
	visited := map[point]int{
		{0, 0}: 1,
	}

	santa := point{0, 0}
	robo := point{0, 0}

	for i, v := range input {
		current := &santa
		if i%2 == 1 {
			current = &robo
		}

		switch v {
		case '>':
			current.X++
		case '<':
			current.X--
		case '^':
			current.Y++
		case 'v':
			current.Y--
		}

		visited[*current] = 1
	}

	return fmt.Sprint(len(visited))
}

func main() {
	input := utils.ReadFromFile("input.txt")

	fmt.Println(runPart1(input))
	fmt.Println(runPart2(input))
}
