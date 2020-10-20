package week2

import (
	"strconv"
)

const (
	// Number 0,1,2,4,5,100
	Number = iota
	// +, -, *, /
	Operator
	// (
	OpenBracket
	// }
	CloseBracket
	// end of string
	EOF
)

// Evaluate infix expression(E.W.Dijkstra)
func Evaluate(input string) float64 {
	numers := NewArrayStack()
	ops := NewArrayStack()
	exp := &expressionParse{input, 0}
	for {
		k, v := exp.Next()
		if k == EOF {
			break
		}
		if k == Number {
			numers.Push(&Item{v})
		} else if k == Operator {
			ops.Push(&Item{v})
		} else if k == CloseBracket {
			v1 := numers.Pop()
			v2 := numers.Pop()
			op := ops.Pop()
			res := float64(0)
			if op.Value.(string) == "+" {
				res = v2.Value.(float64) + v1.Value.(float64)
			}
			if op.Value.(string) == "-" {
				res = v2.Value.(float64) - v1.Value.(float64)
			}
			if op.Value.(string) == "*" {
				res = v2.Value.(float64) * v1.Value.(float64)
			}

			if op.Value.(string) == "/" {
				res = v2.Value.(float64) / v1.Value.(float64)
			}
			numers.Push(&Item{res})
		}
	}
	return numers.Pop().Value.(float64)
}

type expressionParse struct {
	input string
	idx   int
}

func (e *expressionParse) Next() (int, interface{}) {
	if e.idx >= len(e.input) {
		return EOF, nil
	}
	if e.input[e.idx] == '(' {
		e.idx++
		return OpenBracket, "("
	}
	if e.input[e.idx] == ')' {
		e.idx++
		return CloseBracket, ")"
	}
	if e.input[e.idx] == '+' || e.input[e.idx] == '-' || e.input[e.idx] == '*' || e.input[e.idx] == '/' {
		e.idx++
		return Operator, OperatorByteToString(e.input[e.idx-1])
	}
	start := e.idx
	for ; e.idx < len(e.input) && ((e.input[e.idx] >= '0' && e.input[e.idx] <= '9') || e.input[e.idx] == '.'); e.idx++ {

	}
	f, _ := strconv.ParseFloat(e.input[start:e.idx], 64)
	return Number, f
}

func OperatorByteToString(b byte) string {
	if b == '+' {
		return "+"
	}
	if b == '-' {
		return "-"
	}
	if b == '*' {
		return "*"
	}

	return "/"
}
