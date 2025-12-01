package dpt

import "time"

type Dpt8007 time.Duration

// Value implements IDpt.
func (d Dpt8007) Value() time.Duration {
	return time.Duration(d)
}

// Name implements IDpt.
func (d Dpt8007) Name() string {
	return "DPT_DeltaTimeHrs"
}

// String implements IDpt.
func (d Dpt8007) String() string {
	return d.Value().String()
}

// ToBytes implements IDpt.
func (d Dpt8007) Pack() []byte {
	return packV16(V16(d.Value().Hours()))
}

// Unit implements IDpt.
func (d Dpt8007) Unit() string {
	return "h"
}

func UnpackDpt8007(data []byte) (IDpt[time.Duration], error) {
	v, err := unpackV16(data)
	return Dpt8007(time.Duration(v) * time.Hour), err
}
