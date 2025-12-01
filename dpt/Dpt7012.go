package dpt

import "fmt"

type Dpt7012 U16

// Value implements IDpt.
func (d Dpt7012) Value() U16 {
	return U16(d)
}

// Name implements IDpt.
func (d Dpt7012) Name() string {
	return "DPT_UElCurrentmA"
}

// String implements IDpt.
func (d Dpt7012) String() string {
	return fmt.Sprintf("%v%s", U16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt7012) Pack() []byte {
	return packU16(d.Value())
}

// Unit implements IDpt.
func (d Dpt7012) Unit() string {
	return "mA"
}

func UnpackDpt7012(data []byte) (IDpt[U16], error) {
	v, err := unpackU16(data)
	return Dpt7012(v), err
}
