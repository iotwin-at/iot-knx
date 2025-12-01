package dpt

import "fmt"

type Dpt9026 F16

// Value implements IDpt.
func (d Dpt9026) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9026) Name() string {
	return "DPT_Rain_Amount"
}

// String implements IDpt.
func (d Dpt9026) String() string {
	return fmt.Sprintf("%v%s", F16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt9026) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9026) Unit() string {
	return "l/m²"
}

func UnpackDpt9026(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9026(v), err
}
