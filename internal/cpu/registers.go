package cpu

type Registers struct {
	a, x, y byte
	sp      byte
	pc      uint16
}

func newRegisters() Registers {
	return Registers{
		a:  0x00,
		x:  0x00,
		y:  0x00,
		sp: 0xFD,
		// TODO: Actually, it is the init vector
		pc: 0xC000,
	}

}
