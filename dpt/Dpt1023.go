package dpt

type Dpt1023 B1

// Value implements IDpt.
func (d Dpt1023) Value() B1 {
	return B1(d)
}

// Name implements IDpt.
func (d Dpt1023) Name() string {
	return "DPT_ShutterBlinds_Mode"
}

// String implements IDpt.
func (d Dpt1023) String() string {
	if d {
		return "move Up/Down + StepStop mode (blind)"
	}
	return "only move Up/Down mode (shutter)"
}

// ToBytes implements IDpt.
func (d Dpt1023) Pack() []byte {
	return packB1(d.Value())
}

// Unit implements IDpt.
func (d Dpt1023) Unit() string {
	return ""
}

func UnpackDpt1023(data []byte) (IDpt[B1], error) {
	v, err := unpackB1(data)
	return Dpt1023(v), err
}
