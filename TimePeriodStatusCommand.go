package knx

import (
	"encoding/binary"
	"fmt"
)

type TimePeriodStatusCommand struct {
	TimePeriod    uint16
	StatusCommand uint8
}

func (t TimePeriodStatusCommand) String() string {
	return fmt.Sprintf("TimePeriod: %d StatusCommand: %x", t.TimePeriod, t.StatusCommand)
}

func (t TimePeriodStatusCommand) Pack() []byte {
	result := make([]byte, 3)
	binary.BigEndian.PutUint16(result[0:1], t.TimePeriod)
	result[2] = t.StatusCommand
	return result
}

func NewTimePeriodStatusCommand(data []byte) (TimePeriodStatusCommand, error) {
	if len(data) != 3 {
		return TimePeriodStatusCommand{}, NewErrInvalidDataType("given data cannot be unpacked to TimePeriodStatusCommand value")
	}
	return TimePeriodStatusCommand{
		TimePeriod:    binary.BigEndian.Uint16(data[0:1]),
		StatusCommand: data[2],
	}, nil
}
