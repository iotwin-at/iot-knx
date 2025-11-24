package dpt

import "fmt"

type Dpt3008 B1U3

// Value implements IDpt.
func (d Dpt3008) Value() B1U3 {
	return B1U3(d)
}

// Name implements IDpt.
func (d Dpt3008) Name() string {
	return "DPT_Control_Blinds"
}

// String implements IDpt.
func (d Dpt3008) String() string {
	if d.C {
		return fmt.Sprintf("Down(%d)", d.StepCode)
	}
	return fmt.Sprintf("Up(%d)", d.StepCode)
}

// ToBytes implements IDpt.
func (d Dpt3008) Pack() []byte {
	return packB1U3(d.Value())
}

// Unit implements IDpt.
func (d Dpt3008) Unit() string {
	return ""
}

func UnpackDpt3008(data []byte) (IDpt[B1U3], error) {
	v, err := unpackB1U3(data)
	return Dpt3008(v), err
}
