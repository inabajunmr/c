package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Tokens struct {
	tokens []Token
	index  int
}

func (ts *Tokens) ConsumeIdent() *IdentToken {
	if ident, ok := ts.tokens[ts.index].(IdentToken); ok {
		ts.index++
		return &ident
	}
	return nil
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
	log.Fatalf("not ):%+v", ts.tokens[ts.index])
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

func (ts *Tokens) ConsumeGreaterToken() bool {
	if _, ok := ts.tokens[ts.index].(GreaterToken); ok {
		ts.index++
		return true
	}
	return false
}

func (ts *Tokens) ConsumeGreaterThanEqualToken() bool {
	if _, ok := ts.tokens[ts.index].(GreaterThanEqualToken); ok {
		ts.index++
		return true
	}
	return false
}

func (ts *Tokens) ConsumeLessToken() bool {
	if _, ok := ts.tokens[ts.index].(LessToken); ok {
		ts.index++
		return true
	}
	return false
}

func (ts *Tokens) ConsumeLessThanToken() bool {
	if _, ok := ts.tokens[ts.index].(LessThanEqualToken); ok {
		ts.index++
		return true
	}
	return false
}

func (ts *Tokens) ConsumeSemicolonTokenMust() bool {
	if _, ok := ts.tokens[ts.index].(SemicolonToken); ok {
		ts.index++
		return true
	}
	log.Fatal("not ;")
	return false
}

func (ts *Tokens) ConsumeAssignToken() bool {
	if _, ok := ts.tokens[ts.index].(AssignToken); ok {
		ts.index++
		return true
	}
	return false
}

func (ts *Tokens) ConsumeEqualToken() bool {
	if _, ok := ts.tokens[ts.index].(EqualToken); ok {
		ts.index++
		return true
	}
	return false
}

func (ts *Tokens) ConsumeNotEqualToken() bool {
	if _, ok := ts.tokens[ts.index].(NotEqualToken); ok {
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

type LessToken struct{}
type LessThanEqualToken struct{}
type GreaterToken struct{}
type GreaterThanEqualToken struct{}
type EqualToken struct{}
type NotEqualToken struct{}

type NumberToken struct {
	Value int
}

type IdentToken struct {
	Value string
}

type AssignToken struct{}

type LParenthesisToken struct{}
type RParenthesisToken struct{}

type SemicolonToken struct{}

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
		} else if s[i] == "=" && !isComparisonOperatorSymbol(s[i+1]) {
			tokens.tokens = append(tokens.tokens, AssignToken{})
		} else if isComparisonOperatorSymbol(s[i]) {
			if isComparisonOperatorSymbol(s[i+1]) {
				tokens.tokens = append(tokens.tokens, mapComparisonOperator(source[i:i+2]))
				i = i + 1
			} else {
				tokens.tokens = append(tokens.tokens, mapComparisonOperator(s[i]))
			}
		} else if s[i] == "!" && s[i+1] == "=" {
			tokens.tokens = append(tokens.tokens, NotEqualToken{})
			i = i + 1
		} else if isIdent(s[i]) {
			tokens.tokens = append(tokens.tokens, IdentToken{s[i]})
		} else if s[i] == ";" {
			tokens.tokens = append(tokens.tokens, SemicolonToken{})
		}
	}
	tokens.tokens = append(tokens.tokens, EOFToken{})
	return &tokens
}

func isIdent(s string) bool {
	re := regexp.MustCompile("[a-z]")
	return re.Match([]byte(s))
}

func isComparisonOperatorSymbol(s string) bool {
	return s[0] == '=' || s[0] == '>' || s[0] == '<'
}

func mapComparisonOperator(s string) Token {
	switch s {
	case "==":
		return EqualToken{}
	case "<":
		return LessToken{}
	case ">":
		return GreaterToken{}
	case "<=":
		return LessThanEqualToken{}
	case ">=":
		return GreaterThanEqualToken{}
	}

	os.Exit(1)
	return nil
}

func lastIndex(s string, pattern string) int {
	re := regexp.MustCompile(pattern)
	match := re.FindString(s)
	if match == "" {
		return 0
	}
	return len(match) - 1
}
