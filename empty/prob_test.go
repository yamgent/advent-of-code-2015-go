package main

import (
	"aoc2015/utils"
	"testing"
)

func TestRunPart1(t *testing.T) {
	cases := []utils.TestCase{
		{In: "None", Out: ""},
	}

	for _, v := range cases {
		utils.AssertTestCase(t, v, runPart1(v.In))
	}
}

func TestRunPart2(t *testing.T) {
	cases := []utils.TestCase{
		{In: "None", Out: ""},
	}

	for _, v := range cases {
		utils.AssertTestCase(t, v, runPart2(v.In))
	}
}
