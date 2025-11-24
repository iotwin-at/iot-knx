package dpt

type Dpt1008 B1

// Value implements IDpt.
func (d Dpt1008) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1008) Name() string {
	return "DPT_UpDown"
}

// String implements IDpt.
func (d Dpt1008) String() string {
	if d {
		return "Down"
	}
	return "Up"
}

// ToBytes implements IDpt.
func (d Dpt1008) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1008) Unit() string {
	return ""
}

func UnpackDpt1008(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1008(v), err
}
