package ring

import "errors"

// Item for generic type
type Item struct {
	Value interface{}
}

// array base ringBuffer(Algorithms 4th 1.3.39)
type RingBuffer struct {
	// capacity of RingBuffer
	capacity int
	data     []*Item
	// read index
	read int
	// write index
	write int
	// len of current items
	len int
}

// ring buffer is full
var FULL = errors.New("ringBuffer is full")

// ring buffer is empty
var Empty = errors.New("ringBuffer is empty")

// create a new ringBuffer
func NewRingBuffer(cap int) *RingBuffer {
	data := make([]*Item, cap)
	return &RingBuffer{cap, data, 0, 0, 0}
}

// push to ringBuffer
func (rb *RingBuffer) Push(item *Item) error {

	if rb.len == rb.capacity {
		return FULL
	}
	rb.data[rb.write] = item
	rb.write = (rb.write + 1) % rb.capacity
	rb.len += 1
	return nil
}

// pull from ringBuffer
func (rb *RingBuffer) Poll() (*Item, error) {
	if rb.len == 0 {
		return nil, Empty
	}
	item := rb.data[rb.read]
	rb.read = (rb.read + 1) % rb.capacity
	rb.len -= 1
	return item, nil
}
