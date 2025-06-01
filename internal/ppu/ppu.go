package ppu

type Ppu struct {
	bus Bus
}

func NewPpu() Ppu {
	return Ppu{
		bus: newPpuBus(),
	}
}

func (p *Ppu) ReadByteAt(address uint16) byte {
	panic("PPU write not implemented")
}

func (p *Ppu) WriteByteAt(address uint16, data byte) {

}
