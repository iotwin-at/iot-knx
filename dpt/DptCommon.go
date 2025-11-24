package dpt

import knx "github.com/iotwin-at/iot-knx"

// -----------------------------------------------------------------
// B¹ Datatype
// -----------------------------------------------------------------
type B1 bool

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

func packU8(value U8) []byte {
	return []byte{byte(value)}
}

func unpackU8(data []byte) (U8, error) {
	if len(data) != 1 {
		return U8(0), knx.NewErrInvalidDataType("given data cannot be unpacked to A8 value")
	}
	return U8(data[0]), nil
}
