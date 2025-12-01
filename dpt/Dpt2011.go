package dpt

type Dpt2011 B2

// Value implements IDpt.
func (d Dpt2011) Value() B2 {
	return B2(d)
}

// Name implements IDpt.
func (d Dpt2011) Name() string {
	return "DPT_State_Control"
}

// String implements IDpt.
func (d Dpt2011) String() string {
	return B2(d).String()
}

// ToBytes implements IDpt.
func (d Dpt2011) Pack() []byte {
	return packB2(d.Value())
}

// Unit implements IDpt.
func (d Dpt2011) Unit() string {
	return ""
}

func UnpackDpt2011(data []byte) (IDpt[B2], error) {
	v, err := unpackB2(data)
	return Dpt2011(v), err
}
