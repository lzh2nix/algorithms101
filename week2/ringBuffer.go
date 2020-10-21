package week2

import "errors"

// array base ringBuffer
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

var errFULL = errors.New("ringBuffer is full")
var errEmpty = errors.New("ringBuffer is empty")

// create a new ringBuffer
func NewRingBuffer(cap int) *RingBuffer {
	data := make([]*Item, cap)
	return &RingBuffer{cap, data, 0, 0, 0}
}

func (rb *RingBuffer) Push(item *Item) error {

	if rb.len == rb.capacity {
		return errFULL
	}
	rb.data[rb.write] = item
	rb.write = (rb.write + 1) % rb.capacity
	rb.len += 1
	return nil
}

func (rb *RingBuffer) Poll() (*Item, error) {
	if rb.len == 0 {
		return nil, errEmpty
	}
	item := rb.data[rb.read]
	rb.read = (rb.read + 1) % rb.capacity
	rb.len -= 1
	return item, nil
}
