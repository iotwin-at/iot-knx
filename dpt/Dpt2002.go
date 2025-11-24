package dpt

type Dpt2002 B2

// Value implements IDpt.
func (d Dpt2002) Value() B2 {
	return B2(d)
}

// Name implements IDpt.
func (d Dpt2002) Name() string {
	return "DPT_Bool_Control"
}

// String implements IDpt.
func (d Dpt2002) String() string {
	if !d.C {
		return "No control"
	}
	if !d.V {
		return "Control. Function value 0"
	}
	return "Control. Function value 1"
}

// ToBytes implements IDpt.
func (d Dpt2002) Pack() []byte {
	return packB2(d.Value())
}

// Unit implements IDpt.
func (d Dpt2002) Unit() string {
	return ""
}

func UnpackDpt2002(data []byte) (IDpt[B2], error) {
	v, err := unpackB2(data)
	return Dpt2002(v), err
}
