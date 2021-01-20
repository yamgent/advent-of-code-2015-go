package main

import (
	"aoc2015/utils"
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
