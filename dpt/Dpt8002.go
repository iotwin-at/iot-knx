package dpt

import "time"

type Dpt8002 time.Duration

// Value implements IDpt.
func (d Dpt8002) Value() time.Duration {
	return time.Duration(d)
}

// Name implements IDpt.
func (d Dpt8002) Name() string {
	return "DPT_DeltaTimeMsec"
}

// String implements IDpt.
func (d Dpt8002) String() string {
	return d.Value().String()
}

// ToBytes implements IDpt.
func (d Dpt8002) Pack() []byte {
	return packV16(V16(d.Value().Milliseconds()))
}

// Unit implements IDpt.
func (d Dpt8002) Unit() string {
	return "ms"
}

func UnpackDpt8002(data []byte) (IDpt[time.Duration], error) {
	v, err := unpackV16(data)
	return Dpt8002(time.Duration(v) * time.Millisecond), err
}
