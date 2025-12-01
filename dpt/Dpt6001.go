package dpt

type Dpt6001 V8

// Value implements IDpt.
func (d Dpt6001) Value() V8 {
	return V8(d)
}

// Name implements IDpt.
func (d Dpt6001) Name() string {
	return "DPT_Percent_V8"
}

// String implements IDpt.
func (d Dpt6001) String() string {
	return V8(d).String()
}

// ToBytes implements IDpt.
func (d Dpt6001) Pack() []byte {
	return packV8(d.Value())
}

// Unit implements IDpt.
func (d Dpt6001) Unit() string {
	return "%"
}

func UnpackDpt6001(data []byte) (IDpt[V8], error) {
	v, err := unpackV8(data)
	return Dpt6001(v), err
}
