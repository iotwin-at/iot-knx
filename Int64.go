package knx

import (
	"encoding/binary"
	"fmt"
)

type Int64 int64

func (b Int64) String() string {
	return fmt.Sprintf("%v", int64(b))
}

func (b Int64) Pack() []byte {
	var data []byte
	binary.BigEndian.PutUint64(data, uint64(b))
	return []byte{byte(b)}
}

func NewInt64(data []byte) (Int64, error) {
	if len(data) != 4 {
		return Int64(0), NewErrInvalidDataType("given data cannot be unpacked to Int64 value")
	}
	return Int64(int64(binary.BigEndian.Uint64(data))), nil
}
