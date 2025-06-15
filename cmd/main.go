package main

import (
	"chip8/internal/chip8"
	ioPkg "chip8/internal/io"
	"flag"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	var rom string

	flag.StringVar(&rom, "rom", "", "Path to rom which you want to run")
	flag.Parse()

	if rom == "" {
		panic("Path to rom is required. Run with '-rom' parameter.")
	}

	r := chip8.NewRuntime(logger)
	r.LoadRom(rom)
	go r.Run()

	io := ioPkg.NewIO(r, logger)
	io.Run()
}
