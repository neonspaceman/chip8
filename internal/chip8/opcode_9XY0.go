package chip8

// Opcode9XY0
// SNE VX, VY
// Skip the next instruction if the values of VX and VY are not equal.
type Opcode9XY0 struct {
}

func (i *Opcode9XY0) Support(opcode uint16) bool {
	return (opcode & 0xF00F) == 0x9000
}

func (i *Opcode9XY0) Do(opcode uint16, r *Runtime) {
	vx, vy := opcodeX(opcode), opcodeY(opcode)

	if r.v[vx] != r.v[vy] {
		r.pc.Next()
	}

	r.pc.Next()
}
