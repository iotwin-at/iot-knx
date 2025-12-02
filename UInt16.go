package knx

import (
	"encoding/binary"
	"fmt"
)

type UInt16 uint16

func (b UInt16) String() string {
	return fmt.Sprintf("%v", uint16(b))
}

func (b UInt16) Pack() []byte {
	var data []byte
	binary.BigEndian.PutUint16(data, uint16(b))
	return data
}

func NewUInt16(data []byte) (UInt16, error) {
	if len(data) != 2 {
		return UInt16(0), NewErrInvalidDataType("given data cannot be unpacked to UInt16 value")
	}
	return UInt16(binary.BigEndian.Uint16(data)), nil
}
