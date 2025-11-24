package dpt

type Dpt1011 B1

// Value implements IDpt.
func (d Dpt1011) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1011) Name() string {
	return "DPT_State"
}

// String implements IDpt.
func (d Dpt1011) String() string {
	if d {
		return "Active"
	}
	return "Inactive"
}

// ToBytes implements IDpt.
func (d Dpt1011) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1011) Unit() string {
	return ""
}

func UnpackDpt1011(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1011(v), err
}
