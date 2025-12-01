package dpt

import "fmt"

type Dpt9023 F16

// Value implements IDpt.
func (d Dpt9023) Value() F16 {
	return F16(d)
}

// Name implements IDpt.
func (d Dpt9023) Name() string {
	return "DPT_KelvinPerPercent"
}

// String implements IDpt.
func (d Dpt9023) String() string {
	return fmt.Sprintf("%v%s", F16(d), d.Unit())
}

// ToBytes implements IDpt.
func (d Dpt9023) Pack() []byte {
	return packF16(d.Value())
}

// Unit implements IDpt.
func (d Dpt9023) Unit() string {
	return "K/%"
}

func UnpackDpt9023(data []byte) (IDpt[F16], error) {
	v, err := unpackF16(data)
	return Dpt9023(v), err
}
