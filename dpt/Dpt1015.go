package dpt

type Dpt1015 B1

// Value implements IDpt.
func (d Dpt1015) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1015) Name() string {
	return "DPT_Reset"
}

// String implements IDpt.
func (d Dpt1015) String() string {
	if d {
		return "reset command (trigger)"
	}
	return "no action (dummy)"
}

// ToBytes implements IDpt.
func (d Dpt1015) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1015) Unit() string {
	return ""
}

func UnpackDpt1015(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1015(v), err
}
