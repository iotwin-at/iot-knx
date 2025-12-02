package knx

import (
	"fmt"
)

type Int8 int8

func (b Int8) String() string {
	return fmt.Sprintf("%v", int8(b))
}

func (b Int8) Pack() []byte {
	return []byte{byte(b)}
}

func NewInt8(data []byte) (Int8, error) {
	if len(data) != 1 {
		return Int8(0), NewErrInvalidDataType("given data cannot be unpacked to Int8 value")
	}
	return Int8(data[0]), nil
}
