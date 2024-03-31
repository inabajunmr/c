package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Tokens struct {
	tokens []Token
	index  int
}

func (ts *Tokens) ConsumeNumberMust() *NumberToken {
	if number, ok := ts.tokens[ts.index].(NumberToken); ok {
		ts.index++
		return &number
	}
	log.Fatal("not number")
	return nil
}

func (ts *Tokens) ConsumeLParenthesisToken() bool {
	if _, ok := ts.tokens[ts.index].(LParenthesisToken); ok {
		ts.index++
		return true
	}
	return false
}

func (ts *Tokens) ConsumeRParenthesisTokenMust() bool {
	if _, ok := ts.tokens[ts.index].(RParenthesisToken); ok {
		ts.index++
		return true
	}
	log.Fatal("not )")
	return false
}

func (ts *Tokens) ConsumeMultiplicationToken() bool {
	if _, ok := ts.tokens[ts.index].(MultiplicationToken); ok {
		ts.index++
		return true
	}
	return false
}

func (ts *Tokens) ConsumeDivisionToken() bool {
	if _, ok := ts.tokens[ts.index].(DivisionToken); ok {
		ts.index++
		return true
	}
	return false
}

func (ts *Tokens) ConsumePlusToken() bool {
	if _, ok := ts.tokens[ts.index].(PlusToken); ok {
		ts.index++
		return true
	}
	return false
}

func (ts *Tokens) ConsumeMinusToken() bool {
	if _, ok := ts.tokens[ts.index].(MinusToken); ok {
		ts.index++
		return true
	}
	return false
}

type Token interface{}

type PlusToken struct{}
type MinusToken struct{}
type MultiplicationToken struct{}
type DivisionToken struct{}

type NumberToken struct {
	Value int
}

type LParenthesisToken struct{}
type RParenthesisToken struct{}

type EOFToken struct{}

func Tokenize(source string) *Tokens {
	tokens := Tokens{}

	s := strings.Split(source, "")
	len := len(s)
	for i := 0; i < len; i++ {
		if s[i] == "+" {
			tokens.tokens = append(tokens.tokens, PlusToken{})
		} else if s[i] == "-" {
			tokens.tokens = append(tokens.tokens, MinusToken{})
		} else if s[i] == "*" {
			tokens.tokens = append(tokens.tokens, MultiplicationToken{})
		} else if s[i] == "/" {
			tokens.tokens = append(tokens.tokens, DivisionToken{})
		} else if unicode.IsDigit([]rune(s[i])[0]) {
			last := lastIndex(source[i:], "[0-9]+")
			n, _ := strconv.Atoi(source[i : last+i+1])
			tokens.tokens = append(tokens.tokens, NumberToken{Value: n})
			i = i + last
		} else if s[i] == "(" {
			tokens.tokens = append(tokens.tokens, LParenthesisToken{})
		} else if s[i] == ")" {
			tokens.tokens = append(tokens.tokens, RParenthesisToken{})
		}
	}
	tokens.tokens = append(tokens.tokens, EOFToken{})
	return &tokens
}

func lastIndex(s string, pattern string) int {
	re := regexp.MustCompile(pattern)
	match := re.FindString(s)
	if match == "" {
		return 0
	}
	return len(match) - 1
}
