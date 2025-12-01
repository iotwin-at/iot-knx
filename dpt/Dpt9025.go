package dpt

type Dpt9025 F16

// Value implements IDpt.
func (d Dpt9025) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9025) Name() string {
	return "DPT_Value_Volume_Flow"
}

// String implements IDpt.
func (d Dpt9025) String() string {
	return F16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt9025) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9025) Unit() string {
	return "l/h"
}

func UnpackDpt9025(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9025(v), err
}
