package main

import (
	"aoc2015/utils"
	"strings"
	"testing"
)

func TestRunPart1(t *testing.T) {
	cases := []utils.TestCase{
		{In: "2x3x4", Out: "58"},
		{In: "1x1x10", Out: "43"},
	}

	for _, v := range cases {
		utils.AssertTestCase(t, v, runPart1(v.In))
	}
}

func TestRunPart2(t *testing.T) {
	cases := []utils.TestCase{
		{In: "2x3x4", Out: "34"},
		{In: "1x1x10", Out: "14"},
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
