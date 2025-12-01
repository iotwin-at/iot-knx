package dpt

type Dpt5006 U8

// Value implements IDpt.
func (d Dpt5006) Value() U8 {
	return U8(d)
}

// Name implements IDpt.
func (d Dpt5006) Name() string {
	return "DPT_Tariff"
}

// String implements IDpt.
func (d Dpt5006) String() string {
	return U8(d).String()
}

// ToBytes implements IDpt.
func (d Dpt5006) Pack() []byte {
	return packU8(d.Value())
}

// Unit implements IDpt.
func (d Dpt5006) Unit() string {
	return ""
}

func UnpackDpt5006(data []byte) (IDpt[U8], error) {
	v, err := unpackU8(data)
	return Dpt5006(v), err
}
