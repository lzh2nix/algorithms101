package week2_test

import (
	"testing"

	"github.com/lzh2nix/algorithms101/week2"
	"github.com/stretchr/testify/assert"
)

func testQueue(stack week2.Queue, t *testing.T, tc string) {
	stack.Enqueue(&week2.Item{1})
	stack.Enqueue(&week2.Item{2})
	stack.Enqueue(&week2.Item{3})
	stack.Enqueue(&week2.Item{4})
	stack.Enqueue(&week2.Item{5})
	assert.Equal(t, 1, stack.Dequeue().Value, "test dequeue for "+tc)
	assert.Equal(t, 2, stack.Dequeue().Value, "test dequeue for "+tc)
	assert.Equal(t, 3, stack.Dequeue().Value, "test dequeue for "+tc)
	assert.Equal(t, 4, stack.Dequeue().Value, "test dequeue for "+tc)
	assert.False(t, stack.IsEmpty(), "test for is empty"+tc)
	assert.Equal(t, 5, stack.Dequeue().Value, "test dequeue for "+tc)
	assert.True(t, stack.IsEmpty(), "test for is empty"+tc)
	assert.Nil(t, stack.Dequeue(), "test for dequeue nil"+tc)

}
func TestArrayQueue(t *testing.T) {
	s := week2.NewArrayQueue()
	testQueue(s, t, "arrayQueue")
}
func TestLinkedListQueue(t *testing.T) {
	s := week2.NewLinkedListQueue()
	testQueue(s, t, "linkedListQueue")
}
