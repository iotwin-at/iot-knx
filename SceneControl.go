package knx

import (
	"fmt"
)

type SceneControl struct {
	C           bool
	SceneNumber uint8
}

func (s SceneControl) String() string {
	return fmt.Sprintf("%v", uint8(s.SceneNumber))
}

func (s SceneControl) Pack() []byte {
	result := []byte{0}
	// Encode C
	if s.C {
		result[0] |= 0b10000000
	}
	// Encode SceneNumber
	result[0] |= (s.SceneNumber & 0b00111111)
	return result
}

func NewSceneControl(data []byte) (SceneControl, error) {
	if len(data) != 1 {
		return SceneControl{}, NewErrInvalidDataType("given data cannot be unpacked to SceneControl value")
	}
	// Extract C
	c := ((data[0] & 0b10000000) >> 7) == 1
	// Extract SceneNumber
	sceneNumber := uint8(data[0] & 0b00111111)
	return SceneControl{
		C:           c,
		SceneNumber: sceneNumber,
	}, nil
}
