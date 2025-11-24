package dpt

type Dpt1012 B1

// Value implements IDpt.
func (d Dpt1012) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1012) Name() string {
	return "DPT_Invert"
}

// String implements IDpt.
func (d Dpt1012) String() string {
	if d {
		return "Inverted"
	}
	return "Not inverted"
}

// ToBytes implements IDpt.
func (d Dpt1012) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1012) Unit() string {
	return ""
}

func UnpackDpt1012(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1012(v), err
}
