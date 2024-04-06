package main

import (
	"fmt"
	"os"
	"strings"
)

func translateArithmeticOperator(nodeType NodeType) string {
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

func translateComparisonOperator(nodeType NodeType) string {
	switch nodeType {
	case EQ:
		return "cset x0, eq\n"
	case NE:
		return "cset x0, ne\n"
	case GT:
		return "cset x0, gt\n"
	case GE:
		return "cset x0, ge\n"
	case LT:
		return "cset x0, lt\n"
	case LE:
		return "cset x0, le\n"
	default:
		os.Exit(1)
		return ""
	}
}

func GenerateAssembly(result *strings.Builder, node *Node) {
	if node.NodeType == NUM {
		fmt.Fprintf(result, "mov x0, %v\n", node.Number) // 直接できる気もする？
		fmt.Fprintln(result, "str x0, [sp, #-16]!")
	} else if isArithmeticOperator(node) {
		// top of 2 values are target of the operation
		GenerateAssembly(result, node.Left)
		GenerateAssembly(result, node.Right)
		// pop both values
		fmt.Fprintln(result, "ldr x1, [sp], #16")
		fmt.Fprintln(result, "ldr x0, [sp], #16")
		// operation
		fmt.Fprint(result, translateArithmeticOperator(node.NodeType))
		fmt.Fprintln(result, "str x0, [sp, #-16]!")
	} else if isComparisonOperator(node) {
		// top of 2 values are target of the operation
		GenerateAssembly(result, node.Left)
		GenerateAssembly(result, node.Right)
		// pop both values
		fmt.Fprintln(result, "ldr x1, [sp], #16")
		fmt.Fprintln(result, "ldr x0, [sp], #16")
		fmt.Fprintln(result, "cmp x0, x1")
		fmt.Fprintln(result, "mov x0, 0")
		fmt.Fprint(result, translateComparisonOperator(node.NodeType))
		fmt.Fprintln(result, "str x0, [sp, #-16]!")
	}
}

func isArithmeticOperator(n *Node) bool {
	t := n.NodeType
	return t == ADD || t == SUB || t == DIV || t == MUL
}

func isComparisonOperator(n *Node) bool {
	t := n.NodeType
	return t == EQ || t == NE || t == GT || t == GE || t == LT || t == LE

}
