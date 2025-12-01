package dpt

import "time"

type Dpt7004 time.Duration

// Value implements IDpt.
func (d Dpt7004) Value() time.Duration {
	return time.Duration(d)
}

// Name implements IDpt.
func (d Dpt7004) Name() string {
	return "DPT_TimePeriod100Msec"
}

// String implements IDpt.
func (d Dpt7004) String() string {
	return d.Value().String()
}

// ToBytes implements IDpt.
func (d Dpt7004) Pack() []byte {
	return packU16(U16(d.Value().Milliseconds() / 100))
}

// Unit implements IDpt.
func (d Dpt7004) Unit() string {
	return "ms"
}

func UnpackDpt7004(data []byte) (IDpt[time.Duration], error) {
	v, err := unpackU16(data)
	return Dpt7004(time.Duration(v) * time.Millisecond * 100), err
}
