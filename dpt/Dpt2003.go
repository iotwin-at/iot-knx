package dpt

type Dpt2003 B2

// Value implements IDpt.
func (d Dpt2003) Value() B2 {
	return B2(d)
}

// Name implements IDpt.
func (d Dpt2003) Name() string {
	return "DPT_Enable_Control"
}

// String implements IDpt.
func (d Dpt2003) String() string {
	return B2(d).String()
}

// ToBytes implements IDpt.
func (d Dpt2003) Pack() []byte {
	return packB2(d.Value())
}

// Unit implements IDpt.
func (d Dpt2003) Unit() string {
	return ""
}

func UnpackDpt2003(data []byte) (IDpt[B2], error) {
	v, err := unpackB2(data)
	return Dpt2003(v), err
}
