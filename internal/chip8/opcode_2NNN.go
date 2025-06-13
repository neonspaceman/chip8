package chip8

// Opcode2NNN
// CALL NNN
// Call the subroutine at address NNN. It increments SP, puts the current PC at the top of the stack and sets PC to the address NNN
type Opcode2NNN struct {
}

func (i *Opcode2NNN) Support(opcode uint16) bool {
	return (opcode & 0xF000) == 0x2000
}

func (i *Opcode2NNN) Do(opcode uint16, r *Runtime) {
	nnn := opcodeNNN(opcode)

	r.stack.Push(r.pc.Get())

	r.pc.Set(nnn)
}
