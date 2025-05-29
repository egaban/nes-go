package cpu

const (
	// Status flags
	StatusCarry     = 1 << iota // C
	StatusZero                  // Z
	StatusInterrupt             // I
	StatusDecimal               // D
	StatusBreak                 // B
	StatusUnused                // U (this is always 1)
	StatusOverflow              // V
	StatusNegative              // N
)

type Registers struct {
	a, x, y byte
	sp      byte
	pc      uint16
	status  byte
}

func newRegisters() Registers {
	return Registers{
		a:  0x00,
		x:  0x00,
		y:  0x00,
		sp: 0xFD,
		// TODO: Actually, it is the init vector
		pc:     0xC000,
		status: 0x24,
	}

}
