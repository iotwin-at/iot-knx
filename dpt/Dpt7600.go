package dpt

import "fmt"

type Dpt7600 U16

// Value implements IDpt.
func (d Dpt7600) Value() U16 {
	return U16(d)
}

// Name implements IDpt.
func (d Dpt7600) Name() string {
	return "DPT_Absolute_Colour_Temperature"
}

// String implements IDpt.
func (d Dpt7600) String() string {
	return fmt.Sprintf("%v%s", U16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt7600) Pack() []byte {
	return packU16(d.Value())
}

// Unit implements IDpt.
func (d Dpt7600) Unit() string {
	return "K"
}

func UnpackDpt7600(data []byte) (IDpt[U16], error) {
	v, err := unpackU16(data)
	return Dpt7600(v), err
}
