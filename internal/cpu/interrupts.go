package cpu

const irqHandlerAddress = 0xFFFE
const nmiHandlerAdress = 0xFFFA
const resetVectorAddress = 0xFFFC

func (cpu *Cpu) Reset() {
	initVector := cpu.bus.ReadWordAt(resetVectorAddress)
	cpu.registers.reset(initVector)
	cpu.status.byteValue = 0x00 | StatusUnused | StatusBreak | StatusInterrupt

	cpu.fetchedByte = 0x00
	cpu.fetchedAddress = 0x0000
	cpu.cyclesLeft = 8
}

func (cpu *Cpu) interrupt(handlerAddress uint16) {
	// Clear B flag
	cpu.status.byteValue &^= StatusBreak

	// Set I, U flags
	cpu.status.byteValue |= StatusInterrupt | StatusUnused

	cpu.stackPushAddress(cpu.registers.pc)
	cpu.stackPushByte(cpu.status.byteValue)

	cpu.registers.pc = handlerAddress
}

func (cpu *Cpu) irq() {
	// Ignore interrupt request
	if cpu.status.get(StatusInterrupt) {
		return
	}

	irqHandler := cpu.bus.ReadWordAt(irqHandlerAddress)
	cpu.interrupt(irqHandler)
}

func (cpu *Cpu) nmi() {
	nmiHandler := cpu.bus.ReadWordAt(nmiHandlerAdress)
	cpu.interrupt(nmiHandler)
	cpu.cyclesLeft = 8
}
