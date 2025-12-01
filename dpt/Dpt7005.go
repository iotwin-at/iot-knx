package dpt

import "time"

type Dpt7005 time.Duration

// Value implements IDpt.
func (d Dpt7005) Value() time.Duration {
	return time.Duration(d)
}

// Name implements IDpt.
func (d Dpt7005) Name() string {
	return "DPT_TimePeriodSec"
}

// String implements IDpt.
func (d Dpt7005) String() string {
	return d.Value().String()
}

// ToBytes implements IDpt.
func (d Dpt7005) Pack() []byte {
	return packU16(U16(d.Value().Seconds()))
}

// Unit implements IDpt.
func (d Dpt7005) Unit() string {
	return "s"
}

func UnpackDpt7005(data []byte) (IDpt[time.Duration], error) {
	v, err := unpackU16(data)
	return Dpt7005(time.Duration(v) * time.Second), err
}
