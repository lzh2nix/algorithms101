package week2_test

import (
	"testing"

	"github.com/lzh2nix/algorithms101/week2"
	"github.com/stretchr/testify/assert"
)

func TestRingBuffer(t *testing.T) {
	rb := week2.NewRingBuffer(5)
	var err error
	err = rb.Push(&week2.Item{1})
	assert.Nil(t, err)
	err = rb.Push(&week2.Item{2})
	assert.Nil(t, err)
	err = rb.Push(&week2.Item{3})
	assert.Nil(t, err)
	err = rb.Push(&week2.Item{4})
	assert.Nil(t, err)
	err = rb.Push(&week2.Item{5})
	assert.Nil(t, err)
	err = rb.Push(&week2.Item{6})
	assert.NotNil(t, err)

	var v *week2.Item
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
}
