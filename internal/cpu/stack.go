package cpu

func (cpu *Cpu) stackPushByte(data byte) {
	stack_address := 0x0100 + uint16(cpu.registers.sp)
	// Stack address - 1, because the stack pointer decreases, and we are storing a word.
	cpu.bus.WriteByteAt(stack_address, data)
	cpu.registers.sp -= 1
}

func (cpu *Cpu) stackPushAddress(address uint16) {
	stack_address := 0x0100 + uint16(cpu.registers.sp)
	// Stack address - 1, because the stack pointer decreases, and we are storing a word.
	cpu.bus.WriteWordAt(stack_address-1, address)
	cpu.registers.sp -= 2
}

func (cpu *Cpu) stackPopByte() (byte, error) {
	stack_address := 0x0100 + uint16(cpu.registers.sp)
	result, err := cpu.bus.ReadByteAt(stack_address + 1)

	if err != nil {
		return 0, err
	}

	cpu.registers.sp += 1
	return result, nil
}

func (cpu *Cpu) stackPopAddress() (uint16, error) {
	stack_address := 0x0100 + uint16(cpu.registers.sp)
	result, err := cpu.bus.ReadWordAt(stack_address + 1)

	if err != nil {
		return 0, err
	}

	cpu.registers.sp += 2
	return result, nil
}
