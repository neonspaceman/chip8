package chip8

// Opcode5XY0
// SE VX, VY
// Skip the next instruction if VX == VY. Compare the value of register VX with the value of VY and if they are equal, increment PC by two
type Opcode5XY0 struct {
}

func (i *Opcode5XY0) Support(opcode uint16) bool {
	return (opcode & 0xF00F) == 0x5000
}

func (i *Opcode5XY0) Do(opcode uint16, r *Runtime) {
	vx, vy := opcodeX(opcode), opcodeY(opcode)

	if r.v[vx] == r.v[vy] {
		r.pc.Next()
	}

	r.pc.Next()
}
