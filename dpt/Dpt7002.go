package dpt

import "time"

type Dpt7002 time.Duration

// Value implements IDpt.
func (d Dpt7002) Value() time.Duration {
	return time.Duration(d)
}

// Name implements IDpt.
func (d Dpt7002) Name() string {
	return "DPT_TimePeriodMsec"
}

// String implements IDpt.
func (d Dpt7002) String() string {
	return d.Value().String()
}

// ToBytes implements IDpt.
func (d Dpt7002) Pack() []byte {
	return packU16(U16(d.Value().Milliseconds()))
}

// Unit implements IDpt.
func (d Dpt7002) Unit() string {
	return "ms"
}

func UnpackDpt7002(data []byte) (IDpt[time.Duration], error) {
	v, err := unpackU16(data)
	return Dpt7002(time.Duration(v) * time.Millisecond), err
}
