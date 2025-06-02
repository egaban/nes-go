package mappers

type Mapper interface {
	MapPrgRead(address uint16) (uint32, bool)
	MapPrgWrite(address uint16) (uint32, bool)

	MapChrRead(address uint16) (uint32, bool)
	MapChrWrite(address uint16) (uint32, bool)
}

type MapperFactory func() Mapper

func NewMapper(id byte, numPrgBanks, numChrBanks int) Mapper {
	switch id {
	case 0x00:
		return NewMapper000(numPrgBanks, numChrBanks)
	default:
		panic("Mapper not implemented")
	}
}
