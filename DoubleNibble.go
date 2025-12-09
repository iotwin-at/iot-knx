package knx

import "fmt"

type DoubleNibble struct {
	Busy uint8 // High
	Nak  uint8 // Low
}

func (d DoubleNibble) String() string {
	return fmt.Sprintf("Busy: %d Nak: %d", d.Busy, d.Nak)
}

func (d DoubleNibble) Pack() []byte {
	return []byte{
		((d.Busy << 4) | d.Nak),
	}
}

func NewDoubleNibble(data []byte) (DoubleNibble, error) {
	if len(data) <= 0 {
		return DoubleNibble{}, NewErrInvalidDataType("given data cannot be unpacked to DoubleNibble value")
	}
	high := uint8((data[0] & 0b1111_0000) >> 4)
	low := (data[0] & 0b0000_1111)
	return DoubleNibble{
		Busy: high,
		Nak:  low,
	}, nil
}
