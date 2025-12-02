package knx

import (
	"fmt"
)

type UInt8 uint8

func (b UInt8) String() string {
	return fmt.Sprintf("%v", uint8(b))
}

func (b UInt8) Pack() []byte {
	return []byte{byte(b)}
}

func NewUInt8(data []byte) (UInt8, error) {
	if len(data) != 1 {
		return UInt8(0), NewErrInvalidDataType("given data cannot be unpacked to UInt8 value")
	}
	return UInt8(data[0]), nil
}
