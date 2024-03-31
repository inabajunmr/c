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

func translateOperator(nodeType NodeType) string {
	switch nodeType {
	case ADD:
		return "add x0, x0, x1\n"
	case SUB:
		return "sub x0, x0, x1\n"
	case MUL:
		return "mul x0, x0, x1\n"
	case DIV:
		return "udiv x0, x0, x1\n"
	default:
		os.Exit(1)
		return ""
	}
}

func GenerateAssembly(result *strings.Builder, node *Node) {
	if node.NodeType == NUM {
		fmt.Fprintf(result, "mov x0, %v\n", node.Number) // 直接できる気もする？
		fmt.Fprintln(result, "str x0, [sp, #-16]!")
	} else {
		// top of 2 values are target of the operation
		GenerateAssembly(result, node.Left)
		GenerateAssembly(result, node.Right)
		// pop both values
		fmt.Fprintln(result, "ldr x1, [sp], #16")
		fmt.Fprintln(result, "ldr x0, [sp], #16")
		// operation
		fmt.Fprint(result, translateOperator(node.NodeType))
		fmt.Fprintln(result, "str x0, [sp, #-16]!")
	}
}

func compile(source string) string {
	node := Parse(Tokenize(source))

	var result strings.Builder

	fmt.Fprintln(&result, ".global main")
	fmt.Fprintln(&result, "main:")
	GenerateAssembly(&result, node)
	fmt.Fprintln(&result, "ldr x0, [sp], #16")
	fmt.Fprintln(&result, "mov x8, 93")
	fmt.Fprintln(&result, "svc 0")

	return result.String()
}
