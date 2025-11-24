package dpt

type Dpt1003 B1

// Value implements IDpt.
func (d Dpt1003) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1003) Name() string {
	return "DPT_Enable"
}

// String implements IDpt.
func (d Dpt1003) String() string {
	if d {
		return "Enable"
	}
	return "Disable"
}

// ToBytes implements IDpt.
func (d Dpt1003) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1003) Unit() string {
	return ""
}

func UnpackDpt1003(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1003(v), err
}
