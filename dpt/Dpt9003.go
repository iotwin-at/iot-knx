package dpt

type Dpt9003 F16

// Value implements IDpt.
func (d Dpt9003) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9003) Name() string {
	return "DPT_Value_Tempa"
}

// String implements IDpt.
func (d Dpt9003) String() string {
	return F16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt9003) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9003) Unit() string {
	return "K/h"
}

func UnpackDpt9003(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9003(v), err
}
