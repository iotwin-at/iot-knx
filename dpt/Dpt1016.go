package dpt

type Dpt1016 B1

// Value implements IDpt.
func (d Dpt1016) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1016) Name() string {
	return "DPT_Ack"
}

// String implements IDpt.
func (d Dpt1016) String() string {
	if d {
		return "acknowledge command (trigger)"
	}
	return "no action (dummy)"
}

// ToBytes implements IDpt.
func (d Dpt1016) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1016) Unit() string {
	return ""
}

func UnpackDpt1016(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1016(v), err
}
