package dpt

import "fmt"

type Dpt5005 U8

// Value implements IDpt.
func (d Dpt5005) Value() U8 {
	return U8(d)
}

// Name implements IDpt.
func (d Dpt5005) Name() string {
	return "DPT_DecimalFactor"
}

// String implements IDpt.
func (d Dpt5005) String() string {
	return fmt.Sprintf("%d", d)
}

// ToBytes implements IDpt.
func (d Dpt5005) Pack() []byte {
	return packU8(d.Value())
}

// Unit implements IDpt.
func (d Dpt5005) Unit() string {
	return "ratio"
}

func UnpackDpt5005(data []byte) (IDpt[U8], error) {
	v, err := unpackU8(data)
	return Dpt5005(v), err
}
