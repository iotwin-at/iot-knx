package dpt

type Dpt2001 B2

// Value implements IDpt.
func (d Dpt2001) Value() B2 {
	return B2(d)
}

// Name implements IDpt.
func (d Dpt2001) Name() string {
	return "DPT_Switch_Control"
}

// String implements IDpt.
func (d Dpt2001) String() string {
	if !d.C {
		return "No control"
	}
	if !d.V {
		return "Control. Function value 0"
	}
	return "Control. Function value 1"
}

// ToBytes implements IDpt.
func (d Dpt2001) Pack() []byte {
	return packB2(d.Value())
}

// Unit implements IDpt.
func (d Dpt2001) Unit() string {
	return ""
}

func UnpackDpt2001(data []byte) (IDpt[B2], error) {
	v, err := unpackB2(data)
	return Dpt2001(v), err
}
