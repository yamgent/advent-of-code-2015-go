package main

import (
	"aoc2015/utils"
	"testing"
)

func TestRunPart1(t *testing.T) {
	cases := []utils.TestCase{
		{In: "(())", Out: "0"},
		{In: "()()", Out: "0"},
		{In: "(((", Out: "3"},
		{In: "(()(()(", Out: "3"},
		{In: "))(((((", Out: "3"},
		{In: "())", Out: "-1"},
		{In: "))(", Out: "-1"},
		{In: ")))", Out: "-3"},
		{In: ")())())", Out: "-3"},
	}

	for _, v := range cases {
		utils.AssertTestCase(t, v, runPart1(v.In))
	}
}

func TestRunPart2(t *testing.T) {
	cases := []utils.TestCase{
		{In: ")", Out: "1"},
		{In: "()())", Out: "5"},
	}

	for _, v := range cases {
		utils.AssertTestCase(t, v, runPart2(v.In))
	}
}
