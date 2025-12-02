package knx

import (
	"fmt"
)

type StepControl struct {
	C        bool
	StepCode uint8
}

func (b StepControl) String() string {
	return fmt.Sprintf("C: %v, StepCode: %v", b.C, b.StepCode)
}

func (b StepControl) Pack() []byte {
	var result byte
	if b.C {
		result |= 0b00001000
	}
	result |= (b.StepCode & 0b00000111)
	return []byte{result}
}

func NewStepControl(data []byte) (StepControl, error) {
	if len(data) != 1 {
		return StepControl{}, NewErrInvalidDataType("given data cannot be unpacked to StepControl value")
	}
	return StepControl{
		C:        (data[0] & 0b00001000) == 1,
		StepCode: (data[0] & 0b00000111),
	}, nil
}
