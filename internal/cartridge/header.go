package cartridge

import "errors"

type Header struct {
	NumPrgBanks int
	NumChrBanks int
	MapperId    byte
	Format      int
}

// Supported iNES formats
const (
	ines int = iota
	ines2
)

func NewHeader(data []byte) *Header {
	format, err := cartridgeFormat(data)
	if err != nil {
		panic(err)
	}

	return &Header{
		NumPrgBanks: int(data[4]),
		NumChrBanks: int(data[5]),
		MapperId:    byte((data[6] >> 4) | (data[7] & 0xF0)),
		Format:      format,
	}
}

func isINes(data []byte) bool {
	return data[0] == 'N' &&
		data[1] == 'E' &&
		data[2] == 'S' &&
		data[3] == 0x1A
}

func cartridgeFormat(data []byte) (int, error) {
	if isINes(data) {
		if data[7]&0x0C == 0x08 {
			return ines2, nil
		}
		return ines, nil
	}

	return -1, errors.New("not a valid iNES file")
}
