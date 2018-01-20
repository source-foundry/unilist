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
		startInt, startErr := strconv.ParseInt(tempUniList[0], 16, 32)
		endInt, endErr := strconv.ParseInt(tempUniList[1], 16, 32)
		if startErr != nil {
			return returnSlice, startErr
		}
		if endErr != nil {
			return returnSlice, endErr
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
