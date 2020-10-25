package deque_test

import (
	"testing"

	"github.com/lzh2nix/algorithms101/deque"
	"github.com/stretchr/testify/assert"
)

func TestDequeue(t *testing.T) {
	s := deque.NewDeque()
	s.PushLeft(&deque.Item{1})
	// 1
	s.PushLeft(&deque.Item{2})
	// 2 1
	s.PushRight(&deque.Item{3})
	// 2 1 3
	s.PushRight(&deque.Item{4})
	// 2 1 3 4
	s.PushRight(&deque.Item{5})
	// 2 1 4 5
	assert.Equal(t, 2, s.PopLeft().Value, "pop left")
	assert.Equal(t, 1, s.PopLeft().Value, "pop left")
	assert.Equal(t, 3, s.PopLeft().Value, "pop left")
	assert.False(t, s.IsEmpty(), "test is empty")
	assert.Equal(t, 5, s.PopRight().Value, "pop right")
	assert.Equal(t, 4, s.PopLeft().Value, "pop left")
	assert.Nil(t, s.PopLeft(), "pop left")
	assert.Nil(t, s.PopRight(), "pop left")

	assert.True(t, s.IsEmpty(), "test is empty")

}
