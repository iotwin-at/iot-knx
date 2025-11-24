package knx

import (
	"encoding/binary"
	"fmt"
)

// ----------------------------------------------------------------------------------------
// Destination Address
// ----------------------------------------------------------------------------------------

type DestAddr struct {
	Main   uint16
	Middle uint16
	Sub    uint16
}

func (d *DestAddr) ToBytes() []byte {
	rawDest := (d.Main&0x1F)<<11 |
		(d.Middle&0x07)<<8 |
		(d.Sub & 0xFF)
	result := make([]byte, 2)
	binary.BigEndian.PutUint16(result, rawDest)
	return result
}

func parseDestAddr(data []byte) (*DestAddr, error) {
	if len(data) != 2 {
		return nil, NewErrInvalidDataframe("destination address must consist of 2 bytes")
	}
	destAddr := &DestAddr{}
	rawDest := binary.BigEndian.Uint16(data)
	destAddr.Main = (rawDest >> 11) & 0x1F
	destAddr.Middle = (rawDest >> 8) & 0x07
	destAddr.Sub = rawDest & 0xFF
	return destAddr, nil
}

// ----------------------------------------------------------------------------------------
// Source Address
// ----------------------------------------------------------------------------------------

type SrcAddr struct {
	Area   uint16
	Line   uint16
	Device uint16
}

func (s *SrcAddr) ToBytes() []byte {
	rawSrc := (s.Area&0x0F)<<12 |
		(s.Line&0x0F)<<8 |
		(s.Device & 0xFF)
	result := make([]byte, 2)
	binary.BigEndian.PutUint16(result, rawSrc)
	return result
}

func parseSrcAddr(data []byte) (*SrcAddr, error) {
	if len(data) != 2 {
		return nil, NewErrInvalidDataframe("source address must consist of 2 bytes")
	}
	srcAddr := &SrcAddr{}
	rawSrc := binary.BigEndian.Uint16(data)
	srcAddr.Area = (rawSrc >> 12) & 0x0F
	srcAddr.Line = (rawSrc >> 8) & 0x0F
	srcAddr.Device = rawSrc & 0xFF
	return srcAddr, nil
}

// ----------------------------------------------------------------------------------------
// Cemi Frame
// ----------------------------------------------------------------------------------------

type Cemi struct {
	MessageCode       uint8
	AdditionalInfoLen uint8
	AdditionalInfo    []byte
	ControlField1     uint8
	ControlField2     uint8
	SourceAddr        SrcAddr
	DestAddr          DestAddr
	DataLength        uint8
	TPCI              uint8  // Transport Protocol Control Information
	APCI              Apci   // Application Protocol Control Information (Command)
	Data              []byte // The actual payload data
}

func (c *Cemi) ToBytes() []byte {
	// Calculate total frame size
	// 2 (MessageCode + AdditionalInfoLen) + AdditionalInfo + 9 (core fields) + Data
	coreSize := 2 + int(c.AdditionalInfoLen) + 9
	totalSize := coreSize + len(c.Data)

	// Handle short payload case (value embedded in APCI byte)
	payloadLen := int(c.DataLength) - 1
	if payloadLen <= 0 {
		totalSize = coreSize
	}

	result := make([]byte, totalSize)

	// Byte 0: Message Code
	result[0] = c.MessageCode

	// Byte 1: Additional Info Length
	result[1] = c.AdditionalInfoLen

	// Bytes 2 to 2+AdditionalInfoLen: Additional Info
	if c.AdditionalInfoLen > 0 {
		copy(result[2:2+c.AdditionalInfoLen], c.AdditionalInfo)
	}

	// Calculate offset for core fields
	offset := 2 + int(c.AdditionalInfoLen)

	// Offset N: Control Field 1
	result[offset] = c.ControlField1

	// Offset N+1: Control Field 2
	result[offset+1] = c.ControlField2

	// Offset N+2 to N+3: Source Address (using SrcAddr.ToBytes())
	srcBytes := c.SourceAddr.ToBytes()
	copy(result[offset+2:offset+4], srcBytes)

	// Offset N+4 to N+5: Destination Address (using DestAddr.ToBytes())
	destBytes := c.DestAddr.ToBytes()
	copy(result[offset+4:offset+6], destBytes)

	// Offset N+6: Data Length
	result[offset+6] = c.DataLength

	// Offset N+7: TPCI (2 bits) + upper 2 bits of APCI (4 bits total in this byte)
	tpciApciByte1 := (c.TPCI & 0b00000011) << 6
	tpciApciByte1 |= byte((c.APCI >> 2) & 0b00000011)
	result[offset+7] = tpciApciByte1

	// Offset N+8: Lower 2 bits of APCI + data/value (6 bits)
	apciDataByte2 := byte((c.APCI & 0b00000011) << 6)

	// Handle payload
	if payloadLen > 0 {
		// Long payload: data follows in subsequent bytes
		copy(result[offset+9:], c.Data)
	} else {
		// Short payload: value is embedded in the last 6 bits of APCI byte
		if len(c.Data) > 0 {
			apciDataByte2 |= c.Data[0] & 0b00111111
		}
	}
	result[offset+8] = apciDataByte2
	return result
}

