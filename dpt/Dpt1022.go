package dpt

type Dpt1022 B1

// Value implements IDpt.
func (d Dpt1022) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1022) Name() string {
	return "DPT_Scene_AB"
}

// String implements IDpt.
func (d Dpt1022) String() string {
	if d {
		return "scene B"
	}
	return "scene A"
}

// ToBytes implements IDpt.
func (d Dpt1022) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1022) Unit() string {
	return ""
}

func UnpackDpt1022(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1022(v), err
}
