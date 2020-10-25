package queue_test

import (
	"testing"

	"github.com/lzh2nix/algorithms101/queue"
	"github.com/stretchr/testify/assert"
)

func testQueue(stack queue.Queue, t *testing.T, tc string) {
	stack.Enqueue(&queue.Item{1})
	stack.Enqueue(&queue.Item{2})
	stack.Enqueue(&queue.Item{3})
	stack.Enqueue(&queue.Item{4})
	stack.Enqueue(&queue.Item{5})
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
	s := queue.NewArrayQueue()
	testQueue(s, t, "arrayQueue")
}
func TestLinkedListQueue(t *testing.T) {
	s := queue.NewLinkedListQueue()
	testQueue(s, t, "linkedListQueue")
}
