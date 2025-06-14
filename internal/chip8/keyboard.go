package chip8

import (
	"sync"
)

type Keyboard struct {
	mu       sync.RWMutex
	ch       chan uint8
	keyboard [16]bool
}

func NewKeyboard() Keyboard {
	return Keyboard{
		mu: sync.RWMutex{},
		ch: nil,
	}
}

func (k *Keyboard) WaitKey() <-chan uint8 {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.ch != nil {
		panic("Undefined behaviour, keyboard channel is already opened")
	}

	k.ch = make(chan uint8)

	return k.ch
}

func (k *Keyboard) SendKey(key uint8, keyPressed bool) {
	k.mu.Lock()
	defer k.mu.Unlock()

	// If channel is open, wait new key
	if k.ch != nil {
		if !keyPressed {
			k.ch <- key
			k.ch = nil
		}
		return
	}

	k.keyboard[key] = keyPressed
}

func (k *Keyboard) IsPressed(key uint8) bool {
	k.mu.RLock()
	defer k.mu.RUnlock()

	return k.keyboard[key]
}
