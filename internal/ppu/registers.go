package ppu

type Registers struct {
	Control uint8 // $2000
	Mask    uint8 // $2001
	Status  uint8 // $2002
	Scroll  uint8 // $2005
	Address uint8 // $2006
	Data    uint8 // $2007
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