func parseKnxCemi(data []byte) (*Cemi, error) {
	frame := &Cemi{}
	// Check min length
	if len(data) < 11 {
		return frame, NewErrInvalidDataframe("invalid CEMI frame length: got %d, want at least 11", len(data))
	}
	// Byte 0: Message Code
	frame.MessageCode = data[0]
	// Byte 1: Additional Info Length
	frame.AdditionalInfoLen = data[1]
	// --- Variable Offset Section ---
	// Calculate the base offset for the main telegram data.
	offset := 2 + int(frame.AdditionalInfoLen)
	// Extract Additional Info if it exists
	if frame.AdditionalInfoLen > 0 {
		if len(data) < 2+int(frame.AdditionalInfoLen) {
			return frame, fmt.Errorf("not enough data for additional info")
		}
		frame.AdditionalInfo = make([]byte, frame.AdditionalInfoLen)
		copy(frame.AdditionalInfo, data[2:offset])
	}
	// Check if the remaining data is long enough for the core fields (9 bytes)
	if len(data) < offset+9 {
		return frame, fmt.Errorf("not enough data for core CEMI fields")
	}
	// Offset N:   Control Field 1
	frame.ControlField1 = data[offset]
	// Offset N+1: Control Field 2
	frame.ControlField2 = data[offset+1]
	// Decode source address (area.line.device)
	src, err := parseSrcAddr(data[offset+2 : offset+4])
	if err != nil {
		return frame, err
	}
	frame.SourceAddr = *src
	// Decode destination address (main.middle.sub)
	dest, err := parseDestAddr(data[offset+4 : offset+6])
	if err != nil {
		return frame, err
	}
	frame.DestAddr = *dest
	// Offset N+6: Data Length (length of TPCI/APCI + Payload)
	frame.DataLength = data[offset+6]
	// The TPCI and APCI are encoded across the next two bytes.
	tpciApciByte1 := data[offset+7]
	apciDataByte2 := data[offset+8]
	// Extract TPCI (first 2 bits of byte at N+7)
	frame.TPCI = (tpciApciByte1 & 0xC0) >> 6
	// Extract APCI (Command, e.g., GroupValue_Read/Write)
	// It's formed by the last 2 bits of the first byte and the first 2 bits of the second byte.
	frame.APCI = (Apci(tpciApciByte1&0x03) << 2) | Apci((apciDataByte2&0xC0)>>6)
	// The DataLength field includes the TPCI/APCI byte, so the payload is DataLength - 1.
	payloadLen := int(frame.DataLength) - 1
	if payloadLen > 0 {
		// If the payload is longer than 6 bits, it's in the following bytes.
		dataStartOffset := offset + 9
		if len(data) < dataStartOffset+payloadLen {
			return frame, fmt.Errorf("data length mismatch: expected %d payload bytes, but not enough data", payloadLen)
		}
		frame.Data = make([]byte, payloadLen)
		copy(frame.Data, data[dataStartOffset:dataStartOffset+payloadLen])
	} else {
		// For short payloads (e.g., boolean On/Off), the value is in the last 6 bits
		// of the APCI/Data byte. We can store this in the Data field for consistency.
		value := apciDataByte2 & 0x3F
		frame.Data = []byte{value}
	}
	return frame, nil
}
