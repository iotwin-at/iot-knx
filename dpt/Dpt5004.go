package dpt

import "fmt"

type Dpt5004 U8

// Value implements IDpt.
func (d Dpt5004) Value() U8 {
	return U8(d)
}

// Name implements IDpt.
func (d Dpt5004) Name() string {
	return "DPT_Percent_U8"
}

// String implements IDpt.
func (d Dpt5004) String() string {
	return fmt.Sprintf("%d", d)
}

// ToBytes implements IDpt.
func (d Dpt5004) Pack() []byte {
	return packU8(d.Value())
}

// Unit implements IDpt.
func (d Dpt5004) Unit() string {
	return "%"
}

func UnpackDpt5004(data []byte) (IDpt[U8], error) {
	v, err := unpackU8(data)
	return Dpt5004(v), err
}
