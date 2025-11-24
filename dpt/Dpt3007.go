package dpt

import "fmt"

type Dpt3007 B1U3

// Value implements IDpt.
func (d Dpt3007) Value() B1U3 {
	return B1U3(d)
}

// Name implements IDpt.
func (d Dpt3007) Name() string {
	return "DPT_Control_Dimming"
}

// String implements IDpt.
func (d Dpt3007) String() string {
	if d.C {
		return fmt.Sprintf("Increase(%d)", d.StepCode)
	}
	return fmt.Sprintf("Decrease(%d)", d.StepCode)
}

// ToBytes implements IDpt.
func (d Dpt3007) Pack() []byte {
	return packB1U3(d.Value())
}

// Unit implements IDpt.
func (d Dpt3007) Unit() string {
	return ""
}

func UnpackDpt3007(data []byte) (IDpt[B1U3], error) {
	v, err := unpackB1U3(data)
	return Dpt3007(v), err
}
