package dpt

import "time"

type Dpt8006 time.Duration

// Value implements IDpt.
func (d Dpt8006) Value() time.Duration {
	return time.Duration(d)
}

// Name implements IDpt.
func (d Dpt8006) Name() string {
	return "DPT_DeltaTimeMin"
}

// String implements IDpt.
func (d Dpt8006) String() string {
	return d.Value().String()
}

// ToBytes implements IDpt.
func (d Dpt8006) Pack() []byte {
	return packV16(V16(d.Value().Minutes()))
}

// Unit implements IDpt.
func (d Dpt8006) Unit() string {
	return "min"
}

func UnpackDpt8006(data []byte) (IDpt[time.Duration], error) {
	v, err := unpackV16(data)
	return Dpt8006(time.Duration(v) * time.Minute), err
}
