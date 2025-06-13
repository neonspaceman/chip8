package chip8

// Opcode6XNN
// LD VX, NN
// Load the value NN into the register VX.
type Opcode6XNN struct {
}

func (i *Opcode6XNN) Support(opcode uint16) bool {
	return (opcode & 0xF000) == 0x6000
}

func (i *Opcode6XNN) Do(opcode uint16, r *Runtime) {
	vx, nn := opcodeX(opcode), opcodeNN(opcode)

	r.v[vx] = nn

	r.pc.Next()
}
