package dpt

type Dpt4002 A8

// Value implements IDpt.
func (d Dpt4002) Value() A8 {
	return A8(d)
}

// Name implements IDpt.
func (d Dpt4002) Name() string {
	return "DPT_Char_8859_1"
}

// String implements IDpt.
func (d Dpt4002) String() string {
	return string(rune(d))
}

// ToBytes implements IDpt.
func (d Dpt4002) Pack() []byte {
	return packA8(d.Value())
}

// Unit implements IDpt.
func (d Dpt4002) Unit() string {
	return ""
}

func UnpackDpt4002(data []byte) (IDpt[A8], error) {
	v, err := unpackA8(data)
	return Dpt4002(v), err
}
