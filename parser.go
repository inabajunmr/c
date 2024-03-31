package main

type Node struct {
	NodeType NodeType
	Left     *Node
	Right    *Node
	Number   int
}

type NodeType int

const (
	ADD NodeType = iota
	SUB
	MUL
	DIV
	NUM
)

func mul(t *Tokens) *Node {
	node := primary(t)
	for {
		if t.ConsumeMultiplicationToken() {
			node = &Node{NodeType: MUL, Left: node, Right: primary(t)}
		} else if t.ConsumeDivisionToken() {
			node = &Node{NodeType: DIV, Left: node, Right: primary(t)}
		} else {
			return node
		}
	}
}

func primary(t *Tokens) *Node {
	if t.ConsumeLParenthesisToken() {
		e := expr(t)
		t.ConsumeRParenthesisTokenMust()
		return e
	}
	return &Node{
		NodeType: NUM,
		Number:   t.ConsumeNumberMust().Value,
	}
}

func expr(t *Tokens) *Node {
	node := mul(t)
	for {
		if t.ConsumePlusToken() {
			node = &Node{NodeType: ADD, Left: node, Right: mul(t)}
		} else if t.ConsumeMinusToken() {
			node = &Node{NodeType: SUB, Left: node, Right: mul(t)}
		} else {
			return node
		}
	}
}

func Parse(tokens *Tokens) *Node {
	return expr(tokens)
}
