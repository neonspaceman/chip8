package chip8

// Opcode8XY0
// LD VX, VY
// Put the value of register VY into VX.
type Opcode8XY0 struct {
}

func (i *Opcode8XY0) Support(opcode uint16) bool {
	return (opcode & 0xF00F) == 0x8000
}

func (i *Opcode8XY0) Do(opcode uint16, r *Runtime) {
	vx, vy := opcodeX(opcode), opcodeY(opcode)

	r.v[vx] = r.v[vy]

	r.pc.Next()
}
