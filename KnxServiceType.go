package knx

type KnxServiceType uint16

const (
	SEARCH_REQUEST               = KnxServiceType(0x0201)
	SEARCH_RESPONSE              = KnxServiceType(0x0202)
	DESCRIPTION_REQUEST          = KnxServiceType(0x0203)
	DESCRIPTION_RESPONSE         = KnxServiceType(0x0204)
	CONNECT_REQUEST              = KnxServiceType(0x0205)
	CONNECT_RESPONSE             = KnxServiceType(0x0206)
	CONNECTIONSTATE_REQUEST      = KnxServiceType(0x0207)
	CONNECTIONSTATE_RESPONSE     = KnxServiceType(0x0208)
	DISCONNECT_REQUEST           = KnxServiceType(0x0209)
	DISCONNECT_RESPONSE          = KnxServiceType(0x020A)
	TUNNELLING_REQUEST           = KnxServiceType(0x0420)
	TUNNELLING_ACK               = KnxServiceType(0x0421)
	ROUTING_INDICATION           = KnxServiceType(0x0530)
	ROUTING_LOST_MESSAGE         = KnxServiceType(0x0531)
	ROUTING_BUSY                 = KnxServiceType(0x0532)
	DEVICE_CONFIGURATION_REQUEST = KnxServiceType(0x0310)
	DEVICE_CONFIGURATION_ACK     = KnxServiceType(0x0311)
)

// ------------------------------------------------------------------------------------
// SEARCH_REQUEST (0x0201)
// ------------------------------------------------------------------------------------
type SearchRequest struct {
	// TODO: Implement SearchRequest
}

func (r *SearchRequest) Pack() []byte {
	panic("not implemented")
}

func parseSearchRequest([]byte) (*SearchRequest, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// SEARCH_RESPONSE (0x0202)
// ------------------------------------------------------------------------------------
type SearchResponse struct {
	// TODO: Implement SearchResponse
}

func (r *SearchResponse) Pack() []byte {
	panic("not implemented")
}

func parseSearchResponse([]byte) (*SearchResponse, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// DESCRIPTION_REQUEST (0x0203)
// ------------------------------------------------------------------------------------
type DescriptionRequest struct {
	// TODO: Implement Description Request
}

func (r *DescriptionRequest) Pack() []byte {
	panic("not implemented")
}

func parseDescriptionRequest([]byte) (*DescriptionRequest, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// DESCRIPTION_RESPONSE (0x0204)
// ------------------------------------------------------------------------------------
type DescriptionResponse struct {
	// TODO: Implement Description Response
}

func (r *DescriptionResponse) Pack() []byte {
	panic("not implemented")
}

func parseDescriptionResponse([]byte) (*DescriptionResponse, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// CONNECT_REQUEST (0x0205)
// ------------------------------------------------------------------------------------
type ConnectRequest struct {
	// TODO: Connect Request
}

func (r *ConnectRequest) Pack() []byte {
	panic("not implemented")
}

func parseConnectRequest([]byte) (*ConnectRequest, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// CONNECT_RESPONSE (0x0206)
// ------------------------------------------------------------------------------------
type ConnectResponse struct {
	// TODO: Implement Connect Response
}

func (r *ConnectResponse) Pack() []byte {
	panic("not implemented")
}

func parseConnectResponse([]byte) (*ConnectResponse, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// CONNECTIONSTATE_REQUEST (0x0207)
// ------------------------------------------------------------------------------------
type ConnectionStateRequest struct {
	// TODO: Implement Connection State Request
}

func (r *ConnectionStateRequest) Pack() []byte {
	panic("not implemented")
}

func parseConnectionStateRequest([]byte) (*ConnectionStateRequest, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// CONNECTIONSTATE_RESPONSE (0x0208)
// ------------------------------------------------------------------------------------
type ConnectionStateResponse struct {
	// TODO: Implement Connection State Response
}

func (r *ConnectionStateResponse) Pack() []byte {
	panic("not implemented")
}

func parseConnectionStateResponse([]byte) (*ConnectionStateResponse, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// DISCONNECT_REQUEST (0x0209)
// ------------------------------------------------------------------------------------
type DisconnectRequest struct {
	// TODO: Implement Disconnect Request
}

func (r *DisconnectRequest) Pack() []byte {
	panic("not implemented")
}

func parseDisconnectRequest([]byte) (*DisconnectRequest, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// DISCONNECT_RESPONSE (0x020A)
// ------------------------------------------------------------------------------------
type DisconnectResponse struct {
	// TODO: Implement Disconnect Response
}

func (r *DisconnectResponse) Pack() []byte {
	panic("not implemented")
}

func parseDisconnectResponse([]byte) (*DisconnectResponse, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// TUNNELLING_REQUEST (0x0420)
// ------------------------------------------------------------------------------------
type TunnellingRequest struct {
	// TODO: Implement Tunnelling Request
}

func (r *TunnellingRequest) Pack() []byte {
	panic("not implemented")
}

func parseTunnellingRequest([]byte) (*TunnellingRequest, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// TUNNELLING_ACK (0x0421)
// ------------------------------------------------------------------------------------
type TunnellingAck struct {
	// TODO: Implement Tunneling Response
}

func (r *TunnellingAck) Pack() []byte {
	panic("not implemented")
}

func parseTunnellingAck([]byte) (*TunnellingAck, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// ROUTING_INDICATION (0x0530)
// ------------------------------------------------------------------------------------
type RoutingIndication struct {
	Cemi
}

func (r *RoutingIndication) Pack() []byte {
	return r.Cemi.Pack()
}

func parseRoutingIndication(data []byte) (*RoutingIndication, error) {
	cemi, err := parseKnxCemi(data)
	if err != nil {
		return nil, err
	}
	return &RoutingIndication{
		Cemi: *cemi,
	}, nil
}

// ------------------------------------------------------------------------------------
// ROUTING_LOST_MESSAGE (0x0531)
// ------------------------------------------------------------------------------------
type RoutingLostMessage struct {
	// TODO: Implement Routing Lost Message
}

func (r *RoutingLostMessage) Pack() []byte {
	panic("not implemented")
}

func parseRoutingLostMessage([]byte) (*RoutingLostMessage, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// ROUTING_BUSY (0x0532)
// ------------------------------------------------------------------------------------
type RoutingBusy struct {
	// TODO: Implement Routing Busy Message
}

func (r *RoutingBusy) Pack() []byte {
	panic("not implemented")
}

func parseRoutingBusy([]byte) (*RoutingBusy, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// DEVICE_CONFIGURATION_REQUEST (0x0310)
// ------------------------------------------------------------------------------------
type DeviceConfigurationRequest struct {
	// TODO: Implement Device Configuration Request
}

func (r *DeviceConfigurationRequest) Pack() []byte {
	panic("not implemented")
}

func parseDeviceConfigurationRequest([]byte) (*DeviceConfigurationRequest, error) {
	panic("not implemented")
}

// ------------------------------------------------------------------------------------
// DEVICE_CONFIGURATION_ACK (0x0311)
// ------------------------------------------------------------------------------------
type DeviceConfigurationAck struct {
	// TODO: Implment Device Configuration Ack
}

func (r *DeviceConfigurationAck) Pack() []byte {
	panic("not implemented")
}

func parseSeaDeviceConfigurationAck([]byte) (*DeviceConfigurationAck, error) {
	panic("not implemented")
}
