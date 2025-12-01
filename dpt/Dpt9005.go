package dpt

import "fmt"

type Dpt9005 F16

// Value implements IDpt.
func (d Dpt9005) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9005) Name() string {
	return "DPT_Value_Wsp"
}

// String implements IDpt.
func (d Dpt9005) String() string {
	return fmt.Sprintf("%v%s", F16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt9005) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9005) Unit() string {
	return "m/s"
}

func UnpackDpt9005(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9005(v), err
}
