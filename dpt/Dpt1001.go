package dpt

type Dpt1001 B1

// Value implements IDpt.
func (d Dpt1001) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1001) Name() string {
	return "DPT_Switch"
}

// String implements IDpt.
func (d Dpt1001) String() string {
	if d {
		return "On"
	}
	return "Off"
}

// ToBytes implements IDpt.
func (d Dpt1001) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1001) Unit() string {
	return ""
}

func UnpackDpt1001(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1001(v), err
}
