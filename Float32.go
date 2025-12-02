package knx

import (
	"encoding/binary"
	"fmt"
	"math"
)

// Float32Value represents a 4-octet floating point value (F₃₂) datapoint type.
//
// Format: 4 octets organized as follows:
//
//	            ┌─┬────────┬─────────────────────────
//	Field:      │S│Exponent│       Fraction         │
//	            └─┴────────┴─────────────────────────
//	Encoding:   │F│FFFFFFFF│FFFFFFFFFFFFFFFFFFFFFFFF│
//	Bits:       │1│   8    │           23           │
//
// Encoding: Values are encoded in IEEE 754 single precision floating point format.
// The exponent is biased, which allows for negative exponent values.
//
// Field Ranges:
//   - Sign (S): {0, 1}
//   - Exponent: [0 ... 255]
//   - Fraction: [0 ... 8,388,607]
//
// Resolution: Determined by the IEEE 754 format and varies with the exponent used.
//

type Float32 float32

func (f Float32) String() string {
	return fmt.Sprintf("%v", float32(f))
}

func (f Float32) Pack() []byte {
	b := []byte{0, 0, 0, 0}
	bits := math.Float32bits(float32(f))
	binary.BigEndian.PutUint32(b, bits)
	return b
}

func NewFloat32(data []byte) (Float32, error) {
	// Check length
	if len(data) != 4 {
		return Float32(0), NewErrInvalidDataType("given data cannot be unpacked to Float32 value")
	}
	bits := binary.BigEndian.Uint32(data)
	return Float32(math.Float32frombits(bits)), nil
}
