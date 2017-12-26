package goutils

import (
	"strings"
	"unicode/utf8"
)

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

// Left pad a String with a specified character.
// WARNING: string to pad should be uft8-encoded!
//
// goutils.LeftPad("", 3, 'z')     = "zzz"
// goutils.LeftPad("bat", 3, 'z')  = "bat"
// goutils.LeftPad("bat", 5, 'z')  = "zzbat"
// goutils.LeftPad("bat", 1, 'z')  = "bat"
// goutils.LeftPad("bat", -1, 'z') = "bat"
func LeftPad(str string, size int, padChar rune) string {
	return calcPadStr(str, size, padChar) + str
}

// RightPad right pad a String with a specified character.
// WARNING: string to pad should be uft8-encoded!

// goutils.RightPad("", 3, 'z')     = "zzz"
// goutils.RightPad("bat", 3, 'z')  = "bat"
// goutils.RightPad("bat", 5, 'z')  = "batzz"
// goutils.RightPad("bat", 1, 'z')  = "bat"
// goutils.RightPad("bat", -1, 'z') = "bat"
func RightPad(str string, size int, padChar rune) string {
	return str + calcPadStr(str, size, padChar)
}

func calcPadStr(str string, size int, padChar rune) string {
	count := utf8.RuneCountInString(str)
	if size <= count {
		return ""
	}
	return strings.Repeat(string(padChar), size - count)
}
