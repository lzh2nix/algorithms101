package week2

// Item for generic type
type Queue interface {
	// enqueue to tail
	Enqueue(item *Item)
	// dequeue from head
	Dequeue() *Item
	// empty check
	IsEmpty() bool
}

// array based queue
type ArrayQueue struct {
	elements []*Item
}

// create a new array based queue
// [a,b,c,d,e]
func NewArrayQueue() Queue {
	return &ArrayQueue{[]*Item{}}
}

// implement Queue's  Enqueue func
func (s *ArrayQueue) Enqueue(item *Item) {
	s.elements = append(s.elements, item)
}

// implement Queue's Dequeue func
func (s *ArrayQueue) Dequeue() *Item {
	if s.IsEmpty() {
		return nil
	}
	t := s.elements[0]
	s.elements = s.elements[1:]
	return t
}

// check empty queue
func (s *ArrayQueue) IsEmpty() bool {
	return len(s.elements) == 0

}

// LinkedList based queue
type LinkedListQueue struct {
	*node
	tail *node
}

// create a new linkedlist based queue
//  a--->b--->c--->d--->e
func NewLinkedListQueue() Queue {
	return &LinkedListQueue{&node{next: nil}, nil}
}

// implement Queue's enqueue func
func (s *LinkedListQueue) Enqueue(item *Item) {

	oldTail := s.tail
	newNode := &node{item, nil}
	if s.IsEmpty() {
		s.next = newNode
	} else {
		oldTail.next = newNode
	}
	s.tail = newNode
}

// implement Queue's dequeue func
func (s *LinkedListQueue) Dequeue() *Item {
	if s.IsEmpty() {
		return nil
	}
	next := s.next
	s.next = next.next
	next.next = nil
	return next.value
}

// check empty queue
func (s *LinkedListQueue) IsEmpty() bool {
	return s.next == nil

}
