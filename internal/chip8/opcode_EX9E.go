package chip8

// OpcodeEX9E
// SKP VX
// Skip the next instruction if the key with the value of VX is currently pressed. Basically, increase PC by two if the key corresponding to the value in VX is pressed.
type OpcodeEX9E struct {
}

func (i *OpcodeEX9E) Support(opcode uint16) bool {
	return (opcode & 0xF0FF) == 0xE09E
}

func (i *OpcodeEX9E) Do(opcode uint16, r *Runtime) {
	panic("SKP NOT IMPLEMENTED")
}
