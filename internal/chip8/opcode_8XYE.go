package chip8

// Opcode8XYE
// SHL VX {, VY}
// Shift left, or multiply VX by two. Store the most significant bit of VX in VF, and then multiply VX and store its value in VX
type Opcode8XYE struct {
}

func (i *Opcode8XYE) Support(opcode uint16) bool {
	return (opcode & 0xF00F) == 0x800E
}

func (i *Opcode8XYE) Do(opcode uint16, r *Runtime) {
	vx, vy := opcodeX(opcode), opcodeY(opcode)

	vf := (r.v[vy] & 0b1000_0000) >> 7

	r.v[vx] = r.v[vy] << 1
	r.v[0xF] = vf

	r.pc.Next()
}
