package dpt

type Dpt9001 F16

// Value implements IDpt.
func (d Dpt9001) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9001) Name() string {
	return "DPT_Value_Temp"
}

// String implements IDpt.
func (d Dpt9001) String() string {
	return F16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt9001) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9001) Unit() string {
	return "°C"
}

func UnpackDpt9001(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9001(v), err
}
