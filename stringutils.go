package main

import "strings"

func containsString(haystack []string, needle string) bool {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}

func getCommaDelimitedString(list []string) string {
	return strings.Join(list, ",")
}

func getNewlineDelimitedString(list []string) string {
	return strings.Join(list, "\n")
}

func getSpaceDelimitedString(list []string) string {
	return strings.Join(list, " ")
}
