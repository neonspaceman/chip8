package chip8

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"time"
	"unsafe"
)

type InstructionInterface interface {
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
	instructions []InstructionInterface
	pc           PC
	i            uint16    // The index register I (16-bit), used to store memory addresses
	v            [16]uint8 // 16 general purpose 8-bit registers
	dt           uint8     // Delay timer
	st           uint8     // Sound timer

	videoBuffer VideoBuffer

	delayCounter time.Duration
	soundCounter time.Duration
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
		instructions: []InstructionInterface{
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
	runtimeTicker := time.NewTicker(time.Second / RuntimeSpeed)

	lastTick := time.Now()

	//for range 10 {
	for {
		currentTick := time.Now()
		dt := currentTick.Sub(lastTick)
		lastTick = currentTick

		r.updateDt(dt)
		r.updateSt(dt)

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
		case <-runtimeTicker.C:
		}
	}
}

func (r *Runtime) updateDt(dt time.Duration) {
	r.delayCounter += dt

	if r.delayCounter > time.Second/DelayTimerSpeed {
		if int(r.dt-1) < 0 {
			r.dt = 0
		} else {
			r.dt--
		}
		r.delayCounter = 0
	}
}

func (r *Runtime) updateSt(dt time.Duration) {
	r.soundCounter += dt

	if r.soundCounter > time.Second/SoundTimerSpeed {
		if int(r.st-1) < 0 {
			r.st = 0
		} else {
			r.dt--
		}
		r.soundCounter = 0
	}
}
