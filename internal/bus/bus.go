package bus

import "github.com/egaban/nesgo/internal/cartridge"

type Bus struct {
	memory [64 * 1024]byte
}

func NewBus(cartridge *cartridge.Cartridge) *Bus {
	result := &Bus{
		memory: [64 * 1024]byte{0},
	}

	result.loadCartridge(cartridge)
	return result
}
