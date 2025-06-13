package chip8

// OpcodeANNN
// LD I, NNN
// Write16 the value of I to the address NNN.
type OpcodeANNN struct {
}

func (i *OpcodeANNN) Support(opcode uint16) bool {
	return (opcode & 0xF000) == 0xA000
}

func (i *OpcodeANNN) Do(opcode uint16, r *Runtime) {
	nnn := opcodeNNN(opcode)

	r.i = nnn

	r.pc.Next()
}
