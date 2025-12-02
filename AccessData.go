package knx

// AccessData represents a 4-octet access data value (DPT_Access_Data #015.000).
//
// Format: 4 octets (U₄U₄U₄U₄U₄U₄B₄N₄) organized as follows:
//
//	            |--------|--------|--------|--------|--------|--------|----|------|
//	Field:      |   D6   |   D5   |   D4   |   D3   |   D2   |   D1   |EPDC|Index |
//	            |--------|--------|--------|--------|--------|--------|----|------|
//	Encoding:   |UUUUUUUU|UUUUUUUU|UUUUUUUU|UUUUUUUU|UUUUUUUU|UUUUUUUU|bbbb| NNNN |
//	Bits:       |   8    |   8    |   8    |   8    |   8    |   8    | 4  |  4   |
//
// Encoding:
//   - D₆, D₅, D₄, D₃, D₂, D₁: Binary encoded values (unsigned, 8 bits each)
//   - N (Index): Binary encoded value (4 bits)
//   - E, P, D, C: Binary flags (1 bit each) - see specification for details
//

type AccessData struct {
	AccessIdentificationCode [6]uint8
	DetectionError           bool
	Permission               bool
	ReadDirection            bool
	Encrytion                bool
	Index                    uint8
}

func (a AccessData) Pack() []byte {
	result := []byte{0, 0, 0, 0}
	// Set D1-D6
	result[0] |= ((a.AccessIdentificationCode[0] & 0x0F) << 4) // D6
	result[0] |= (a.AccessIdentificationCode[1] & 0x0F)        // D5
	result[1] |= ((a.AccessIdentificationCode[2] & 0x0F) << 4) // D4
	result[1] |= (a.AccessIdentificationCode[3] & 0x0F)        // D3
	result[2] |= ((a.AccessIdentificationCode[4] & 0x0F) << 4) // D2
	result[2] |= (a.AccessIdentificationCode[5] & 0x0F)        // D1
	// Set E
	if a.DetectionError {
		result[3] |= 0b10000000
	}
	// Set P
	if a.Permission {
		result[3] |= 0b01000000
	}
	// Set D
	if a.ReadDirection {
		result[3] |= 0b00100000
	}
	// Set C
	if a.Encrytion {
		result[3] |= 0b00010000
	}

	return result
}

func NewAccessData(data []byte) (AccessData, error) {
	if len(data) != 4 {
		return AccessData{}, NewErrInvalidDataType("given data cannot be unpacked to UInt32 value")
	}
	// Extract D1-D6
	d6 := uint8(data[0]&0xF0) >> 4
	d5 := uint8(data[0] & 0x0F)
	d4 := uint8(data[1]&0xF0) >> 4
	d3 := uint8(data[1] & 0x0F)
	d2 := uint8(data[2]&0xF0) >> 4
	d1 := uint8(data[2] & 0x0F)
	// Extract E
	e := (data[3]&0b10000000)>>7 == 1
	// Extract P
	p := (data[3]&0b01000000)>>6 == 1
	// Extract D
	d := (data[3]&0b00100000)>>5 == 1
	// Extract C
	c := (data[3]&0b00010000)>>4 == 1
	// Extract Index
	index := uint8(data[3] & 0x0F)
	// Return
	return AccessData{
		AccessIdentificationCode: [6]uint8{d6, d5, d4, d3, d2, d1},
		DetectionError:           e,
		Permission:               p,
		ReadDirection:            d,
		Encrytion:                c,
		Index:                    index,
	}, nil
}
