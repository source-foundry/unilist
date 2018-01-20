// unilist is a command line executable that generates lists of Unicode code points
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	version = "0.1.0"

	usage = `Usage: unilist (options) [arg 1]...[arg n]
`

	help = `=================================================
 unilist
 Copyright 2018 Christopher Simpkins
 MIT License

 Source: https://github.com/source-foundry/unilist
=================================================
 Usage:
   $ unilist (options) [arg 1]...[arg n]

 Default:
   Write a space glyph delimited list of Unicode code point values

 Options:
  -c, --comma          Use comma delimiter
  -h, --help           Application help
  -n, --newline        Use newline delimiter
      --usage          Application usage
  -v, --version        Application version

 Notes:
   arg values should be single Unicode hexadecimal code point values,
   or ranges that are indicated with a '-' between values. Exclude
   hexadecimal values from output by using a case-insensitive 'x'
   as the first character of the value or the range (e.g., x0020)
`
)

var versionShort, versionLong, helpShort, helpLong, usageLong *bool
var useCommasLong, useCommasShort, useNewLinesLong, useNewLinesShort *bool

func init() {
	// define available command line flags
	versionShort = flag.Bool("v", false, "Application version")
	versionLong = flag.Bool("version", false, "Application version")
	helpShort = flag.Bool("h", false, "Help")
	helpLong = flag.Bool("help", false, "Help")
	usageLong = flag.Bool("usage", false, "Usage")
	useCommasLong = flag.Bool("comma", false, "Comma delimiters in lists")
	useCommasShort = flag.Bool("c", false, "Comma delimiters in lists")
	useNewLinesLong = flag.Bool("newline", false, "Newline delimiters in lists")
	useNewLinesShort = flag.Bool("n", false, "Newline delimiters in lists")
}

func main() {
	flag.Parse()

	// parse command line flags and handle them
	switch {
	case *versionShort, *versionLong:
		fmt.Printf("unilist v%s\n", version)
		os.Exit(0)
	case *helpShort, *helpLong:
		fmt.Println(help)
		os.Exit(0)
	case *usageLong:
		fmt.Println(usage)
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		os.Stderr.WriteString("[Error] Please include at least one Unicode code point range argument in your command.\n")
		os.Exit(1)
	}

	var keepUniList []string
	var removeUniList []string

	// create exclude unicode value slice
	for _, remArg := range flag.Args() {
		if strings.HasPrefix(remArg, "x") || strings.HasPrefix(remArg, "X") {
			tempRemoveSlice, err := getUnicodeValueSlice(remArg)
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("[Error] %v", err))
				os.Exit(1)
			}

			removeUniList = append(removeUniList, tempRemoveSlice...)
		}
	}

	for _, keepArg := range flag.Args() {
		if !strings.HasPrefix(keepArg, "x") || strings.HasPrefix(keepArg, "X") {
			tempKeepSlice, err := getUnicodeValueSlice(keepArg)
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("[Error] %v", err))
				os.Exit(1)
			}

			for _, keepValue := range tempKeepSlice {
				if !containsString(removeUniList, keepValue) {
					// only insert into the keep values slice if it is not in the exclude values slice
					keepUniList = append(keepUniList, keepValue)
				}
			}
		}
	}

	// format the standard out string with proper default or command line defined delimiter
	if *useCommasShort || *useCommasLong {
		fmt.Print(strings.Join(keepUniList, ","))
	} else if *useNewLinesShort || *useNewLinesLong {
		fmt.Print(strings.Join(keepUniList, "\n"))
	} else {
		fmt.Print(strings.Join(keepUniList, " "))
	}
}

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

func containsString(haystack []string, needle string) bool {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}
