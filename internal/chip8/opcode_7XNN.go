package chip8

// Opcode7XNN
// ADD VX, NN
// Add the value NN to the value of register VX and store the result in VX.
type Opcode7XNN struct {
}

func (i *Opcode7XNN) Support(opcode uint16) bool {
	return (opcode & 0xF000) == 0x7000
}

func (i *Opcode7XNN) Do(opcode uint16, r *Runtime) {
	vx, nn := opcodeX(opcode), opcodeNN(opcode)

	r.v[vx] = r.v[vx] + nn

	r.pc.Next()
}
