package dpt

type Dpt7011 U16

// Value implements IDpt.
func (d Dpt7011) Value() U16 {
	return U16(d)
}

// Name implements IDpt.
func (d Dpt7011) Name() string {
	return "DPT_Length_mm"
}

// String implements IDpt.
func (d Dpt7011) String() string {
	return U16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt7011) Pack() []byte {
	return packU16(d.Value())
}

// Unit implements IDpt.
func (d Dpt7011) Unit() string {
	return "mm"
}

func UnpackDpt7011(data []byte) (IDpt[U16], error) {
	v, err := unpackU16(data)
	return Dpt7011(v), err
}
