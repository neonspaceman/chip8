package chip8

type PC struct {
	address uint16
}

func NewPC(offset uint16) PC {
	return PC{
		address: offset,
	}
}

func (pc *PC) Get() uint16 {
	return pc.address
}

func (pc *PC) Next() {
	pc.address += 2 // 1 opcode is 2 bytes
}

func (pc *PC) Set(offset uint16) {
	pc.address = offset
}
