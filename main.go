package main

import (
	"fmt"
	"runtime"

	"github.com/egaban/nes-go/internal/cartridge"
	"github.com/egaban/nes-go/internal/nes"
)

// MacOS requires the main thread to be locked to the OS thread
func init() {
	runtime.LockOSThread()
}

func main() {
	cartridge, err := cartridge.LoadCartridge("nestest.nes")

	if err != nil {
		fmt.Printf("Error loading cartridge: %v\n", err)
	}

	emulator := nes.NewEmulator(cartridge)
	emulator.Run()
}
