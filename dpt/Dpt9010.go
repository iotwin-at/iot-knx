package dpt

import "fmt"

type Dpt9010 F16

// Value implements IDpt.
func (d Dpt9010) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9010) Name() string {
	return "DPT_Value_Time1"
}

// String implements IDpt.
func (d Dpt9010) String() string {
	return fmt.Sprintf("%v%s", F16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt9010) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9010) Unit() string {
	return "s"
}

func UnpackDpt9010(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9010(v), err
}
