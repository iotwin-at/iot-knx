package knx

type KnxNetIpPackage[T ISerializable] struct {
	Header KnxNetIpHeader
	Body   T
}

func (k *KnxNetIpPackage[T]) Pack() []byte {
	header := k.Header.Pack()
	body := k.Body.Pack()
	result := make([]byte, len(header)+len(body))
	copy(result[0:6], header)
	copy(result[6:], body)
	return result
}

func parseKnxNetIpPackage[T ISerializable](data []byte) (*KnxNetIpPackage[T], error) {
	// Parse Header
	header, remainingData, err := parseKnxNetIpHeader(data)
	if err != nil {
		return nil, err
	}
	// Parse Body
	var zero T
	var body any
	switch any(zero).(type) {
	case *SearchRequest:
		body, err = parseSearchRequest(remainingData)
		if err != nil {
			return nil, err
		}
	case *SearchResponse:
		body, err = parseSearchRequest(remainingData)
		if err != nil {
			return nil, err
		}
	case *DescriptionRequest:
		body, err = parseDescriptionRequest(remainingData)
		if err != nil {
			return nil, err
		}
	case *DescriptionResponse:
		body, err = parseDescriptionResponse(remainingData)
		if err != nil {
			return nil, err
		}
	case *ConnectRequest:
		body, err = parseConnectRequest(remainingData)
		if err != nil {
			return nil, err
		}
	case *ConnectResponse:
		body, err = parseConnectResponse(remainingData)
		if err != nil {
			return nil, err
		}
	case *ConnectionStateRequest:
		body, err = parseConnectionStateRequest(remainingData)
		if err != nil {
			return nil, err
		}
	case *ConnectionStateResponse:
		body, err = parseConnectionStateResponse(remainingData)
		if err != nil {
			return nil, err
		}
	case *DisconnectRequest:
		body, err = parseDisconnectRequest(remainingData)
		if err != nil {
			return nil, err
		}
	case *DisconnectResponse:
		body, err = parseDisconnectResponse(remainingData)
		if err != nil {
			return nil, err
		}
	case *TunnellingRequest:
		body, err = parseTunnellingRequest(remainingData)
		if err != nil {
			return nil, err
		}
	case *TunnellingAck:
		body, err = parseTunnellingAck(remainingData)
		if err != nil {
			return nil, err
		}
	case *RoutingIndication:
		body, err = parseRoutingIndication(remainingData)
		if err != nil {
			return nil, err
		}
	case *RoutingLostMessage:
		body, err = parseRoutingLostMessage(remainingData)
		if err != nil {
			return nil, err
		}
	case *RoutingBusy:
		body, err = parseRoutingBusy(remainingData)
		if err != nil {
			return nil, err
		}
	case *DeviceConfigurationRequest:
		body, err = parseDeviceConfigurationRequest(remainingData)
		if err != nil {
			return nil, err
		}
	case *DeviceConfigurationAck:
		body, err = parseSeaDeviceConfigurationAck(remainingData)
		if err != nil {
			return nil, err
		}
	default:
		return nil, NewErrInvalidDataType("Given type T is not a supported ServiceType payload")
	}
	return &KnxNetIpPackage[T]{
		Header: *header,
		Body:   body.(T),
	}, nil
}
