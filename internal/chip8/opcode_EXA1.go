package chip8

// OpcodeEXA1
// SKNP VX
// Skip the next instruction if the key with the value of VX is currently not pressed. Basically, increase PC by two if the key corresponding to the value in VX is not pressed.
type OpcodeEXA1 struct {
}

func (i *OpcodeEXA1) Support(opcode uint16) bool {
	return (opcode & 0xF0FF) == 0xE0A1
}

func (i *OpcodeEXA1) Do(opcode uint16, r *Runtime) {
	panic("SKPN NOT IMPLEMENTED")
}
