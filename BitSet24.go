package knx

import "fmt"

type BitSet24 [24]bool

func (b BitSet24) String() string {
	return fmt.Sprintf("[%t %t %t %t %t %t %t %t %t %t %t %t %t %t %t %t %t %t %t %t %t %t %t %t]", b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23])
}

func (b BitSet24) Pack() []byte {
	result := make([]byte, 3)
	bIdx := -1
	for i := 0; i < 24; i++ {
		if i%8 == 0 {
			bIdx++
		}
		if b[i] {
			result[bIdx] |= (0b0000_0001 << i)
		}
	}
	return result
}

func NewBitSet24(data []byte) (BitSet24, error) {
	if len(data) != 3 {
		return BitSet24{}, NewErrInvalidDataType("given data cannot be unpacked to BitSet24 value")
	}
	result := BitSet24{}
	bIdx := -1
	for i := 0; i < 24; i++ {
		if i%8 == 0 {
			bIdx++
		}
		result[i] = (data[bIdx]&(0b0000_0001<<i))>>i == 1
	}
	return result, nil
}
