package io

import "math"

const (
	sampleRate = 48000
	frequency  = 440
)

// SineWave is an infinite SineWave of 440 Hz sine wave.
type SineWave struct {
	pos int64
}

// Read is io.Reader's Read.
//
// Read fills the data with sine wave samples.
func (s *SineWave) Read(buf []byte) (int, error) {
	const bytesPerSample = 8

	n := len(buf) / bytesPerSample * bytesPerSample

	const length = sampleRate / frequency
	for i := 0; i < n/bytesPerSample; i++ {
		v := math.Float32bits(float32(math.Sin(2 * math.Pi * float64(s.pos/bytesPerSample+int64(i)) / length)))
		buf[8*i] = byte(v)
		buf[8*i+1] = byte(v >> 8)
		buf[8*i+2] = byte(v >> 16)
		buf[8*i+3] = byte(v >> 24)
		buf[8*i+4] = byte(v)
		buf[8*i+5] = byte(v >> 8)
		buf[8*i+6] = byte(v >> 16)
		buf[8*i+7] = byte(v >> 24)
	}

	s.pos += int64(n)
	s.pos %= length * bytesPerSample

	return n, nil
}

// Close is io.Closer's Close.
func (s *SineWave) Close() error {
	return nil
}
