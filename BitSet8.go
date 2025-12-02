package knx

import "fmt"

type BitSet8 [8]bool

func (b BitSet8) String() string {
	return fmt.Sprintf("[%t %t %t %t %t %t %t %t]", b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7])
}

func (b BitSet8) Pack() []byte {
	result := []byte{0}
	bIdx := -1
	for i := 0; i < 8; i++ {
		if i%8 == 0 {
			bIdx++
		}
		if b[i] {
			result[bIdx] |= (0b0000_0001 << i)
		}
	}
	return result
}

func NewBitSet8(data []byte) (BitSet8, error) {
	if len(data) != 1 {
		return BitSet8{}, NewErrInvalidDataType("given data cannot be unpacked to BitSet8 value")
	}
	result := BitSet8{}
	bIdx := -1
	for i := 0; i < 8; i++ {
		if i%8 == 0 {
			bIdx++
		}
		result[i] = (data[bIdx]&(0b0000_0001<<i))>>i == 1
	}
	return result, nil
}
