package week2

// double end queue(can sadd/remove from front && end)
//Algorithms 4th 1.3.33
type Deque struct {
	elements []*Item
}

// create a new dqueue
func NewDeque() *Deque {
	return &Deque{[]*Item{}}
}

// push from left end
func (s *Deque) PushLeft(item *Item) {
	s.elements = append([]*Item{item}, s.elements...)
}

// push from right end
func (s *Deque) PushRight(item *Item) {
	s.elements = append(s.elements, item)
}

// pop from left end
func (s *Deque) PopLeft() *Item {
	if s.IsEmpty() {
		return nil
	}
	t := s.elements[0]
	s.elements = s.elements[1:]
	return t
}

// pop from right end
func (s *Deque) PopRight() *Item {
	if s.IsEmpty() {
		return nil
	}
	t := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return t
}

// check empty queue
func (s *Deque) IsEmpty() bool {
	return len(s.elements) == 0

}
