package dpt

type Dpt1007 B1

// Value implements IDpt.
func (d Dpt1007) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1007) Name() string {
	return "DPT_Step"
}

// String implements IDpt.
func (d Dpt1007) String() string {
	if d {
		return "Increase"
	}
	return "Decrease"
}

// ToBytes implements IDpt.
func (d Dpt1007) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1007) Unit() string {
	return ""
}

func UnpackDpt1007(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1007(v), err
}
