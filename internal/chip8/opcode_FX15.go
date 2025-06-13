package chip8

// OpcodeFX15
// LD DT, VX
// Load the value of VX into the delay timer DT.
type OpcodeFX15 struct {
}

func (i *OpcodeFX15) Support(opcode uint16) bool {
	return (opcode & 0xF0FF) == 0xF015
}

func (i *OpcodeFX15) Do(opcode uint16, r *Runtime) {
	vx := opcodeX(opcode)

	r.dt = r.v[vx]

	r.pc.Next()
}
