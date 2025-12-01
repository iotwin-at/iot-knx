package dpt

type Dpt2007 B2

// Value implements IDpt.
func (d Dpt2007) Value() B2 {
	return B2(d)
}

// Name implements IDpt.
func (d Dpt2007) Name() string {
	return "DPT_Step_Control"
}

// String implements IDpt.
func (d Dpt2007) String() string {
	return B2(d).String()
}

// ToBytes implements IDpt.
func (d Dpt2007) Pack() []byte {
	return packB2(d.Value())
}

// Unit implements IDpt.
func (d Dpt2007) Unit() string {
	return ""
}

func UnpackDpt2007(data []byte) (IDpt[B2], error) {
	v, err := unpackB2(data)
	return Dpt2007(v), err
}
