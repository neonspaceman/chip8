package chip8

// OpcodeFX55
// LD [I], VX
// Store registers from V0 to VX in the main memory, starting at location I.
// Note that X is the number of the register, so we can use it in the loop.
type OpcodeFX55 struct {
}

func (i *OpcodeFX55) Support(opcode uint16) bool {
	return (opcode & 0xF0FF) == 0xF055
}

func (i *OpcodeFX55) Do(opcode uint16, r *Runtime) {
	vx := opcodeX(opcode)

	address := r.i

	for i := range vx + 1 {
		v := r.v[i]
		r.mem.Write(address, v)
		address += 1
	}

	r.pc.Next()
}
