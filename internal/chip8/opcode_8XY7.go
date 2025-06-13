package chip8

// Opcode8XY7
// SUBN VX, VY
// Subtract the value of VY from VX and store the result in VX. Write16 VF to 1 if there is no borrow, to 0 otherwise.
type Opcode8XY7 struct {
}

func (i *Opcode8XY7) Support(opcode uint16) bool {
	return (opcode & 0xF00F) == 0x8007
}

func (i *Opcode8XY7) Do(opcode uint16, r *Runtime) {
	vx, vy := opcodeX(opcode), opcodeY(opcode)

	vf := uint8(0)

	if r.v[vy] >= r.v[vx] {
		vf = 1
	}

	r.v[vx] = r.v[vy] - r.v[vx]
	r.v[0xF] = vf

	r.pc.Next()
}
