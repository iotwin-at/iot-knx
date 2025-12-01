package dpt

type Dpt7013 U16

// Value implements IDpt.
func (d Dpt7013) Value() U16 {
	return U16(d)
}

// Name implements IDpt.
func (d Dpt7013) Name() string {
	return "DPT_Brightness"
}

// String implements IDpt.
func (d Dpt7013) String() string {
	return U16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt7013) Pack() []byte {
	return packU16(d.Value())
}

// Unit implements IDpt.
func (d Dpt7013) Unit() string {
	return "lux"
}

func UnpackDpt7013(data []byte) (IDpt[U16], error) {
	v, err := unpackU16(data)
	return Dpt7013(v), err
}
