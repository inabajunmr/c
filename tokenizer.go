package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type TokenType int

const (
	OPERATOR TokenType = iota
	NUMBER
)

type Token interface{}

type OperatorToken struct {
	Value string
}

type NumberToken struct {
	Value int
}

func Tokenize(source string) []Token {
	var tokens []Token

	s := strings.Split(source, "")
	len := len(s)
	for i := 0; i < len; i++ {
		if isOperator(s[i]) {
			tokens = append(tokens, OperatorToken{Value: s[i]})
		} else if unicode.IsDigit([]rune(s[i])[0]) {
			last := lastIndex(source[i:], "[0-9]+")
			n, _ := strconv.Atoi(source[i : last+i+1])
			tokens = append(tokens, NumberToken{Value: n})
			i = i + last
		}
	}
	return tokens
}

func isOperator(v string) bool {
	return v == "-" || v == "+"
}

func lastIndex(s string, pattern string) int {
	re := regexp.MustCompile(pattern)
	match := re.FindString(s)
	if match == "" {
		return 0
	}
	return len(match) - 1
}
