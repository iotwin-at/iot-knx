package dpt

type Dpt2005 B2

// Value implements IDpt.
func (d Dpt2005) Value() B2 {
	return B2(d)
}

// Name implements IDpt.
func (d Dpt2005) Name() string {
	return "DPT_Alarm_Control"
}

// String implements IDpt.
func (d Dpt2005) String() string {
	return B2(d).String()
}

// ToBytes implements IDpt.
func (d Dpt2005) Pack() []byte {
	return packB2(d.Value())
}

// Unit implements IDpt.
func (d Dpt2005) Unit() string {
	return ""
}

func UnpackDpt2005(data []byte) (IDpt[B2], error) {
	v, err := unpackB2(data)
	return Dpt2005(v), err
}
