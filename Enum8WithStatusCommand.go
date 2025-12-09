package knx

import "fmt"

type Enum8WithStatusCommand struct {
	Enum          uint8
	StatusCommand uint8
}

func (e Enum8WithStatusCommand) String() string {
	return fmt.Sprintf("Enum: %d StatusCommand: %x", e.Enum, e.StatusCommand)
}

func (e Enum8WithStatusCommand) Pack() []byte {
	return []byte{
		e.Enum,
		e.StatusCommand,
	}
}

func NewEnum8WithStatusCommand(data []byte) (Enum8WithStatusCommand, error) {
	if len(data) != 2 {
		return Enum8WithStatusCommand{}, NewErrInvalidDataType("given data cannot be unpacked to Enum8WithStatusCommand value")
	}
	return Enum8WithStatusCommand{
		Enum:          data[0],
		StatusCommand: data[1],
	}, nil
}
