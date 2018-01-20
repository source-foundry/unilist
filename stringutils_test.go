package main

import (
	"testing"
)

func TestContainsStringValid(t *testing.T) {
	testSlice := []string{"test", "another", "last"}
	if !containsString(testSlice, "test") {
		t.Errorf("[FAIL] Expected test for string that is present to return true and it returned false")
	}
	if !containsString(testSlice, "another") {
		t.Errorf("[FAIL] Expected test for string that is present to return true and it returned false")
	}
	if !containsString(testSlice, "last") {
		t.Errorf("[FAIL] Expected test for string that is present to return true and it returned false")
	}
}

func TestContainsStringInvalid(t *testing.T) {
	testSlice := []string{"test", "another", "last"}
	if containsString(testSlice, "bogus") {
		t.Errorf("[FAIL] Expected test for string that is not present to return false and it returned true")
	}
}

func TestGetCommaDelimitedString(t *testing.T) {
	testSlice := []string{"test", "another", "last"}
	responseString := getCommaDelimitedString(testSlice)
	if responseString != "test,another,last" {
		t.Errorf("[FAIL] Expected test to return comma delimited string and instead returned '%s'", responseString)
	}
}

func TestGetNewlineDelimitedString(t *testing.T) {
	testSlice := []string{"test", "another", "last"}
	responseString := getNewlineDelimitedString(testSlice)
	if responseString != "test\nanother\nlast" {
		t.Errorf("[FAIL] Expected test to return newline delimited string and instead returned '%s'", responseString)
	}
}

func TestGetSpaceDelimitedString(t *testing.T) {
	testSlice := []string{"test", "another", "last"}
	responseString := getSpaceDelimitedString(testSlice)
	if responseString != "test another last" {
		t.Errorf("[FAIL] Expected test to return space delimited string and instead returned '%s'", responseString)
	}
}
