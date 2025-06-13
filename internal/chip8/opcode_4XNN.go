package chip8

// Opcode4XNN
// SNE VX, NN
// Skip the next instruction if VX != NN. Compare the value of register VX with NN and if they are not equal, increment PC by two.
type Opcode4XNN struct {
}

func (i *Opcode4XNN) Support(opcode uint16) bool {
	return (opcode & 0xF000) == 0x4000
}

func (i *Opcode4XNN) Do(opcode uint16, r *Runtime) {
	vx, nn := opcodeX(opcode), opcodeNN(opcode)

	if r.v[vx] != nn {
		r.pc.Next()
	}

	r.pc.Next()
}
