package ppu

func (p *Ppu) setScrollRegister(data byte) {
	if p.firstWrite {
		p.registers.x = data & 0x07
		p.updateCoarseX(data >> 3)
	} else {
		p.updateFineY(data & 0x07)
		p.updateCoarseY(data >> 3)
	}

	p.firstWrite = !p.firstWrite
}

// Updates the coarse X. This should be the 5-bit value, not the whole byte.
func (p *Ppu) updateCoarseX(value byte) {
	// Clears Bits 0-4
	p.registers.t &^= 0x001F
	p.registers.t |= uint16(value & 0x1F)
}

// Updates the coarse Y. This should be the 5-bit value, not the whole byte.
func (p *Ppu) updateCoarseY(value byte) {
	// Clears bits 5-9
	p.registers.t &^= 0x03E0
	p.registers.t |= (uint16(value) << 5)
}

// Updates fine Y scroll. This should be the 3-bit value, not the whole byte.
func (p *Ppu) updateFineY(value byte) {
	// Clears bits 12-14
	p.registers.t &^= 0x7000
	p.registers.t |= (uint16(value) << 12)
}
