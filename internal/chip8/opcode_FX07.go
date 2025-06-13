package chip8

// OpcodeFX07
// LD VX, DT
// Read16 the delay timer register value into VX.
type OpcodeFX07 struct {
}

func (i *OpcodeFX07) Support(opcode uint16) bool {
	return (opcode & 0xF0FF) == 0xF007
}

func (i *OpcodeFX07) Do(opcode uint16, r *Runtime) {
	vx := opcodeX(opcode)

	r.v[vx] = r.dt

	r.pc.Next()
}
