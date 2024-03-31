package main

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		expected []Token
	}{
		{"only numbers", "12", []Token{NumberToken{12}, EOFToken{}}},
		{"only operators", "+-+*/", []Token{PlusToken{}, MinusToken{}, PlusToken{}, MultiplicationToken{}, DivisionToken{}, EOFToken{}}},
		{"numbers and operators", "12345+234", []Token{NumberToken{12345}, PlusToken{}, NumberToken{234}, EOFToken{}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := Tokenize(tc.source)

			if !reflect.DeepEqual(actual.tokens, tc.expected) {
				t.Errorf("got %d, want %d", actual, tc.expected)
			}
		})
	}
}

func TestLastIndex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		pattern  string
		expected int
	}{
		{"alphabet then numbers", "abc1234abc", "[a-z]+", 2},
		{"only numbers", "12345", "[a-z]+", 0},
		{"numbers at end", "12345", "[0-9]+", 4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := lastIndex(tc.input, tc.pattern)
			if actual != tc.expected {
				t.Errorf("got %d, want %d", actual, tc.expected)
			}
		})
	}
}
