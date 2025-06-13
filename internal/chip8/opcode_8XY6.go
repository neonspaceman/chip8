package chip8

// Opcode8XY6
// SHR VX {, VY}
// Shift right, or divide VX by two. Store the least significant bit of VX in VF, and then divide VX and store its value in VX
type Opcode8XY6 struct {
}

func (i *Opcode8XY6) Support(opcode uint16) bool {
	return (opcode & 0xF00F) == 0x8006
}

func (i *Opcode8XY6) Do(opcode uint16, r *Runtime) {
	vx, vy := opcodeX(opcode), opcodeY(opcode)

	vf := r.v[vy] & 0b1

	r.v[vx] = r.v[vy] >> 1
	r.v[0xF] = vf

	r.pc.Next()
}
