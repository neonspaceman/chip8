package chip8

// OpcodeFX29
// LD F, VX
// Set the location of the sprite for the digit VX to I.
// The font sprites start at address 0x000, and contain the hexadecimal digits from 1..F.
// Each font has a length of 0x05 bytes. The memory address for the value in VX is put in I.
type OpcodeFX29 struct {
}

func (i *OpcodeFX29) Support(opcode uint16) bool {
	return (opcode & 0xF0FF) == 0xF029
}

func (i *OpcodeFX29) Do(opcode uint16, r *Runtime) {
	vx := opcodeX(opcode)

	r.i = uint16(r.v[vx] * 0x05)

	r.pc.Next()
}
