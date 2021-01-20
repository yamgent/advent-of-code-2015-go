package utils

import (
	"fmt"
	"io/ioutil"
	"testing"
)

// ReadFromFile into a string, panics if file does not exist
func ReadFromFile(filepath string) string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(fmt.Sprintf("Cannot read file '%v'", filepath))
	}
	return string(content)
}

// Assert compare expected to actual, fails fatally if doesn't match
func Assert(t *testing.T, testname string, expected string, actual string) {
	if expected != actual {
		t.Fatalf("%v fail, expected '%v', got '%v'", testname, expected, actual)
	}
}

// TestCase represents a test case with an input and expected output,
// and optionally a Testname (if blank, use In as Testname instead).
type TestCase struct {
	In       string
	Out      string
	Testname string
}

// AssertTestCase same as Assert, but accepts a TestCase for input
// instead (which will also intelligently determine test name - if
// Testname is omitted, we take In as the name instead)
func AssertTestCase(t *testing.T, testCase TestCase, actual string) {
	name := testCase.Testname
	if name == "" {
		name = testCase.In
	}
	Assert(t, name, testCase.Out, actual)
}
