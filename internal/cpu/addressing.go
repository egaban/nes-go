package cpu

type addressingMode int

// Addressing Modes Enum
const (
	badAddressing addressingMode = iota
	absolute
	absoluteX
	absoluteY
	accumulator
	immediate
	implied
	indirect
	indirectX
	indirectY
	relative
	zeroPage
	zeroPageX
	zeroPageY
)

// Returns if page cross happened.
func (cpu *Cpu) loadMemory() bool {
	switch cpu.currentInstruction.addressMode {
	case absolute:
		return cpu.loadAbsolute()
	case absoluteX:
		return cpu.loadAbsX()
	case absoluteY:
		return cpu.loadAbsY()
	case accumulator, implied:
		return false
	case immediate:
		return cpu.loadImmediate()
	case indirect:
		return cpu.loadIndirect()
	case indirectX:
		return cpu.loadIndirectX()
	case indirectY:
		return cpu.loadIndirectY()
	case relative:
		return cpu.loadRelative()
	case zeroPage:
		return cpu.loadZeroPage()
	case zeroPageX:
		return cpu.loadZeroPageX()
	case zeroPageY:
		return cpu.loadZeroPageY()
	default:
		panic("Invalid addressing mode")
	}
}

func (cpu *Cpu) writeBack(value uint8) {
	switch cpu.currentInstruction.addressMode {
	case accumulator:
		cpu.registers.a = value
	case absolute, absoluteX, absoluteY, zeroPage, zeroPageX, zeroPageY:
		cpu.bus.WriteByteAt(cpu.fetchedAddress, value)
	default:
		panic("Addressing mode writeback not implemented")
	}
}

func (cpu *Cpu) loadAbsX() bool {
	baseAddress := cpu.bus.ReadWordAt(cpu.registers.pc)

	cpu.fetchedAddress = baseAddress + uint16(cpu.registers.x)
	cpu.fetchedByte = cpu.bus.ReadByteAt(cpu.fetchedAddress)

	cpu.registers.pc += 2
	return didPageCross(cpu.fetchedAddress, baseAddress)
}

func (cpu *Cpu) loadAbsY() bool {
	baseAddress := cpu.bus.ReadWordAt(cpu.registers.pc)

	cpu.fetchedAddress = baseAddress + uint16(cpu.registers.y)
	cpu.fetchedByte = cpu.bus.ReadByteAt(cpu.fetchedAddress)

	cpu.registers.pc += 2
	return didPageCross(cpu.fetchedAddress, baseAddress)
}

func (cpu *Cpu) loadAbsolute() bool {
	address := cpu.bus.ReadWordAt(cpu.registers.pc)

	cpu.fetchedAddress = address
	cpu.fetchedByte = cpu.bus.ReadByteAt(address)

	cpu.registers.pc += 2
	return false
}

func (cpu *Cpu) loadImmediate() bool {
	fetchedByte := cpu.bus.ReadByteAt(cpu.registers.pc)

	cpu.fetchedAddress = cpu.registers.pc
	cpu.fetchedByte = fetchedByte

	cpu.registers.pc += 1
	return false
}

func (cpu *Cpu) loadIndirect() bool {
	address := cpu.bus.ReadWordAt(cpu.registers.pc)

	cpu.fetchedAddress = cpu.bus.ReadSamePageWord(address)

	cpu.registers.pc += 2
	return false
}

func (cpu *Cpu) loadIndirectX() bool {
	lo := cpu.bus.ReadByteAt(cpu.registers.pc)

	// Effective address
	cpu.fetchedAddress = cpu.bus.ReadSamePageWord(uint16(lo + cpu.registers.x))

	// Effective content
	cpu.fetchedByte = cpu.bus.ReadByteAt(cpu.fetchedAddress)

	cpu.registers.pc += 1
	return false
}

func (cpu *Cpu) loadIndirectY() bool {
	lo := cpu.bus.ReadByteAt(cpu.registers.pc)

	// Base to the effective address
	baseAddress := cpu.bus.ReadSamePageWord(uint16(lo))

	// Effective address
	cpu.fetchedAddress = baseAddress + uint16(cpu.registers.y)
	cpu.fetchedByte = cpu.bus.ReadByteAt(cpu.fetchedAddress)

	cpu.registers.pc += 1
	return didPageCross(baseAddress, cpu.fetchedAddress)
}

func (cpu *Cpu) loadRelative() bool {
	offset := cpu.bus.ReadByteAt(cpu.registers.pc)

	cpu.registers.pc += 1
	cpu.fetchedAddress = cpu.registers.pc + uint16(int8(offset))

	return false
}

func (cpu *Cpu) loadZeroPage() bool {
	lo := cpu.bus.ReadByteAt(cpu.registers.pc)

	cpu.fetchedAddress = uint16(lo)
	cpu.fetchedByte = cpu.bus.ReadByteAt(cpu.fetchedAddress)

	cpu.registers.pc += 1
	return false
}

func (cpu *Cpu) loadZeroPageX() bool {
	lo := cpu.bus.ReadByteAt(cpu.registers.pc)

	cpu.fetchedAddress = uint16(lo + cpu.registers.x)
	cpu.fetchedByte = cpu.bus.ReadByteAt(cpu.fetchedAddress)

	cpu.registers.pc += 1
	return false

}

func (cpu *Cpu) loadZeroPageY() bool {
	lo := cpu.bus.ReadByteAt(cpu.registers.pc)

	cpu.fetchedAddress = uint16(lo + cpu.registers.y)
	cpu.fetchedByte = cpu.bus.ReadByteAt(cpu.fetchedAddress)

	cpu.registers.pc += 1
	return false

}

func didPageCross(a, b uint16) bool {
	return (a & 0xFF00) != (b & 0xFF00)
}
