package dpt

import "fmt"

type Dpt9029 F16

// Value implements IDpt.
func (d Dpt9029) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9029) Name() string {
	return "DPT_Value_Absolute_Humidity"
}

// String implements IDpt.
func (d Dpt9029) String() string {
	return fmt.Sprintf("%v%s", F16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt9029) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9029) Unit() string {
	return "gm³"
}

func UnpackDpt9029(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9029(v), err
}
