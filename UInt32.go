package knx

import (
	"encoding/binary"
	"fmt"
)

type UInt32 uint32

func (b UInt32) String() string {
	return fmt.Sprintf("%v", UInt32(b))
}

func (b UInt32) Pack() []byte {
	var data []byte
	binary.BigEndian.PutUint32(data, uint32(b))
	return data
}

func NewUInt32(data []byte) (UInt32, error) {
	if len(data) != 4 {
		return UInt32(0), NewErrInvalidDataType("given data cannot be unpacked to UInt32 value")
	}
	return UInt32(binary.BigEndian.Uint32(data)), nil
}
