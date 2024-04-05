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
	EQ
	NE
	GT
	GE
	LT
	LE
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

func add(t *Tokens) *Node {
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

func relational(t *Tokens) *Node {
	node := add(t)
	for {
		if t.ConsumeGreaterToken() {
			node = &Node{NodeType: GT, Left: node, Right: add(t)}
		} else if t.ConsumeGreaterThanEqualToken() {
			node = &Node{NodeType: GE, Left: node, Right: add(t)}
		} else if t.ConsumeLessToken() {
			node = &Node{NodeType: LT, Left: node, Right: add(t)}
		} else if t.ConsumeLessThanToken() {
			node = &Node{NodeType: LE, Left: node, Right: add(t)}
		} else {
			return node
		}
	}
}

func equality(t *Tokens) *Node {
	node := relational(t)
	for {
		if t.ConsumeEqualToken() {
			node = &Node{NodeType: EQ, Left: node, Right: relational(t)}
		} else if t.ConsumeNotEqualToken() {
			node = &Node{NodeType: NE, Left: node, Right: relational(t)}
		} else {
			return node
		}
	}
}

func expr(t *Tokens) *Node {
	return equality(t)
}

func Parse(tokens *Tokens) *Node {
	return expr(tokens)
}
