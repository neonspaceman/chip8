package chip8

// Opcode00E0
// CLS
// clearVideoBuffer the display by setting all pixels to ‘off’.
type Opcode00E0 struct {
}

func (i *Opcode00E0) Support(opcode uint16) bool {
	return opcode == 0x00E0
}

func (i *Opcode00E0) Do(opcode uint16, r *Runtime) {
	r.videoBuffer.Clear()

	r.pc.Next()
}
