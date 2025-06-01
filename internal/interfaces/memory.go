package interfaces

type Memory interface {
	ReadByteAt(address uint16) byte
	WriteByteAt(address uint16, data byte)
}

type WordMemory interface {
	ReadWordAt(address uint16) uint16
	WriteWordAt(address uint16, data uint16)
}

type PpuRegisters interface {
	ReadByteAt(address uint16) byte
	WriteByteAt(address uint16, data byte)
}
