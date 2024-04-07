package main

type Node struct {
	NodeType NodeType
	Left     *Node
	Right    *Node
	Number   int
	Offset   int
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
	ASSIGN
	LVAR
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

	ident := t.ConsumeIdent()
	if ident != nil {
		return &Node{
			NodeType: LVAR,
			Offset:   (int([]rune(ident.Value)[0]) - int('a') + 1) * 16,
		}
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
	return assign(t)
}

func stmt(t *Tokens) *Node {
	node := expr(t)
	t.ConsumeSemicolonTokenMust()
	return node
}

func assign(t *Tokens) *Node {
	node := equality(t)
	if t.ConsumeAssignToken() {
		node = &Node{NodeType: ASSIGN, Left: node, Right: assign(t)}
	}
	return node

}

func Parse(tokens *Tokens) []*Node {
	var result []*Node
	for len(tokens.tokens)-1 > tokens.index {
		result = append(result, stmt(tokens))
	}

	return result
}
