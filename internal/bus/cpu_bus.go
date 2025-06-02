package bus

import (
	"github.com/egaban/nesgo/internal/cartridge"
	"github.com/egaban/nesgo/internal/interfaces"
)

type CpuBus struct {
	ram       [2 * 1024]byte
	ppu       interfaces.PpuRegisters
	cartridge *cartridge.Cartridge
}

func NewBus(ppu interfaces.PpuRegisters) *CpuBus {
	result := &CpuBus{
		ram:       [2 * 1024]byte{0},
		cartridge: nil,
		ppu:       ppu,
	}

	return result
}

func (b *CpuBus) LoadCartridge(cartridge *cartridge.Cartridge) {
	b.cartridge = cartridge
}

func (b *CpuBus) ReadByteAt(address uint16) byte {
	if value, success := b.cartridge.TryReadPrgAt(address); success {
		return value
	} else if address < 0x2000 {
		return b.ram[address&0x7FF]
	} else if address < 0x4000 {
		return b.ppu.ReadRegister(address & 0x0007)
	}
	return 0x00
}

func (b *CpuBus) WriteByteAt(address uint16, data uint8) {
	if b.cartridge.TryWritePrgAt(address, data) {
		return
	} else if address < 0x2000 {
		b.ram[address&0x7FF] = data
	} else if address < 0x4000 {
		b.ppu.WriteRegister(address&0x0007, data)
	}
}

func (b *CpuBus) ReadWordAt(address uint16) uint16 {
	lo := b.ReadByteAt(address)
	hi := b.ReadByteAt(address + 1)

	return uint16(lo) | (uint16(hi) << 8)
}

func (b *CpuBus) ReadSamePageWord(address uint16) uint16 {
	page := 0xFF00 & address
	lo := b.ReadByteAt(address)
	hi_address := page | (0x00FF & (address + 1))
	hi := b.ReadByteAt(hi_address)

	return uint16(lo) | (uint16(hi) << 8)
}

func (b *CpuBus) WriteWordAt(address uint16, data uint16) {
	lo := uint8(data)
	hi := uint8(data >> 8)

	b.WriteByteAt(address, lo)
	b.WriteByteAt(address+1, hi)
}
