package chip8

// Opcode00EE
// RET
// Return from a subroutine. Pops the value at the top of the stack (indicated by the stack pointer SP) and puts it in PC.
type Opcode00EE struct {
}

func (i *Opcode00EE) Support(opcode uint16) bool {
	return opcode == 0x00EE
}

func (i *Opcode00EE) Do(opcode uint16, r *Runtime) {
	r.pc.Set(r.stack.Pop())
	r.pc.Next()
}
