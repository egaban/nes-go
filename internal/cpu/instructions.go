package cpu

type Instruction struct {
	opcode              Opcode
	cycles              uint8
	addressMode         addressingMode
	pageCrossExtraCycle bool
}

// Returns the amount of extra cycles.
func (cpu *Cpu) execute(instruction *Instruction) uint8 {
	switch instruction.opcode {
	case jmp:
		return cpu.jmp()
	default:
		panic("NOT IMPLEMENTED YET")
	}
}

func (cpu *Cpu) jmp() uint8 {
	cpu.registers.pc = cpu.fetchedAddress
	return 0
}
