package dpt

type Dpt6020 B5N3

// Value implements IDpt.
func (d Dpt6020) Value() B5N3 {
	return B5N3(d)
}

// Name implements IDpt.
func (d Dpt6020) Name() string {
	return "DPT_Status_Mode3"
}

// String implements IDpt.
func (d Dpt6020) String() string {
	return B5N3(d).String()
}

// ToBytes implements IDpt.
func (d Dpt6020) Pack() []byte {
	return packB5N3(d.Value())
}

// Unit implements IDpt.
func (d Dpt6020) Unit() string {
	return ""
}

func UnpackDpt6020(data []byte) (IDpt[B5N3], error) {
	v, err := unpackB5N3(data)
	return Dpt6020(v), err
}
