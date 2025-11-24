package knx

type Apci uint16

const (
	GroupValueRead     = Apci(0x00)
	GroupValueResponse = Apci(0x01)
	GroupValueWrite    = Apci(0x02)
)
