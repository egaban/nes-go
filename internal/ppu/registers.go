package ppu

type Registers struct {
	control uint8
	mask    uint8
	status  uint8
	t       uint16
	v       uint16
	x       byte // Fine X scroll
}

const (
	statusSpriteOverflow = 1 << 5
	statusSpriteZeroHit  = 1 << 6
	statusVerticalBlank  = 1 << 7
)

const (
	maskGrayscale = 1 << iota
	maskRenderBackgroundLeft
	maskRenderSpritesLeft
	maskRenderBackground
	maskRenderSprites
	maskEmphasizeRed
	maskEmphasizeGreen
	maskEmphasizeBlue
)

const (
	controlNametableX = 1 << iota
	controlNametableY
	controlIncrement
	controlSpritePatternTable
	controlBackgroundPatternTable
	controlSpriteSize
	controlOutputMode // Renamed from ControlMasterSlave
	controlGenerateNMI
)
