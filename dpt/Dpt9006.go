package dpt

import "fmt"

type Dpt9006 F16

// Value implements IDpt.
func (d Dpt9006) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9006) Name() string {
	return "DPT_Value_Pres"
}

// String implements IDpt.
func (d Dpt9006) String() string {
	return fmt.Sprintf("%v%s", F16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt9006) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9006) Unit() string {
	return "Pa"
}

func UnpackDpt9006(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9006(v), err
}
