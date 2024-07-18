package main

import (
	"unicode/utf8"
)

func StringLength(s string) int {
	return utf8.RuneCountInString(s)
}