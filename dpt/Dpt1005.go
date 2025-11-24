package dpt

type Dpt1005 B1

// Value implements IDpt.
func (d Dpt1005) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1005) Name() string {
	return "DPT_Alarm"
}

// String implements IDpt.
func (d Dpt1005) String() string {
	if d {
		return "Alarm"
	}
	return "No alarm"
}

// ToBytes implements IDpt.
func (d Dpt1005) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1005) Unit() string {
	return ""
}

func UnpackDpt1005(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1005(v), err
}
