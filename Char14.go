package knx

type Char14 [14]byte

func (c Char14) String() string {
	return string(c[:])
}

func (c Char14) Pack() []byte {
	return c[:]
}

func NewChar14(data []byte) (Char14, error) {
	if len(data) != 14 {
		return Char14{}, NewErrInvalidDataType("given data cannot be unpacked to Char14 value")
	}
	result := Char14{}
	copy(result[:], data)
	return result, nil
}
