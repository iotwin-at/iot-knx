package dpt

type Dpt1014 B1

// Value implements IDpt.
func (d Dpt1014) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1014) Name() string {
	return "DPT_InputSource"
}

// String implements IDpt.
func (d Dpt1014) String() string {
	if d {
		return "Calculated"
	}
	return "Fixed"
}

// ToBytes implements IDpt.
func (d Dpt1014) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1014) Unit() string {
	return ""
}

func UnpackDpt1014(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1014(v), err
}
