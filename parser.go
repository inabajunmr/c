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
	node := unary(t)
	for {
		if t.ConsumeMultiplicationToken() {
			node = &Node{NodeType: MUL, Left: node, Right: unary(t)}
		} else if t.ConsumeDivisionToken() {
			node = &Node{NodeType: DIV, Left: node, Right: unary(t)}
		} else {
			return node
		}
	}
}

func unary(t *Tokens) *Node {
	if t.ConsumePlusToken() {
		return primary(t)
	} else if t.ConsumeMinusToken() {
		// 0 - value = -value
		return &Node{NodeType: SUB, Left: &Node{
			NodeType: NUM,
			Number:   0,
		}, Right: primary(t)}
	}
	return primary(t)
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
