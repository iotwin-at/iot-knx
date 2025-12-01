package dpt

import "fmt"

type Dpt9020 F16

// Value implements IDpt.
func (d Dpt9020) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9020) Name() string {
	return "DPT_Value_Volt"
}

// String implements IDpt.
func (d Dpt9020) String() string {
	return fmt.Sprintf("%v%s", F16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt9020) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9020) Unit() string {
	return "mV"
}

func UnpackDpt9020(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9020(v), err
}
