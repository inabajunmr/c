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
	node := Parse(Tokenize(source))

	var result strings.Builder

	fmt.Fprintln(&result, ".global main")
	fmt.Fprintln(&result, "main:")
	GenerateAssembly(&result, node[0]) // TODO
	fmt.Fprintln(&result, "ldr x0, [sp], #16")
	fmt.Fprintln(&result, "mov x8, 93")
	fmt.Fprintln(&result, "svc 0")

	return result.String()
}
