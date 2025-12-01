package dpt

import "time"

type Dpt7006 time.Duration

// Value implements IDpt.
func (d Dpt7006) Value() time.Duration {
	return time.Duration(d)
}

// Name implements IDpt.
func (d Dpt7006) Name() string {
	return "DPT_TimePeriodMin"
}

// String implements IDpt.
func (d Dpt7006) String() string {
	return d.Value().String()
}

// ToBytes implements IDpt.
func (d Dpt7006) Pack() []byte {
	return packU16(U16(d.Value().Minutes()))
}

// Unit implements IDpt.
func (d Dpt7006) Unit() string {
	return "min"
}

func UnpackDpt7006(data []byte) (IDpt[time.Duration], error) {
	v, err := unpackU16(data)
	return Dpt7006(time.Duration(v) * time.Minute), err
}
