package stack_test

import (
	"testing"

	"github.com/lzh2nix/algorithms101/stack"
	"github.com/stretchr/testify/assert"
)

func testStack(s stack.Stack, t *testing.T, tc string) {
	s.Push(&stack.Item{1})
	s.Push(&stack.Item{2})
	s.Push(&stack.Item{3})
	s.Push(&stack.Item{4})
	s.Push(&stack.Item{5})
	assert.Equal(t, 5, s.Pop().Value, "test pop for "+tc)
	assert.Equal(t, 4, s.Pop().Value, "test pop for "+tc)
	assert.Equal(t, 3, s.Pop().Value, "test pop for "+tc)
	assert.Equal(t, 2, s.Pop().Value, "test pop for "+tc)
	assert.Equal(t, 1, s.Peek().Value, "test peek for "+tc)
	assert.False(t, s.IsEmpty(), "test for is empty"+tc)
	assert.Equal(t, 1, s.Pop().Value, "test pop for "+tc)
	assert.True(t, s.IsEmpty(), "test for is empty"+tc)
	assert.Nil(t, s.Pop(), "test for pop nil"+tc)
	assert.Nil(t, s.Peek(), "test for peek nil"+tc)

}
func TestArrayStack(t *testing.T) {
	s := stack.NewArrayStack()
	testStack(s, t, "arrayStack")
}
func TestLinkedListStack(t *testing.T) {
	s := stack.NewLinkedListStack()
	testStack(s, t, "linkedListStack")
}
