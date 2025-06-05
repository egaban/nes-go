package ppu

import (
	"log/slog"

	"github.com/egaban/nesgo/internal/cartridge"
	"github.com/egaban/nesgo/internal/sdl"
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

// This is the CPU interface for the PPU.
func (p *Ppu) ReadRegister(address uint16) byte {
	address &= 0x0007 // Just to be sure the CPU isn't doing anything wrong.

	switch address {
	case 0x0000: // PPU Control Register
		slog.Warn("Trying to read PPU Control Register from the CPU")
		return 0
	case 0x0001: // PPU Mask Register
		slog.Warn("Trying to read PPU Mask Register from the CPU")
		return 0
	case 0x0002: // PPU Status Register
		p.registers.Status |= statusVerticalBlank // TEMPORARY
		result := (p.registers.Status & 0xE0) | (p.dataBuffer & 0x1F)
		p.registers.Status &^= statusVerticalBlank
		p.firstWrite = true
		return result
	case 0x0003: // OAM Address Register
		slog.Warn("Trying to read OAM Address Register from the CPU")
		return 0
	case 0x0004: // OAM Data Register
		return 0
	case 0x0005: // PPU Scroll Register
		slog.Warn("Trying to read PPU Scroll Register from the CPU")
		return 0
	case 0x0006: // PPU Address Register
		slog.Warn("Trying to read PPU Address Register from the CPU")
		return 0
	case 0x0007: // PPU Data Register
		result := p.dataBuffer
		p.dataBuffer = p.bus.ReadByteAt(p.address)

		if p.address < 0x3F00 {
			result = p.dataBuffer
		}

		return result
	default:
		panic("Unknown PPU register address")
	}
}

// This is the CPU interface for the PPU.
func (p *Ppu) WriteRegister(address uint16, data byte) {
	address &= 0x0007 // Just to be sure the CPU isn't doing anything wrong.

	switch address {
	case 0x0000: // PPU Control Register
		p.registers.Control = data
	case 0x0001: // PPU Mask Register
		p.registers.Mask = data
	case 0x0002: // PPU Status Register
		p.firstWrite = false
		slog.Warn("Trying to write read-only PPU Status Register from the CPU")
	case 0x0003: // OAM Address Register
		break
	case 0x0004: // OAM Data Register
		break
	case 0x0005: // PPU Scroll Register
		p.setScrollRegister(data)
	case 0x0006: // PPU Address Register
		p.setAddressRegister(data)
	case 0x0007: // PPU Data Register
		p.bus.WriteByteAt(p.address, data)
	}
}

func (p *Ppu) Destroy() {
	p.renderer.Destroy()
}
