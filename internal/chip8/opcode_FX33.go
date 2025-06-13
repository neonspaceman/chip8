package chip8

// OpcodeFX33
// LD B, VX
// Store the binary-coded decimal in VX and put it in three consecutive memory slots starting at I.
// VX is a byte, so it is in 0…255. The interpreter takes the value in VX (for example the decimal value 174, or 0xAE in hex),
// converts it into a decimal and separates the hundreds, the tens and the ones (1, 7 and 4 respectively).
// Then, it stores them in three memory locations starting at I (1 to I, 7 to I+1 and 4 to I+2).
type OpcodeFX33 struct {
}

func (i *OpcodeFX33) Support(opcode uint16) bool {
	return (opcode & 0xF0FF) == 0xF033
}

func (i *OpcodeFX33) Do(opcode uint16, r *Runtime) {
	vx := opcodeX(opcode)

	value := r.v[vx]

	hundreds := [3]uint8{}

	hundreds[0] = value / 100
	hundreds[1] = value % 100 / 10
	hundreds[2] = value % 10

	address := r.i

	for _, v := range hundreds {
		r.mem.Write(address, v)
		address += 1
	}

	r.pc.Next()
}
