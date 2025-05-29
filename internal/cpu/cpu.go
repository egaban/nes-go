package cpu

import (
	"fmt"

	"github.com/egaban/nesgo/internal/bus"
)

type Cpu struct {
	registers      Registers
	bus            *bus.Bus
	cyclesLeft     uint8
	totalCycles    uint64
	fetchedByte    byte
	fetchedAddress uint16
}

func NewCpu(bus *bus.Bus) *Cpu {
	return &Cpu{
		registers:   newRegisters(),
		bus:         bus,
		cyclesLeft:  0,
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
		cpu.registers.status,
		cpu.registers.sp,
		cpu.totalCycles,
	)
}

func (cpu *Cpu) fetchAndExecute() uint8 {
	instruction, err := cpu.fetch()
	if err != nil {
		fmt.Printf("Error fetching instruction: %v\n", err)
		return 0
	}
	pageCross := cpu.loadMemory(&instruction)
	extraCycles := cpu.execute(&instruction)

	numCycles := instruction.cycles + extraCycles
	if pageCross && instruction.pageCrossExtraCycle {
		numCycles += 1
	}

	return numCycles
}
