package knx

type String string

func (s String) String() string {
	return string(s)
}

func (s String) Pack() []byte {
	value := string(s)
	result := make([]byte, len(value)+1)
	for i, char := range value {
		if char > 255 {
			result[i] = '?' // No valid ISO 8859-1 character
		} else {
			result[i] = byte(char)
		}
	}
	result[len(result)-1] = 0x00
	return result
}

func NewString(data []byte) (String, error) {
	if len(data) <= 0 {
		return String(""), NewErrInvalidDataType("given data cannot be unpacked to String value")
	}
	length := len(data)
	for i := 0; i < length; i++ {
		if data[i] == 0x00 { // Search NULL character (like C-String)
			length = i
			break
		}
	}
	chars := make([]rune, length)
	for i := 0; i < length; i++ {
		chars[i] = rune(data[i])
	}
	return String(chars), nil
}
