package dpt

type Dpt1018 B1

// Value implements IDpt.
func (d Dpt1018) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1018) Name() string {
	return "DPT_Occupancy"
}

// String implements IDpt.
func (d Dpt1018) String() string {
	if d {
		return "occupied"
	}
	return "not occupied"
}

// ToBytes implements IDpt.
func (d Dpt1018) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1018) Unit() string {
	return ""
}

func UnpackDpt1018(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1018(v), err
}
