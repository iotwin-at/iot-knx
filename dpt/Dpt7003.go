package dpt

import "time"

type Dpt7003 time.Duration

// Value implements IDpt.
func (d Dpt7003) Value() time.Duration {
	return time.Duration(d)
}

// Name implements IDpt.
func (d Dpt7003) Name() string {
	return "DPT_TimePeriod10Msec"
}

// String implements IDpt.
func (d Dpt7003) String() string {
	return d.Value().String()
}

// ToBytes implements IDpt.
func (d Dpt7003) Pack() []byte {
	return packU16(U16(d.Value().Milliseconds() / 10))
}

// Unit implements IDpt.
func (d Dpt7003) Unit() string {
	return "ms"
}

func UnpackDpt7003(data []byte) (IDpt[time.Duration], error) {
	v, err := unpackU16(data)
	return Dpt7003(time.Duration(v) * time.Millisecond * 10), err
}
