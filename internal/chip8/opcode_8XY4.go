package chip8

// Opcode8XY4
// ADD VX, VY
// Add the values of VX and VY and store the result in VX. Put the carry bit in VF (if there is overflow, set VF to 1, otherwise 0).
type Opcode8XY4 struct {
}

func (i *Opcode8XY4) Support(opcode uint16) bool {
	return (opcode & 0xF00F) == 0x8004
}

func (i *Opcode8XY4) Do(opcode uint16, r *Runtime) {
	vx, vy := opcodeX(opcode), opcodeY(opcode)

	vf := uint8(0)

	if uint16(r.v[vx])+uint16(r.v[vy]) > 0xFF {
		vf = 1
	}

	r.v[vx] = r.v[vx] + r.v[vy]
	r.v[0xF] = vf

	r.pc.Next()
}
