package chip8

// OpcodeFX0A
// LD VX, K
// Wait for a key press, and then store the value of the key to VX.
type OpcodeFX0A struct {
}

func (i *OpcodeFX0A) Support(opcode uint16) bool {
	return (opcode & 0xF0FF) == 0xF00A
}

func (i *OpcodeFX0A) Do(opcode uint16, r *Runtime) {
	panic("LD VX, K NOT IMPLEMENTED")
}
