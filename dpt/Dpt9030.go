package dpt

import "fmt"

type Dpt9030 F16

// Value implements IDpt.
func (d Dpt9030) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9030) Name() string {
	return "DPT_Concentration_µgm3"
}

// String implements IDpt.
func (d Dpt9030) String() string {
	return fmt.Sprintf("%v%s", F16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt9030) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9030) Unit() string {
	return "µg"
}

func UnpackDpt9030(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9030(v), err
}
