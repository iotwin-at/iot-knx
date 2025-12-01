package dpt

import "time"

type Dpt8003 time.Duration

// Value implements IDpt.
func (d Dpt8003) Value() time.Duration {
	return time.Duration(d)
}

// Name implements IDpt.
func (d Dpt8003) Name() string {
	return "DPT_DeltaTime10Msec"
}

// String implements IDpt.
func (d Dpt8003) String() string {
	return d.Value().String()
}

// ToBytes implements IDpt.
func (d Dpt8003) Pack() []byte {
	return packV16(V16(d.Value().Milliseconds() / 10))
}

// Unit implements IDpt.
func (d Dpt8003) Unit() string {
	return "ms"
}

func UnpackDpt8003(data []byte) (IDpt[time.Duration], error) {
	v, err := unpackV16(data)
	return Dpt8003(time.Duration(v) * time.Millisecond * 10), err
}
