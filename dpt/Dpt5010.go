package dpt

import "fmt"

type Dpt5010 U8

// Value implements IDpt.
func (d Dpt5010) Value() U8 {
	return U8(d)
}

// Name implements IDpt.
func (d Dpt5010) Name() string {
	return "DPT_Value_1_Ucount"
}

// String implements IDpt.
func (d Dpt5010) String() string {
	return fmt.Sprintf("%v%s", U8(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt5010) Pack() []byte {
	return packU8(d.Value())
}

// Unit implements IDpt.
func (d Dpt5010) Unit() string {
	return "counter pulses"
}

func UnpackDpt5010(data []byte) (IDpt[U8], error) {
	v, err := unpackU8(data)
	return Dpt5010(v), err
}
