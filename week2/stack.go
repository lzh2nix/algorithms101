package week2

// Item for generic type
type Item struct {
	Value interface{}
}
type Stack interface {
	// push to stack
	Push(item *Item)
	// pop from stack and remove for top
	Pop() *Item
	// empty check
	IsEmpty() bool
	// just return the top of stack
	Peek() *Item
}

// array based stack
type ArrayStack struct {
	elements []*Item
}

// create a new array based stack
// [a,b,c,d,e]
func NewArrayStack() Stack {
	return &ArrayStack{[]*Item{}}
}

// implement Stack's push func
func (s *ArrayStack) Push(item *Item) {
	s.elements = append(s.elements, item)
}

// implement Stack's pop func
func (s *ArrayStack) Pop() *Item {
	if s.IsEmpty() {
		return nil
	}
	t := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return t
}

// implement Stack's push func
func (s *ArrayStack) Peek() *Item {
	if s.IsEmpty() {
		return nil
	}
	return s.elements[len(s.elements)-1]
}

// check empty stack
func (s *ArrayStack) IsEmpty() bool {
	return len(s.elements) == 0

}

type node struct {
	value *Item
	next *node
}
// LinkedList based stack
type LinkedListStack struct {
	*node
}

// create a new linkedlist based stack
//  a--->b--->c--->d--->e
func NewLinkedListStack() Stack {
	return &LinkedListStack{&node{next:nil}}
}

// implement Stack's push func
func (s *LinkedListStack) Push(item *Item) {
	old := s.next
	newNode := &node{item, old}
	s.next = newNode
}

// implement Stack's pop func
func (s *LinkedListStack) Pop() *Item {
	if s.IsEmpty() {
		return nil
	}
	next := s.next
	s.next = next.next
	next.next = nil
	return next.value
}

// implement Stack's push func
func (s *LinkedListStack) Peek() *Item {
	if s.IsEmpty() {
		return nil
	}
	return s.next.value
}


// check empty stack
func (s *LinkedListStack) IsEmpty() bool {
	return s.next == nil

}
