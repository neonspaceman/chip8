package chip8

func opcodeX(opcode uint16) uint8 {
	return uint8((opcode & 0x0F00) >> 8)
}

func opcodeY(opcode uint16) uint8 {
	return uint8((opcode & 0x00F0) >> 4)
}

func opcodeN(opcode uint16) uint8 {
	return uint8(opcode & 0x000F)
}

func opcodeNN(opcode uint16) uint8 {
	return uint8(opcode & 0x00FF)
}

func opcodeNNN(opcode uint16) uint16 {
	return opcode & 0x0FFF
}
