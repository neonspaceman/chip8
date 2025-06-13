package chip8

// OpcodeDXYN
// DRW VX, VY, N
// The draw instruction. This is arguably the most involved operation.
// The n-byte sprite starting at the address I is drawn to the display at the coordinates [VX, VY].
// Then, set VF to 1 if there has been a collision (a display bit was changed from 1 to 0).
// The interpreter must read N bytes from the I address in memory.
// These bytes are interpreted as a sprite and drawn at the display coordinates [VX, VY].
// The bits are set using an XOR with the current display state.
type OpcodeDXYN struct {
}

func (i *OpcodeDXYN) Support(opcode uint16) bool {
	return (opcode & 0xF000) == 0xD000
}

func (i *OpcodeDXYN) Do(opcode uint16, r *Runtime) {
	vx, vy, n := opcodeX(opcode), opcodeY(opcode), opcodeN(opcode)

	sprite := make([]byte, n)

	address := r.i

	for i := range n {
		v := r.mem.Read(address)
		sprite[i] = v
		address++
	}

	if r.videoBuffer.Render(r.v[vx], r.v[vy], sprite) {
		r.v[0xF] = 1
	} else {
		r.v[0xF] = 0
	}

	r.pc.Next()
}
