package dpt

type Dpt9009 F16

// Value implements IDpt.
func (d Dpt9009) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9009) Name() string {
	return "DPT_Value_AirFlow"
}

// String implements IDpt.
func (d Dpt9009) String() string {
	return F16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt9009) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9009) Unit() string {
	return "m³/h"
}

func UnpackDpt9009(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9009(v), err
}
