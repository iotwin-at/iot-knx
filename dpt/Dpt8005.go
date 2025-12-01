package dpt

import "time"

type Dpt8005 time.Duration

// Value implements IDpt.
func (d Dpt8005) Value() time.Duration {
	return time.Duration(d)
}

// Name implements IDpt.
func (d Dpt8005) Name() string {
	return "DPT_DeltaTimeSec"
}

// String implements IDpt.
func (d Dpt8005) String() string {
	return d.Value().String()
}

// ToBytes implements IDpt.
func (d Dpt8005) Pack() []byte {
	return packV16(V16(d.Value().Seconds()))
}

// Unit implements IDpt.
func (d Dpt8005) Unit() string {
	return "s"
}

func UnpackDpt8005(data []byte) (IDpt[time.Duration], error) {
	v, err := unpackV16(data)
	return Dpt8005(time.Duration(v) * time.Second), err
}
