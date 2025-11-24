package dpt

type Dpt1009 B1

// Value implements IDpt.
func (d Dpt1009) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1009) Name() string {
	return "DPT_OpenClose"
}

// String implements IDpt.
func (d Dpt1009) String() string {
	if d {
		return "Close"
	}
	return "Open"
}

// ToBytes implements IDpt.
func (d Dpt1009) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1009) Unit() string {
	return ""
}

func UnpackDpt1009(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1009(v), err
}
