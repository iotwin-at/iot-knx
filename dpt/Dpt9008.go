package dpt

type Dpt9008 F16

// Value implements IDpt.
func (d Dpt9008) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9008) Name() string {
	return "DPT_Value_AirQuality"
}

// String implements IDpt.
func (d Dpt9008) String() string {
	return F16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt9008) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9008) Unit() string {
	return "ppm"
}

func UnpackDpt9008(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9008(v), err
}
