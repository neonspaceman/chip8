package chip8

// OpcodeFX18
// LD ST, VX
// Load the value of VX into the sound time ST.
type OpcodeFX18 struct {
}

func (i *OpcodeFX18) Support(opcode uint16) bool {
	return (opcode & 0xF0FF) == 0xF018
}

func (i *OpcodeFX18) Do(opcode uint16, r *Runtime) {
	vx := opcodeX(opcode)

	r.st = r.v[vx]

	r.pc.Next()
}
