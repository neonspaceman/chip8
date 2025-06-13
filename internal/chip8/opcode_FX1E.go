package chip8

// OpcodeFX1E
// ADD I, VX
// Add the values of I and VX, and store the result in I.
type OpcodeFX1E struct {
}

func (i *OpcodeFX1E) Support(opcode uint16) bool {
	return (opcode & 0xF0FF) == 0xF01E
}

func (i *OpcodeFX1E) Do(opcode uint16, r *Runtime) {
	vx := opcodeX(opcode)

	r.i += uint16(r.v[vx])

	r.pc.Next()
}
