package cpu

type Instruction struct {
	opcode              Opcode
	cycles              uint8
	addressMode         addressingMode
	pageCrossExtraCycle bool
}

// Returns the amount of extra cycles.
func (cpu *Cpu) execute() uint8 {
	switch cpu.currentInstruction.opcode {
	case adc:
		cpu.adc()
	case and:
		cpu.and()
	case asl:
		cpu.asl()
	case bcc:
		return cpu.bcc()
	case bcs:
		return cpu.bcs()
	case beq:
		return cpu.beq()
	case bit:
		cpu.bit()
	case bmi:
		return cpu.bmi()
	case bne:
		return cpu.bne()
	case bpl:
		return cpu.bpl()
	case brk:
		panic("brk not implemented yet.")
	case bvc:
		return cpu.bvc()
	case bvs:
		return cpu.bvs()
	case clc:
		cpu.clc()
	case cld:
		cpu.cld()
	case cli:
		cpu.cli()
	case clv:
		cpu.clv()
	case cmp:
		cpu.cmp()
	case cpx:
		cpu.cpx()
	case cpy:
		cpu.cpy()
	case dec:
		cpu.dec()
	case dex:
		cpu.dex()
	case dey:
		cpu.dey()
	case eor:
		cpu.eor()
	case inc:
		cpu.inc()
	case inx:
		cpu.inx()
	case iny:
		cpu.iny()
	case jmp:
		cpu.jmp()
	case jsr:
		cpu.jsr()
	case lda:
		cpu.lda()
	case ldx:
		cpu.ldx()
	case ldy:
		cpu.ldy()
	case lsr:
		cpu.lsr()
	case nop:
		break
	case ora:
		cpu.ora()
	case pha:
		cpu.pha()
	case php:
		cpu.php()
	case pla:
		cpu.pla()
	case plp:
		cpu.plp()
	case rol:
		cpu.rol()
	case ror:
		cpu.ror()
	case rti:
		cpu.rti()
	case rts:
		cpu.rts()
	case sbc:
		cpu.sbc()
	case sec:
		cpu.sec()
	case sed:
		cpu.sed()
	case sei:
		cpu.sei()
	case sta:
		cpu.sta()
	case stx:
		cpu.stx()
	case sty:
		cpu.sty()
	case tax:
		cpu.tax()
	case tay:
		cpu.tay()
	case tsx:
		cpu.tsx()
	case txa:
		cpu.txa()
	case txs:
		cpu.txs()
	case tya:
		cpu.tya()

	default:
		panic("NOT IMPLEMENTED YET")
	}

	return 0
}

func (cpu *Cpu) adc() {
	a := int16(cpu.registers.a)
	operand := int16(cpu.fetchedByte)
	carry := int16(boolToBit(cpu.status.get(StatusCarry)))
	result := a + operand + carry
	overflow := ((result ^ a) & (result ^ operand) & (0x80)) != 0

	cpu.status.set(StatusCarry, result > 0xFF)
	cpu.status.set(StatusOverflow, overflow)
	cpu.updateZnFlags(uint8(result))

	cpu.registers.a = uint8(result)
}

func (cpu *Cpu) and() {
	cpu.registers.a = cpu.registers.a & cpu.fetchedByte
	cpu.updateZnFlags(cpu.registers.a)
}

func (cpu *Cpu) asl() {
	carry := (cpu.fetchedByte & 0x80) != 0
	result := cpu.fetchedByte << 1

	cpu.updateZnFlags(result)
	cpu.status.set(StatusCarry, carry)

	cpu.writeBack(result)
}

func (cpu *Cpu) bcc() uint8 {
	return cpu.branchIf(!cpu.status.get(StatusCarry))
}

func (cpu *Cpu) bcs() uint8 {
	return cpu.branchIf(cpu.status.get(StatusCarry))
}

func (cpu *Cpu) beq() uint8 {
	return cpu.branchIf(cpu.status.get(StatusZero))
}

func (cpu *Cpu) bit() {
	negative := (0x80 & cpu.fetchedByte) != 0
	overflow := (0x40 & cpu.fetchedByte) != 0
	result := cpu.fetchedByte & cpu.registers.a

	cpu.status.set(StatusNegative, negative)
	cpu.status.set(StatusOverflow, overflow)
	cpu.status.set(StatusZero, result == 0)
}

func (cpu *Cpu) bmi() uint8 {
	return cpu.branchIf(cpu.status.get(StatusNegative))
}

