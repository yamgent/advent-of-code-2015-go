package main

import (
	"aoc2015/utils"
	"testing"
)

func TestRunPart1(t *testing.T) {
	cases := []utils.TestCase{
		{In: "abcdef", Out: "609043"},
		{In: "pqrstuv", Out: "1048970"},
	}

	for _, v := range cases {
		utils.AssertTestCase(t, v, runPart1(v.In))
	}
}

// func TestRunPart2(t *testing.T) {
// 	cases := []utils.TestCase{
// 		{In: "None", Out: ""},
// 	}

// 	for _, v := range cases {
// 		utils.AssertTestCase(t, v, runPart2(v.In))
// 	}
// }

func TestRunInput(t *testing.T) {
	input := "iwrupvqb"
	outputs := []string{"346386", "9958218"}

	utils.Assert(t, "Actual input part 1", outputs[0], runPart1(input))
	utils.Assert(t, "Actual input part 2", outputs[1], runPart2(input))
}
