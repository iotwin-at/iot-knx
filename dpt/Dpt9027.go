package dpt

import "fmt"

type Dpt9027 F16

// Value implements IDpt.
func (d Dpt9027) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9027) Name() string {
	return "DPT_Value_Temp_F"
}

// String implements IDpt.
func (d Dpt9027) String() string {
	return fmt.Sprintf("%v%s", F16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt9027) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9027) Unit() string {
	return "°F"
}

func UnpackDpt9027(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9027(v), err
}
