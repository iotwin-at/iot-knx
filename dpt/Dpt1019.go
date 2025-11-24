package dpt

type Dpt1019 B1

// Value implements IDpt.
func (d Dpt1019) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1019) Name() string {
	return "DPT_Window_Door"
}

// String implements IDpt.
func (d Dpt1019) String() string {
	if d {
		return "open"
	}
	return "closed"
}

// ToBytes implements IDpt.
func (d Dpt1019) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1019) Unit() string {
	return ""
}

func UnpackDpt1019(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1019(v), err
}
