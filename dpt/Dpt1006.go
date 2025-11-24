package dpt

type Dpt1006 B1

// Value implements IDpt.
func (d Dpt1006) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1006) Name() string {
	return "DPT_BinaryValue"
}

// String implements IDpt.
func (d Dpt1006) String() string {
	if d {
		return "High"
	}
	return "Low"
}

// ToBytes implements IDpt.
func (d Dpt1006) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1006) Unit() string {
	return ""
}

func UnpackDpt1006(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1006(v), err
}
