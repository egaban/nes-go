package ppu

import "log/slog"

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
		result := (p.registers.status & 0xE0)
		p.registers.status &^= statusVerticalBlank
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
		result := p.bus.ReadByteAt(p.registers.v)
		if p.registers.control&controlIncrement == 0 {
			p.registers.v += 1
		} else {
			p.registers.v += 32
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
		p.registers.control = data
		p.registers.t &^= 0x0C00                   // Clear bits 10-11
		p.registers.t |= (uint16(data&0x03) << 10) // Set bits 10-11 based on control register
	case 0x0001: // PPU Mask Register
		p.registers.mask = data
	case 0x0002: // PPU Status Register
		p.firstWrite = false
		slog.Warn("Write to $2002 (PPU Status Register) skipped")
	case 0x0003: // OAM Address Register
		slog.Warn("Write to $2003 (OAM Address Register) skipped")
	case 0x0004: // OAM Data Register
		slog.Warn("Write to $2004 (OAM Data Register) skipped")
	case 0x0005: // PPU Scroll Register
		p.setScrollRegister(data)
	case 0x0006: // PPU Address Register
		p.setAddressRegister(data)
	case 0x0007: // PPU Data Register
		p.bus.WriteByteAt(p.registers.v, data)
		if p.registers.control&controlIncrement == 0 {
			p.registers.v += 1
		} else {
			p.registers.v += 32
		}
	}
}
