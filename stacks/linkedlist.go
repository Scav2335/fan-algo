package stacks

type Node struct {
	val  string
	next *Node
}

type LinkedStackOfStrings struct {
	first, current *Node
}

func (s *LinkedStackOfStrings) IsEmpty() bool {
	return s.first == nil
}

func (s *LinkedStackOfStrings) Push(val string) {
	n := &Node{
		val:  val,
		next: s.first,
	}
	s.first = n
}

func (s *LinkedStackOfStrings) Pop() string {
	if s.IsEmpty() {
		panic("stack is empty")
	}

	first := s.first
	s.first = first.next
	return first.val
}

func (s *LinkedStackOfStrings) HasNext() bool {
	return s.current != nil
}

func (s *LinkedStackOfStrings) Next() string {
	if !s.HasNext() {
		panic("stack is empty")
	}

	cur := s.current
	s.current = cur.next
	return cur.val
}
