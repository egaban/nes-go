package bus

import (
	"errors"

	"github.com/egaban/nesgo/internal/cartridge"
)

func (bus *Bus) initVector() uint16 {
	return uint16(bus.memory[0xFFFD]) | (uint16(bus.memory[0xFFFC]) << 8)
}

func (bus *Bus) loadCartridge(cartridge *cartridge.Cartridge) {
	switch cartridge.NumPrgBanks {
	case 1:
		copy(bus.memory[0x8000:0xC000], cartridge.PrgRom)
		copy(bus.memory[0xC000:], cartridge.PrgRom)
	case 2:
		copy(bus.memory[0x8000:], cartridge.PrgRom)
	default:
		panic("Unsupported number of PRG banks.")
	}

}

func (bus *Bus) ReadByteAt(address uint16) (byte, error) {
	if int(address) >= len(bus.memory) {
		return 0, errors.New("Read out of bounds.")
	}
	return bus.memory[address], nil
}

func (bus *Bus) ReadWordAt(address uint16) (uint16, error) {
	if int(address)+2 >= len(bus.memory) {
		return 0, errors.New("Read out of bounds.")
	}

	lo := bus.memory[address]
	hi := bus.memory[address+1]

	return uint16(lo) | (uint16(hi) << 8), nil
}

func (bus *Bus) ReadSamePageWord(address uint16) (uint16, error) {
	if int(address) >= len(bus.memory)-1 {
		return 0, errors.New("Read out of bounds.")
	}

	page := 0xFF00 & address
	lo := bus.memory[address]
	hi_address := page | (0x00FF & (address + 1))
	hi := bus.memory[hi_address]

	return uint16(lo) | (uint16(hi) << 8), nil
}

func (bus *Bus) WriteByteAt(address uint16, data uint8) error {
	if int(address) >= len(bus.memory) {
		return errors.New("Write out of bounds")
	}

	bus.memory[address] = data
	return nil
}

func (bus *Bus) WriteWordAt(address uint16, data uint16) error {
	if int(address)+1 >= len(bus.memory) {
		return errors.New("Write out of bounds")
	}

	lo := uint8(data)
	hi := uint8(data >> 8)

	bus.memory[address] = lo
	bus.memory[address+1] = hi

	return nil
}
