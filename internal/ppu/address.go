package ppu

func (p *Ppu) setAddressRegister(data byte) {
	if p.firstWrite {
		// Clears the first 8 bits. The PPU only has 14-bit address space.
		p.registers.t &^= 0xFF00
		p.registers.t |= (uint16(data&0x1F) << 8)
	} else {
		p.registers.t &^= 0x00FF
		p.registers.t |= uint16(data)
		p.registers.v = p.registers.t
	}
	p.firstWrite = !p.firstWrite
}
