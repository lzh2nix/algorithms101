package week2_test

import (
	"testing"

	"github.com/lzh2nix/algorithms101/week2"
	"github.com/stretchr/testify/assert"
)

func testStack(stack week2.Stack, t *testing.T, tc string) {
	stack.Push(&week2.Item{1})
	stack.Push(&week2.Item{2})
	stack.Push(&week2.Item{3})
	stack.Push(&week2.Item{4})
	stack.Push(&week2.Item{5})
	assert.Equal(t, 5, stack.Pop().Value, "test pop for "+tc)
	assert.Equal(t, 4, stack.Pop().Value, "test pop for "+tc)
	assert.Equal(t, 3, stack.Pop().Value, "test pop for "+tc)
	assert.Equal(t, 2, stack.Pop().Value, "test pop for "+tc)
	assert.Equal(t, 1, stack.Peek().Value, "test peek for "+tc)
	assert.False(t, stack.IsEmpty(), "test for is empty"+tc)
	assert.Equal(t, 1, stack.Pop().Value, "test pop for "+tc)
	assert.True(t, stack.IsEmpty(), "test for is empty"+tc)
	assert.Nil(t, stack.Pop(), "test for pop nil"+tc)
	assert.Nil(t, stack.Peek(), "test for peek nil"+tc)

}
func TestArrayStack(t *testing.T) {
	s := week2.NewArrayStack()
	testStack(s, t, "arrayStack")
}
func TestLinkedListStack(t *testing.T) {
	s := week2.NewLinkedListStack()
	testStack(s, t, "linkedListStack")
}
