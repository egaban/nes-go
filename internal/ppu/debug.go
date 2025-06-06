package ppu

import "github.com/egaban/nes-go/internal/sdl"

func (p *Ppu) getPatternTable(patternTable int) [256][16]byte {
	if patternTable < 0 || patternTable > 1 {
		panic("Invalid pattern table index")
	}

	var table [256][16]byte

	for i := 0; i < 256; i++ {
		for j := 0; j < 16; j++ {
			address := uint16(patternTable*0x1000 + i*16 + j)
			table[i][j] = p.bus.ReadByteAt(address)
		}
	}

	return table
}

var temporaryPalette = []sdl.Color{
	{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}, // Black
	{R: 0x55, G: 0x55, B: 0x55, A: 0xFF}, // Dark gray
	{R: 0xAA, G: 0xAA, B: 0xAA, A: 0xFF}, // Light gray
	{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, // White
}

func (p *Ppu) drawTile(tile [16]byte, x, y int) {
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			loPixelValue := (tile[row] & (1 << (7 - col))) >> (7 - col)
			hiPixelValue := (tile[row+8] & (1 << (7 - col))) >> (7 - col)

			pixelValue := hiPixelValue<<1 | loPixelValue
			color := temporaryPalette[pixelValue]
			p.renderer.DrawPixel(x+col, y+row, color)
		}
	}
}

func (p *Ppu) RenderPatternTable(patternTableId int, x, y int) {
	patternTable := p.getPatternTable(patternTableId)

	for tile := 0; tile < 256; tile++ {
		p.drawTile(patternTable[tile], x+(tile%16)*8, y+(tile/16)*16)
	}

	p.renderer.Present()
}
