package chip8

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"time"
	"unsafe"
)

type InstructionImpl interface {
	Support(opcode uint16) bool
	Do(opcode uint16, r *Runtime)
}

type DisplayInterface interface {
}

const MemOffset uint16 = 0x200
const MemSize uint16 = 2 << 11 // 4096KB
const StackSize = 16
const RuntimeSpeed = 540   // Hz (540 per sec)
const DelayTimerSpeed = 60 // Hz (60 per sec)
const SoundTimerSpeed = 60 // Hz (60 per sec)

type Runtime struct {
	stack        Stack
	mem          Mem
	instructions []InstructionImpl
	pc           PC
	i            uint16    // The index register I (16-bit), used to store memory addresses
	v            [16]uint8 // 16 general purpose 8-bit registers
	dt           uint8     // Delay timer
	st           uint8     // Sound timer

	videoBuffer VideoBuffer

	keyboard Keyboard
}

func NewRuntime() *Runtime {
	fonts := [...]uint8{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0xF0, // C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}

	r := Runtime{
		pc:    NewPC(MemOffset),
		mem:   NewMem(MemSize),
		stack: NewStack(StackSize),
		instructions: []InstructionImpl{
			&Opcode00E0{},
			&Opcode00EE{},
			&Opcode1NNN{},
			&Opcode2NNN{},
			&Opcode3XNN{},
			&Opcode4XNN{},
			&Opcode5XY0{},
			&Opcode6XNN{},
			&Opcode7XNN{},
			&Opcode8XY0{},
			&Opcode8XY1{},
			&Opcode8XY2{},
			&Opcode8XY3{},
			&Opcode8XY4{},
			&Opcode8XY5{},
			&Opcode8XY6{},
			&Opcode8XY7{},
			&Opcode8XYE{},
			&Opcode9XY0{},
			&OpcodeANNN{},
			&OpcodeBNNN{},
			&OpcodeCXNN{},
			&OpcodeDXYN{},
			&OpcodeEX9E{},
			&OpcodeEXA1{},
			&OpcodeFX0A{},
			&OpcodeFX1E{},
			&OpcodeFX07{},
			&OpcodeFX15{},
			&OpcodeFX18{},
			&OpcodeFX29{},
			&OpcodeFX33{},
			&OpcodeFX55{},
			&OpcodeFX65{},
		},
		keyboard: NewKeyboard(),
	}

	address := uint16(0x0)

	for _, v := range fonts {
		r.mem.Write(address, v)
		address++
	}

	return &r
}

func (r *Runtime) LoadRom(filepath string) {
	f, err := os.Open(filepath)

	if err != nil {
		panic("Failed to load rom")
	}

	var opcode uint8

	offset := MemOffset

	for {
		err = binary.Read(f, binary.NativeEndian, &opcode)

		if err == io.EOF {
			break
		}

		if err != nil {
			panic("Failed to load rom: " + err.Error())
		}

		r.mem.Write(offset, opcode)

		offset += uint16(unsafe.Sizeof(opcode))
	}
}

func (r *Runtime) Run() {
	go r.updateDt()
	go r.updateSt()

	ticker := time.NewTicker(time.Second / RuntimeSpeed)

	for {
		// 2 bytes, big endian
		opcode := (uint16(r.mem.Read(r.pc.Get())) << 8) | uint16(r.mem.Read(r.pc.Get()+1))

		opcodeProceeded := false

		for _, i := range r.instructions {
			if i.Support(opcode) {
				opcodeProceeded = true
				i.Do(opcode, r)
			}
		}

		if !opcodeProceeded {
			panic(fmt.Sprintf("UNKNOWN opcode: %04X", opcode))
		}

		select {
		case <-ticker.C:
		}
	}
}

func (r *Runtime) VideoBuffer() VideoBufferType {
	return r.videoBuffer.Get()
}

func (r *Runtime) SendKey(key uint8, keyPressed bool) {
	r.keyboard.SendKey(key, keyPressed)
}

func (r *Runtime) updateDt() {
	ticker := time.NewTicker(time.Second / DelayTimerSpeed)

	for {
		if r.dt > 0 {
			r.dt = r.dt - 1
		}

		select {
		case <-ticker.C:
		}
	}
}

func (r *Runtime) updateSt() {
	ticker := time.NewTicker(time.Second / SoundTimerSpeed)

	for {
		if r.st > 0 {
			r.st = r.st - 1
		}

		select {
		case <-ticker.C:
		}
	}
}
