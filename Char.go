package knx

type Char byte

func (b Char) String() string {
	return string(rune(b))
}

func (b Char) Pack() []byte {
	return []byte{byte(b)}
}

func NewChar(data []byte) (Char, error) {
	if len(data) != 1 {
		return Char(0), NewErrInvalidDataType("given data cannot be unpacked to Char value")
	}
	return Char(data[0]), nil
}
