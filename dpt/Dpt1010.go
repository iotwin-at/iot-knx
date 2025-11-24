package dpt

type Dpt1010 B1

// Value implements IDpt.
func (d Dpt1010) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1010) Name() string {
	return "DPT_Start"
}

// String implements IDpt.
func (d Dpt1010) String() string {
	if d {
		return "Start"
	}
	return "Stop"
}

// ToBytes implements IDpt.
func (d Dpt1010) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1010) Unit() string {
	return ""
}

func UnpackDpt1010(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1010(v), err
}
