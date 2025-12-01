package dpt

type Dpt8012 V16

// Value implements IDpt.
func (d Dpt8012) Value() V16 {
	return V16(d)
}

// Name implements IDpt.
func (d Dpt8012) Name() string {
	return "DPT_Length_m"
}

// String implements IDpt.
func (d Dpt8012) String() string {
	return V16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt8012) Pack() []byte {
	return packV16(d.Value())
}

// Unit implements IDpt.
func (d Dpt8012) Unit() string {
	return "m"
}

func UnpackDpt8012(data []byte) (IDpt[V16], error) {
	v, err := unpackV16(data)
	return Dpt8012(v), err
}
