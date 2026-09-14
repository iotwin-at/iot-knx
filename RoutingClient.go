package knx

import (
	"context"
	"log/slog"
	"net"
	"sync"
	"time"

	"github.com/uoul/go-async"
)

type RoutingClient struct {
	multicastAddr string
	retryInterval time.Duration

	conn    *net.UDPConn
	subs    map[async.Sequence[Cemi]]bool
	subMux  sync.RWMutex
	connMux sync.RWMutex
}

// ------------------------------------------------------------------------------------
// Public
// ------------------------------------------------------------------------------------

// Describe implements IClient.
func (r *RoutingClient) Describe(ctx context.Context, req DescriptionRequest) async.Result[KnxNetIpPackage[*DescriptionResponse]] {
	panic("unimplemented")
}

// Search implements IClient.
func (r *RoutingClient) Search(ctx context.Context, req SearchRequest) async.Result[KnxNetIpPackage[*SearchResponse]] {
	panic("unimplemented")
}

// Send implements IClient.
func (r *RoutingClient) Send(ctx context.Context, command GroupCommand) error {
	r.connMux.RLock()
	defer r.connMux.RUnlock()
	if r.conn == nil {
		return NewErrNetConnection("no udp connection available")
	}
	// TODO: Implement Frame conversion
	return NewErrNotImplemented("not implemented")
}

// Subscribe implements IClient.
func (r *RoutingClient) Subscribe() async.Sequence[Cemi] {
	r.subMux.Lock()
	defer r.subMux.Unlock()
	sub := make(async.Sequence[Cemi])
	r.subs[sub] = true
	return sub
}

// Unsubscribe implements IClient.
func (r *RoutingClient) Unsubscribe(sub async.Sequence[Cemi]) {
	r.subMux.Lock()
	defer r.subMux.Unlock()
	delete(r.subs, sub)
}

// ------------------------------------------------------------------------------------
// Private
// ------------------------------------------------------------------------------------

func readUdp(mux *sync.RWMutex, conn *net.UDPConn) async.Sequence[[]byte] {
	ch := make(async.Sequence[[]byte])
	go func() {
		defer close(ch)
		for {
			buffer := make([]byte, 1024)
			mux.RLock()
			n, err := conn.Read(buffer)
			mux.RUnlock()
			if err != nil {
				ch <- async.Fail[[]byte](NewErrNetConnection("failed to read from udp - %v", err))
				return
			}
			ch <- async.Success(buffer[:n])
		}
	}()
	return ch
}

func (k *RoutingClient) notify(p Cemi) {
	k.subMux.RLock()
	defer k.subMux.RUnlock()
	for sub := range k.subs {
		sub <- async.Success(p)
	}
}

func (k *RoutingClient) run(ctx context.Context, iface *net.Interface) error {
	// Create UDP address
	addr, err := net.ResolveUDPAddr("udp4", k.multicastAddr)
	if err != nil {
		return NewErrAddressResolution("failed to create udp address for %s - %v", k.multicastAddr, err)
	}
	// Connect to multicast address
	k.conn, err = net.ListenMulticastUDP("udp4", nil, addr)
	if err != nil {
		return NewErrNetConnection("failed to listen on %s - %v", addr.String(), err)
	}
	defer func() {
		k.conn.Close()
		k.connMux.Lock()
		k.conn = nil
		k.connMux.Unlock()
	}()
	// Create Reader
	reader := readUdp(&k.connMux, k.conn)
	slog.Debug("KnxRoutingClient listening...", slog.String("interface", iface.Name), slog.String("address", k.multicastAddr))
	// Run
	for {
		select {
		case <-ctx.Done():
			return nil
		case data := <-reader:
			// Check for error
			if data.Error != nil {
				return data.Error
			}
			// Parse Header
			header, _, err := parseKnxNetIpHeader(data.Value[:6])
			if err != nil {
				slog.Warn("failed to parse net header", slog.Any("error", err))
				continue
			}
			// Process Package depending on ServiceType
			switch header.ServiceType {
			case SEARCH_REQUEST:
				err = k.processSearchReq(data.Value)
			case SEARCH_RESPONSE:
				err = k.processSearchRes(data.Value)
			case DESCRIPTION_REQUEST:
				err = k.processDescriptionReq(data.Value)
			case DESCRIPTION_RESPONSE:
				err = k.processDescriptionResp(data.Value)
			case ROUTING_INDICATION:
				err = k.processRoutingIndication(data.Value)
			case ROUTING_LOST_MESSAGE:
				err = k.processRoutingLostMsg(data.Value)
			case ROUTING_BUSY:
				err = k.processRoutingBusy(data.Value)
			default:
				continue // ignore
			}
			if err != nil {
				slog.Warn("failed to process request", slog.Any("error", err))
				continue
			}
		}
	}
}

func (r *RoutingClient) processSearchReq(data []byte) error {
	return NewErrNotImplemented("not implemented")
}

func (r *RoutingClient) processSearchRes(data []byte) error {
	return NewErrNotImplemented("not implemented")
}

func (r *RoutingClient) processDescriptionReq(data []byte) error {
	return NewErrNotImplemented("not implemented")
}

func (r *RoutingClient) processDescriptionResp(data []byte) error {
	return NewErrNotImplemented("not implemented")
}

func (r *RoutingClient) processRoutingIndication(data []byte) error {
	msg, err := parseKnxNetIpPackage[*RoutingIndication](data)
	if err != nil {
		return err
	}
	r.notify(msg.Body.Cemi)
	return nil
}

func (r *RoutingClient) processRoutingLostMsg(data []byte) error {
	return NewErrNotImplemented("not implemented")
}

func (r *RoutingClient) processRoutingBusy(data []byte) error {
	return NewErrNotImplemented("not implemented")
}

// ------------------------------------------------------------------------------------
// Options
// ------------------------------------------------------------------------------------

func WithRoutingClientMulticastAddr(addr string) func(*RoutingClient) {
	return func(rc *RoutingClient) {
		rc.multicastAddr = addr
	}
}

func WithRoutingClientRetryInterval(interval time.Duration) func(*RoutingClient) {
	return func(rc *RoutingClient) {
		rc.retryInterval = interval
	}
}

// ------------------------------------------------------------------------------------
// Constructor
// ------------------------------------------------------------------------------------

func NewRoutingClient(ctx context.Context, iface *net.Interface, opts ...func(*RoutingClient)) IClient {
	c := &RoutingClient{
		multicastAddr: "224.0.23.12:3671",
		retryInterval: 30 * time.Second,
		conn:          nil,
		subs:          map[async.Sequence[Cemi]]bool{},
	}
	for _, o := range opts {
		o(c)
	}
	go func() {
		for {
			err := c.run(ctx, iface)
			if err == nil {
				slog.Debug("RoutingClient shutdown successfully - context ended")
				return // Context dead
			}
			slog.Error("KNX RoutingClient failed", slog.Any("error", err))
			time.Sleep(c.retryInterval)
		}
	}()
	return c
}
