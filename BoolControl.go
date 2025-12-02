package knx

type BoolControl struct {
	C bool // 0..no control, 1..control
	V bool // According to Type 1.xxx
}

func (b BoolControl) String() string {
	if !b.C {
		return "No control"
	}
	if !b.V {
		return "Control. Function value 0"
	}
	return "Control. Function value 1"
}

func (b BoolControl) Pack() []byte {
	var result byte = 0
	if b.V {
		result |= 0b00000001
	}
	if b.C {
		result |= 0b00000010
	}
	return []byte{result}
}

func NewBoolControl(data []byte) (BoolControl, error) {
	if len(data) != 1 {
		return BoolControl{}, NewErrInvalidDataType("given data cannot be unpacked to BoolControl value")
	}
	return BoolControl{
		V: (data[0] & 0b00000001) == 1,
		C: ((data[0] & 0b00000010) >> 1) == 1,
	}, nil
}
