package dpt

import "time"

type Dpt8004 time.Duration

// Value implements IDpt.
func (d Dpt8004) Value() time.Duration {
	return time.Duration(d)
}

// Name implements IDpt.
func (d Dpt8004) Name() string {
	return "DPT_DeltaTime100Msec"
}

// String implements IDpt.
func (d Dpt8004) String() string {
	return d.Value().String()
}

// ToBytes implements IDpt.
func (d Dpt8004) Pack() []byte {
	return packV16(V16(d.Value().Milliseconds() / 100))
}

// Unit implements IDpt.
func (d Dpt8004) Unit() string {
	return "ms"
}

func UnpackDpt8004(data []byte) (IDpt[time.Duration], error) {
	v, err := unpackV16(data)
	return Dpt8004(time.Duration(v) * time.Millisecond * 100), err
}
