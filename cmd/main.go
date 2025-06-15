package main

import (
	"chip8/internal/chip8"
	ioPkg "chip8/internal/io"
	"context"
	"flag"
	"log/slog"
	"os"
	"sync"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	var rom string

	flag.StringVar(&rom, "rom", "", "Path to rom which you want to run")
	flag.Parse()

	if rom == "" {
		panic("Path to rom is required. Run with '-rom' parameter.")
	}

	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := chip8.NewRuntime(logger)
	r.LoadRom(rom)

	wg.Add(1)
	go func() {
		defer wg.Done()
		r.Run(ctx)
	}()

	io := ioPkg.NewIO(r, logger)
	io.Run()

	// Graceful shutdown runtime
	cancel()

	wg.Wait()

	logger.Info("Bye!")
}
