package mappers

type Mapper000 struct {
	numPrgBanks int
	numChrBanks int
}

// MapPrgRead implements Mapper.
func (m *Mapper000) MapPrgRead(address uint16) (uint32, bool) {
	if address >= 0x8000 {
		if m.numPrgBanks == 1 {
			return uint32(address & 0x3FFF), true
		}
		return uint32(address & 0x7FFF), true
	}
	return 0, false
}

// MapPrgWrite implements Mapper.
func (m *Mapper000) MapPrgWrite(address uint16) (uint32, bool) {
	if address >= 0x8000 {
		if m.numPrgBanks == 1 {
			return uint32(address & 0x3FFF), true
		}
		return uint32(address & 0x7FFF), true
	}
	return 0, false
}

// MapChrRead implements Mapper.
func (m *Mapper000) MapChrRead(address uint16) (uint32, bool) {
	if address < 0x2000 {
		return uint32(address), true
	}
	return 0, false
}

// MapChrWrite implements Mapper.
func (m *Mapper000) MapChrWrite(address uint16) (uint32, bool) {
	return 0, false
}

func NewMapper000(numPrgBanks, numChrBanks int) Mapper {
	if numPrgBanks == 0 || numPrgBanks > 2 {
		panic("Invalid number of PRG banks: must be 1 or 2")
	}

	if numChrBanks != 1 {
		panic("Invalid number of CHR banks: must be 1")
	}

	return &Mapper000{
		numPrgBanks: numPrgBanks,
		numChrBanks: numChrBanks,
	}
}
