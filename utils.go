package main

import "fmt"

func Error(line int, message string) error {
	return report(line, "", message)

}
func report(line int, where string, message string) error {
	error_string := fmt.Errorf("line %d Error %s : %s ", line, where, message)
	return (error_string)
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

func ParsingError(token Token, message string) error {
	if token.Type == EOF {
		return report(token.Line, " at end", message)
	} else {
		return report(token.Line, " at '"+token.Lexeme+"'", message)
	}
}

func IsTruthy(object interface{}) bool {
	if object == nil {
		return false
	}
	value, ok := object.(bool)
	if !(ok) {
		return true
	}
	return value
}

func IsEqual(a interface{}, b interface{}) bool {
	return a == b

}
