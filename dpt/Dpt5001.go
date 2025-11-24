package dpt

import (
	"fmt"
)

type Dpt5001 U8

// Value implements IDpt.
func (d Dpt5001) Value() U8 {
	return U8(d)
}

// Name implements IDpt.
func (d Dpt5001) Name() string {
	return "DPT_Scaling"
}

// String implements IDpt.
func (d Dpt5001) String() string {
	return fmt.Sprintf("%d", d)
}

// ToBytes implements IDpt.
func (d Dpt5001) Pack() []byte {
	return packU8(d.Value())
}

// Unit implements IDpt.
func (d Dpt5001) Unit() string {
	return "%"
}

func UnpackDpt5001(data []byte) (IDpt[U8], error) {
	v, err := unpackU8(data)
	return Dpt5001(v), err
}
