package chip8

import "unsafe"

// OpcodeFX65
// LD VX, [I]
// Load the memory data starting at address I into the registers V0 to VX.
type OpcodeFX65 struct {
}

func (i *OpcodeFX65) Support(opcode uint16) bool {
	return (opcode & 0xF0FF) == 0xF065
}

func (i *OpcodeFX65) Do(opcode uint16, r *Runtime) {
	vx := opcodeX(opcode)

	address := r.i

	for i := range vx + 1 {
		v := r.mem.Read(address)
		r.v[i] = v
		address += uint16(unsafe.Sizeof(v))
	}

	r.pc.Next()
}
