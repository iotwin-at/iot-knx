package dpt

type Dpt1021 B1

// Value implements IDpt.
func (d Dpt1021) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1021) Name() string {
	return "DPT_LogicalFunction"
}

// String implements IDpt.
func (d Dpt1021) String() string {
	if d {
		return "logical function AND"
	}
	return "logical function OR"
}

// ToBytes implements IDpt.
func (d Dpt1021) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1021) Unit() string {
	return ""
}

func UnpackDpt1021(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1021(v), err
}
