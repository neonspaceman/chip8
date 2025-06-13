package chip8

import (
	"math"
	"math/rand"
)

// OpcodeCXNN
// RND VX, NN
// Generate a random byte (from 0 to 255), do a bitwise AND with NN and store the result to VX.
type OpcodeCXNN struct {
}

func (i *OpcodeCXNN) Support(opcode uint16) bool {
	return (opcode & 0xF000) == 0xC000
}

func (i *OpcodeCXNN) Do(opcode uint16, r *Runtime) {
	vx, nn := opcodeX(opcode), opcodeNN(opcode)

	r.v[vx] = uint8(rand.Intn(math.MaxUint8)) & nn

	r.pc.Next()
}
