package chip8

import (
	tm "github.com/buger/goterm"
	"time"
)

type Display struct {
	r *Runtime
}

func NewDisplay(r *Runtime) Display {
	return Display{
		r: r,
	}
}

func (d *Display) Display() {
	tm.Clear() // clearVideoBuffer current screen

	for {
		//By moving cursor to top-left position we ensure that console output
		//will be overwritten each time, instead of adding new.
		tm.MoveCursor(1, 1)

		buf := d.r.videoBuffer.Get()

		for i := range VideoBufferHeight {
			//tm.Println(fmt.Sprintf("%064b", buf[i]))
			mask := uint64(1) << (VideoBufferWidth - 1)
			for range VideoBufferWidth {
				if (buf[i] & mask) > 0 {
					tm.Print("▉▉")
				} else {
					tm.Print("  ")
				}
				mask >>= 1
			}
			tm.Println()
		}

		tm.Flush() // Call it every time at the end of rendering

		time.Sleep(time.Second / 24)
	}
}
