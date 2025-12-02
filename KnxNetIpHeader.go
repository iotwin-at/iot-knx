package knx

import "encoding/binary"

type KnxNetIpHeader struct {
	HeaderLength    uint8
	ProtocolVersion uint8
	ServiceType     KnxServiceType
	TotalLength     uint16
}

func (k *KnxNetIpHeader) Pack() []byte {
	result := make([]byte, 6)
	result[0] = k.HeaderLength
	result[1] = k.ProtocolVersion
	binary.BigEndian.PutUint16(result[2:4], uint16(k.ServiceType))
	binary.BigEndian.PutUint16(result[4:6], k.TotalLength)
	return result[:]
}

func parseKnxNetIpHeader(data []byte) (*KnxNetIpHeader, []byte, error) {
	if len(data) < 6 {
		return nil, data, NewErrInvalidDataframe("KnxNetIpHeader needs at least 6 bytes")
	}
	header := KnxNetIpHeader{
		HeaderLength:    data[0],
		ProtocolVersion: data[1],
		ServiceType:     KnxServiceType(binary.BigEndian.Uint16(data[2:4])),
		TotalLength:     binary.BigEndian.Uint16(data[4:6]),
	}
	return &header, data[6:], nil
}
