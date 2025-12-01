package dpt

type Dpt9011 F16

// Value implements IDpt.
func (d Dpt9011) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9011) Name() string {
	return "DPT_Value_Time2"
}

// String implements IDpt.
func (d Dpt9011) String() string {
	return F16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt9011) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9011) Unit() string {
	return "ms"
}

func UnpackDpt9011(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9011(v), err
}
