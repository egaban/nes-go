package cpu

import "fmt"

type Opcode = int

const (
	adc Opcode = iota
	and
	asl
	bcc
	bcs
	beq
	bit
	bmi
	bne
	bpl
	brk
	bvc
	bvs
	clc
	cld
	cli
	clv
	cmp
	cpx
	cpy
	dec
	dex
	dey
	eor
	inc
	inx
	iny
	jmp
	jsr
	lda
	ldx
	ldy
	lsr
	nop
	ora
	pha
	php
	pla
	plp
	rol
	ror
	rti
	rts
	sbc
	sec
	sed
	sei
	sta
	stx
	sty
	tax
	tay
	tsx
	txa
	txs
	tya
)

func (cpu *Cpu) fetch() (Instruction, error) {
	opcode, err := cpu.bus.ReadByteAt(cpu.registers.pc)
	cpu.registers.pc += 1
	if err != nil {
		return Instruction{}, err
	}

	switch opcode {
	case 0xA2:
		return Instruction{
			opcode:      ldx,
			cycles:      2,
			addressMode: immediate,
		}, nil
	case 0x4C:
		return Instruction{
			opcode:      jmp,
			cycles:      3,
			addressMode: absolute,
		}, nil
	default:
		panic(fmt.Sprintf("TODO: Implement instruction %02X", opcode))
		// return Instruction{}, nil
	}

}
