package dpt

type Dpt9024 F16

// Value implements IDpt.
func (d Dpt9024) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9024) Name() string {
	return "DPT_Power"
}

// String implements IDpt.
func (d Dpt9024) String() string {
	return F16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt9024) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9024) Unit() string {
	return "kW"
}

func UnpackDpt9024(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9024(v), err
}
