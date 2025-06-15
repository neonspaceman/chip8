package io

import (
	"chip8/internal/chip8"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log/slog"
)

type RuntimeImpl interface {
	VideoBuffer() chip8.VideoBufferType
	SendKey(key uint8, keyPressed bool)
	Beep() bool
}

type IO struct {
	r RuntimeImpl

	pixels []byte

	pressedKeys  []ebiten.Key
	releasedKeys []ebiten.Key

	keyboardMapping map[ebiten.Key]uint8

	beep *audio.Player

	logger *slog.Logger
}

func NewIO(r RuntimeImpl, logger *slog.Logger) IO {
	audioCtx := audio.NewContext(sampleRate)
	beep, err := audioCtx.NewPlayerF32(&SineWave{})

	if err != nil {
		panic(err)
	}

	return IO{
		r:      r,
		pixels: make([]byte, chip8.VideoBufferHeight*chip8.VideoBufferWidth*4),
		/*
			Chip8				QWERTY
			┌───┬───┬───┬───┐	┌───┬───┬───┬───┐
			│ 1 │ 2 │ 3 │ C │	│ 1 │ 2 │ 3 │ 4 │
			│ 4 │ 5 │ 6 │ D │	│ Q │ W │ D │ E │
			│ 7 │ 8 │ 9 │ E │	│ A │ S │ D │ F │
			│ A │ 0 │ B │ F │	│ Z │ X │ C │ V │
			└───┴───┴───┴───┘	└───┴───┴───┴───┘
		*/
		keyboardMapping: map[ebiten.Key]uint8{
			ebiten.KeyDigit1: 0x1, ebiten.KeyDigit2: 0x2, ebiten.KeyDigit3: 0x3, ebiten.KeyDigit4: 0xC,
			ebiten.KeyQ: 0x4, ebiten.KeyW: 0x5, ebiten.KeyE: 0x6, ebiten.KeyR: 0xD,
			ebiten.KeyA: 0x7, ebiten.KeyS: 0x8, ebiten.KeyD: 0x9, ebiten.KeyF: 0xE,
			ebiten.KeyZ: 0xA, ebiten.KeyX: 0x0, ebiten.KeyC: 0xB, ebiten.KeyV: 0xF,
		},
		beep:   beep,
		logger: logger,
	}
}

func (io *IO) Update() error {
	io.pressedKeys = inpututil.AppendJustPressedKeys(io.pressedKeys[:0])
	io.releasedKeys = inpututil.AppendJustReleasedKeys(io.releasedKeys[:0])

	for _, k := range io.pressedKeys {
		if key, ok := io.keyboardMapping[k]; ok {
			io.r.SendKey(key, true)
		}
	}

	for _, k := range io.releasedKeys {
		if key, ok := io.keyboardMapping[k]; ok {
			io.r.SendKey(key, false)
		}
	}

	if io.r.Beep() {
		io.beep.Play()
	} else {
		io.beep.Pause()
	}

	return nil
}

func (io *IO) Draw(screen *ebiten.Image) {
	buf := io.r.VideoBuffer()

	for row := range chip8.VideoBufferHeight {
		mask := uint64(1) << (chip8.VideoBufferWidth - 1)
		for col := range chip8.VideoBufferWidth {
			if (buf[row] & mask) > 0 {
				io.pixels[row*chip8.VideoBufferWidth*4+col*4] = 0xFF
				io.pixels[row*chip8.VideoBufferWidth*4+col*4+1] = 0xFF
				io.pixels[row*chip8.VideoBufferWidth*4+col*4+2] = 0xFF
				io.pixels[row*chip8.VideoBufferWidth*4+col*4+3] = 0xFF
			} else {
				io.pixels[row*chip8.VideoBufferWidth*4+col*4] = 0x00
				io.pixels[row*chip8.VideoBufferWidth*4+col*4+1] = 0x00
				io.pixels[row*chip8.VideoBufferWidth*4+col*4+2] = 0x00
				io.pixels[row*chip8.VideoBufferWidth*4+col*4+3] = 0x00
			}
			mask >>= 1
		}
	}

	screen.WritePixels(io.pixels)
}

func (io *IO) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 64, 32
}

func (io *IO) Run() {
	io.logger.Info("Run the io")

	ebiten.SetWindowSize(1024, 512)
	ebiten.SetWindowTitle("Chip8")

	if err := ebiten.RunGame(io); err != nil {
		panic(err)
	}

	io.logger.Info("The io is stopped")
}
