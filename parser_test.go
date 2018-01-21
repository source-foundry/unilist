package main

import (
	"testing"
)

//
//  INCLUDE VALUE TESTS
//

func TestGetUnicodeValueSingleValidIncludeValue(t *testing.T) {
	testArg := "0020"
	resSlice, err := getUnicodeValueSlice(testArg)
	if err != nil {
		t.Errorf("[FAIL] Expected no error and received error %v", err)
	}
	if len(resSlice) != 1 {
		t.Errorf("[FAIL] Expected one return value in the slice, received %d", len(resSlice))
	}
	if resSlice[0] != "U+0020" {
		t.Errorf("[FAIL] Did not expect returned value %s", resSlice[0])
	}
}

func TestGetUnicodeValueRangeValidIncludeValues(t *testing.T) {
	testArg := "0020-0022"
	resSlice, err := getUnicodeValueSlice(testArg)
	if err != nil {
		t.Errorf("[FAIL] Expected no error and received error %v", err)
	}
	if len(resSlice) != 3 {
		t.Errorf("[FAIL] Expected one return value in the slice, received %d", len(resSlice))
	}
	if resSlice[0] != "U+0020" {
		t.Errorf("[FAIL] Did not expect returned value %s", resSlice[0])
	}
	if resSlice[1] != "U+0021" {
		t.Errorf("[FAIL] Did not expect returned value %s", resSlice[0])
	}
	if resSlice[2] != "U+0022" {
		t.Errorf("[FAIL] Did not expect returned value %s", resSlice[0])
	}
}

func TestGetUnicodeValueSingleInvalidIncludeValueNotHexadecimal(t *testing.T) {
	testArg := "ZZZZ"
	_, err := getUnicodeValueSlice(testArg)
	if err == nil {
		t.Errorf("[FAIL] Expected error for invalid hexadecimal string and received no error")
	}
}

func TestGetUnicodeValueSingleInvalidIncludeIncorrectLengthHexadecimal(t *testing.T) {
	testArg := "20"
	_, err := getUnicodeValueSlice(testArg)
	if err == nil {
		t.Errorf("[FAIL] Expected error for invalid length hexadecimal string and received no error")
	}
}

func TestGetUnicodeValueRangeInvalidIncludeFirstValueNotHexadecimal(t *testing.T) {
	testArg := "ZZZZ-0020"
	_, err := getUnicodeValueSlice(testArg)
	if err == nil {
		t.Errorf("[FAIL] Expected error for invalid hexadecimal string and received no error")
	}
}

func TestGetUnicodeValueRangeInvalidIncludeLastValueNotHexadecimal(t *testing.T) {
	testArg := "0020-ZZZZ"
	_, err := getUnicodeValueSlice(testArg)
	if err == nil {
		t.Errorf("[FAIL] Expected error for invalid hexadecimal string and received no error")
	}
}


func TestGetUnicodeValueRangeInvalidIncludeTooShortStart(t *testing.T) {
	testArg := "20-0030"
	_, err := getUnicodeValueSlice(testArg)
	if err == nil {
		t.Errorf("[FAIL] Expected error for invalid hexadecimal string and received no error")
	}
}


func TestGetUnicodeValueRangeInvalidIncludeTooShortEnd(t *testing.T) {
	testArg := "0020-30"
	_, err := getUnicodeValueSlice(testArg)
	if err == nil {
		t.Errorf("[FAIL] Expected error for invalid hexadecimal string and received no error")
	}
}

//
// EXCLUDE VALUE TESTS
//

func TestGetUnicodeValueSingleValidExcludeValue(t *testing.T) {
	testArg := "x0020"
	resSlice, err := getUnicodeValueSlice(testArg)
	if err != nil {
		t.Errorf("[FAIL] Expected no error and received error %v", err)
	}
	if len(resSlice) != 1 {
		t.Errorf("[FAIL] Expected one return value in the slice, received %d", len(resSlice))
	}
	if resSlice[0] != "U+0020" {
		t.Errorf("[FAIL] Did not expect returned value %s", resSlice[0])
	}
}

func TestGetUnicodeValueRangeValidExcludeValues(t *testing.T) {
	testArg := "x0020-0022"
	resSlice, err := getUnicodeValueSlice(testArg)
	if err != nil {
		t.Errorf("[FAIL] Expected no error and received error %v", err)
	}
	if len(resSlice) != 3 {
		t.Errorf("[FAIL] Expected one return value in the slice, received %d", len(resSlice))
	}
	if resSlice[0] != "U+0020" {
		t.Errorf("[FAIL] Did not expect returned value %s", resSlice[0])
	}
	if resSlice[1] != "U+0021" {
		t.Errorf("[FAIL] Did not expect returned value %s", resSlice[0])
	}
	if resSlice[2] != "U+0022" {
		t.Errorf("[FAIL] Did not expect returned value %s", resSlice[0])
	}
}

func TestGetUnicodeValueSingleInvalidExcludeValueNotHexadecimal(t *testing.T) {
	testArg := "xZZZZ"
	_, err := getUnicodeValueSlice(testArg)
	if err == nil {
		t.Errorf("[FAIL] Expected error for invalid hexadecimal string and received no error")
	}
}

func TestGetUnicodeValueSingleInvalidExcludeIncorrectLengthHexadecimal(t *testing.T) {
	testArg := "x20"
	_, err := getUnicodeValueSlice(testArg)
	if err == nil {
		t.Errorf("[FAIL] Expected error for invalid length hexadecimal string and received no error")
	}
}

func TestGetUnicodeValueRangeInvalidExcludeFirstValueNotHexadecimal(t *testing.T) {
	testArg := "xZZZZ-0020"
	_, err := getUnicodeValueSlice(testArg)
	if err == nil {
		t.Errorf("[FAIL] Expected error for invalid hexadecimal string and received no error")
	}
}

func TestGetUnicodeValueRangeInvalidExcludeLastValueNotHexadecimal(t *testing.T) {
	testArg := "x0020-ZZZZ"
	_, err := getUnicodeValueSlice(testArg)
	if err == nil {
		t.Errorf("[FAIL] Expected error for invalid hexadecimal string and received no error")
	}
}
