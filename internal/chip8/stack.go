package chip8

type Stack struct {
	current int
	size    int
	stack   []uint16
}

func NewStack(size int) Stack {
	return Stack{
		current: -1,
		size:    size,
		stack:   make([]uint16, size),
	}
}

func (s *Stack) Push(value uint16) {
	if s.current+1 >= s.size {
		panic("Stack overflow")
	}

	s.current++

	s.stack[s.current] = value
}

func (s *Stack) Pop() uint16 {
	if s.current == -1 {
		panic("Stack out of range")
	}

	value := s.stack[s.current]

	s.current--

	return value
}
