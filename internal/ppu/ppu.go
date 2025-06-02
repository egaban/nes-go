package ppu

import "github.com/egaban/nesgo/internal/cartridge"

type Ppu struct {
	bus Bus
}

func NewPpu() *Ppu {
	return &Ppu{
		bus: newPpuBus(),
	}
}

func (p *Ppu) LoadCartridge(cartridge *cartridge.Cartridge) {
	p.bus.loadCartridge(cartridge)
}

func (p *Ppu) Tick() {}

// This is the CPU interface for the PPU.
func (p *Ppu) ReadRegister(address uint16) byte {
	address &= 0x0007 // Just to be sure the CPU isn't doing anything wrong.

	switch address {
	case 0x0000: // PPU Control Register
		break
	case 0x0001: // PPU Mask Register
		break
	case 0x0002: // PPU Status Register
		break
	case 0x0003: // OAM Address Register
		break
	case 0x0004: // OAM Data Register
		break
	case 0x0005: // PPU Scroll Register
		break
	case 0x0006: // PPU Address Register
		break
	case 0x0007: // PPU Data Register
		break
	}

	panic("PPU write not implemented")
}

// This is the CPU interface for the PPU.
func (p *Ppu) WriteRegister(address uint16, data byte) {
	address &= 0x0007 // Just to be sure the CPU isn't doing anything wrong.

	switch address {
	case 0x0000: // PPU Control Register
		break
	case 0x0001: // PPU Mask Register
		break
	case 0x0002: // PPU Status Register
		break
	case 0x0003: // OAM Address Register
		break
	case 0x0004: // OAM Data Register
		break
	case 0x0005: // PPU Scroll Register
		break
	case 0x0006: // PPU Address Register
		break
	case 0x0007: // PPU Data Register
		break
	}

	panic("PPU write not implemented")
}
