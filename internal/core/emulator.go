package core

import (
	"github.com/egaban/nesgo/internal/bus"
	"github.com/egaban/nesgo/internal/cartridge"
	"github.com/egaban/nesgo/internal/cpu"
)

type Emulator struct {
	cpu *cpu.Cpu
	bus *bus.Bus
}

func NewEmulator(cartridge *cartridge.Cartridge) *Emulator {
	bus := bus.NewBus(cartridge)
	return &Emulator{
		cpu: cpu.NewCpu(bus),
		bus: bus,
	}
}

func (emulator *Emulator) Run() {
	for {
		emulator.cpu.Tick()
	}
}
