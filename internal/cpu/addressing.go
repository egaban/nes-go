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
	baseAddress, err := cpu.bus.ReadWordAt(cpu.registers.pc)
	if err != nil {
		panic(err)
	}

	cpu.fetchedAddress = baseAddress + uint16(cpu.registers.x)
	cpu.fetchedByte, err = cpu.bus.ReadByteAt(cpu.fetchedAddress)

	if err != nil {
		panic(err)
	}

	cpu.registers.pc += 2
	return didPageCross(cpu.fetchedAddress, baseAddress)
}

func (cpu *Cpu) loadAbsY() bool {
	baseAddress, err := cpu.bus.ReadWordAt(cpu.registers.pc)
	if err != nil {
		panic(err)
	}

	cpu.fetchedAddress = baseAddress + uint16(cpu.registers.y)
	cpu.fetchedByte, err = cpu.bus.ReadByteAt(cpu.fetchedAddress)

	if err != nil {
		panic(err)
	}

	cpu.registers.pc += 2
	return didPageCross(cpu.fetchedAddress, baseAddress)
}

func (cpu *Cpu) loadAbsolute() bool {
	address, err := cpu.bus.ReadWordAt(cpu.registers.pc)
	if err != nil {
		panic(err)
	}

	cpu.fetchedAddress = address
	cpu.fetchedByte, err = cpu.bus.ReadByteAt(address)

	if err != nil {
		panic(err)
	}

	cpu.registers.pc += 2
	return false
}

func (cpu *Cpu) loadImmediate() bool {
	fetchedByte, err := cpu.bus.ReadByteAt(cpu.registers.pc)
	if err != nil {
		panic(err)
	}

	cpu.fetchedAddress = cpu.registers.pc
	cpu.fetchedByte = fetchedByte

	cpu.registers.pc += 1
	return false
}

func (cpu *Cpu) loadIndirect() bool {
	address, err := cpu.bus.ReadWordAt(cpu.registers.pc)
	if err != nil {
		panic(err)
	}

	cpu.fetchedAddress, err = cpu.bus.ReadSamePageWord(address)
	if err != nil {
		panic(err)
	}

	cpu.registers.pc += 2
	return false
}

func (cpu *Cpu) loadIndirectX() bool {
	lo, err := cpu.bus.ReadByteAt(cpu.registers.pc)
	if err != nil {
		panic(err)
	}

	// Effective address
	cpu.fetchedAddress, err = cpu.bus.ReadSamePageWord(uint16(lo + cpu.registers.x))
	if err != nil {
		panic(err)
	}

	// Effective content
	cpu.fetchedByte, err = cpu.bus.ReadByteAt(cpu.fetchedAddress)
	if err != nil {
		panic(err)
	}

	cpu.registers.pc += 1
	return false
}

func (cpu *Cpu) loadIndirectY() bool {
	lo, err := cpu.bus.ReadByteAt(cpu.registers.pc)
	if err != nil {
		panic(err)
	}

	// Base to the effective address
	baseAddress, err := cpu.bus.ReadSamePageWord(uint16(lo))
	if err != nil {
		panic(err)
	}

	// Effective address
	cpu.fetchedAddress = baseAddress + uint16(cpu.registers.y)
	cpu.fetchedByte, err = cpu.bus.ReadByteAt(cpu.fetchedAddress)
	if err != nil {
		panic(err)
	}

	cpu.registers.pc += 1
	return didPageCross(baseAddress, cpu.fetchedAddress)
}

func (cpu *Cpu) loadRelative() bool {
	offset, err := cpu.bus.ReadByteAt(cpu.registers.pc)
	if err != nil {
		panic(err)
	}

	cpu.registers.pc += 1
	cpu.fetchedAddress = cpu.registers.pc + uint16(int8(offset))

	return false
}

func (cpu *Cpu) loadZeroPage() bool {
	lo, err := cpu.bus.ReadByteAt(cpu.registers.pc)
	if err != nil {
		panic(err)
	}

	cpu.fetchedAddress = uint16(lo)
	cpu.fetchedByte, err = cpu.bus.ReadByteAt(cpu.fetchedAddress)
	if err != nil {
		panic(err)
	}

	cpu.registers.pc += 1
	return false
}

func (cpu *Cpu) loadZeroPageX() bool {
	lo, err := cpu.bus.ReadByteAt(cpu.registers.pc)
	if err != nil {
		panic(err)
	}

	cpu.fetchedAddress = uint16(lo + cpu.registers.x)
	cpu.fetchedByte, err = cpu.bus.ReadByteAt(cpu.fetchedAddress)
	if err != nil {
		panic(err)
	}

	cpu.registers.pc += 1
	return false

}

func (cpu *Cpu) loadZeroPageY() bool {
	lo, err := cpu.bus.ReadByteAt(cpu.registers.pc)
	if err != nil {
		panic(err)
	}

	cpu.fetchedAddress = uint16(lo + cpu.registers.y)
	cpu.fetchedByte, err = cpu.bus.ReadByteAt(cpu.fetchedAddress)
	if err != nil {
		panic(err)
	}

	cpu.registers.pc += 1
	return false

}

func didPageCross(a, b uint16) bool {
	return (a & 0xFF00) != (b & 0xFF00)
}
