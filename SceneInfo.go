package knx

import "fmt"

type SceneInfo struct {
	IsActive    bool // B
	SceneNumber uint8
}

func (s SceneInfo) String() string {
	return fmt.Sprintf("Active: %v SceneNumber: %d", s.IsActive, s.SceneNumber)
}

func (s SceneInfo) Pack() []byte {
	result := make([]byte, 1)
	if s.IsActive {
		result[0] |= 0b0100_0000
	}
	result[0] |= s.SceneNumber
	return result
}

func NewSceneInfo(data []byte) (SceneInfo, error) {
	if len(data) != 1 {
		return SceneInfo{}, NewErrInvalidDataType("given data cannot be unpacked to SceneInfo value")
	}
	return SceneInfo{
		IsActive:    ((data[0] & 0b0100_0000) >> 6) == 1,
		SceneNumber: (data[0] & 0b0011_1111),
	}, nil
}
