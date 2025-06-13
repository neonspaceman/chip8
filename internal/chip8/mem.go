package chip8

type Mem struct {
	size uint16
	heap []uint8
}

func NewMem(size uint16) Mem {
	return Mem{
		size: size,
		heap: make([]uint8, size),
	}
}

func (m *Mem) Read(address uint16) uint8 {
	m.assertCorrectAddress(address)
	return m.heap[address]
}

func (m *Mem) Write(address uint16, data uint8) {
	m.assertCorrectAddress(address)
	m.heap[address] = data
}

func (m *Mem) assertCorrectAddress(address uint16) {
	if address >= m.size {
		panic("Address out of memory")
	}
}
