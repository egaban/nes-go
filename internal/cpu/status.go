package cpu

const (
	// Status flags
	StatusCarry     byte = 1 << iota // C
	StatusZero                       // Z
	StatusInterrupt                  // I
	StatusDecimal                    // D
	StatusBreak                      // B
	StatusUnused                     // U (this is always 1)
	StatusOverflow                   // V
	StatusNegative                   // N
)

type Status struct {
	byteValue byte
}

func newStatus() Status {
	return Status{
		byteValue: 0x24,
	}
}

func (status *Status) get(flag byte) bool {
	return (status.byteValue & flag) != 0
}

func (status *Status) set(flag byte, value bool) {
	if value {
		status.byteValue = status.byteValue | flag
	} else {
		status.byteValue = status.byteValue &^ flag
	}
}

// Updates the zero and negative flags, based on the passed value.
func (cpu *Cpu) updateZnFlags(value byte) {
	cpu.status.set(StatusZero, value == 0)
	cpu.status.set(StatusNegative, (value&0x80) != 0)
}
