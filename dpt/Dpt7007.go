package dpt

import "time"

type Dpt7007 time.Duration

// Value implements IDpt.
func (d Dpt7007) Value() time.Duration {
	return time.Duration(d)
}

// Name implements IDpt.
func (d Dpt7007) Name() string {
	return "DPT_TimePeriodHrs"
}

// String implements IDpt.
func (d Dpt7007) String() string {
	return d.Value().String()
}

// ToBytes implements IDpt.
func (d Dpt7007) Pack() []byte {
	return packU16(U16(d.Value().Hours()))
}

// Unit implements IDpt.
func (d Dpt7007) Unit() string {
	return "h"
}

func UnpackDpt7007(data []byte) (IDpt[time.Duration], error) {
	v, err := unpackU16(data)
	return Dpt7007(time.Duration(v) * time.Hour), err
}
