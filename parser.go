package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getUnicodeValueSlice(arg string) ([]string, error) {
	var returnSlice []string
	// strip the case insensitive 'x' value from exclude request arguments (detected in calling code)
	if strings.HasPrefix(arg, "x") || strings.HasPrefix(arg, "X") {
		arg = arg[1:]
	}
	if strings.Contains(arg, "-") {
		tempUniList := strings.Split(arg, "-")
		startInt, startParseErr := strconv.ParseInt(tempUniList[0], 16, 32)
		endInt, endParseErr := strconv.ParseInt(tempUniList[1], 16, 32)
		if startParseErr != nil {
			return returnSlice, startParseErr
		}
		if endParseErr != nil {
			return returnSlice, endParseErr
		}
		for i := startInt; i <= endInt; i++ {
			returnSlice = append(returnSlice, fmt.Sprintf("%U", i))
		}
	} else {
		if len(arg) > 3 {
			anInt, parseErr := strconv.ParseInt(arg, 16, 32)
			if parseErr != nil {
				return returnSlice, parseErr
			}
			returnSlice = append(returnSlice, fmt.Sprintf("%U", anInt))

		} else {
			// handle argument that is not formatted as unicode hexadecimal value
			lenErr := fmt.Errorf("the argment '%s' is not properly formatted as a Unicode hexadecimal value", arg)
			return returnSlice, lenErr
		}
	}
	return returnSlice, nil
}
