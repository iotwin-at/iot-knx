package dpt

type Dpt8001 V16

// Value implements IDpt.
func (d Dpt8001) Value() V16 {
	return V16(d)
}

// Name implements IDpt.
func (d Dpt8001) Name() string {
	return "DPT_Value_2_Count"
}

// String implements IDpt.
func (d Dpt8001) String() string {
	return V16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt8001) Pack() []byte {
	return packV16(d.Value())
}

// Unit implements IDpt.
func (d Dpt8001) Unit() string {
	return "pulses"
}

func UnpackDpt8001(data []byte) (IDpt[V16], error) {
	v, err := unpackV16(data)
	return Dpt8001(v), err
}
