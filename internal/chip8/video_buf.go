package chip8

import (
	"sync"
)

const VideoBufferWidth = 64
const VideoBufferHeight = 32

type VideoBufferType [32]uint64

type VideoBuffer struct {
	mu  sync.Mutex
	buf VideoBufferType
}

func (v *VideoBuffer) Render(x, y uint8, sprite []byte) bool {
	v.mu.Lock()
	defer v.mu.Unlock()

	x %= VideoBufferWidth

	collision := false

	for _, s := range sprite {
		shift := VideoBufferWidth - int(x) - 8 // 8 - length of byte

		renderSpriteLine := uint64(s)

		renderLine := &v.buf[y%VideoBufferHeight]

		if shift > 0 {
			renderSpriteLine <<= shift
		} else {
			renderSpriteLine >>= -shift
		}

		// print sprite on the scree
		collision = collision || ((*renderLine & renderSpriteLine) > 0)
		*renderLine = *renderLine ^ renderSpriteLine

		// print outside part of sprite on opposite part of screen
		if shift < 0 {
			renderSpriteLine = uint64(s) << (VideoBufferWidth + shift)
			collision = collision || ((*renderLine & renderSpriteLine) > 0)
			*renderLine = *renderLine ^ renderSpriteLine
		}

		y++
	}

	return collision
}

func (v *VideoBuffer) Clear() {
	v.mu.Lock()
	defer v.mu.Unlock()

	for i := range VideoBufferHeight {
		v.buf[i] = 0
	}
}

func (v *VideoBuffer) Get() VideoBufferType {
	v.mu.Lock()
	defer v.mu.Unlock()

	return v.buf
}
