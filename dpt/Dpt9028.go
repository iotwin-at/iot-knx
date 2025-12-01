package dpt

type Dpt9028 F16

// Value implements IDpt.
func (d Dpt9028) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9028) Name() string {
	return "DPT_Value_Wsp_kmh"
}

// String implements IDpt.
func (d Dpt9028) String() string {
	return F16(d).String()
}

// ToBytes implements IDpt.
func (d Dpt9028) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9028) Unit() string {
	return "km/h"
}

func UnpackDpt9028(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9028(v), err
}
