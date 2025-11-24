package dpt

import "fmt"

type Dpt5003 U8

// Value implements IDpt.
func (d Dpt5003) Value() U8 {
	return U8(d)
}

// Name implements IDpt.
func (d Dpt5003) Name() string {
	return "DPT_Angle"
}

// String implements IDpt.
func (d Dpt5003) String() string {
	return fmt.Sprintf("%d", d)
}

// ToBytes implements IDpt.
func (d Dpt5003) Pack() []byte {
	return packU8(d.Value())
}

// Unit implements IDpt.
func (d Dpt5003) Unit() string {
	return "°"
}

func UnpackDpt5003(data []byte) (IDpt[U8], error) {
	v, err := unpackU8(data)
	return Dpt5003(v), err
}
