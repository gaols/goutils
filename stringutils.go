package goutils

import "strings"

// IsBlank Checks if a string is whitespace, empty ("").
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// IsBlank Checks if a string is not empty (""), not null and not whitespace only.
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// IsEmpty Checks if a string is whitespace, empty ("").
func IsEmpty(str string) bool {
	return len(str) == 0
}

// IsNotEmpty Checks if a string is not empty ("").
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// Trim returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.
func Trim(str string) string {
	return strings.TrimSpace(str)
}

// DefaultIfEmpty Returns either the passed in string, or if the string is
// empty (""), the value of default string.
func DefaultIfEmpty(str, defaultStr string) string {
	if IsEmpty(str) {
		return defaultStr
	}

	return str
}

// DefaultIfBlank Returns either the passed in string, or if the string is
// whitespace, empty (""), the value of default string.
func DefaultIfBlank(str, defaultStr string) string {
	if IsBlank(str) {
		return defaultStr
	}

	return str
}
