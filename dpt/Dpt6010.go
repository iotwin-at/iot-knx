package dpt

import "fmt"

type Dpt6010 V8

// Value implements IDpt.
func (d Dpt6010) Value() V8 {
	return V8(d)
}

// Name implements IDpt.
func (d Dpt6010) Name() string {
	return "DPT_Value_1_Count"
}

// String implements IDpt.
func (d Dpt6010) String() string {
	return fmt.Sprintf("%v%s", V8(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt6010) Pack() []byte {
	return packV8(d.Value())
}

// Unit implements IDpt.
func (d Dpt6010) Unit() string {
	return "counter pulses"
}

func UnpackDpt6010(data []byte) (IDpt[V8], error) {
	v, err := unpackV8(data)
	return Dpt6010(v), err
}
