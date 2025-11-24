package dpt

type Dpt1013 B1

// Value implements IDpt.
func (d Dpt1013) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1013) Name() string {
	return "DPT_DimSendStyle"
}

// String implements IDpt.
func (d Dpt1013) String() string {
	if d {
		return "Cyclically"
	}
	return "Start/stop"
}

// ToBytes implements IDpt.
func (d Dpt1013) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1013) Unit() string {
	return ""
}

func UnpackDpt1013(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1013(v), err
}
