package knx

import (
	"fmt"
)

type SceneNumber uint8

func (s SceneNumber) String() string {
	return fmt.Sprintf("%v", uint8(s))
}

func (s SceneNumber) Pack() []byte {
	result := []byte{0}
	result[0] |= (uint8(s) & 0b00111111)
	return result
}

func NewSceneNumber(data []byte) (SceneNumber, error) {
	if len(data) != 1 {
		return SceneNumber(0), NewErrInvalidDataType("given data cannot be unpacked to SceneNumber value")
	}
	return SceneNumber(data[0] & 0b00111111), nil
}
