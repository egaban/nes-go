package cpu

import (
	"fmt"

	"github.com/egaban/nesgo/internal/bus"
)

type Cpu struct {
	bus                *bus.CpuBus
	currentInstruction *Instruction
	cyclesLeft         uint8
	fetchedAddress     uint16
	fetchedByte        byte
	registers          Registers
	status             Status
	totalCycles        uint64
}

func NewCpu(bus *bus.CpuBus) *Cpu {
	initInstructionTable()
	initVector := bus.ReadWordAt(resetVectorAddress)
	return &Cpu{
		bus:         bus,
		cyclesLeft:  0,
		registers:   newRegisters(initVector),
		status:      newStatus(),
		totalCycles: 7,
	}
}

func (cpu *Cpu) Tick() {
	if cpu.cyclesLeft == 0 {
		cpu.printCurrentCpuStatus()
		cpu.cyclesLeft = cpu.fetchAndExecute()
	} else {
		cpu.cyclesLeft--
		cpu.totalCycles += 1
	}
}

func (cpu *Cpu) printCurrentCpuStatus() {
	fmt.Printf("%04X %02X %02X %02X %02X %02X %d\n",
		cpu.registers.pc,
		cpu.registers.a,
		cpu.registers.x,
		cpu.registers.y,
		cpu.status.byteValue,
		cpu.registers.sp,
		cpu.totalCycles,
	)
}

func (cpu *Cpu) fetchAndExecute() uint8 {
	err := cpu.fetch()
	if err != nil {
		fmt.Printf("Error fetching instruction: %v\n", err)
		panic("NOT IMPLEMENTED")
		// return 0
	}
	pageCross := cpu.loadMemory()
	extraCycles := cpu.execute()

	numCycles := cpu.currentInstruction.cycles + extraCycles
	if pageCross && cpu.currentInstruction.pageCrossExtraCycle {
		numCycles += 1
	}

	return numCycles
}
