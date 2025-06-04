package ppu

import "github.com/egaban/nesgo/internal/cartridge"

type Bus struct {
	cartridge *cartridge.Cartridge

	nametableVRAM [2048]byte // Two name tables, each 1024 bytes
	paletteVRAM   [32]byte   // Palette memory, 32 bytes
}

func newPpuBus() Bus {
	return Bus{}
}

func (b *Bus) ReadByteAt(address uint16) byte {
	address &= 0x3FFF // PPU memory space is 16KB, so we mask to that range

	if value, success := b.cartridge.TryReadChrAt(address); success {
		return value
	}

	// Nametable VRAM (Almost 8kb range, but actually 1kb per nametable)
	// 0x3000-0x3EFF is always a mirror of 0x2000-0x2EFF.
	if address >= 0x2000 && address < 0x3F00 {
		offset := address & 0x0FFF
		address = 0x2000 | offset // Now we are in the 0x2000-0x2FFF range

		var effectiveAddress uint16
		switch b.cartridge.GetNametableMirroring() {
		case cartridge.HorizontalMirroring:
			// Offsets 0x0400 if address is in the second nametable
			nametableOffset := (address & 0x0800) >> 1
			effectiveAddress = (address & 0x03FF) | nametableOffset
		case cartridge.VerticalMirroring:
			effectiveAddress = address & 0x07FF
		}
		return b.nametableVRAM[effectiveAddress]
	}

	panic("Not implemented: PPU palette (and mirror)")
}

func (b *Bus) WriteByteAt(address uint16, data byte) {
	address &= 0x3FFF // PPU memory space is 16KB, so we mask to that range

	if b.cartridge.TryWriteChrAt(address, data) {
		return
	}
}

func (b *Bus) loadCartridge(cartridge *cartridge.Cartridge) {
	b.cartridge = cartridge
}
