package main

import (
	"aoc2015/utils"
	"strings"
	"testing"
)

func TestRunPart1(t *testing.T) {
	cases := []utils.TestCase{
		{In: "ugknbfddgicrmopn", Out: "1"},
		{In: "aaa", Out: "1"},
		{In: "jchzalrnumimnmhp", Out: "0"},
		{In: "haegwjzuvuyypxyu", Out: "0"},
		{In: "dvszwmarrgswjxmb", Out: "0"},
	}

	for _, v := range cases {
		utils.AssertTestCase(t, v, runPart1(v.In))
	}
}

func TestRunPart2(t *testing.T) {
	cases := []utils.TestCase{
		{In: "qjhvhtzxzqqjkmpb", Out: "1"},
		{In: "xxyxx", Out: "1"},
		{In: "uurcxstgmygtbstg", Out: "0"},
		{In: "ieodomkazucvgmuy", Out: "0"},
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
