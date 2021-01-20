package main

import (
	"aoc2015/utils"
	"testing"
)

func TestRunPart1(t *testing.T) {
	cases := map[string]string{
		"(())":    "0",
		"()()":    "0",
		"(((":     "3",
		"(()(()(": "3",
		"))(((((": "3",
		"())":     "-1",
		"))(":     "-1",
		")))":     "-3",
		")())())": "-3",
	}

	for k, v := range cases {
		utils.Assert(t, k, v, runPart1(k))
	}
}

func TestRunPart2(t *testing.T) {
	cases := map[string]string{
		")":     "1",
		"()())": "5",
	}

	for k, v := range cases {
		utils.Assert(t, k, v, runPart2(k))
	}
}
