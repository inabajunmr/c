package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		expected *Node
	}{
		{
			name:   "just division",
			source: "4/1",
			expected: &Node{
				NodeType: DIV,
				Left:     &Node{NodeType: NUM, Number: 4},
				Right:    &Node{NodeType: NUM, Number: 1},
			},
		},
		{
			name:   "complex",
			source: "(3+2)*(4/2)+(3+2*3)",
			expected: &Node{
				NodeType: ADD,
				Left: &Node{
					NodeType: MUL,
					Left: &Node{
						NodeType: ADD,
						Left:     &Node{NodeType: NUM, Number: 3},
						Right:    &Node{NodeType: NUM, Number: 2},
					},
					Right: &Node{
						NodeType: DIV,
						Left:     &Node{NodeType: NUM, Number: 4},
						Right:    &Node{NodeType: NUM, Number: 2},
					},
				},
				Right: &Node{
					NodeType: ADD,
					Left:     &Node{NodeType: NUM, Number: 3},
					Right: &Node{
						NodeType: MUL,
						Left:     &Node{NodeType: NUM, Number: 2},
						Right:    &Node{NodeType: NUM, Number: 3}},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := Parse(Tokenize(tc.source))

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("actual:%+v expected:%+v", actual, tc.expected)
			}
		})
	}
}