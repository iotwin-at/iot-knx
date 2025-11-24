package dpt

type Dpt1004 B1

// Value implements IDpt.
func (d Dpt1004) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1004) Name() string {
	return "DPT_Ramp"
}

// String implements IDpt.
func (d Dpt1004) String() string {
	if d {
		return "Ramp"
	}
	return "No ramp"
}

// ToBytes implements IDpt.
func (d Dpt1004) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1004) Unit() string {
	return ""
}

func UnpackDpt1004(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1004(v), err
}
