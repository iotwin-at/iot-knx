package dpt

type Dpt1017 B1

// Value implements IDpt.
func (d Dpt1017) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1017) Name() string {
	return "DPT_Trigger"
}

// String implements IDpt.
func (d Dpt1017) String() string {
	if d {
		return "trigger 1"
	}
	return "trigger 0"
}

// ToBytes implements IDpt.
func (d Dpt1017) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1017) Unit() string {
	return ""
}

func UnpackDpt1017(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1017(v), err
}
