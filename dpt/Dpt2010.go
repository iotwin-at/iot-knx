package dpt

type Dpt2010 B2

// Value implements IDpt.
func (d Dpt2010) Value() B2 {
	return B2(d)
}

// Name implements IDpt.
func (d Dpt2010) Name() string {
	return "DPT_Start_Control"
}

// String implements IDpt.
func (d Dpt2010) String() string {
	return B2(d).String()
}

// ToBytes implements IDpt.
func (d Dpt2010) Pack() []byte {
	return packB2(d.Value())
}

// Unit implements IDpt.
func (d Dpt2010) Unit() string {
	return ""
}

func UnpackDpt2010(data []byte) (IDpt[B2], error) {
	v, err := unpackB2(data)
	return Dpt2010(v), err
}
