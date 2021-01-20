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
