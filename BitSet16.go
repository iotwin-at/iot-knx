package knx

import "fmt"

type BitSet16 [16]bool

func (b BitSet16) String() string {
	return fmt.Sprintf("[%t %t %t %t %t %t %t %t %t %t %t %t %t %t %t %t]", b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15])
}

func (b BitSet16) Pack() []byte {
	result := make([]byte, 2)
	bIdx := -1
	for i := 0; i < 16; i++ {
		if i%8 == 0 {
			bIdx++
		}
		if b[i] {
			result[bIdx] |= (0b0000_0001 << i)
		}
	}
	return result
}

func NewBitSet16(data []byte) (BitSet16, error) {
	if len(data) != 2 {
		return BitSet16{}, NewErrInvalidDataType("given data cannot be unpacked to BitSet16 value")
	}
	result := BitSet16{}
	bIdx := -1
	for i := 0; i < 16; i++ {
		if i%8 == 0 {
			bIdx++
		}
		result[i] = (data[bIdx]&(0b0000_0001<<i))>>i == 1
	}
	return result, nil
}
