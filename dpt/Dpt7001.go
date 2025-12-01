package dpt

import "fmt"

type Dpt7001 U16

// Value implements IDpt.
func (d Dpt7001) Value() U16 {
	return U16(d)
}

// Name implements IDpt.
func (d Dpt7001) Name() string {
	return "DPT_Value_2_Ucount"
}

// String implements IDpt.
func (d Dpt7001) String() string {
	return fmt.Sprintf("%v%s", U16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt7001) Pack() []byte {
	return packU16(d.Value())
}

// Unit implements IDpt.
func (d Dpt7001) Unit() string {
	return "pulses"
}

func UnpackDpt7001(data []byte) (IDpt[U16], error) {
	v, err := unpackU16(data)
	return Dpt7001(v), err
}
