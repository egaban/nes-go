package nes

import (
	"github.com/egaban/nesgo/internal/bus"
	"github.com/egaban/nesgo/internal/cartridge"
	"github.com/egaban/nesgo/internal/cpu"
	"github.com/egaban/nesgo/internal/ppu"
)

type Emulator struct {
	cpu *cpu.Cpu
	ppu *ppu.Ppu
	bus *bus.CpuBus

	totalCycles int64
}

func NewEmulator(cartridge *cartridge.Cartridge) *Emulator {
	ppu := ppu.NewPpu()
	bus := bus.NewBus(ppu)

	ppu.LoadCartridge(cartridge)
	bus.LoadCartridge(cartridge)

	return &Emulator{
		cpu: cpu.NewCpu(bus),
		bus: bus,
		ppu: ppu,
	}
}

func (e *Emulator) Run() {
	for {
		e.Tick()
	}
}

func (e *Emulator) Reset() {
	e.cpu.Reset()
	e.totalCycles = 0
}

func (e *Emulator) Tick() {
	e.cpu.Tick()
	e.totalCycles += 1
}
