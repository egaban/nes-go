package main

import (
	"fmt"

	"github.com/egaban/nesgo/internal/cartridge"
	"github.com/egaban/nesgo/internal/core"
)

func main() {
	cartridge, err := cartridge.LoadCartridge("nestest.nes")

	if err != nil {
		fmt.Printf("Error loading cartridge: %v\n", err)
	}

	emulator := core.NewEmulator(cartridge)
	emulator.Run()
}
