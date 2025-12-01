package dpt

import "fmt"

type Dpt8010 float32

// Value implements IDpt.
func (d Dpt8010) Value() float32 {
	return float32(d)
}

// Name implements IDpt.
func (d Dpt8010) Name() string {
	return "DPT_Percent_V16"
}

// String implements IDpt.
func (d Dpt8010) String() string {
	return fmt.Sprintf("%.2f%%%s", float32(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt8010) Pack() []byte {
	return packV16(V16(d * 100))
}

// Unit implements IDpt.
func (d Dpt8010) Unit() string {
	return "%"
}

func UnpackDpt8010(data []byte) (IDpt[float32], error) {
	v, err := unpackV16(data)
	return Dpt8010(float32(v) / 100), err
}
