package dpt

type Dpt9002 F16

// Value implements IDpt.
func (d Dpt9002) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9002) Name() string {
	return "DPT_Value_Tempd"
}

// String implements IDpt.
func (d Dpt9002) String() string {
	return F16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt9002) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9002) Unit() string {
	return "K"
}

func UnpackDpt9002(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9002(v), err
}
