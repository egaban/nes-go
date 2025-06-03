package ppu

import "github.com/egaban/nesgo/internal/cartridge"

type Bus struct {
	cartridge *cartridge.Cartridge

	nameTables [2][1024]byte // Two name tables, each 1024 bytes
	palette    [32]byte      // Palette memory, 32 bytes
}

func newPpuBus() Bus {
	return Bus{}
}

func (b *Bus) ReadByteAt(address uint16) byte {
	address &= 0x3FFF // PPU memory space is 16KB, so we mask to that range

	if value, success := b.cartridge.TryReadChrAt(address); success {
		return value
	}

	panic("Not implemented: PPU read operation")
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
