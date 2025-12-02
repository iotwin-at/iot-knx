package knx

import (
	"encoding/binary"
	"fmt"
)

type Int16 int16

func (b Int16) String() string {
	return fmt.Sprintf("%v", int16(b))
}

func (b Int16) Pack() []byte {
	var data []byte
	binary.BigEndian.PutUint16(data, uint16(b))
	return []byte{byte(b)}
}

func NewInt16(data []byte) (Int16, error) {
	if len(data) != 2 {
		return Int16(0), NewErrInvalidDataType("given data cannot be unpacked to Int16 value")
	}
	return Int16(int16(binary.BigEndian.Uint16(data))), nil
}
