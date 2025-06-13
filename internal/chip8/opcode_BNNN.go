package chip8

// OpcodeBNNN
// JMP V0, NNN
// Jump to the location NNN + V0.
type OpcodeBNNN struct {
}

func (i *OpcodeBNNN) Support(opcode uint16) bool {
	return (opcode & 0xF000) == 0xB000
}

func (i *OpcodeBNNN) Do(opcode uint16, r *Runtime) {
	nnn := opcodeNNN(opcode)

	r.pc.Set(uint16(r.v[0]) + nnn)
}
