package knx

import "fmt"

type BoolWithStatusCommand struct {
	Bool          bool
	StatusCommand uint8
}

func (b BoolWithStatusCommand) String() string {
	return fmt.Sprintf("Bool: %v StatusCommand: %x", b.Bool, b.StatusCommand)
}

func (b BoolWithStatusCommand) Pack() []byte {
	result := make([]byte, 2)
	if b.Bool {
		result[0] |= 0b0000_0001
	}
	result[1] |= b.StatusCommand
	return result
}

func NewBoolWithStatusCommand(data []byte) (BoolWithStatusCommand, error) {
	if len(data) != 2 {
		return BoolWithStatusCommand{}, NewErrInvalidDataType("given data cannot be unpacked to BoolWithStatusCommand value")
	}
	return BoolWithStatusCommand{
		Bool:          (data[1] & 0b0000_0001) == 1,
		StatusCommand: data[0],
	}, nil
}
