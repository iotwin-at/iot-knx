package dpt

import (
	"encoding/binary"
	"fmt"

	knx "github.com/iotwin-at/iot-knx"
)

// -----------------------------------------------------------------
// B¹ Datatype
// -----------------------------------------------------------------
type B1 bool

func (b B1) String() string {
	return fmt.Sprintf("%v", bool(b))
}

func packB1(value B1) []byte {
	if value {
		return []byte{0x01}
	}
	return []byte{0x00}
}

func unpackB1(data []byte) (B1, error) {
	if len(data) != 1 {
		return false, knx.NewErrInvalidDataType("given data cannot be unpacked as B1 value")
	}
	return (data[0] & 0x01) == 1, nil
}

// -----------------------------------------------------------------
// B² Datatype
// -----------------------------------------------------------------
type B2 struct {
	C bool // 0..no control, 1..control
	V bool // According to Type 1.xxx
}

func (b B2) String() string {
	if !b.C {
		return "No control"
	}
	if !b.V {
		return "Control. Function value 0"
	}
	return "Control. Function value 1"
}

func packB2(value B2) []byte {
	var b byte = 0
	if value.V {
		b |= 0b00000001
	}
	if value.C {
		b |= 0b00000010
	}
	return []byte{b}
}

func unpackB2(data []byte) (B2, error) {
	if len(data) != 1 {
		return B2{}, knx.NewErrInvalidDataType("given data cannot be unpacked to B2 value")
	}
	return B2{
		V: (data[0] & 0b00000001) == 1,
		C: ((data[0] & 0b00000010) >> 1) == 1,
	}, nil
}

// -----------------------------------------------------------------
// B¹U³ Datatype
// -----------------------------------------------------------------
type B1U3 struct {
	C        bool
	StepCode uint8
}

func (b B1U3) String() string {
	return fmt.Sprintf("C: %v, StepCode: %v", b.C, b.StepCode)
}

func packB1U3(value B1U3) []byte {
	var result byte
	if value.C {
		result |= 0b00001000
	}
	result |= (value.StepCode & 0b00000111)

	return []byte{result}
}

func unpackB1U3(data []byte) (B1U3, error) {
	if len(data) != 1 {
		return B1U3{}, knx.NewErrInvalidDataType("given data cannot be unpacked to B1U3 value")
	}
	return B1U3{
		C:        (data[0] & 0b00001000) == 1,
		StepCode: (data[0] & 0b00000111),
	}, nil
}

// -----------------------------------------------------------------
// A8 Datatype
// -----------------------------------------------------------------
type A8 byte

func (b A8) String() string {
	return string(rune(b))
}

func packA8(value A8) []byte {
	return []byte{byte(value)}
}

func unpackA8(data []byte) (A8, error) {
	if len(data) != 1 {
		return A8(0), knx.NewErrInvalidDataType("given data cannot be unpacked to A8 value")
	}
	return A8(data[0]), nil
}

// -----------------------------------------------------------------
// U8 Datatype
// -----------------------------------------------------------------
type U8 uint8

func (b U8) String() string {
	return fmt.Sprintf("%v", uint8(b))
}

func packU8(value U8) []byte {
	return []byte{byte(value)}
}

func unpackU8(data []byte) (U8, error) {
	if len(data) != 1 {
		return U8(0), knx.NewErrInvalidDataType("given data cannot be unpacked to A8 value")
	}
	return U8(data[0]), nil
}

// -----------------------------------------------------------------
// V8 Datatype
// -----------------------------------------------------------------
type V8 int8

func (b V8) String() string {
	return fmt.Sprintf("%v", int8(b))
}

func packV8(value V8) []byte {
	return []byte{byte(value)}
}

func unpackV8(data []byte) (V8, error) {
	if len(data) != 1 {
		return V8(0), knx.NewErrInvalidDataType("given data cannot be unpacked to V8 value")
	}
	return V8(data[0]), nil
}

// -----------------------------------------------------------------
// B5N3 Datatype
// -----------------------------------------------------------------
type B5N3 struct {
	A    bool
	B    bool
	C    bool
	D    bool
	E    bool
	Mode int8 // Active mode (F) - 0..Mode_0, 1..Mode_1, 2..Mode_2, -1..Invalid mode
}

func (b B5N3) String() string {
	return fmt.Sprintf("A: %v, B: %v, C: %v, D: %v, E: %v, Mode(F): %v", b.A, b.B, b.C, b.D, b.E, b.Mode)
}

func packB5N3(value B5N3) []byte {
	var result byte
	if value.A {
		result |= 0b10000000
	}
	if value.B {
		result |= 0b01000000
	}
	if value.C {
		result |= 0b00100000
	}
	if value.D {
		result |= 0b00010000
	}
	if value.E {
		result |= 0b00001000
	}
	switch value.Mode {
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

func unpackB5N3(data []byte) (B5N3, error) {
	if len(data) != 1 {
		return B5N3{}, knx.NewErrInvalidDataType("given data cannot be unpacked to V8 value")
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
	return B5N3{
		A:    ((data[0] & 0b10000000) >> 7) == 1,
		B:    ((data[0] & 0b01000000) >> 6) == 1,
		C:    ((data[0] & 0b00100000) >> 5) == 1,
		D:    ((data[0] & 0b00010000) >> 4) == 1,
		E:    ((data[0] & 0b00001000) >> 3) == 1,
		Mode: mode,
	}, nil
}

// -----------------------------------------------------------------
// U16 Datatype
// -----------------------------------------------------------------
type U16 uint16

func (b U16) String() string {
	return fmt.Sprintf("%v", uint16(b))
}

func packU16(value U16) []byte {
	var data []byte
	binary.BigEndian.PutUint16(data, uint16(value))
	return data
}

func unpackU16(data []byte) (U16, error) {
	if len(data) != 2 {
		return U16(0), knx.NewErrInvalidDataType("given data cannot be unpacked to U16 value")
	}
	return U16(binary.BigEndian.Uint16(data)), nil
}

// -----------------------------------------------------------------
// V16 Datatype
// -----------------------------------------------------------------
type V16 int16

func (b V16) String() string {
	return fmt.Sprintf("%v", int16(b))
}

func packV16(value V16) []byte {
	var data []byte
	binary.BigEndian.PutUint16(data, uint16(value))
	return []byte{byte(value)}
}

func unpackV16(data []byte) (V16, error) {
	if len(data) != 2 {
		return V16(0), knx.NewErrInvalidDataType("given data cannot be unpacked to V16 value")
	}
	return V16(int16(binary.BigEndian.Uint16(data))), nil
}
