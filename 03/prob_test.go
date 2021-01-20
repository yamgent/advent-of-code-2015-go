package main

import (
	"aoc2015/utils"
	"strings"
	"testing"
)

func TestRunPart1(t *testing.T) {
	cases := []utils.TestCase{
		{In: ">", Out: "2"},
		{In: "^>v<", Out: "4"},
		{In: "^v^v^v^v^v", Out: "2"},
	}

	for _, v := range cases {
		utils.AssertTestCase(t, v, runPart1(v.In))
	}
}

func TestRunPart2(t *testing.T) {
	cases := []utils.TestCase{
		{In: "^v", Out: "3"},
		{In: "^>v<", Out: "3"},
		{In: "^v^v^v^v^v", Out: "11"},
	}

	for _, v := range cases {
		utils.AssertTestCase(t, v, runPart2(v.In))
	}
}

func TestRunInput(t *testing.T) {
	input := utils.ReadFromFile("input.txt")
	outputs := strings.Split(utils.ReadFromFile("output.txt"), "\n")

	utils.Assert(t, "Actual input part 1", outputs[0], runPart1(input))
	utils.Assert(t, "Actual input part 2", outputs[1], runPart2(input))
}
