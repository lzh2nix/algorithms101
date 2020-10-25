package stack_test

import (
	"testing"

	"github.com/lzh2nix/algorithms101/stack"
	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
	tcs := []struct {
		exp    string
		expect float64
	}{
		{"(1+2)", 3.0},
		{"((1-2)*3)", -3.0},
		{"((1+2)*3)", 9.0},
		{"(1+((2+3)*(4*5)))", 101.0},
		{"((100+1)*2)", 202.0},
		{"((100.1+1)*2)", 202.2},
		{"((100.1+1)/2)", 50.55},
	}

	for _, tc := range tcs {
		res := stack.Evaluate(tc.exp)
		assert.Equal(t, tc.expect, res, tc.exp)
	}
}
