package cartridge

import (
	"bytes"
	"errors"
	"os"
)

type Cartridge struct {
	PrgRom      []byte
	NumPrgBanks int
}

var iNesHeader = []byte{'N', 'E', 'S', 0x1A}

func LoadCartridge(filename string) (*Cartridge, error) {
	data, error := os.ReadFile(filename)

	if error != nil {
		return nil, error
	}

	if len(data) < 16 {
		return nil, errors.New("File too small to be a valid iNES ROM.")
	}

	if !bytes.Equal(data[:4], iNesHeader) {
		return nil, errors.New("Invalid iNES header.")
	}

	numPrgBanks := int(data[0x04])
	prgRomSize := numPrgBanks * 16 * 1024

	if len(data) < 16+prgRomSize {
		return nil, errors.New("File too small for the PRG-ROM data.")
	}

	return &Cartridge{
		PrgRom:      data[16 : 16+prgRomSize],
		NumPrgBanks: numPrgBanks,
	}, nil
}
