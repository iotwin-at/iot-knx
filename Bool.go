package knx

import (
	"fmt"
)

type Bool bool

func (b Bool) String() string {
	return fmt.Sprintf("%v", bool(b))
}

func (b Bool) Pack() []byte {
	if b {
		return []byte{0x01}
	}
	return []byte{0x00}
}

func NewBool(data []byte) (Bool, error) {
	if len(data) != 1 {
		return Bool(false), NewErrInvalidDataType("given data cannot be unpacked as Bool value")
	}
	return Bool((data[0] & 0x01) == 1), nil
}
