package chip8

// Opcode8XY5
// SUB VX, VY
// Subtract the value of VY from VX and store the result in VX. Put the borrow in VF (if there is borrow, VX > VY, set VF to 1, otherwise 0).
type Opcode8XY5 struct {
}

func (i *Opcode8XY5) Support(opcode uint16) bool {
	return (opcode & 0xF00F) == 0x8005
}

func (i *Opcode8XY5) Do(opcode uint16, r *Runtime) {
	vx, vy := opcodeX(opcode), opcodeY(opcode)

	vf := uint8(0)

	if r.v[vx] >= r.v[vy] {
		vf = 1
	}

	r.v[vx] = r.v[vx] - r.v[vy]
	r.v[0xF] = vf

	r.pc.Next()
}
