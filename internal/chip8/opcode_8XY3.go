package chip8

// Opcode8XY3
// XOR VX, VY
// Perform a bitwise XOR between the values of VX and VY and store the result in VX.
type Opcode8XY3 struct {
}

func (i *Opcode8XY3) Support(opcode uint16) bool {
	return (opcode & 0xF00F) == 0x8003
}

func (i *Opcode8XY3) Do(opcode uint16, r *Runtime) {
	vx, vy := opcodeX(opcode), opcodeY(opcode)

	r.v[vx] = r.v[vx] ^ r.v[vy]
	r.v[0xF] = 0

	r.pc.Next()
}
