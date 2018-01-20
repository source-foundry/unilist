// unilist is a command line executable that generates lists of Unicode code points
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	version = "0.2.0"

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
   Write a list of Unicode code point values delimited by space glyphs

 Options:
  -c, --comma          Use comma delimiter
  -h, --help           Application help
  -n, --newline        Use newline delimiter
      --usage          Application usage
  -v, --version        Application version

 Notes:
   Define arg values as single Unicode hexadecimal code point
   values or as ranges that are indicated with a '-' between 
   values (e.g., 0020 or 0020-0024). 

   Exclude hexadecimal values from output by using a 
   case-insensitive 'x' as the first character of the single
   Unicode code point arg or range arg (e.g., x0020 or x0020-0024)
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
		fmt.Print(getCommaDelimitedString(keepUniList))
	} else if *useNewLinesShort || *useNewLinesLong {
		fmt.Print(getNewlineDelimitedString(keepUniList))
	} else {
		fmt.Print(getSpaceDelimitedString(keepUniList))
	}
}
