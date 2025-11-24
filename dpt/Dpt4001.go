package dpt

type Dpt4001 A8

// Value implements IDpt.
func (d Dpt4001) Value() A8 {
	return A8(d)
}

// Name implements IDpt.
func (d Dpt4001) Name() string {
	return "DPT_Char_ASCII"
}

// String implements IDpt.
func (d Dpt4001) String() string {
	return string(d)
}

// ToBytes implements IDpt.
func (d Dpt4001) Pack() []byte {
	return packA8(d.Value())
}

// Unit implements IDpt.
func (d Dpt4001) Unit() string {
	return ""
}

func UnpackDpt4001(data []byte) (IDpt[A8], error) {
	v, err := unpackA8(data)
	return Dpt4001(v), err
}
