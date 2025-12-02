package knx

import (
	"fmt"
)

type StatusWithMode struct {
	A    bool
	B    bool
	C    bool
	D    bool
	E    bool
	Mode int8 // Active mode (F) - 0..Mode_0, 1..Mode_1, 2..Mode_2, -1..Invalid mode
}

func (b StatusWithMode) String() string {
	return fmt.Sprintf("A: %v, B: %v, C: %v, D: %v, E: %v, Mode(F): %v", b.A, b.B, b.C, b.D, b.E, b.Mode)
}

func (b StatusWithMode) Pack() []byte {
	var result byte
	if b.A {
		result |= 0b10000000
	}
	if b.B {
		result |= 0b01000000
	}
	if b.C {
		result |= 0b00100000
	}
	if b.D {
		result |= 0b00010000
	}
	if b.E {
		result |= 0b00001000
	}
	switch b.Mode {
	case 0: // 001b
		result |= 0b00000001
	case 1: // 010b
		result |= 0b00000010
	case 2: // 100b
		result |= 0b00000100
	default:
		// ignore mode if not valid
	}
	return []byte{result}
}

func NewStatusWithMode(data []byte) (StatusWithMode, error) {
	if len(data) != 1 {
		return StatusWithMode{}, NewErrInvalidDataType("given data cannot be unpacked to V8 value")
	}

	// C: ((data[0] & 0b00000010) >> 1) == 1,
	var mode int8 = -1
	switch data[0] & 0b00000111 {
	case 0b00000001:
		mode = 0
	case 0b00000010:
		mode = 1
	case 0b00000100:
		mode = 2
	}
	return StatusWithMode{
		A:    ((data[0] & 0b10000000) >> 7) == 1,
		B:    ((data[0] & 0b01000000) >> 6) == 1,
		C:    ((data[0] & 0b00100000) >> 5) == 1,
		D:    ((data[0] & 0b00010000) >> 4) == 1,
		E:    ((data[0] & 0b00001000) >> 3) == 1,
		Mode: mode,
	}, nil
}
