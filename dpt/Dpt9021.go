package dpt

import "fmt"

type Dpt9021 F16

// Value implements IDpt.
func (d Dpt9021) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9021) Name() string {
	return "DPT_Value_Curr"
}

// String implements IDpt.
func (d Dpt9021) String() string {
	return fmt.Sprintf("%v%s", F16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt9021) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9021) Unit() string {
	return "mA"
}

func UnpackDpt9021(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9021(v), err
}
