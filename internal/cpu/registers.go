package cpu

type Registers struct {
	a, x, y byte
	sp      byte
	pc      uint16
}

func newRegisters(initVector uint16) Registers {
	return Registers{
		a:  0x00,
		x:  0x00,
		y:  0x00,
		sp: 0xFD,
		pc: initVector,
	}

}

func (registers *Registers) reset(initVector uint16) {
	registers.a, registers.x, registers.y = 0x00, 0x00, 0x00
	registers.pc = initVector
	registers.sp = 0xFD
}
