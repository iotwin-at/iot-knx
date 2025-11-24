package dpt

type Dpt2004 B2

// Value implements IDpt.
func (d Dpt2004) Value() B2 {
	return B2(d)
}

// Name implements IDpt.
func (d Dpt2004) Name() string {
	return "DPT_Ramp_Control"
}

// String implements IDpt.
func (d Dpt2004) String() string {
	if !d.C {
		return "No control"
	}
	if !d.V {
		return "Control. Function value 0"
	}
	return "Control. Function value 1"
}

// ToBytes implements IDpt.
func (d Dpt2004) Pack() []byte {
	return packB2(d.Value())
}

// Unit implements IDpt.
func (d Dpt2004) Unit() string {
	return ""
}

func UnpackDpt2004(data []byte) (IDpt[B2], error) {
	v, err := unpackB2(data)
	return Dpt2004(v), err
}
