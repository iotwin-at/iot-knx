package dpt

import "fmt"

type Dpt9004 F16

// Value implements IDpt.
func (d Dpt9004) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9004) Name() string {
	return "DPT_Value_Lux"
}

// String implements IDpt.
func (d Dpt9004) String() string {
	return fmt.Sprintf("%v%s", F16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt9004) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9004) Unit() string {
	return "Lux"
}

func UnpackDpt9004(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9004(v), err
}
