package knx

import (
	"encoding/binary"
	"fmt"
	"math"
)

type Float16 float32

func (f Float16) String() string {
	return fmt.Sprintf("%v", float32(f))
}

func (f Float16) Pack() []byte {
	var sign uint16 = 0
	var exponent uint16 = 0
	var mantissa int16
	// Handle sign
	absValue := float32(f)
	if f < 0 {
		sign = 1
		absValue = -absValue
	}
	// Scale value by 100 to work with 0.01 resolution
	scaledValue := absValue * 100.0
	// Find appropriate exponent (0-15) and mantissa (0-2047)
	// Formula: value = (0.01 * mantissa) * 2^exponent
	for exponent = 0; exponent < 15; exponent++ {
		mantissa = int16(float64(scaledValue) / math.Pow(2, float64(exponent)))
		if mantissa <= 2047 {
			break
		}
	}
	// limit mantissa
	if mantissa > 2047 {
		mantissa = 2047
	}
	// Apply sign to mantissa using two's complement if negative
	var mantissaUnsigned uint16
	if sign == 1 {
		mantissaUnsigned = uint16(^mantissa+1) & 0x07FF
	} else {
		mantissaUnsigned = uint16(mantissa) & 0x07FF
	}

	// Combine: [S][EEEE][MMMMMMMMMMM]
	// S = sign bit (bit 15)
	// E = exponent (bits 14-11)
	// M = mantissa (bits 10-0)
	result := (sign << 15) | (exponent << 11) | mantissaUnsigned

	// Convert to byte array (big-endian)
	return []byte{
		byte(result >> 8),   // High byte
		byte(result & 0xFF), // Low byte
	}
}

func NewFloat16(data []byte) (Float16, error) {
	// Check length
	if len(data) != 2 {
		return Float16(0), NewErrInvalidDataType("given data cannot be unpacked to Float16 value")
	}
	// Read as big-endian (MSB first)
	raw := binary.BigEndian.Uint16(data)

	// Extract components
	sign := ((raw & 0b1000000000000000) >> 15) == 1
	exponent := (raw & 0b0111100000000000) >> 11
	mantissa := raw & 0b0000011111111111

	// Convert mantissa from two's complement if negative
	var mantissaSigned int16
	if sign {
		// Two's complement: invert and add 1
		mantissaSigned = -int16((^mantissa & 0x07FF) + 1)
	} else {
		mantissaSigned = int16(mantissa)
	}

	// Calculate value: (0.01 * mantissa) * 2^exponent
	value := (0.01 * float64(mantissaSigned)) * math.Pow(2, float64(exponent))

	return Float16(value), nil
}
