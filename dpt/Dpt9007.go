package dpt

type Dpt9007 F16

// Value implements IDpt.
func (d Dpt9007) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9007) Name() string {
	return "DPT_Value_Humidity"
}

// String implements IDpt.
func (d Dpt9007) String() string {
	return F16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt9007) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9007) Unit() string {
	return "%"
}

func UnpackDpt9007(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9007(v), err
}
