package ppu

import (
	"github.com/egaban/nes-go/internal/cartridge"
	"github.com/egaban/nes-go/internal/sdl"
)

type Ppu struct {
	bus       Bus
	registers Registers
	renderer  *sdl.Renderer

	firstWrite bool
	dataBuffer byte
}

func NewPpu(renderer *sdl.Renderer) *Ppu {
	return &Ppu{
		bus:        newPpuBus(),
		renderer:   renderer,
		firstWrite: true,
	}
}

func (p *Ppu) LoadCartridge(cartridge *cartridge.Cartridge) {
	p.bus.loadCartridge(cartridge)
}

func (p *Ppu) Tick() {}

func (p *Ppu) ClearScreen() {
	p.renderer.Clear()
}

func (p *Ppu) Destroy() {
	p.renderer.Destroy()
}
