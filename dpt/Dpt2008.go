package dpt

type Dpt2008 B2

// Value implements IDpt.
func (d Dpt2008) Value() B2 {
	return B2(d)
}

// Name implements IDpt.
func (d Dpt2008) Name() string {
	return "DPT_Direction1_Control"
}

// String implements IDpt.
func (d Dpt2008) String() string {
	return B2(d).String()
}

// ToBytes implements IDpt.
func (d Dpt2008) Pack() []byte {
	return packB2(d.Value())
}

// Unit implements IDpt.
func (d Dpt2008) Unit() string {
	return ""
}

func UnpackDpt2008(data []byte) (IDpt[B2], error) {
	v, err := unpackB2(data)
	return Dpt2008(v), err
}
