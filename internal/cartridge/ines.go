package cartridge

func (c *Cartridge) parseInes1(data []byte) {
	hasTrainer := data[6]&0x04 != 0

	if hasTrainer {
		panic("iNES 1.0 with trainer not supported yet")
	}

	prgRomStart := 16
	prgRomSize := c.numPrgBanks * 16 * 1024
	prgRomEnd := prgRomStart + prgRomSize

	if prgRomEnd > len(data) {
		panic("PRG ROM data exceeds input size")
	}
	c.PrgRom = data[prgRomStart:prgRomEnd]

	chrRomStart := prgRomEnd
	chrRomSize := c.numChrBanks * 8 * 1024
	chrRomEnd := chrRomStart + chrRomSize

	if chrRomEnd > len(data) {
		panic("CHR ROM data exceeds input size")
	}
	c.ChrRom = data[chrRomStart:chrRomEnd]
}
