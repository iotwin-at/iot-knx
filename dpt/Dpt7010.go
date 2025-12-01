package dpt

type Dpt7010 U16

// Value implements IDpt.
func (d Dpt7010) Value() U16 {
	return U16(d)
}

// Name implements IDpt.
func (d Dpt7010) Name() string {
	return "DPT_PropDataType"
}

// String implements IDpt.
func (d Dpt7010) String() string {
	return U16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt7010) Pack() []byte {
	return packU16(d.Value())
}

// Unit implements IDpt.
func (d Dpt7010) Unit() string {
	return ""
}

func UnpackDpt7010(data []byte) (IDpt[U16], error) {
	v, err := unpackU16(data)
	return Dpt7010(v), err
}
