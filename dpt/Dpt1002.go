package dpt

type Dpt1002 B1

// Value implements IDpt.
func (d Dpt1002) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1002) Name() string {
	return "DPT_Bool"
}

// String implements IDpt.
func (d Dpt1002) String() string {
	if d {
		return "True"
	}
	return "False"
}

// ToBytes implements IDpt.
func (d Dpt1002) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1002) Unit() string {
	return ""
}

func UnpackDpt1002(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1002(v), err
}