func (cpu *Cpu) bne() uint8 {
	return cpu.branchIf(!cpu.status.get(StatusZero))
}

func (cpu *Cpu) bpl() uint8 {
	return cpu.branchIf(!cpu.status.get(StatusNegative))
}

func (cpu *Cpu) bvc() uint8 {
	return cpu.branchIf(!cpu.status.get(StatusOverflow))
}

func (cpu *Cpu) bvs() uint8 {
	return cpu.branchIf(cpu.status.get(StatusOverflow))
}

func (cpu *Cpu) clc() {
	cpu.status.set(StatusCarry, false)
}

func (cpu *Cpu) cld() {
	cpu.status.set(StatusDecimal, false)
}

func (cpu *Cpu) cli() {
	cpu.status.set(StatusInterrupt, false)
}

func (cpu *Cpu) clv() {
	cpu.status.set(StatusOverflow, false)
}

func (cpu *Cpu) cmp() {
	result := cpu.registers.a - cpu.fetchedByte
	cpu.status.set(StatusCarry, cpu.registers.a >= cpu.fetchedByte)
	cpu.updateZnFlags(result)
}

func (cpu *Cpu) cpx() {
	result := cpu.registers.x - cpu.fetchedByte
	cpu.status.set(StatusCarry, cpu.registers.x >= cpu.fetchedByte)
	cpu.updateZnFlags(result)
}

func (cpu *Cpu) cpy() {
	result := cpu.registers.y - cpu.fetchedByte
	cpu.status.set(StatusCarry, cpu.registers.y >= cpu.fetchedByte)
	cpu.updateZnFlags(result)
}

func (cpu *Cpu) dec() {
	result := cpu.fetchedByte - 1
	cpu.updateZnFlags(result)
	cpu.writeBack(result)
}

func (cpu *Cpu) dex() {
	cpu.registers.x = cpu.registers.x - 1
	cpu.updateZnFlags(cpu.registers.x)
}

func (cpu *Cpu) dey() {
	cpu.registers.y = cpu.registers.y - 1
	cpu.updateZnFlags(cpu.registers.y)
}

func (cpu *Cpu) eor() {
	cpu.registers.a = cpu.registers.a ^ cpu.fetchedByte
	cpu.updateZnFlags(cpu.registers.a)
}

func (cpu *Cpu) inc() {
	result := cpu.fetchedByte + 1
	cpu.updateZnFlags(result)
	cpu.writeBack(result)
}

func (cpu *Cpu) inx() {
	cpu.registers.x = cpu.registers.x + 1
	cpu.updateZnFlags(cpu.registers.x)
}

func (cpu *Cpu) iny() {
	cpu.registers.y = cpu.registers.y + 1
	cpu.updateZnFlags(cpu.registers.y)
}

func (cpu *Cpu) jmp() {
	cpu.registers.pc = cpu.fetchedAddress
}

func (cpu *Cpu) jsr() {
	returnAddress := cpu.registers.pc - 1
	cpu.stackPushAddress(returnAddress)
	cpu.registers.pc = cpu.fetchedAddress
}

func (cpu *Cpu) lda() {
	cpu.registers.a = cpu.fetchedByte
	cpu.updateZnFlags(cpu.fetchedByte)
}

func (cpu *Cpu) ldx() {
	cpu.registers.x = cpu.fetchedByte
	cpu.updateZnFlags(cpu.fetchedByte)
}

func (cpu *Cpu) ldy() {
	cpu.registers.y = cpu.fetchedByte
	cpu.updateZnFlags(cpu.fetchedByte)
}

func (cpu *Cpu) lsr() {
	carry := (cpu.fetchedByte & 0x01) != 0
	result := cpu.fetchedByte >> 1

	cpu.updateZnFlags(result)
	cpu.status.set(StatusCarry, carry)

	cpu.writeBack(result)
}

func (cpu *Cpu) ora() {
	cpu.registers.a = cpu.registers.a | cpu.fetchedByte
	cpu.updateZnFlags(cpu.registers.a)
}

func (cpu *Cpu) pha() {
	cpu.stackPushByte(cpu.registers.a)
}

func (cpu *Cpu) php() {
	toPush := cpu.status
	toPush.set(StatusUnused, true)
	toPush.set(StatusBreak, true)
	cpu.stackPushByte(toPush.byteValue)
}

