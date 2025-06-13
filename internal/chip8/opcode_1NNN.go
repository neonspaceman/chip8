package chip8

// Opcode1NNN
// JMP
// Jump to the address in NNN. Sets the PC to NNN.
type Opcode1NNN struct {
}

func (i *Opcode1NNN) Support(opcode uint16) bool {
	return (opcode & 0xF000) == 0x1000
}

func (i *Opcode1NNN) Do(opcode uint16, r *Runtime) {
	nnn := opcodeNNN(opcode)

	r.pc.Set(nnn)
}
