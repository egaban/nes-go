package nes

import (
	"time"

	"github.com/egaban/nes-go/internal/bus"
	"github.com/egaban/nes-go/internal/cartridge"
	"github.com/egaban/nes-go/internal/cpu"
	"github.com/egaban/nes-go/internal/ppu"
	"github.com/egaban/nes-go/internal/sdl"
)

type Emulator struct {
	cpu    *cpu.Cpu
	ppu    *ppu.Ppu
	bus    *bus.CpuBus
	window *sdl.Window

	totalCycles int64
	shouldStop  bool
}

func NewEmulator(cartridge *cartridge.Cartridge) *Emulator {
	sdl.Init()
	window := sdl.CreateWindow()

	ppu := ppu.NewPpu(window.CreateRenderer())
	bus := bus.NewBus(ppu)

	ppu.LoadCartridge(cartridge)
	bus.LoadCartridge(cartridge)

	return &Emulator{
		cpu:        cpu.NewCpu(bus),
		bus:        bus,
		ppu:        ppu,
		window:     window,
		shouldStop: false,
	}
}

func (e *Emulator) Run() {
	for {
		e.Tick()
		e.ppu.ClearScreen()
		e.ppu.RenderPatternTable(0, 240, 0)
		e.ppu.RenderPatternTable(1, 240+128, 0)
		if e.shouldStop {
			e.Destroy()
			return
		}
	}
}

func (e *Emulator) Reset() {
	e.cpu.Reset()
	e.totalCycles = 0
}

func (e *Emulator) Tick() {
	if e.totalCycles%3 == 0 {
		e.cpu.Tick()
	}

	events := e.window.PollEvents()

	for _, event := range events {
		if event == sdl.QUIT {
			e.shouldStop = true
		}
	}

	e.ppu.Tick()

	// Sleep for 1ms
	time.Sleep(1 * time.Millisecond)
	e.totalCycles += 1
}

func (e *Emulator) Destroy() {
	e.window.Destroy()
	e.ppu.Destroy()
}
