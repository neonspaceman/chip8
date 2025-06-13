package chip8

// Opcode8XY1
// OR VX, VY
// Perform a bitwise OR between the values of VX and VY and store the result in VX.
type Opcode8XY1 struct {
}

func (i *Opcode8XY1) Support(opcode uint16) bool {
	return (opcode & 0xF00F) == 0x8001
}

func (i *Opcode8XY1) Do(opcode uint16, r *Runtime) {
	vx, vy := opcodeX(opcode), opcodeY(opcode)

	r.v[vx] = r.v[vx] | r.v[vy]
	r.v[0xF] = 0

	r.pc.Next()
}
