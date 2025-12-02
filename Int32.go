package knx

import (
	"encoding/binary"
	"fmt"
)

type Int32 int32

func (b Int32) String() string {
	return fmt.Sprintf("%v", int32(b))
}

func (b Int32) Pack() []byte {
	var data []byte
	binary.BigEndian.PutUint32(data, uint32(b))
	return []byte{byte(b)}
}

func NewInt32(data []byte) (Int32, error) {
	if len(data) != 4 {
		return Int32(0), NewErrInvalidDataType("given data cannot be unpacked to Int32 value")
	}
	return Int32(int32(binary.BigEndian.Uint32(data))), nil
}
