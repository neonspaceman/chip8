package chip8

// Opcode3XNN
// SE VX, NN
// Skip the next instruction if VX == NN. Compare the value of register VX with NN and if they are equal, increment PC by two
type Opcode3XNN struct {
}

func (i *Opcode3XNN) Support(opcode uint16) bool {
	return (opcode & 0xF000) == 0x3000
}

func (i *Opcode3XNN) Do(opcode uint16, r *Runtime) {
	vx, nn := opcodeX(opcode), opcodeNN(opcode)

	if r.v[vx] == nn {
		r.pc.Next()
	}

	r.pc.Next()
}
