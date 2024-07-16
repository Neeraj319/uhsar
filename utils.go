package main

import "fmt"

func Error(line int, message string) {
	report(line, "", message)

}
func report(line int, where string, message string) {
	error_string := fmt.Errorf("line %d Error %s : %s ", line, where, message)
	fmt.Println(error_string)

}

func IsDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func IsAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func IsAlphaNumeric(c rune) bool {
	return IsAlpha(c) || IsDigit(c)

}
