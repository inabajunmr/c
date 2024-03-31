package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("invalid args %v", os.Args)
		os.Exit(1)
	}

	fmt.Println(compile(os.Args[1]))
}

func compile(source string) string {
	tokens := Tokenize(source)

	var result strings.Builder

	fmt.Fprintln(&result, ".global main")
	fmt.Fprintln(&result, "main:")

	if number, ok := tokens[0].(NumberToken); ok {
		fmt.Fprintf(&result, "mov x0, %d\n", number.Value)
	} else {
		os.Exit(1)
	}

	index := 1
	for len(tokens) > index {
		t := tokens[index]
		if operator, ok := t.(OperatorToken); ok {
			index++
			if number, ok := tokens[index].(NumberToken); ok {
				if operator.Value == "+" {
					fmt.Fprintf(&result, "add x0, x0, %d\n", number.Value)
				} else if operator.Value == "-" {
					fmt.Fprintf(&result, "sub x0, x0, %d\n", number.Value)
				} else {
					os.Exit(1)
				}
			} else {
				os.Exit(1)
			}
		}
		index++
	}

	fmt.Fprintln(&result, "mov x8, 93")
	fmt.Fprintln(&result, "svc 0")

	return result.String()
}
