package main

import (
	"fmt"
	"os"
	"strconv"
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
	arg, err := strconv.Atoi(source)
	if err != nil {
		fmt.Println("invalid")
		os.Exit(1)
	}

	var result strings.Builder

	fmt.Fprintln(&result, ".global main")
	fmt.Fprintln(&result, "main:")
	fmt.Fprintf(&result, "mov x0, %d\n", arg)
	fmt.Fprintln(&result, "mov x8, 93")
	fmt.Fprintln(&result, "svc 0")

	return result.String()
}