func (cpu *Cpu) pla() {
	result, err := cpu.stackPopByte()
	if err != nil {
		panic(err)
	}
	cpu.registers.a = result
	cpu.updateZnFlags(result)
}

func (cpu *Cpu) plp() {
	oldB := cpu.status.get(StatusBreak)
	newStatus, err := cpu.stackPopByte()
	if err != nil {
		panic(err)
	}
	cpu.status.byteValue = newStatus | StatusUnused
	cpu.status.set(StatusBreak, oldB)
}

func (cpu *Cpu) rol() {
	oldCarry := boolToBit(cpu.status.get(StatusCarry))
	newCarry := (cpu.fetchedByte & 0x80) != 0
	result := (cpu.fetchedByte << 1) | oldCarry

	cpu.updateZnFlags(result)
	cpu.status.set(StatusCarry, newCarry)
	cpu.writeBack(result)
}

func (cpu *Cpu) ror() {
	oldCarry := boolToBit(cpu.status.get(StatusCarry))
	newCarry := (cpu.fetchedByte & 0x01) != 0
	result := (cpu.fetchedByte >> 1) | (oldCarry << 7)

	cpu.updateZnFlags(result)
	cpu.status.set(StatusCarry, newCarry)
	cpu.writeBack(result)
}

func (cpu *Cpu) rti() {
	popped, err := cpu.stackPopByte()
	if err != nil {
		panic(err)
	}

	breakAndUnused := StatusUnused | (popped & StatusBreak)

	cpu.status.byteValue = (popped &^ StatusBreak) | breakAndUnused
	cpu.registers.pc, err = cpu.stackPopAddress()
	if err != nil {
		panic(err)
	}
}

func (cpu *Cpu) rts() {
	address, err := cpu.stackPopAddress()
	if err != nil {
		panic(err)
	}
	cpu.registers.pc = address + 1
}

func (cpu *Cpu) sbc() {
	a := cpu.registers.a
	operand := cpu.fetchedByte
	oldCarry := boolToBit(cpu.status.get(StatusCarry))

	result := uint16(a) + uint16(^operand) + uint16(oldCarry)
	carry := result > 0x00FF

	resultByte := byte(result)
	cpu.updateZnFlags(resultByte)
	cpu.status.set(StatusCarry, carry)
	cpu.status.set(StatusOverflow, ((a^resultByte)&(a^operand)&0x80) != 0)
	cpu.registers.a = resultByte
}

func (cpu *Cpu) sec() {
	cpu.status.set(StatusCarry, true)
}

func (cpu *Cpu) sed() {
	cpu.status.set(StatusDecimal, true)
}

func (cpu *Cpu) sei() {
	cpu.status.set(StatusInterrupt, true)
}

func (cpu *Cpu) sta() {
	address := cpu.fetchedAddress
	cpu.bus.WriteByteAt(address, cpu.registers.a)
}

func (cpu *Cpu) stx() {
	address := cpu.fetchedAddress
	cpu.bus.WriteByteAt(address, cpu.registers.x)
}

func (cpu *Cpu) sty() {
	address := cpu.fetchedAddress
	cpu.bus.WriteByteAt(address, cpu.registers.y)
}

func (cpu *Cpu) tax() {
	cpu.registers.x = cpu.registers.a
	cpu.updateZnFlags(cpu.registers.x)
}

func (cpu *Cpu) tay() {
	cpu.registers.y = cpu.registers.a
	cpu.updateZnFlags(cpu.registers.y)
}

func (cpu *Cpu) tsx() {
	cpu.registers.x = cpu.registers.sp
	cpu.updateZnFlags(cpu.registers.x)
}

func (cpu *Cpu) txa() {
	cpu.registers.a = cpu.registers.x
	cpu.updateZnFlags(cpu.registers.a)
}

func (cpu *Cpu) txs() {
	cpu.registers.sp = cpu.registers.x
}

func (cpu *Cpu) tya() {
	cpu.registers.a = cpu.registers.y
	cpu.updateZnFlags(cpu.registers.a)
}

func (cpu *Cpu) branchIf(condition bool) uint8 {
	if condition {
		pageCross := didPageCross(cpu.registers.pc, cpu.fetchedAddress)
		cpu.registers.pc = cpu.fetchedAddress

		if pageCross {
			return 2
		} else {
			return 1
		}
	}

	return 0
}

// true => 1, false => 0
func boolToBit(value bool) byte {
	if value {
		return 1
	}
	return 0
}
