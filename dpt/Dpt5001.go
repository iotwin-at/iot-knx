package dpt

import "fmt"

type Dpt5001 float32

// Value implements IDpt.
func (d Dpt5001) Value() float32 {
	return float32(d)
}

// Name implements IDpt.
func (d Dpt5001) Name() string {
	return "DPT_Scaling"
}

// String implements IDpt.
func (d Dpt5001) String() string {
	return fmt.Sprintf("%.2f%%", float32(d))
}

// ToBytes implements IDpt.
func (d Dpt5001) Pack() []byte {
	if d <= 0 {
		return packU8(0)
	} else if d >= 100 {
		return packU8(255)
	} else {
		return packU8(U8(d * 255.0 / 100.0))
	}
}

// Unit implements IDpt.
func (d Dpt5001) Unit() string {
	return "%"
}

func UnpackDpt5001(data []byte) (IDpt[float32], error) {
	v, err := unpackU8(data)
	return Dpt5001(100.0 / 255.0 * float32(v)), err
}
