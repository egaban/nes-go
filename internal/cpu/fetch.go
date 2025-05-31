package cpu

import (
	"errors"
	"fmt"
)

type Opcode = int

var instructionTable [256]Instruction

const (
	bad Opcode = iota
	adc
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

func initInstructionTable() {
	// ADC
	instructionTable[0x69] = Instruction{opcode: adc, addressMode: immediate, cycles: 2}
	instructionTable[0x65] = Instruction{opcode: adc, addressMode: zeroPage, cycles: 3}
	instructionTable[0x75] = Instruction{opcode: adc, addressMode: zeroPageX, cycles: 4}
	instructionTable[0x6D] = Instruction{opcode: adc, addressMode: absolute, cycles: 4}
	instructionTable[0x7D] = Instruction{opcode: adc, addressMode: absoluteX, cycles: 4, pageCrossExtraCycle: true}
	instructionTable[0x79] = Instruction{opcode: adc, addressMode: absoluteY, cycles: 4, pageCrossExtraCycle: true}
	instructionTable[0x61] = Instruction{opcode: adc, addressMode: indirectX, cycles: 6}
	instructionTable[0x71] = Instruction{opcode: adc, addressMode: indirectY, cycles: 5, pageCrossExtraCycle: true}

	// AND
	instructionTable[0x29] = Instruction{opcode: and, addressMode: immediate, cycles: 2}
	instructionTable[0x25] = Instruction{opcode: and, addressMode: zeroPage, cycles: 3}
	instructionTable[0x35] = Instruction{opcode: and, addressMode: zeroPageX, cycles: 4}
	instructionTable[0x2D] = Instruction{opcode: and, addressMode: absolute, cycles: 4}
	instructionTable[0x3D] = Instruction{opcode: and, addressMode: absoluteX, cycles: 4, pageCrossExtraCycle: true}
	instructionTable[0x39] = Instruction{opcode: and, addressMode: absoluteY, cycles: 4, pageCrossExtraCycle: true}
	instructionTable[0x21] = Instruction{opcode: and, addressMode: indirectX, cycles: 6}
	instructionTable[0x31] = Instruction{opcode: and, addressMode: indirectY, cycles: 5, pageCrossExtraCycle: true}

	// ASL
	instructionTable[0x0A] = Instruction{opcode: asl, addressMode: accumulator, cycles: 2}
	instructionTable[0x06] = Instruction{opcode: asl, addressMode: zeroPage, cycles: 5}
	instructionTable[0x16] = Instruction{opcode: asl, addressMode: zeroPageX, cycles: 6}
	instructionTable[0x0E] = Instruction{opcode: asl, addressMode: absolute, cycles: 6}
	instructionTable[0x1E] = Instruction{opcode: asl, addressMode: absoluteX, cycles: 7}

	// BCC
	instructionTable[0x90] = Instruction{opcode: bcc, addressMode: relative, cycles: 2}

	// BCS
	instructionTable[0xB0] = Instruction{opcode: bcs, addressMode: relative, cycles: 2}

	// BEQ
	instructionTable[0xF0] = Instruction{opcode: beq, addressMode: relative, cycles: 2}

	// BIT
	instructionTable[0x24] = Instruction{opcode: bit, addressMode: zeroPage, cycles: 3}
	instructionTable[0x2C] = Instruction{opcode: bit, addressMode: absolute, cycles: 4}

	// BMI
	instructionTable[0x30] = Instruction{opcode: bmi, addressMode: relative, cycles: 2}

	// BNE
	instructionTable[0xD0] = Instruction{opcode: bne, addressMode: relative, cycles: 2}

	// BPL
	instructionTable[0x10] = Instruction{opcode: bpl, addressMode: relative, cycles: 2}

	// BRK
	instructionTable[0x00] = Instruction{opcode: brk, addressMode: implied, cycles: 7}

	// BVC
	instructionTable[0x50] = Instruction{opcode: bvc, addressMode: relative, cycles: 2}

	// BVS
	instructionTable[0x70] = Instruction{opcode: bvs, addressMode: relative, cycles: 2}

	// CLC
	instructionTable[0x18] = Instruction{opcode: clc, addressMode: implied, cycles: 2}

	// CLD
	instructionTable[0xD8] = Instruction{opcode: cld, addressMode: implied, cycles: 2}

	// CLI
	instructionTable[0x58] = Instruction{opcode: cli, addressMode: implied, cycles: 2}

	// CLV
	instructionTable[0xB8] = Instruction{opcode: clv, addressMode: implied, cycles: 2}

	// CMP
	instructionTable[0xC9] = Instruction{opcode: cmp, addressMode: immediate, cycles: 2}
	instructionTable[0xC5] = Instruction{opcode: cmp, addressMode: zeroPage, cycles: 3}
	instructionTable[0xD5] = Instruction{opcode: cmp, addressMode: zeroPageX, cycles: 4}
	instructionTable[0xCD] = Instruction{opcode: cmp, addressMode: absolute, cycles: 4}
	instructionTable[0xDD] = Instruction{opcode: cmp, addressMode: absoluteX, cycles: 4, pageCrossExtraCycle: true}
	instructionTable[0xD9] = Instruction{opcode: cmp, addressMode: absoluteY, cycles: 4, pageCrossExtraCycle: true}
	instructionTable[0xC1] = Instruction{opcode: cmp, addressMode: indirectX, cycles: 6}
	instructionTable[0xD1] = Instruction{opcode: cmp, addressMode: indirectY, cycles: 5, pageCrossExtraCycle: true}

	// CPX
	instructionTable[0xE0] = Instruction{opcode: cpx, addressMode: immediate, cycles: 2}
	instructionTable[0xE4] = Instruction{opcode: cpx, addressMode: zeroPage, cycles: 3}
	instructionTable[0xEC] = Instruction{opcode: cpx, addressMode: absolute, cycles: 4}

	// CPY
	instructionTable[0xC0] = Instruction{opcode: cpy, addressMode: immediate, cycles: 2}
	instructionTable[0xC4] = Instruction{opcode: cpy, addressMode: zeroPage, cycles: 3}
	instructionTable[0xCC] = Instruction{opcode: cpy, addressMode: absolute, cycles: 4}

	// DEC
	instructionTable[0xC6] = Instruction{opcode: dec, addressMode: zeroPage, cycles: 5}
	instructionTable[0xD6] = Instruction{opcode: dec, addressMode: zeroPageX, cycles: 6}
	instructionTable[0xCE] = Instruction{opcode: dec, addressMode: absolute, cycles: 6}
	instructionTable[0xDE] = Instruction{opcode: dec, addressMode: absoluteX, cycles: 7}

	// DEX
	instructionTable[0xCA] = Instruction{opcode: dex, addressMode: implied, cycles: 2}

	// DEY
	instructionTable[0x88] = Instruction{opcode: dey, addressMode: implied, cycles: 2}

	// EOR
	instructionTable[0x49] = Instruction{opcode: eor, addressMode: immediate, cycles: 2}
	instructionTable[0x45] = Instruction{opcode: eor, addressMode: zeroPage, cycles: 3}
	instructionTable[0x55] = Instruction{opcode: eor, addressMode: zeroPageX, cycles: 4}
	instructionTable[0x4D] = Instruction{opcode: eor, addressMode: absolute, cycles: 4}
	instructionTable[0x5D] = Instruction{opcode: eor, addressMode: absoluteX, cycles: 4, pageCrossExtraCycle: true}
	instructionTable[0x59] = Instruction{opcode: eor, addressMode: absoluteY, cycles: 4, pageCrossExtraCycle: true}
	instructionTable[0x41] = Instruction{opcode: eor, addressMode: indirectX, cycles: 6}
	instructionTable[0x51] = Instruction{opcode: eor, addressMode: indirectY, cycles: 5, pageCrossExtraCycle: true}

	// INC
	instructionTable[0xE6] = Instruction{opcode: inc, addressMode: zeroPage, cycles: 5}
	instructionTable[0xF6] = Instruction{opcode: inc, addressMode: zeroPageX, cycles: 6}
	instructionTable[0xEE] = Instruction{opcode: inc, addressMode: absolute, cycles: 6}
	instructionTable[0xFE] = Instruction{opcode: inc, addressMode: absoluteX, cycles: 7}

	// INX
	instructionTable[0xE8] = Instruction{opcode: inx, addressMode: implied, cycles: 2}

	// INY
	instructionTable[0xC8] = Instruction{opcode: iny, addressMode: implied, cycles: 2}

	// JMP
	instructionTable[0x4C] = Instruction{opcode: jmp, cycles: 3, addressMode: absolute}
	instructionTable[0x6C] = Instruction{opcode: jmp, cycles: 5, addressMode: indirect}

	// JSR
	instructionTable[0x20] = Instruction{opcode: jsr, cycles: 6, addressMode: absolute}

	// LDA
	instructionTable[0xA9] = Instruction{opcode: lda, cycles: 2, addressMode: immediate}
	instructionTable[0xA5] = Instruction{opcode: lda, cycles: 3, addressMode: zeroPage}
	instructionTable[0xB5] = Instruction{opcode: lda, cycles: 4, addressMode: zeroPageX}
	instructionTable[0xAD] = Instruction{opcode: lda, cycles: 4, addressMode: absolute}
	instructionTable[0xBD] = Instruction{opcode: lda, cycles: 4, addressMode: absoluteX, pageCrossExtraCycle: true}
	instructionTable[0xB9] = Instruction{opcode: lda, cycles: 4, addressMode: absoluteY, pageCrossExtraCycle: true}
	instructionTable[0xA1] = Instruction{opcode: lda, cycles: 6, addressMode: indirectX}
	instructionTable[0xB1] = Instruction{opcode: lda, cycles: 5, addressMode: indirectY, pageCrossExtraCycle: true}

	// LDX
	instructionTable[0xA2] = Instruction{opcode: ldx, cycles: 2, addressMode: immediate}
	instructionTable[0xA6] = Instruction{opcode: ldx, cycles: 3, addressMode: zeroPage}
	instructionTable[0xB6] = Instruction{opcode: ldx, cycles: 4, addressMode: zeroPageY}
	instructionTable[0xAE] = Instruction{opcode: ldx, cycles: 4, addressMode: absolute}
	instructionTable[0xBE] = Instruction{opcode: ldx, cycles: 4, addressMode: absoluteY, pageCrossExtraCycle: true}

	// LDY
	instructionTable[0xA0] = Instruction{opcode: ldy, addressMode: immediate, cycles: 2}
	instructionTable[0xA4] = Instruction{opcode: ldy, addressMode: zeroPage, cycles: 3}
	instructionTable[0xB4] = Instruction{opcode: ldy, addressMode: zeroPageX, cycles: 4}
	instructionTable[0xAC] = Instruction{opcode: ldy, addressMode: absolute, cycles: 4}
	instructionTable[0xBC] = Instruction{opcode: ldy, addressMode: absoluteX, cycles: 4, pageCrossExtraCycle: true}

	// LSR
	instructionTable[0x4A] = Instruction{opcode: lsr, addressMode: accumulator, cycles: 2}
	instructionTable[0x46] = Instruction{opcode: lsr, addressMode: zeroPage, cycles: 5}
	instructionTable[0x56] = Instruction{opcode: lsr, addressMode: zeroPageX, cycles: 6}
	instructionTable[0x4E] = Instruction{opcode: lsr, addressMode: absolute, cycles: 6}
	instructionTable[0x5E] = Instruction{opcode: lsr, addressMode: absoluteX, cycles: 7}

	// NOP
	instructionTable[0xEA] = Instruction{opcode: nop, cycles: 2, addressMode: implied}

	// ORA
	instructionTable[0x09] = Instruction{opcode: ora, addressMode: immediate, cycles: 2}
	instructionTable[0x05] = Instruction{opcode: ora, addressMode: zeroPage, cycles: 3}
	instructionTable[0x15] = Instruction{opcode: ora, addressMode: zeroPageX, cycles: 4}
	instructionTable[0x0D] = Instruction{opcode: ora, addressMode: absolute, cycles: 4}
	instructionTable[0x1D] = Instruction{opcode: ora, addressMode: absoluteX, cycles: 4, pageCrossExtraCycle: true}
	instructionTable[0x19] = Instruction{opcode: ora, addressMode: absoluteY, cycles: 4, pageCrossExtraCycle: true}
	instructionTable[0x01] = Instruction{opcode: ora, addressMode: indirectX, cycles: 6}
	instructionTable[0x11] = Instruction{opcode: ora, addressMode: indirectY, cycles: 5, pageCrossExtraCycle: true}

	// PHA
	instructionTable[0x48] = Instruction{opcode: pha, addressMode: implied, cycles: 3}

	// PHP
	instructionTable[0x08] = Instruction{opcode: php, addressMode: implied, cycles: 3}

	// PLA
	instructionTable[0x68] = Instruction{opcode: pla, addressMode: implied, cycles: 4}

	// PLP
	instructionTable[0x28] = Instruction{opcode: plp, addressMode: implied, cycles: 4}

	// ROL
	instructionTable[0x2A] = Instruction{opcode: rol, addressMode: accumulator, cycles: 2}
	instructionTable[0x26] = Instruction{opcode: rol, addressMode: zeroPage, cycles: 5}
	instructionTable[0x36] = Instruction{opcode: rol, addressMode: zeroPageX, cycles: 6}
	instructionTable[0x2E] = Instruction{opcode: rol, addressMode: absolute, cycles: 6}
	instructionTable[0x3E] = Instruction{opcode: rol, addressMode: absoluteX, cycles: 7}

	// ROR
	instructionTable[0x6A] = Instruction{opcode: ror, addressMode: accumulator, cycles: 2}
	instructionTable[0x66] = Instruction{opcode: ror, addressMode: zeroPage, cycles: 5}
	instructionTable[0x76] = Instruction{opcode: ror, addressMode: zeroPageX, cycles: 6}
	instructionTable[0x6E] = Instruction{opcode: ror, addressMode: absolute, cycles: 6}
	instructionTable[0x7E] = Instruction{opcode: ror, addressMode: absoluteX, cycles: 7}

	// RTI
	instructionTable[0x40] = Instruction{opcode: rti, cycles: 6, addressMode: implied}

	// RTS
	instructionTable[0x60] = Instruction{opcode: rts, cycles: 6, addressMode: implied}

	// SBC
	instructionTable[0xE9] = Instruction{opcode: sbc, addressMode: immediate, cycles: 2}
	instructionTable[0xE5] = Instruction{opcode: sbc, addressMode: zeroPage, cycles: 3}
	instructionTable[0xF5] = Instruction{opcode: sbc, addressMode: zeroPageX, cycles: 4}
	instructionTable[0xED] = Instruction{opcode: sbc, addressMode: absolute, cycles: 4}
	instructionTable[0xFD] = Instruction{opcode: sbc, addressMode: absoluteX, cycles: 4, pageCrossExtraCycle: true}
	instructionTable[0xF9] = Instruction{opcode: sbc, addressMode: absoluteY, cycles: 4, pageCrossExtraCycle: true}
	instructionTable[0xE1] = Instruction{opcode: sbc, addressMode: indirectX, cycles: 6}
	instructionTable[0xF1] = Instruction{opcode: sbc, addressMode: indirectY, cycles: 5, pageCrossExtraCycle: true}

	// SEC
	instructionTable[0x38] = Instruction{opcode: sec, cycles: 2, addressMode: implied}

	// SED
	instructionTable[0xF8] = Instruction{opcode: sed, cycles: 2, addressMode: implied}

	// SEI
	instructionTable[0x78] = Instruction{opcode: sei, cycles: 2, addressMode: implied}

	// STA
	instructionTable[0x85] = Instruction{opcode: sta, addressMode: zeroPage, cycles: 3}
	instructionTable[0x95] = Instruction{opcode: sta, addressMode: zeroPageX, cycles: 4}
	instructionTable[0x8D] = Instruction{opcode: sta, addressMode: absolute, cycles: 4}
	instructionTable[0x9D] = Instruction{opcode: sta, addressMode: absoluteX, cycles: 5}
	instructionTable[0x99] = Instruction{opcode: sta, addressMode: absoluteY, cycles: 5}
	instructionTable[0x81] = Instruction{opcode: sta, addressMode: indirectX, cycles: 6}
	instructionTable[0x91] = Instruction{opcode: sta, addressMode: indirectY, cycles: 6}

	// STX
	instructionTable[0x86] = Instruction{opcode: stx, addressMode: zeroPage, cycles: 3}
	instructionTable[0x96] = Instruction{opcode: stx, addressMode: zeroPageY, cycles: 4}
	instructionTable[0x8E] = Instruction{opcode: stx, addressMode: absolute, cycles: 4}

	// STY
	instructionTable[0x84] = Instruction{opcode: sty, addressMode: zeroPage, cycles: 3}
	instructionTable[0x94] = Instruction{opcode: sty, addressMode: zeroPageX, cycles: 4}
	instructionTable[0x8C] = Instruction{opcode: sty, addressMode: absolute, cycles: 4}

	// TAX
	instructionTable[0xAA] = Instruction{opcode: tax, cycles: 2, addressMode: implied}

	// TAY
	instructionTable[0xA8] = Instruction{opcode: tay, cycles: 2, addressMode: implied}

	// TSX
	instructionTable[0xBA] = Instruction{opcode: tsx, cycles: 2, addressMode: implied}

	// TXA
	instructionTable[0x8A] = Instruction{opcode: txa, cycles: 2, addressMode: implied}

	// TXS
	instructionTable[0x9A] = Instruction{opcode: txs, cycles: 2, addressMode: implied}

	// TYA
	instructionTable[0x98] = Instruction{opcode: tya, cycles: 2, addressMode: implied}
}

func (cpu *Cpu) fetch() error {
	opcode, err := cpu.bus.ReadByteAt(cpu.registers.pc)
	cpu.registers.pc += 1
	if err != nil {
		return err
	}

	cpu.currentInstruction = &instructionTable[opcode]
	if cpu.currentInstruction.opcode == bad {
		return errors.New(fmt.Sprintf("Invalid opcode %02X", opcode))
	}

	return nil
}
