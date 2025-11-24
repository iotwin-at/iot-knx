package dpt

type Dpt1024 B1

// Value implements IDpt.
func (d Dpt1024) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1024) Name() string {
	return "DPT_DayNight"
}

// String implements IDpt.
func (d Dpt1024) String() string {
	if d {
		return "Night"
	}
	return "Day"
}

// ToBytes implements IDpt.
func (d Dpt1024) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1024) Unit() string {
	return ""
}

func UnpackDpt1024(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1024(v), err
}
