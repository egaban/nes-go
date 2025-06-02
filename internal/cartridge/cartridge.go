package cartridge

import (
	"errors"
	"os"

	"github.com/egaban/nesgo/internal/cartridge/mappers"
)

type Cartridge struct {
	PrgRom []byte
	ChrRom []byte

	numPrgBanks int
	numChrBanks int
	mapper      mappers.Mapper
}

func LoadCartridge(filename string) (*Cartridge, error) {
	data, error := os.ReadFile(filename)

	if error != nil {
		return nil, error
	}

	if len(data) < 16 {
		return nil, errors.New("file too small to be a valid iNES ROM")
	}

	header := NewHeader(data[:16])
	result := &Cartridge{
		numPrgBanks: header.NumPrgBanks,
		numChrBanks: header.NumChrBanks,
		mapper:      mappers.NewMapper(header.MapperId, header.NumPrgBanks, header.NumChrBanks),
	}

	switch header.Format {
	case ines:
		result.parseInes1(data)
	case ines2:
		panic("iNES 2.0 format not supported yet")
	}

	return result, nil
}

// Reads a byte from the cartridge at the specified PRG address. If the address
// is out of bounds (given the mapper), it returns 0 and false.
func (c *Cartridge) TryReadPrgAt(address uint16) (byte, bool) {
	if address, ok := c.mapper.MapPrgRead(address); ok {
		return c.PrgRom[address], true
	}
	return 0, false
}

// Writes a byte to the cartridge at the specified PRG address. If the address
// is out of bounds (given the mapper), it returns false.
func (c *Cartridge) TryWritePrgAt(address uint16, data byte) bool {
	if address, ok := c.mapper.MapPrgWrite(address); ok {
		c.PrgRom[address] = data
		return true
	}
	return false
}

// Reads a byte from the cartridge at the specified CHR address. If the address
// is out of bounds (given the mapper), it returns 0 and false.
func (c *Cartridge) TryReadChrAt(address uint16) (byte, bool) {
	if address, ok := c.mapper.MapChrRead(address); ok {
		return c.ChrRom[address], true
	}
	return 0, false
}

// Writes a byte to the cartridge at the specified CHR address. If the address
// is out of bounds (given the mapper), it returns false.
func (c *Cartridge) TryWriteChrAt(address uint16, data byte) bool {
	if address, ok := c.mapper.MapChrWrite(address); ok {
		c.ChrRom[address] = data
		return true
	}
	return false
}
