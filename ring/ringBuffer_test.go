package ring_test

import (
	"testing"

	"github.com/lzh2nix/algorithms101/ring"
	"github.com/stretchr/testify/assert"
)

func TestRingBuffer(t *testing.T) {
	rb := ring.NewRingBuffer(5)
	var err error
	err = rb.Push(&ring.Item{1})
	assert.Nil(t, err)
	err = rb.Push(&ring.Item{2})
	assert.Nil(t, err)
	err = rb.Push(&ring.Item{3})
	assert.Nil(t, err)
	err = rb.Push(&ring.Item{4})
	assert.Nil(t, err)
	err = rb.Push(&ring.Item{5})
	assert.Nil(t, err)
	err = rb.Push(&ring.Item{6})
	assert.NotNil(t, err)

	var v *ring.Item
	v, err = rb.Poll()
	assert.Nil(t, err)
	assert.Equal(t, 1, v.Value)
	v, err = rb.Poll()
	assert.Nil(t, err)
	assert.Equal(t, 2, v.Value)
	v, err = rb.Poll()
	assert.Nil(t, err)
	assert.Equal(t, 3, v.Value)
	v, err = rb.Poll()
	assert.Nil(t, err)
	assert.Equal(t, 4, v.Value)
	v, err = rb.Poll()
	assert.Nil(t, err)
	assert.Equal(t, 5, v.Value)
	v, err = rb.Poll()
	assert.NotNil(t, err)
	err = rb.Push(&ring.Item{6})
	assert.Nil(t, err)
}
