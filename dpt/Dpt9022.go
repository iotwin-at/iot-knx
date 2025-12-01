package dpt

import "fmt"

type Dpt9022 F16

// Value implements IDpt.
func (d Dpt9022) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9022) Name() string {
	return "DPT_PowerDensity"
}

// String implements IDpt.
func (d Dpt9022) String() string {
	return fmt.Sprintf("%v%s", F16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt9022) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9022) Unit() string {
	return "W/m²"
}

func UnpackDpt9022(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9022(v), err
}
